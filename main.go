package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
)

type Project struct {
    Name        string
    TechStack   string
    Description string
    Details     []string
}

type Experience struct {
    Company     string
    Role        string
    Duration    string
    Description []string
}

type PortfolioData struct {
    Name        string
    Title       string
    Summary     string
    Email       string
    Github      string
    LinkedIn    string
    Skills      []string
    Experiences []Experience
    Projects    []Project
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }

    data := PortfolioData{
        Name:  "Dharhshini V Jaisankar",
        Title: "AI Engineer + Full Stack Developer",
        Summary: "Electronics and Computer Engineering student with hands-on experience building full stack web-applications, MCP servers, Web3 Contracts, and full-stack mobile apps. Comfortable across Flutter/React Native, Node.js/Python, and Neo4j/MongoDB/Firebase. Exploring AI harnesses, agentic tooling, and system design patterns.",
        Email: "dharhshinivj30@gmail.com",
        Github: "github.com/DharhshiniVJ",
        LinkedIn: "LinkedIn",
        Skills: []string{
            "Python", "Java", "C++", "C", "TypeScript", "JavaScript", "GoLang",
            "Flutter", "Dart", "React Native", "Firebase", "REST APIs",
            "MCP Servers", "Neo4j", "Graph RAG", "MongoDB", "RAG", "FAISS", "Node.js", "FastAPI",
            "Solidity", "Hardhat", "MetaMask", "React", "HTML", "CSS","Langchain", "LLM", "Agentic Tooling",
            "Arduino", "ESP32", "Docker", "Git", "GitHub","Jitsi SDK", "Gemini API", "OpenAI API", "Cerebras LLM",
        },
        Experiences: []Experience{
            {
                Company:  "Educational Initiatives",
                Role:     "Software Summer Intern",
                Duration: "June 2026",
                Description: []string{
                    "Learned how AI harnesses are designed and orchestrated, including MCP (Model Context Protocol)-based tool integrations for connecting LLMs to real systems.",
                    "Built a deeper understanding of software design patterns and applied them to improve code structure and maintainability in an existing production codebase.",
                    "Gained hands-on exposure to Neo4j graph database modeling and broader backend development practices.",
                },
            },
            {
                Company:  "Zhians On-Site",
                Role:     "Software Development Engineer (Internship)",
                Duration: "Jul 2025 - Aug 2025",
                Description: []string{
                    "Built a multi-role fertility tracking application ecosystem (Patient, Doctor, Seller apps) using Flutter and Firebase.",
                    "Developed Firestore data models, authentication, Cloud Functions, storage management, and push notifications.",
                    "Integrated the Gemini API for AI-driven guidance and the Jitsi SDK for real-time video calling.",
                    "Implemented responsive UI/UX and REST API integrations; collaborated using Git-based workflows, improving performance and code maintainability.",
                },
            },
        },
        Projects: []Project{
            {
                Name:        "StudyDB",
                TechStack:   "React, FastAPI, Neo4j, Cerebras LLM, Docker, Redis",
                Description: "Adaptive quiz platform with an MCP-based AI agent and GraphRAG for hallucination-free, context-aware learning.",
                Details: []string{
                    "Agentic MCP Orchestrator: Bridged FastAPI and React via SSE and Action Interceptor Loops, enabling autonomous LLM tool execution.",
                    "GraphRAG & Dual-Layer Cache: Eliminated LLM hallucinations via PyPDF chunking and Neo4j Cypher traversals; achieved sub-second latency.",
                    "Containerized Cognitive Framework: Deployed Docker microservices with JWT auth; ran graph queries for real-time adaptive cognitive profiling.",
                },
            },
            {
                Name:        "DeWorkHub",
                TechStack:   "Next.js, Solidity, Hardhat, Ethers.js, MongoDB",
                Description: "Trustless Web3 freelance platform — contracts, payments, and reputation fully on-chain, no middleman.",
                Details: []string{
                    "Next.js & MongoDB: Built a full-stack platform handling user authentication, real-time chat, job boards, and tracking.",
                    "Solidity & Hardhat: Implemented smart contracts and deployment scripts.",
                },
            },
        },
    }

    // Parse the HTML template used for rendering the portfolio page.
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        log.Fatalf("failed to parse template: %v", err)
    }
    // Ensure the variable is used even if the handler does not reference it directly.
    _ = tmpl

    // Simple HTTP handler that renders the parsed template with the portfolio data.
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if err := tmpl.Execute(w, data); err != nil {
            log.Printf("template execution error: %v", err)
        }
    })

    log.Printf("Server starting on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}
