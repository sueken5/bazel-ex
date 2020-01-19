package main

import logger "github.com/sueken5/bazel-ex/src/logger/go"

func main() {
	l := logger.NewLogger()
	l.Info("echo")
}