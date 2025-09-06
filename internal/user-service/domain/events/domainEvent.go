package events

import (
	"time"

	"github.com/google/uuid"
)

type DomainEvent interface {
	EventId() string
	EventType() string
	OccurredAt() time.Time
	AggregateId() string
	AggregateType() string
	EventVersion() int
	EventData() interface{}
}

type BaseDomainEvent struct {
	eventId       string
	eventType     string
	occurredAt    time.Time
	aggregateId   string
	aggregateType string
	eventVersion  int
}

func NewBaseDomainEvent(eventType, aggregateId, aggregateType string, eventVersion int) BaseDomainEvent {
	id, _ := uuid.NewUUID()
	return BaseDomainEvent{
		eventId:       id.String(),
		eventType:     eventType,
		occurredAt:    time.Now(),
		aggregateId:   aggregateId,
		aggregateType: aggregateType,
		eventVersion:  eventVersion,
	}
}

func (e *BaseDomainEvent) EventId() string {
	return e.eventId
}

func (e *BaseDomainEvent) EventType() string {
	return e.eventType
}
func (e *BaseDomainEvent) OccurredAt() time.Time {
	return e.occurredAt
}
func (e *BaseDomainEvent) AggregateId() string {
	return e.aggregateId
}
func (e *BaseDomainEvent) AggregateType() string {
	return e.aggregateType
}
func (e *BaseDomainEvent) EventVersion() int {
	return e.eventVersion
}
