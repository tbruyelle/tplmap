package tplmap

import "html/template"

type TemplateMap map[string]*template.Template

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
