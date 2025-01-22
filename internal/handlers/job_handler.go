package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"retail-pulse-image-processor/internal/models"
)

var jobs = make(map[string]models.Job)

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var job models.Job

	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if job.Count != len(job.Visits) {
		http.Error(w, "Count does not match number of visits", http.StatusBadRequest)
		return
	}

	jobID := uuid.New().String()
	job.JobID = jobID
	job.Status = "ongoing"
	jobs[jobID] = job

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"job_id": jobID})
}
