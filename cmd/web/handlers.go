package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{"./ui/html/base.tmpl.html", "./ui/html/pages/home.tmpl.html","./ui/html/partials/nav.tmpl.html",}
	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we use ... to pass the contents
	// of the files slice as variadic argument
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
	}
	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w,"base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue()
	// and try to convert it to an integer using the strconv.Atoi() function. If
	// it can't be converted to an integer, or the value is less than 1, we
	// return a 404 page not found response.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.logger.Error(err.Error())
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w,"Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Set a new cache-control header. If an existing "Cache-Control" header exists
	// it will be overwritten.
	w.Header().Set("Cache-Control","public, max-age=31536000")
	// In contrast, the Add() method appends a new "Cache-Control" header and can
	// be called multiple times.
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=31536000")
	// Delete all values for the "Cache-Control" header.
	w.Header().Del("Cache-Control")
	// Retrieve the first value for the "Cache-Control" header.
	w.Header().Get("Cache-Control")
	// Retrieve a slice of all values for the "Cache-Control" header.
	w.Header().Values("Cache-Control")
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request){
	// Use the w.WriteHeader() method to send a 201 status code.
	w.WriteHeader(http.StatusCreated)
	// Then w.Write() method to write the response body as normal.
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"Alex"}`))

}