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

// flags
const (
	ToEdit = 1 << iota // a == 1 (iota has been reset)
)

// Photo is used by pop to map your photos database table to your go code.
type Photo struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Path      string    `json:"path" db:"path"`
	Rating    int       `json:"rating" db:"rating"`
	Flags     int64     `json:"flags" db:"flags"`
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
