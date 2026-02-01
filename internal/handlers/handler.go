package handlers

import (
	"candipack-pdf/internal/generator"
	"candipack-pdf/internal/parser"
)

type Handler struct {
	parser    *parser.HTMLParser
	generator *generator.Generator
}

func New() *Handler {
	return &Handler{
		parser:    parser.NewHTMLParser(),
		generator: generator.NewGenerator(),
	}
}
