package models

import (
	"encoding/base64"
	"github.com/satori/go.uuid"
	"math/rand"
)

func RandomByte() int {
	return (5 + rand.Int()%20) * 10
}

func RandomColor() Color {
	return Color{R: RandomByte(), G: RandomByte(), B: RandomByte()}
}

func Solid16x16gifDatauri(r, g, b uint8) string {
	bytes := make([]byte, 6)
	bytes[0] = 0
	bytes[1] = r
	bytes[2] = g
	bytes[3] = b
	bytes[4] = 0
	bytes[5] = 0
	encodeString := base64.StdEncoding.EncodeToString(bytes)
	return "data:image/gif;base64,R0lGODlhEAAQAIAA" + encodeString + "ACH5BAQAAAAALAAAAAAQABAAAAIOhI+py+0Po5y02ouzPgUAOw=="
}

func TinyId4() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()[0:4]
}

func TinyId6() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()[0:6]
}
