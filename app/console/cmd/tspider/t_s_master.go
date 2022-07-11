package tspider

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"math/rand"
	"net/url"
	"path"
	"strings"
	"sync"
	"thh/arms"
	"thh/arms/config"
	"time"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "tMaster",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   shadow3,
		//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func ifErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

type vTw struct {
	ScreenName string
	Name       string
	VerList    map[string]string
	Desc       string
}

var prefix string
var screenNameMap map[string]int
var key string
var vMap map[string]vTw
var vMapLock = &sync.Mutex{}

func shadow3(cmd *cobra.Command, args []string) {
	screenNameMap = make(map[string]int, 100)
	var maxRoutineNum = 3

	vMap = make(map[string]vTw, 0)

	prefix = "./storage/tmp/" + time.Now().Format("20060102_150405")
	key = "twitter:screenName:list"
	superT := super4Twitter{}
	allUsePush := config.GetBool("t_allUsePush", false)
	downMedia := config.GetBool("t_downMedia", false)
	screenNamesFromEnv := config.GetString("t_screenName", "")
	tOnlyMaster := config.GetBool("t_onlyMaster", false)
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
			superT.superT(superTConfig{
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

	arms.FilePutContents("./storage/tmp/"+time.Now().Format("20060102_150405")+"tjson.json", []byte(jsonStr), false)

	if tOnlyMaster {
		fmt.Println("仅仅抓取 screenName 列表，不向外扩展抓取")
		return
	}

	time.Sleep(15 * time.Second)
	var wg sync.WaitGroup
	go func() {
		for {
			fmt.Println("当前队列长度:", arms.QueueLen(key))
			time.Sleep(5 * time.Second)
		}
	}()
	for {
		screenName, err := arms.QueueLPop(key)
		if err != nil {
			break
		}
		if _, ok := screenNameMap[screenName]; ok {
			fmt.Println(screenName + "当前已经查询过，跳过")
			continue
		}
		screenNameMap[screenName] = 1
		ch <- 1
		wg.Add(1)
		fmt.Println(screenName, "ready~~~~~")
		go func(screenName string, usePush bool, ch chan int) {
			defer wg.Done()
			fmt.Println(screenName, "start~~~~~")
			superT.superT(superTConfig{
				screenName: screenName,
				usePush:    usePush,
				downMedia:  downMedia,
			})
			<-ch
		}(screenName, allUsePush, ch)
	}
	wg.Wait()

}

type super4Twitter struct {
}

type superTConfig struct {
	screenName string
	usePush    bool
	downMedia  bool
}

func (itself super4Twitter) superT(sConfig superTConfig) {
	tScreenNameList := config.GetString("t_screenName", "")
	screenName := sConfig.screenName
	usePush := sConfig.usePush
	downMedia := sConfig.downMedia
	// Create a Resty Client
	client := newTClient()

	sourcePrefix := prefix + "/zzzzzzzzzzzsource_" + screenName + "/"
	tMediaDir := prefix + "/" + screenName + "/"
	screenPath := prefix + "/zzzzzzzzzzzsource_" + screenName
	tListPath := prefix + "/" + screenName + "/list.txt"
	allListPath := prefix + "/list.txt"

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
	for {
		r, err = client.getTList(restId, 40, cursor)
		sourcePath := sourcePrefix + cast.ToString(time.Now().UnixMilli())
		GetDocumentByR(*r, sourcePath)
		TList := arms.JsonDecode[TList](r.String())
		//fmt.Println(TList)
		//fmt.Println(TList.Data.User.Result.TimelineV2.Timeline.Instructions)
		//fmt.Println(len(TList.Data.User.Result.TimelineV2.Timeline.Instructions))

		i++

		if len(TList.Data.User.Result.TimelineV2.Timeline.Instructions) == 0 {
			fmt.Println(screenName + "完成··································")
			break
		}

		activeCount := 0
		for _, value := range TList.Data.User.Result.TimelineV2.Timeline.Instructions {
			switch value.Type {
			case "TimelineAddEntries":
				for _, entry := range value.Entries {
					entryContent := entry.Content
					if entryContent.EntryType == "TimelineTimelineItem" {
						activeCount += 1
						userId := entry.Content.ItemContent.TweetResults.Result.RestId
						// 用户
						// masterUserResult = entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result
						// 当前作者
						// userResult := entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result
						// 原文作者（如果没有可能为非转发）

						userResult := entry.Content.ItemContent.TweetResults.Result.Legacy.RetweetedStatusResult.Result.Core.UserResults.Result
						masterUserResult := entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result

						conversationIdStr := entry.Content.ItemContent.TweetResults.Result.Legacy.ConversationIdStr
						itemScreenName := "https://twitter.com/" + screenName
						itemName := userResult.Legacy.Name
						itemUserDesc := userResult.Legacy.Description
						itemUserId := userId
						itemConversationIdStr := conversationIdStr
						itemTLink := "https://twitter.com/" + userResult.Legacy.ScreenName
						itemFullText := entry.Content.ItemContent.TweetResults.Result.Legacy.FullText

						// 当前非转发
						if userResult.Legacy.ScreenName == "" {
							itemScreenName = "https://twitter.com/" + screenName
							itemName = masterUserResult.Legacy.Name
							itemUserDesc = masterUserResult.Legacy.Description
							itemUserId = masterUserResult.RestId
						}

						msgStr := fmt.Sprintf("from:%v,name:%v\ndesc:%v\nuserId:%v\nConversationIdStr:%v\n%v\ntext:%v\n++++++++++++++++++++++++++++++++++++++++++++\n",
							itemScreenName,
							itemName,
							itemUserDesc,
							itemUserId,
							itemConversationIdStr,
							itemTLink,
							itemFullText,
						)

						// 允许后续扩散查询
						if usePush && userResult.Legacy.ScreenName != "" {
							arms.QueueRPush(key, userResult.Legacy.ScreenName)
							fmt.Println(userResult.Legacy.ScreenName, "进入后续查询队列")
						}

						if len(userResult.Legacy.ScreenName) > 0 && arms.InArray(screenName, strings.Split(tScreenNameList, ",")) {
							vMapLock.Lock()
							if v, ok := vMap[userResult.Legacy.ScreenName]; ok {
								if _, ok := v.VerList[screenName]; !ok {
									v.VerList[screenName] = fmt.Sprintf("https://twitter.com/%v/status/%v", screenName, conversationIdStr)
								}
							} else {
								vMap[userResult.Legacy.ScreenName] = vTw{
									ScreenName: userResult.Legacy.ScreenName,
									Name:       userResult.Legacy.Name,
									VerList: map[string]string{
										screenName: fmt.Sprintf("https://twitter.com/%v/status/%v", screenName, conversationIdStr),
									},
									Desc: userResult.Legacy.Description,
								}
							}
							vMapLock.Unlock()
						}
						// 推文
						// entry.Content.ItemContent.TweetResults.Result.Legacy

						if downMedia {
							arms.FilePutContents(tListPath, []byte(msgStr), true)
							arms.FilePutContents(allListPath, []byte(msgStr), true)
							medias := entry.Content.ItemContent.TweetResults.Result.Legacy.ExtendedEntities.Media
							for _, media := range medias {
								switch media.Type {
								case "photo":
									//u, _ := url.Parse(media.MediaUrlHttps)
									basename := path.Base(media.MediaUrlHttps)
									stdToolClient.downMedia(media.MediaUrlHttps, tMediaDir+conversationIdStr+basename)
									break
								case "video":
									// 下载封面
									basename := path.Base(media.MediaUrlHttps)
									stdToolClient.downMedia(media.MediaUrlHttps, tMediaDir+conversationIdStr+basename)
									// 下载视频
									variants := media.VideoInfo.Variants
									tmpBitrate := 0
									tmpUrl := ""
									for _, variant := range variants {
										if variant.Bitrate > tmpBitrate {
											u, err := url.Parse(variant.Url)
											if err != nil {
												fmt.Println("url解析失败")
												continue
											}
											basename = path.Base(u.Path)
											tmpUrl = variant.Url
										}
									}
									if len(tmpUrl) == 0 {
										fmt.Println("视频下载失败")
										break
									}
									stdToolClient.downMedia(tmpUrl, tMediaDir+conversationIdStr+basename)
									break
								default:
									fmt.Println(media.Type)
								}
							}
						}
					}
					// 选择下次节点
					if entryContent.CursorType == "Bottom" {
						cursor = entryContent.Value
					}
				}
				break
			case "TimelinePinEntry":
				for _, entry := range value.Entries {
					userId := entry.Content.ItemContent.TweetResults.Result.RestId
					fmt.Println(userId)
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
		fmt.Println(screenName, "下一轮", i*40, "-", (i+1)*40)

		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	}

	// Want to remove proxy setting
	//client.RemoveProxy()
}
