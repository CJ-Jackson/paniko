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
	checked  bool
}

func NewFormHelper(template FormHelperTemplate) *FormHelper {
	return &FormHelper{
		template: template,
		valid:    true,
		checked:  false,
	}
}

func (h *FormHelper) Set(checked, valid bool) string {
	h.checked = checked
	h.valid = valid
	return ""
}

func (h *FormHelper) ValidClass(err error) string {
	if !h.checked && h.valid {
		return ""
	}
	if err == nil {
		return "is-valid"
	}

	return "is-invalid"
}

func (h *FormHelper) CheckErr(err error) template.HTML {
	if !h.valid && err != nil {
		buf := &bytes.Buffer{}
		strs := strings.Split(err.Error(), "\n")
		h.template.errorTemplate.Execute(buf, strs)

		return template.HTML(buf.Bytes())
	}

	return template.HTML("")
}
