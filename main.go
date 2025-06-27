package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// === 定义命令行参数 ===
	endpoint := flag.String("endpoint", os.Getenv("MINIO_SERVER_URL"), "MinIO server endpoint (e.g., play.min.io or 127.0.0.1:9000)")
	accessKey := flag.String("access-key", os.Getenv("MINIO_ACCESS_KEY"), "MinIO access key")
	secretKey := flag.String("secret-key", os.Getenv("MINIO_SECRET_KEY"), "MinIO secret key")
	bucket := flag.String("bucket", "shopee", "Target bucket name")
	object := flag.String("object", "ecommerce_automation.exe", "Object key (destination name in bucket)")
	file := flag.String("file", "./dist/ecommerce_automation.exe", "Local file path to upload")
	secure := flag.Bool("secure", false, "Use HTTPS to connect to MinIO")

	flag.Parse()

	// === 参数检查 ===
	if *endpoint == "" || *accessKey == "" || *secretKey == "" || *bucket == "" || *object == "" || *file == "" {
		flag.Usage()
		log.Fatal("❌ 所有参数均为必填，请检查输入或环境变量。")
	}

	// === 创建 MinIO 客户端 ===
	client, err := minio.New(*endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(*accessKey, *secretKey, ""),
		Secure: *secure,
	})
	if err != nil {
		log.Fatalf("❌ 连接 MinIO 失败: %v", err)
	}

	// === 上传文件 ===
	ctx := context.Background()
	info, err := client.FPutObject(ctx, *bucket, *object, *file, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalf("❌ 上传失败: %v", err)
	}

	fmt.Printf("✅ 上传成功: %s (%d bytes)\n", info.Key, info.Size)
}
