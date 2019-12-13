##Git

* 以下[]内都是用户自定义内容

工作区：项目目录

暂存区：临时存放被修改文件

版本库：存放提交记录

###1.本地库及基本使用

####1.1配置用户名、地址

	$git config --global user.name "[name]"
	$git config --global user.email "[email]"

####1.2创建版本库

	$mkdir [filename]
	$cd [filename]
	$git init

####1.3添加文件至暂存区

	单个添加
	$git add [file]
	多个添加
	$git add [file1] [file2]
	or
	$git add .
	or
	$git add [dirname/.]
	
####1.4将文件提交至仓库(全部提交)

	$git commit -m "[提交注释]"
	
#####1.5基本使用

	查看当前状态
	$git status
	查看文件修改内容
	$git diff [file]
	查看提交日志
	$git log
	简化信息
	$git log --pretty=oneline
	查看历史操作
	$git reflog

####1.6删除文件

	删除、提交
	$git rm [file1] or $rm [file1]
	$git commit -m "[rm file]"
	删错撤销
	$git checkout -- [file1]

###2.版本回退
HEAD相当于`commit id`的指针

	两种回退方式
	$git reset --hard HEAD^
	(HEAD为当前版本，HEAD^表示上一版本，HEAD^^表示上上版本，HEAD~100表示往上100版本)
	or
	$git reset --hard [commit id]
	
####2.1撤销修改

	应用场景
	1.错误修改未存入暂存区,丢弃工作区修改
	$git checkout -- [file]
	2.错误修改已存入暂存区
	$git reset HEAD [file]
	$git checkout -- [file]
	3.错误修改已提交至「本地」版本库,此时撤销相当于版本回退
	$git reset --hard [commit id]
	4.错误修改已提交至「远程」版本库
	GG
	
###3.远程仓库
	创建SSH Key
	$ssh-keygen -t rsa -C "email@example.com"
	完成后用户主目录里找到.ssh目录，里面有id_rsa和id_rsa.pub两个文件，
	这两个就是SSH Key的秘钥对，id_rsa是私钥，不能泄露出去，id_rsa.pub是公钥，
	可以放心地告诉任何人.
	登陆GitHub进入“Setting”--“SSH key”，填上任意title，
	在Key文本框里粘贴id_rsa.pub文件的内容.
	
	关联仓库
	$git remote add origin git@server-name:path/repo-name.git
	origin可自定义，本地对远程仓库的名字
	
	查看关联远程库
	$git remote
	$git remote -v
	
	取消关联
	$git remote remove origin
	
	本地内容推送至远程库
	第一次推送
	$git push -u origin master
	-u顺便关联master分支
	后续推送
	$git push
	
	从远程库克隆
	$git clone git@github.com:path/repo-name.git
	
	在本地创建远程库分支
	$git checkout -b dev origin/dev
	
	建立本地分支和远程分支的关联(未执行上述步骤，git pull再关联)[未实验]
	$git branch --set-upstream-to=origin/branch-name branch-name

	抓取分支
	$git pull
	
	git2.9.0后抓取合并分支[未实验]
	$git pull origin master --allow-unrelated-histories
	
	

###4.分支管理

	创建分支
	$git branch [name]
	$git checkout [name]
	合并
	$git checkout -b [name]（-b表示创建并切换）
	
	查看当前分支
	$git branch
	切换分支
	$git checkout [name]
	
	合并分支，git中的合并分支指的是将[name]分支合并到当前分支
	$git merge [name]
	
	通常，合并分支时，如果可能，Git会用Fast forward模式，但这种模式下，删除分支后，会丢掉分支信息。合并分支时，加上--no-ff参数就可以用普通模式合并，合并后的历史有分支，能看出来曾经做过合并，而fast forward合并就看不出来曾经做过合并
	$git merge --no-ff -m "[msg]" [branchname]
	
	删除分支
	$git branch -d [name]
	强制删除分支，git默认分支未合并不允许删除，一般用于已提交但未合并
	$git branch -D [name]
	
	Bug分支
	情景：当前分支工作未完成，需临时进行文件修改。此时需要，存储临时改动
	$git stash
	将当前分支的工作现场“储藏“起来，再新建临时工作需要的分支进行Bug修复之类的工作
	$git stash 可以多次使用，可以使用
	$git stash list 查看具体需要恢复哪个工作现场
	恢复改动
	方法一 精准恢复
	$git stash apply [index] stash为stash中的index，可选填，不填时逐级恢复；使用此方法恢复的改动，未删除stash，使用
	$git stash drop [index]
	方法二 恢复最新stash
	$git stash pop 恢复同时删除stash

###5.设置.gitignore
.gitignore文件用于忽略文件，某些不想让git管理的文件就可以在该文件中编辑规则，一行为一条规则。

* 所有空行或者以注释符号 ＃ 开头的行都会被 git 忽略，空行可以为了可读性分隔段落，# 表明注释。
* 第一个 / 会匹配路径的根目录，举个栗子，”/*.html”会匹配”index.html”，而不是”d/index.html”。
* 通配符 * 匹配任意个任意字符，? 匹配一个任意字符。需要注意的是通配符不会匹配文件路径中的 /，举个栗子，”d/*.html”会匹配”d/index.html”，但不会匹配”d/a/b/c/index.html”。
* 两个连续的星号 ** 有特殊含义：
	- 以 **/ 开头表示匹配所有的文件夹，例如 **/test.md 匹配所有的test.md文件。
	- 以 /** 结尾表示匹配文件夹内所有内容，例如 a/** 匹配文件夹a中所有内容。
	- 连续星号 ** 前后分别被 / 夹住表示匹配0或者多层文件夹，例如 a/**/b 匹配到 a/b 、a/x/b 、a/x/y/b 等。
* 前缀 ! 的模式表示如果前面匹配到被忽略，则重新添加回来。如果匹配到的父文件夹还是忽略状态，该文件还是保持忽略状态。如果路径名第一个字符为 ! ，则需要在前面增加 \ 进行转义。