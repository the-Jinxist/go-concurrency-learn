package firstchallenge

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("monkey", &wg)
	wg.Wait()

	if !strings.Contains(msg, "monkey") {
		t.Error("msg doesn't contain `monkey` ")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "monkey"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "monkey") {
		t.Errorf("monkey doesn't exist")
	}
}

func Test_main(t *testing.T) {

	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	Main()

	//Refusing to close the writer caused the test to timeout
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("no universe")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("no cosmos")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("no world")
	}

}
