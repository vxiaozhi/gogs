package staticsite

import (
	"encoding/json"
	"fmt"
	"strconv"

	//"text/template"

	log "unknwon.dev/clog/v2"
)

const (
	MaxUrlItemsPerRow int = 4
)

var (
	// 随机定义一个整数值
	CateLinkInt int = 10011112301

	CategoryIdStart uint32 = 10001
	ItemIdStart     uint32 = 20001
)

// NavContent 定义
type UrlItem struct {
	ItemId uint32 `json:"itemid"`
	Title  string `json:"title"`
	Url    string `json:"url"`
	Img    string `json:"img"`
	Desc   string `json:"desc"`
}

type NavCategoryItem struct {
	CategoryId   uint32    `json:"cateid"`
	CategoryName string    `json:"category"`
	IClass       string    `json:"iclass"`
	Data         []UrlItem `json:"data"`

	// 校验 CategoryId 是否唯一
	itemIdMap map[uint32]int
}

type NavContentStruct struct {
	Title           string            `json:"title"`
	Keywords        string            `json:"keywords"`
	Description     string            `json:"description"`
	NavCategoryList []NavCategoryItem `json:"navcategories"`

	// 校验 CategoryId 是否唯一
	categoryIdMap map[uint32]int
}

// SiteDataTmpl 定义

type SiteDataRow struct {
	UrlItems []UrlItem
}

type CategoryItem struct {
	Category     string `json:"category"`
	CategoryLink string
	IClass       string `json:"iclass"`
	RowsData     []SiteDataRow
}

type SiteDataTmpl struct {
	Title        string         `json:"title"`
	Keywords     string         `json:"keywords"`
	Description  string         `json:"description"`
	CategoryList []CategoryItem `json:"categories"`
}

func (s *SiteDataTmpl) GetCateNameList() []string {
	var cate_name_list []string
	for _, v := range s.CategoryList {
		cate_name_list = append(cate_name_list, v.Category)
	}
	return cate_name_list
}

// nav 函数定义
func NewNavContentStruct(nav_content string) (*NavContentStruct, error) {
	var nav_content_struct NavContentStruct
	err := json.Unmarshal([]byte(nav_content), &nav_content_struct)
	if err != nil {
		log.Error("Error: %v", err)
		return &NavContentStruct{}, err
	}

	log.Info("nav_content_struct: %v", nav_content_struct)
	return &nav_content_struct, nil
}

// 对 nav_content_struct 进行修正，填充ID值。
func (nav *NavContentStruct) FixAndToJson() (string, error) {
	// use categoryIdMap to fix CategoryId

	nav.categoryIdMap = make(map[uint32]int)

	// Step 1 只要存在 CategoryId = 0 的情况， 则全部重新生成新的 CategoryId、ItemId
	has_category_id_zero := false
	for _, v := range nav.NavCategoryList {
		if v.CategoryId == 0 {
			has_category_id_zero = true
		}
	}

	if has_category_id_zero {
		for i, v := range nav.NavCategoryList {

			nav.NavCategoryList[i].CategoryId = CategoryIdStart
			CategoryIdStart += 1

			nav.categoryIdMap[v.CategoryId] = 1

			for item_i, _ := range v.Data {

				nav.NavCategoryList[i].Data[item_i].ItemId = ItemIdStart
				ItemIdStart += 1
			}
		}
	}

	json_data, err := json.MarshalIndent(nav, "", "  ")
	if err != nil {
		log.Error("Error: %v", err)
		return "", err
	}
	return string(json_data), nil
}

// 将 NavContentStruct 转化为 站点需要的数据结构， 包括生成
// 1 生成对应的图标
// 2 将item 分割成多行，每行 4 个item。
func (nav *NavContentStruct) GetSiteDataTmpl(site_data_content string) (SiteDataTmpl, error) {

	site_data_tmpl := SiteDataTmpl{
		Title:        nav.Title,
		Keywords:     nav.Keywords,
		Description:  nav.Description,
		CategoryList: []CategoryItem{},
	}
	for cate_index, v := range nav.NavCategoryList {
		//v.RowsData = []SiteDataRow{}

		CateLinkInt += 1
		len := len(v.Data)
		cate_item := CategoryItem{
			Category:     v.CategoryName,
			CategoryLink: strconv.Itoa(CateLinkInt),
			IClass:       v.IClass,
			RowsData:     []SiteDataRow{},
		}
		if len <= MaxUrlItemsPerRow {
			//此处只能通过下标修改，否则不生效。参考：https://blog.csdn.net/qq_37102984/article/details/117850578
			cate_item.RowsData = append(cate_item.RowsData, SiteDataRow{
				UrlItems: v.Data,
			})

		} else {
			for i := 0; i < len; i += MaxUrlItemsPerRow {
				end := i + MaxUrlItemsPerRow
				if end > len {
					end = len
				}
				cate_item.RowsData = append(cate_item.RowsData, SiteDataRow{
					UrlItems: v.Data[i:end],
				})
				// site_data_tmpl.CategoryList[cate_index].RowsData = append(site_data_tmpl.CategoryList[cate_index].RowsData, SiteDataRow{
				// 	UrlItems: v.Data[i:end],
				// })
			}
		}
		site_data_tmpl.CategoryList = append(site_data_tmpl.CategoryList, cate_item)

		// 修正img
		for i, v_row := range site_data_tmpl.CategoryList[cate_index].RowsData {
			for j := range v_row.UrlItems {
				if v_row.UrlItems[j].Img == "" {
					img_letter := GetFirstLetter(v_row.UrlItems[j].Title)
					site_data_tmpl.CategoryList[cate_index].RowsData[i].UrlItems[j].Img = fmt.Sprintf("/static/webstack/assets/images/letters/%s.png", img_letter)
					//log.Info("img_base64: %s", v_row.UrlItems[j].Img)
				}
			}
		}

	}
	log.Info("site_data_tmpl: %v", site_data_tmpl)
	return site_data_tmpl, nil
}

func GetSiteDataTmpl(site_data_content string) (SiteDataTmpl, error) {
	nav_content, err := NewNavContentStruct(site_data_content)
	if err != nil {
		log.Error("Error: %v", err)
		return SiteDataTmpl{}, err
	}
	return nav_content.GetSiteDataTmpl(site_data_content)
}
