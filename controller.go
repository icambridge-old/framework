package framework

import (
	"net/http"
)

type Controller struct {
	request *http.Request

	Response http.ResponseWriter
}
