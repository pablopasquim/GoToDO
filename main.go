package main

import (
	"html/template"
	"net/http"
	"sync"
)

// Task representa uma tarefa simples
type Task struct {
	Title string
	Done  bool
}

// slice de tarefas (memória)
var tasks []Task
var mu sync.Mutex

// handler da página principal
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	mu.Lock()
	defer mu.Unlock()
	tmpl.Execute(w, tasks)
}

// handler para adicionar tarefas
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		if title != "" {
			mu.Lock()
			tasks = append(tasks, Task{Title: title, Done: false})
			mu.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addTaskHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
