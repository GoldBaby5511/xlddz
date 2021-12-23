package ddz

import (
	"github.com/golang/protobuf/proto"
	"xlddz/api/gameddz"
	"xlddz/cmd/table/business/table"
	"xlddz/pkg/log"
	n "xlddz/pkg/network"
)

type Sink struct {
	frame table.Frame
}

func (s *Sink) StartGame(f table.Frame) {
	s.frame = f
	var start gameddz.GameStart
	start.CurrentSeat = proto.Uint32(0)

	bm := n.BaseMessage{MyMessage: &start, TraceId: ""}
	bm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDTable), SubCmdID: uint16(gameddz.CMDID_Gameddz_IDGameStart)}
	s.frame.SendTableData(table.InvalidSeadID, bm)
}

func (s *Sink) EndGame() {

}

func (s *Sink) GameMessage(seatID, subCmdID uint32, data []byte) {
	switch subCmdID {
	case uint32(gameddz.CMDID_Gameddz_IDCallLandReq):
		s.CallLandReq(seatID, data)
	case uint32(gameddz.CMDID_Gameddz_IDOutCardReq):
		s.OutCardReq(seatID, data)
	case uint32(gameddz.CMDID_Gameddz_IDGameDataReq):
		s.GameDataReq(seatID, data)
	default:
		log.Warning("", "未定义消息,seatID=%d,subCmdID=%d", seatID, subCmdID)
	}
}

func (s *Sink) CallLandReq(seatID uint32, data []byte) {
	var m gameddz.CallLandReq
	_ = proto.Unmarshal(data, &m)

	var rsp gameddz.CallLandRsp
	bm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	bm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDTable), SubCmdID: uint16(gameddz.CMDID_Gameddz_IDCallLandRsp)}
	s.frame.SendTableData(table.InvalidSeadID, bm)

	log.Debug("", "叫地主消息,seatID=%d", seatID)
}

func (s *Sink) OutCardReq(seatID uint32, data []byte) {
	var m gameddz.OutCardReq
	_ = proto.Unmarshal(data, &m)

	var rsp gameddz.OutCardRsp
	bm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	bm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDTable), SubCmdID: uint16(gameddz.CMDID_Gameddz_IDOutCardRsp)}
	s.frame.SendTableData(seatID, bm)

	log.Debug("", "出牌消息,seatID=%d", seatID)
}

func (s *Sink) GameDataReq(seatID uint32, data []byte) {
	var m gameddz.GameDataReq
	_ = proto.Unmarshal(data, &m)

	var rsp gameddz.GameDataRsp
	bm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	bm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDTable), SubCmdID: uint16(gameddz.CMDID_Gameddz_IDGameDataRsp)}
	s.frame.SendTableData(table.InvalidSeadID, bm)

	log.Debug("", "数据消息,seatID=%d", seatID)
}
