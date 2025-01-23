package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"net/http"
	"retail-pulse-image-processor/internal/models"
	"retail-pulse-image-processor/internal/services"
	"time"
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

	var errors []map[string]string

	for _, visit := range job.Visits {
		if visit.StoreID == "" {
			errors = append(errors, map[string]string{"store_id": "Missing store_id"})
			continue
		}

		if len(visit.ImageURLs) == 0 {
			errors = append(errors, map[string]string{"store_id": visit.StoreID, "error": "Missing image_url"})
			continue
		}

		if visit.VisitTime == "" {
			errors = append(errors, map[string]string{"store_id": visit.StoreID, "error": "Missing visit_time"})
			continue
		}

		if _, exists := services.StoreMaster[visit.StoreID]; !exists {
			errors = append(errors, map[string]string{"store_id": visit.StoreID, "error": "Invalid store_id"})
			continue
		}
	}

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "failed",
			"job_id": "",
			"error":  errors,
		})
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

	go processJob(jobID)

	// Respond with the job ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"job_id": jobID})
}

func processJob(jobID string) {
	job := jobs[jobID]
	var errors []map[string]string

	for _, visit := range job.Visits {
		for _, imageURL := range visit.ImageURLs {
			result := models.ImageResult{
				StoreID:   visit.StoreID,
				ImageURL:  imageURL,
				Processed: false,
			}

			_, width, height, err := DownloadImage(imageURL)
			if err != nil {

				errors = append(errors, map[string]string{"store_id": visit.StoreID, "error": "Failed to download image"})
				continue
			}

			result.Perimeter = CalculatePerimeter(width, height)
			IntroduceRandomDelay()
			result.Processed = true

			job.Results = append(job.Results, result)
		}
	}

	if len(errors) > 0 {
		job.Status = "failed"
		job.Results = append(job.Results, models.ImageResult{Error: "Download failed"})
		jobs[jobID] = job

	} else {
		job.Status = "completed"
		jobs[jobID] = job
	}

}

func DownloadImage(url string) (data []byte, width, height int, err error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, 0, 0, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, 0, 0, err
	}

	bounds := img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()

	return nil, width, height, nil
}

func CalculatePerimeter(width, height int) int {
	return 2 * (width + height)
}

func IntroduceRandomDelay() {
	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
}
