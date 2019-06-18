package Bootstrap

import (
	"github.com/larisgo/framework/Foundation"
	"github.com/larisgo/framework/Foundation/Http"
	"os"
	"os/exec"
	"path/filepath"
)

func App() (app *Foundation.Application) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	app = Foundation.NewApplication(filepath.Dir(path))

	app.Singleton("kernel", Http.NewKernel())

	return app
}
