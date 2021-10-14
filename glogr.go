/*
Copyright 2019 The logr Authors.

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

// Package glogr implements github.com/go-logr/logr.Logger in terms of
// github.com/golang/glog.
package glogr

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	"github.com/golang/glog"
)

// New returns a logr.Logger which is implemented by glog.
func New() logr.Logger {
	return NewWithOptions(Options{})
}

// NewWithOptions returns a logr.Logger which is implemented by glog.
func NewWithOptions(opts Options) logr.Logger {
	if opts.Depth < 0 {
		opts.Depth = 0
	}

	fopts := funcr.Options{
		LogCaller: funcr.MessageClass(opts.LogCaller),
	}

	gl := &glogger{
		Formatter: funcr.NewFormatter(fopts),
	}

	// For skipping glogger.Info and glogger.Error.
	gl.Formatter.AddCallDepth(opts.Depth + 1)

	return logr.New(gl)
}

// Options carries parameters which influence the way logs are generated.
type Options struct {
	// Depth biases the assumed number of call frames to the "true" caller.
	// This is useful when the calling code calls a function which then calls
	// glogr (e.g. a logging shim to another API).  Values less than zero will
	// be treated as zero.
	Depth int

	// LogCaller tells glogr to add a "caller" key to some or all log lines.
	// The glog implementation always logs this information in its per-line
	// header, whether this option is set or not.
	LogCaller MessageClass

	// TODO: add an option to log the date/time
}

// MessageClass indicates which category or categories of messages to consider.
type MessageClass int

const (
	// None ignores all message classes.
	None MessageClass = iota
	// All considers all message classes.
	All
	// Info only considers info messages.
	Info
	// Error only considers error messages.
	Error
)

type glogger struct {
	funcr.Formatter
}

var _ logr.LogSink = &glogger{}
var _ logr.CallDepthLogSink = &glogger{}

func (l glogger) Enabled(level int) bool {
	return bool(glog.V(glog.Level(level)))
}

func (l glogger) Info(level int, msg string, kvList ...interface{}) {
	prefix, args := l.FormatInfo(level, msg, kvList)
	if prefix != "" {
		args = prefix + ": " + args
	}
	glog.InfoDepth(l.Formatter.GetDepth(), args)
}

func (l glogger) Error(err error, msg string, kvList ...interface{}) {
	prefix, args := l.FormatError(err, msg, kvList)
	if prefix != "" {
		args = prefix + ": " + args
	}
	glog.ErrorDepth(l.Formatter.GetDepth(), args)
}

func (l glogger) WithName(name string) logr.LogSink {
	l.Formatter.AddName(name)
	return &l
}

func (l glogger) WithValues(kvList ...interface{}) logr.LogSink {
	l.Formatter.AddValues(kvList)
	return &l
}

func (l glogger) WithCallDepth(depth int) logr.LogSink {
	l.Formatter.AddCallDepth(depth)
	return &l
}
