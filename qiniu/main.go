package main

import (
	"context"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func main() {
	download()
}

func download() {
	domain := "http://q2p54f9at.bkt.clouddn.com"
	key := "a/32/3232/1.jpg"

	publicAccessURL := storage.MakePublicURL(domain, key)
	fmt.Println(publicAccessURL)
}

func upload() {
	accessKey := "oLvb2EEL238a-jivZsQpZn7pqJvncY2HqzWA9VnS"
	secretKey := "vvW9hwoML0l1HBrT3x1KFrtp4ai05OKDFAy2jKrc"

	var localFile, bucket, key string
	localFile = "/Users/liuwei/Downloads/1.jpg"
	bucket = "levonfly1"
	key = "a/32/3232/1.jpg"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
