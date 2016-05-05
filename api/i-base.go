package api

import (
	"../fs/"
)
type IBase interface {
	ReadContent(path string) (fs.IContent, error)
}
