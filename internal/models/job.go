package models

type Job struct {
	JobID   string        `json:"job_id"`
	Count   int           `json:"count"`
	Visits  []Visit       `json:"visits"`
	Status  string        `json:"status"`
	Results []ImageResult `json:"results"`
}

type ImageResult struct {
	StoreID   string `json:"store_id"`
	ImageURL  string `json:"image_url"`
	Perimeter int    `json:"perimeter"`
	Processed bool   `json:"processed"`
	Error     string `json:"error,omitempty"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}
