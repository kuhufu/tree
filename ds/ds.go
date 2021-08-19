package ds

type Request interface {
	Path() string
	Body() []byte
	SetBody(body []byte)
}

type Response interface {
	SetContentType(contentType ContentType)
	SetCode(code int)
	SetBody(data []byte)
	SetMsg(msg string)

	Code() int
	Body() []byte
	Msg() string
}
