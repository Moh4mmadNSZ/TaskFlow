package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"taskflow/models"
)

type TaskHandler struct {
	tasks  map[int]models.Task
	nextID int
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (h *TaskHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/tasks", h.getTasks)
	mux.HandleFunc("POST /api/v1/tasks", h.createTask)
	mux.HandleFunc("GET /api/v1/tasks/{id}", h.getTask)
	mux.HandleFunc("PUT /api/v1/tasks/{id}", h.updateTask)
	mux.HandleFunc("DELETE /api/v1/tasks/{id}", h.deleteTask)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}

func validateTask(task models.Task) string {
	if strings.TrimSpace(task.Title) == "" {
		return "Title cannot be empty"
	}
	return ""
}

func (h *TaskHandler) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := make([]models.Task, 0, len(h.tasks))

	for _, task := range h.tasks {
		tasks = append(tasks, task)
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if message := validateTask(task); message != "" {
		writeError(w, http.StatusBadRequest, message)
		return
	}

	task.ID = h.nextID
	h.nextID++
	h.tasks[task.ID] = task

	writeJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) getTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, exists := h.tasks[id]
	if !exists {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	if _, exists := h.tasks[id]; !exists {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}

	var updated models.Task
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if message := validateTask(updated); message != "" {
		writeError(w, http.StatusBadRequest, message)
		return
	}

	updated.ID = id
	h.tasks[id] = updated

	writeJSON(w, http.StatusOK, updated)
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	if _, exists := h.tasks[id]; !exists {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}

	delete(h.tasks, id)
	w.WriteHeader(http.StatusNoContent)
}
