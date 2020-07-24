package api

import (
	"fmt"
	"net/http"
	"time"
)

func LogPrint(r *http.Request, respCode int) {
	nowTimeStr := time.Now().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s] \"%s %s %s\" %d %s", r.RemoteAddr, nowTimeStr, r.Method, r.RequestURI, r.Proto, respCode, r.UserAgent())
	fmt.Println(msg)
}
