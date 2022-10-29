package rest

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"prep/internal/storage"

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
	writeRequest(w, "Template responce")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	if err == nil {
		logrus.Infof("Get data %s", string(data))
	} else {
		logrus.Error("Failed read body")
		writeRequest(w, errors.New("Failed read body"))
	}

	if err := storage.SendPrepData(data); err != nil {
		logrus.Errorf("Invalid send data to store %s", err)
		writeRequest(w, errors.New("Internal server error"))
	}

	writeRequest(w, "Template responce")
}
