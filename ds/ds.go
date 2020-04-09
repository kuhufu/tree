package ds

type Request interface {
	Method() string
	Path() string
	Body() []byte
	SetBody(body []byte)
}

type Response interface {
	SetContentType(contentType ContentType)
	SetCode(code int)
	SetBody(data []byte)
	SetMsg(msg string)
}
