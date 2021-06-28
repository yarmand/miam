package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// rating
const ( // iota is reset to 0
	Reject = iota // c0 == 0
	Love   = iota // c1 == 1
	Keep   = iota // c2 == 2
)

// states
// use to recover on massive image list processing
const (
	Importing = iota
	Imported  = iota
)

// flags
const (
	ImportToFinish = 1 << iota // a == 1 (iota has been reset)
	ToEdit         = 1 << iota
	HasRaw         = 1 << iota
	IsACopy        = 1 << iota // when phot is a copy, the original ID is stored in the field OriginalPhoto
)

// Photo is used by pop to map your photos database table to your go code.
type Photo struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Path          string    `json:"path" db:"path"`
	JpegFilename  string    `json:"jpeg_filename" db:"jpeg_filename"`
	RawFilename   string    `json:"raw_filename" db:"raw_filename"`
	Rating        int       `json:"rating" db:"rating"`
	Flags         int64     `json:"flags" db:"flags"`
	OriginalPhoto uuid.UUID `json:"original_photo" db:"original_photo"` // the ID of original if this one is a copy
}

// String is not required by pop and may be deleted
func (p Photo) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// photos is not required by pop and may be deleted
type photos []Photo

// String is not required by pop and may be deleted
func (p photos) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Photo) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Photo) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Photo) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
