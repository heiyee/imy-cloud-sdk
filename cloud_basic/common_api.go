package cloud_basic

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var (
	OK = &Err{Code: 0, Message: "OK"}
	// 系统错误, 前缀为 100
	InternalServerError = &Err{Code: 10001, Message: "内部服务器错误"}
	ParamError          = &Err{Code: 10002, Message: "请求参数错误"}
	ErrTokenSign        = &Err{Code: 10003, Message: "签名 jwt 时发生错误"}
)

type RequestHandler interface {
	Handle(c *gin.Context, request interface{}) error
}

type ResponseHandler interface {
}

type JsonRequestHandler struct {
	internalErrorHandler InternalErrorHandler
}

type Err struct {
	Code    int    // 错误码
	Message string // 展示给用户看的
}

func (e Err) Error() string {
	return e.Message
}

type ErrorField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// json response
type BaseResponse struct {
	Code        int          `json:"code"`
	Msg         string       `json:"msg"`
	Data        interface{}  `json:"data,omitempty"`
	ErrorFields []ErrorField `json:"errorFields,omitempty"`
}

func (j JsonRequestHandler) Handle(c *gin.Context, request interface{}) (e error) {
	s, _ := ioutil.ReadAll(c.Request.Body)
	e = json.Unmarshal(s, &request)
	if e != nil {
		j.internalErrorHandler.Handle(c)
		return InternalServerError
	}
	return nil
}

type InternalErrorHandler interface {
	Handle(c *gin.Context)
}

type JsonInternalErrorHandler struct {
}

func (s JsonInternalErrorHandler) Handle(c *gin.Context) {
	c.AbortWithStatusJSON(500, BaseResponse{Code: InternalServerError.Code, Msg: InternalServerError.Message})
}

type OctetStreamInternalErrorHandler struct {
}

func (s OctetStreamInternalErrorHandler) Handle(c *gin.Context) {
	c.AbortWithStatus(500)
}

type Initializer interface {
	GetHandlers() []WebHandler
	GetAddr() string
}

type WebHandler struct {
	ApiId   string
	Method  int
	Url     string
	Handler func(apiId string, c *gin.Context)
}
