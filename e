[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] Loaded HTML Templates (20): 
	- forbidden.tmpl
	- not_found.tmpl
	- create_post.tmpl
	- footer.tmpl
	- blog_post.tmpl
	- comment_post.tmpl
	- contact.tmpl
	- index.tmpl
	- login.tmpl
	- resume.tmpl
	- 
	- blog.tmpl
	- wip.tmpl
	- comment_post_spam.tmpl
	- footer_no_banners.tmpl
	- header.tmpl
	- internal_server_error.tmpl
	- nav.tmpl
	- admin_success_redirect.tmpl
	- comment_post_failed.tmpl

[GIN-debug] GET    /blog/                    --> gitlab.com/rawleyifowler/site-rework/controllers.RenderBlogPage (4 handlers)
[GIN-debug] GET    /blog/post/:url           --> gitlab.com/rawleyifowler/site-rework/controllers.RenderIndividualBlogPost (4 handlers)
[GIN-debug] POST   /blog/post                --> github.com/gin-gonic/gin.CustomRecoveryWithWriter.func1 (3 handlers)
[GIN-debug] POST   /blog/post/comment        --> gitlab.com/rawleyifowler/site-rework/controllers.CreateComment (4 handlers)
[GIN-debug] GET    /admin/                   --> gitlab.com/rawleyifowler/site-rework/utils.ServePage.func1 (4 handlers)
[GIN-debug] POST   /admin/login              --> gitlab.com/rawleyifowler/site-rework/controllers.HandleLogin (4 handlers)
[GIN-debug] GET    /admin/post               --> gitlab.com/rawleyifowler/site-rework/utils.ServePage.func1 (5 handlers)
[GIN-debug] POST   /admin/post               --> gitlab.com/rawleyifowler/site-rework/controllers.CreatePost (5 handlers)
[GIN-debug] GET    /                         --> gitlab.com/rawleyifowler/site-rework/utils.ServePage.func1 (4 handlers)
[GIN-debug] GET    /resume                   --> gitlab.com/rawleyifowler/site-rework/utils.ServePage.func1 (4 handlers)
[GIN-debug] GET    /contact                  --> gitlab.com/rawleyifowler/site-rework/utils.ServePage.func1 (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
signal: interrupt
