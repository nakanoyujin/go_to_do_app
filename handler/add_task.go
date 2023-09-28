package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/nakanoyujin/go_to_do_app/entity"
	"github.com/nakanoyujin/go_to_do_app/store"
)

type AddTask struct {
	store     *store.TaskStore
	validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return

	}
	err := at.validator.Struct(b)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t := &entity.Task{
		Title:   b.Title,
		Status:  entity.TaskStatusTodo,
		Created: time.Now(),
	}

	id, err := store.Tasks.Add(t)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID int `json:"id"`
	}{ID: int(id)}

	RespondJSON(ctx, w, rsp, http.StatusOK)
}
