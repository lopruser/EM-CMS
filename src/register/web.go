package register

import (
	"encoding/json"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/model"
	"github.com/Etpmls/EM-CMS/src/application/web/webService"
	em "github.com/Etpmls/Etpmls-Micro"
	"github.com/Etpmls/Etpmls-Micro/define"
	em_library "github.com/Etpmls/Etpmls-Micro/library"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"strconv"
)


func initRoute(r *gin.Engine) {
	var p webService.ServicePage
	r.GET("/", p.PageShowIndex)
	r.GET("/category/:url_path", p.PageShowCategory)
	r.GET("/post/:url_path", p.PageShowPost)
	r.GET("/search", p.PageShowSearch)

	webRoute := r.Group("/data")
	{
		webRoute.GET("/get/index", p.PageGetIndex)
		webRoute.POST("/get/category/page", p.PageGetCategroyPage)
		webRoute.POST("/get/category/list", p.PageGetCategroyList)
		webRoute.POST("/get/post", p.PageGetPost)
		webRoute.POST("/get/search", p.PageGetSearch)

		webRoute.GET("/get/category/tree", p.PageGetCategoryTree)
		webRoute.GET("/get/category/children/:id", p.PageGetCategoryChildren)
		webRoute.POST("/get/category/treeByUrlPath", p.PageGetCategoryTreeByUrlPath)
		webRoute.GET("/get/category/byCategoryId/:id", p.PageGetCategoryByCategoryId)
		webRoute.POST("/get/category/manyByCategoryIds", p.PageGetCategoryManyByCategoryIds)

		webRoute.POST("/get/post/manyByIds", p.PageGetPostManyByIds)
		webRoute.POST("/get/post/byCategoryUrlPath", p.PageGetPostByCategoryUrlPath)
		webRoute.POST("/get/post/byCategoryId", p.PageGetPostByCategoryId)
	}
}

func RunWebServer()  {
	gin.SetMode(gin.ReleaseMode)

	// Get port
	p := em.MustGetServiceIdKvKey(application.KvServiceWebPort)

	// Run server
	router := gin.Default()
	// Load Front End Files
	initRouteStatic(router)
	// Init Template
	initTemplate(router)
	// WEB Route
	initRoute(router)
	_ = router.Run(":" + p)
}

func RegisterWebServer() error {
	// Service Check
	var (
		// Id key
		ckUrl = em.MustGetServiceIdKvKey(define.KvServiceCheckUrl)
		httpPort = em.MustGetServiceIdKvKey(define.KvServiceHttpPort)
		addr = em.MustGetServiceIdKvKey(define.KvServiceAddress)
		webId = em.MustGetServiceIdKvKey(application.KvServiceWebId)
		webPort = em.MustGetServiceIdKvKey(application.KvServiceWebPort)
		// Id or Name key
		webName = em.MustGetServiceKvKey(application.KvServiceWebName)
		interval = em.MustGetServiceKvKey(define.KvServiceCheckInterval)

		webTag = em.IfEmptyChangeJsonFormat(em.MustGetServiceKvKey(application.KvServiceWebTag))
	)

	Check := api.AgentServiceCheck{
		Interval: interval,
		HTTP:     "http://" + addr + ":" + httpPort + ckUrl,
	}

	// Configuration service
	conf := api.AgentServiceRegistration{
		Address:   addr,
		Check: &Check,
	}

	pt, err := strconv.Atoi(webPort)
	if err != nil {
		em_library.InitLog.Println("[ERROR]", "Web Port format error. Error:" + err.Error())
		panic("Web Port format error. Error:" + err.Error())
	}

	var t []string
	err = json.Unmarshal([]byte(webTag), &t)
	if err != nil {
		em_library.InitLog.Println("[ERROR]", "Web Tag format error. Error:" + err.Error())
	}

	webConf := conf
	webConf.ID = webId
	webConf.Name = webName
	webConf.Tags = t
	webConf.Port = pt

	err = em_library.Instance_Consul.Agent().ServiceRegister(&webConf)
	return err
}

// Cancel Service
// 取消服务
func ExitWebServer() error {
	err := em_library.Instance_Consul.Agent().ServiceDeregister(em.MustGetServiceIdKvKey(application.KvServiceWebId))
	return err
}

func initRouteStatic(router *gin.Engine) {
	// router.LoadHTMLGlob("storage/view/dist/*.html")
	router.Static("/static", "./storage/static")
	// router.Static("/storage/upload", "storage/upload")
	return
}

// Init Template
// 初始化模板

func initTemplate(router *gin.Engine) {
	application.HTMLRender = multitemplate.NewRenderer()

	var s model.Setting
	err := s.ReloadTemplate(application.HTMLRender)
	if err != nil {
		panic(err)
	}

	router.HTMLRender = application.HTMLRender
}

