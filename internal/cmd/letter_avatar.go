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
}

var names = []string{
	"Alice",
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

func drawAvatar() {
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
func runLetterAvatar(_ *cli.Context) error {
	fmt.Println("run Command letter_avatar")

	drawAvatar()
	os.Exit(1)
	return nil
}
