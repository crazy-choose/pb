// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: instrument.proto

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

// 合约信息
type InstrumentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reserve_1              string  `protobuf:"bytes,1,opt,name=reserve_1,json=reserve1,proto3" json:"reserve_1,omitempty"`                                                // 保留的无效字段
	ExchangeId             string  `protobuf:"bytes,2,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`                                          // 交易所代码
	InstrumentName         string  `protobuf:"bytes,3,opt,name=instrument_name,json=instrumentName,proto3" json:"instrument_name,omitempty"`                              // 合约名称
	Reserve_2              string  `protobuf:"bytes,4,opt,name=reserve_2,json=reserve2,proto3" json:"reserve_2,omitempty"`                                                // 保留的无效字段
	Reserve_3              string  `protobuf:"bytes,5,opt,name=reserve_3,json=reserve3,proto3" json:"reserve_3,omitempty"`                                                // 保留的无效字段
	ProductClass           string  `protobuf:"bytes,6,opt,name=product_class,json=productClass,proto3" json:"product_class,omitempty"`                                    // 产品类型
	DeliveryYear           int32   `protobuf:"varint,7,opt,name=delivery_year,json=deliveryYear,proto3" json:"delivery_year,omitempty"`                                   // 交割年份
	DeliveryMonth          int32   `protobuf:"varint,8,opt,name=delivery_month,json=deliveryMonth,proto3" json:"delivery_month,omitempty"`                                // 交割月
	MaxMarketOrderVolume   int32   `protobuf:"varint,9,opt,name=max_market_order_volume,json=maxMarketOrderVolume,proto3" json:"max_market_order_volume,omitempty"`       // 市价单最大下单量
	MinMarketOrderVolume   int32   `protobuf:"varint,10,opt,name=min_market_order_volume,json=minMarketOrderVolume,proto3" json:"min_market_order_volume,omitempty"`      // 市价单最小下单量
	MaxLimitOrderVolume    int32   `protobuf:"varint,11,opt,name=max_limit_order_volume,json=maxLimitOrderVolume,proto3" json:"max_limit_order_volume,omitempty"`         // 限价单最大下单量
	MinLimitOrderVolume    int32   `protobuf:"varint,12,opt,name=min_limit_order_volume,json=minLimitOrderVolume,proto3" json:"min_limit_order_volume,omitempty"`         // 限价单最小下单量
	VolumeMultiple         int32   `protobuf:"varint,13,opt,name=volume_multiple,json=volumeMultiple,proto3" json:"volume_multiple,omitempty"`                            // 合约数量乘数
	PriceTick              float64 `protobuf:"fixed64,14,opt,name=price_tick,json=priceTick,proto3" json:"price_tick,omitempty"`                                          // 最小变动价位
	CreateDate             string  `protobuf:"bytes,15,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`                                         // 创建日
	OpenDate               string  `protobuf:"bytes,16,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`                                               // 上市日
	ExpireDate             string  `protobuf:"bytes,17,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`                                         // 到期日
	StartDelivDate         string  `protobuf:"bytes,18,opt,name=start_deliv_date,json=startDelivDate,proto3" json:"start_deliv_date,omitempty"`                           // 开始交割日
	EndDelivDate           string  `protobuf:"bytes,19,opt,name=end_deliv_date,json=endDelivDate,proto3" json:"end_deliv_date,omitempty"`                                 // 结束交割日
	InstLifePhase          string  `protobuf:"bytes,20,opt,name=inst_life_phase,json=instLifePhase,proto3" json:"inst_life_phase,omitempty"`                              // 合约生命周期状态
	IsTrading              int32   `protobuf:"varint,21,opt,name=is_trading,json=isTrading,proto3" json:"is_trading,omitempty"`                                           // 当前是否交易
	PositionType           string  `protobuf:"bytes,22,opt,name=position_type,json=positionType,proto3" json:"position_type,omitempty"`                                   // 持仓类型
	PositionDateType       string  `protobuf:"bytes,23,opt,name=position_date_type,json=positionDateType,proto3" json:"position_date_type,omitempty"`                     // 持仓日期类型
	LongMarginRatio        float64 `protobuf:"fixed64,24,opt,name=long_margin_ratio,json=longMarginRatio,proto3" json:"long_margin_ratio,omitempty"`                      // 多头保证金率
	ShortMarginRatio       float64 `protobuf:"fixed64,25,opt,name=short_margin_ratio,json=shortMarginRatio,proto3" json:"short_margin_ratio,omitempty"`                   // 空头保证金率
	MaxMarginSideAlgorithm string  `protobuf:"bytes,26,opt,name=max_margin_side_algorithm,json=maxMarginSideAlgorithm,proto3" json:"max_margin_side_algorithm,omitempty"` // 是否使用大额单边保证金算法
	Reserve_4              string  `protobuf:"bytes,27,opt,name=reserve_4,json=reserve4,proto3" json:"reserve_4,omitempty"`                                               // 保留的无效字段
	StrikePrice            float64 `protobuf:"fixed64,28,opt,name=strike_price,json=strikePrice,proto3" json:"strike_price,omitempty"`                                    // 执行价
	OptionsType            string  `protobuf:"bytes,29,opt,name=options_type,json=optionsType,proto3" json:"options_type,omitempty"`                                      // 期权类型
	UnderlyingMultiple     float64 `protobuf:"fixed64,30,opt,name=underlying_multiple,json=underlyingMultiple,proto3" json:"underlying_multiple,omitempty"`               // 合约基础商品乘数
	CombinationType        string  `protobuf:"bytes,31,opt,name=combination_type,json=combinationType,proto3" json:"combination_type,omitempty"`                          // 组合类型
	InstrumentId           string  `protobuf:"bytes,32,opt,name=instrument_id,json=instrumentId,proto3" json:"instrument_id,omitempty"`                                   // 合约代码
	ExchangeInstId         string  `protobuf:"bytes,33,opt,name=exchange_inst_id,json=exchangeInstId,proto3" json:"exchange_inst_id,omitempty"`                           // 合约在交易所的代码
	ProductId              string  `protobuf:"bytes,34,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`                                            // 产品代码
	UnderlyingInstrId      string  `protobuf:"bytes,35,opt,name=underlying_instr_id,json=underlyingInstrId,proto3" json:"underlying_instr_id,omitempty"`                  // 基础商品代码
}

func (x *InstrumentInfo) Reset() {
	*x = InstrumentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instrument_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstrumentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstrumentInfo) ProtoMessage() {}

func (x *InstrumentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_instrument_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstrumentInfo.ProtoReflect.Descriptor instead.
func (*InstrumentInfo) Descriptor() ([]byte, []int) {
	return file_instrument_proto_rawDescGZIP(), []int{0}
}

func (x *InstrumentInfo) GetReserve_1() string {
	if x != nil {
		return x.Reserve_1
	}
	return ""
}

func (x *InstrumentInfo) GetExchangeId() string {
	if x != nil {
		return x.ExchangeId
	}
	return ""
}

func (x *InstrumentInfo) GetInstrumentName() string {
	if x != nil {
		return x.InstrumentName
	}
	return ""
}

func (x *InstrumentInfo) GetReserve_2() string {
	if x != nil {
		return x.Reserve_2
	}
	return ""
}

func (x *InstrumentInfo) GetReserve_3() string {
	if x != nil {
		return x.Reserve_3
	}
	return ""
}

func (x *InstrumentInfo) GetProductClass() string {
	if x != nil {
		return x.ProductClass
	}
	return ""
}

func (x *InstrumentInfo) GetDeliveryYear() int32 {
	if x != nil {
		return x.DeliveryYear
	}
	return 0
}

func (x *InstrumentInfo) GetDeliveryMonth() int32 {
	if x != nil {
		return x.DeliveryMonth
	}
	return 0
}

func (x *InstrumentInfo) GetMaxMarketOrderVolume() int32 {
	if x != nil {
		return x.MaxMarketOrderVolume
	}
	return 0
}

func (x *InstrumentInfo) GetMinMarketOrderVolume() int32 {
	if x != nil {
		return x.MinMarketOrderVolume
	}
	return 0
}

func (x *InstrumentInfo) GetMaxLimitOrderVolume() int32 {
	if x != nil {
		return x.MaxLimitOrderVolume
	}
	return 0
}

func (x *InstrumentInfo) GetMinLimitOrderVolume() int32 {
	if x != nil {
		return x.MinLimitOrderVolume
	}
	return 0
}

func (x *InstrumentInfo) GetVolumeMultiple() int32 {
	if x != nil {
		return x.VolumeMultiple
	}
	return 0
}

func (x *InstrumentInfo) GetPriceTick() float64 {
	if x != nil {
		return x.PriceTick
	}
	return 0
}

func (x *InstrumentInfo) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *InstrumentInfo) GetOpenDate() string {
	if x != nil {
		return x.OpenDate
	}
	return ""
}

func (x *InstrumentInfo) GetExpireDate() string {
	if x != nil {
		return x.ExpireDate
	}
	return ""
}

func (x *InstrumentInfo) GetStartDelivDate() string {
	if x != nil {
		return x.StartDelivDate
	}
	return ""
}

func (x *InstrumentInfo) GetEndDelivDate() string {
	if x != nil {
		return x.EndDelivDate
	}
	return ""
}

func (x *InstrumentInfo) GetInstLifePhase() string {
	if x != nil {
		return x.InstLifePhase
	}
	return ""
}

func (x *InstrumentInfo) GetIsTrading() int32 {
	if x != nil {
		return x.IsTrading
	}
	return 0
}

func (x *InstrumentInfo) GetPositionType() string {
	if x != nil {
		return x.PositionType
	}
	return ""
}

func (x *InstrumentInfo) GetPositionDateType() string {
	if x != nil {
		return x.PositionDateType
	}
	return ""
}

func (x *InstrumentInfo) GetLongMarginRatio() float64 {
	if x != nil {
		return x.LongMarginRatio
	}
	return 0
}

func (x *InstrumentInfo) GetShortMarginRatio() float64 {
	if x != nil {
		return x.ShortMarginRatio
	}
	return 0
}

func (x *InstrumentInfo) GetMaxMarginSideAlgorithm() string {
	if x != nil {
		return x.MaxMarginSideAlgorithm
	}
	return ""
}

func (x *InstrumentInfo) GetReserve_4() string {
	if x != nil {
		return x.Reserve_4
	}
	return ""
}

func (x *InstrumentInfo) GetStrikePrice() float64 {
	if x != nil {
		return x.StrikePrice
	}
	return 0
}

func (x *InstrumentInfo) GetOptionsType() string {
	if x != nil {
		return x.OptionsType
	}
	return ""
}

func (x *InstrumentInfo) GetUnderlyingMultiple() float64 {
	if x != nil {
		return x.UnderlyingMultiple
	}
	return 0
}

func (x *InstrumentInfo) GetCombinationType() string {
	if x != nil {
		return x.CombinationType
	}
	return ""
}

func (x *InstrumentInfo) GetInstrumentId() string {
	if x != nil {
		return x.InstrumentId
	}
	return ""
}

func (x *InstrumentInfo) GetExchangeInstId() string {
	if x != nil {
		return x.ExchangeInstId
	}
	return ""
}

func (x *InstrumentInfo) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *InstrumentInfo) GetUnderlyingInstrId() string {
	if x != nil {
		return x.UnderlyingInstrId
	}
	return ""
}

// InstrumentStatus 合约状态
type InstrumentStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeId          string `protobuf:"bytes,1,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`                                  // 交易所代码	[9]byte
	Reserve_1           string `protobuf:"bytes,2,opt,name=reserve_1,json=reserve1,proto3" json:"reserve_1,omitempty"`                                        // 保留的无效字段	[31]byte
	SettlementGroupId   string `protobuf:"bytes,3,opt,name=settlement_group_id,json=settlementGroupId,proto3" json:"settlement_group_id,omitempty"`           // 结算组代码 [9]byte
	Reserve_2           string `protobuf:"bytes,4,opt,name=reserve_2,json=reserve2,proto3" json:"reserve_2,omitempty"`                                        // 保留的无效字段 [31]byte
	InstrumentStatus    string `protobuf:"bytes,5,opt,name=instrument_status,json=instrumentStatus,proto3" json:"instrument_status,omitempty"`                // 合约交易状态 byte
	TradingSegmentSn    int32  `protobuf:"varint,6,opt,name=trading_segment_sn,json=tradingSegmentSn,proto3" json:"trading_segment_sn,omitempty"`             // 交易阶段编号 int32
	EnterTime           string `protobuf:"bytes,7,opt,name=enter_time,json=enterTime,proto3" json:"enter_time,omitempty"`                                     // 进入本状态时间 [9]byte
	EnterReason         string `protobuf:"bytes,8,opt,name=enter_reason,json=enterReason,proto3" json:"enter_reason,omitempty"`                               // 进入本状态原因 byte
	ExchangeInstId      string `protobuf:"bytes,9,opt,name=exchange_inst_id,json=exchangeInstId,proto3" json:"exchange_inst_id,omitempty"`                    // 合约在交易所的代码 [81]byte
	InstrumentId        string `protobuf:"bytes,10,opt,name=instrument_id,json=instrumentId,proto3" json:"instrument_id,omitempty"`                           // 合约代码 [81]byte
	StatusZn            string `protobuf:"bytes,11,opt,name=status_zn,json=statusZn,proto3" json:"status_zn,omitempty"`                                       //当前状态中文
	PreTradingSegmentSn int32  `protobuf:"varint,12,opt,name=pre_trading_segment_sn,json=preTradingSegmentSn,proto3" json:"pre_trading_segment_sn,omitempty"` // 前交易阶段编号(自加)
	PreStatus           string `protobuf:"bytes,13,opt,name=pre_status,json=preStatus,proto3" json:"pre_status,omitempty"`                                    //前状态(自加)
	PreEnterTime        string `protobuf:"bytes,14,opt,name=pre_enter_time,json=preEnterTime,proto3" json:"pre_enter_time,omitempty"`                         //前状态进入时间(自加)
	PreEnterReason      string `protobuf:"bytes,15,opt,name=pre_enter_reason,json=preEnterReason,proto3" json:"pre_enter_reason,omitempty"`                   //前状态进入原因(自加)
}

func (x *InstrumentStatus) Reset() {
	*x = InstrumentStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instrument_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstrumentStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstrumentStatus) ProtoMessage() {}

func (x *InstrumentStatus) ProtoReflect() protoreflect.Message {
	mi := &file_instrument_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstrumentStatus.ProtoReflect.Descriptor instead.
func (*InstrumentStatus) Descriptor() ([]byte, []int) {
	return file_instrument_proto_rawDescGZIP(), []int{1}
}

func (x *InstrumentStatus) GetExchangeId() string {
	if x != nil {
		return x.ExchangeId
	}
	return ""
}

func (x *InstrumentStatus) GetReserve_1() string {
	if x != nil {
		return x.Reserve_1
	}
	return ""
}

func (x *InstrumentStatus) GetSettlementGroupId() string {
	if x != nil {
		return x.SettlementGroupId
	}
	return ""
}

func (x *InstrumentStatus) GetReserve_2() string {
	if x != nil {
		return x.Reserve_2
	}
	return ""
}

func (x *InstrumentStatus) GetInstrumentStatus() string {
	if x != nil {
		return x.InstrumentStatus
	}
	return ""
}

func (x *InstrumentStatus) GetTradingSegmentSn() int32 {
	if x != nil {
		return x.TradingSegmentSn
	}
	return 0
}

func (x *InstrumentStatus) GetEnterTime() string {
	if x != nil {
		return x.EnterTime
	}
	return ""
}

func (x *InstrumentStatus) GetEnterReason() string {
	if x != nil {
		return x.EnterReason
	}
	return ""
}

func (x *InstrumentStatus) GetExchangeInstId() string {
	if x != nil {
		return x.ExchangeInstId
	}
	return ""
}

func (x *InstrumentStatus) GetInstrumentId() string {
	if x != nil {
		return x.InstrumentId
	}
	return ""
}

func (x *InstrumentStatus) GetStatusZn() string {
	if x != nil {
		return x.StatusZn
	}
	return ""
}

func (x *InstrumentStatus) GetPreTradingSegmentSn() int32 {
	if x != nil {
		return x.PreTradingSegmentSn
	}
	return 0
}

func (x *InstrumentStatus) GetPreStatus() string {
	if x != nil {
		return x.PreStatus
	}
	return ""
}

func (x *InstrumentStatus) GetPreEnterTime() string {
	if x != nil {
		return x.PreEnterTime
	}
	return ""
}

func (x *InstrumentStatus) GetPreEnterReason() string {
	if x != nil {
		return x.PreEnterReason
	}
	return ""
}

var File_instrument_proto protoreflect.FileDescriptor

var file_instrument_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfd, 0x0a, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x31, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x5f, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x33, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x59, 0x65, 0x61, 0x72,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x17, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x14, 0x6d, 0x69, 0x6e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6d, 0x69,
	0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6d, 0x69, 0x6e, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x44, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x5f, 0x6c,
	0x69, 0x66, 0x65, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x69, 0x73, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x5f,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x18, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6c, 0x6f, 0x6e,
	0x67, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x2c, 0x0a, 0x12,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x18, 0x19, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4d,
	0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x39, 0x0a, 0x19, 0x6d, 0x61,
	0x78, 0x5f, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6d,
	0x61, 0x78, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x53, 0x69, 0x64, 0x65, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x5f, 0x34, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x34, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x75, 0x6e, 0x64, 0x65,
	0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e,
	0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x21, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x49, 0x64, 0x22, 0xca, 0x04, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x5f, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x31, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x5f, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x32, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x72, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x7a, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5a, 0x6e, 0x12, 0x33, 0x0a, 0x16, 0x70,
	0x72, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x70, 0x72, 0x65,
	0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x24, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42,
	0x0c, 0x5a, 0x0a, 0x2f, 0x6f, 0x75, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_instrument_proto_rawDescOnce sync.Once
	file_instrument_proto_rawDescData = file_instrument_proto_rawDesc
)

func file_instrument_proto_rawDescGZIP() []byte {
	file_instrument_proto_rawDescOnce.Do(func() {
		file_instrument_proto_rawDescData = protoimpl.X.CompressGZIP(file_instrument_proto_rawDescData)
	})
	return file_instrument_proto_rawDescData
}

var file_instrument_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_instrument_proto_goTypes = []interface{}{
	(*InstrumentInfo)(nil),   // 0: InstrumentInfo
	(*InstrumentStatus)(nil), // 1: InstrumentStatus
}
var file_instrument_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_instrument_proto_init() }
func file_instrument_proto_init() {
	if File_instrument_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_instrument_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstrumentInfo); i {
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
		file_instrument_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstrumentStatus); i {
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
			RawDescriptor: file_instrument_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_instrument_proto_goTypes,
		DependencyIndexes: file_instrument_proto_depIdxs,
		MessageInfos:      file_instrument_proto_msgTypes,
	}.Build()
	File_instrument_proto = out.File
	file_instrument_proto_rawDesc = nil
	file_instrument_proto_goTypes = nil
	file_instrument_proto_depIdxs = nil
}