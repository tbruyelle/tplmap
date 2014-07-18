// Package tplmap is an helper for the native html/template package
package tplmap

import "html/template"

// TemplateMap is simple map where the key is name of the template
// and the value is the html/template
type TemplateMap map[string]*template.Template

// New returns a TemplateMap filled with the templates in parameters.
// The templates have a common parent which is represented by the parent
// parameter, which is a path to the template.
// The template parameter contains the template name as the key (user
// defined) and the path to the template as the value.
func New(parent string, templates map[string]string) TemplateMap {
	tplparent := template.Must(template.ParseFiles(parent))
	t := make(TemplateMap, len(templates))
	for k, v := range templates {
		tpl := template.Must(tplparent.Clone())
		tpl = template.Must(tpl.ParseFiles(v))
		t[k] = tpl
	}
	return t
}
