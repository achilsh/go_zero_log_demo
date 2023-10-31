package main

import (
	"flag"
	"fmt"

	// "os"

	// "gopkg.in/yaml.v3"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	logCfgFile *string = flag.String("f", "cfg.yml", "-f cfg.ymal")
)

type DemoConfig struct {
	//自定义的配置文件，加载go-zero日志。
	Log logx.LogConf
}

var Cfg DemoConfig

func init_log() {
	flag.Parse()

	if logCfgFile == nil || len(*logCfgFile) <= 0 {
		panic("log file is empty, check file is ready")
		return
	}

	// data, e := os.ReadFile(*logCfgFile)
	// if e != nil {
	// 	fmt.Println("read file fail, e: ", e)
	// 	//panic("")
	// 	return
	// }

	// if e := yaml.Unmarshal(data, &cfg); e != nil {
	// 	fmt.Println("parse data from ymal fail, e: ", e)
	// 	//painc("")
	// 	return
	// }

	conf.MustLoad(*logCfgFile, &Cfg, conf.UseEnv())
	if e := logc.SetUp(Cfg.Log); e != nil {
		panic(fmt.Sprintf("init log config from file fail, e: %v", e))
	}
}

func destory() {
	logc.Close()
}
