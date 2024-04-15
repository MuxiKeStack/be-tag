package domain

type TagType uint8

func (t TagType) Uint8() uint8 {
	return uint8(t)
}

const (
	TagTypeAssessment TagType = iota
	TagTypeFeature
)
