# DCD

DCD是一个用来做简单数据收集的服务器软件.(通过Go的Echo Web Framework开发) 比如统计软件使用率什么的. 开发这个项目的初衷是想收集`BangumiOpen`功能使用的情况.
DCD的特点就是可以无外部依赖部署, 可以不需要数据库就能部署, 而且配置简单, 通过简单的YAML文件, 就可以指定路由以及数据存贮方式.

## Handlers

目前提供了以下统计器:

- 计数统计