package cloud_basic

import (
	"github.com/gin-gonic/gin"
)

var requestMap = make(map[string]*RequestConfig)

func init() {
	requestMap[""] = getApi_1_context()
}

func getApi_1_context() *RequestConfig {
	var handlerFunc = make([]HandlerFunc, 0)
	var requestHandler = JsonRequestHandler{internalErrorHandler: JsonInternalErrorHandler{}}
	return &RequestConfig{RequestHandler: requestHandler, Chains: handlerFunc}
}

func handle_1_api(c *gin.Context) {
	chains := make(HandlerChain, 0)
	var request interface{}
	config := requestMap[""]
	requestContext := &RequestContext{Request: request}
	err := config.RequestHandler.Handle(c, &request)
	if err != nil {
		return
	}
	for _, v := range chains {
		err := v.Handle(requestContext)
		if err != nil {
			return
		}
	}
}
