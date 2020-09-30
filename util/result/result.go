package result

import (
	"github.com/kuhufu/tree"
	"github.com/kuhufu/tree/util/errs"
	"log"
)

var GetCurUserIdFunc func(c *tree.Context) int64

func Ok(c *tree.Context, data interface{}, msg ...string) {
	c.Abort()
	var tmpMsg string
	if len(msg) > 0 {
		tmpMsg = msg[0]
	}
	c.JSON(200, data, tmpMsg)
}

func OkMsg(c *tree.Context, msg string) {
	c.Abort()
	c.JSON(200, nil, msg)
}

func Custom(c *tree.Context, code int, data interface{}, msg ...string) {
	c.Abort()
	var tmpMsg string
	if len(msg) > 0 {
		tmpMsg = msg[0]
	}
	c.JSON(code, data, tmpMsg)
}

func Fail(c *tree.Context, err error) {
	c.Abort()
	var (
		logMsg  string
		msg     string
		msgCode int
		uid     int64
		data    interface{}
	)

	if GetCurUserIdFunc != nil {
		uid = GetCurUserIdFunc(c)
	}

	switch {
	case errs.IsErrInternal(err):
		msgCode = err.(*errs.ErrInternal).Code
		msg = "内部错误"
		logMsg = "内部错误"
	case errs.IsErrParam(err):
		msgCode = err.(*errs.ErrParam).Code
		msg = err.Error()
		logMsg = "参数错误"
	case errs.IsErrBusiness(err):
		msgCode = err.(*errs.ErrBusiness).Code
		msg = err.Error()
		logMsg = "业务错误"
	case errs.IsErrCustom(err):
		msgCode = err.(*errs.ErrCustom).Code
		msg = err.Error()
		logMsg = "自定义错误"
		data = err.(*errs.ErrCustom).Data
	default:
		msgCode = 500
		msg = "未知错误"
		logMsg = "未处理的错误"
	}

	log.Printf("[error] %v\t uid=%v\t %v\t %v: %v",
		msgCode,
		uid,
		c.Request.Path(),
		logMsg,
		err,
	)

	c.JSON(msgCode, data, msg)
}
