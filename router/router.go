package router

import (
)

var basePaseFiles []string = []string{
	"templates/base.html",
	"templates/header.html",
}

type Router struct {
}

func NewRouter() *Router {
	return &Router{
	}
}

