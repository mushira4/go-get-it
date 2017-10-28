package logger

import (
	"errors"
	"testing"
)

var checkMark = "\u2713"
var ballotX = "\u2717"

func TestAppend(t *testing.T) {
	logger := Append("keyA", "valueA").Append("keyB", "valueB").Append("keyC", "valueC").Append("keyD", "valueD")

	t.Log("Check the correct number of elements")
	{
		if len(logger.elements) == 4 {
			t.Logf("Correct number of elements %s", checkMark)
		} else {
			t.Fatalf("Wrong number of elements. Elements=%s %s", len(logger.elements), ballotX)
		}
	}

	t.Log("Check the elements")
	{
		if logger.elements["keyA"] == "valueA" && logger.elements["keyB"] == "valueB" && logger.elements["keyC"] == "valueC" && logger.elements["keyD"] == "valueD" {
			t.Logf("Correct elements values %s", checkMark)
		} else {
			t.Fatalf("Incorrect elements  %s", ballotX)
		}
	}
}

func TestLogInfo(t *testing.T) {
	Append("keyA", "valueA").Append("KeyB", "valueB").Append("keyC", "valueC").Append("keyD", "valueD").Info()
}

func TestLogWarn(t *testing.T) {
	Append("keyA", "valueA").Append("KeyB", "valueB").Append("keyC", "valueC").Append("keyD", "valueD").Warn()
}

func TestLogError(t *testing.T) {
	Append("keyA", "valueA").Append("KeyB", "valueB").Append("keyC", "valueC").Append("keyD", "valueD").Error()
}

func TestLogErrorWithContent(t *testing.T) {
	Append("keyA", "valueA").Append("KeyB", "valueB").Append("keyC", "valueC").Append("keyD", "valueD").Error(errors.New("My Error"))
}
