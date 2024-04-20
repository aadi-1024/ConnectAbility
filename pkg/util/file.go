package util

import (
	"bufio"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

// CreateFile creates file in static from multipart form and returns the filename
func CreateFile(mFile *multipart.FileHeader) (string, error) {
	file, err := mFile.Open()
	if err != nil {
		return "", err
	}
	filename := strconv.Itoa(rand.Int()%1000) + strconv.Itoa(int(time.Now().UnixMilli()))
	f, err := os.Create("static/" + filename)
	if err != nil {
		return "", err
	}

	_, err = bufio.NewReader(file).WriteTo(bufio.NewWriter(f))
	return filename, err
}
