/*
Copyright © 2021 Aurelio Calegari

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package plist

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/subchen/go-xmldom"
)

type pltype int

const (
	Bool pltype = iota
	Dict
	String
	Array
	Real
	Integer
	Date
	Data
)

type Target int

const (
	Json Target = iota
	Yaml
	Html
)

type Config struct {
	Target       Target
	HighFidelity bool
	Beatify      bool
}

func Parse(in []byte, config *Config) ([]byte, error) {
	doc := xmldom.Must(xmldom.ParseXML(string(in)))
	root := doc.Root
	var output string
	if config.HighFidelity {
		output = toJsonFill(root.FirstChild(), nil)
	} else {
		itm, _ := makePLTypeFromNode(root.FirstChild())
		output = itm.toJson(nil)
	}
	bout := []byte(output)
	var err error
	switch config.Target {
	case Json:
		if config.Beatify {
			var prettyJSON bytes.Buffer
			err = json.Indent(&prettyJSON, bout, "", "  ")
			if err != nil {
				return nil, err
			}
			bout = prettyJSON.Bytes()
		}
		return bout, nil
	case Yaml:
		bout, err = yaml.JSONToYAML(bout)
		if err != nil {
			return nil, err
		}
		return bout, err
	}
	return nil, fmt.Errorf("not implemented type conversion")
}

type plnode struct {
	name   string
	pltype pltype
	value  interface{}
}

func toJsonFill(n, parent *xmldom.Node) string {
	v := ""
	if n == nil {
		return v
	}
	switch n.Name {
	case "key":
		v += fmt.Sprintf("{\"k\":\"%s\",", n.Text)
		next := n.NextSibling()
		v += fmt.Sprintf("\"v\":%s}", toJsonFill(next, nil))
		return v
	case "dict":
		v += fmt.Sprintf("{\"type\":\"dict\",\"value\":[")
		var arr []string
		for i, cn := range n.Children {
			if i%2 == 0 {
				arr = append(arr, toJsonFill(cn, n))
			}
		}
		v += strings.Join(arr, ",") + "]}"
		return v
	case "true", "false":
		v += fmt.Sprintf("{\"type\":\"bool\",\"value\":\"%s\"}", n.Name)
		return v
	case "string", "real", "integer", "date":
		v += fmt.Sprintf("{\"type\":\"%s\",\"value\":\"%s\"}", n.Name, n.Text)
		return v
	case "data":
		v += fmt.Sprintf("{\"type\":\"%s\",\"value\":\"%s\"}", n.Name, trimData(n.Text))
		return v
	case "array":
		v += fmt.Sprintf("{\"type\":\"array\",\"value\":[")
		var arr []string
		for _, cn := range n.Children {
			itm := toJsonFill(cn, n)
			arr = append(arr, fmt.Sprintf("%s", itm))
		}
		v += strings.Join(arr, ",") + "]}"
		//if parent == nil {
		//	v = fmt.Sprintf("{%s}", v)
		//}
		return v
	}
	return ""
}

func trimData(data interface{}) string {
	s := bufio.NewScanner(strings.NewReader(
		fmt.Sprintf("\"%v\"", data)))
	dv := ""
	for s.Scan() {
		dv = dv + strings.TrimSpace(s.Text())
	}
	return fmt.Sprintf("%s", dv)
}

func (node *plnode) toJson(parent *plnode) string {
	v := ""
	switch node.pltype {
	case Data:
		if parent == nil || parent.pltype != Array {
			v = v + fmt.Sprintf("\"%s\":", node.name)
		}
		v = v + trimData(node.value)
	case String, Date:
		if parent == nil || parent.pltype != Array {
			v = v + fmt.Sprintf("\"%s\":", node.name)
		}
		v = v + fmt.Sprintf("\"%v\"", node.value)
	case Bool, Real, Integer:
		if parent == nil || parent.pltype != Array {
			v = v + fmt.Sprintf("\"%s\":", node.name)
		}
		v = v + fmt.Sprintf("%v", node.value)
	case Array:
		if parent != nil && parent.pltype != Array {
			v = v + fmt.Sprintf("\"%s\":", node.name)
		}
		v = v + "["
		vv := node.value.([]*plnode)
		for i, vi := range vv {
			v = v + vi.toJson(node)
			if i < len(vv)-1 {
				v = v + ","
			}
		}
		v = v + "]"
	case Dict:
		if node.name != "" {
			v = v + fmt.Sprintf("\"%s\":", node.name)
		}
		v = v + "{"
		vv := node.value.([]*plnode)
		for i, vi := range vv {
			v = v + vi.toJson(node)
			if i < len(vv)-1 {
				v = v + ","
			}
		}
		v = v + "}"
	}
	return v
}

func makePLTypeFromNode(n *xmldom.Node) (*plnode, *xmldom.Node) {
	switch n.Name {
	case "key":
		next := n.NextSibling()
		tp, _ := makePLTypeFromNode(next)
		tp.name = n.Text
		return makePLType(n.Text, tp.pltype, tp.value), next.NextSibling()
	case "dict":
		var dArr []*plnode
		for i, cn := range n.Children {
			if i%2 == 0 {
				currChild := cn
				itm, _ := makePLTypeFromNode(currChild)
				dArr = append(dArr, itm)
			}
		}
		return makePLType("", Dict, dArr), n.NextSibling()
	case "true":
		return makePLType("Bool", Bool, true), nil
	case "false":
		return makePLType("Bool", Bool, false), nil
	case "string":
		return makePLType("String", String, n.Text), nil
	case "real":
		v, err := strconv.ParseFloat(n.Text, 64)
		if err != nil {
			panic(err)
		}
		return makePLType("Real", Real, v), nil
	case "integer":
		v, err := strconv.ParseInt(n.Text, 10, 64)
		if err != nil {
			panic(err)
		}
		return makePLType("Integer", Integer, v), nil
	case "date":
		return makePLType("Date", Date, n.Text), nil
	case "data":
		return makePLType("Data", Data, n.Text), nil
	case "array":
		var dArr []*plnode
		for _, cn := range n.Children {
			currChild := cn
			itm, _ := makePLTypeFromNode(currChild)
			dArr = append(dArr, itm)
		}
		return makePLType("array", Array, dArr), n.NextSibling()
	}
	return nil, nil
}

func makePLType(name string, pltype pltype, value interface{}) *plnode {
	return &plnode{
		name:   name,
		pltype: pltype,
		value:  value,
	}
}
