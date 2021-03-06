package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mattn/go-sqlite3"
	"graffiti/ent"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var client *ent.Client
var ctx context.Context

func main() {
	var err error
	client, err = ent.Open("sqlite3", "file:graffiti.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx = context.Background()

	if _, err := os.Stat("graffiti.db"); os.IsNotExist(err) {
		if err = client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", todoPage)
	r.Post("/", postTodoPage)

	fmt.Println("Now serving on http://127.0.0.1:8000")
	http.ListenAndServe(":8000", r)
}

func todoPage(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		filepath.Join("templates", "base.gohtml"),
		filepath.Join("templates", "todopage.gohtml"),
	}

	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		log.Println(err.Error())
	}

	tasks, err := QueryTasks(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.ExecuteTemplate(w, "base", tasks); err != nil {
		log.Println("err:", err)
	}
}

func postTodoPage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	activity := r.FormValue("activity")
	CreateTask(ctx, client, activity)

	paths := []string{
		filepath.Join("templates", "base.gohtml"),
		filepath.Join("templates", "todopage.gohtml"),
	}

	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		log.Println(err.Error())
	}

	tasks, err := QueryTasks(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.ExecuteTemplate(w, "base", tasks); err != nil {
		log.Println("err:", err)
	}
}

func CreateTask(ctx context.Context, client *ent.Client, activity string) (*ent.Task, error) {
	task, err := client.Task.
		Create().
		SetActivity(activity).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Task: %v", err)
	}
	log.Println("Task was created: ", task)
	return task, nil
}

func QueryTasks(ctx context.Context, client *ent.Client) ([]*ent.Task, error) {
	tasks, err := client.Task.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying Tasks: %v", err)
	}
	log.Println("Tasks returned: ", tasks)
	return tasks, nil
}