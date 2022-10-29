package server

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

type TestResponse struct {
	OK bool `json:"ok"`
}

func getMocRoute(t *testing.T) *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := json.Marshal(&TestResponse{OK: true})
		if err != nil {
			t.Fatal(err)
		}
		w.Write(body)
	}).Methods("GET")

	return route
}

func TestNewServer(t *testing.T) {
	ok := make(chan struct{})

	t.Run("Start and stop server", func(t *testing.T) {

		srv, err := NewServer(
			nil,
			"",
			6004,
			func() {
				ok <- struct{}{}
			},
			getMocRoute(t),
		)

		if err != nil {
			t.Fatal(err)
			return
		}

		for {
			if srv.server != nil {
				break
			}

			time.Sleep(time.Second * 1)
		}

		if err := srv.RunServer(nil, "", 6005, nil); err == nil {
			t.Fatal("Double run server")
			return
		}

		srv.Stop()

		<-ok
	})

	// TODO: add test with http client
	// t.Run("Http server test", func(t *testing.T) {
	// })
}
