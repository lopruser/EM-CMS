package model

import (
	"encoding/json"
	"github.com/Etpmls/EM-CMS/src/application"
	em "github.com/Etpmls/Etpmls-Micro"
	"github.com/Etpmls/Etpmls-Micro/define"
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
func (this *Variable) GetAll() interface{} {
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
func (this *Variable) GetAll_Cache() interface{} {
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

	return cache
}
func (this *Variable) GetAll_NotCache() interface{} {
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
		} else {
			em.Cache.SetString(application.Cache_CmsVariableGetAll, string(b), 0)
		}
	}

	return data
}
