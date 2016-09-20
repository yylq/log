package log

import (
	"runtime"
	"strconv"
	"testing"
	"time"
)

func writelog(i int) {
	Debug("Debug>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Warn("Warn>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
}
func writeflog(i int) {
	Debugf("Debug>>>>>>>>>>>>>>>>>>>>>>%d", i)
	Infof("Info>>>>>>>>>>>>>>>>>>>>>>>>>%d", i)
	Warnf("Warn>>>>>>>>>>>>>>>>>>>>>>>>>%d", i)
	Errorf("Error>>>>>>>>>>>>>>>>>>>>>>>>>%d", i)
	Fatalf("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>%d", i)
}

func testLog(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//指定是否控制台打印，默认为true
	SetConsole(true)

	SetRollingDaily("./", "test.log")

	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	//一般习惯是测试阶段为debug，生成环境为info以上
	SetLevel(ERROR)

	for i := 10; i > 0; i-- {
		go writelog(i)
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(15 * time.Second)
}
func TestLogf(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//指定是否控制台打印，默认为true
	SetConsole(true)

	SetRollingDaily("./", "test.log")

	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	//一般习惯是测试阶段为debug，生成环境为info以上
	SetLevel(ERROR)

	for i := 10; i > 0; i-- {
		writeflog(i)
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(15 * time.Second)
}
