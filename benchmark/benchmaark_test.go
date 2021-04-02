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

package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/go-logr/glogr"
	"github.com/go-logr/logr"
	"github.com/golang/glog"
)

func init() {
	flag.Set("v", "1")
	flag.Set("logtostderr", "true")
	os.Stderr, _ = os.Open("/dev/null")
}

func BenchmarkGlogInfoOneArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		glog.Info("this is", "a", "string")
	}
}

func BenchmarkInfoOneArg(b *testing.B) {
	var log logr.Logger = glogr.New()

	for i := 0; i < b.N; i++ {
		log.Info("this is", "a", "string")
	}
}

func BenchmarkGlogInfoSeveralArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		glog.Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

func BenchmarkInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = glogr.New()

	for i := 0; i < b.N; i++ {
		log.Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

func BenchmarkGlogInfoV0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		glog.V(0).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

func BenchmarkV0Info(b *testing.B) {
	var log logr.Logger = glogr.New()

	for i := 0; i < b.N; i++ {
		log.V(0).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

func BenchmarkGlogV9Info(b *testing.B) {
	for i := 0; i < b.N; i++ {
		glog.V(9).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

func BenchmarkV9Info(b *testing.B) {
	var log logr.Logger = glogr.New()

	for i := 0; i < b.N; i++ {
		log.V(9).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

func BenchmarkError(b *testing.B) {
	var log logr.Logger = glogr.New()

	err := fmt.Errorf("error message")
	for i := 0; i < b.N; i++ {
		log.Error(err, "multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

func BenchmarkWithValues(b *testing.B) {
	var log logr.Logger = glogr.New()

	for i := 0; i < b.N; i++ {
		l := log.WithValues("k1", "v1", "k2", "v2")
		_ = l
	}
}

func BenchmarkWithName(b *testing.B) {
	var log logr.Logger = glogr.New()

	for i := 0; i < b.N; i++ {
		l := log.WithName("name")
		_ = l
	}
}
