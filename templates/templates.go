package templates

import (
	"github.com/albrow/temple"
	"os"
	"path/filepath"
	"strings"
)

var ( // Split GOPATH using os.PathListSeparator and choose first path in GOPATH as scripts path.
	// Otherwise you would receive an open file error if using multiple paths in your GOPATH.
	gopath            = strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))[0]
	rootTemplatesPath = filepath.Join(gopath, "src", "github.com", "albrow", "temple-example", "tmpl")
	templatesPath     = filepath.Join(rootTemplatesPath, "templates")
	partialsPath      = filepath.Join(rootTemplatesPath, "partials")
	layoutsPath       = filepath.Join(rootTemplatesPath, "layouts")
)

func init() {
	if err := temple.AddAllFiles(templatesPath, partialsPath, layoutsPath); err != nil {
		panic(err)
	}
}
