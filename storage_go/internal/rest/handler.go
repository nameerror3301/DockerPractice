package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"store/internal/storage"

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
		logrus.Infof("Get data for store %s", string(data))
	} else {
		logrus.Error("Failed read body %v", err)
		writeRequest(w, errors.New("failed read body"))
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()

	if err := storage.SavePayload(&ctx, data); err != nil {
		logrus.Errorf("Failed save data in database %v", err)
		writeRequest(w, errors.New("internal server error"))
		return
	}

	writeRequest(w, true)
}
