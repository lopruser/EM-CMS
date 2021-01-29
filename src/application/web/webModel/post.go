package webModel

import (
	"encoding/json"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/model"
	em "github.com/Etpmls/Etpmls-Micro/v2"
	"github.com/Etpmls/Etpmls-Micro/v2/define"

	"github.com/go-redis/redis/v8"

	"strconv"
	"strings"
)

type post struct {

}

func NewPost() *post {
	return &post{}
}

// Get Post according to url_path [relationship:thumbnail]
// CmsPost	[relationship:thumbnail]
func (this *post) GetByUrlPath(url_path string) (model.Post, error) {
	s, _ := em.Kv.ReadKey(define.KvCacheEnable)
	if strings.ToLower(s) == "true" {
		return this.getByUrlPath_Cache(url_path, strings.ToLower(s) == "true")
	} else {
		return this.getByUrlPath_NoCache(url_path, strings.ToLower(s) == "true")
	}
}
func (this *post) getByUrlPath_Cache(url_path string, enableCache bool) (model.Post, error) {
	s, err := em.Cache.GetHash(application.Cache_CmsPostGetByUrlPath, url_path)
	if err != nil {
		if err == redis.Nil {
			return this.getByUrlPath_NoCache(url_path, enableCache)
		}
		em.LogError.OutputFullPath(err)
		return model.Post{}, err
	}

	var p model.Post
	err = json.Unmarshal([]byte(s), &p)
	if err != nil {
		em.LogError.OutputFullPath(err)
		return model.Post{}, err
	}
	em.LogDebug.OutputFullPath("Get Post from the cache according to UrlPath.")	// debug

	return p, nil
}
func (this *post) getByUrlPath_NoCache(url_path string, enableCache bool) (model.Post, error) {
	em.LogDebug.OutputFullPath("Post cache not found")	// debug

	var p model.Post
	result := em.DB.Where("url_path = ?", url_path).Order("sort desc").First(&p)
	if result.Error != nil {
		em.LogDebug.OutputFullPath(result.Error)	// DEBUG
		return model.Post{}, result.Error
	}

	// If the cache function is turned on
	// 如果开启缓存功能
	if enableCache {
		b, err := json.Marshal(p)
		if err != nil  {
			em.LogError.OutputFullPath(err)
			return model.Post{}, err
		}
		var m = make(map[string]string)
		m[url_path] = string(b)
		em.Cache.SetHash(application.Cache_CmsPostGetByUrlPath, m)
	}

	return p, nil
}


// Get an article with the highest sort attribute according to the category ID
// 根据分类ID获取一个sort属性最高的文章
func (this *post) GetOnePostByCategoryIdWithHighestSort(categoryId uint) (model.Post, error) {
	s, _ := em.Kv.ReadKey(define.KvCacheEnable)
	if strings.ToLower(s) == "true" {
		return this.getOnePostByCategoryIdWithHighestSort_Cache(categoryId, strings.ToLower(s) == "true")
	} else {
		return this.getOnePostByCategoryIdWithHighestSort_NoCache(categoryId, strings.ToLower(s) == "true")
	}
}
func (this *post) getOnePostByCategoryIdWithHighestSort_Cache(categoryId uint, enableCache bool) (model.Post, error) {
	s, err := em.Cache.GetHash(application.Cache_CmsGetOnePostByCategoryIdWithHighestSort, strconv.Itoa(int(categoryId)))
	if err != nil {
		if err == redis.Nil {
			return this.getOnePostByCategoryIdWithHighestSort_NoCache(categoryId, enableCache)
		}
		em.LogError.OutputFullPath(err)
		return model.Post{}, err
	}
	var p model.Post
	err = json.Unmarshal([]byte(s), &p)
	if err != nil {
		em.LogError.Output("CmsPost getOnePostByCategoryIdWithHighestSort_Cache:" + err.Error())
		return model.Post{}, err
	}

	return p, nil
}
func (this *post) getOnePostByCategoryIdWithHighestSort_NoCache(categoryId uint, enableCache bool) (model.Post, error) {
	var p model.Post
	result := em.DB.Where("category_id = ?", categoryId).Order("sort desc").First(&p)
	if result.Error != nil {
		return model.Post{}, result.Error
	}

	// If the cache function is turned on
	// 如果开启缓存功能
	if enableCache {
		b, err := json.Marshal(p)
		if err != nil  {
			em.LogError.OutputFullPath(err)
			return model.Post{}, err
		}
		var m = make(map[string]string)
		m[strconv.Itoa(int(categoryId))] = string(b)
		em.Cache.SetHash(application.Cache_CmsGetOnePostByCategoryIdWithHighestSort, m)
	}

	return p, nil
}
