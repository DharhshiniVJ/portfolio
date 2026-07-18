package main

import (
    "log"
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
    // Use the port variable to avoid "declared and not used" errors.
    log.Printf("Server would start on port %s (placeholder)", port)
}
