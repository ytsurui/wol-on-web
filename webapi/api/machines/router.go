package machines

import (
	"net/http"

	"../../api"
)

func GetMachineList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "HEAD":
	case "OPTIONS":
	case "GET":
		statusCode := getMachineList(w, r)
		api.LogPrint(r, statusCode)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	api.LogPrint(r, http.StatusMethodNotAllowed)
}

func GetMachineItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "HEAD":
	case "OPTIONS":
	case "GET":
		statusCode := getItem(w, r)
		api.LogPrint(r, statusCode)
		return
	case "POST":
		statusCode := changeItem(w, r)
		api.LogPrint(r, statusCode)
		return
	case "DELETE":
		statusCode := deleteItem(w, r)
		api.LogPrint(r, statusCode)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	api.LogPrint(r, http.StatusMethodNotAllowed)
}
