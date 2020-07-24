package machines

import (
	"fmt"
	"net/http"
	"strconv"
)

func deleteItem(w http.ResponseWriter, r *http.Request) int {
	params := r.URL.Query()

	if _, ok := params["id"]; ok == false {
		w.WriteHeader(http.StatusBadRequest)
		return (http.StatusBadRequest)
	}

	machineID, err := strconv.Atoi(params["id"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return (http.StatusBadRequest)
	}

	err = deleteMachineItemData(machineID)
	if err != nil {
		fmt.Println("[api.machines.deleteItem] Delete Machine Error, err=", err)
		w.WriteHeader(http.StatusInternalServerError)
		return (http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	return (http.StatusOK)
}
