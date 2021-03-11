package sinsp

// FieldEntry represents a single field entry that an extractor plugin can expose.
// Should be used when implementing plugin_get_fields().
type FieldEntry struct {
	Type string `json:"type"`
	ID   uint32 `json: "ID"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
