package gtool

import (
	"fmt"
	"os"
)

var logDirName = "gtool"  //log dir name
var logFileTimeOutDay = 5 //log file will keep 5 days

// set log dir name
func SetLogDirName(name string) {
	logDirName = name
}

// get log dir name
func LogDirName() string {
	return logDirName
}

// set log file time out day
func SetLogTimeOutDay(day int) {
	logFileTimeOutDay = day
}

// get log file time out day
func LogTimeOutDay() int {
	return logFileTimeOutDay
}

//log level assert
func LogA(title string, value interface{}) {
	fileDir := logPathSon("assert")
	logWriteFile("Assert", fileDir, title, value)
}

//log level info
func LogI(title string, value interface{}) {
	fileDir := logPathSon("info")
	logWriteFile("Info", fileDir, title, value)
}

//log level debug
func LogD(title string, value interface{}) {
	fileDir := logPathSon("debug")
	logWriteFile("Debug", fileDir, title, value)
}

//log level warn
func LogW(title string, value interface{}) {
	fileDir := logPathSon("warn")
	logWriteFile("Warn", fileDir, title, value)
}

//log level error
func LogE(title string, value interface{}) {
	fileDir := logPathSon("error")
	logWriteFile("Error", fileDir, title, value)
}

//log
func Log(format LogFormat) {
	fmt.Println(logFormatContent(format))
}

//log sql
func LogSql() {

}

func logBasePath() {

}

//log content format
type LogFormat struct {
	Type    string      //type key
	TypeVal interface{} //type val
	Title   string      //log title
}

//log format content
func logFormatContent(format LogFormat) (res string) {
	if format.TypeVal != nil {
		res = fmt.Sprintf("[%s-%s] %s | %s \n %v", LogDirName(), format.Type, TimeNow(), format.Title, format.TypeVal)
	} else {
		res = fmt.Sprintf("[%s-%s] %s | %s", LogDirName(), format.Type, TimeNow(), format.Title)
	}

	return
}

//write log to file
func logWriteFile(fileName, fileDir, title string, value interface{}) {
	//create dir
	FileDirAutoCreate(fileDir)

	//delete time out file
	logDelTimeOutFile(fileDir)

	path := fileDir + logFileName("")

	//format content struct
	formatContent := LogFormat{
		Type:    fileName,
		TypeVal: value,
		Title:   title,
	}

	content := logFormatContent(formatContent)
	status := FileWrite(path, content)

	if !status {
		Log(formatContent)
	}
}

//time out date
var timeOutDate string

//delete time out log file
func logDelTimeOutFile(fileDir string) {
	toDate := DateYmdDay(-LogTimeOutDay())

	if timeOutDate == toDate {
		return
	}

	fileName := logFileName("")
	pathName := fileDir + fileName

	if FileDirIsExist(pathName) { //if file exist and remove it
		err := os.Remove(pathName)

		if err == nil {
			timeOutDate = toDate //set time out date
		}
	}
}

//log file name
func logFileName(name string) string {
	if name == "" {
		name = DateToday()
	}

	return fmt.Sprintf("%s.log", name)
}

//log base path
func logPath() string {
	return fmt.Sprintf("%s/%s/", PathCurrent(), LogDirName())
}

func logPathSon(name string) string {
	return logPath() + name + "/"
}
