# EK

kubernetes工具

## 安装
项目使用golang 1.15，开始之前请先安装golang。
```shell
$ git clone https://github.com/ripdent0532/ek.git
$ cd ek
$ make GOOS=linux
```
##用法
```shell
k8s终端工具

Usage:
  ek [flags]
  ek [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  images      获取k8s中指定命名空间中的容器镜像
  migrate     迁移k8s中的镜像

Flags:
  -h, --help   help for ek

```