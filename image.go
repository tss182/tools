package tools

import (
	"bytes"
	"encoding/base64"
	"golang.org/x/image/webp"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func CreateFileFromBase64(base64file, filename string) (string, error) {
	if base64file == "" {
		return "", nil
	}
	splitFile := strings.Split(base64file, ",")
	typeFile := strings.TrimPrefix(splitFile[0], "data:")
	typeFile = strings.TrimSuffix(typeFile, ";base64")
	base64File := splitFile[1]
	dec, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		return "", err
	}
	r := bytes.NewReader(dec)
	var file *os.File
	switch typeFile {
	case "image/png", "image/PNG":
		img, err := png.Decode(r)
		if err != nil {
			return "", err
		}
		filename += ".png"
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", err
		}
		_ = png.Encode(file, img)
	case "image/jpeg", "image/jpg":
		img, err := jpeg.Decode(r)
		if err != nil {
			return "", err
		}
		filename += ".jpg"
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", err
		}
		_ = jpeg.Encode(file, img, nil)
	case "image/webp":
		img, err := webp.Decode(r)
		if err != nil {
			return "", err
		}
		filename += ".jpg"
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", err
		}
		_ = jpeg.Encode(file, img, nil)
	}
	_ = file.Close()
	file, err = os.Open(filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}
