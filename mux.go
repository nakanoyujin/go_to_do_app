package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/nakanoyujin/go_to_do_app/handler"
	"github.com/nakanoyujin/go_to_do_app/store"
)

// どんなハンドラーの実装をどんなURLで公開するかルーティングするかを担当
// NewServerMuxはメソッドの違いやパラメータの違いを解釈しないので別のに変える
// func NewMux() http.Handler {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-TYpe", "application/json;charset=utf-8")

// 		_, _ = w.Write([]byte(`{"status":"ok"}`))
// 	})
// 	return mux
// }

func NewMux() http.Handler {
	mux := chi.NewRouter()
	// chiはhttp.Handlerインターフェースを実装しているので同じような使い方ができる。
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-TYpe", "application/json;charset=utf-8")

		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	v := validator.New()
	at := &handler.AddTask{Store: store.Tasks, Validator: v}

	// mux.method は第一引数と関数を内部的にHandleFuncとして処理する。
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux
}
