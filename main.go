package main

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
    // No operation – placeholder to satisfy the compiler.
}
