package uploadbiz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"g07/common"
	"g07/component/uploadprovider"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type uploadBiz struct {
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{
		provider: provider,
	}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, errors.New("file is not image")
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
