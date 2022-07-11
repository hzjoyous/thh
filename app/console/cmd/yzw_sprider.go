package cmd

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"golang.org/x/net/html"
	"strings"
	"sync"
	"thh/arms"
	"thh/arms/logger"
	"thh/arms/restyUtil"
)

func init() {
	appendCommand(YzwSpider)
}

var YzwSpider = &cobra.Command{
	Use:   "yzw_spider",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runYzwSpider,
	// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

type SpecialityInfo struct {
	Zz          string
	Wy          string
	B1          string
	B2          string
	Zsdw        string //招生单位：
	Kxfs        string //考试方式
	Zy          string // 专业：	(083500)软件工程
	Yx          string //院系所
	Fx          string //方向
	StudyFunc   string //方式
	Number      string //人数
	TeacherInfo string //指导老师：
}
type Speciality struct {
	Name     string
	Code     string
	IsDoctor string
	Info     []SpecialityInfo
}
type v struct {
	Name string
	City string

	SpecialityList []Speciality
}

var vList []v

type SimpleSType struct {
	Name  string
	Code  string
	Cckey string
}

func runYzwSpider(cmd *cobra.Command, args []string) {
	var sm sync.Map
	if v, ok := sm.Load("a"); ok {
		if t, tok := v.(int); tok {
			t = t + 1
		}
	}
	specialityList := []SimpleSType{
		{"软件工程", "083500", "10"},
		{"计算机科学与技术", "081200", "10"},
		{"计算机系统结构", "081201", "10"},
		{"计算机软件与理论", "081202", "10"},
		{"计算机应用技术", "081203", "10"},
		{"人工智能", "085410", "10"},
		{"大数据技术与工程", "085411", "10"},
		{"网络与信息安全", "085412", "10"},
		{"计算机技术", "085404", "10"},
		{"软件工程", "085405", "10"},
		{"电子信息", "085400", "20"},
	}
	vJson := arms.StorageGet("vlist.json")
	vList = arms.JsonDecode[[]v](vJson)
	client := newYZWClient()

	vMap := map[string]v{}

	if len(vList) == 0 {
		r, _ := client.getSpecialCategory()
		for _, speciality := range specialityList {
			r, _ = client.specialityDetail(speciality.Name, speciality.Code, speciality.Cckey)
			wr := restyUtil.GetCurlByR(*r)
			fmt.Println(wr.String())
			doc, _ := htmlquery.Parse(strings.NewReader(r.String()))
			list := htmlquery.Find(doc, "//div[@class=\"tab-container zyk-zyfb-tab\"]/div/div[@class=\"item-content active\"]/ul/li")
			for _, item := range list {
				vName := strings.TrimSpace(htmlquery.InnerText(item))
				if value, ok := vMap[vName]; ok {
					//value.SpecialityList = append(value.SpecialityList, Speciality{Name: speciality.Name, Code: speciality.Code})
					vMap[vName] = value
				} else {
					vMap[vName] = v{Name: vName, SpecialityList: []Speciality{
						//{Name: speciality.Name, Code: speciality.Code},
					}}
					fmt.Println(vMap)
				}

			}
		}
		for _, value := range vMap {
			vList = append(vList, value)
		}
		vList = append(vList, v{Name: "北京大学"})
	}
	vList = arms.ArrayMap(func(item v) v {
		city := ""
		item.SpecialityList = arms.ArrayMap(func(sItem SimpleSType) Speciality {

			isDoctor := ""

			r, _ := client.querySchAction(item.Name, sItem.Code[:4], sItem.Name)
			wr := restyUtil.GetCurlByR(*r)
			logger.Info(wr.String())
			doc, _ := htmlquery.Parse(strings.NewReader(r.String()))
			list := htmlquery.Find(doc, "//tbody/*")
			sTmp := Speciality{Name: sItem.Name, Code: sItem.Code}
			sTmp.Info = arms.ArrayMap(func(infoItem *html.Node) SpecialityInfo {

				hItem := htmlquery.FindOne(infoItem, "/td[8]/a/@href")
				vUrl := strings.TrimSpace(htmlquery.InnerText(hItem))
				fmt.Println(vUrl)

				r, _ := client.kskm(vUrl)
				wr := restyUtil.GetCurlByR(*r)
				logger.Info(wr.String())

				mItem, _ := htmlquery.Parse(strings.NewReader(r.String()))

				zsdw := XpathFirst(mItem, "//tbody[1]/tr[1]/td[2]")        // 招生单位
				fangshi := XpathFirst(mItem, "//tbody[1]/tr[1]/td[4]")     // 考试方式
				yx := XpathFirst(mItem, "//tbody[1]/tr[2]/td[2]")          //院系所
				zhuangye := XpathFirst(mItem, "//tbody[1]/tr[2]/td[4]")    //专业
				studyFunc := XpathFirst(mItem, "//tbody[1]/tr[3]/td[2]")   //学习方式
				fangxiang := XpathFirst(mItem, "//tbody[1]/tr[3]/td[4]")   //研究方向
				teacherInfo := XpathFirst(mItem, "//tbody[1]/tr[4]/td[2]") //学习方式
				number := XpathFirst(mItem, "//tbody[1]/tr[4]/td[4]")      // 拟招人数

				Zz := XpathFirst(mItem, "//tbody[@class=\"zsml-res-items\"]/tr/td[1]")
				Zz = strings.Split(Zz, "\n")[0]
				Wy := XpathFirst(mItem, "//tbody[@class=\"zsml-res-items\"]/tr/td[2]")
				Wy = strings.Split(Wy, "\n")[0]
				B1 := XpathFirst(mItem, "//tbody[@class=\"zsml-res-items\"]/tr/td[3]")
				B1 = strings.Split(B1, "\n")[0]
				B2 := XpathFirst(mItem, "//tbody[@class=\"zsml-res-items\"]/tr/td[4]")
				sTmp := SpecialityInfo{Zz: Zz, Wy: Wy, B1: B1, B2: B2,
					Zsdw:        zsdw,        //招生单位：
					Kxfs:        fangshi,     //考试方式
					Zy:          zhuangye,    // 专业：	(083500)软件工程
					Yx:          yx,          //院系所
					Fx:          fangxiang,   //方向
					StudyFunc:   studyFunc,   //方式
					Number:      number,      //人数
					TeacherInfo: teacherInfo, //指导老师：
				}
				logger.Info(arms.JsonEncode(sTmp))

				r, _ = client.queryAction(item.Name, sItem.Code[:4], sItem.Name)
				mItem, _ = htmlquery.Parse(strings.NewReader(r.String()))
				tmpNode := htmlquery.FindOne(mItem, "//tbody/tr/td[2]")
				if tmpNode != nil {
					if len(city) == 0 {
						city = strings.TrimSpace(htmlquery.InnerText(tmpNode)) //
					}
					tmpNode = htmlquery.FindOne(mItem, "//tbody/tr/td[5]/i") //博士点
					if tmpNode != nil {
						isDoctor = "有"
					}
				}

				return sTmp
			}, list)
			sTmp.IsDoctor = isDoctor
			return sTmp
		}, specialityList)
		item.City = city
		return item
	}, vList)

	vListString := arms.JsonEncode(vList)
	arms.StoragePut("vlist.json", vListString, false)

}

func XpathFirst(item *html.Node, expr string) string {
	tmpNode := htmlquery.FindOne(item, expr)
	if tmpNode == nil {
		return ""
	}
	return strings.TrimSpace(htmlquery.InnerText(tmpNode)) //院系所
}

type yzwClient struct {
	httpClient *resty.Client
}

func newYZWClient() yzwClient {
	client := resty.New()
	// Setting a Proxy URL and Port
	//client.SetProxy(config.GetString("t_Proxy"))
	client.SetBaseURL("https://yz.chsi.com.cn/")
	return yzwClient{client}
}

// 分类页面
func (itself *yzwClient) getSpecialCategory() (*resty.Response, error) {
	return itself.httpClient.R().
		SetFormData(map[string]string{
			"method": "subCategoryXk",
			"key":    "100812",
		}).Post("zyk/specialityCategory.do")
}

// 专业页面 可以查询这个专业有哪些学校
func (itself *yzwClient) specialityDetail(zymc, zydm, cckey string) (*resty.Response, error) {
	return itself.httpClient.R().
		SetQueryParams(map[string]string{
			"zymc":   zymc, // 专业名称
			"zydm":   zydm, // 专业代码
			"cckey":  cckey,
			"method": "distribution",
		}).Get("zyk/specialityDetail.do")
}

// 学校专业查询，可以学校的具体信息
func (itself *yzwClient) querySchAction(dwmc, yjxkdm, zymc string) (*resty.Response, error) {
	// 内部有一个数据可以继续 /zsml/kskm.jsp?id=1026921160083500021 查询考试范围
	//https://yz.chsi.com.cn/zsml/queryAction.do
	return itself.httpClient.R().
		SetQueryParams(map[string]string{
			"ssdm":   "",
			"dwmc":   dwmc, //华东师范大学
			"mldm":   "",
			"mlmc":   "",
			"yjxkdm": yjxkdm, //0835
			"zymc":   zymc,   //软件工程
			"xxfs":   "",
		}).Post("zsml/querySchAction.do")
}

//学校专业查询，可以学校的具体信息
func (itself *yzwClient) queryAction(dwmc, yjxkdm, zymc string) (*resty.Response, error) {
	// 内部有一个数据可以继续 /zsml/kskm.jsp?id=1026921160083500021 查询考试范围
	//https://yz.chsi.com.cn/zsml/queryAction.do
	return itself.httpClient.R().
		SetFormData(map[string]string{
			"ssdm":   "",
			"dwmc":   dwmc, //华东师范大学
			"mldm":   "",
			"mlmc":   "",
			"yjxkdm": yjxkdm, //0835
			"zymc":   zymc,   //软件工程
			"xxfs":   "",
		}).Post("zsml/queryAction.do")
}

func (itself *yzwClient) kskm(uri string) (*resty.Response, error) {
	// 内部有一个数据可以继续 /zsml/kskm.jsp?id=1026921160083500021 查询考试范围
	//https://yz.chsi.com.cn/zsml/queryAction.do
	return itself.httpClient.R().Get(uri)
}

//for key, vItem := range vList {
//
//	subList := []Speciality{}
//	for _, sItem := range specialityList {
//
//		r, _ := client.querySchAction(vItem.Name, sItem.Code[:4], sItem.Name)
//		wr := restyUtil.GetCurlByR(*r)
//		fmt.Println(wr.String())
//		doc, _ := htmlquery.Parse(strings.NewReader(r.String()))
//		list := htmlquery.Find(doc, "//tbody/*")
//		slist := []SpecialityInfo{}
//		for _, item := range list {
//
//			//fmt.Println(htmlquery.OutputHTML(item, true))
//			hItem := htmlquery.FindOne(item, "/td[8]/a/@href")
//			vUrl := strings.TrimSpace(htmlquery.InnerText(hItem))
//			fmt.Println(vUrl)
//
//			// //tbody/tr/td[8]/a/@href
//			r, _ := client.kskm(vUrl)
//			wr := restyUtil.GetCurlByR(*r)
//			fmt.Println(wr.String())
//
//			mItem, _ := htmlquery.Parse(strings.NewReader(r.String()))
//
//			tmpNode := htmlquery.FindOne(mItem, "//tbody[1]/tr[1]/td[2]")
//			zsdw := strings.TrimSpace(htmlquery.InnerText(tmpNode)) // 招生单位：
//
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[1]/tr[1]/td[4]")
//			fangshi := strings.TrimSpace(htmlquery.InnerText(tmpNode)) // 考试方式
//
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[1]/tr[2]/td[2]")
//			yx := strings.TrimSpace(htmlquery.InnerText(tmpNode)) //院系所
//
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[1]/tr[2]/td[4]")
//			zhuangye := strings.TrimSpace(htmlquery.InnerText(tmpNode)) //专业
//
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[1]/tr[3]/td[2]")
//			studyFunc := strings.TrimSpace(htmlquery.InnerText(tmpNode)) //学习方式：
//
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[1]/tr[3]/td[4]")
//			fangxiang := strings.TrimSpace(htmlquery.InnerText(tmpNode)) //研究方向：：
//
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[1]/tr[4]/td[2]")
//			teacherInfo := strings.TrimSpace(htmlquery.InnerText(tmpNode)) //学习方式
//
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[1]/tr[4]/td[4]")
//			number := strings.TrimSpace(htmlquery.InnerText(tmpNode)) //拟招人数：：
//
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[@class=\"zsml-res-items\"]/tr/td[1]")
//			Zz := strings.TrimSpace(htmlquery.InnerText(tmpNode))
//			Zz = strings.Split(Zz, "\n")[0]
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[@class=\"zsml-res-items\"]/tr/td[2]")
//			Wy := strings.TrimSpace(htmlquery.InnerText(tmpNode))
//			Wy = strings.Split(Wy, "\n")[0]
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[@class=\"zsml-res-items\"]/tr/td[3]")
//			B1 := strings.TrimSpace(htmlquery.InnerText(tmpNode))
//			B1 = strings.Split(B1, "\n")[0]
//			tmpNode = htmlquery.FindOne(mItem, "//tbody[@class=\"zsml-res-items\"]/tr/td[4]")
//			B2 := strings.TrimSpace(htmlquery.InnerText(tmpNode))
//			sTmp := SpecialityInfo{Zz: Zz, Wy: Wy, B1: B1, B2: B2,
//				Zsdw:        zsdw,        //招生单位：
//				Kxfs:        fangshi,     //考试方式
//				Zy:          zhuangye,    // 专业：	(083500)软件工程
//				Yx:          yx,          //院系所
//				Fx:          fangxiang,   //方向
//				StudyFunc:   studyFunc,   //方式
//				Number:      number,      //人数
//				TeacherInfo: teacherInfo, //指导老师：
//			}
//			helpers.FilePutContents("slog.log", helpers.JsonEncode(sTmp)+"\n", true)
//			fmt.Println(sTmp)
//			slist = append(slist, sTmp)
//		}
//		subList = append(subList, Speciality{Name: sItem.Name, Code: sItem.Code, Info: slist})
//	}
//	vList[key].SpecialityList = subList
//
//	city := ""
//	for sKey, sItem := range vList[key].SpecialityList {
//		r, _ := client.queryAction(vItem.Name, sItem.Code[:4], sItem.Name)
//		mItem, _ := htmlquery.Parse(strings.NewReader(r.String()))
//		tmpNode := htmlquery.FindOne(mItem, "//tbody/tr/td[2]")
//		if tmpNode == nil {
//			continue
//		}
//		if len(city) == 0 {
//			city = strings.TrimSpace(htmlquery.InnerText(tmpNode)) //
//		}
//		tmpNode = htmlquery.FindOne(mItem, "//tbody/tr/td[5]/i") //博士点
//		IsDoctor := "没有"
//		if tmpNode != nil {
//			IsDoctor = "有"
//		}
//
//		fmt.Println(sKey, city, IsDoctor)
//
//		vList[key].SpecialityList[sKey].IsDoctor = IsDoctor
//	}
//	fmt.Println(city)
//	vList[key].City = city
//
//}
