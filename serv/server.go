package serv

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	asciiart "asciiart/src"
)

var asciiex string

var (
	tmpl  = template.Must(template.ParseFiles("templates/index.html"))
	tmpl2 = template.Must(template.ParseFiles("templates/output.html"))
	tmpl3 = template.Must(template.ParseFiles("templates/errors.html"))
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		data := map[string]any{"code": http.StatusMethodNotAllowed, "msg": "method not allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		tmpl3.Execute(w, data)
		return
	}
	if r.URL.Path != "/" {
		data := map[string]any{"code": http.StatusNotFound, "msg": "not found"}
		w.WriteHeader(http.StatusNotFound)
		tmpl3.Execute(w, data)
		return
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AsciiWeb(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		data := map[string]any{"code": http.StatusMethodNotAllowed, "msg": "method not allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		tmpl3.Execute(w, data)
		return
	}
	textInput := asciiart.CheckInput(r.FormValue("text"))
	textLines := strings.Split(textInput, "\r\n")
	banner := r.FormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		data := map[string]any{"code": http.StatusBadRequest, "msg": "bad request"}
		w.WriteHeader(http.StatusBadRequest)
		tmpl3.Execute(w, data)
		return
	}
	maps, err := asciiart.MapBanner(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	asciiArt := asciiart.Draw(maps, textLines)
	err = tmpl2.Execute(w, asciiArt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ExportAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		data := map[string]any{"code": http.StatusMethodNotAllowed, "msg": "method not allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		tmpl3.Execute(w, data)
		return
	}
	fileName := "ascii_art.txt"
	file, err := os.Create(fileName)
	if err != nil {
		panic("Unable to create file: " + err.Error())
	}
	defer file.Close()

	file.WriteString(asciiex)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprint(len(asciiex)))

	w.Write([]byte(asciiex))
}
