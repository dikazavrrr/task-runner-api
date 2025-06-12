package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-runner-api/internal/task"
	"task-runner-api/pkg/logger"

	"github.com/gorilla/mux"
)

type Handler struct {
	Manager *task.Manager
}

func NewHandler(m *task.Manager) *Handler {
	return &Handler{Manager: m}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", h.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	t := h.Manager.CreateTask()
	logger.Info(fmt.Sprintf("CreateTask: new task created with ID=%s, status=%s", t.ID, t.Status))

	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if task, ok := h.Manager.GetTask(id); ok {
		logger.Info(fmt.Sprintf("GetTask: task found with ID=%s, status=%s", task.ID, task.Status))
		writeJSON(w, http.StatusOK, task)
	} else {
		logger.Error(fmt.Sprintf("GetTask: task with ID=%s not found", id))
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, ok := h.Manager.GetTask(id)
	if !ok {
		logger.Error(fmt.Sprintf("DeleteTask Error: %d (id: %s)", http.StatusNotFound, id))
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	h.Manager.DeleteTask(id)
	logger.Info(fmt.Sprintf("DeleteTask Status: %d (id: %s)", http.StatusOK, id))
	writeJSON(w, http.StatusOK, map[string]string{"message": "task deleted"})
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
