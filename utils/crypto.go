package utils

import (
	"encoding/base64"
	"time"
)

func EncodeBase64(byte []byte) string {
	return base64.StdEncoding.EncodeToString(byte)
}

func ParseTime(timeStr []uint8) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", string(timeStr))
}
