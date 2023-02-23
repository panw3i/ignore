# 打包
使用gox可以轻松地打包Go应用程序，它可以跨多个操作系统和架构生成二进制文件，并将它们打包到一个压缩文件中。

```bash
go install github.com/mitchellh/gox
gox -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"
```




