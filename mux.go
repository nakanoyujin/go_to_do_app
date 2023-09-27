package main

import (
	"net/http"
)

// どんなハンドラーの実装をどんなURLで公開するかルーティングするかを担当
func NewMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-TYpe", "application/json;charset=utf-8")

		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	return mux
}
