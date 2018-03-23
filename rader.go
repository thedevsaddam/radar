// Copyright @2018 Saddam Hossain.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package radar help to debug nested function call and Trace current file/line
package radar

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Beam print the current file, function and line information
func Beam(skip ...int) (err error) {
	skp := 1
	if skip != nil && skip[0] > 0 {
		skp = skip[0]
	}
	pc, file, line, _ := runtime.Caller(skp)
	_, err = fmt.Fprintf(os.Stdout,
		" -> %s: %s(%d)\n",
		file, filepath.Base(runtime.FuncForPC(pc).Name()), line)
	return
}

type caller struct {
	File, Func string
	Line       int
}

// Trace print the current file, function and line information for all program counter
func Trace() (err error) {
	clrs := []caller{}
	for depth := 0; ; depth++ {
		pc, src, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		clrs = append(clrs, caller{
			File: src,
			Func: filepath.Base(runtime.FuncForPC(pc).Name()),
			Line: line,
		})
	}

	for i := 1; i < len(clrs)-2; i++ {
		_, err = fmt.Fprintf(os.Stdout, " -> %d: %s: %s(%d)\n", i, clrs[i].File, clrs[i].Func, clrs[i].Line)
	}

	return
}
