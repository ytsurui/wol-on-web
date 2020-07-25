package machines

import (
	"encoding/json"
	"net/http"
)

type MachineInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	MacAddr string `json:"macaddr`
	IpAddr  string `json:"ipaddr"`
	NetMask int    `json:"netmask`
	NetAddr string `json:"netaddr"`
}

func getMachineList(w http.ResponseWriter, r *http.Request) int {
	machineList := getMachineListData()
	respDataByte, err := json.Marshal(machineList)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return (http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(respDataByte)
	return (http.StatusOK)
}
