package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StreamAnalytic struct {
	BaseModel
	ReleaseID        uuid.UUID `gorm:"type:uuid;column:release_id;index:idx_release_period;not null" json:"release_id"`
	StreamCount      int64     `gorm:"type:bigint;column:stream_count;not null;default:0" json:"stream_count"`
	RevenueGenerated float64   `gorm:"type:numeric(12,4);column:revenue_generated;not null;default:0.0000" json:"revenue_generated"`
	RecordedMonth    time.Time `gorm:"type:date;column:recorded_month;index:idx_release_period;not null" json:"recorded_month"`
}

func (StreamAnalytic) TableName() string { return "stream_analytics" }

func (sa *StreamAnalytic) BeforeCreate(tx *gorm.DB) (err error) {
	sa.ID = uuid.New()
	return
}
