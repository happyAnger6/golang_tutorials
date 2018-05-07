package lib

import "time"

type RawReq struct {
	ID  int64
	Req []byte
}

type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

type RetCode int

const (
	RET_CODE_SUCCESS              RetCode = 0
	RET_CODE_WARNING_CALL_TIMEOUT         = 1001 // 调用超时警告。
	RET_CODE_ERROR_CALL                   = 2001 // 调用错误。
	RET_CODE_ERROR_RESPONSE               = 2002 // 响应内容错误。
	RET_CODE_ERROR_CALEE                  = 2003 // 被调用方（被测软件）的内部错误。
	RET_CODE_FATAL_CALL                   = 3001 // 调用过程中发生了致命错误！
)

func GetRetCodePlain(code RetCode) string {
	var codePlain string
	switch code {
	case RET_CODE_SUCCESS:
		codePlain = "success"
	case RET_CODE_WARNING_CALL_TIMEOUT:
		codePlain = "Call Timeout Warning"
	case RET_CODE_ERROR_CALL:
		codePlain = "Call Error"
	case RET_CODE_ERROR_RESPONSE:
		codePlain = "Response Error"
	case RET_CODE_ERROR_CALEE:
		codePlain = "Callee Error"
	case RET_CODE_FATAL_CALL:
		codePlain = "Call Fatal Error"
	default:
		codePlain = "Unknown result code"
	}
	return codePlain
}

type CallResult struct {
	ID int64
	Req RawReq
	Resp RawResp
	Code RetCode
	Msg string
	Elapse time.Duration
}

const (
	STATUS_ORIGINAL uint32 = 0
	// STATUS_STARTING 代表正在启动。
	STATUS_STARTING uint32 = 1
	// STATUS_STARTED 代表已启动。
	STATUS_STARTED uint32 = 2
	// STATUS_STOPPING 代表正在停止。
	STATUS_STOPPING uint32 = 3
	// STATUS_STOPPED 代表已停止。
	STATUS_STOPPED uint32 = 4
)

type Generator interface {
	Start() bool
	Stop() bool
	Status() uint32
	CallCount() int64
}
