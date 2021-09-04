package log

import (
	"time"
)

type LogError struct {
	ID           string    `json:"id"`
	Header       string    `json:"request_header"`
	Body         string    `json:"request_body"`
	URL          string    `json:"url"`
	HttpMethod   string    `json:"http_method"`
	Email        string    `json:"email"`
	ErrorMessage string    `json:"error_message"`
	Level        string    `json:"level"`
	AppName      string    `json:"app_name"`
	Version      string    `json:"version"`
	Env          string    `json:"env"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LogActivity struct {
	ID        string    `json:"id"`
	TableName string    `json:"table_name"`
	Email     string    `json:"email"`
	Row       string    `json:"row"`
	NewData   string    `json:"new_data"`
	OldData   string    `json:"old_data"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LogLogin struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DateLogin   string `json:"date_login"`
	DateLogout  string `json:"date_logout"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
