package requests

import validation "github.com/go-ozzo/ozzo-validation"

// PostNews request.
type PostNews struct {
	Text string `json:"text"`
}

// ValidateJSON data.
func (r *PostNews) ValidateJSON() error {
	return r.Validate()
}

// Validate me.
func (r *PostNews) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Text, validation.Required),
	)
}
