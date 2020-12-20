package model

import "time"

type Grader interface {
	Run(path string) (string, int64, error)
	RunCase(path string, inPath string, outPath string, execTime time.Duration) error
	RunAllCases(path string, timePerCase time.Duration, cases ...string) error
}
