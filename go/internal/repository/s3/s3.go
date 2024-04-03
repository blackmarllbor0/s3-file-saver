package s3

import (
	"app/internal/cfg"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type a interface {
	SaveFile(file *os.File) error
	DeleteFile(fileName string) error
}

type S3Repo struct {
	uploader *s3manager.Uploader
	client   *s3.S3

	cfgService cfg.ConfigService
}

func NewS3Session(cfgService cfg.ConfigService) (*S3Repo, error) {
	s3Cfg := cfgService.GetRepoConfig().Minio

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(s3Cfg.Region),
		Endpoint:         aws.String(fmt.Sprintf("%s:%d", s3Cfg.Host, s3Cfg.Port)),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(s3Cfg.User, s3Cfg.Pwd, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("s3: failed to create session: %v", err)
	}

	return &S3Repo{
		uploader:   s3manager.NewUploader(sess),
		cfgService: cfgService,
		client:     s3.New(sess),
	}, nil
}

func (s S3Repo) SaveFile(file *os.File) error {
	if _, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.cfgService.GetRepoConfig().Minio.Bucket),
		Key:    aws.String(file.Name()),
		Body:   file,
	}); err != nil {
		return fmt.Errorf("s3: failed to upload file %v: %v", file.Name(), err)
	}

	return nil
}

func (s S3Repo) DeleteFile(fileName string) error {
	if _, err := s.client.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(s.cfgService.GetRepoConfig().Minio.Bucket),
		Delete: &s3.Delete{
			Objects: []*s3.ObjectIdentifier{{
				Key: aws.String(fileName),
			}},
		},
	}); err != nil {
		return fmt.Errorf("s3: failed to delete file by key: %s: %v", fileName, err)
	}

	return nil
}
