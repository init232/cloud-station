package aliyun

import (
	"cloud-oss/store"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	_ store.Uploader = &AliOssStore{}
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSercet string
}

func (o *Options) Validate() error {
	//校验传参
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSercet == "" {
		return nil, error("")
	}

}

// AliOssStore 对象构造函数
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSercet)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		clinet: c,
	}, nil
}

type AliOssStore struct {
	clinet *oss.Client
}

func (s *AliOssStore) Upload(BucketName string, ObjectKey string, FileName string) error {
	//获取bucket对象
	bucket, err := s.clinet.Bucket(BucketName)
	if err != nil {
		fmt.Println("获取bucket失败", err.Error())
	}
	//上传文件到bucket
	if err = bucket.PutObjectFromFile(ObjectKey, FileName); err != nil {
		return err
	}
	//打印下载链接
	DownLoadUrl, err := bucket.SignURL(FileName, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL: %s 请在一天之内下载\n", DownLoadUrl)
	return nil
}
