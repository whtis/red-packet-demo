package main

import (
	"context"
	"fmt"
	"ginDemo/rpc_rewite/thrift/gen-go/com/red_packet/rpc"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"net"
	"os"
	"time"
)

func main() {
	startTime := currentTimeMillis()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "19091"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, _ := transportFactory.GetTransport(transport)
	client := rpc.NewRedPacketServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19090", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	for i := 0; i < 1000; i++ {
		req := &rpc.SendRpReq{
			UserId:   "rpc001",
			GroupId:  "rpcGroup001",
			Amount:   1000,
			Number:   10,
			BizOutNo: "biztestrpc001",
		}

		resp, err := client.SendRp(context.Background(), req)
		fmt.Println(i, "golang client: SendRp->", resp, err)
	}

	endTime := currentTimeMillis()
	fmt.Println("Program exit. time->", endTime, startTime, endTime-startTime)
}

// 转换成毫秒
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
