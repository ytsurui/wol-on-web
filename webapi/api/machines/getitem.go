package machines

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getItem(w http.ResponseWriter, r *http.Request) int {

	params := r.URL.Query()

	if _, ok := params["id"]; ok {
		machineID, err := strconv.Atoi(params["id"][0])
		if err == nil {
			data, err := getMachineItemData(machineID)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return (http.StatusNotFound)
			}

			respDataByte, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return (http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(respDataByte)
			return (http.StatusOK)
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	return (http.StatusBadRequest)
}
