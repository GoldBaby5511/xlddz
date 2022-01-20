package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

const (
	Send2All    uint32 = 1
	Send2AnyOne uint32 = 2
)

const (
	AppLogger   uint32 = 1
	AppCenter   uint32 = 2
	AppConfig   uint32 = 3
	AppGate     uint32 = 4
	AppLogin    uint32 = 5
	AppList     uint32 = 6
	AppProperty uint32 = 7
	AppTable    uint32 = 8
	AppRoom     uint32 = 9
	AppRobot    uint32 = 10
)

const (
	DataIndex  = 0 //数据
	AgentIndex = 1 //网络代理
	CMDIndex   = 2 //TCPCommon
	OtherIndex = 3 //其他
)

const (
	MinRouteArgsCount = 3
)

//网络命令
type TCPCommand struct {
	AppType uint16
	CmdId   uint16
}

//消息头内other字段常量
const (
	FlagOtherTraceId = 1
	TraceIdLen       = 16
)

type PackageHeader struct {
	version uint8
	encrypt uint8
	AppType uint16
	CmdId   uint16
	Other   string // 0xFF字节以内
}

// BaseMessage 基础消息结构
type BaseMessage struct {
	MyMessage interface{} //消息体
	AgentInfo BaseAgentInfo
	Cmd       TCPCommand //命令
	TraceId   string     //traceId
}

// MsgParser --------------
// | msgSize | headSize | header | msgData |
// | 4bit(msgSize) | 2bit(headSize) | 1bit(version) + 1bit(encrypt) + 2bit(AppType) + 2bit(CmdId) + Xbit(other) | msgData
// --------------
type MsgParser struct {
	lenMsgLen int
	minMsgLen uint32
	maxMsgLen uint32
}

func NewMsgParser() *MsgParser {
	p := new(MsgParser)
	p.lenMsgLen = 2
	p.minMsgLen = 8
	p.maxMsgLen = 8 * 1024

	return p
}

// It's dangerous to call the method on reading or writing
func (p *MsgParser) SetMsgLen(lenMsgLen int, minMsgLen uint32, maxMsgLen uint32) {
	if lenMsgLen == 1 || lenMsgLen == 2 || lenMsgLen == 4 {
		p.lenMsgLen = lenMsgLen
	}
	if minMsgLen != 0 {
		p.minMsgLen = minMsgLen
	}
	if maxMsgLen != 0 {
		p.maxMsgLen = maxMsgLen
	}

	var max uint32
	switch p.lenMsgLen {
	case 1:
		max = math.MaxUint8
	case 2:
		max = math.MaxUint16
	case 4:
		max = math.MaxUint32
	}
	if p.minMsgLen > max {
		p.minMsgLen = max
	}
	if p.maxMsgLen > max {
		p.maxMsgLen = max
	}
}

// |	msgSize	 |	headSize		| 						header 												| msgData
// |4bit(msgSize)| 2bit(headSize) 	| 1bit(version) + 1bit(encrypt) + 2bit(AppType) + 2bit(CmdId) + Xbit(other) | msgData
func (p *MsgParser) Read(conn *TCPConn) (BaseMessage, []byte, error) {
	msgSizeBuf := make([]byte, 4)
	if _, err := io.ReadFull(conn, msgSizeBuf); err != nil {
		return BaseMessage{}, nil, fmt.Errorf("消息头读取失败,%v", err)
	}

	var msgSize uint32 = 0
	if err := binary.Read(bytes.NewBuffer(msgSizeBuf), binary.BigEndian, &msgSize); err != nil {
		return BaseMessage{}, nil, fmt.Errorf("消息体长度读取失败,%v", err)
	}

	if msgSize < p.minMsgLen || msgSize > p.maxMsgLen {
		return BaseMessage{}, nil, fmt.Errorf("消息长度有问题,msgSize=%v,minMsgLen=%d,maxMsgLen=%d", msgSize, p.minMsgLen, p.maxMsgLen)
	}

	// data
	allData := make([]byte, msgSize)
	if _, err := io.ReadFull(conn, allData); err != nil {
		return BaseMessage{}, nil, fmt.Errorf("消息体内容读取失败,%v", err)
	}

	var headSize uint16 = 0
	_ = binary.Read(bytes.NewBuffer(allData[0:2]), binary.BigEndian, &headSize)

	header := &PackageHeader{}
	dataBuf := bytes.NewBuffer(allData[2:])
	_ = binary.Read(dataBuf, binary.BigEndian, &header.version)
	_ = binary.Read(dataBuf, binary.BigEndian, &header.encrypt)
	_ = binary.Read(dataBuf, binary.BigEndian, &header.AppType)
	_ = binary.Read(dataBuf, binary.BigEndian, &header.CmdId)

	//获取traceId，不做通用按字节去读，前8个字节是固定的，第9位如果等于1则紧跟在后边的16位就是traceId
	traceId := ""
	if len(allData) >= 8+1+TraceIdLen {
		//获取traceId == 1为具体标识
		var traceIdFlag uint8 = 0
		_ = binary.Read(bytes.NewBuffer(allData[8:8+1]), binary.BigEndian, &traceIdFlag)
		if traceIdFlag == FlagOtherTraceId {
			traceId = string(allData[8+1 : 8+1+TraceIdLen])
		}
	}

	//构造参数
	headCmd := &TCPCommand{AppType: header.AppType, CmdId: header.CmdId}
	msgData := allData[headSize+2:]
	bm := BaseMessage{Cmd: *headCmd, TraceId: traceId}
	return bm, msgData, nil
}

// |	msgSize	 |	headSize		| 						header 												| msgData
// |4bit(msgSize)| 2bit(headSize) 	| 1bit(version) + 1bit(encrypt) + 2bit(AppType) + 2bit(CmdId) + Xbit(other) | msgData
func (p *MsgParser) Write(appType, cmdId uint16, conn *TCPConn, msgData, otherData []byte) error {
	var headSize uint16 = 1 + 1 + 2 + 2 + uint16(len(otherData))
	var msgSize uint32 = 2 + uint32(headSize) + uint32(len(msgData))

	header := PackageHeader{uint8(99), uint8(104), appType, cmdId, ""}
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, msgSize)
	_ = binary.Write(buf, binary.BigEndian, headSize)
	_ = binary.Write(buf, binary.BigEndian, header.version)
	_ = binary.Write(buf, binary.BigEndian, header.encrypt)
	_ = binary.Write(buf, binary.BigEndian, header.AppType)
	_ = binary.Write(buf, binary.BigEndian, header.CmdId)
	if len(otherData) > 0 {
		_ = binary.Write(buf, binary.BigEndian, otherData)
	}
	_ = binary.Write(buf, binary.BigEndian, msgData)

	conn.Write(buf.Bytes())

	return nil
}
