###Mac 下go环境部署及vscode使用

####1.下载

#####1.1 go

golang官网：[golang](https://golang.org/dl/)

国内访问：

[golang](https://golang.google.cn/dl/)

[go语言中文网](https://studygolang.com/dl)


![dl](/Users/banz/Desktop/gopkg.jpg)

#####1.2 vscode

Visual Studio Code:[vscode](https://code.visualstudio.com)

####2.安装

#####2.1go包解压
`tar -C /usr/local -xzf go1.12.7.darwin-amd64.tar.gz`
在解压时遇到无法创建文件的情况（本人第一次用Mac，小声bb），经查，原因是在/usr/local目录下新增文件夹需要密码。Mac下/usr/local目录默认是对于Finder隐藏的，打开Finder，使用command+shift+G ，在弹出的目录中填写/usr/local就可以进入路径，然后手动新建一个名为go的文件夹，再执行`tar -C /usr/local -xzf go1.12.7.darwin-amd64.tar.gz`就可以了。

#####2.2Homebrew安装
易升级、管理

	brew install go

####3.配置

#####3.1Go环境变量

	1.打开配置文件
	vim ~/.bash_profile
	2.添加环境变量
	export GOROOT=/usr/local/Cellar/go/1.12.7/libexec
	export GOPATH=/Users/xxx/go
	export GOBIN=$GOPATH/bin
	export PATH=$PATH:$GOBIN:$GOROOT/bin
* GOROOT： go安装目录
* GOPATH：go工作目录
* GOBIN：go可执行文件目录
* PATH：将go可执行文件加入PATH中，使GO命令与我们编写的GO应用可以全局调用

编辑完成保存退出
	
	3.使配置生效
	source ~/.bash_profile
	4.检查
	go env或go version
	
###4.VsCode

在GOPATH/src下新建github.com/和golang.org/x/两个文件目录，进入/x/目录下使用git clone https://github.com/golang/tools.git tools下载tools，此时会有go-outline、gocode等包不存在的情况，使用：

	go get -u -v github.com/ramya-rao-a/go-outline
	go get -u -v github.com/acroca/go-symbols
	go get -u -v github.com/nsf/gocode
	go get -u -v github.com/rogpeppe/godef
	go get -u -v github.com/zmb3/gogetdoc
	go get -u -v github.com/golang/lint/golint
	go get -u -v github.com/fatih/gomodifytags
	go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
	go get -u -v sourcegraph.com/sqs/goreturns
	go get -u -v github.com/cweill/gotests/...
	go get -u -v github.com/josharian/impl
	go get -u -v github.com/haya14busa/goplay/cmd/goplay
获取.

安装：

	go install github.com/ramya-rao-a/go-outline
	
	go install github.com/acroca/go-symbols
	
	go install golang.org/x/tools/cmd/guru
	
	go install golang.org/x/tools/cmd/gorename
	
	go install github.com/josharian/impl
	
	go install github.com/rogpeppe/godef
	
	go install github.com/sqs/goreturns
	
	go install github.com/golang/lint/golint
	
	go install github.com/cweill/gotests/gotests
