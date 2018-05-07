package loadgen

import (
	"gopcp.v2/helper/log"
	"golang_tutorials/gocp/chapter4/loadgen/lib"
	"context"
	"time"
	"bytes"
	"math"
	"fmt"
)

var logger = log.DLogger()


// myGenerator 代表载荷发生器的实现类型。
type myGenerator struct {
	caller      lib.Caller           // 调用器。
	timeoutNS   time.Duration        // 处理超时时间，单位：纳秒。
	lps         uint32               // 每秒载荷量。
	durationNS  time.Duration        // 负载持续时间，单位：纳秒。
	concurrency uint32               // 载荷并发量。
	tickets     lib.GoTickets        // Goroutine票池。
	ctx         context.Context      // 上下文。
	cancelFunc  context.CancelFunc   // 取消函数。
	callCount   int64                // 调用计数。
	status      uint32               // 状态。
	resultCh    chan *lib.CallResult // 调用结果通道。
}

// NewGenerator 会新建一个载荷发生器。
func NewGenerator(pset ParamSet) (lib.Generator, error) {

	logger.Infoln("New a load generator...")
	if err := pset.Check(); err != nil {
		return nil, err
	}
	gen := &myGenerator{
		caller:     pset.Caller,
		timeoutNS:  pset.TimeoutNS,
		lps:        pset.LPS,
		durationNS: pset.DurationNS,
		status:     lib.STATUS_ORIGINAL,
		resultCh:   pset.ResultCh,
	}
	if err := gen.init(); err != nil {
		return nil, err
	}
	return gen, nil
}

func (gt *myGenerator) init() error {
	var buf bytes.Buffer
	buf.WriteString("Initialize Generator...")

	total64 := int64(gt.timeoutNS)/int64(1e9/gt.lps) + 1
	if total64 > math.MaxInt64 {
		total64 = math.MaxInt64
	}

	gt.concurrency = uint32(total64)
	tickets, err := lib.NewGoTickets(gt.concurrency)
	if err != nil {
		return err
	}
	gt.tickets = tickets

	buf.WriteString(fmt.Sprintf("Done concurrency(%d)", gt.concurrency))
	logger.Infoln(buf.String())
	return nil
}