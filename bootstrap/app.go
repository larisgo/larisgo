package bootstrap

import (
	"github.com/larisgo/larisgo/Foundation"
	"os"
	"os/exec"
	"path/filepath"
)

func App() (app *Foundation.App) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	app = Foundation.Application(filepath.Dir(filepath.Dir(path)))

	app.Singleton("test", func() {

	})

	return app
}
