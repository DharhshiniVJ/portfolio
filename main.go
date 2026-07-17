package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Project represents a single portfolio project
type Project struct {
	Name        string
	Description string
	Tags        []string
	Link        string
}

// PortfolioData holds all the data to render the portfolio
type PortfolioData struct {
	Name     string
	Role     string
	About    string
	Projects []Project
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Mock data for the portfolio
	data := PortfolioData{
		Name:  "Dharhshini V J",
		Role:  "Software Engineer & AI Integrator",
		About: "I build robust CI/CD pipelines, intelligent automation tools, and scalable web services using Go, Docker, and AI agents.",
		Projects: []Project{
			{
				Name:        "ShipIt",
				Description: "An AI-powered, language-agnostic CI/CD platform built in Go. Features automated webhook handling, AI build planning, and a multi-agent AI pipeline to fix build failures automatically.",
				Tags:        []string{"Go", "Kafka", "Docker", "LangChain"},
				Link:        "#",
			},
			{
				Name:        "Demo App",
				Description: "A Go REST API used to test the ShipIt CI/CD platform. Features endpoints for task management and automated testing.",
				Tags:        []string{"Go", "REST API"},
				Link:        "#",
			},
		},
	}

	// Serve static files (CSS, JS, Images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle the main route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	log.Printf("🚀 Starting portfolio server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
