package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

		slog.Info("request", "Headers", fmt.Sprintf("+%v", r.Header), "Body", string(body))

		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":3001", mux); err != nil {
		slog.Error("http listen and server error", "error", err.Error())
	}
}
