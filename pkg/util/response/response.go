package response

import "codeid-boiler/internal/abstraction"

type Meta struct {
	Success bool                        `json:"success" default:"true"`
	Message string                      `json:"message" default:"true"`
	Info    *abstraction.PaginationInfo `json:"info"`
}
