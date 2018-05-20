package shared

import (
	"bytes"
	"html/template"
	"strings"
)

type FormHelperTemplate struct {
	errorTemplate *template.Template
}

func NewFormHelperTemplate() FormHelperTemplate {
	return FormHelperTemplate{
		errorTemplate: template.Must(template.New("errorTemplate").Parse(helperErrorTemplate)),
	}
}

type FormHelper struct {
	template FormHelperTemplate
	valid    bool
}

func NewFormHelper(template FormHelperTemplate) *FormHelper {
	return &FormHelper{
		template: template,
		valid:    true,
	}
}

func (h *FormHelper) Valid(value bool) string {
	h.valid = value
	return ""
}

func (h *FormHelper) Check(err error) template.HTML {
	if !h.valid && err != nil {
		buf := &bytes.Buffer{}
		strs := strings.Split(err.Error(), "\n")
		h.template.errorTemplate.Execute(buf, strs)

		return template.HTML(buf.Bytes())
	}

	return template.HTML("")
}