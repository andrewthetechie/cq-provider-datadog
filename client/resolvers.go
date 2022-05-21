package client

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func CreateMetaID(accountName string, id int64) string {
	hasher := sha1.New()
	hasher.Write([]byte(fmt.Sprintf("%s-%d", accountName, id)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func CreateMetaIDStr(accountName string, id string) string {
	hasher := sha1.New()
	hasher.Write([]byte(fmt.Sprintf("%s-%s", accountName, id)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
