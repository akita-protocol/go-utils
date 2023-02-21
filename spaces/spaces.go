package spaces

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
		Body:        strings.NewReader(string(file)),
		ACL:         &acl,
		ContentType: &mime,
	}

	_, err = fm.PutObject(&object)
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
