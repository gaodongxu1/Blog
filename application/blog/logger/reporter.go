package logger

import (
	"fmt"
	"strings"
)

type LogReporterInterface interface {
	ReportDebug(v ...interface{})
	ReportInfo(v ...interface{})
	ReportWarning(v ...interface{})
	ReportError(error)
}

type LogReport struct {
	provider LogReporterInterface
}

var GlobalLogReporter *LogReport

func init() {
	GlobalLogReporter = &LogReport{
		provider: &SentryLogReport{},
	}
}

func (manager *LogReport) Debug(v ...interface{}) {
	manager.provider.ReportDebug(v...)
}

func (manager *LogReport) Info(v ...interface{}) {
	manager.provider.ReportInfo(v...)
}

func (manager *LogReport) Warn(v ...interface{}) {
	manager.provider.ReportWarning(v...)
}

func (manager *LogReport) Error(err error) {
	manager.provider.ReportError(err)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}

func formatLogMessage(v ...interface{}) string {
	format := generateFmtStr(len(v))
	msg := fmt.Sprintf(format, v...)

	return msg
}
