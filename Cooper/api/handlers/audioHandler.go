package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	assambly "example.com/coomper/assistantCore/processors"
)

func Command(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := base64.StdEncoding.DecodeString(r.FormValue("data"))
	if err != nil {
		http.Error(w, "Error decoding base64 data: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, err := os.Create("audio/" + time.Now().GoString() + ".wav")
	if err != nil {
		http.Error(w, "Error creating the file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		http.Error(w, "Error saving the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	processor, err := assambly.NewAssamblyS2T()
	if err != nil {
		http.Error(w, "Error creating the processor: "+err.Error(), http.StatusInternalServerError)
		return
	}

	text, err := processor.Process(file.Name())
	if err != nil {
		http.Error(w, "Error processing the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(map[string]string{"text": text})
	w.Write(response)
}
