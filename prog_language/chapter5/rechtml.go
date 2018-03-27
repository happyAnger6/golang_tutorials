package main

import "golang.org/x/net/html"

type NodeType int32

const (
	ErrNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Node struct {
	Type NodeType
	Data string
	Attr []html.Attribute
	FirstChild, NextSibling *Node
}
