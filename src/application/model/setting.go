package model

import (
	"github.com/Etpmls/EM-CMS/src/application/web"
	em "github.com/Etpmls/Etpmls-Micro"
	"github.com/gin-contrib/multitemplate"
	"html/template"
	"path/filepath"
)

type Setting struct {

}

// Reload Template
// 重载模板
func (this *Setting) ReloadTemplate(r multitemplate.Renderer) error {
	// Get the default template name
	defaultTplName := em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)

	// Get all templates including layout
	allTpl := em.MustListServiceNameKvKey(web.KvServiceTemplate + "/" + defaultTplName)

	// Template classfication
	layouts := make(map[string]string)
	pages := make(map[string]string)
	for k, v := range allTpl {
		if filepath.Base(filepath.Dir(k)) == web.LayoutDirName {
			layouts[k] = v
			continue
		}
		pages[k] = v
	}


	render := r.(multitemplate.Render)
	for k ,_ := range render {
		delete(render, k)
	}

	// tpl = current page + all layout
	for k, v := range pages {
		tpl := template.New(web.BaseTemplate)

		filename := filepath.Base(k)
		dirname := filepath.Base(filepath.Dir(k))
		for lk, lv := range layouts {
			lfilename := filepath.Base(lk)
			// if lfilename = web.BaseTemplate + extension name = app.html No need to new
			if tpl.Name() + filepath.Ext(lfilename) == lfilename {
				_, err := tpl.Parse(lv)
				if err != nil {
					return err
				}
			} else {
				_, err := tpl.New(lfilename).Parse(lv)
				if err != nil {
					em.LogError.OutputFullPath(err)
					return err
				}
			}
		}
		// current page template
		_, err := tpl.New(filename).Parse(v)
		if err != nil {
			return err
		}

		r.Add(dirname + "/" + filename, tpl)
	}
	return nil
}
