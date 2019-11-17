package main

import (
	"logger"
)

func main() {
	var standardLogger = logger.ElevatorLogger()

	standardLogger.GenericInfoMessage("This is a generic message!")
}
