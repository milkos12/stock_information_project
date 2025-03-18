package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetListStocksHandler(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"

	nextPage := r.URL.Query().Get("next_page")
	fullURL := baseURL

	if nextPage != "" {
		fullURL = baseURL + "?next_page=" + nextPage
	}

	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MSwiZW1haWwiOiJtaWxsZXJhbmRyZXNvY0BnbWFpbC5jb20iLCJleHAiOjE3NDIzMDQ0ODcsImlkIjoiMCIsInBhc3N3b3JkIjoiJyBPUiAnYSc9J2EifQ.rA2OUi_VPV6ipXuvPjCu69j9zK_xXTKQcBAElZ6VcDk"

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error creando la solicitud"})
		return
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error haciendo la solicitus"})
		http.Error(w, "Error haciendo la solicitud: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "error leyendo la respuesta"})
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}
