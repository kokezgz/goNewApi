package log

import (
	"log"
	"os"
	"time"

	"github.com/Koke/BC/bdservice/utils/setting"
)

const Trace = 99
const Info = 100
const Fatal = 101
const Warning = 102

type ILogger interface {
	WriteLog(str string, t int)
}

type Logger struct {
}

func (l *Logger) WriteLog(str string, t int) {
	settings := setting.GetSettings()
	logDate, fileDate :=
		func() (string, string) {
			now := time.Now()
			return now.Format("01-02-2006 15:04:05"), now.Format("20060102")
		}()

	dir := settings.Log.Folder
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModeDir)
	}

	file := dir + "/" + settings.Log.File + fileDate + settings.Log.Ext
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	defer f.Close()

	if err != nil {
		log.Printf("ERROR: %s \n", err)
	}

	switch t {
	case Info:
		f.WriteString("INFO: " + logDate + ": " + str + "\n")
		log.Printf("INFO: %s \n", str)
	case Fatal:
		f.WriteString("FATAL ERROR : " + logDate + ": " + str + "\n")
		log.Fatal("FATAL ERROR: " + str)
	case Warning:
		f.WriteString("WARNING: " + logDate + ": " + str + "\n")
		log.Printf("WARNING %s \n", str)
	default:
		f.WriteString("TRACE: " + logDate + ": " + str + "\n")
		log.Printf("TRACE: %s \n", str)
	}
}
