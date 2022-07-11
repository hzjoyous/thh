package tspider

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"math/rand"
	"strings"
	"sync"
	"thh/arms"
	"thh/arms/config"
	"time"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "tfollow",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   tfollow,
		//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func tfollow(cmd *cobra.Command, args []string) {
	screenNameMap = make(map[string]int, 100)
	var maxRoutineNum = 3

	vMap = make(map[string]vTw, 0)

	prefix = "./storage/tmp/" + time.Now().Format("20060102_150405")
	key = "twitter:screenName:list"
	superT := super4Twitter{}
	downMedia := config.GetBool("t_downMedia", false)
	screenNamesFromEnv := config.GetString("t_screenName", "")
	ch := make(chan int, maxRoutineNum)
	var wg4master sync.WaitGroup

	stdToolClient = newToolClient()

	dataList := strings.Split(screenNamesFromEnv, ",")

	if len(dataList) == 0 {
		fmt.Println("当前无配置")
		return
	}

	for _, jobScreenName := range dataList {
		wg4master.Add(1)
		ch <- 1
		go func(screenName string, ch chan int) {
			defer wg4master.Done()
			superT.superFollow(superTConfig{
				screenName: screenName,
				usePush:    true,
				downMedia:  downMedia,
			})
			<-ch
		}(jobScreenName, ch)
	}

	wg4master.Wait()

	vlist := []vTw{}
	for _, v := range vMap {
		vlist = append(vlist, v)
	}
	jsonStr := arms.JsonEncode(vlist)

	arms.FilePutContents("./storage/tmp/"+time.Now().Format("20060102150405")+"tfollow.json", []byte(jsonStr), false)

	fmt.Println("抓取关乎人列表完毕")
}

func (itself super4Twitter) superFollow(sConfig superTConfig) {
	tScreenNameList := config.GetString("t_screenName", "")
	screenName := sConfig.screenName
	// Create a Resty Client
	client := newTClient()

	sourcePrefix := prefix + "/zzzzzzzzzzzsource_" + screenName + "/"

	screenPath := prefix + "/zzzzzzzzzzzsource_" + screenName

	r, err := client.getUserInfo(screenName)
	GetDocumentByR(*r, screenPath)
	if ifErr(err) {
		return
	}
	userInfo := arms.JsonDecode[TUserInfo](r.String())
	restId := userInfo.Data.User.Result.RestID
	desc := userInfo.Data.User.Result.Legacy.Description
	fmt.Println(desc)
	linkList := []string{}
	for _, value := range userInfo.Data.User.Result.Legacy.Entities.URL.Urls {
		linkList = append(linkList, value.ExpandedURL)
	}

	cursor := ""

	// 谁 发了 什么
	// 谁 用户id 用户名 用户简介
	// 发了什么 内容 图片 视频地址
	i := 0
	pageCount := 20
	for {
		r, err = client.getFollowList(restId, pageCount, cursor)

		sourcePath := sourcePrefix + cast.ToString(time.Now().UnixMilli())
		GetDocumentByR(*r, sourcePath)

		TList := arms.JsonDecode[TFollowList](r.String())
		//fmt.Println(TList)
		//fmt.Println(TList.Data.User.Result.TimelineV2.Timeline.Instructions)
		//fmt.Println(len(TList.Data.User.Result.TimelineV2.Timeline.Instructions))

		i++

		if len(TList.Data.User.Result.Timeline.Timeline.Instructions) == 0 {
			fmt.Println(screenName + "完成··································")
			break
		}

		activeCount := 0
		for _, value := range TList.Data.User.Result.Timeline.Timeline.Instructions {
			switch value.Type {
			case "TimelineAddEntries":
				for _, entry := range value.Entries {
					entryContent := entry.Content
					if entryContent.EntryType == "TimelineTimelineItem" {
						activeCount += 1
						// 用户
						// masterUserResult = entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result
						// 当前作者
						// userResult := entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result
						// 原文作者（如果没有可能为非转发）

						userResult := entry.Content.ItemContent.UserResults.Result

						if len(userResult.Legacy.ScreenName) > 0 && arms.InArray(screenName, strings.Split(tScreenNameList, ",")) {
							vMapLock.Lock()
							if v, ok := vMap[userResult.Legacy.ScreenName]; ok {
								if _, ok := v.VerList[screenName]; !ok {
									v.VerList[screenName] = fmt.Sprintf("https://twitter.com/%v", screenName)
								}
							} else {
								vMap[userResult.Legacy.ScreenName] = vTw{
									ScreenName: userResult.Legacy.ScreenName,
									Name:       userResult.Legacy.Name,
									VerList: map[string]string{
										screenName: fmt.Sprintf("https://twitter.com/%v", screenName),
									},
									Desc: userResult.Legacy.Description,
								}
							}
							vMapLock.Unlock()
						}
						// 推文
						// entry.Content.ItemContent.TweetResults.Result.Legacy

					}
					// 选择下次节点
					if entryContent.CursorType == "Bottom" {
						cursor = entryContent.Value
					}
				}
				break
			default:
				fmt.Println(value.Type)
			}
		}

		if activeCount == 0 {
			fmt.Println(screenName + "完成··································")
			break
		}
		fmt.Println(screenName, "下一轮", i*pageCount, "-", (i+1)*pageCount)

		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	}

	// Want to remove proxy setting
	//client.RemoveProxy()
}
