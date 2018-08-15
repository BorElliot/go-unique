# 生成唯一字符串

生成windows可执行文件

`OOS=windows GOARCH=386 go build -o random_windows.exe main.go`

生成Centos二进制文件
`GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o random_linux main.go`



## 提高数据库写入速度

[https://blog.csdn.net/m0_38132420/article/details/78964433](https://blog.csdn.net/m0_38132420/article/details/78964433)

[https://juejin.im/post/5a11a3c76fb9a04512388a50](https://juejin.im/post/5a11a3c76fb9a04512388a50)



### 参考

[如何生成跨平台二进制文件](https://stackoverflow.com/questions/12168873/cross-compile-go-on-osx)

[windows Compiling](https://github.com/golang/go/wiki/WindowsCrossCompiling)