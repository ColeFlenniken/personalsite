package API

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ColeFlenniken/personalsite/Templates"
	"github.com/ColeFlenniken/personalsite/models"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	Templates.Index().Render(r.Context(), w)

}

func GetResume(w http.ResponseWriter, r *http.Request) {
	Templates.ResumePage().Render(r.Context(), w)
}

func GetJobCardHTML(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("card")
	for i := 0; i < len(models.Jobs); i++ {
		if id == models.Jobs[i].Id {
			Templates.ResumePageCardHTML(models.Jobs[i]).Render(r.Context(), w)
		}
	}
}

type CanvasData struct {
	Pixels []int `json:"pixels"`
}

var canvasData CanvasData = CanvasData{Pixels: make([]int, 550*300*4)}

func UpdateCanvasData(w http.ResponseWriter, r *http.Request) {
	var jsonData CanvasData

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	expectedSize := 300 * 550 * 4
	if len(jsonData.Pixels) != expectedSize {
		http.Error(w, fmt.Sprintf("Invalid pixel array size. Expected %d, got %d", expectedSize, len(jsonData.Pixels)), http.StatusBadRequest)
		return
	}

	// Update the global canvas data
	for i := 0; i < len(canvasData.Pixels); i++ {
		if jsonData.Pixels[i] != 0 && canvasData.Pixels[i] != jsonData.Pixels[i] {
			canvasData.Pixels[i] = jsonData.Pixels[i]
		}
	}
	w.WriteHeader(http.StatusOK)
}

func GetCanvasPixels(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(canvasData)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(data)
}
