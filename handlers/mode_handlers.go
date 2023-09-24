package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MAXXKUMAR/Exp/controllers"
)

func GetModeCountsByArea(w http.ResponseWriter, r *http.Request) {
	// Extract areaCode from the request URL and pass it to the controller
	// For simplicity, let's assume areaCode is passed as a path variable.
	areaCode := r.URL.Path[len("/modes/"):]
	modeCounts, err := controllers.GetModeCounts(areaCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return modeCounts as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modeCounts)
}
