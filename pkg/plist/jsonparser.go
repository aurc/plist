package plist

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

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

func Parse(in []byte) (string, error) {
	doc := xmldom.Must(xmldom.ParseXML(string(in)))
	root := doc.Root
	var dArr []*plnode
	for _, cn := range root.Children {
		itm, _ := makePLTypeFromNode(cn)
		dArr = append(dArr, itm)
	}
	return dArr[0].toJson(nil), nil
}

type plnode struct {
	Name  string
	Type  pltype
	Value interface{}
}

func (node *plnode) toJson(parent *plnode) string {
	v := ""
	switch node.Type {
	case Data:
		if parent == nil || parent.Type != Array {
			v = v + fmt.Sprintf("\"%s\":", node.Name)
		}

		s := bufio.NewScanner(strings.NewReader(
			fmt.Sprintf("\"%v\"", node.Value)))
		dv := ""
		for s.Scan() {
			dv = dv + strings.TrimSpace(s.Text())
		}
		v = v + fmt.Sprintf("%v", dv)
	case String, Date:
		if parent == nil || parent.Type != Array {
			v = v + fmt.Sprintf("\"%s\":", node.Name)
		}
		v = v + fmt.Sprintf("\"%v\"", node.Value)
	case Bool, Real, Integer:
		if parent == nil || parent.Type != Array {
			v = v + fmt.Sprintf("\"%s\":", node.Name)
		}
		v = v + fmt.Sprintf("%v", node.Value)
	case Array:
		if parent != nil && parent.Type != Array {
			v = v + fmt.Sprintf("\"%s\":", node.Name)
		}
		v = v + "["
		vv := node.Value.([]*plnode)
		for i, vi := range vv {
			v = v + vi.toJson(node)
			if i < len(vv)-1 {
				v = v + ","
			}
		}
		v = v + "]"
	case Dict:
		if node.Name != "" {
			v = v + fmt.Sprintf("\"%s\":", node.Name)
		}
		v = v + "{"
		vv := node.Value.([]*plnode)
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
		//if n.Text == "all_tasks" {
		// fmt.Sprintf("I did!")
		//}
		tp.Name = n.Text
		return makePLType(n.Text, tp.Type, tp.Value), next.NextSibling()
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
		Name:  name,
		Type:  pltype,
		Value: value,
	}
}
