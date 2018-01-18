package main

import (
	"context"
	"encoding/base64"
	"flag"
	"log"
	"os"
)

func main() {
	// project := flag.String("project", "", "The Google Cloud Platform project ID. required.")
	cmd := flag.String("cmd", "", "cmd is required. cmd is upload or download or printkey")
	bucket := flag.String("bucket", "", "bucket is required.")
	object := flag.String("object", "", "object is required.")
	flag.Parse()
	for _, f := range []string{"cmd", "bucket", "object"} {
		if flag.Lookup(f).Value.String() == "" {
			log.Fatalf("The %s flag is required.", f)
		}
	}

	ctx := context.Background()
	s, err := NewStorageService(ctx)
	if err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}

	encryptKey := GetEncryptKey(32)

	switch *cmd {
	case "upload":
		size, err := s.Upload(ctx, encryptKey, *bucket, *object, []byte("Hello World"))
		if err != nil {
			log.Fatalf("%+v", err)
			os.Exit(1)
		}
		log.Printf("done.upload size = %d", size)
	case "download":
		file, err := s.Download(ctx, encryptKey, *bucket, *object)
		if err != nil {
			log.Fatalf("%+v", err)
			os.Exit(1)
		}
		log.Printf("done.file = %s", string(file))
	case "printkey":
		encryptKey := GetEncryptKey(32)
		log.Printf(base64.URLEncoding.EncodeToString(encryptKey))
	default:
		log.Println("The cmd flog is upload or download")
	}
}

// GetEncryptKey is EncryptKeyを返す
func GetEncryptKey(n int) []byte {
	b := make([]byte, n)
	// TODO 実際にはここでKEYを設定する
	return b
}
