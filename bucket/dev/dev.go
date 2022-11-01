package dev

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

type Stream struct {
	Connection io.Closer
	Reader     io.Reader
	Buffer     bytes.Buffer
	Headers    []string
	BytesRead  uint64
}

func Upload(ctx context.Context, cfg aws.Config, bucket, key string, r io.Reader) (*manager.UploadOutput, error) {
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    aws.String(key),
		Body:   r,
	})

	return result, err
}
func (s *Stream) Read(b []byte) (n int, err error) {
	n, err = s.Buffer.Read(b)
	s.BytesRead += uint64(n)
	if err == io.EOF {
		err = s.LoadNextLine()
		if err != nil {
			return n, err
		}
		return n, nil
	}
	return n, err
}

func (s *Stream) LoadNextLine() error {
	line, err := s.Buffer.ReadBytes('\n')
	if err != nil {
		if err == io.EOF {
			if s.Connection != nil {
				_ = s.Connection.Close()
			}
		}
		return err
	}

	s.Buffer.Write([]byte(fmt.Sprintf("%s", line)))
	return nil
}
