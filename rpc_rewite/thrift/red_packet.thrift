namespace go com.red_packet.rpc

// IDL  interface define language

struct SendRpReq {
    1: required string userId
    2: required string groupId
    3: required i64 amount
    4: required i64 number
    5: required string bizOutNo
}

struct SendRpResp {
    1: required string rpId
    2: required i64 errCode
    3: required string errMsg
}


// 接口定义
service RedPacketService {

    SendRpResp SendRp(1: SendRpReq req)

}
