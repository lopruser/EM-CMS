package model

import (
	"github.com/Etpmls/EM-CMS/src/application/web"
	em "github.com/Etpmls/Etpmls-Micro/v2"

	"path/filepath"
)

type Page struct {

}

// Get template path
// 获取模板路径
func (this *Page)PageGetTemplatePath() []string {
	// Get the default template name
	defaultTplName := em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)
	// Get all templates including layout
	allTpl := em.MustListServiceNameKvKey(web.KvServiceTemplate + "/" + defaultTplName)

	var tpls []string
	for k, _ := range allTpl {
		if filepath.Base(filepath.Dir(k)) == web.LayoutDirName {
			continue
		}
		tpls = append(tpls, filepath.Base(filepath.Dir(k)) + "/" + filepath.Base(k))
	}
	return tpls
}

