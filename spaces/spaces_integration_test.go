//go:build integration

package spaces

// Runs AddStream against an S3-compatible server:
//
//	docker run -d -p 127.0.0.1:39000:9000 -e MINIO_ROOT_USER=test -e MINIO_ROOT_PASSWORD=testtest minio/minio server /data
//	SPACES_TEST_ENDPOINT=http://127.0.0.1:39000 go test -tags integration ./spaces/

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type failAfter struct {
	r io.Reader
	n int64
}

func (f *failAfter) Read(p []byte) (int, error) {
	if f.n <= 0 {
		return 0, errors.New("too large")
	}
	if int64(len(p)) > f.n {
		p = p[:f.n]
	}
	n, err := f.r.Read(p)
	f.n -= int64(n)
	return n, err
}

func TestAddStream(t *testing.T) {
	endpoint := os.Getenv("SPACES_TEST_ENDPOINT")
	if endpoint == "" {
		t.Skip("SPACES_TEST_ENDPOINT not set")
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("test", "testtest", ""), Endpoint: aws.String(endpoint),
		Region: aws.String("us-east-1"), S3ForcePathStyle: aws.Bool(true),
	}))
	fm := &Manager{S3: s3.New(sess), spaceName: "akita"}
	fm.CreateBucket(&s3.CreateBucketInput{Bucket: aws.String("akita")})

	data := bytes.Repeat([]byte("0123456789abcdef"), 20<<20/16) // 20 MB, three parts
	if _, err := fm.AddStream(context.Background(), "ipfs/video", bytes.NewReader(data), "video/mp4", true); err != nil {
		t.Fatal(err)
	}
	out, err := fm.GetObject(&s3.GetObjectInput{Bucket: aws.String("akita"), Key: aws.String("ipfs/video")})
	if err != nil {
		t.Fatal(err)
	}
	got, _ := io.ReadAll(out.Body)
	if !bytes.Equal(got, data) || aws.StringValue(out.ContentType) != "video/mp4" {
		t.Fatalf("stored %d bytes (%s)", len(got), aws.StringValue(out.ContentType))
	}

	if _, err := fm.AddStream(context.Background(), "ipfs/broken", &failAfter{r: bytes.NewReader(data), n: 12 << 20}, "video/mp4", true); err == nil {
		t.Fatal("a failing reader must abort the upload")
	}
	if _, err := fm.HeadObject(&s3.HeadObjectInput{Bucket: aws.String("akita"), Key: aws.String("ipfs/broken")}); err == nil {
		t.Fatal("an aborted upload must not leave an object")
	}
}
