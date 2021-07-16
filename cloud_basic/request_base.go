package cloud_basic

type HandlerChain []HandlerFunc

var requestConfigMap = make(map[string]*RequestConfig)

type RequestContext struct {
	Request  interface{}
	Response interface{}
}

type RequestConfig struct {
	Chains         HandlerChain
	CurrentChain   int
	RequestHandler RequestHandler
}

type HandlerFunc interface {
	Handle(context *RequestContext) error
}

func AddRequestConfig(apiId string, f func() *RequestConfig) {
	if apiId != "" {
		requestConfigMap[apiId] = f()
	}
}

func GetRequestConfig(apiId string) *RequestConfig {
	return requestConfigMap[apiId]
}
