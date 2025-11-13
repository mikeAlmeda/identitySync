package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/store"
)

func GetStudentsHandler(st store.StudentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		students, err := st.GetAll(ctx, map[string]interface{}{})
		if err != nil {
			http.Error(w, "failed to fetch students: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, err := json.MarshalIndent(students, "", " ")
		if err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}
