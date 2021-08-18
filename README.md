# EK

kubernetes工具

## 安装
项目使用golang 1.15，开始之前请先安装golang。
```shell
$ git clone https://github.com/ripdent0532/ek.git
$ cd ek
$ make GOOS=linux
```
## 用法
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
  version     查看版本信息

Flags:
  -h, --help   help for ek

Use "ek [command] --help" for more information about a command.
```

## 配置
```yaml
repositories:
  - host: xxxxx.dkr.ecr.cn-northwest-1.amazonaws.com.cn
    username: AWS
    password:
    type: AWS
```


| 项目  | 必填 | 备注 |
| ------------- | ------------- | ------------- |
| host  | 是  | docker仓库地址 |
| username  | 是  | 用户名 |
| password | 否
| type | 否 | `AWS`:类型为AWS时，不需要填写password