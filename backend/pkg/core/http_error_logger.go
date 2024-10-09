package core

type errorLoggerData struct {
	info       string
	requestId  string
	stackTrace string
}

func NewErrorLoggerData(info string, requestId string, stackTrace string) *errorLoggerData {
	return &errorLoggerData{
		info:       info,
		requestId:  requestId,
		stackTrace: stackTrace,
	}
}

func (e *errorLoggerData) GetFields() map[string]interface{} {
	return map[string]interface{}{
		"request_id":  e.requestId,
		"stack_trace": e.stackTrace,
	}
}

func (e *errorLoggerData) GetInfo() string {
	return e.info
}
