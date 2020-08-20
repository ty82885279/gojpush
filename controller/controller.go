package controller

import (
	"fmt"
	"gojpush/conf"
	"gojpush/model"

	jpushclient "github.com/ylywyn/jpush-api-go-client"

	"github.com/gin-gonic/gin"
)

type res struct {
	status int64
	msg    string
	data   interface{}
}

func SendMsg(c *gin.Context) {
	AppInfo := &conf.AppInfo{AppKey: c.Query("AppKey"), ResID: c.Query("ResID")}
	App := model.SendPush(AppInfo)
	//var r res
	//r = res{
	//	status: 200,
	//	msg:    "获取成功",
	//	data:   App,
	//}
	//c.JSON(http.StatusOK, gin.H{})

	var pf jpushclient.Platform
	pf.Add(jpushclient.IOS)

	var ad jpushclient.Audience

	ad.SetID([]string{AppInfo.ResID})

	var msg jpushclient.Message
	msg.Content = App.Address
	var op jpushclient.Option
	op.ApnsProduction = true

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetMessage(&msg)
	payload.SetAudience(&ad)
	payload.SetOptions(&op)

	bytes, _ := payload.ToBytes()
	fmt.Printf("%s\r\n", string(bytes))

	//push
	push := jpushclient.NewPushClient(App.MasterSecret, App.AppKey)
	str, err := push.Send(bytes)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	} else {
		fmt.Printf("ok:%s\n", str)
	}
}
