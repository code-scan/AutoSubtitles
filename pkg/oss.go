package pkg

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path"
)

type OSS struct {
	EndPoint   string
	BucketName string
	Aliyun
}

func NewOSS(AK, AS, EndPoint, Bucket string) *OSS {
	oss := &OSS{
		EndPoint:   EndPoint,
		BucketName: Bucket,
	}
	oss.AK = AK
	oss.AS = AS
	return oss
}
func (o *OSS) UploadFile(filepath string) string {
	client, err := oss.New(o.EndPoint, o.AK, o.AS)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	storageType := oss.ObjectStorageClass(oss.StorageStandard)

	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 上传字符串。
	filename := path.Base(filepath)
	fileIO, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	err = bucket.PutObject(filename, fileIO, storageType, objectAcl)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("http://%s.%s/%s", o.BucketName, o.EndPoint, filename)
}
