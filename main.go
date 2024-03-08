package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelInfo, TimeFormat: time.Kitchen}))

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

		logger.Info("request", "Headers", fmt.Sprintf("+%v", r.Header), "Body", string(body))

		w.WriteHeader(http.StatusOK)
	})

	logger.Info("start server")
	if err := http.ListenAndServe(":3001", mux); err != nil {
		logger.Error("http listen and server error", "error", err.Error())
	}
}
