package logging

import (
	"testing"
)

func test1() {

	subsytem := "CPLive"
	module := "LoggerUnitTest"

	logger, slogger, err := LogInit(subsytem, module)
	//logger, sLogger, err := LogInit()
	if err != nil {
		panic(err)
	}
	yes := "Yes"
	maybe := "Maybe!"
	logger.Info("This is a info test log")
	logger.Warn("This is a warn test log")
	logger.Error("This is an error test log")
	slogger.Infow("This is also a test log, but a bit different",
		"Does it work?", yes,
		"Is it useful?", maybe)
	slogger.Warnw("This is also a test log, but with a warning",
		"Does it work?", yes,
		"Is it useful?", maybe)
	slogger.Errorw("This is also a test log, but with an error",
		"Does it work?", yes,
		"Is it useful?", maybe)
}

func test2() {
	subsytem := "CPLive"
	module := "LoggerUnitTest"

	logger, slogger, err := CommonLogInit(subsytem, module)
	//logger, sLogger, err := LogInit()
	if err != nil {
		panic(err)
	}
	yes := "Yes"
	maybe := "Maybe!"
	logger.Info("This is a info test log")
	logger.Warn("This is a warn test log")
	logger.Error("This is an error test log")
	slogger.Infow("This is also a test log, but a bit different",
		"Does it work?", yes,
		"Is it useful?", maybe)
	slogger.Warnw("This is also a test log, but with a warning",
		"Does it work?", yes,
		"Is it useful?", maybe)
	slogger.Errorw("This is also a test log, but with an error",
		"Does it work?", yes,
		"Is it useful?", maybe)

}

func TestLogInit(t *testing.T) {
	test1()
	test2()
}
