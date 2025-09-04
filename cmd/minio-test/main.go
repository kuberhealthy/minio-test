package main

import (
	"io/ioutil"
	"math/rand"
	"time"

	checkclient "github.com/kuberhealthy/kuberhealthy/v3/pkg/checkclient"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"context"
	"log"
	"os"
)

var minioEndpoint string
var accessKey string
var secretKey string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	minioEndpoint = os.Getenv("MINIO_ENDPOINT")
	accessKey = os.Getenv("ACCESS_KEY")
	secretKey = os.Getenv("SECRET_KEY")

	if len(minioEndpoint) < 1 || len(accessKey) < 1 || len(secretKey) < 1 {
		log.Fatalln("MINIO_ENDPOINT, ACCESS_KEY, and SECRET_KEY are required.")
	}
}

func randSeq(n int) []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return []byte(string(b))
}

func main() {
	ctx := context.Background()

	// Initialize minio client object.
	minioClient, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
	})

	if err != nil {
		log.Printf("Reporting failure: Unable to setup minio client: %s", err)
		err = checkclient.ReportFailure([]string{err.Error()})
		if err != nil {
			log.Printf("Unable to report failure: %s", err)
		}
		os.Exit(0)
	}

	bucketName := "khcheck-test"
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		log.Printf("Unable to make a bucket in minio: %s", err)
		err = checkclient.ReportFailure([]string{err.Error()})
		if err != nil {
			log.Printf("Unable to report failure: %s", err)
		}
		os.Exit(0)
	}

	log.Printf("Successfully created %s\n", bucketName)

	// create a file
	contents := randSeq(1000)
	err = ioutil.WriteFile("/app/test", contents, 0755)
	if err != nil {
		log.Printf("Unable to write a new file: %s", err)
		os.Exit(0)
	}

	// put the file object into our new bucket
	if _, err := minioClient.FPutObject(ctx, bucketName, "test", "/app/test", minio.PutObjectOptions{}); err != nil {
		log.Printf("Unable to put file object into minio: %s", err)
		err = checkclient.ReportFailure([]string{err.Error()})
		if err != nil {
			log.Printf("Unable to report failure: %s", err)
		}
		os.Exit(0)
	}

	// delete file
	err = os.Remove("/app/test")
	if err != nil {
		log.Printf("Unable to remove file: %s", err)
	}

	// delete object
	err = minioClient.RemoveObject(ctx, bucketName, "test", minio.RemoveObjectOptions{})
	if err != nil {
		log.Printf("Unable to remove object we wrote: %s", err)
		err = checkclient.ReportFailure([]string{err.Error()})
		if err != nil {
			log.Printf("Unable to report failure: %s", err)
		}
		os.Exit(0)
	}

	// delete bucket
	err = minioClient.RemoveBucket(ctx, bucketName)
	if err != nil {
		log.Printf("Unable to remove bucket we created: %s", err)
		err = checkclient.ReportFailure([]string{err.Error()})
		if err != nil {
			log.Printf("Unable to report failure: %s", err)
		}
		os.Exit(0)
	}

	// report success
	err = checkclient.ReportSuccess()
	if err != nil {
		log.Printf("Unable to report success to kuberhealthy: %s", err)
	}

	log.Println("Successfully reported to Kuberhealthy")
}
