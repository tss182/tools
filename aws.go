package tools

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type (
	AwsConfig struct {
		Key      string
		Secret   string
		Endpoint string
		Domain   string
		Region   string
		Bucket   string
		client   *s3.S3
	}
	AwsUploadReq struct {
		Bucket        string
		Folder        string
		Prefix        string
		Cache         bool
		CacheTime     int
		File          *os.File
		FileMultipart *multipart.FileHeader
	}
	AwsUploadResp struct {
		Url string `json:"url"`
	}
	AwsDeletedReq struct {
		Bucket string
		File   string
	}
)

func InitAws(cfg AwsConfig) (*AwsConfig, error) {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(cfg.Key, cfg.Secret, ""),
		Endpoint:    aws.String(cfg.Endpoint),
		Region:      aws.String(cfg.Region),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}
	cfg.client = s3.New(newSession)
	return &cfg, nil
}

func (s *AwsConfig) Upload(req AwsUploadReq) (*AwsUploadResp, error) {
	//get file
	fileInfo, _ := req.File.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	_, _ = req.File.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	//folder
	path := strings.TrimSpace(req.Folder)
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")
	filename := fileInfo.Name()
	if req.Prefix != "" {
		filename = req.Prefix + filename
	}
	path += "/" + filename
	expired := ""
	if req.Cache {
		expired = fmt.Sprintf("max-age=%s", req.CacheTime)
	}

	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(req.Bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
		ACL:           aws.String("public-read"),
		CacheControl:  aws.String(expired),
	})
	if err != nil {
		return nil, err
	}

	return &AwsUploadResp{Url: s.Domain + "/" + path}, nil
}

func (s *AwsConfig) UploadMultipart(req AwsUploadReq) (*AwsUploadResp, error) {
	//get file
	file, err := req.FileMultipart.Open()
	var size = req.FileMultipart.Size
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)

	//folder
	path := strings.TrimSpace(req.Folder)
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")
	filename := Slug(req.FileMultipart.Filename)
	if req.Prefix != "" {
		filename = req.Prefix + filename
	}
	path += "/" + filename
	expired := ""
	if req.Cache {
		expired = fmt.Sprintf("max-age=%s", req.CacheTime)
	}

	_, err = s.client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(req.Bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(req.FileMultipart.Header.Get("Content-Type")),
		ACL:           aws.String("public-read"),
		CacheControl:  aws.String(expired),
	})
	if err != nil {
		return nil, err
	}

	return &AwsUploadResp{Url: s.Domain + "/" + path}, nil
}

func (s *AwsConfig) Delete(req AwsDeletedReq) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(req.Bucket),
		Key:    aws.String(req.File),
	})
	if err != nil {
		return err
	}
	return nil
}
