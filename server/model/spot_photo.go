package model

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
)

type SpotPhotoModelImpl struct {
}

func NewSpotPhotoModel() SpotPhotoModel {
	return &SpotPhotoModelImpl{}
}

func (model *SpotPhotoModelImpl) Add(image *multipart.FileHeader) (fileName string, err error) {
	// 画像のの読み込み
	srcFileName := image.Filename
	srcFileExt := path.Ext(srcFileName)
	srcFile, err := image.Open()
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	// 書き込み先ディレクトリがあるか確認し、なければ作る
	dstFileName, dstDir, dstFilePath := model.generateFilePathSet(srcFileExt)
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		err = os.MkdirAll(dstDir, 0777)
		if err != nil {
			return "", nil
		}
	}

	// 書き込む画像をとりあえず作る
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	// 書き込む
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return "", err
	}

	return dstFileName, nil
}

func (model *SpotPhotoModelImpl) generateFilePathSet(ext string) (fileName, pathDir, filePath string) {
	uuidStr := uuid.New().String()
	fileName = fmt.Sprintf("%s%s", uuidStr, ext)
	pathDir = "./assets/spots/"
	filePath = path.Join("./assets/spots/", fileName)
	log.Printf("%+v\n", filePath)

	return fileName, pathDir, filePath
}
