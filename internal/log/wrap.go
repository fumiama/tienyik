package log

import (
	"github.com/sirupsen/logrus"

	"github.com/fumiama/tienyik/internal/textio"
)

func Debug(args ...any) {
	if debug {
		args = append([]any{textio.Logger(2)}, args...)
		logrus.Debug(args...)
	}
}

func Debugf(format string, args ...any) {
	if debug {
		args = append([]any{textio.Logger(2)}, args...)
		logrus.Debugf(format, args...)
	}
}

func Debugln(args ...any) {
	if debug {
		args = append([]any{textio.Logger(2)}, args...)
		logrus.Debugln(args...)
	}
}

func Info(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Info(args...)
}

func Infof(format string, args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Infof(format, args...)
}

func Infoln(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Infoln(args...)
}

func Warn(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Warn(args...)
}

func Warnf(format string, args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Warnf(format, args...)
}

func Warnln(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Warnln(args...)
}

func Error(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Error(args...)
}

func Errorf(format string, args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Errorf(format, args...)
}

func Errorln(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Errorln(args...)
}

func Fatal(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Fatalf(format, args...)
}

func Fatalln(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Fatalln(args...)
}

func Panic(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Panic(args...)
}

func Panicf(format string, args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Panicf(format, args...)
}

func Panicln(args ...any) {
	args = append([]any{textio.Logger(2)}, args...)
	logrus.Panicln(args...)
}
