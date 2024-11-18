package staticsite

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"
	"unicode"
	"unicode/utf8"

	"github.com/disintegration/letteravatar"
	"github.com/mozillazg/go-pinyin"
)

func DrawAvatarToPngFile(name string, isHan bool, filename string) error {
	if isHan {
		a := pinyin.NewArgs()
		name_pinyin := pinyin.Pinyin(name, a)
		fmt.Println(name_pinyin)
		name = name_pinyin[0][0]
	}
	firstLetter, _ := utf8.DecodeRuneInString(name)

	img, err := letteravatar.Draw(120, firstLetter, nil)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = png.Encode(file, img)
	return err

}

func DrawAvatarToPngBase64(name string, isHan bool) (string, error) {
	if isHan {
		a := pinyin.NewArgs()
		name_pinyin := pinyin.Pinyin(name, a)
		fmt.Println(name_pinyin)
		name = name_pinyin[0][0]
	}
	firstLetter, _ := utf8.DecodeRuneInString(name)

	img, err := letteravatar.Draw(120, firstLetter, nil)
	if err != nil {
		return "", err
	}

	// 创建一个 bytes.Buffer
	buf := new(bytes.Buffer)

	// 将图片编码成 PNG 格式，并以 base64 形式写入到 buffer 中
	err = png.Encode(buf, img)
	if err != nil {
		return "", err
	}

	// 将 buffer 中的内容转换为 base64 字符串
	encodedString := base64.StdEncoding.EncodeToString(buf.Bytes())

	return encodedString, nil
}

func DrawAvatarToPngBase64Smart(name string) (string, error) {
	// 先判断 name 是否是英文字母开头
	if len(name) > 0 && unicode.IsLetter(rune(name[0])) {
		return DrawAvatarToPngBase64(name, false)
	}
	a := pinyin.NewArgs()
	name_pinyin := pinyin.Pinyin(name, a)
	fmt.Println(name_pinyin)
	if len(name_pinyin) > 0 && len(name_pinyin[0]) > 0 {
		return DrawAvatarToPngBase64(name, true)
	}
	return DrawAvatarToPngBase64("A", false)
}

func GetFirstLetter(s string) string {
	if len(s) == 0 {
		return "A"
	}
	firstChar := []rune(s)[0]

	if unicode.IsUpper(firstChar) {
		return string(firstChar)
	} else if unicode.IsLower(firstChar) {
		return string(unicode.ToUpper(firstChar))
	} else {
		return "A"
	}
}
