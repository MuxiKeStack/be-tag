package domain

type TagType uint8

func (t TagType) Uint8() uint8 {
	return uint8(t)
}

const (
	TagTypeAssessment TagType = iota
	TagTypeFeature
)

type CountTagItem struct {
	Tag   int32
	Count int64
}
