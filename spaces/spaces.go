package spaces

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/h2non/filetype"
)

type Manager struct {
	*s3.S3
	spaceName string
}

func New() *Manager {
	key := os.Getenv("SPACES_KEY")
	secret := os.Getenv("SPACES_SECRET")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String("https://sfo3.digitaloceanspaces.com"),
		Region:      aws.String("us-east-1"),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatalln(err)
	}
	s3Client := s3.New(newSession)

	return &Manager{
		S3:        s3Client,
		spaceName: "akita",
	}
}

// func NewDownloader() {
// 	key := os.Getenv("SPACES_KEY")
// 	secret := os.Getenv("SPACES_SECRET")

// 	s3Config := &aws.Config{
// 		Credentials: credentials.NewStaticCredentials(key, secret, ""),
// 		Endpoint:    aws.String("https://sfo3.digitaloceanspaces.com"),
// 		Region:      aws.String("us-east-1"),
// 	}

// 	newSession, err := session.NewSession(s3Config)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// downloader := s3.NewDownloader(newSession)

// 	s3Client := s3.New(newSession)

// 	return &Manager{
// 		S3:        s3Client,
// 		spaceName: "akita",
// 	}
// }

func (fm *Manager) AddFile(name string, file []byte, public bool) (string, error) {

	fType, err := filetype.Get(file)
	if err != nil {
		return "", err
	}

	var mime string

	if fType == filetype.Unknown {
		contentType := http.DetectContentType(file)
		if !strings.Contains(contentType, "text/plain") {
			return "", err
		}
		mime = contentType
	} else {
		mime = fType.MIME.Value
	}

	var acl string
	if public {
		acl = "public-read"
	} else {
		acl = "private"
	}

	object := s3.PutObjectInput{
		Bucket:      &fm.spaceName,
		Key:         &name,
		Body:        bytes.NewReader(file),
		ACL:         &acl,
		ContentType: &mime,
	}

	_, err = fm.PutObject(&object)
	if err != nil {
		return "", err
	}

	return name, nil
}

// AddStream uploads body without holding it in memory: the multipart
// uploader buffers at most Concurrency parts of PartSize bytes. A read error
// from body (for example a size cap) aborts the upload.
func (fm *Manager) AddStream(ctx context.Context, name string, body io.Reader, contentType string, public bool) (string, error) {
	acl := "private"
	if public {
		acl = "public-read"
	}
	uploader := s3manager.NewUploaderWithClient(fm.S3, func(u *s3manager.Uploader) {
		u.PartSize = 8 << 20
		u.Concurrency = 2
	})
	_, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket:      &fm.spaceName,
		Key:         &name,
		Body:        body,
		ACL:         &acl,
		ContentType: &contentType,
	})
	if err != nil {
		return "", err
	}
	return name, nil
}

func (fm *Manager) RemoveFile(name, ext string) error {
	input := &s3.DeleteObjectInput{
		Bucket: &fm.spaceName,
		Key:    &name,
	}

	_, err := fm.DeleteObject(input)
	if err != nil {
		return err
	}

	return nil
}
