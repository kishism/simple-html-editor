package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

const PORT = ":8000"

func main() {
	fmt.Println("Server running on port", PORT)

	http.HandleFunc("/", editorHandler)       // GET - Show the editor
	http.HandleFunc("/save", saveHandler)     // POST - Save content
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/format/bold", boldHandler)
	http.HandleFunc("/format/italic", italicHandler)
	http.HandleFunc("/format/underline", underlineHandler)

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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	savedContent := ""

	if data, err := os.ReadFile("saved.html"); err == nil {
		savedContent = string(data) 
	}

	tmpl := template.Must(template.ParseFiles("templates/view.tmpl.html"))

	viewData := struct {
		ViewContent template.HTML
	}{
		ViewContent: template.HTML(savedContent),
	}

	if err := tmpl.Execute(w, viewData); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Println("View template error:", err)
	}
}

func boldHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content-hidden")
	startStr := r.FormValue("start")
	endStr := r.FormValue("end")

	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)

	if start >= end || start < 0 || end > len(content) {
		http.Error(w, "Invalid selection", http.StatusBadRequest)
		return
	}

	newContent := content[:start] + "<b>" + content[start:end] + "</b>" + content[end:]

	tmpl := template.Must(template.ParseFiles("templates/partial_editable.tmpl.html"))
    err = tmpl.Execute(w, struct {
        InitialContent string
    }{InitialContent: newContent})

    if err != nil {
        http.Error(w, "Server Error", http.StatusInternalServerError)
        return
    }
}

func italicHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content-hidden")
	startStr := r.FormValue("start")
	endStr := r.FormValue("end")

	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)

	if start >= end || start < 0 || end > len(content) {
		http.Error(w, "Invalid selection", http.StatusBadRequest)
		return
	}

	newContent := content[:start] + "<i>" + content[start:end] + "</i>" + content[end:]

	tmpl := template.Must(template.ParseFiles("templates/partial_editable.tmpl.html"))
    err = tmpl.Execute(w, struct {
        InitialContent string
    }{InitialContent: newContent})

    if err != nil {
        http.Error(w, "Server Error", http.StatusInternalServerError)
        return
    }
}

func underlineHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content-hidden")
	startStr := r.FormValue("start")
	endStr := r.FormValue("end")

	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)

	if start >= end || start < 0 || end > len(content) {
		http.Error(w, "Invalid selection", http.StatusBadRequest)
		return
	}

	newContent := content[:start] + "<u>" + content[start:end] + "</u>" + content[end:]

	tmpl := template.Must(template.ParseFiles("templates/partial_editable.tmpl.html"))
    err = tmpl.Execute(w, struct {
        InitialContent string
    }{InitialContent: newContent})

    if err != nil {
        http.Error(w, "Server Error", http.StatusInternalServerError)
        return
    }
}