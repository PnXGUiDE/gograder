package java

import (
	"log"
	"os/exec"
	"time"
	"errors"
	"bytes"
	"io/ioutil"
	"fmt"
)

// Run :
func Run(path string) (string, int64, error) {
	cmd := exec.Command("java", path)
	
	var buf bytes.Buffer
	cmd.Stdout = &buf

	cmd.Start()
	start := time.Now().UnixNano()

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(2 * time.Second)

	select {
	case <-timeout:
		cmd.Process.Kill()
		return "", 2000, errors.New("Command Timed Out!")
	case err := <-done:
		if err != nil {
			log.Println("Non-zero exit code:", err)
		}
		return buf.String(), (time.Now().UnixNano() - start) / int64(time.Millisecond), nil
	}
}

// RunCase :
func RunCase(path string, inPath string, outPath string, execTime time.Duration) error {
	cmd := exec.Command("java", path, "<", inPath)
	
	var buf bytes.Buffer
	cmd.Stdout = &buf

	out, err := ioutil.ReadFile(outPath)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(execTime * time.Second)

	select {
	case <-timeout:
		cmd.Process.Kill()
		return errors.New("Command timed out")
	case err := <-done:
		if err != nil {
			log.Println("Non-zero exit code:", err)
		}
		if string(out) == buf.String() {
			return nil
		}
		return errors.New("Output mismatch")
	}
}

// RunAllCases :
func RunAllCases(path string, timePerCase time.Duration, cases ...string) error {
	if len(cases) % 2 != 0 {
		return errors.New("Test cases are invalid")
	}

	for i := 0; i < len(cases); i += 2 {
		err := RunCase(path, cases[i], cases[i + 1], timePerCase)
		if err != nil {
			msg := fmt.Sprintf("Wrong output on test case #%d", (i / 2) + 1)
			return errors.New(msg)
		}
	}

	return nil
}
