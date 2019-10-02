package main

import (
	"flag"
	"os"
	"log"

	"github.com/minio/minio-go/v6"

)

var key = flag.String("k", "", "Set file key to get from s3")
var outputFile = flag.String("o", "output", "Set output file name")

const apiEndpoint = "obs.eu-de.otc.t-systems.com"
const bucket = "obs-csm"


func main() {
	flag.Parse()
	s3Client, err := minio.New(apiEndpoint, os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), false)
	if err != nil {
		log.Fatalln(err)
	}

	if err := s3Client.FGetObject(bucket, *key, *outputFile, minio.GetObjectOptions{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully saved ", *outputFile)
}
