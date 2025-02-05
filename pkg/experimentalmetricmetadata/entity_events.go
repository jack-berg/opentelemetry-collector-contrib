// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package experimentalmetricmetadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// See entity event design document:
// https://docs.google.com/document/d/1Tg18sIck3Nakxtd3TFFcIjrmRO_0GLMdHXylVqBQmJA/edit#heading=h.pokdp8i2dmxy

const (
	semconvOtelEntityEventName    = "otel.entity.event.type"
	semconvEventEntityEventState  = "entity_state"
	semconvEventEntityEventDelete = "entity_delete"

	semconvOtelEntityID         = "otel.entity.id"
	semconvOtelEntityType       = "otel.entity.type"
	semconvOtelEntityAttributes = "otel.entity.attributes"
)

// EntityEventsSlice is a slice of EntityEvent.
type EntityEventsSlice struct {
	orig plog.LogRecordSlice
}

// NewEntityEventsSlice creates an empty EntityEventsSlice.
func NewEntityEventsSlice() EntityEventsSlice {
	return EntityEventsSlice{orig: plog.NewLogRecordSlice()}
}

// AppendEmpty will append to the end of the slice an empty EntityEvent.
// It returns the newly added EntityEvent.
func (s EntityEventsSlice) AppendEmpty() EntityEvent {
	return EntityEvent{orig: s.orig.AppendEmpty()}
}

// Len returns the number of elements in the slice.
func (s EntityEventsSlice) Len() int {
	return s.orig.Len()
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
func (s EntityEventsSlice) EnsureCapacity(newCap int) {
	s.orig.EnsureCapacity(newCap)
}

// At returns the element at the given index.
func (s EntityEventsSlice) At(i int) EntityEvent {
	return EntityEvent{orig: s.orig.At(i)}
}

// EntityEvent is an entity event.
type EntityEvent struct {
	orig plog.LogRecord
}

// ID of the entity.
func (e EntityEvent) ID() pcommon.Map {
	m, ok := e.orig.Attributes().Get(semconvOtelEntityID)
	if !ok {
		return e.orig.Attributes().PutEmptyMap(semconvOtelEntityID)
	}
	return m.Map()
}

// SetEntityState makes this an EntityStateDetails event.
func (e EntityEvent) SetEntityState() EntityStateDetails {
	e.orig.Attributes().PutStr(semconvOtelEntityEventName, semconvEventEntityEventState)
	return e.EntityStateDetails()
}

// EntityStateDetails returns the entity state details of this event.
func (e EntityEvent) EntityStateDetails() EntityStateDetails {
	return EntityStateDetails(e)
}

// SetEntityDelete makes this an EntityDeleteDetails event.
func (e EntityEvent) SetEntityDelete() EntityDeleteDetails {
	e.orig.Attributes().PutStr(semconvOtelEntityEventName, semconvEventEntityEventDelete)
	return e.EntityDeleteDetails()
}

// EntityDeleteDetails return the entity delete details of this event.
func (e EntityEvent) EntityDeleteDetails() EntityDeleteDetails {
	return EntityDeleteDetails(e)
}

// EventType is the type of the entity event.
type EventType int

const (
	// EventTypeNone indicates an invalid or unknown event type.
	EventTypeNone EventType = iota
	// EventTypeState is the "entity state" event.
	EventTypeState
	// EventTypeDelete is the "entity delete" event.
	EventTypeDelete
)

// EventType returns the type of the event.
func (e EntityEvent) EventType() EventType {
	eventType, ok := e.orig.Attributes().Get(semconvOtelEntityEventName)
	if !ok {
		return EventTypeNone
	}

	switch eventType.Str() {
	case semconvEventEntityEventState:
		return EventTypeState
	case semconvEventEntityEventDelete:
		return EventTypeDelete
	default:
		return EventTypeNone
	}
}

// EntityStateDetails represents the details of an EntityState event.
type EntityStateDetails struct {
	orig plog.LogRecord
}

// Attributes returns the attributes of the entity.
func (s EntityStateDetails) Attributes() pcommon.Map {
	m, ok := s.orig.Attributes().Get(semconvOtelEntityAttributes)
	if !ok {
		return s.orig.Attributes().PutEmptyMap(semconvOtelEntityAttributes)
	}
	return m.Map()
}

// EntityType returns the type of the entity.
func (s EntityStateDetails) EntityType() string {
	t, ok := s.orig.Attributes().Get(semconvOtelEntityType)
	if !ok {
		return ""
	}
	return t.Str()
}

// SetEntityType sets the type of the entity.
func (s EntityStateDetails) SetEntityType(t string) {
	s.orig.Attributes().PutStr(semconvOtelEntityType, t)
}

// EntityDeleteDetails represents the details of an EntityDelete event.
type EntityDeleteDetails struct {
	orig plog.LogRecord
}
