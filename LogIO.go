package Logger

//可以挂载到别的框架中
func NewLoggerIO(loglv LogLevel) *LoggerIO {
	return &LoggerIO{
		LogLv: loglv,
	}
}

//一般第三方的框架都是可以设置一个支持io.Write接口的对象传入就可以了
type LoggerIO struct {
	LogLv LogLevel
}

func (this *LoggerIO) Write(p []byte) (n int, err error) {
	PAlart(this.LogLv, string(p))
	return 0, nil
}
