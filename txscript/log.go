// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2015-2019 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

import (
	"github.com/Utopia/slog"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
// The default amount of logging is none.
var log = slog.Disabled

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger slog.Logger) {
	log = logger
}

// LogClosure is a closure that can be printed with %v to be used to
// generate expensive-to-create data for a detailed log level and avoid doing
// the work if the data isn't printed.
type logClosure func() string

func (c logClosure) String() string {
	return c()
}

func newLogClosure(c func() string) logClosure {
	return logClosure(c)
}
