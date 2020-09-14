package response

var (
	SUCCESS = NewResponse(20000, "success", nil)
	FAIL = NewResponse(40000, "fail", nil)
	ERROR = NewResponse(50000, "error", nil)
)
