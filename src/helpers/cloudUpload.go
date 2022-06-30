package helpers

import (
	"context"
	"errors"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudUpload(src string) (string, error) {
	cloud := os.Getenv("CLOUD_NAME")
	keys := os.Getenv("CLOUD_KEY")
	secret := os.Getenv("CLOUD_SEC")
	ctx := context.Background()

	cld, err := cloudinary.NewFromParams(cloud, keys, secret)
	if err != nil {
		return "", errors.New("gorm failed to connect")
	}

	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{PublicID: "imageID"})
	if err != nil {
		return "", errors.New("gorm failed to connect")
	}

	return resp.SecureURL, nil
}
