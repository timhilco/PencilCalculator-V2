package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

const (
	LOG_LEVEL_INFO int = iota
	LOG_LEVEL_DEBUG
)

// HilcoLogger with zeroLog
type HilcoLogger struct {
	logger *zerolog.Logger
}

// NewMultiWithFile with file
func NewMultiWithFile(logLevel int, isAppend bool, iFileName string, createConsole bool) zerolog.Logger {
	var zlogLevel zerolog.Level
	switch logLevel {
	case LOG_LEVEL_INFO:
		zlogLevel = zerolog.InfoLevel
	case LOG_LEVEL_DEBUG:
		zlogLevel = zerolog.DebugLevel
	default:
		zlogLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(zlogLevel)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	//file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	var file *os.File
	var err error
	createFile := false
	if iFileName != "" {
		createFile = true
		fileName := "Consumer-Group-X_logs.txt"
		fileName = iFileName

		flag := os.O_CREATE
		if isAppend {
			flag = os.O_APPEND
		}
		file, err = os.OpenFile(fileName, flag|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Error: ZeroLogger file error: %s\n", err.Error())
		}
	}
	//multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout, file)
	var multi zerolog.LevelWriter
	if createFile {
		if createConsole {
			multi = zerolog.MultiLevelWriter(consoleWriter, file)
		} else {
			multi = zerolog.MultiLevelWriter(file)
		}
	} else {
		if createConsole {
			multi = zerolog.MultiLevelWriter(consoleWriter, file)
		} else {
			consoleWriter = zerolog.ConsoleWriter{Out: io.Discard}
			multi = zerolog.MultiLevelWriter(consoleWriter)
		}

	}

	logger := zerolog.New(multi).With().Timestamp().Logger()

	return logger
}

// New for zerolog
func New(isDebug bool) zerolog.Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return logger
}

// NewConsole with Console
func NewConsole(isDebug bool) zerolog.Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return logger
}

// Output duplicates the global logger and sets w as its output.
func (l *HilcoLogger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}

// Info is a standard error message
func (l *HilcoLogger) Info(message string) {
	l.logger.Info().Msg(message)
}

// Debug is a standard error message
func (l *HilcoLogger) Debug(message string) {
	l.logger.Debug().Msg(message)
}

// Fatal is a standard error message
func (l *HilcoLogger) Fatal(message string) {
	l.logger.Fatal().Msg(message)
}

// Error is a standard error message
func (l *HilcoLogger) Error(message string) {
	l.logger.Error().Msg(message)
}
