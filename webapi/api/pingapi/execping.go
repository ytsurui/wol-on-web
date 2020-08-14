package pingapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sparrc/go-ping"
)

func execPing(w http.ResponseWriter, r *http.Request) int {
	params := r.URL.Query()

	if addrArray, ok := params["ipaddr"]; ok {
		if len(addrArray) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return (http.StatusBadRequest)
		}
		addr := addrArray[0]
		pinger, err := ping.NewPinger(addr)
		if err != nil {
			fmt.Println("[api.pingapi.execPing] Ping Object Error, err=", err)
			w.WriteHeader(http.StatusInternalServerError)
			return (http.StatusInternalServerError)
		}
		pinger.SetPrivileged(true)
		pinger.Timeout = time.Second * 3
		pinger.Count = 1
		pinger.Run()

		stats := pinger.Statistics()

		if stats.PacketsRecv == 0 {
			w.WriteHeader(http.StatusNotFound)
			return (http.StatusNotFound)
		}

		respMap := map[string]string{"ipaddr": addr}
		respData, err := json.Marshal(respMap)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return (http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(respData)
		return (http.StatusOK)
	}
	w.WriteHeader(http.StatusBadRequest)
	return (http.StatusBadRequest)
}
