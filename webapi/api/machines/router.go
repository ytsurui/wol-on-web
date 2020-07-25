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

func GetMachineConfig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "HEAD":
	case "OPTIONS":
	case "GET":
		data, err := downloadConfig()
		if (data == nil) && (err == nil) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			api.LogPrint(r, http.StatusMethodNotAllowed)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			api.LogPrint(r, http.StatusInternalServerError)
		} else {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			api.LogPrint(r, http.StatusOK)
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	api.LogPrint(r, http.StatusMethodNotAllowed)
}

func CheckConfigFlag(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "HEAD":
	case "OPTIONS":
	case "GET":
		status := checkConfigFlag(w, r)
		api.LogPrint(r, status)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	api.LogPrint(r, http.StatusMethodNotAllowed)
}
