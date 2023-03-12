package service

import (
	"context"
	"encoding/json"
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
)

func TestSendRedPacket(t *testing.T) {
	ctx := context.Background()
	db.InitDB()
	kv.InitRedis(ctx)

	// 关键点1， 使用gin的Router
	r := setupRouter()
	// 关键点2 构造请求body

	sReq := model.SendRpReq{
		UserId:   "test001",
		GroupId:  "testgroup001",
		Amount:   1000,
		Number:   20,
		BizOutNo: "biz_000001",
	}

	reqbody := strings.NewReader(utils.Json2String(sReq))
	req, err := http.NewRequest(http.MethodPost, "/red-packet/send", reqbody)
	if err != nil {
		t.Fatalf("构建请求失败, err: %v", err)
	}
	// 关键点3， 设置请求头，一定
	req.Header.Set("Content-Type", "application/json")
	// 构造一个记录
	rec := httptest.NewRecorder()
	//关键点4， 调用web服务的方法
	r.ServeHTTP(rec, req)
	result := rec.Result()
	if result.StatusCode != 200 {
		t.Fatalf("请求状态码不符合预期")
	}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("读取返回内容失败, err:%v", err)
	}
	defer result.Body.Close()
	var res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Fatalf("转换结果失败,err: %v", err)
	}
	if res.Code != 0 {
		t.Fatalf("结果不符合预期, 预期为:%v, 实际为：%v", 0, res.Code)
	}
	t.Log("用例测试通过")

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/red-packet/send", SendRedPacket)
	return r
}
