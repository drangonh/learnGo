"# learnGo" 

*如果遇到运行报错：*
`can't load package: package learnGo/tree/entry: malformed module path "learnGo/tree/entry": missing dot in first path element` 

请设置go env -w GO111MODULE=off

*GOPATH目录下的src保存的是每一个项目*

##### 扩展已有包
*包需要注意的内容*
- 为结构定义的包必须放在同一个包内
- 可以是不同文件

*如何定义包*
* 定义别名
* 使用组合