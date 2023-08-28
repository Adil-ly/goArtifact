# goArtifact 工具

##### 安装拓展包命令

`go get -u github.com/Adil-ly/goArtifact@latest`

##### 使用

###### _1、初始化_

```
util := goArtifact.NewArtifactUtil()
```

###### _2、调用对应的方法_


```
// 验证某个元素是否存在int切片中
util.IsInSliceWithInt() 
util.IsInSliceWithInt8()    
util.IsInSliceWithInt32()
util.IsInSliceWithInt64()

-------
// 验证某个元素是否存在string切片中
util.IsInSliceWithString()

-------
// int类型的去重
util.ReDupWithInt() 
util.ReDupWithInt8()
util.ReDupWithInt32()
util.ReDupWithInt64()

-------
// string类型的去重
util.ReDupWithString()

-------

...后续工具以提供快捷开发使用
```


