
// +build !dev

package frontend

import (
	"embed"
	"io/fs"
	"fmt"
)

//go:embed homeauto_frontend/dist
var embedFrontend embed.FS

func GetFrontendAssets() fs.FS {
	fmt.Println("wrong one, asshole")
	f, err := fs.Sub(embedFrontend, "homeauto_frontend/dist")
	if err != nil {
		panic(err)
	}

	return f
}
