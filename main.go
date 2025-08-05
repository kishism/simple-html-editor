package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const PORT = ":8000"

func main() {
	fmt.Println("Server running on port", PORT)

	http.HandleFunc("/", editorHandler)       // GET - Show the editor
	http.HandleFunc("/save", saveHandler)     // POST - Save content

	log.Fatal(http.ListenAndServe(PORT, nil))
}

func editorHandler(w http.ResponseWriter, r *http.Request) {
	loadedcontent := "" // this is not the same one in <textarea name="content"> previously
	if saved, err := os.ReadFile("saved.html"); err == nil {
		loadedcontent = string(saved) // []byte into string 
	}

	tmpl := template.Must(template.ParseFiles("templates/editor.tmpl.html"))

	data := struct {
		InitialContent string
	}{
		InitialContent: loadedcontent,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Println("Template error:", err)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {	// Go does this automatically but it's a good practice to explicity call this 
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	editorInput := r.FormValue("content")

	err := os.WriteFile("saved.html", []byte(editorInput), 0644)
	if err != nil {
		http.Error(w, "Failed to save content", http.StatusInternalServerError)
		log.Println("Save error:", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<div id="save-status" style="color: green;">Saved</div>`)
}
