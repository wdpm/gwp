# notes

## chapter 03

net/http package

- Client 、Response 、Header 、Request 和 Cookie 对客户端进⾏⽀持
- Server 、ServeMux 、Handler/HandleFunc 、 ResponseWriter 、Header 、Request 和 Cookie 则对服务器进⾏⽀持

## chapter 04

- Form 字段
- 如果⼀个键同时拥有表单键值对和 URL 键值对，但是⽤户只想要获取表单键值对⽽不是 URL 键值对，那么可以访问 Request 结构的
  PostForm 字段。PostFormValue ⽅法只会返回表单键值对⽽不会返回 URL 键值对。
- MultipartForm 字段：取得 multipart/form-data 编码的表单数据，一般用于包含 file 的表单

一个疑问： ServeHTTP 为什么要接 受ResponseWriter 接⼝和⼀个指向Request 结构的指针作为参数呢？

- 接受Request 结构指针的原因很简单：为了让服务器能够察觉到处理器对Request 结构的修改，必须
以传引⽤（pass by reference）⽽不是传值（pass by value）的⽅式传递Request 结构。
- 但是为什么ServeHTTP 却是以传值的⽅式接受ResponseWriter 呢？难道服务器不 需要知道处理器对ResponseWriter 所做的修改吗？
对于这个问题，如果我们深⼊探究net/http 库的源码，就会发现ResponseWriter 实际上就是response 这个⾮导出结构的接⼝，
⽽ResponseWriter 在使⽤response 结构时，传递的也是指向response 结构的指针，这也就是说，ResponseWriter 是以传引
⽤⽽不是传值的⽅式在使⽤response 结构。 换句话说，实际上ServeHTTP 函数的两个参数传递的都是引⽤⽽不是值——虽然
ResponseWriter 看上去像是⼀个值，但它实际上却是⼀个带有结构指针的接⼝。

ResponseWriter ：
- write => 写响应内容
- writeHeader => 写 status code。注意：WriteHeader ⽅法在执⾏完毕之后就不允许再对⾸部进⾏写⼊。
- Header => 写headers

### cookie
```go
type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
```

- 没有设置Expires 字段的cookie通常称为会话cookie或者临时
  cookie，这种cookie在浏览器关闭的时候就会⾃动被移除。相对⽽⾔，
  设置了Expires 字段的cookie通常称为持久cookie，这种cookie会⼀直
  存在，直到指定的过期时间来临或者被⼿动删除为⽌。
- Expires 字段⽤于明确地指定cookie应该在什么时候过期，
  ⽽MaxAge 字段则指明了cookie在被浏览器创建出来之后能够存活多少
  秒。一个是绝对时间，一个是相对时间。

---

flash message using cookie

实现闪现消息的⽅法有很多种，但最常⽤的⽅法是把这些消息存
储在⻚⾯刷新时就会被移除的会话cookie⾥⾯。参考 cookie-flash/

## chapter 05