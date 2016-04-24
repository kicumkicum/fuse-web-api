package api

import (
	"net/http"
)

type IAbstract interface {
	Get(url string) (*http.Response, error)
	Post(url string, bodyType string, body string) (*http.Response, error)
}
