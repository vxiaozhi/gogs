package staticsite

import (
	"encoding/json"
	//"text/template"
	"net/url"
	"strconv"

	log "unknwon.dev/clog/v2"
)

// var site_data_list string = `
// [
// 	{
// 		"category": "常用推荐",
// 		"iclass":"linecons-star",
// 		"data": [
// 			{
// 				"title": "Dribbble",
// 				"url": "https://dribbble.com/",
// 				"img": "/static/webstack/assets/images/logos/dribbble.png",
// 				"desc": "全球UI设计师作品分享平台。"
// 			},
// 			{
// 				"title": "UI中国",
// 				"url": "http://www.ui.cn/",
// 				"img": "/static/webstack/assets/images/logos/uicn.png",
// 				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
// 			}
// 		]
// 	},
// 	{
// 		"category": "社区资讯",
// 			"iclass":"linecons-doc",
// 		"data": [
// 			{
// 				"title": "雷锋网",
// 				"url": "https://www.leiphone.com/",
// 				"img": "/static/webstack/assets/images/logos/dribbble.png",
// 				"desc": "人工智能和智能硬件领域的互联网科技媒体"
// 			},
// 			{
// 				"title": "36kr",
// 				"url": "http://36kr.com/",
// 				"img": "/static/webstack/assets/images/logos/uicn.png",
// 				"desc": "创业资讯、科技新闻"
// 			}
// 		]
// 	}

// ]

// `

var site_data_list string = `
[
	{
		"category": "常用推荐",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐2",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐3",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐4",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐5",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐6",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐7",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐8",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐9",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐21",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐22",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "常用推荐23",  
		"iclass":"linecons-star",
		"data": [
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{ 
				"title": "Dribbble",
				"url": "https://dribbble.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "全球UI设计师作品分享平台。"
			},
			{
				"title": "UI中国",
				"url": "http://www.ui.cn/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "图形交互与界面设计交流、作品展示、学习平台。"
			}
		]
	},
	{
		"category": "社区资讯",  
			"iclass":"linecons-doc",
		"data": [
			{ 
				"title": "雷锋网",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网2",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网3",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网4",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网5",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网6",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网7",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网8",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网9",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{ 
				"title": "雷锋网10",
				"url": "https://www.leiphone.com/",
				"img": "/static/webstack/assets/images/logos/dribbble.png",
				"desc": "人工智能和智能硬件领域的互联网科技媒体"
			},
			{
				"title": "36kr",
				"url": "http://36kr.com/",
				"img": "/static/webstack/assets/images/logos/uicn.png",
				"desc": "创业资讯、科技新闻"
			}
		]
	}
	
]

`

type UrlItem struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Img   string `json:"img"`
	Desc  string `json:"desc"`
}

type SiteDataRaw struct {
	Category string    `json:"category"`
	IClass   string    `json:"iclass"`
	Data     []UrlItem `json:"data"`
}

const (
	MaxUrlItemsPerRow int = 4
)

var (
	// 随机定义一个整数值
	CateLinkInt int = 10011112301
)

type SiteDataRow struct {
	UrlItems []UrlItem
}
type SiteDataTmpl struct {
	Category string
	//CategoryLink  template.URL
	CategoryLink  string
	CategoryLink2 *url.URL
	IClass        string
	Rows          []SiteDataRow
}

func GetSiteDataList() ([]SiteDataTmpl, error) {
	var raw_data_list []SiteDataRaw
	err := json.Unmarshal([]byte(site_data_list), &raw_data_list)
	if err != nil {
		log.Error("Error: %v", err)
		return nil, err
	}
	//log.Info("raw_data_list: %v", raw_data_list)

	var tmpl_data_list []SiteDataTmpl
	for _, v := range raw_data_list {
		//cate_link, _ := url.Parse(v.Category)
		tmpl_data := SiteDataTmpl{
			Category: v.Category,
			//CategoryLink: template.URL(v.Category),
			CategoryLink: strconv.Itoa(CateLinkInt),
			IClass:       v.IClass,
			Rows:         []SiteDataRow{},
		}
		CateLinkInt += 1
		len := len(v.Data)
		if len <= MaxUrlItemsPerRow {
			tmpl_data.Rows = append(tmpl_data.Rows, SiteDataRow{
				UrlItems: v.Data,
			})

		} else {
			for i := 0; i < len; i += MaxUrlItemsPerRow {
				end := i + MaxUrlItemsPerRow
				if end > len {
					end = len
				}
				tmpl_data.Rows = append(tmpl_data.Rows, SiteDataRow{
					UrlItems: v.Data[i:end],
				})
			}
		}

		// 修正img
		for i := range tmpl_data.Rows {
			for j := range tmpl_data.Rows[i].UrlItems {
				if tmpl_data.Rows[i].UrlItems[j].Img == "" {
					img_base64, err := DrawAvatarToPngBase64Smart(tmpl_data.Rows[i].UrlItems[j].Title)
					if err != nil {
						log.Error("Error: %v", err)
						continue
					}
					tmpl_data.Rows[i].UrlItems[j].Img = "data:image/png;base64," + img_base64
					log.Info("img_base64: %s", tmpl_data.Rows[i].UrlItems[j].Img)
				}
			}
		}

		tmpl_data_list = append(tmpl_data_list, tmpl_data)

	}
	log.Info("tmpl_data_list: %v", tmpl_data_list)
	return tmpl_data_list, nil
}
