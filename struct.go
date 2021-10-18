package main

type Atom struct {
	ID         int
	Author     string
	Createtime string
	Edittime   string
	Title      string
	Tags       []string
	Hide       bool
	Contents   []Content
}

type Content struct {
	Type string `json:"type" bson:type"`
	Body string `json:"body" bson:body"`
}
