package protocol

import (
	"encoding/json"
)

type RequestMessage struct {
	MsgType      int32
	RequiredBody RegisterReqBody
	OptionalBody interface{}
}

type ResponseMessage struct {
	MsgType      int32
	RequiredBody RegisterRespBody
	OptionalBody interface{}
}

/*
 * ResponseMessage转string
 */
func (m *ResponseMessage) String() string {
	byteM, _ := json.Marshal(m)
	return string(byteM)
}

/*
 * string转RequestMessage
 */
func FromString(s string) (msg *RequestMessage, err error) {
	msg = &RequestMessage{}
	if err := json.Unmarshal([]byte(s), msg); err != nil {
		return nil, err
	}

	optionalBodyMap, ok := msg.OptionalBody.(map[string]interface{})
	if !ok {
		msg.OptionalBody = nil
		return msg, nil
	}

	switch msg.MsgType {
	case RegisterReq, LoginReq:
		msg.OptionalBody = &UserInfo{
			UserName: optionalBodyMap["UserName"].(string),
			Password: optionalBodyMap["Password"].(string),
		}
	case NewTableReq, JoinTableReq, ExitTableReq:
		msg.OptionalBody = &TableInfo{
			TableId: optionalBodyMap["TableId"].(int64),
		}
	case DealPokersReq, ShotPokersReq:
		msg.OptionalBody = &PokersInfo{
			Pokers: optionalBodyMap["Pokers"].([]int),
		}
	case SyncTableInfoReq:
		// todo 同步信息struct定义
	default:
		msg.OptionalBody = nil
	}

	return msg, nil
}
