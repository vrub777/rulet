package Controllers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type TemplateHtml struct {
}

func (t *TemplateHtml) showHtml(responseWriter gin.ResponseWriter, page interface{}, nameTemplate string) {
	var template = template.Must(template.ParseGlob("Template/*"))
	err := template.ExecuteTemplate(responseWriter, "header", page)
	err = template.ExecuteTemplate(responseWriter, nameTemplate, page)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateHtml) showHtmlWithHeader(responseWriter gin.ResponseWriter, page interface{},
	nameTemplate string) {
	var template = template.Must(template.ParseGlob("Template/*"))
	//err := template.ExecuteTemplate(responseWriter, "header", page)
	err := template.ExecuteTemplate(responseWriter, nameTemplate, page)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateHtml) showHtmlWithoutHeader(responseWriter gin.ResponseWriter, page interface{},
	nameTemplate string) {
	var template = template.Must(template.ParseGlob("Template/*"))
	err := template.ExecuteTemplate(responseWriter, nameTemplate, page)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateHtml) showHtmlWithoutHeaderR(responseWriter gin.ResponseWriter, model interface{},
	nameTemplate string) {
	var template = template.Must(template.ParseFiles("Template/" + nameTemplate + ".html"))
	err := template.ExecuteTemplate(responseWriter, nameTemplate, model)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateHtml) showHtmlWithoutHeaderRR(responseWriter gin.ResponseWriter, model interface{},
	nameTemplate string) {

	var tt, _ = template.ParseGlob("Template/*")
	err := tt.ExecuteTemplate(responseWriter, nameTemplate, model)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateHtml) showHeader(responseWriter gin.ResponseWriter, page interface{},
	nameTemplate string) {
	var template = template.Must(template.ParseGlob("Template/*"))
	err := template.ExecuteTemplate(responseWriter, "head-header", page)
	err = template.ExecuteTemplate(responseWriter, nameTemplate, page)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
