package models

import (
	"time"
)

type Receipe struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Tags        []string  `json:"tags"`
	PublishedAt time.Time `json:"publishedAt"`
}
