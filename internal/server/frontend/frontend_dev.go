// +build dev

package frontend

import (
	"io/fs"
	"os"
)


func GetFrontendAssets() fs.FS {
	return os.DirFS("internal/server/frontend/homeauto_frontend/dist") 
}
