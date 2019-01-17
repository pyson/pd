package main

import (
	"os"
	"strings"
	"testing"

	// ensure vendor contains it.
	_ "github.com/pingcap/gofail/runtime"
)

func TestRunMain(t *testing.T) {
	var args []string
	for _, arg := range os.Args {
		switch {
		case arg == "DEVEL":
		case strings.HasPrefix(arg, "-test."):
		default:
			args = append(args, arg)
		}
	}

	waitCh := make(chan struct{}, 1)

	os.Args = args
	go func() {
		main()
		close(waitCh)
	}()

	<-waitCh
}