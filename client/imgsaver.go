package client

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	CoverAssetDir = "www/img/covers/"
	NoCoverImg    = "default.jpg"
	BaseCoverUrl  = "/img/covers/"
)

func SaveCover(isbn, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	fn := fmt.Sprintf("%s%s", generateFileName(isbn), filepath.Ext(url))
	fp := fmt.Sprintf("%s%s", CoverAssetDir, fn)
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	defer f.Close()

	imgsrc := fmt.Sprintf("%s%s", BaseCoverUrl, fn)

	return imgsrc, nil
}

func generateFileName(isbn string) string {
	str := fmt.Sprintf("%s%d", isbn, time.Now().Unix())
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
