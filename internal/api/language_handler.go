package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/AdelYarR/Data-Blossom/pkg/models"
)

func (s *APIServer) languageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var lang models.Languages
		err := json.NewDecoder(r.Body).Decode(&lang)
		if err != nil {
			s.logger.Error("Failed to decode")
			return
		}

		err = s.db.NewLanguage(lang)
		if err != nil {
			s.logger.Error("Unsuccessful POST")
			return
		}

	case http.MethodGet:
		data, err := s.db.GetLanguages()
		if err != nil {
			s.logger.Error("Unsuccessful GET")
			return
		}

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			s.logger.Error("Unsuccesful encoding")
			return
		}

	case http.MethodPut:
		var lang models.Languages
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			s.logger.Error("Failed to decode")
			return
		}

		err := json.NewDecoder(r.Body).Decode(&lang)
		if err != nil {
			s.logger.Error("Failed to decode")
			return
		}

		id_, err := strconv.Atoi(id)
		if err != nil {
			s.logger.Error("Failed to decode")
			return
		}

		err = s.db.UpdateLanguage(id_, lang)
		if err != nil {
			s.logger.Error("Unsuccessful DELETE")
			return
		}

	case http.MethodDelete:
		vars := mux.Vars(r)
		name, ok := vars["name"]
		if !ok {
			s.logger.Error("Failed to decode")
			return
		}

		err := s.db.DeleteLanguage(name)
		if err != nil {
			s.logger.Error("Unsuccessful DELETE")
			return
		}
	}
}
