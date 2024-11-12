package logging

import (
	"encoding/json"
	"eppv2/internal/constants"
	_ "fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogInit(subsytem string, module string) (*zap.Logger, *zap.SugaredLogger, error) {
	initFields := make(map[string]interface{})
	initFields["subsytem"] = subsytem
	initFields["module"] = module

	rawJSONCfg := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
		"timeKey": "time",
		"callerKey": "caller",
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSONCfg, &cfg); err != nil {
		log.Fatalf("can't initialize json config in logger: %v", err)
		return nil, nil, err
	}
	cfg.InitialFields = initFields

	// fetch log level from env variables and update the config
	logLevel := os.Getenv(constants.LOG_LEVEL)
	if logLevel == "" {
		logLevel = "info"
	}

	atomicLogLevel, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		// if there is any error while parsing the logLevel, create a default atomicLogLevel
		// default atomicLogLevel is `info`
		log.Printf("Error occured while creating atomic log level: %v, Creating atomicLogLevel with default settings [info]", err)
		atomicLogLevel = zap.NewAtomicLevel()
	}
	cfg.Level = atomicLogLevel

	// Discuss and decide on the TimeEncoder type
	//cfg.EncoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
		return nil, nil, err
	}
	defer logger.Sync()
	slogger := logger.Sugar()
	logger.Info("logger construction succeeded")
	return logger, slogger, nil
}

// NOTE: Don't know the exact usecase of ReplaceGlobals but will see in future if need arises
func CommonLogInit(subsystem string, module string) (*zap.Logger, *zap.SugaredLogger, error) {
	logger, _, err := LogInit(subsystem, module)
	if err != nil {
		return nil, nil, err
	}

	undo := zap.ReplaceGlobals(logger)
	defer undo()
	return logger, logger.Sugar(), err
}
