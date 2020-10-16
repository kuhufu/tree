package binding

import (
	. "github.com/kuhufu/tree/ds"
	"net/url"

	"strings"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req Request, obj interface{}) error {
	//
	//values := extract(req.path)
	//if err := mapForm(obj, values); err != nil {
	//	return err
	//}
	return nil
}

func extract(path string) url.Values {
	split := strings.Split(path, "?")
	var values url.Values
	if len(split) >= 2 {
		values, _ = url.ParseQuery(split[1])
	}

	return values
}
