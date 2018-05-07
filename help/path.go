package help

import (
	"os"
	"path/filepath"
	"strings"
)

//GetBinaryPath will return the location of the binary or the project in go run mode.
func GetBinaryPath() (path string) {
	goPath := os.Getenv("GOPATH")

	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	if strings.Contains(exPath, "go-build") { //This means we are running in go run and need to use the goPath
		path = goPath + PathSeparator + "src" + PathSeparator + "github.com" + PathSeparator + "DanielRenne" + PathSeparator + "nsqManager" + PathSeparator + ""
	} else {
		path = exPath + PathSeparator
	}
	return
}

//IsGoRun returns true if the binary is run from a go run command.
func IsGoRun() bool {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	if strings.Contains(exPath, "go-build") {
		return true
	}
	return false

}
