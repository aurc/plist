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
	"io/ioutil"
	"os"
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
)

// Config sets the available conversions parameters such as:
// - Target: Either Json or Yaml;
// - HighFidelity: Generate a natural format or a more verbose format to explicitly identify each field datatype
// (default is false).
// - Beatify: Only available for Json, as by default the produced Json is in a minified format.
type Config struct {
	Target       Target
	HighFidelity bool
	Beatify      bool
}

// Convert a given plist file as the `in` input and the Config into either a Json or Yaml format.
func Convert(in []byte, config *Config) (bout []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	doc := xmldom.Must(xmldom.ParseXML(string(in)))
	root := doc.Root
	var output string
	if config.HighFidelity {
		output = toJsonFill(root.FirstChild())
	} else {
		itm, _ := makePLTypeFromNode(root.FirstChild())
		output = itm.toJson(nil)
	}
	bout = []byte(output)
	switch config.Target {
	case Json:
		if config.Beatify {
			var prettyJSON bytes.Buffer
			_ = json.Indent(&prettyJSON, bout, "", "  ")
			if prettyJSON.Len() > 0 {
				bout = prettyJSON.Bytes()
			}
		}
	case Yaml:
		bout, err = yaml.JSONToYAML(bout)
	}
	return bout, err
}

// ReadInput will first check if there's any incoming content from
// the Standard Input (via pipe) otherwise it assumes the input parameter
// contains a valid input file name. This is file format agnostic.
func ReadInput(input string) ([]byte, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		if input == "" {
			return nil, fmt.Errorf("bad input")
		} else {
			inputPayload, err := ioutil.ReadFile(input)
			if err != nil {
				return nil, err
			}
			return inputPayload, nil
		}
	}

	reader := bufio.NewReader(os.Stdin)

	inputPayload, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return inputPayload, nil
}

type plnode struct {
	name   string
	pltype pltype
	value  interface{}
}

func toJsonFill(n *xmldom.Node) string {
	v := ""
	switch n.Name {
	case "key":
		v += fmt.Sprintf("{\"k\":\"%s\",", n.Text)
		next := n.NextSibling()
		v += fmt.Sprintf("\"v\":%s}", toJsonFill(next))

	case "dict":
		v += fmt.Sprintf("{\"type\":\"dict\",\"value\":[")
		var arr []string
		for i, cn := range n.Children {
			if i%2 == 0 {
				arr = append(arr, toJsonFill(cn))
			}
		}
		v += strings.Join(arr, ",") + "]}"
	case "true", "false":
		v += fmt.Sprintf("{\"type\":\"bool\",\"value\":\"%s\"}", n.Name)
	case "string", "real", "integer", "date":
		v += fmt.Sprintf("{\"type\":\"%s\",\"value\":\"%s\"}", n.Name, n.Text)
	case "data":
		v += fmt.Sprintf("{\"type\":\"%s\",\"value\":%s}", n.Name, trimData(n.Text))
	case "array":
		v += fmt.Sprintf("{\"type\":\"array\",\"value\":[")
		var arr []string
		for _, cn := range n.Children {
			itm := toJsonFill(cn)
			arr = append(arr, fmt.Sprintf("%s", itm))
		}
		v += strings.Join(arr, ",") + "]}"
	}
	return v
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
	var pl *plnode
	var ns *xmldom.Node
	switch n.Name {
	case "key":
		next := n.NextSibling()
		tp, _ := makePLTypeFromNode(next)
		tp.name = n.Text
		pl, ns = makePLType(n.Text, tp.pltype, tp.value), next.NextSibling()
	case "dict":
		var dArr []*plnode
		for i, cn := range n.Children {
			if i%2 == 0 {
				currChild := cn
				itm, _ := makePLTypeFromNode(currChild)
				dArr = append(dArr, itm)
			}
		}
		pl, ns = makePLType("", Dict, dArr), n.NextSibling()
	case "true":
		pl, ns = makePLType("Bool", Bool, true), nil
	case "false":
		pl, ns = makePLType("Bool", Bool, false), nil
	case "string":
		pl, ns = makePLType("String", String, n.Text), nil
	case "real":
		v, err := strconv.ParseFloat(n.Text, 64)
		if err != nil {
			panic(err)
		}
		pl, ns = makePLType("Real", Real, v), nil
	case "integer":
		v, err := strconv.ParseInt(n.Text, 10, 64)
		if err != nil {
			panic(err)
		}
		pl, ns = makePLType("Integer", Integer, v), nil
	case "date":
		pl, ns = makePLType("Date", Date, n.Text), nil
	case "data":
		pl, ns = makePLType("Data", Data, n.Text), nil
	case "array":
		var dArr []*plnode
		for _, cn := range n.Children {
			currChild := cn
			itm, _ := makePLTypeFromNode(currChild)
			dArr = append(dArr, itm)
		}
		pl, ns = makePLType("array", Array, dArr), n.NextSibling()
	}
	return pl, ns
}

func makePLType(name string, pltype pltype, value interface{}) *plnode {
	return &plnode{
		name:   name,
		pltype: pltype,
		value:  value,
	}
}
