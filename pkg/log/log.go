// Copyright © 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: MIT

package log

import (
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar = logger.Sugar()
}

func Infof(msg string, args ...interface{}) {
	sugar.Infof(msg, args)
}

func Debugf(msg string, args ...interface{}) {
	sugar.Debugf(msg, args)
}

func Errorf(msg string, args ...interface{}) {
	sugar.Errorf(msg, args)
}
