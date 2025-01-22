package models

type Job struct {
	JobID  string  `json:"job_id"`
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
	Status string  `json:"status"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}
