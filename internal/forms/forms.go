package forms

import (
	"net/http"
	"net/url"
)

// Form is a collection of form values.
type Form struct {
	url.Values
	Error errors
}

// New returns a new form.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has gets a value from the form.
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
