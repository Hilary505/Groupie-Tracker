package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"tracker/models"
)

// FetchLocations retrieves location data from the API.
func FetchLocations(w http.ResponseWriter, r *http.Request) (wrapper []models.Location, e error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		HandleError(w, err, 500, "500.html")
		e = err
		return
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Println("Error closing response body:", closeErr)
		}
	}()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, fmt.Errorf("failed to read response body: %w", readErr)
	}

	var locationIndex models.LocationsResponse
	if jsonErr := json.Unmarshal(body, &locationIndex); jsonErr != nil {
		return nil, fmt.Errorf("failed to decode locations: %w", jsonErr)
	}
	wrapper = locationIndex.Index
	return wrapper, nil
}
