# notes

## chapter 02 迷你论坛

## chapter 03 接受请求

net/http package

- Client 、Response 、Header 、Request 和 Cookie 对客户端进⾏⽀持
- Server 、ServeMux 、Handler/HandleFunc 、 ResponseWriter 、Header 、Request 和 Cookie 则对服务器进⾏⽀持

## chapter 04 处理请求

- Form 字段
- 如果⼀个键同时拥有表单键值对和 URL 键值对，但是⽤户只想要获取表单键值对⽽不是 URL 键值对，那么可以访问 Request 结构的
  PostForm 字段。PostFormValue ⽅法只会返回表单键值对⽽不会返回 URL 键值对。
- MultipartForm 字段：取得 multipart/form-data 编码的表单数据，一般用于包含 file 的表单

一个疑问： ServeHTTP 为什么要接 受 ResponseWriter 接⼝和⼀个指向 Request 结构的指针作为参数呢？

- 接受 Request 结构指针的原因很简单：为了让服务器能够察觉到处理器对 Request 结构的修改，必须
以传引⽤（pass by reference）⽽不是传值（pass by value）的⽅式传递 Request 结构。
- 但是为什么 ServeHTTP 却是以传值的⽅式接受 ResponseWriter 呢？难道服务器不 需要知道处理器对 ResponseWriter 所做的修改吗？
对于这个问题，如果我们深⼊探究 net/http 库的源码，就会发现 ResponseWriter 实际上就是 response 这个⾮导出结构的接⼝，
⽽ResponseWriter 在使⽤response 结构时，传递的也是指向 response 结构的指针，这也就是说，ResponseWriter 是以传引
⽤⽽不是传值的⽅式在使⽤response 结构。 换句话说，实际上 ServeHTTP 函数的两个参数传递的都是引⽤⽽不是值——虽然
ResponseWriter 看上去像是⼀个值，但它实际上却是⼀个带有结构指针的接⼝。

ResponseWriter ：
- write => 写响应内容
- writeHeader => 写 status code。注意：WriteHeader ⽅法在执⾏完毕之后就不允许再对⾸部进⾏写⼊。
- Header => 写 headers

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

- 没有设置 Expires 字段的 cookie 通常称为会话 cookie 或者临时
  cookie，这种 cookie 在浏览器关闭的时候就会⾃动被移除。相对⽽⾔，
  设置了 Expires 字段的 cookie 通常称为持久 cookie，这种 cookie 会⼀直
  存在，直到指定的过期时间来临或者被⼿动删除为⽌。
- Expires 字段⽤于明确地指定 cookie 应该在什么时候过期，
  ⽽MaxAge 字段则指明了 cookie 在被浏览器创建出来之后能够存活多少
  秒。一个是绝对时间，一个是相对时间。

---

flash message using cookie

实现闪现消息的⽅法有很多种，但最常⽤的⽅法是把这些消息存
储在⻚⾯刷新时就会被移除的会话 cookie⾥⾯。参考 cookie-flash/

## chapter 05 内容展示

Web 模板引擎演变⾃SSI（服务器端包含）技术，并最
终衍⽣出了诸如 PHP、ColdFusion 和 JSP 这样的 Web 编程语⾔。后端渲染。

## chapter 06 数据持久化

- 数据保存到内存：一般是一个 map。
- 数据保存到文件：csv、gob（二进制）、raw
- 数据保存到数据库：利用 `database/sql` 内置库，sqlx 或者 gorm 框架，保存到传统数据库。

## chapter 07 Go Web 服务

xml 和 json 的读取和保存，以及一个简易的 Web API Server。

## chapter 08 应用测试

- Go 语⾔的 testing 库
  - 基本测试
  - bench 测试：-bench
  - 常用 options：-cover -v
  - 并行测试：t.Parallel()
  - 跳过测试 cmd option `-short` 配合代码中的 Short()/Skip()
- 单元测试
- HTTP 测试
  - httptest 库，一开始使用真实数据库链接来测试。
- 使⽤依赖注⼊进⾏测试
  - ! 使用依赖注入来 mock db 到 post 的一个字段，测试时构造一个 FakePost 实例传入，就能很轻便地进行测试。这个办法很好。
- 使⽤第三⽅测试库
  - Gocheck 较为简单，整合并扩展了 testing 包；github库最后更新2020年。
    - `go get gopkg.in/check.v1`
    - `go test -check.v` not work
  - Ginkgo 实现⾏为驱动开发，⽐较复杂。
    - 转化：ginkgo convert .  => 被启用，参考这个[链接](https://onsi.github.io/ginkgo/MIGRATING_TO_V2#removed-ginkgo-convert)
    - 执行测试： `ginkgo -v` => not work, `go test -v` work but no verbose log
    - 创建全新的测试：ginkgo bootstrap && ginkgo generate

## chapter 09 并发

- goroutine 与 waitGroup 精准等待
- chan
  - w1 <- true 写channel 与 <-w1 阻塞读channel
  - 带缓冲的channel实现消费者/生产者模型：消息传递
  - select 对多个chan进行选择。
- mosaic 示例
  - 这种预置tiles db的方式很诡异。
  - 本地运行，得到了错误的结果。部分区域全白，部分区域全黑。可以肯定是寻找近似tile时发生了边界情况。

## chapter 10 部署

App 的部署。
- IaaS: EC2 GCE Droplets(Digital Ocean)
- Paas: Heroku、AWS的Elastic Beanstalk以及Google公司的App Engine
- Saas: AWS 的RDS服务、Google公司的Cloud SQL（云SQL）

---
- IaaS：基于云的服务，按需付费，用于存储，网络和虚拟化等服务。
- PaaS： Internet上可用的硬件和软件工具。
- SaaS： 可通过互联网通过第三方获得的软件

---

- 进程守护：使用 upstart或者systemd 。原生。
- heroku：使⽤Godep⽣成本地依赖关系并创建Procfile⽂件。godeps和heroku这种方式已经过时。与平台绑定。
- GAE：部署⽅法复杂，优点在于被部署的Web服务将获得⾮常好的可扩展性。与平台密切绑定。

## chapter 11 使用 Web 框架改写
一些web框架。
- Beego：github.com/astaxie/beego
- martini：github.com/go-martini/martini
- goji：github.com/zenazn/goji

都是些很古老的框架，目前来看，gin或许更为主流。

## summary

下面是一些值得回顾和注意的章节：

- 02 的迷你论坛模型很不错，值得扩展和继续完善。
- 06 数据持久化，使用的是原生自带库
- 08 应用测试
- 10 部署方式介绍