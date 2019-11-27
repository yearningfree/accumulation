#Linux

##目录结构

	系统启动必须：
	/boot /etc /lib /sys
	指令集合：
	/bin /sbin
	外部文件管理：
	/dev /media /mnt
	临时文件：
	/run /lost+found /tmp
	账户：
	/root /home /usr /usr/bin /usr/sbin /usr/src
	运行过程中使用：
	/var /proc
	扩展用：
	/opt /srv

***

	/root 系统管理员，超级用户主目录
	
	/bin Binary，存放常用程序和指令
	
	/boot 存放启动Linux时使用的核心文件，包括一些连接文件以及镜像文件
	
	/dev Device，存放Linux的外部设备
	
	/etc 存放所有的系统管理所需的配置文件和目录
	
	/home 用户主目录，在Linux中每个用户都有一个自己的目录，一般以用户名账号命名
	
	/lib 存放系统最基本的动态连接共享库
	
	/usr 存放用户大多数应用程序和文件
	
	/usr/bin 系统用户使用的应用程序与指令
	
	/usr/sbin 超级用户使用的比较高级的管理程序和系统守护程序
	
	/usr/src 内核源代码默认的放置目录
	
	/media 类似Windows的其他设备，U盘、光驱等设备识别后挂载在此目录
	
	/tmp 存放一些临时文件
	
	/var 存放经常修改的数据，比如日志文件
	
	/mnt 临时挂在别的文件系统，可以将光驱挂载在此目录，然后直接进入查看光驱里的内容
	
	/opt 默认是空的，可以安装额外软件
	
	/proc 虚拟目录，用于管理内存空间，是系统内存映射，此目录不在硬盘上而在内存里
	
	/sbin 存放只有系统管理员能使用的程序和指令
	
	/selinux Radhat/CentOS特有，SeLinux是一个安全机制，类似于windows防火墙
	
	/srv 存放服务启动之后需要提取的数据
	
	/sys 这是linux2.6内核的一个很大的变化。该目录下安装了2.6内核中新出现的一个文件系统 sysfs 。sysfs文件系统集成了下面3种文件系统的信息：针对进程信息的proc文件系统、针对设备的devfs文件系统以及针对伪终端的devpts文件系统。
	该文件系统是内核设备树的一个直观反映。
	当一个内核对象被创建的时候，对应的文件和目录也在内核对象子系统中被创建
	
	/run 是一个临时文件系统，存储系统启动以来的信息。当系统重启时，这个目录下的文件应该被删掉或清除。如果你的系统上有 /var/run 目录，应该让它指向 run
	
	/lost+found 一般情况下为空的，系统非法关机后，这里就存放一些文件
	
	
	

##关机

	sync 必须先将数据从内存同步到硬盘中
	
	shutdown 关机指令，关闭机器
	
	shutdown -h 10 'This server will shutdown after 10 mins' 通知所有登陆用户，计算机将在10分钟后关机
	
	shutdown -h +10 10分钟后关机
	
	shutdown -h now 立刻关机
	
	shutdown -r now 重启系统
	
	shutdown -r +10 10分钟后重启系统
	
	reboot 重启 等于 shutdown -r now
	
	reboot --halt 停止机器
	
	reboot -p 关闭机器
	
	halt 停止cpu所有功能，但仍保持通电，处于低层维护状态，但有些情况会完全关闭
	
	halt -p 关闭机器
	
	halt --reboot 重启机器
	
	poweroff 关闭机器
	
	poweroff --halt 停止机器
	
	poweroff --reboot 重启机器
	
##文件权限

* 所有者权限（属主 user）：文件所有者能够进行的操作
* 组权限（属组 group）：文件所属用户组能够进行的操作
* 外部权限（其他用户 others）：其他用户能够进行的操作

第一列：

1. 文件类型：[d]目录 [-]文件
2. 属主权限：rwx 读 写 执行
3. 属组权限：r-x 读 - 执行
4. 其他用户权限 ：r-x 读 - 执行

第二列：若是文件则代表该文件的硬链接数。若是目录，则代表该目录下的子目录数。无子目录的目录也会显示有2个子目录，因为任何目录都有，当前目录[.]和子目录[..]

第三列：当前文件所属用户名，若该用户被删除，则会显示该用户删除前的用户id

第四列：当前文件所属的用户组

第五列：若是文件则代表该文件的大小，若是目录则代表该目录的大小（不包括目录下的子目录和文件的大小）

第六列：该文件最近修改或者查看的时间

第七列：文件名称


例

	drwxr-xr-x 2 root  root  4096 Oct 13 14:46 btmp
	
---
###更改文件属性

	chown 属主名 文件名        更改文件属主
	chown 属主名:数组名 文件名  同时更改属组
	chgrp groupname filename 更改文件属组
	chgrp -R groupname filename 递归更改文件属组，目录用
	
	rwx 对应十进制 4 2 1
	文件三个用户权限user、group、others，简称u、g、o；另外a则代表all;使用`+`加入、`-`除去或`=`设定可以相应操作权限
	chmod [u、g、o、a] [+、-、=] [r、w、x] [filename]     任意组合，修改文件属性
	
##文件与目录管理

	ls 列出目录
	语法：
	ls [dir] 
	参数：
	-a ：全部的文件，连同隐藏档(开头为 . 的文件) 一起列出来
	-d ：仅列出目录本身，而不是列出目录内的文件数据
	-l ：长数据串列出，包含文件的属性与权限等等数据
---

	cd 	切换目录
---

	pwd 显示当前所在目录
	pwd -P 显示真是目录，而非link路径
---

	mkdir 创建新目录
	语法：
	mkdir [-mp] 目录名
	参数：
	-m：配置权限
	-p：递归创建目录
	
	mkdir -m 711 test
	mkdir -p test1/test2/test3
---

	rmdir 删除空目录
	语法：
	rmdir [-p] 目录名
	参数：
	-p：连同上级空目录一起删除	
	
	rmdir test
	rmdir -p test1/test2/test3
---

	cp 复制文件或目录
	cp [-adfilprsu] 来源档(source) 目标档(destination)
	参数：
	-a：相当於 -pdr 的意思
	-d：若来源档为连结档的属性(link file)，则复制连结档属性而非文件本身
	-f：为强制(force)的意思，若目标文件已经存在且无法开启，则移除后再尝试一次
	-i：若目标档(destination)已经存在时，在覆盖时会先询问动作的进行
	-l：进行硬式连结(hard link)的连结档创建，而非复制文件本身
	-p：连同文件的属性一起复制过去，而非使用默认属性(备份常用)
	-r：递归持续复制，用於目录的复制行为
	-s：复制成为符号连结档 (symbolic link)，亦即『捷径』文件
	-u：若 destination 比 source 旧才升级 destination 
---

	rm 移除文件或目录
	语法：
	rm [-fir] 文件或目录
	参数：
	-f：忽略不存在文件，不提示警告
	-i：互动模式，删除前会询问
	-r：递归删除（危险）
---

	mv 移动文件与目录可改文件名
	语法：
	mv [-fiu] source destination
	参数：
	-f：如果目标文件已经存在，不会询问而直接覆盖
	-i：若目标文件 (destination) 已经存在时，就会询问是否覆盖
	-u：若目标文件已经存在，且 source 比较新，才会升级 (update)

##内容查看
	cat 从第一行开始显示文件内容
	语法：
	cat [-AbEnTv]
	参数：
	-A：相当於 -vET 的整合选项，可列出一些特殊字符而不是空白而已
	-b：列出行号，仅针对非空白行做行号显示，空白行不标行号
	-E：将结尾的断行字节 $ 显示出来
	-n：列印出行号，连同空白行也会有行号，与 -b 的选项不同
	-T：将 [tab] 按键以 ^I 显示出来
	-v：列出一些看不出来的特殊字符
___
	tac
	tac与cat命令刚好相反，文件内容从最后一行开始显示	
---
	head 取出文件前面几行
	语法：
	head [-n number] 文件 
	参数：
	-n：后面接数字，代表显示几行的意思
---
	tail 取出文件后面几行
	语法：
	tail [-n number] 文件 
	参数：
	-n：后面接数字，代表显示几行的意思
	-f：表示持续侦测后面所接的档名，要等到按下[ctrl]-c才会结束tail的侦测

##用户管理

	useradd 选项 用户名
	选项：
	-c：comment 指定一段注释性描述。
	-d：目录 指定用户主目录，如果此目录不存在，则同时使用-m选项，可以创建主目录。
	-g：用户组 指定用户所属的用户组。
	-G：用户组，用户组 指定用户所属的附加组。
	-s：Shell文件 指定用户的登录Shell。
	-u：用户号 指定用户的用户号，如果同时有-o选项，则可以重复使用其他用户的标识号

跳

##yum
语法：
	
	yum [options] [command] [package..]
	options:
	-y 安装过程提示选择全部“yes”
	-q 不显示安装过程
	
	常用命令：
	yum list 列出所有可安装的软件清单
	yum check-update 列出可更新软件
	yum update更新所有软件
	yum install <package> 安装指定软件
	yum update <package> 更新指定软件包
	yum remove <package> 删除指定软件包
	yum serach <package> 查找指定软件包
	yum clean packages 清楚缓存目录下的软件包
	yum clean headers 清除缓存目录下的 headers
	yum clean oldheaders 清除缓存目录下旧的 headers
	yum clean all 全部清除
	yum makecache 生成缓存

##管道和过滤器
把两个命令连起来用，一个命令的输出作为另一个命令的输入，这就是管道。在两个命令之间用竖线`|`连接。

能够接收数据，过滤（处理或筛选）后再输出的工具，称为过滤器。

grep是一个文本搜索工具，可以使用正则表达式，并返回匹配的行。

	语法：grep [options] pattern file
	例如：grep -i 'hello world' menu.h main.c
	options有很多，-i不匹配大小写；-n输出匹配行及行号；-l输出匹配的行所在文件名；-v输出不匹配的行；-c输出匹配的总行数等。

结合管道

	ls -l | grep -i "carol.*aug"
	-rw-rw-r--   1 carol doc      1605 Aug 23 07:35 macros


##进程管理

	ps 查看正在运行的进程
	参数:
	-a 显示所有用户的所有进程
	-x 显示无终端进程
	-f 显示更多信息
	-u 显示更多信息
	-e 显示所有进程

	kill -9 PID 结束进程
	
	top 类似于windows的任务管理器
	
	创建后台进程在命令后加`&`
	
	jobs -l 查看当前任务包含的进程ID，显示任务号、进程id、运行状态、启动任务的命令
	
	fg %jobnumber 将后台任务调到前台jobnumber是任务号不是pid
	
	bg %jobnumber 将前台任务放到后台处理，先
	

##Shell脚本
第一行约定标记`#！`，告诉系统使用什么解释器来执行

	#!/bin/bash
###定义变量
定义变量不需要加`$`，命名只能使用字母、数字和下划线，首个字符不能以数字开头

变量名和等号之间不能有空格

	var="ab"
	#允许重新定义变量
	var="abc"

###赋值
	var="abc" 显示赋值
	for var in 'aaa' 语句赋值
		
###使用变量
	echo $var 
	echo ${var} 加花括号方便解释器识别变量边界

###只读变量
使用readonly关键字
	
	str="abc"
	readonly str

先赋值再设置只读，否则就相当于修改只读变量，会报错

###删除变量
unset命令

	unset str

变量删除后不可使用，unset不能删除只读变量
	
###注释
以`#`开头，单行注释；可以通过定义函数而不调用的方式实现注释效果

多行注释

	:<<EOF
	注释内容。。。
	注释内容。。。
	注释内容。。。
	EOF
	
	EOF也可以使用其他符号
	:<<'
	注释内容。。。
	注释内容。。。
	'
	
	:<<!
	注释内容。。。
	注释内容。。。
	!
	
###字符串
可以单引号也可以双引号
####单引号

	name='my name' 单引号里的任何字符都会原样输出，单引号内的变量无效
	无法在单引号内通过转义符再使用单引号
####双引号
	name='my name'
	str="${name} is \"xxx\"!"
	双引号内可以有变量，可以使用转义符
####拼接字符串

	name='my name'
	str="${name} is \"xxx\"!"
	str1="hello,world"
	echo $str $str1
	#输出 my name is "xxx"! hello,world
	
	name="abc"
	str='my name is '${name}'!'
	echo $str 
	#输出 my name is abc!

####获取字符串长度
	string="abcd"
	echo ${#string} #输出 4

####提取子字符串
从字符串第 2 个字符开始截取 4 个字符

	str="abcdefg"
	echo ${string:1:4} # 输出 bcde

####查找子字符串
查找字符e、f位置，哪个先出现就计算哪个

	str="abcdefg"
	echo `expr index "$name" fe`
	#输出5，先找到e
此处使用的是反引号

###数组

bash支持一维数组（不支持多维数组），并且没有限定数组的大小。数组元素的下标由 0 开始编号。获取数组中的元素要利用下标，下标可以是整数或算术表达式，其值应大于或等于 0

####定义数组

用括号表示，元素间用空格分割

	array=(n1 n2 n3 ... n)
	
	array0=(
	n1
	n2
	n3)
	
	array1[0]=val
	array1[2]=val1
	array1[4]=val2
	#下标可以不连续
	
####读取数组

	${数组名[下标]}#单个读取
	${数组名[@]}#@全部读取
	
####获取数组长度

	length=${#array[@]}
	或者
	length=${#array[*]}
	
	单个元素长度
	length=${#array[n]}

###传递参数
获取参数格式为`$n`，n代表数字，0是执行的文件名，从1开始是第1个参数

	$# 传递到脚本的参数个数
	$* 引用所有参数，把所有参数当1个参数使用
	#@ 引用所有参数，单个区分开来

###基本运算符
expr表达式工具，能够完成表达式求值操作

表达式和运算符之间要有空格

	#运算符
	#a=10,b=20
	`expr $a + $b`#30，加
	`expr $a - $b`#-10，减
	`expr $a \* $b`#200，乘，mac中的乘号不需要\转义
	`expr $b / $a`#2，除
	`expr $b % $a`#0，取余
	a=$b`#赋值
	[ $a == $b ]#相等，返回false，注意空格
	[ $a != $b ]#相等，返回true

###关系运算符

	[ $a -eq $b ] #是否相等，返回false
	[ $a -ne $b ] #是否不相等，返回true
	[ $a -gt $b ] #左边是否大于右边，返回false
	[ $a -lt $b ] #左边是否小于右边，返回true
	[ $a -ge $b ] #左边是否大于等于右边，返回false
	[ $a -ne $b ] #左边是否小于等于右边，返回true

###布尔运算符

	[ ! false] #非运算，true
	[ $a -lt 20 -o $b -gt 100 ] #或运算，true
	[ $a -lt 20 -a $b -gt 100 ] #与运算，false
	
###逻辑运算符

	[[ $a -lt 20 && $b -gt 100 ]] #逻辑的AND，false
	[[ $a -lt 20 && $b -gt 100 ]] #逻辑的OR，false
	
###字符串运算符

	#a="abc",b="def"
	[ $a = $b ] #false,字符串是否相等，相等返回true
	[ $a != $b ] #false,字符串是否相等，不相等返回false
	[ -z $a ] #false，字符串长度是否为0，为0返回 true
	[ -n "$a" ] #true，字符串长度是否为0，不为0返回 true
	[ $a ] #true，字符串是否为空，不为空返回 true

###文件测试运算符
	