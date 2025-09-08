package models

import (
	"time"

	"gorm.io/gorm"
)

type EventModel struct {
	Id            string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	EventId       string    `gorm:"type:uuid;uniqueIndex;not null" json:"event_id"`
	EventType     string    `gorm:"type:varchar(255);not null;index" json:"event_type"`
	AggregateId   string    `gorm:"type:uuid;not null;index" json:"aggregate_id"`
	AggregateType string    `gorm:"type:varchar(255);not null;index" json:"aggregate_type"`
	EventData     string    `gorm:"type:jsonb;not null" json:"event_data"`
	OccurredAt    time.Time `gorm:"not null" json:"occurred_at"`
	CreatedAt     time.Time `gorm:"not null" json:"created_at"`
}

func (EventModel) TableName() string {
	return "events"
}

func (e *EventModel) BeforeCreate(tx *gorm.DB) error {
	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now()
	}
	return nil
}
