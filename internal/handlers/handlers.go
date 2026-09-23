package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// send a client an html from a file
func HandlerRoot(w http.ResponseWriter, r *http.Request){
	// read an html file
	fileData, err := os.ReadFile("index.html")
	if err != nil{
    	http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
	
	// send a data to the client
	w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)
    w.Write(fileData)
}

// use html and send a client a converted string 
func HandlerUpload(w http.ResponseWriter, r *http.Request){
	// parse an html form
	err := r.ParseMultipartForm(2 << 20)
	if err != nil{
    	http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

	// find a name buy the key
	var key string
	for k := range r.MultipartForm.File{
		key = k
	}
	
	// get the file and the header
	file, header, err := r.FormFile(key)
	if err != nil{
    	http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
	defer file.Close()

	// read the file
	data, err := io.ReadAll(file)
	if err != nil{
    	http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
	
	// convert to morse or to text
	dst := service.ConvertTextMorse(string(data))

	// make a file name
	decFileName := time.Now().UTC().Format("02_01_2006") + filepath.Ext(header.Filename)

	// create a new file
	decodedFile, err := os.OpenFile(decFileName, os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil{
    	http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
	defer decodedFile.Close()

	// write into the new file
	_, err = decodedFile.Write([]byte(dst))
	if err != nil{
    	http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

	// send a data to the client
	w.Write([]byte(dst))
}