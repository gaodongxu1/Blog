package logger

import (
	"github.com/astaxie/beego"
	"github.com/getsentry/raven-go"
)

type SentryLogReport struct {
}

func (sentry *SentryLogReport) ReportDebug(v ...interface{}) {
	beego.Debug(v...)
}

func (sentry *SentryLogReport) ReportInfo(v ...interface{}) {
	beego.Info(v...)

	msg := formatLogMessage(v...)

	raven.CaptureMessage(msg, map[string]string{
		"Server": beego.AppConfig.String("sentry::server"),
		"Level":  "INFO",
	})
}

func (sentry *SentryLogReport) ReportWarning(v ...interface{}) {
	beego.Warn(v...)

	msg := formatLogMessage(v...)

	raven.CaptureMessage(msg, map[string]string{
		"Server": beego.AppConfig.String("sentry::server"),
		"Level":  "WARNING",
	})
}

func (sentry *SentryLogReport) ReportError(err error) {
	beego.Error(err)

	raven.CaptureError(err, map[string]string{
		"Server": beego.AppConfig.String("sentry::server"),
		"Level":  "ERROR",
	})
}
