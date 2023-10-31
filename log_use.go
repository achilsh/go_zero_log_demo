package main

import (
	"context"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

func logc_demo() {
	type Data struct {
		A int     `json:"a"`
		B float32 `json:"b"`
	}
	d := Data{
		A: 11,
		B: 1.12,
	}
	//debug level logs
	logc.Debug(context.Background(), "this logc debug log")
	logc.Debugf(context.Background(), "this is demo: %v: %d", "123123", 000111222)
	logc.Debugv(context.Background(), &d)
	logc.Debugw(context.Background(), "is debug w", logc.LogField{Key: "aa", Value: 123})

	//error level logs
	logc.Error(context.Background(), "is error log, 1")
	logc.Errorf(context.Background(), "log2 :%v: %d", "aaa", 123)
	logc.Errorv(context.Background(), d)
	logc.Errorw(context.Background(), "aisdfadfa", logc.Field("aaa", 12321))

	// info level logs
	logc.Info(context.Background(), "is info demo")
	logc.Infof(context.Background(), "info2: %v: %d", "isdfa", 1231231)
	logc.Infov(context.Background(), d)
	logc.Infow(context.Background(), "info3", logc.Field("aa", 123.123))

	//logc.Must()
	//logc.MustSetup(logc.LogConf{})
	//logc.SetLevel(level)

	//其中 level 值为:
	// DebugLevel logs everything
	// DebugLevel uint32 = iota
	// InfoLevel does not include debugs
	// InfoLevel
	// ErrorLevel includes errors, slows, stacks
	// ErrorLevel
	// SevereLevel only log severe messages
	// SevereLevel

	// logc.SetUp(logc.LogConf{})

	//slow level logs
	logc.Slow(context.Background(), "is slow")
	logc.Slowf(context.Background(), "slow: %v:%d", "isfad", 123)
	logc.Slowv(context.Background(), d)
	logc.Sloww(context.Background(), "is slow add field", logc.Field("aa", "bbb"))

	//alert
	logc.Alert(context.Background(), "is alert to error file")
}

func logx_demo() {
	//logx.Severe
	logx.Severe("is serve")
	//logx.Stat
	logx.Stat("is stat...")
	//logx.Stack()
	logx.ErrorStack("is stack...")
}
