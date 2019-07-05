package Logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	//正常退出日志模块
	defer LogClose()

	defer func() {
		r := recover()
		if r != nil {
			PFatal(r)

		}
	}()
	//初始化日志模块
	Init(0, "logs", LogModeFmt)
	SetListenKeyID(1001)
	PDebug("test1")
	PInfo("test2")
	PInfoKey("test3", 1001)
	PDebugKey("test4", 1002)
	PDebugKey("test5", 1001)
	//errs := fmt.Errorf(string(debug.Stack()))
	panic("panicinfo")

}
