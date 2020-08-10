package models

import (
	"encoding/base64"
	"github.com/satori/go.uuid"
)

func Solid16x16gifDatauri(color *Color) string {
	bytes := make([]byte, 6)
	bytes[0] = 0
	bytes[1] = uint8(color.R)
	bytes[2] = uint8(color.G)
	bytes[3] = uint8(color.B)
	bytes[4] = 0
	bytes[5] = 0
	encodeString := base64.StdEncoding.EncodeToString(bytes)
	return "data:image/gif;base64,R0lGODlhEAAQAIAA" + encodeString + "ACH5BAQAAAAALAAAAAAQABAAAAIOhI+py+0Po5y02ouzPgUAOw=="
}

func TinyId(len int) string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return u1.String()[0:len]
}
