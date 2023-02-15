package req

// 常见几种请求类型
func Head(origurl string, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// call request Head
	resp, err = req.Send("HEAD", origurl, args...)
	return resp, err
}
func Get(origurl string, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// call request Get
	resp, err = req.Send("GET", origurl, args...)
	return resp, err
}

func Post(origurl string, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// call request Post
	resp, err = req.Send("POST", origurl, args...)
	return resp, err
}

func Put(origurl string, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// call request Put
	resp, err = req.Send("PUT", origurl, args...)
	return resp, err
}

func Patch(origurl string, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// call request Patch
	resp, err = req.Send("PATCH", origurl, args...)
	return resp, err
}

func Delete(origurl string, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// call request Delete
	resp, err = req.Send("DELETE", origurl, args...)
	return resp, err
}
