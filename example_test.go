/*
Copyright 2021 The logr Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package glogr_test

import (
	"errors"
	"flag"
	"os"
	"sync"

	"github.com/go-logr/glogr"
)

var glogInit = sync.Once{}

func initGlog() {
	glogInit.Do(func() {
		_ = flag.Set("v", "1")
		_ = flag.Set("logtostderr", "true")
		flag.Parse()
	})
	os.Stderr = os.Stdout
}

var errSome = errors.New("some error")

func ExampleNew() {
	initGlog()
	log := glogr.New()
	log.Info("info message with default options")
	log.Error(errSome, "error message with default options")
	log.Info("invalid key", 42, "answer")
	log.Info("missing value", "answer")
	// I1015 08:59:26.952954 2385059 example_test.go:44] "level"=0 "msg"="info message with default options"
	// E1015 08:59:26.953000 2385059 example_test.go:45] "msg"="error message with default options" "error"="some error"
	// I1015 08:59:26.953005 2385059 example_test.go:46] "level"=0 "msg"="invalid key" "<non-string-key: 42>"="answer"
	// I1015 08:59:26.953013 2385059 example_test.go:47] "level"=0 "msg"="missing value" "answer"="<no-value>"
}

func ExampleNew_withName() {
	initGlog()
	log := glogr.New()
	log.WithName("hello").WithName("world").Info("thanks for the fish")
	// I1015 08:59:26.953089 2385059 example_test.go:54] hello/world: "level"=0 "msg"="thanks for the fish"
}

func ExampleNewWithOptions() {
	initGlog()
	log := glogr.NewWithOptions(glogr.Options{LogCaller: glogr.Error})
	log.Info("Info log with LogCaller=Error")
	log.Error(nil, "Error log with LogCaller=All")
	// I1015 09:10:36.754010 2392854 example_test.go:64] "level"=0 "msg"="Info log with LogCaller=Error"
	// E1015 09:10:36.754057 2392854 example_test.go:65] "caller"={"file":"example_test.go","line":65} "msg"="Error log with LogCaller=All" "error"=null
}
