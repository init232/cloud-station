package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	Endpoint        = ""
	AccessKeyId     = ""
	AccessKeySecret = ""
	BucketName      = ""
	UploadFile      = ""
)

// 实现文件上传
func Upload(filepath string) error {
	//1. 实例化客户端
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		panic("oss连接失败" + err.Error())
	}
	//2. 获取bucket对象
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("获取bucket失败", err.Error())
	}
	//3. 上传文件到bucket
	if err = bucket.PutObjectFromFile(filepath, filepath); err != nil {
		return err
	}
	//4. 打印下载链接
	DownLoadUrl, err := bucket.SignURL(filepath, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL: %s 请在一天之内下载\n", DownLoadUrl)

	return nil
}

// 参数和上传文件合法性检查
func validate() error {
	if Endpoint == "" || AccessKeyId == "" || AccessKeySecret == "" {
		return fmt.Errorf("参数为空")
	}
	if UploadFile == "" {
		return fmt.Errorf("上传文件为空")
	}
	return nil
}

func LoadParams() {
	flag.StringVar(&Endpoint, "Endpoint", "", "地域节点")
	flag.StringVar(&AccessKeyId, "key", "", "key")
	flag.StringVar(&AccessKeySecret, "secret", "", "secret")
	flag.StringVar(&BucketName, "bucketname", "", "bucket名称")
	flag.StringVar(&UploadFile, "file", "", "上传的文件")
	flag.Parse()
}

func Usage() {
	fmt.Fprintf(os.Stderr, "cloud-station version:0.0.1 Usage: cloud-station [-help]\n")
	flag.PrintDefaults()
}

func main() {
	Usage()
	//参数加载
	LoadParams()
	//参数验证
	if err := validate(); err != nil {
		fmt.Println("参数校验失败")
		os.Exit(1)
	}
	if err := Upload(UploadFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("文件上传成功")
	}
}
