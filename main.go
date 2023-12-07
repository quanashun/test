package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var fileWriter *os.File
var fileLock *sync.Mutex

func main() {
	bootstrap()
	tick := time.Tick(1 * time.Second)
	loopCount := 0
	for {
		<-tick
		for i := 0; i < 3; i++ {
			go func(i int) {
				fileLock.Lock()
				fileWriter.WriteString(strconv.Itoa(1+loopCount*3+i) + "\n")
				logrus.Debugf("write: %v", strconv.Itoa(1+loopCount*3+i))
				fileLock.Unlock()
			}(i)
		}
		loopCount++
	}
}
func bootstrap() {
	logrus.SetLevel(logrus.DebugLevel)
	dir := filepath.Join("./logs")
	os.MkdirAll(dir, os.ModePerm)
	dir = filepath.Join(dir, "test.log")
	fileWriter, _ = os.OpenFile(dir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileLock = new(sync.Mutex)
}
