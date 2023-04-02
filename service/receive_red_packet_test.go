package service

import (
	"context"
	"encoding/json"
	"fmt"
	"ginDemo/dal/db"
	"ginDemo/dal/kv"
	"ginDemo/model"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestReceiveRedPacket(t *testing.T) {
	ctx := context.Background()
	db.InitDB()
	kv.InitRedis(ctx)

	rpId := "00cbaf2ad6d64a2c9cc68780f4a692cc"
	uniqueId := "fff"
	for i := 0; i < 2000; i++ {
		userid := fmt.Sprintf("userid_random%s%d", uniqueId, i)
		bizOutNo := fmt.Sprintf("biz_out_receive_00%s%d", uniqueId, i)
		// go 协程
		go func() {
			receiveTest(rpId, userid, bizOutNo)
		}()
	}

	time.Sleep(5 * time.Second)

}

func receiveTest(rpId, userId, bizOutNo string) {
	fmt.Println(fmt.Sprintf("start receive rp..."))
	// 关键点1， 使用gin的Router
	r := setupReceiveRouter()
	// 关键点2 构造请求body

	rReq := model.ReceiveRpReq{
		UserId:   userId,
		GroupId:  "testgroup001",
		RpId:     rpId,
		BizOutNo: bizOutNo,
	}

	reqbody := strings.NewReader(utils.Json2String(rReq))
	req, err := http.NewRequest(http.MethodPost, "/red-packet/receive", reqbody)
	if err != nil {
		fmt.Println(fmt.Sprintf("构建请求失败, err: %v", err))
	}
	// 关键点3， 设置请求头，一定
	req.Header.Set("Content-Type", "application/json")
	// 构造一个记录
	rec := httptest.NewRecorder()
	//关键点4， 调用web服务的方法
	r.ServeHTTP(rec, req)
	result := rec.Result()
	if result.StatusCode != 200 {
		fmt.Println("请求状态码不符合预期")
		return
	}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("读取返回内容失败, err:%v", err))
	}
	defer result.Body.Close()
	var res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(fmt.Sprintf("转换结果失败,err: %v", err))
	}
	if res.Code != 0 {
		fmt.Println(fmt.Sprintf("结果不符合预期, 预期为:%v, 实际为：%v", 0, res.Code))
	}
	fmt.Println(fmt.Sprintf("用例测试通过, resp: %v", res.Data))
}

func setupReceiveRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/red-packet/receive", ReceiveRedPacket)
	return r
}
