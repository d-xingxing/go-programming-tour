### 标签管理

* 新增标签

```shell
$ curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=Golang' -F created_by=xizhao
{}
$ curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=PHP' -F created_by=xizhao
{}
$ curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=Rust' -F created_by=xizhao
{}
```

* 获取标签列表

```shell
$ curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'
{"list":[{"id":1,"created_by":"xizhao","modified_by":"","created_on":1574493416,"modified_on":1574493416,"deleted_on":0,"is_del":0,"name":"Go 语言","state":1},{"id":2,"created_by":"xizhao","modified_by":"","created_on":1574493813,"modified_on":1574493813,"deleted_on":0,"is_del":0,"name":"PHP","state":1}],"pager":{"page":1,"page_size":2,"total_rows":3}}

$ curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=2&page_size=2'
{"list":[{"id":3,"created_by":"xizhao","modified_by":"","created_on":1574493817,"modified_on":1574493817,"deleted_on":0,"is_del":0,"name":"Rust","state":1}],"pager":{"page":2,"page_size":2,"total_rows":3}}

```

* 修改标签

```shell
$  curl -X PUT http://127.0.0.1:8000/api/v1/tags/1 -F state=0 -F modified_by=zhang123
{}
```

* 删除标签

```shell
$ curl -X DELETE  http://127.0.0.1:8000/api/v1/tags/{id}
{}
```

* 文件上传

```shell
curl -X POST http://127.0.0.1:8000/upload/file -F file=@/Users/xizhao/Desktop/course01.jpg -F type=1
{"file_access_url":"http://127.0.0.1:8000/static/a86dbd27118f444fe379e0692cb5dfe5.jpg"}%
}
```

* 验证获取token的接口

```shell
curl -X POST http://127.0.0.1:8000/auth -F app_key=xizhao -F app_secret=go-programming-tour
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiNGE4ODhlOWZhZmZlMDJiOTcxN2ZiZGMyNWI0NzU4NGMiLCJhcHBfc2VjcmV0IjoiMWZhNTFhMjk1MjIyYTNmMzMxZjQwMzFkM2MzY2M2MjYiLCJleHAiOjE2NTIxMTM4OTQsImlzcyI6ImJsb2ctc2VydmljZSJ9.lP9zsu4aaLyfJBDwbZsaufu92mdT7IVjNlTSAyd5f3I"}%
```

```shell
curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2' -H 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiNGE4ODhlOWZhZmZlMDJiOTcxN2ZiZGMyNWI0NzU4NGMiLCJhcHBfc2VjcmV0IjoiMWZhNTFhMjk1MjIyYTNmMzMxZjQwMzFkM2MzY2M2MjYiLCJleHAiOjE2NTIxMTM4OTQsImlzcyI6ImJsb2ctc2VydmljZSJ9.lP9zsu4aaLyfJBDwbZsaufu92mdT7IVjNlTSAyd5f3I'
## 需要带token鉴权
```


