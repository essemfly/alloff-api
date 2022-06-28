package productinfo

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/aws"
	"go.uber.org/zap"
)

func CacheProductImages(pdInfoDao *domain.ProductMetaInfoDAO) *domain.ProductMetaInfoDAO {
	if pdInfoDao.IsImageCached {
		return pdInfoDao
	}

	newImageUrls, err := cacheImages(pdInfoDao.ID.Hex(), pdInfoDao.Images)
	if err != nil {
		return pdInfoDao
	}
	newDescImageUrls, err := cacheImages(pdInfoDao.ID.Hex()+"-desc", pdInfoDao.SalesInstruction.Description.Images)
	if err != nil {
		return pdInfoDao
	}

	pdInfoDao.CachedImages = pdInfoDao.Images
	pdInfoDao.Images = newImageUrls
	pdInfoDao.SalesInstruction.Description.Images = newDescImageUrls
	pdInfoDao.IsImageCached = true

	pdInfoDao, _ = Update(pdInfoDao)
	return pdInfoDao
}

func cacheImages(pdInfoID string, images []string) ([]string, error) {
	newImageUrls := []string{}
	for idx, imgURL := range images {
		imgResp, err := http.Get(imgURL)
		if err != nil {
			config.Logger.Error("failed to get image from url: "+imgURL, zap.Error(err))
			return nil, err
		}
		defer imgResp.Body.Close()

		extension, err := getFileExtensionFromUrl(imgURL)
		if err != nil {
			config.Logger.Error("failed to get extension from url: "+imgURL, zap.Error(err))
		}

		filename := pdInfoID + "-" + strconv.Itoa(idx)
		filekey, err := aws.UploadImage(imgResp.Body, filename, extension)
		if err != nil {
			config.Logger.Error("failed to upload image for pdinfo ID: "+pdInfoID, zap.Error(err))
		}
		newImageUrls = append(newImageUrls, filekey)
	}

	return newImageUrls, nil
}

func getFileExtensionFromUrl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	pos := strings.LastIndex(u.Path, ".")
	if pos == -1 {
		return "", errors.New("couldn't find a period to indicate a file extension")
	}
	return u.Path[pos+1 : len(u.Path)], nil
}
