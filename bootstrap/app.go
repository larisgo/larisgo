package bootstrap

import (
	"fmt"
	"github.com/larisgo/larisgo/Foundation"
	"os"
	"os/exec"
	"path/filepath"
)

func GetCurPath() string {
	file, err := exec.LookPath(os.Args[0])
	fmt.Println(file)
	if err != nil {
		panic(err)
	}

	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}

	return filepath.Dir(path)
}

func App() (app *Foundation.App) {
	path, err := filepath.Abs(GetCurPath() + "/../")
	if err != nil {
		panic(err)
	}
	return Foundation.Application(path)
}
