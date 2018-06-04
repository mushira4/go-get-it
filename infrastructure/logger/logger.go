package logger

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

//APP name 
const APP = "go-get-it"

//BLUE color
const BLUE = "\033[1;34m"

//YELLOW color
const YELLOW = "\033[1;33m"

//RED color
const RED = "\033[0;31m"

//END of the color
const END_COLOR = "\033[0m"

func init() {
	log.SetFlags(0)
}

//SimpleLogger creates a simple non-structured logger
func EmptyLogger() *Logger {
	logger := Logger{}
	return &logger
}

//Logger creates logger
func SimpleLogger() *Logger {
	logger := Logger{}
	logger.elements = make(map[string]string)
	return &logger
}

// Append initialize the logger insetion
func Append(key string, value string) *Logger {
	logger := Logger{}
	logger.elements = make(map[string]string)
	logger.elements[key] = value
	return &logger
}

//Logger is responsible for keeping the structured logs elements
type Logger struct {
	elements map[string]string
}

//Append adds a new structured to be logged
func (logger *Logger) Append(key string, value string) *Logger {
	logger.elements[key] = value
	return logger
}

//Info prints a simple message
func (logger *Logger) Info() {
	log.Printf(BLUE + fmt.Sprintf("app=%s ", APP) + logger.writeLog("INFO") + END_COLOR)
}

//Warn prints a simple message, with a higher danger than info
func (logger *Logger) Warn() {
	log.Printf(YELLOW + fmt.Sprintf("app=%s ", APP) + logger.writeLog("WARN") + END_COLOR)
}

//Error prints a simple message, with a higher danger than Warn
func (logger *Logger) Error(anError ...error) {
	log.Printf(RED + fmt.Sprintf("app=%s ", APP) + logger.writeErrorLog("ERROR", anError) + END_COLOR)
}

//Panic prints a simple message, with a the same danger than Error and kill the current go routine
func (logger *Logger) Panic(anError ...error) {
	log.Panic(RED+fmt.Sprintf("app=%s ", APP)+logger.writeErrorLog("FATAL", anError)+END_COLOR, anError)
}

//Fatal prints a simple message, with a higher danger than Error and kill the application
func (logger *Logger) Fatal(anError ...error) {
	log.Fatal(RED+fmt.Sprintf("app=%s ", APP)+logger.writeErrorLog("FATAL", anError)+END_COLOR, anError)
}

func (logger *Logger) writeLog(messageLevel string) string {
	return fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05.000"), messageLevel, logger.writeElements())
}

func (logger *Logger) writeErrorLog(messageLevel string, anError []error) string {
	if anError == nil {
		return fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05.000"), messageLevel, logger.writeElements())
	} else {
		return fmt.Sprintf("%s %s %s %s", time.Now().Format("2006-01-02 15:04:05.000"), messageLevel, logger.writeElements(), anError)
	}
}

func (logger *Logger) writeElements() string {
	var buffer bytes.Buffer

	for k := range logger.elements {
		buffer.WriteString(k + "=" + logger.elements[k] + " ")
	}

	return buffer.String()
}
