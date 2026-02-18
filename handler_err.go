package main

import "net/http"

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	ResponseWithError(w, 400, "something went wrong")
}
