# Template / 模板



### Directory structure and file content / 目录结构及文件内容

view/template/[YOUR_TEMPLATE_NAME]/layout

#### /Layout/

##### app.html

> ```
> {{ block "app" .}}{{end}}
> ```



##### layout.html

> ```
> {{ define "layout" }}
> ...
> <body>
> <div id="vue">
> 		[YOUR_HTML]
> 	</div>
> 	{{ template "javascript" . }}
> </body>
> ...
> {{end}}
> ```
>
> 
>
> Variables that may be used / 可能用到的变量
>
> 1. Template reservation / 模板预留位
>
>    `{{ template "head" . }}`
>
>    `{{ template "header" . }}`
>
>    `{{ block "banner" . }}{{ end }}`
>
>    `{{ block "content" . }}{{ end }}`
>
>    `{{ template "nav" . }}`
>
>    `{{ template "footer" . }}`
>
>    `{{ template "javascript" . }}`
>
> 
>
> Simple Example / 简单的例子:
>
> ```html
> {{ define "layout" }}
>     <!DOCTYPE html>
>     <html>
>     <head>
>         {{ template "head" . }}
>     </head>
>     <body>
>     <div id="vue">
>     <header>
>         {{ template "header" . }}
>     </header>
> 
> 
> 
>     {{ block "banner" . }}{{ end }}
>     {{ block "content" . }}{{ end }}
> 
>     {{ template "nav" . }}
> 
>     <footer id="footer">
> 
>         {{ template "footer" . }}
> 
>     </footer>
> 
>     <div class="scroll" id="scroll" style="display:none;">Top</div>
> 
>     </div>
> 
>     {{ template "javascript" . }}
> 
>     </body>
>     </html>
> {{end}}
> ```





##### head.html

> ```
> {{ define "head" }}
> ...
> [YOUR_HTML]
> ...
> {{end}}
> ```
>
> 
>
> Variables that may be used / 可能用到的变量
>
> 1. Title / 标题: `{{ block "title" . }}{{.site_name}}{{ end }}`
>
> 2. CSS: `{{ block "css" .}}{{end}}`
>
> 
>
> Simple Example / 简单的例子:
>
> ```html
> {{ define "head" }}
>     <meta charset="utf-8">
> 	...
>     <title>{{ block "title" . }}{{.site_name}}{{ end }}</title>
>     <meta name="description" content="" />
> 	...
>     <link href="/static/w4zfy8vd/css/bootstrap.css" rel="stylesheet">
> 	...
>     {{ block "css" .}}{{end}}
> 	...
>     <script src="/static/js/vue.js"></script>
>     <!-- use the latest release -->
>     <script src="/static/js/vuejs-paginate.js"></script>
> 	...
> {{end}}
> ```



#####  header.html

> ```
> {{ define "header" }}
> ...
> [YOUR_HTML]
> ...
> {{ end }}
> ```
>
> 
>
> Variables that may be used / 可能用到的变量
>
> 1. Background variables / 后台变量
>
>    Phone / 电话: {{.phone}}
>
> 2. Style / 样式
>
>    Is it the homepage / 是否为首页: `data.active_index`
>
>    > Style binding / 样式绑定: `:class="{ 'current': data.active_index }`
>
>    Check if it is the current page / 检查是否为当前页: `checkActive()`
>
>    >  `:class="{ 'current': checkActive(item.id) }"`
>
> 3. Category tree / 分类树: `data.category_tree`
>
>    >  ```
>    >  v-for="item in data.category_tree" :key="item.id"
>    >  ```
>
>    Whether to include sub-categories / 是否包含子分类
>
>    >  ```
>    >  v-if="item.children"
>    >  ```
>
> 
>
> Simple Example / 简单的例子:
>
> ```html
> {{ define "header" }}
> 	...
> 	<i class="fa fa-phone icon-secondary" aria-hidden="true"></i>&nbsp;&nbsp;{{.phone}}<span class="xcang"><i class="fa fa-envelope-o icon-secondary" aria-hidden="true"></i>&nbsp;&nbsp;<span><a href="mailto:{{.email}}">{{.email}}</a></span></span> </div>
> 	...
> 
>                 <ul class="nav navbar-nav">
> 
>                     <li :class="{ 'active': data.active_index }">
>                         <a href="/">首页</a>
>                     </li>
> 
>                     <li v-for="item in data.category_tree" :key="item.id" class="dropdown" :class="{ 'active': checkActive(item.id) }">
>                         <a :href="'/category/' + item.url_path">[[item.name]]</a>
>                         <a :href="'/category/' + item.url_path" id="app_menudown" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
>                             <span class="glyphicon glyphicon-menu-down btn-xs"></span>
>                         </a>
>                         <ul v-if="item.children" class="dropdown-menu nav_small" role="menu">
>                             <li v-for="sub in item.children" :key="sub.id"><a :href="'/category/' + sub.url_path">[[sub.name]]</a></li>
>                         </ul>
>                     </li>
>                 </ul>
> 
> {{end}}
> ```



#####  footer.html

> ```
> {{ define "footer" }}
> ...
> [YOUR_HTML]
> ...
> {{ end }}
> ```
>
> 
>
> Variables that may be used / 可能用到的变量
>
> 1. Background variables / 后台变量
>
>    Phone / 电话: {{.phone}}
>
>    Address / 地址: {{.address}}
>
>    Email / 邮箱: {{.email}}
>
>    ICP record number / ICP备案号 : {{.icp}}
>
> 2. Get the sub-category of the specified category based on the ID / 根据ID获取指定分类的子分类: `PageGetCategoryChildren()`
>
> 
>
> Simple Example / 简单的例子:
>
> ```html
> {{ define "footer" }}
> 
> 
>                         <h4 class="f-h4"><span class="f-h4-span">热门分类</span></h4>
>                         <div class="footer-block">
>                                 <p v-for="item in data_PageGetCategoryChildren" :key="item.id"><a :href="'/category/' + item.url_path">[[item.name]]</a>&nbsp;&nbsp;</p>
>                         </div>
> 
>                     <p><img src="/static/w4zfy8vd/images/icon-phone.png" height="30" width="30"> <span class="footer-add">Tel：{{.phone}}</span> </p>
>                     <p><img src="/static/w4zfy8vd/images/icon-whatsapp.png" height="30" width="30"> <span class="footer-add">手机：{{.phone}}</span> </p>
>                     <p><img src="/static/w4zfy8vd/images/icon-email.png" height="30" width="30"> <span class="footer-add">E-mail: {{.email}}</span> </p>
>                     <p><img src="/static/w4zfy8vd/images/icon-map.png" height="30" width="30"> <span class="footer-add">Add：{{.address}}</span> </p>
> 
>     <div class="footer-bar">
>         <p>{{.site_name}} <a href="http://beian.miit.gov.cn/"> 沪ICP备18037708号-1</a></p>
>     </div>
> {{ end }}
> ```





##### javascript.html

> ```
> {{ define "javascript" }}
>  {{ block "js_begin" .}}{{end}}
>  [TEMPLATE JS]
>  {{ block "js" .}}{{end}}
>  [PLUGIN JS]
>  {{ block "js_end" .}}{{end}}
> {{end}}
> ```
>
> 



##### paginate.html

> Document / 文档: https://github.com/lokyoung/vuejs-paginate
>
> Simple Example:
>
> ```
> {{ define "paginate" }}
>  <div class="pages">
>      <paginate
>              v-model="paginate.current_page"
>              :page-count="paginate.page_count"
>              :container-class="''"
>              :active-class="'thisclass'"
>              :prev-text="'<'"
>              :next-text="'>'"
>              :first-last-button="true"
>              :first-button-text="'首页'"
>              :last-button-text="'末页'"
>              :click-handler="clickCallback"
>      >
>      </paginate>
>  </div>
> {{ end }}
> ```
>
> 



##### sidebar.html

> Variables that may be used / 可能用到的变量:
>
>  1. Data / 数据: 
>
>     (1)Get the sub-columns of the current column / 获取当前栏目的子级栏目: `data.sidebar.children`
>
> Simple Example:
>
>  ```
> {{ define "sidebar" }}
> 
>                     <li v-for="item in data.sidebar.children" :key="item.id">
> 
>                         <a :href="'/category/'+item.url_path">[[item.name]]</a>
>                         <ul class="prolist_side_son close" v-for="v in item.children" :key="v.id" class='children' >
> 
>                             <li style="padding-left:10px;">
>                                 <a :href="'/category/'+v.url_path">[[v.name]]</a>
>                             </li>
> 
>                         </ul>
>                     </li>
> 
> {{ end }}
>  ```
>
> 



view/template/[YOUR_TEMPLATE_NAME]/

#### /

##### 404.html

> ```javascript
> {{define "app"}}
>  {{ block "content" .}}{{end}}
> {{end}}
> 
> {{define "content"}}
>  [YOUR_HTML]
> {{end}}
> ```
>
> 
>
> Simple Example:
>
> ```javascript
> {{define "app"}}
>  {{ block "content" .}}{{end}}
> {{end}}
> 
> {{define "content"}}
>  404
> {{end}}
> ```
>
> 





##### index.html

> 
>
> Variables that may be used / 可能用到的变量
>
> 1. Title / 标题: `{{define "title"}}{{.site_name}}{{end}}`
>
> 2. Thumbnail / 缩略图: getFirstThumbnail()
>
>    > ```
>    > `:src="getFirstThumbnail(item.thumbnail)"`
>    > ```



##### category_image.html / category_text.html

> ```
> {{define "app"}}
> {{ template "layout" .}}
> {{end}}
> 
> {{define "banner"}}
> ...
> {{end}}
> 
> {{define "content"}}
> ...
> {{end}}
> 
> {{ define "js" }}
> ...
> {{end}}
> ```
>
> Variables that may be used / 可能用到的变量
>
> 1. Template / 模板: 
>
>    (1) Sidebar / 侧边栏: `{{template "sidebar" . }}`
>
>    (2) Paginate / 分页: `{{template "paginate" . }}`
>
> 2. Data / 数据:
>
>    (1) Post / 文章: `data.post`
>
>    (2) Category / 顶级分类: `data.top_category`
>
>    (3) Breadcrumbs / 面包屑: `data.breadcrumbs`
>
>    ```vue
>    <li v-for="item in data.breadcrumbs" :key="item.id"><a :href="'/category/'+item.url_path">[[item.name]]</a></li>
>    ```
>
>    
>
> Simple Example:
>
> ```javascript
> <li v-for="item in data.post" :key="item.id" class="col-lg-4 col-md-4 col-sm-4 col-xs-6">
>     <div class="pic">
>         <img :src="getFirstThumbnail(item.thumbnail)" :alt="item.name">
>             <div class="pic-bg"><a :href="'/post/'+item.url_path">
>                 <i class="iconfont"> </i>
> <h3>[[item.name]]</h3>
> <b></b></a> </div>
> </div>
> <div class="title"> <a :href="'/post/'+item.url_path">[[item.name]] </a></div>
>     </li>
> ```
>
> 



##### category_page.html

> ```
> {{define "app"}}
> {{ template "layout" .}}
> {{end}}
> 
> {{define "banner"}}
> ...
> {{end}}
> 
> {{define "content"}}
> ...
> {{end}}
> 
> {{ define "js" }}
> ...
> {{end}}
> ```
>
> Variables that may be used / 可能用到的变量
>
> 1. Data / 数据:
>
>    (1) Top category / 顶级分类: `data.top_category`
>
>    (2) Current category / 当前分类: `data.category`
>
>    (3) Current Post / 当前文章: `data.post`



##### post_image.html / post_text.html

> ```
> {{define "app"}}
> {{ template "layout" .}}
> {{end}}
> 
> {{define "banner"}}
> ...
> {{end}}
> 
> {{define "content"}}
> ...
> {{end}}
> 
> {{ define "js" }}
> ...
> {{end}}
> ```
>
> Variables that may be used / 可能用到的变量
>
> 1. Data / 数据:
>
>    (1) Post / 文章: `data.post`
>
>    - Post title/ 文章标题: `data.post.name`
>
>    - Post content / 文章内容: data.post.content
>
>    > ```
>    > v-if="data.post" v-html="data.post.content"
>    > ```
>
>    (2) Category / 分类
>
>    - Top category / 顶级分类: `data.top_category`
>      - Top category name / 顶级分类名: `data.top_category.name`
>
>    (3) active_category_ids: ID of the current column location / 当前栏目所在位置的ID
>
>    (4) breadcrumbs: breadcrumbs / 面包屑
>
>    (5) category_tree: category tree / 导航树
>
>    (6) related_post: Related posts / 相关文章
>
>    (7) sidebar: sidebar,Contains a subset of the current top column / 侧边栏,包含当前顶级栏目的子集
>
>    (8)top_category: Top-level section information / 顶级栏目信息



##### search.html

> ```javascript
> {{define "app"}}
> {{ template "layout" .}}
> {{end}}
> 
> {{define "banner"}}
> ...
> {{end}}
> 
> {{define "content"}}
> ...
> {{end}}
> 
> {{ define "js" }}
> ...
> {{end}}
> ```
>
> 
>
> Simple Example:
>
> ```javascript
> <form action="/search" method="get" id="search-block-form" accept-charset="UTF-8">
> 	<div class="form-type-textfield form-item-search-block-form form-item form-group">
> 		<label for="edit-search-block-form--2">Search </label>
> 		<input title="Enter the terms you wish to search for." placeholder="Search here" type="text" id="edit-search-block-form--2" name="keyword" value="" size="15" maxlength="128" />
> 	</div>
> 	<button class="search-submit"></button>
> 	<div class="form-actions form-wrapper" id="edit-actions--11">
> 		<input class="btn form-submit" type="submit" id="edit-submit" name="op" value="Search" />
> 	</div>
> </form>
> ```
>
> 





### Template function / 模板函数

#### General function / 通用函数



##### getFirstThumbnail()

Get thumbnail / 获取缩略图

```javascript
getFirstThumbnail: function(thumbnailArray) {
    if (thumbnailArray != null) {
        if (thumbnailArray.length > 0) {
            if ("path" in thumbnailArray[0]) {
                return thumbnailArray[0].path
            }
        }
    }
    return "/static/images/nopic.jpg"
},
```



##### checkActive()

Check whether the current column ID belongs to the ID of the current URL or the parent ID of the current URL / 检查当前栏目ID是否属于当前URL所属的ID或当前URL所属父级ID下

```javascript
checkActive: function(id) {
    if (this.data.active_category_ids) {
        if (this.data.active_category_ids.indexOf(id) > -1){
            this.category_show_id = id
            return true
        }
    }
    return false
},
```



##### PageGetCategoryChildren()

Get the subset according to the parent category_ID / 根据父级category_ID得到子集

```javascript
PageGetCategoryChildren: function(id) {
    var self = this;
    // Make a request for a user with a given ID
    jQuery.ajax({
        type: 'GET',
        url: '/data/get/category/children/' + id,
        async: false,
        dataType:'json',
        success: function(response){
            // handle success
            self.data_PageGetCategoryChildren = response.data
        },
        error: function(response){
            // handle error
            console.log(response);
        }
    });
},
```



##### PageGetPostByCategoryId()

Get Post by Categroy ID (Include sub-category articles) / 根据分类ID得到对应文章（包括子分类文章）

1. data_PageGetPostByCategoryId => data
2. 2. this.PageGetPostByCategoryId(Your ID);  => created

```javascript
        // Code:
        PageGetPostByCategoryId: function(id) {
        var self = this;
        // Make a request for a user with a given ID
        let data = {"id":id};
        jQuery.ajax({
            type: 'POST',
            url: '/data/get/post/byCategoryId',
            data: JSON.stringify(data),
            contentType: "application/json",
            async: false,
            dataType:'json',
            success: function(response){
                // handle success
                self.data_PageGetPostByCategoryId = response.data
            },
            error: function(response){
                // handle error
                console.log(response);
            }
        });
    },
```







#### index.html

```javascript
<script>
    var app = new Vue({
        delimiters: ['[[', ']]'],
        el: '#vue',
        data: function() {
            return {
                data: [],
            }
        },
        created: function() {},
        methods: {},
    })
</script>
```



##### PageGetIndex()

```javascript
PageGetIndex: function() {
    var self = this;
    // Make a request for a user with a given ID
    jQuery.ajax({
        type: 'GET',
        url: '/data/get/index',
        async: false,
        dataType:'json',
        success: function(response){
            // handle success
            self.data = response.data;
            self.data.active_index = self.data.active_category_ids === undefined;
        },
        error: function(response){
            // handle error
            console.log(response);
        }
    });
},
```



#### category_page.html

```javascript
<script>
    var app = new Vue({
        delimiters: ['[[', ']]'],
        el: '#vue',
        data: function() {
            return {
                data: [],
            }
        },
        created: function() {},
        methods: {}
    })
</script>
```



##### PageGetCategroyPage()

```javascript
PageGetCategroyPage: function() {
    var self = this;
    let data = {"url_path":document.location.pathname};
    jQuery.ajax({
        type: 'POST',
        url: '/data/get/category/page',
        data: JSON.stringify(data),
        contentType: "application/json",
        async: false,
        dataType:'json',
        success: function(response){
            // handle success
            self.data = response.data;
            // Set title
            var str = "";
            for (var i= self.data.breadcrumbs.length-1; i>=0; i--){
                str += self.data.breadcrumbs[i].name + " - "
            }
            str += {{.site_name}}
        document.title = str
    },
                error: function(response){
        // handle error
        console.log(response);
    }
})
},
```



#### category_image.html / category_text.html

```javascript
    <script>
        Vue.component('paginate', VuejsPaginate)

        var app = new Vue({
            delimiters: ['[[', ']]'],
            el: '#vue',
            data: function() {
                return {
                    data: [],
                    paginate: {
                        current_page: {{.current_page}},
                        per_page: 12,
                        page_count: 0,
                        post_count: 0,
                    }
                }
            },
            created: function() {},
            methods: {}
        })
    </script>
```





##### PageGetCategroyList()

Get Categroy List / 获取列表分类栏目数据

```javascript
PageGetCategroyList: function() {
    var self = this;
    // Make a request for a user with a given ID
    let data = {"url_path":document.location.pathname, "current_page": this.paginate.current_page, "per_page": this.paginate.per_page};
    jQuery.ajax({
        type: 'POST',
        url: '/data/get/category/list',
        data: JSON.stringify(data),
        contentType: "application/json",
        async: false,
        dataType:'json',
        success: function(response){
            // handle success
            self.data = response.data;
            self.post_count = self.data.post_count;
            self.paginate.page_count = Math.ceil( self.post_count / self.paginate.per_page);

            var str = "";
            for (var i= self.data.breadcrumbs.length-1; i>=0; i--){
                str += self.data.breadcrumbs[i].name + " - "
            }
            str += {{.site_name}}
        document.title = str
    },
                error: function(response){
        // handle error
        console.log(response);
    }
});
},
```



##### clickCallback()

Jump paging function 跳转分页函数

```javascript
clickCallback: function(pageNum) {
    page_url = location.protocol + '//' + location.host + location.pathname + '?page=' + pageNum;
    window.location.href=page_url
},
```



#### post_image.html / post_text.html

```javascript
<script>
    var app = new Vue({
        delimiters: ['[[', ']]'],
        el: '#vue',
        data: function() {
            return {
                data: [],
            }
        },
        created: function() {},
        methods: {},
    })
</script>
```



##### PageGetPost()

```javascript
PageGetPost: function() {
    var self = this;
    // Make a request for a user with a given ID
    let data = {"url_path":document.location.pathname};
    jQuery.ajax({
        type: 'POST',
        url: '/data/get/post',
        data: JSON.stringify(data),
        contentType: "application/json",
        dataType:'json',
        async: false, // 同步
        success: function(response){
            // handle success
            self.data = response.data;
            if (self.data.post.parameter && self.data.post.parameter !== ""){
                self.data.post.parameter = JSON.parse(self.data.post.parameter);
            }
            // Set Title
            var str = self.data.post.name + " - ";
            for (var i= self.data.breadcrumbs.length-1; i>=0; i--){
                str += self.data.breadcrumbs[i].name + " - "
            }
            str += {{.site_name}}
        document.title = str
    },
                error: function(response){
        // handle error
        console.log(response);
    }
});
},
```



#### search.html

```javascript
    <script>
        Vue.component('paginate', VuejsPaginate)

        var app = new Vue({
            delimiters: ['[[', ']]'],
            el: '#vue',
            data: function() {
                return {
                    data: [],
                    paginate: {
                        current_page: {{.current_page}},
                        per_page: 12,
                        page_count: 0,
                        post_count: 0,
                    }
                }
            },
            created: function() {},
            methods: {}
        })
    </script>
```

##### PageGetSearch()

```javascript
PageGetSearch: function() {
    var self = this;
    // Make a request for a user with a given ID
    let data = {"url_path":document.location.pathname, "current_page": this.paginate.current_page, "per_page": this.paginate.per_page, "keyword": {{.keyword}}, "sidebar_category_id": this.sidebar_category_id};
jQuery.ajax({
    type: 'POST',
    url: '/data/get/search',
    data: JSON.stringify(data),
    contentType: "application/json",
    dataType:'json',
    async: false, // 同步
    success: function(response){
        // handle success
        self.data = response.data;
        self.post_count = self.data.post_count;
        self.paginate.page_count = Math.ceil( self.post_count / self.paginate.per_page);
        // Set Title
        var str = "{{.keyword}} - Search - {{.site_name}}"
        document.title = str
    },
    error: function(response){
        // handle error
        console.log(response);
    }
});
},
```





### Common Vue syntax / 常用Vue语法

- Filter custom labels / 过滤自定义标签: `v-pre`

