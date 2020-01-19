module github.com/sueken5/bazel-ex/src/echo

go 1.13

replace github.com/sueken5/bazel-ex/src/logger/go => ../logger/go

require github.com/sueken5/bazel-ex/src/logger/go v0.0.0-20200119053734-cc98b4b3561b // indirect
