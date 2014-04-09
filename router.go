package framework

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"strings"
)

type Router struct {
	Routes map[string]Route
}

func (r *Router) ParseXml(src io.Reader) error {
	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	routes := Routes{}
	err = xml.Unmarshal(b, &routes)

	if err != nil {
		return err
	}

	r.Routes = map[string]Route{}
	for _, route := range routes.Routes {
		key := strings.ToLower(route.Path)
		r.Routes[key] = route
	}

	return nil
}

type Routes struct {
	Routes []Route `xml:"route"`
}

type Route struct {
	Controller string `xml:"controller"`
	Action     string `xml:"action"`
	Path       string `xml:"path"`
	Id         string `xml:"id"`
}
