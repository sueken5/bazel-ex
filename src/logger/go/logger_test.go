package logger_test

import (
	logger "github.com/sueken5/bazel-ex/src/logger/go"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	testCases := map[string]struct {
		input string
	} {
		"ok": {
			input:"hello",
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			l := logger.NewLogger()
			l.Info(tc.input)
		})
	}
}
