package main

import (
	"cloud-oss/store/aliyun"
)

func main() {
	aliyun.NewAliOssStore("oss-cn-shanghai.aliyuncs.com", "", "")

}
