package s3_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/external/aws/session"
	"gitlab.com/trivery-id/skadi/external/storage"
	. "gitlab.com/trivery-id/skadi/external/storage/s3"
)

var (
	testBucketName = "trivery-dev-skadi-public"
	testUploader   *Uploader
)

func TestMain(m *testing.M) {
	sess, err := session.GetDefaultSession()
	if err != nil {
		log.Fatal(err)
	}

	uploader, err := NewUploader(testBucketName, sess)
	if err != nil {
		log.Fatal(err)
	}

	testUploader = uploader
	os.Exit(m.Run())
}

func TestUpload(t *testing.T) {
	textFile, err := os.Open("assets/test_txt_01.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer textFile.Close()

	imageFile, err := os.Open("assets/test_image_01.png")
	if err != nil {
		t.Fatal(err)
	}
	defer imageFile.Close()

	t.Run("ok - upload txt", func(t *testing.T) {
		got, err := testUploader.Upload(context.Background(), storage.UploadInput{
			Key:             "test/test_txt_01.txt",
			File:            textFile,
			AllowPublicRead: true,
		})

		assert.NotEmpty(t, got)
		assert.Equal(t, "https://trivery-dev-skadi-public.s3.ap-southeast-1.amazonaws.com/test/test_txt_01.txt", got)
		assert.Nil(t, err)
	})

	t.Run("ok - upload image", func(t *testing.T) {
		got, err := testUploader.Upload(context.Background(), storage.UploadInput{
			Key:             "test/test_image_01.png",
			File:            imageFile,
			AllowPublicRead: true,
		})

		assert.NotEmpty(t, got)
		assert.Equal(t, "https://trivery-dev-skadi-public.s3.ap-southeast-1.amazonaws.com/test/test_image_01.png", got)
		assert.Nil(t, err)
	})
}
