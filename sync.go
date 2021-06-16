package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	serverDirectory = "MinecraftServer"
	s3Bucket        = "pit-enigmatica-expert-server"
)

// upload uploads the minecraft server to S3.
func upload(client *s3.Client) {
	uploader := manager.NewUploader(client)
	_ = filepath.WalkDir(serverDirectory, func(path string, d fs.DirEntry, err error) error {
		if path != serverDirectory && !d.IsDir() {
			if _, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
				Bucket: aws.String(s3Bucket),
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

// download downloads the server from s3.
func download(client *s3.Client) {
	downloader := manager.NewDownloader(client)
	_ = filepath.WalkDir(serverDirectory, func(path string, d fs.DirEntry, err error) error {
		if path != serverDirectory && !d.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			if _, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
				Bucket: aws.String(s3Bucket),
				Key:    aws.String(path),
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	//cfg, err := config.LoadDefaultConfig(context.TODO())
	//if err != nil {
	//	panic(err)
	//}
	//
	//client := s3.NewFromConfig(cfg, func(o *s3.Options) {
	//	o.Region = "us-west-2"
	//	o.UseAccelerate = true
	//})

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Improper directive specified. Directives: Upload, Download")
		os.Exit(1)
	}

	switch strings.ToLower(args[0]) {
	case "upload":
		fmt.Println("Uploading")
	case "downloading":
		fmt.Println("Downloading")
	}
}
