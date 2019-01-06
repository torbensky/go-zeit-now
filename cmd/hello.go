package main

import (
	"fmt"
	"net/http"

	"github.com/common-nighthawk/go-figure"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	myFigure := figure.NewFigure("Hello World", "", true)
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<pre>%s</pre>", myFigure.String())
	fmt.Fprintf(w, "<br /><h3><a href='/bye/bye'>Leave</a>")
}
