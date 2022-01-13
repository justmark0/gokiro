package app

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func calculateMD5FileName(s string) string {
	name := strings.TrimSuffix(s, filepath.Ext(s))
	extension := filepath.Ext(s)
	data := md5.Sum([]byte(name))
	return hex.EncodeToString(data[:]) + extension
}

// SaveFileAndReturnPath saves file and creates md5 of this name to avoid file rewrite
func SaveFileAndReturnPath(file multipart.File, handler *multipart.FileHeader) (string, error) {
	path := "./uploads/" + calculateMD5FileName(handler.Filename)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	log.Default()
	if err != nil {
		return "", err
	}
	_, _ = io.Copy(f, file)
	return path, nil
}
