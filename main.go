package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fghwett/juejin/config"
	"github.com/fghwett/juejin/notify"
	"github.com/fghwett/juejin/task"
)

var path = flag.String("path", "./config.yml", "配置文件地址")

func main() {
	flag.Parse()

	conf, err := config.Init(*path)
	if err != nil {
		fmt.Printf("读取配置文件失败 err: %s", err)
		os.Exit(-1)
	}

	t := task.New(conf.Config)
	t.Do()

	if err := notify.Send(conf.ServerChan.SecretKey, "掘金保活任务", t.GetResult()); err != nil {
		log.Printf("通知发送失败 %s\n", err)
	}
}
