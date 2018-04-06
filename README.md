# 1、首先上效果
![登录页面](https://github.com/winlion/restgo-admin/blob/master/asset/images/1.png)
![登入首页](https://github.com/winlion/restgo-admin/blob/master/asset/images/2.png)
# 2、如何使用
## 2.1、使用如下指令克隆
```
cd $GOPATH/src
git get github.com/winlion/deepthink
```
你将在gopath/src/github.com/winlion/下得到deepthink目录  
进入目录  
```
cd deepthink
```

## 2.2、数据库
新建数据库名称为restgo-admin,编码为utf-8  
将restgo-admin.sql导入到数据库中  
## 2.3、初始化依赖包
使用前先使用如下指令安装指令安装文件  
···
go get github.com/go-sql-driver/mysql  
go get -v -u github.com/alecthomas/log4go  
go get github.com/gin-gonic/gin  
go get github.com/go-sql-driver/mysql  
go get github.com/go-xorm/xorm  
go get github.com/tommy351/gin-sessions  
···

## 2.4、启动
使用前先使用如下指令启动应用  
···
go run main.go  
···
初始化账号和密码
账号keytoany@qq.com
密码123456
使用前先使用如下指令打包应用  
···
build.bat  
···
# 3、FAQ
## 3.1 如何安装开发环境
如果你使用的是vscode,安装问题请访问  
···
https://www.cnblogs.com/Leo_wl/p/8242628.html  

go get github.com/nsf/gocode  
go get github.com/uudashr/gopkgs/cmd/gopkgs  
go get github.com/fatih/gomodifytags  
go get github.com/haya14busa/goplay/cmd/goplay  
go get github.com/derekparker/delve/cmd/dlv  
···
## 3.2 如何联系我
我的微信 jiepool-winlion  
我的qq 271151388
