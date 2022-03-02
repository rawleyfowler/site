

[31m2022/03/02 07:29:33 [Recovery] 2022/03/02 - 07:29:33 panic recovered:
GET /admin/post HTTP/1.1
Host: localhost:8080
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-US,en;q=0.5
Cache-Control: max-age=0
Connection: keep-alive
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Sec-Gpc: 1
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0


Cannot redirect with status code 406
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/render/redirect.go:22 (0x7ff947)
	Redirect.Render: panic(fmt.Sprintf("Cannot redirect with status code %d", r.Code))
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:927 (0x806ea6)
	(*Context).Render: if err := r.Render(c.Writer); err != nil {
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:1008 (0x8cab1d)
	(*Context).Redirect: c.Render(-1, render.Redirect{
/home/rf/Projects/personal/site-rework/controllers/admin.go:23 (0x8caabd)
	AdminOnly.func1: c.Redirect(http.StatusNotAcceptable, "/admin/")
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x80f321)
	(*Context).Next: c.handlers[c.index](c)
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/recovery.go:99 (0x80f30c)
	CustomRecoveryWithWriter.func1: c.Next()
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x80f321)
	(*Context).Next: c.handlers[c.index](c)
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/recovery.go:99 (0x80f30c)
	CustomRecoveryWithWriter.func1: c.Next()
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x80e586)
	(*Context).Next: c.handlers[c.index](c)
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/logger.go:241 (0x80e569)
	LoggerWithConfig.func1: c.Next()
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x80dad0)
	(*Context).Next: c.handlers[c.index](c)
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/gin.go:555 (0x80d738)
	(*Engine).handleHTTPRequest: c.Next()
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/gin.go:511 (0x80d271)
	(*Engine).ServeHTTP: engine.handleHTTPRequest(c)
/usr/lib/go/src/net/http/server.go:2879 (0x6b66da)
	serverHandler.ServeHTTP: handler.ServeHTTP(rw, req)
/usr/lib/go/src/net/http/server.go:1930 (0x6b1d87)
	(*conn).serve: serverHandler{c.server}.ServeHTTP(w, w.req)
/usr/lib/go/src/runtime/asm_amd64.s:1581 (0x464d80)
	goexit: BYTE	$0x90	// NOP
[0m


[31m2022/03/02 07:29:33 [Recovery] 2022/03/02 - 07:29:33 panic recovered:
GET /admin/post HTTP/1.1
Host: localhost:8080
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-US,en;q=0.5
Cache-Control: max-age=0
Connection: keep-alive
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Sec-Gpc: 1
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0


Cannot redirect with status code 406
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/render/redirect.go:22 (0x7ff947)
	Redirect.Render: panic(fmt.Sprintf("Cannot redirect with status code %d", r.Code))
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:927 (0x806ea6)
	(*Context).Render: if err := r.Render(c.Writer); err != nil {
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:1008 (0x8cab1d)
	(*Context).Redirect: c.Render(-1, render.Redirect{
/home/rf/Projects/personal/site-rework/controllers/admin.go:23 (0x8caabd)
	AdminOnly.func1: c.Redirect(http.StatusNotAcceptable, "/admin/")
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x80f321)
	(*Context).Next: c.handlers[c.index](c)
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/recovery.go:99 (0x80f30c)
	CustomRecoveryWithWriter.func1: c.Next()
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x80f321)
	(*Context).Next: c.handlers[c.index](c)
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/recovery.go:99 (0x80f30c)
	CustomRecoveryWithWriter.func1: c.Next()
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x80e586)
	(*Context).Next: c.handlers[c.index](c)
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/logger.go:241 (0x80e569)
	LoggerWithConfig.func1: c.Next()
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x80dad0)
	(*Context).Next: c.handlers[c.index](c)
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/gin.go:555 (0x80d738)
	(*Engine).handleHTTPRequest: c.Next()
/home/rf/.go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/gin.go:511 (0x80d271)
	(*Engine).ServeHTTP: engine.handleHTTPRequest(c)
/usr/lib/go/src/net/http/server.go:2879 (0x6b66da)
	serverHandler.ServeHTTP: handler.ServeHTTP(rw, req)
/usr/lib/go/src/net/http/server.go:1930 (0x6b1d87)
	(*conn).serve: serverHandler{c.server}.ServeHTTP(w, w.req)
/usr/lib/go/src/runtime/asm_amd64.s:1581 (0x464d80)
	goexit: BYTE	$0x90	// NOP
[0m
