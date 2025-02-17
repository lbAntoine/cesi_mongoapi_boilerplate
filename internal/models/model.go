package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func (m *BaseModel) GetID() primitive.ObjectID {
	return m.ID
}

func (m *BaseModel) SetID(id primitive.ObjectID) {
	m.ID = id
}

func (m *BaseModel) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *BaseModel) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

func (m *BaseModel) SetCreatedAt(t time.Time) {
	m.CreatedAt = t
}

func (m *BaseModel) SetUpdatedAt(t time.Time) {
	m.UpdatedAt = t
}
