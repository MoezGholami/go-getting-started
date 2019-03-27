package main

import (
    "fmt"
    "reflect"
    "runtime/debug"
)

type node struct {
    data int
    next *nNode
    prev *pNode
}

type nNode node
type pNode node

func (n *node) doStuff() {
    debug.PrintStack()
}
func (n *nNode) doStuff() {
    debug.PrintStack()
}
func (n *pNode) doStuff() {
    debug.PrintStack()
}

func makeNode() node {
    result := node{data:0}
    result.next = (*nNode)(&result)
    result.prev = (*pNode)(&result)
    return result
}

func main() {
    n := makeNode()
    n.prev.next.doStuff()
    fmt.Println(reflect.TypeOf(n.prev).String())
}
