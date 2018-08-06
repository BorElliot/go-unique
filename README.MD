# 生成唯一字符串

生成windows可执行文件

`OOS=windows GOARCH=386 go build -o random_windows.exe main.go`

生成Centos二进制文件
`GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o random_linux main.go`






### 参考

[如何生成跨平台二进制文件](https://stackoverflow.com/questions/12168873/cross-compile-go-on-osx)

[windows Compiling](https://github.com/golang/go/wiki/WindowsCrossCompiling)