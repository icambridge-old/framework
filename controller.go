package framework

import (
	"net/http"
)

type Controller struct {

	request *http.Request

	response http.ResponseWriter

}
