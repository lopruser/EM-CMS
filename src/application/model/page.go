package model

import (
	"io/ioutil"
	"strings"
)

type Page struct {

}

// Get template path
// 获取模板路径
func (this *Page)PageGetTemplatePath() []string {
	var list []string
	files,_ := ioutil.ReadDir("storage/view/template")

	// Search Files In Dir
	// 在目录中搜索文件
	var f func(path string, list *[]string)
	f = func(path string, list *[]string)  {
		files,_ := ioutil.ReadDir(path)
		for _, v := range files {
			if v.IsDir() {
				if v.Name() != "layout" {
					f(path + "/" + v.Name(), list)
				}
			} else {
				*list = append(*list, path + "/" + v.Name())
			}
		}
	}

	for _,v := range files {
		if v.IsDir() {
			if v.Name() != "layout" {
				f("storage/view/template/" + v.Name(), &list)
			}
		} else {
			if v.Name() != ".gitignore" {
				list = append(list, v.Name())
			}
		}
	}

	// Trim Prefix
	for k, v := range list {
		list[k] = strings.TrimPrefix(v, "storage/view/template/")
	}
	return list
}