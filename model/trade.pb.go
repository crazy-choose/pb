// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: trade.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//成交
type TradeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrokerId       string  `protobuf:"bytes,1,opt,name=broker_id,json=brokerId,proto3" json:"broker_id,omitempty"`                       // 经纪公司代码
	InvestorId     string  `protobuf:"bytes,2,opt,name=investor_id,json=investorId,proto3" json:"investor_id,omitempty"`                 // 投资者代码
	Reserve_1      string  `protobuf:"bytes,3,opt,name=reserve_1,json=reserve1,proto3" json:"reserve_1,omitempty"`                       // 保留的无效字段
	OrderRef       string  `protobuf:"bytes,4,opt,name=order_ref,json=orderRef,proto3" json:"order_ref,omitempty"`                       // 报单引用
	UserId         string  `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                             // 用户代码
	ExchangeId     string  `protobuf:"bytes,6,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`                 // 交易所代码
	TradeId        string  `protobuf:"bytes,7,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`                          // 成交编号
	Direction      string  `protobuf:"bytes,8,opt,name=direction,proto3" json:"direction,omitempty"`                                     // 买卖方向
	OrderSysId     string  `protobuf:"bytes,9,opt,name=order_sys_id,json=orderSysId,proto3" json:"order_sys_id,omitempty"`               // 报单编号
	ParticipantId  string  `protobuf:"bytes,10,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`       // 会员代码
	ClientId       string  `protobuf:"bytes,11,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`                      // 客户代码
	TradingRole    string  `protobuf:"bytes,12,opt,name=trading_role,json=tradingRole,proto3" json:"trading_role,omitempty"`             // 交易角色
	Reserve_2      string  `protobuf:"bytes,13,opt,name=reserve_2,json=reserve2,proto3" json:"reserve_2,omitempty"`                      // 保留的无效字段
	OffsetFlag     string  `protobuf:"bytes,14,opt,name=offset_flag,json=offsetFlag,proto3" json:"offset_flag,omitempty"`                // 开平标志
	HedgeFlag      string  `protobuf:"bytes,15,opt,name=hedge_flag,json=hedgeFlag,proto3" json:"hedge_flag,omitempty"`                   // 投机套保标志
	Price          float64 `protobuf:"fixed64,16,opt,name=price,proto3" json:"price,omitempty"`                                          // 价格
	Volume         int32   `protobuf:"varint,17,opt,name=volume,proto3" json:"volume,omitempty"`                                         // 数量
	TradeDate      string  `protobuf:"bytes,18,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"`                   // 成交时期
	TradeTime      string  `protobuf:"bytes,19,opt,name=trade_time,json=tradeTime,proto3" json:"trade_time,omitempty"`                   // 成交时间
	TradeType      string  `protobuf:"bytes,20,opt,name=trade_type,json=tradeType,proto3" json:"trade_type,omitempty"`                   // 成交类型
	PriceSource    string  `protobuf:"bytes,21,opt,name=price_source,json=priceSource,proto3" json:"price_source,omitempty"`             // 成交价来源
	TraderId       string  `protobuf:"bytes,22,opt,name=trader_id,json=traderId,proto3" json:"trader_id,omitempty"`                      // 交易所交易员代码
	OrderLocalId   string  `protobuf:"bytes,23,opt,name=order_local_id,json=orderLocalId,proto3" json:"order_local_id,omitempty"`        // 本地报单编号
	ClearingPartId string  `protobuf:"bytes,24,opt,name=clearing_part_id,json=clearingPartId,proto3" json:"clearing_part_id,omitempty"`  // 结算会员编号
	BusinessUnit   string  `protobuf:"bytes,25,opt,name=business_unit,json=businessUnit,proto3" json:"business_unit,omitempty"`          // 业务单元
	SequenceNo     int32   `protobuf:"varint,26,opt,name=sequence_no,json=sequenceNo,proto3" json:"sequence_no,omitempty"`               // 序号
	TradingDay     string  `protobuf:"bytes,27,opt,name=trading_day,json=tradingDay,proto3" json:"trading_day,omitempty"`                // 交易日
	SettlementId   int32   `protobuf:"varint,28,opt,name=settlement_id,json=settlementId,proto3" json:"settlement_id,omitempty"`         // 结算编号
	BrokerOrderSeq int32   `protobuf:"varint,29,opt,name=broker_order_seq,json=brokerOrderSeq,proto3" json:"broker_order_seq,omitempty"` // 经纪公司报单编号
	TradeSource    string  `protobuf:"bytes,30,opt,name=trade_source,json=tradeSource,proto3" json:"trade_source,omitempty"`             // 成交来源
	InvestUnitId   string  `protobuf:"bytes,31,opt,name=invest_unit_id,json=investUnitId,proto3" json:"invest_unit_id,omitempty"`        // 投资单元代码
	InstrumentId   string  `protobuf:"bytes,32,opt,name=instrument_id,json=instrumentId,proto3" json:"instrument_id,omitempty"`          // 合约代码
	ExchangeInstId string  `protobuf:"bytes,33,opt,name=exchange_inst_id,json=exchangeInstId,proto3" json:"exchange_inst_id,omitempty"`  // 合约在交易所的代码
	DirectionZn    string  `protobuf:"bytes,34,opt,name=direction_zn,json=directionZn,proto3" json:"direction_zn,omitempty"`             // 买卖方向(转)
	OffsetFlagZn   string  `protobuf:"bytes,35,opt,name=offset_flag_zn,json=offsetFlagZn,proto3" json:"offset_flag_zn,omitempty"`
}

func (x *TradeInfo) Reset() {
	*x = TradeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeInfo) ProtoMessage() {}

func (x *TradeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeInfo.ProtoReflect.Descriptor instead.
func (*TradeInfo) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{0}
}

func (x *TradeInfo) GetBrokerId() string {
	if x != nil {
		return x.BrokerId
	}
	return ""
}

func (x *TradeInfo) GetInvestorId() string {
	if x != nil {
		return x.InvestorId
	}
	return ""
}

func (x *TradeInfo) GetReserve_1() string {
	if x != nil {
		return x.Reserve_1
	}
	return ""
}

func (x *TradeInfo) GetOrderRef() string {
	if x != nil {
		return x.OrderRef
	}
	return ""
}

func (x *TradeInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TradeInfo) GetExchangeId() string {
	if x != nil {
		return x.ExchangeId
	}
	return ""
}

func (x *TradeInfo) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

func (x *TradeInfo) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *TradeInfo) GetOrderSysId() string {
	if x != nil {
		return x.OrderSysId
	}
	return ""
}

func (x *TradeInfo) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *TradeInfo) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *TradeInfo) GetTradingRole() string {
	if x != nil {
		return x.TradingRole
	}
	return ""
}

func (x *TradeInfo) GetReserve_2() string {
	if x != nil {
		return x.Reserve_2
	}
	return ""
}

func (x *TradeInfo) GetOffsetFlag() string {
	if x != nil {
		return x.OffsetFlag
	}
	return ""
}

func (x *TradeInfo) GetHedgeFlag() string {
	if x != nil {
		return x.HedgeFlag
	}
	return ""
}

func (x *TradeInfo) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *TradeInfo) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *TradeInfo) GetTradeDate() string {
	if x != nil {
		return x.TradeDate
	}
	return ""
}

func (x *TradeInfo) GetTradeTime() string {
	if x != nil {
		return x.TradeTime
	}
	return ""
}

func (x *TradeInfo) GetTradeType() string {
	if x != nil {
		return x.TradeType
	}
	return ""
}

func (x *TradeInfo) GetPriceSource() string {
	if x != nil {
		return x.PriceSource
	}
	return ""
}

func (x *TradeInfo) GetTraderId() string {
	if x != nil {
		return x.TraderId
	}
	return ""
}

func (x *TradeInfo) GetOrderLocalId() string {
	if x != nil {
		return x.OrderLocalId
	}
	return ""
}

func (x *TradeInfo) GetClearingPartId() string {
	if x != nil {
		return x.ClearingPartId
	}
	return ""
}

func (x *TradeInfo) GetBusinessUnit() string {
	if x != nil {
		return x.BusinessUnit
	}
	return ""
}

func (x *TradeInfo) GetSequenceNo() int32 {
	if x != nil {
		return x.SequenceNo
	}
	return 0
}

func (x *TradeInfo) GetTradingDay() string {
	if x != nil {
		return x.TradingDay
	}
	return ""
}

func (x *TradeInfo) GetSettlementId() int32 {
	if x != nil {
		return x.SettlementId
	}
	return 0
}

func (x *TradeInfo) GetBrokerOrderSeq() int32 {
	if x != nil {
		return x.BrokerOrderSeq
	}
	return 0
}

func (x *TradeInfo) GetTradeSource() string {
	if x != nil {
		return x.TradeSource
	}
	return ""
}

func (x *TradeInfo) GetInvestUnitId() string {
	if x != nil {
		return x.InvestUnitId
	}
	return ""
}

func (x *TradeInfo) GetInstrumentId() string {
	if x != nil {
		return x.InstrumentId
	}
	return ""
}

func (x *TradeInfo) GetExchangeInstId() string {
	if x != nil {
		return x.ExchangeInstId
	}
	return ""
}

func (x *TradeInfo) GetDirectionZn() string {
	if x != nil {
		return x.DirectionZn
	}
	return ""
}

func (x *TradeInfo) GetOffsetFlagZn() string {
	if x != nil {
		return x.OffsetFlagZn
	}
	return ""
}

var File_trade_proto protoreflect.FileDescriptor

var file_trade_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x09,
	0x0a, 0x09, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x5f, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x31, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x66, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x79, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x79, 0x73, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x32, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x32, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1d, 0x0a,
	0x0a, 0x68, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x68, 0x65, 0x64, 0x67, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x71, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x55, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x7a, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x7a, 0x6e, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x5a, 0x6e, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61,
	0x7a, 0x79, 0x2d, 0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trade_proto_rawDescOnce sync.Once
	file_trade_proto_rawDescData = file_trade_proto_rawDesc
)

func file_trade_proto_rawDescGZIP() []byte {
	file_trade_proto_rawDescOnce.Do(func() {
		file_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_proto_rawDescData)
	})
	return file_trade_proto_rawDescData
}

var file_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_trade_proto_goTypes = []interface{}{
	(*TradeInfo)(nil), // 0: TradeInfo
}
var file_trade_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trade_proto_init() }
func file_trade_proto_init() {
	if File_trade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_proto_goTypes,
		DependencyIndexes: file_trade_proto_depIdxs,
		MessageInfos:      file_trade_proto_msgTypes,
	}.Build()
	File_trade_proto = out.File
	file_trade_proto_rawDesc = nil
	file_trade_proto_goTypes = nil
	file_trade_proto_depIdxs = nil
}
