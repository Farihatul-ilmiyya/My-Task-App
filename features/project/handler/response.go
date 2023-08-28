package handler

import (
	"mia/my_task_app/features/project"
	"time"
)

type ProjectResponse struct {
	Title       string    `json:"title"`
	UserID      uint      `json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// Mapping CorePrject to ProjectResponsee
func MapCoreProjToProjRes(core project.CoreProject) ProjectResponse {
	return ProjectResponse{
		Title:       core.Title,
		UserID:      core.UserID,
		Description: core.Description,
		CreatedAt:   core.CreatedAt,
	}
}
