// Package main contém a API Rest de teste da ilhasoft para manipular todo lists
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Todo representa uma tarefa
type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// API Contém informações sobre essa API
type API struct {
	Version     string `json:"version"`
	Description string `json:"description"`
}

type todoHandler struct {
	store map[string]Todo
}

func storageTodo() *todoHandler {
	return &todoHandler{
		store: map[string]Todo{},
	}
}

// Home ponto de entrada da API, exibe as informações deste projeto
func Home(w http.ResponseWriter, r *http.Request) {

	var api API = API{Version: "1.0.0", Description: "Uma API de todo list"}
	jsonBytes, _ := json.Marshal(api)

	switch r.Method {
	case "GET":
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

// APITodos trata as ações de GET e POST para uma coleção de todo
func (h *todoHandler) APITodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		todos := make([]Todo, len(h.store))

		i := 0
		for _, todo := range h.store {
			todos[i] = todo
			i++
		}

		jsonBytes, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	case "POST":
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
			return
		}

		var todo Todo
		err = json.Unmarshal(bodyBytes, &todo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		todo.ID = fmt.Sprintf("%d", time.Now().UnixNano())
		h.store[todo.ID] = todo
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

// APITodo trata as ações de GET, DELETE e PATCH para uma tarefa
func (h *todoHandler) APITodo(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		todo, ok := h.store[parts[2]]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		jsonBytes, err := json.Marshal(todo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	case "DELETE":
		_, ok := h.store[parts[2]]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		delete(h.store, parts[2])
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func main() {
	fmt.Println("API server is running... on 3030")

	todo := storageTodo()

	http.HandleFunc("/", Home)
	http.HandleFunc("/todos", todo.APITodos)
	http.HandleFunc("/todo/", todo.APITodo)
	err := http.ListenAndServe(":3030", nil)

	if err != nil {
		panic(err)
	}
}
