package controllers

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is your local web server!")
}
