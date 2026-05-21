package tmpl

import (
	"embed"
)

//go:embed *
var tmplFiles embed.FS

func GetFS() (embed.FS) {
	return tmplFiles
}
