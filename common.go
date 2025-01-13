package otelkit

import (
	"runtime"
	"strings"
)

func packageName(skip int) string {
	var (
		name    string
		callers [1]uintptr
	)

	if runtime.Callers(skip+2, callers[:]) > 0 {
		frames := runtime.CallersFrames(callers[:])
		frame, _ := frames.Next()

		if frame.Function != "" {
			lastSlash := strings.LastIndex(frame.Function, "/")
			if lastSlash < 0 {
				lastSlash = 0
			}

			if funcStart := strings.Index(frame.Function[lastSlash:], "."); funcStart >= 0 {
				name = frame.Function[:lastSlash+funcStart]
			} else {
				name = frame.Function
			}
		}
	}
	return name
}
