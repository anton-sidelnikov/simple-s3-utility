package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var outputPath = flag.String("path", "", "Set location of key file to save")

const apiEndpoint = "https://obs.eu-de.otc.t-systems.com"
const bucket = "obs-csm"

func main() {
	var item = "key/scn1_instance_rsa"
	file, err := os.Create(*outputPath)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	sess, _ := session.NewSession(&aws.Config{
		Endpoint:    aws.String(apiEndpoint),
		Credentials: credentials.NewEnvCredentials(),
	})

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
