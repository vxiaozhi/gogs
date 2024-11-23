// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package app

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/macaron.v1"
)

// Go-bluemonday是一个在Golang中用于处理和清除HTML中的有害内容的库，它的设计灵感来自于OWASP Java HTML Sanitizer。OWASP（Open Web Application Security Project）是一个专注于网络安全的开源项目，
//它提供了多种工具和指南，其中包括防止跨站脚本（XSS）攻击的方法。XSS攻击是一种常见的网络攻击手段，攻击者通过注入恶意代码到网页中，来窃取用户的敏感信息或执行非法操作。

// bluemonday库的主要目标是帮助开发者安全地处理用户生成的内容（UGC），如博客评论、论坛帖子或社交媒体消息。在接收这些内容时，如果不进行适当的清理，可能会导致XSS漏洞，使得攻击者能够执行恶意脚本。
//bluemonday通过严格但可配置的策略来移除或转义可能危险的HTML标签、属性和值，从而防止XSS攻击。

// 使用bluemonday库，开发者可以创建自定义的策略来允许或禁止特定的HTML元素和属性。例如，你可以允许`<p>`和`<a>`标签，但不允许`<script>`和`<iframe>`等可能引入恶意代码的标签。
//同时，你还可以控制哪些属性是安全的，比如在`<a>`标签中只允许`href`属性，不允许`onclick`等可以执行JavaScript的属性。

func ipynbSanitizer() *bluemonday.Policy {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class", "data-prompt-number").OnElements("div")
	p.AllowAttrs("class").OnElements("img")
	p.AllowURLSchemes("data")
	return p
}

func SanitizeIpynb() macaron.Handler {
	p := ipynbSanitizer()

	return func(c *macaron.Context) {
		html, err := c.Req.Body().String()
		if err != nil {
			c.Error(http.StatusInternalServerError, "read body")
			return
		}

		c.PlainText(http.StatusOK, []byte(p.Sanitize(html)))
	}
}
