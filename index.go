package main

import (
	"sync"
	"html/template"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}