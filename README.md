# req
**像python一样优雅的Go爬虫库**

通过对net/http的再一次封装，使得req的语法更接近python requests的写法

## 一、安装

> go get github.com/serialt/req

## 二、Demo

1.get请求

```go
package main

import (
	"fmt"

	"github.com/serialt/req"
)

func main() {
	resp, err := req.Get("https://www.httpbin.org/get?a=1&b=2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())

	// Params
	params := req.Params{"a": "1", "b": "2"}
	res, err := req.Get("https://www.httpbin.org/get", params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Text())
}
```

2.post请求

```go
package main

import (
	"fmt"

	"github.com/serialt/req"
)

func main() {
	// form post
	data := req.Data{
		"name": "xiaobulv",
	}
	resp, _ := req.Post("https://www.httpbin.org/post", data)
	fmt.Println(resp.Text())

	// json post
	jdata := req.Json{"AGE": "131"}
	respb, _ := req.Post("https://www.httpbin.org/post", jdata)
	fmt.Println(respb.Text())

}
```

3.put delete patch等

```go
package main

import (
	"fmt"

	"github.com/serialt/req"
)

func main() {
	// put
	putdata := req.Json{"AGE": "131"}
	respp, _ := req.Put("https://www.httpbin.org/put", putdata)
	fmt.Println(respp.Text())
}
```

4.auth

```go
package main

import (
	"github.com/serialt/req"
)

func main() {
	respa, _ := req.Get("https://www.httpbin.org/basic-auth/admin/123456", req.Auth{"admin", "123456"})
	println(respa.Text())
}
```

5.header, cookie, timeout, proxy, json等设置

```go
package main

import (
	"fmt"

	"github.com/serialt/req"
)

func main() {
	// 设置header
	header := req.Header{"user-agent": "hello"}

	// post 请求参数
	data := req.Data{
		"name": "xiaobulv",
	}

	// 设置cookie
	// 当Header里有Cookie时, 此设置无效
	ck := req.Cookie{"BIDUPSID": "C855441CA6145FBB2741293580"}

	// timeout
	timeout := req.SetTimeout(10)

	// 代理
	// proxy := req.Proxy("http://xxx.ip.com")

	resp, _ := req.Post("https://www.httpbin.org/post", data, header, ck, timeout)

	// content
	fmt.Println(resp.Content())

	// text
	fmt.Println(resp.Text())

	// response.Json
	m := make(map[string]interface{})
	resp.Json(&m)
	fmt.Println(m["headers"])
	c := m["headers"]
	fmt.Println(c.(map[string]interface{})["Host"])

	// 状态码
	fmt.Println(resp.StatusCode)

	// 响应cookie
	fmt.Println(resp.Cookies())
}
```
6.tls设置
```go
package main

import (
	"fmt"

	"github.com/serialt/req"
)
func main() {

	// tls双向认证
	// certificate, _ := req.NewCertificateFromFiles("xxx.crt", "xxx.key")
	testConfig := req.TLSConfig{
		InsecureSkipVerify: true,
		// 自签名证书的ca
		ClientCAs:          req.NewCAPoolFromFiles("/tmp/ccc/ca.crt"),
		// Certificates:       []tls.Certificate{certificate},
	}

	resp, err := req.Get("https://example.com", testConfig)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Text())
	}
}
```

7.上传文件
```go
package main

import (
	"fmt"

	"github.com/serialt/req"
)

func main() {
	// form post
	files := req.Files{
		"file1": "/tmp/ccc/file1.txt",
	}
	header := req.Header{
		"content-type": "multipart/form-data",
	}
	resp, _ := req.Post("https://www.httpbin.org/post", files)
	fmt.Println(resp.Text())

	// json post
	jdata := req.Json{"AGE": "131"}
	respb, _ := req.Post("https://www.httpbin.org/post", header, jdata)
	fmt.Println(respb.Text())

}

```


## Thanks
* Forked from [sixgad/gorequests](https://github.com/sixgad/gorequests),
