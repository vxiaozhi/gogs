//go:build !cert

// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"gogs.io/gogs/internal/staticsite"
)

var LetterAvatar = cli.Command{
	Name:        "LetterAvatar",
	Usage:       "run command letter_avatar",
	Description: ``,
	Action:      runLetterAvatar,
	Flags: []cli.Flag{
		stringFlag("subcmd", "base64", "base64/names/upperletters/lowletters"),
	},
}

var names = []string{
	"Alice",
	"alice",
	"Bob",
	"Carol",
	"Dave",
	"Eve",
	"Frank",
	"Gloria",
	"Henry",
	"Isabella",
	"James",
	"Жозефина",
	"Ярослав",
	"中国",
	"推荐",
}

func drawNameAvatar() {
	for _, name := range names {
		filename := name + ".png"
		is_han := false
		if name == "中国" {
			is_han = true
		}
		err := staticsite.DrawAvatarToPngFile(name, is_han, filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func outBase64() {

	for _, name := range names {
		is_han := false
		if name == "中国" {
			is_han = true
		}
		out_base64, err := staticsite.DrawAvatarToPngBase64(name, is_han)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name, out_base64)
	}

}
func drawLettersAvatar(isUpper bool) {
	if isUpper {
		// 大写字母
		for i := 65; i <= 90; i++ {
			filename := string(i) + ".png"
			err := staticsite.DrawAvatarToPngFile(string(i), false, filename)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		// 小写字母
		for i := 97; i <= 122; i++ {
			filename := string(i) + ".png"
			err := staticsite.DrawAvatarToPngFile(string(i), false, filename)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func runLetterAvatar(ctx *cli.Context) error {
	fmt.Println("run Command letter_avatar")
	if ctx.String("subcmd") == "base64" {
		outBase64()
	} else if ctx.String("subcmd") == "names" {
		drawNameAvatar()
	} else if ctx.String("subcmd") == "upperletters" {
		drawLettersAvatar(true)
	} else if ctx.String("subcmd") == "lowletters" {
		drawLettersAvatar(false)
	}

	os.Exit(1)
	return nil
}
