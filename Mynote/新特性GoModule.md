#Go Modules
Go modules是官方推出1.11版本之后增加的新的依赖管理模块，允许用户在$GOPATH/src外使用go.mod新建项目。在 $GOPATH/src 中，为了兼容性，Go 命令仍然在旧的 GOPATH模式下运行。从 Go 1.13 开始，模块模式将成为默认模式。

	# 设置环境变量、代理
	$ export GO111MODULE=on
	$ export GOPROXY=https://mirrors.aliyun.com/goproxy/
	# 在GOPATH/src以外创建项目
	$ mkdir project && cd project
	# 初始化后目录中会生成一个go.mod文件
	$go mod init project
	
	# 示例
	# 新建main.go，添加Gin依赖
	package main

	import "github.com/gin-gonic/gin"
	
	func main() {
	    r := gin.Default()
	    r.GET("/ping", func(c *gin.Context) {
	        c.JSON(200, gin.H{
	            "message": "pong",
	        })
	    })
	    r.Run() 
	}
	
	$ go build main.go

此时会在go.mod中引入gin等依赖项

	# 查看依赖项列表
	$go list -m all
	project
	github.com/davecgh/go-spew v1.1.1
	github.com/gin-contrib/sse v0.1.0
	github.com/gin-gonic/gin v1.5.0
	...
	
	# 查看依赖项历史版本
	$ go list -m -versions github.com/gin-gonic/gin
	# 会看到
	github.com/gin-gonic/gin v1.1.1 v1.1.2 v1.1.3 v1.1.4 v1.3.0 v1.3.0+incompatible v1.4.0 v1.5.0

更新依赖项到某个版本

	# 在依赖项后加@version
	$ go get github.com/gin-gonic/gin@v1.1.4
	或
	$ go mod edit -require="github.com/gin-gonic/gin@v1.1.4"
	$ go tidy //下载更新
	# 法二需要两步

删除未使用的依赖项

	$ go mod tidy

更多命令

	$ go mod
	