package constants

// TagType represents the type of a tag (system or custom)
type TagType int8

func (t TagType) Int8() int8 {
	return int8(t)
}

const (
	// TagTypeSystem represents a system tag
	TagTypeSystem TagType = 1
	// TagTypeCustom represents a custom user-defined tag
	TagTypeCustom TagType = 2
)
