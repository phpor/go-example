package main


// partNumber 可以从 1 开始，不能从 0 开始
// 上传的每个部分都需要有上行的content-length 头，并且需要和实际内容长度一致，所以，需要知道流的长度才能上传
// 小于 100KB的文件分片上传会报错： Your proposed upload smaller than the minimum allowed size

import (
	"github.com/phpor/aliyun-oss-go-sdk/oss"
	"os"
	"io"
	"encoding/xml"
	"fmt"
)
func main()  {

	/*
	fd,err := os.OpenFile("/tmp/test.txt", os.O_RDONLY, 0)
	bucket.PutObject("aa", fd)
	*/
/*	res, err := bucket.ListObjects()
	for _, object := range res.Objects {
		fmt.Println("Objects:", object.Key)
	}*/
	fd, err := os.OpenFile("/tmp/test.txt", os.O_RDONLY, 0)
	assertNil(err)
	stat,err := fd.Stat()
	assertNil(err)
	upload("test1.txt", fd, stat.Size())
}

func newBucket() * oss.Bucket {
	endpoint := os.Getenv("OSS_ENDPOINT")
	accesskey_id := os.Getenv("OSS_ACCESSKEY_ID")
	accesskey_secret := os.Getenv("OSS_ACCESSKEY_SECRET")
	bucket_name := os.Getenv("OSS_BUCKET")

	client, err := oss.New(endpoint, accesskey_id,accesskey_secret)
	assertNil(err)
	bucket, err := client.Bucket(bucket_name)
	assertNil(err)
	return bucket
}
func list() {
	res, err := newBucket().ListObjects()
	assertNil(err)
	for _, object := range res.Objects {
		fmt.Println("Objects:", object.Key)
	}
}

func mv(src string, dst string) error {

	return nil
}

func upload(objKey string, reader io.Reader, size int64) error {
	if size <= 1024*1024*100 {
		return newBucket().PutObject(objKey, reader)
	}
	mpUpload(objKey, reader, size)
	return nil
}


func mpUpload(objKey string, reader io.Reader, size int64) oss.CompleteMultipartUploadResult{
	bucket := newBucket()
	mp, err := bucket.InitiateMultipartUpload(objKey)
	if err != nil {
		panic(err)
	}
	var parts []oss.UploadPart
	i := 1
	part_size := int64(102400)
	for{
		upload_size := size
		if size > part_size {
			upload_size = part_size
		}
		part, err := bucket.UploadPart(mp, reader, upload_size, i)
		assertNil(err)
		parts = append(parts, part)

		println("----------------")
		printXml(part)
		i++
		lpr,err := bucket.ListUploadedParts(mp)
		assertNil(err)
		println("----------------")
		printXml(lpr)
		size -= upload_size
		if size <= 0 {
			break
		}

	}
	println("==============")

	res, err := bucket.CompleteMultipartUpload(mp, parts)
	if err != nil {
		panic(err)
	}
	printXml(res)
	return res
}
func printXml(info interface{})  {
	bytes, err := xml.Marshal(info)
	assertNil(err)
	println(string(bytes))
}
func assertNil(err error) {
	if err != nil {
		panic(err)
	}
}
