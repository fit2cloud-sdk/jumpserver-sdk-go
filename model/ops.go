package model

// OpsJobRequest is the payload for creating a quick-command job.
type OpsJobRequest struct {
	Assets      []string `json:"assets"`
	Nodes       []string `json:"nodes"`
	Module      string   `json:"module"`
	Args        string   `json:"args"`
	RunAs       string   `json:"runas"`
	RunAsPolicy string   `json:"runas_policy"`
	Instant     bool     `json:"instant"`
	IsPeriodic  bool     `json:"is_periodic"`
	Timeout     int      `json:"timeout"`
}

// OpsJobResponse is the response after creating a job.
type OpsJobResponse struct {
	TaskID string `json:"task_id"`
}

// OpsJobResultSummary contains the execution summary of a job.
type OpsJobResultSummary struct {
	Ok       []string            `json:"ok"`
	Dark     map[string]any      `json:"dark"`
	Skipped  []string            `json:"skipped"`
	Failures map[string]any      `json:"failures"`
}

// OpsJobResult is the execution result of a quick-command job.
type OpsJobResult struct {
	Status     LabelValue          `json:"status"`
	IsFinished bool                `json:"is_finished"`
	IsSuccess  bool                `json:"is_success"`
	TimeCost   float64             `json:"time_cost"`
	JobID      string              `json:"job_id"`
	Summary    OpsJobResultSummary `json:"summary"`
}
