package raft

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Debugging
/*const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}*/

type logTopic string

const (
	dInfo logTopic = "INFO"
)

var debugStart time.Time
var debugVerbosity int

//通过检查环境变量获取日志等级
func getVerbosity() int {
	v := os.Getenv("VERBOSE")
	level := 0
	if v != "" {
		var err error
		level, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Invalid verbosity %v", v)
		}
	}
	return level
}

func init() {
	debugVerbosity = getVerbosity()
	debugStart = time.Now()

	// 自定义日志的抬头信息，log库的使用方式可以参考如下的链接
	// https://www.flysnow.org/2017/05/06/go-in-action-go-log
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

func Debug(topic logTopic, format string, a ...interface{}) {
	if debugVerbosity >= 1 {
		time := time.Since(debugStart).Microseconds()
		time /= 100
		prefix := fmt.Sprint("%06d %v ", time, string(topic))
		format = prefix + format
		log.Printf(format, a...)
	}
}

func DebugGetInfo(rf *Raft) {
	Debug(dInfo, "Server%d Term%d State: %s Log:%v", rf.me, rf.currentTerm,
		State2String[rf.state], rf.log)
}
