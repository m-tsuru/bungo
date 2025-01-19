package main

import (
	"encoding/json"
	"html/template"

	"log"
	"math/rand"
	"os"
	"path/filepath"

	"net/http"
)

type Information struct {
	Title        string
	Quote        string
	Author       string
	Copyright    string
	License      string
	CreationDate string
}

type Value struct {
	Information Information
	Content     template.HTML
}

func getfileBaseName(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func parseInformation(filePath string) (Information, error) {
	var info Information
	file, err := os.Open(filePath)
	if err != nil {
		return info, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		return info, err
	}

	return info, nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	staticFileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", staticFileServer))
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	files, _ := filepath.Glob("./content/**.partial.html")
	num := rand.Intn(len(files))
	contentPath, infoPath := files[num], getfileBaseName(files[num])+".json"
	log.Println("[Choice]", contentPath, infoPath)

	var Value Value
	information, err := parseInformation("./content/" + infoPath)
	if err != nil {
		log.Fatalf("[JSON Parsing Error], %s", err)
	}
	Value.Information = information

	contentBytes, err := os.ReadFile(contentPath)
	if err != nil {
		log.Fatalf("[File Reading Error], %s", err)
	}
	Value.Content = template.HTML(contentBytes)

	template := template.Must(template.ParseFiles("./template/index.html"))
	err = template.Execute(w, Value)
	if err != nil {
		log.Fatalf("[HTML Generative Error] %s", err)
	}
}
