package static

import (
	"net/http"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
)

const version = "0.0.1"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(fs *assetfs.AssetFS) *binaryFileSystem {
	return &binaryFileSystem{fs}
}

func ServeRoot(urlPrefix string, fs *assetfs.AssetFS) echo.MiddlewareFunc {
	return Serve(urlPrefix, BinaryFileSystem(fs))
}

// Serve Static returns a middleware handler that serves static files in the given directory.
func Serve(urlPrefix string, fs ServeFileSystem) echo.MiddlewareFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(before echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := before(c)
			if err != nil && err.Error() != http.StatusText(http.StatusNotFound) {
				return err
			}

			w, r := c.Response(), c.Request()
			if fs.Exists(urlPrefix, r.URL.Path) {
				fileserver.ServeHTTP(w, r)
				return nil
			}
			return err
		}
	}
}
