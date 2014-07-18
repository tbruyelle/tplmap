// Package tplmap is an helper for the native html/template package.
package tplmap

import "html/template"

// TemplateMap is simple map where the key is name of the template
// and the value is the html/template.
type TemplateMap map[string]*template.Template

// New returns a TemplateMap filled with the templates in parameters.
// The templates have a common parent which is represented by the parent
// parameter, which is a path to a common template.
// The templates parameter contains multiple templates where the template
// name is the key (user defined) and the path to the template as the value.
//
// Example:
//   // Create the TemplateMap
//   templates = tplmap.New("views/layout.html", map[string]string{
//		"main":  "views/main.html",
//		"login": "views/login.html",
//	})
//
//   // Use it
//   templates["main"].Execute(w, data)
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
