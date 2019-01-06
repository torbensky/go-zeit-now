package main

import (
	"fmt"
	"net/http"

	figure "github.com/common-nighthawk/go-figure"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	myFigure := figure.NewFigure("Byeee", "", true)

	fmt.Fprintf(w, myFigure.String())
}
