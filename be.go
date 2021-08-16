package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 正则路由
func main() {
	// 匹配数字格式 (\d+) 的路由，并将匹配的值存放到 :id 参数
	web.Get("/user/:id(\\d+)/", func(ctx *context.Context) {
		id := ctx.Input.Param(":id")
		ctx.WriteString(fmt.Sprintf("id: %s", id))
		/*
			127.0.0.1:8080/user/1/
			id: 1
		*/
	})
	// 匹配 \b\w+\b
	web.Get("/user/:name(\\b\\w+\\b)/", func(ctx *context.Context) {
		name := ctx.Input.Param(":name")
		ctx.WriteString(fmt.Sprintf("name: %s", name))
		/*
			127.0.0.1:8080/user/zze/
			name: zze
		*/
	})
	// 匹配任意内容
	web.Get("/:content/", func(ctx *context.Context) {
		content := ctx.Input.Param(":content")
		ctx.WriteString(fmt.Sprintf("content: %s", content))
		/*
			127.0.0.1:8080/user-zze/
			content: user-zze
		*/
	})
	// 匹配文件
	web.Get("/file/*", func(ctx *context.Context) {
		content := ctx.Input.Param(":splat")
		ctx.WriteString(fmt.Sprintf("filename: %s", content))
		/*
			127.0.0.1:8080/file/xxx.txt
			filename: xxx.txt
		*/
	})
	// 匹配文件路径和后缀
	web.Get("/files/*.*", func(ctx *context.Context) {
		path := ctx.Input.Param(":path")
		ext := ctx.Input.Param(":ext")
		ctx.WriteString(fmt.Sprintf("path: %s, ext: %s", path, ext))
	})
	/*
		127.0.0.1:8080/files/xxx.txt
		path: xxx, ext: txt

		127.0.0.1:8080/files/xxx/xxx.txt
		path: xxx/xxx, ext: txt
	*/
	web.Run(":8088")
}