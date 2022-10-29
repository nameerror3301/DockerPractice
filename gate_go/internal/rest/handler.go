package rest

import (
	"encoding/json"
	"fmt"
	"gateway/internal/prepare"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type MResponse struct {
	OK     bool        `json:"ok"`
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

func writeRequest(w http.ResponseWriter, data interface{}) {
	var body []byte

	if val, ok := data.(error); ok {
		body, _ = json.Marshal(&MResponse{
			OK:    false,
			Error: val.Error(),
		})

	} else {
		body, _ = json.Marshal(&MResponse{
			OK:     true,
			Result: data,
		})

	}

	w.Write(body)
}

func tmpHandler(w http.ResponseWriter, r *http.Request) {
	writeRequest(w, "Check services")
}

func setData(w http.ResponseWriter, r *http.Request) {
	prepData, err := io.ReadAll(r.Body)
	if err != nil {
		writeRequest(w, fmt.Errorf("invalid read body fir prep %v", err))
		return
	}

	logrus.Warnf("Im get data for forward on prepare:\n%s", string(prepData))

	prepare.SendPrepData(prepData)
	writeRequest(w, "Data ready")
}
