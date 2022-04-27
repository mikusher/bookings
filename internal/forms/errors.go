package forms

type errors map[string][]string

// Add an error to a field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get Returns the first error for the given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) > 0 {
		return es[0]
	}
	return ""
}
