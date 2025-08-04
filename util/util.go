package util

import (
	"html/template"
	logging "log"
	"os"
)

var RootDirectory string;

func CheckErrPanic(e error) {
	if e != nil {
		logging.Fatal(e)
	}
}
  
func CheckErrLog(e error) bool{
	if e != nil {
		logging.Println(e)
		return true;
	}
	return false;
}

func Log(v... any) {
	logging.Println(v...)
}

func LogFatal(v... any) {
	logging.Fatal(v...)
}

func GetStaticFilesPath() string {
	return RootDirectory + "/static";
}

func GetTemplatePath(filename string) string {
	return GetStaticFilesPath() + "/templates/" + filename;
}
func GetTemplate(filename string) *template.Template{
	file, err := os.ReadFile(GetTemplatePath(filename))
	CheckErrLog(err)
	tpl, err := template.New(filename).Parse(string(file))
	CheckErrLog(err)
	return tpl
}

func GetDataPath(filename string) string {
	return GetStaticFilesPath() + "/data/" + filename;
}
func GetData(filename string) []byte {
	file, err := os.ReadFile(GetDataPath(filename))
	CheckErrLog(err)
	return file
}