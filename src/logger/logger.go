package logger

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// ElevatorLogger initializes the standard logger
func ElevatorLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
	genericInfoMessage     = Event{4, "Generic message: %s"}
)

// InvalidArg is a standard error message
func (logger *StandardLogger) InvalidArg(argumentName string) {
	logger.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (logger *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	logger.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (logger *StandardLogger) MissingArg(argumentName string) {
	logger.Errorf(missingArgMessage.message, argumentName)
}

// GenericInfoMessage is a standard info message
func (logger *StandardLogger) GenericInfoMessage(argumentName string) {
	logger.Infof(genericInfoMessage.message, argumentName)
}
