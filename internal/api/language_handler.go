package api

import (
	"encoding/json"
	"net/http"
)

func (s *APIServer) languageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		
	case http.MethodGet:
		data, err := s.db.GetLanguages()
		if err != nil {
			s.logger.Error("Unsuccesful GET")
		}

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			s.logger.Error("Unsuccesful encoding")
		}
	case http.MethodPatch:

	case http.MethodDelete:

	}
}
