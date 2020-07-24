package wol

import (
	"fmt"
	"net"
	"net/http"

	"github.com/mdlayher/wol"
)

func sendWolPacket(w http.ResponseWriter, r *http.Request) int {
	params := r.URL.Query()

	if macaddr, ok := params["macaddr"]; ok {
		//fmt.Println("[api.wol.sendWolPacket] MAC Addr=" + macaddr[0])
		var broadcastAddr string
		if netaddr, ok := params["broadcast"]; ok {
			broadcastAddr = netaddr[0] + ":7"
		} else {
			broadcastAddr = "255.255.255.255:7"
		}

		statusCode, err := sendWol(macaddr[0], broadcastAddr)

		if err != nil {
			fmt.Println("[api.wol.sendWolPacket] Error, err=", err)
			w.WriteHeader(statusCode)
			return (statusCode)
		}

		w.WriteHeader(http.StatusOK)
		return (http.StatusOK)
	}

	w.WriteHeader(http.StatusBadRequest)

	return (http.StatusBadRequest)
}

func sendWol(macaddr string, networkAddr string) (respCode int, respErr error) {
	targetAddr, err := net.ParseMAC(macaddr)
	if err != nil {
		respErr = err
		respCode = http.StatusBadRequest
		return
	}

	wolClient, err := wol.NewClient()
	if err != nil {
		respErr = err
		respCode = http.StatusInternalServerError
		return
	}

	err = wolClient.Wake(networkAddr, targetAddr)
	if err != nil {
		respErr = err
		respCode = http.StatusInternalServerError
		return
	}

	respErr = nil
	respCode = http.StatusOK
	return
}
