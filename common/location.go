package common

import (
	"strings"
)

type Location struct {
	SavePath string
	PackageName string
}

func (l *Location) Fill(defaultSavePath, defaultPackageName string) {
	if l.SavePath == "" {
		l.SavePath = defaultSavePath
	}

	if l.PackageName == "" {
		l.PackageName = defaultPackageName
	}
}

func NewContent(l Location) *strings.Builder {
	content := new(strings.Builder)
	content.WriteString("package ")
	content.WriteString(l.PackageName)
	content.WriteByte('\n')
	
	return content
}