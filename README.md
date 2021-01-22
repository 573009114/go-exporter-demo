# go-exporter-demo

### 开启模块，配置国内代理
#### 修改profile环境变量
```
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```
保存退出，即时生效使用 source /etc/profile 

go evn 看到GO111MODULE=“on” 即为成功。

### 项目使用module 
#### 在终端中切换目录到项目根目录
```
1)初始化  go mod init 初始化项目使用mod ,此时项目根目录下会生成go.mod
2)检测依赖 go mod tidy 检测当前项目所使用的依赖项目，并增加进go.mod，同时生成go.sum 包含所有依赖包。
3)下载依赖 go mod download 下载module .
4)导入依赖 go mod vendor 依赖导下项目.
```