package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"regexp"
	"runtime"
	"strconv"
)

const reg = `\.[a-zA-Z_]+`

//todo? it is possible to do this is asynchronous way? i know that simply answer is yes(channel and forget about message :D), but how to do this in proper way?

var l *logrus.Logger

func Init(level string) {
	l = logrus.New()
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		// default level
		parseLevel = logrus.InfoLevel
	}
	l.SetLevel(parseLevel)
	l.SetFormatter(&logrus.JSONFormatter{})
}

func Infof(ctx context.Context, format string, args ...any) {
	l.WithFields(addFields(ctx)).Infof(format, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	l.WithFields(addFields(ctx)).Errorf(format, args...)
}

func Debugf(ctx context.Context, format string, args ...any) {
	l.WithFields(addFields(ctx)).Debugf(format, args...)
}

func Warnf(ctx context.Context, format string, args ...any) {
	l.WithFields(addFields(ctx)).Warnf(format, args...)
}

func Fatalf(ctx context.Context, format string, args ...any) {
	l.WithFields(addFields(ctx)).Fatalf(format, args...)
}

func addFields(ctx context.Context) logrus.Fields {

	fields := basicFields()

	if ctx == nil {
		return fields
	}

	//TODO add fields from context
	// example: requestId, userId, etc

	return fields
}

func basicFields() logrus.Fields {

	file, numberLine, funcName := "unknown", "unknown", "unknown"
	if pc, f, nl, ok := runtime.Caller(3); ok {
		funcName = findFunctionName(pc)
		file = f
		numberLine = strconv.Itoa(nl)
	}

	return logrus.Fields{
		"file": file,
		"line": numberLine,
		"func": funcName,
	}
}

func findFunctionName(pc uintptr) string {
	funcName := "unknown"

	funcNameBytes := []byte(runtime.FuncForPC(pc).Name())
	matchedBytes := regexp.MustCompile(reg).Find(funcNameBytes)
	funcNameString := string(matchedBytes)

	if len(funcNameString) > 0 {
		funcName = funcNameString[1:]
	}

	return funcName
}
