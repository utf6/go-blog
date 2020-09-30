package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/utf6/go-blog/pkg/file"
	"github.com/utf6/go-blog/pkg/setting"
	"github.com/utf6/go-blog/pkg/util"
	"image/jpeg"
)

type QrCode struct {
	URL 	string
	Width 	int
	Height 	int
	Ext 	string
	Level 	qr.ErrorCorrectionLevel
	Mode 	qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

/**
产生二维码
 */
func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL: url,
		Width: width,
		Height: height,
		Level: level,
		Mode: mode,
		Ext: EXT_JPG,
	}
}

/**
获取二维码路径
 */
func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath
}

/**
获取二维码完整路径
 */
func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.QrCodeSavePath
}

/**
获取二维码路由
 */
func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}

/**
获取二维码文件名
 */
func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

/**
获取二维码文件后缀名
 */
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

/**
检验解码二维码
 */
func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()

	if file.CheckNotExist(src) == true {
		return false
	}
	
	return true
}

/**
加密二维码
 */
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name

	if file.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", nil
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", nil
		}
	}
	return name, path, nil
}