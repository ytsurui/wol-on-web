package health

import (
	"encoding/json"
	"net/http"
)

func HealthResponse(w http.ResponseWriter, r *http.Request) {
	respData := make(map[string]string)
	respData["health"] = "ok"

	res, _ := json.Marshal(respData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}
