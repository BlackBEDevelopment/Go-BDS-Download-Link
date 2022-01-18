# Go-BDS-Download-Link
Golang包，用于获取Bedrock Server最新版本下载链接和版本号

# 如何使用
## 安装
``` shell
go get -u github.com/BlackBEDevelopment/Go-BDS-Download-Link
```
## 例子
``` go
package main

import (
	"fmt"

	"github.com/BlackBEDevelopment/Go-BDS-Download-Link"
)

func main() {
	DownloadLink, Version, err := BDSDownloadLink.GetWindows()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(DownloadLink, Version)
}
```