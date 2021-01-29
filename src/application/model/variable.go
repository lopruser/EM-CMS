package model

import (
	"encoding/json"
	"github.com/Etpmls/EM-CMS/src/application"
	em "github.com/Etpmls/Etpmls-Micro/v2"
	"github.com/Etpmls/Etpmls-Micro/v2/define"

	"strings"
	"time"
)

type Variable struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name string	`json:"name"`
	Value string	`json:"value"`
	Remark string `json:"remark"`
}

// Get all variable
// 获取全部变量
func (this *Variable) GetAll() ([]Variable, error) {
	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}

	if strings.ToLower(e) == "true" {
		return this.GetAll_Cache()
	} else {
		return this.GetAll_NotCache()
	}
}
func (this *Variable) GetAll_Cache() ([]Variable, error) {
	j, err := em.Cache.GetString(application.Cache_CmsVariableGetAll)
	if err != nil {
		return this.GetAll_NotCache()
	}

	var cache []Variable
	err = json.Unmarshal([]byte(j), &cache)
	if err != nil {
		em.LogError.OutputFullPath(err)
		return this.GetAll_NotCache()
	}

	em.LogDebug.OutputFullPath("Get CmsVariable from the cache according to UrlPath.")	// debug

	return cache, nil
}
func (this *Variable) GetAll_NotCache() ([]Variable, error) {
	em.LogDebug.OutputSimplePath("Variable cache not found")	// debug

	var data []Variable
	em.DB.Find(&data)

	// If caching is enabled
	// 如果开启了缓存的功能
	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}

	if strings.ToLower(e) == "true" {
		b, err := json.Marshal(data)
		if err != nil {
			em.LogError.OutputFullPath(err)
			return nil, err
		}
		em.Cache.SetString(application.Cache_CmsVariableGetAll, string(b), 0)
	}

	return data, nil
}

func (this *Variable) GetMap() (map[string]interface{}, error) {
	vars, err := this.GetAll()
	if err != nil {
		return nil, err
	}
	var m = make(map[string]interface{})
	for _, v := range vars {
		m[v.Name] = v.Value
	}
	return m, nil
}