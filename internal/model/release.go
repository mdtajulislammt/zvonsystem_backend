package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReleaseStatus string

const (
	StatusDraft        ReleaseStatus = "draft"
	StatusInModeration ReleaseStatus = "in_moderation"
	StatusApproved     ReleaseStatus = "approved"
	StatusChangesReq   ReleaseStatus = "action_required"
)

type Release struct {
	BaseModel
	UserID             uuid.UUID     `gorm:"type:uuid;column:user_id;index;not null" json:"user_id"`
	Title              string        `gorm:"type:varchar(255);column:title;not null" json:"title"`
	UPC                *string       `gorm:"type:varchar(50);column:upc;uniqueIndex" json:"upc,omitempty"`
	Grid               *string       `gorm:"type:varchar(50);column:grid" json:"grid,omitempty"`
	Language           string        `gorm:"type:varchar(50);column:language;not null" json:"language"`
	PrimaryGenre       string        `gorm:"type:varchar(100);column:primary_genre;not null" json:"primary_genre"`
	SecondaryGenre     *string       `gorm:"type:varchar(100);column:secondary_genre" json:"secondary_genre,omitempty"`
	CoverArtURL        string        `gorm:"type:text;column:cover_art_url;not null" json:"cover_art_url"`
	Status             ReleaseStatus `gorm:"type:varchar(50);column:status;not null;default:'draft'" json:"status"`
	ReleaseDate        time.Time     `gorm:"type:date;column:release_date;not null" json:"release_date"`
	SpotifyPublishDate *time.Time    `gorm:"type:date;column:spotify_publish_date" json:"spotify_publish_date,omitempty"`

	// Relations
	Tracks    []Track          `gorm:"foreignKey:ReleaseID;constraint:OnDelete:CASCADE" json:"tracks,omitempty"`
	Stores    []ReleaseStore   `gorm:"foreignKey:ReleaseID;constraint:OnDelete:CASCADE" json:"stores,omitempty"`
	Artists   []ReleaseArtist  `gorm:"foreignKey:ReleaseID;constraint:OnDelete:CASCADE" json:"artists,omitempty"`
	Analytics []StreamAnalytic `gorm:"foreignKey:ReleaseID;constraint:OnDelete:CASCADE" json:"analytics,omitempty"`
}

func (Release) TableName() string { return "releases" }

func (r *Release) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}

// Track - Dynamic audio asset parts
type Track struct {
	BaseModel
	ReleaseID    uuid.UUID `gorm:"type:uuid;column:release_id;index;not null" json:"release_id"`
	Title        string    `gorm:"type:varchar(255);column:title;not null" json:"title"`
	ISRC         *string   `gorm:"type:varchar(50);column:isrc;uniqueIndex" json:"isrc,omitempty"`
	AudioFileURL string    `gorm:"type:text;column:audio_file_url;not null" json:"audio_file_url"`
	Duration     int       `gorm:"type:integer;column:duration;not null" json:"duration"` // In seconds
}

func (Track) TableName() string { return "tracks" }

func (t *Track) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

// ReleaseArtist maps Contributors/Artists to Releases
type ReleaseArtist struct {
	BaseModel
	ReleaseID  uuid.UUID `gorm:"type:uuid;column:release_id;uniqueIndex:idx_release_artist_role;not null" json:"release_id"`
	ArtistName string    `gorm:"type:varchar(255);column:artist_name;uniqueIndex:idx_release_artist_role;not null" json:"artist_name"`
	Role       string    `gorm:"type:varchar(50);column:role;uniqueIndex:idx_release_artist_role;not null" json:"role"` // main_artist, feature, remixer, producer
}

func (ReleaseArtist) TableName() string { return "release_artists" }

func (ra *ReleaseArtist) BeforeCreate(tx *gorm.DB) (err error) {
	ra.ID = uuid.New()
	return
}

// ReleaseStore represents platforms target checklist
type ReleaseStore struct {
	BaseModel
	ReleaseID uuid.UUID `gorm:"type:uuid;column:release_id;uniqueIndex:idx_release_store;not null" json:"release_id"`
	StoreName string    `gorm:"type:varchar(100);column:store_name;uniqueIndex:idx_release_store;not null" json:"store_name"` // Spotify, Apple Music, etc.
	IsActive  bool      `gorm:"type:boolean;column:is_active;not null;default:true" json:"is_active"`
}

func (ReleaseStore) TableName() string { return "release_stores" }

func (rs *ReleaseStore) BeforeCreate(tx *gorm.DB) (err error) {
	rs.ID = uuid.New()
	return
}
