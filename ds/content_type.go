package ds

type ContentType int

const (
	ContentType_JSON = ContentType(iota + 1)
	ContentType_Text
	ContentType_ProtoBuf
	ContentType_Binary
)
