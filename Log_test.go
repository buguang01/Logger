package Logger_test

import (
	"testing"

	"github.com/buguang01/Logger"
)

func TestLog(t *testing.T) {
	//正常退出日志模块
	defer Logger.LogClose()

	defer func() {
		r := recover()
		if r != nil {
			Logger.PFatal(r)

		}
	}()
	//初始化日志模块
	Logger.Init(0, "logs")
	Logger.SetListenKeyID(1001)
	Logger.PDebug("test1")
	Logger.PInfo("test2")
	Logger.PInfoKey("test3", 1001)
	Logger.PDebugKey("test4", 1002)
	Logger.PDebugKey("test5", 1001)
	//errs := fmt.Errorf(string(debug.Stack()))
	panic("panicinfo")

}
