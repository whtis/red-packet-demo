package main

import (
	"context"
	"fmt"
	"ginDemo/rpc_rewite/thrift/gen-go/com/red_packet/rpc"
	"ginDemo/rpc_rewite/utils"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

const (
	NetworkAddr = "127.0.0.1:19099"
)

type RedPacketServiceImpl struct {
}

func (s RedPacketServiceImpl) SendRp(ctx context.Context, req *rpc.SendRpReq) (*rpc.SendRpResp, error) {
	fmt.Println("golang server: go req: ", utils.Json2String(req))
	resp := &rpc.SendRpResp{
		RpId:    "goRPC:rpId0001",
		ErrCode: 0,
		ErrMsg:  "success",
	}
	return resp, nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RedPacketServiceImpl{}
	processor := rpc.NewRedPacketServiceProcessor(*handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
