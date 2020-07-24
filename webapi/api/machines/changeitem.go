package machines

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func changeItem(w http.ResponseWriter, r *http.Request) int {
	params := r.URL.Query()

	machineID := 0
	if _, ok := params["id"]; ok {
		var err error
		machineID, err = strconv.Atoi(params["id"][0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return (http.StatusBadRequest)
		}
		fmt.Println("machineID:", machineID)
	}

	jsonbytes, err := getJsonData(r)
	if err != nil {
		fmt.Println("JSON data not found, err=", err)
		w.WriteHeader(http.StatusBadRequest)
		return (http.StatusBadRequest)
	}

	var changeMachineInfo MachineInfo
	err = json.Unmarshal(jsonbytes, &changeMachineInfo)
	if err != nil {
		fmt.Println("JSON decode error, err=", err)
		w.WriteHeader(http.StatusInternalServerError)
		return (http.StatusInternalServerError)
	}

	if machineID != 0 {
		changeMachineInfo.ID = machineID
	}

	// Scan Data
	if len(changeMachineInfo.Name) == 0 || len(changeMachineInfo.MacAddr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return (http.StatusBadRequest)
	}

	err = writeMachineItemData(machineID, changeMachineInfo)
	if err != nil {
		fmt.Println("[api.machines.changeItem] Write Config Error, err=", err)
		w.WriteHeader(http.StatusInternalServerError)
		return (http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	return (http.StatusOK)
}

func getJsonData(r *http.Request) (data []byte, err error) {
	headerContentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(headerContentType, "application/json") == false {
		data = nil
		err = errors.New("Invalid content type")
		return
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		data = nil
		err = errors.New("Invalid Content Length")
		return
	}

	data = make([]byte, 0)
	receiveCount := 0

	for {
		respBuf := make([]byte, length)
		respLength, err2 := r.Body.Read(respBuf)

		if err2 != nil && err2 != io.EOF {
			data = nil
			err = errors.New("Data Receiving Error")
			return
		}

		data = append(data, respBuf[:respLength]...)
		receiveCount += respLength

		if receiveCount >= length {
			break
		}
	}

	err = nil
	return
}
