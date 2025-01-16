package main

import (
    "bytes"
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func uploadToS3(imageBytes []byte, bucket, key string) error {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    })
    if err != nil {
        return fmt.Errorf("failed to create session: %v", err)
    }

    svc := s3.New(sess)

    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   bytes.NewReader(imageBytes),
        ACL:    aws.String("public-read"), // Optional, adjust according to your needs
    })
    if err != nil {
        return fmt.Errorf("failed to upload to S3: %v", err)
    }

    return nil
}

func fetchAndUploadImage(imageUrl, bucket, key string) error {
    // Fetch the image
    resp, err := http.Get(imageUrl)
    if err != nil {
        return fmt.Errorf("failed to fetch image: %v", err)
    }
    defer resp.Body.Close()

    // Read the image bytes
    imageBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("failed to read image: %v", err)
    }

    // Upload to S3
    return uploadToS3(imageBytes, bucket, key)
}

func main() {
    imageUrl := "https://example.com/image.jpg" // Replace with image URL from Mediastack
    bucket := "your-s3-bucket-name"
    key := "images/fetched_image.jpg"

    err := fetchAndUploadImage(imageUrl, bucket, key)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Image uploaded successfully!")
    }
}
