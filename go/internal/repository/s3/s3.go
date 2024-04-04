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
	"sync"
)

// todo: understand the point:
// todo: creating folders as a structure (possible emulating folders through obj naming)

type FileWorker interface {
	SaveFile(file *os.File) error
	SaveFiles([]*os.File) error
	DeleteFile(fileName string) error
	GetFolderFiles(folderName string) ([]string, error)
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

func (s S3Repo) SaveFiles(files []*os.File) error {
	var wg sync.WaitGroup
	wg.Add(len(files))
	errorsCh := make(chan error, len(files))

	for _, file := range files {
		go func(file *os.File) {
			defer wg.Done()
			if err := s.SaveFile(file); err != nil {
				errorsCh <- err
				return
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(errorsCh)
	}()

	for err := range errorsCh {
		// todo: figure out how to properly handle multiply errs.
		// todo: look at transaction account in s3.
		if err != nil {
			return err
		}
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

func (s S3Repo) GetFolderFiles(folderName string) ([]string, error) {
	res, err := s.client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(s.cfgService.GetRepoConfig().Minio.Bucket),
		Prefix: aws.String(folderName),
	})
	if err != nil {
		return nil, fmt.Errorf("s3: failed to get field from %s folder: %v", folderName, err)
	}

	fileList := make([]string, len(res.Contents))
	for i, content := range res.Contents {
		fileList[i] = *content.Key
	}

	return fileList, nil
}
