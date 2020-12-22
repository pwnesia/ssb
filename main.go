package main

import "ktbs.dev/ssb/internal/runner"

func main() {
	opt := runner.Parse()
	runner.New(opt)
}
