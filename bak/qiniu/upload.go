package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func main() {
	dir := "/var/lib/kinema/preview"
	uploadDir(dir)

	dir = "/var/lib/oceans/preview"
	uploadDir(dir)
}

func uploadDir(dir string) {
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		var key string
		if dir == "/var/lib/kinema/preview" {
			key = makeUrl1(path)
		} else {
			key = makeUrl2(path)
		}

		fmt.Println(path, key)

		uploadFile(path, key)
		return nil
	})
}

func makeUrl1(path string) string {
	Prefix := "media-preview"
	PreivewRoot := "/var/lib/kinema/preview"
	if strings.HasPrefix(path, PreivewRoot) {
		path = path[len(PreivewRoot):]
	}
	return filepath.Join(Prefix, path)
}

func makeUrl2(path string) string {
	Prefix := "upload-preview"
	PreivewRoot := "/var/lib/oceans/preview"
	if strings.HasPrefix(path, PreivewRoot) {
		path = path[len(PreivewRoot):]
	}
	return filepath.Join(Prefix, path)
}

func uploadFile(localFile, key string) {

	// accessKey := "oLvb2EEL238a-jivZsQpZn7pqJvncY2HqzWA9VnS"
	// secretKey := "vvW9hwoML0l1HBrT3x1KFrtp4ai05OKDFAy2jKrc"
	// bucket := "levonfly1"

	accessKey := "u3SrAzgMe20hwv3Xz5aXwT5qPpxT6yXrnNxPjakY"
	secretKey := "lmGQ3Nk2k4cgsB339SZ7vbh55wva06p23sM01pNW"
	bucket := "fhyx-preview"
	zone := storage.ZoneHuanan

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &zone
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
