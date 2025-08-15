package pipe_filter

// Request 定义 input 支持各种类型
type Request interface{}

// Response 定义 output 支持各种类型
type Response interface{}

// Filter 定义 filter 接口
type Filter interface {
	Process(Request) (Response, error)
}
