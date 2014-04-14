package framework

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type Controller struct {
	request *http.Request

	Response http.ResponseWriter
}

func (c Controller) RenderTemplate(view string, data map[string]interface{}) string {
	t := template.Must(template.New("*").ParseFiles("templates/base.html", fmt.Sprintf("templates/%s.html", view)))
	data["Section"] = view
	var doc bytes.Buffer
	t.ExecuteTemplate(&doc, "base", data)
	return  doc.String()
}
