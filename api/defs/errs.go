package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrResponse{HttpSC: 401, Error: Err{Error: "user authentication failed", ErrorCode: "002"}}
	ErrorDBError                = ErrResponse{HttpSC: 500, Error: Err{Error: "db ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrResponse{HttpSC: 500, Error: Err{Error: "internal server error", ErrorCode: "004"}}
)
