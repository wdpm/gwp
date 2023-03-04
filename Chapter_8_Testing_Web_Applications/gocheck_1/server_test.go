package main

import (
	"encoding/json"
	"fmt"

	// 以点（. ）⽅式导⼊check 包的，包中所有被导出的标识符在测试程序⾥⾯都可以以不带前缀的⽅式访问
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PostTestSuite struct {
}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuite) TestHandleGet(c *C) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	c.Check(writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)

	fmt.Println(post)
	c.Check(post.Id, Equals, 1)
}
