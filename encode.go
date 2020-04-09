package tree

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/kuhufu/tree/ds"
)

func Encode(contentType ds.ContentType, v interface{}) ([]byte, error) {
	switch contentType {
	case ds.ContentType_Text:
		switch res := v.(type) {
		case string:
			return []byte(res), nil
		default:
			return nil, errors.New(fmt.Sprintf("类型错误%T，对于contentType为text要求类型为string", v))
		}
	case ds.ContentType_Binary:
		if data, ok := v.([]byte); ok {
			return data, nil
		}
		return nil, errors.New(fmt.Sprintf("类型错误%T，对于contentType为binary要求类型为[]byte", v))
	case ds.ContentType_JSON:
		return json.Marshal(v)
	case ds.ContentType_ProtoBuf:
		if p, ok := v.(proto.Message); ok {
			return proto.Marshal(p)
		} else {
			return nil, errors.New(fmt.Sprintf("类型错误%T，对于contentType为protobuf要求类型为proto.Message", v))
		}

	default:
		return nil, errors.New(fmt.Sprintf("不支持的contentType%v", contentType))
	}
}
