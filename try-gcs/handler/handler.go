package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	imageupload "github.com/olahol/go-imageupload"
)

var (
	projectID  = os.Getenv("PROJECT_ID")
	bucketName = os.Getenv("BUCKET_NAME")
)

// Handler - upload gcs bucket
func Handler(c *gin.Context) {
	id := c.PostForm("uuid")
	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	r := bytes.NewReader(thumb.Data)
	ctx := context.Background()
	name := fmt.Sprintf("%s.png", id)

	_, objAttrs, err := upload(ctx, r, projectID, bucketName, name, true)
	if err != nil {
		switch err {
		case storage.ErrBucketNotExist:
			log.Fatal("Please create the bucket first e.g. with `gsutil mb")
		default:
			log.Fatal(err)
		}
	}

	log.Printf("URL: %s", objectURL(objAttrs))
	log.Printf("Size: %d", objAttrs.Size)
}

func upload(ctx context.Context, r io.Reader, projectID, bucketName, name string, public bool) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}

	bh := client.Bucket(bucketName)

	obj := bh.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	if public {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, nil, err
		}
	}

	attrs, err := obj.Attrs(ctx)
	return obj, attrs, err
}

func objectURL(objAttrs *storage.ObjectAttrs) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", objAttrs.Bucket, objAttrs.Name)
}
