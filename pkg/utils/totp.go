package utils

import (
	"bytes"
	base "encoding/base64"
	"image/png"

	"github.com/pquerna/otp/totp"
)

type MFAModel struct {
}

func NewMFAModel() *MFAModel {
	return &MFAModel{}
}

func (m *MFAModel) GenerateTotp(userName string) (secret string, base64 string, err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Mute-tiktok",
		AccountName: userName,
	})
	if err != nil {
		return
	}

	img, err := key.Image(200, 200)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return
	}

	return key.Secret(), base.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func (m *MFAModel) VerifyOtp(token, secret string) bool {
	valid := totp.Validate(token, secret)
	return valid
}
