package main

import "net/http"

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	ResponseWithJson(w, 200, struct{}{})
}
