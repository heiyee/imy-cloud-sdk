package cloud_basic

type HandlerChain []HandlerFunc

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
