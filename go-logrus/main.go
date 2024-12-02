package main

import log "github.com/sirupsen/logrus"

func main() {
	//设置log级别
	log.SetLevel(log.TraceLevel)
	//定位代码调用
	log.SetReportCaller(true)
	//设置log格式
	log.SetFormatter(&log.TextFormatter{
		//定义日期格式
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 指定字段
	log.WithFields(log.Fields{
		"name": "test",
	}).Infof("to do %v", "log")

	log.Trace("trace")
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Fatal("fatal")
	log.Panic("panic")
}
