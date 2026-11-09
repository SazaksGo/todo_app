package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SazaksGo/todo_app/task"
	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *task.List
}

func NewHTTPHandlers(todoList *task.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}

}

func (h *HTTPHandlers) HandlerCreateTasks(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Massage: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Massage: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.Time.String(), http.StatusBadRequest)
	}
	todoTask := task.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {

		errDTO := ErrorDTO{
			Massage: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, task.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
		return

	}
}

func (h *HTTPHandlers) HandlerGetTasks(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	taskk, err := h.todoList.GetTask(title)
	if err != nil {
		errDTO := ErrorDTO{
			Massage: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, task.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(taskk, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
		return
	}

}

func (h *HTTPHandlers) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
		return
	}

}

func (h *HTTPHandlers) HandlerGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.ListNotCompletedTasks()
	b, err := json.MarshalIndent(uncompletedTasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
		return
	}

}

func (h *HTTPHandlers) HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeDTO CompleteTaskDTO

	if err := json.NewDecoder(r.Body).Decode(&completeDTO); err != nil {
		errDTO := ErrorDTO{
			Massage: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	title := mux.Vars(r)["title"]

	var (
		changedTask task.Task
		err         error
	)
	if completeDTO.Complete {
		changedTask, err = h.todoList.CompleteTask(title)
	} else {
		changedTask, err = h.todoList.UnCompleteTask(title)
	}

	if err != nil {
		errDTO := ErrorDTO{
			Massage: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, task.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(changedTask, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
	}

}

func (h *HTTPHandlers) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		errDTO := ErrorDTO{
			Massage: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, task.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return

	}
	w.WriteHeader(http.StatusNoContent)
}
