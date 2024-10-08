// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: forwarder/realtime.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*SnapshotMessage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SnapshotResponse) Reset() {
	*x = SnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_realtime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotResponse) ProtoMessage() {}

func (x *SnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_realtime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotResponse.ProtoReflect.Descriptor instead.
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	return file_forwarder_realtime_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotResponse) GetData() []*SnapshotMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type YahooFinancePrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Last  float64 `protobuf:"fixed64,2,opt,name=last,proto3" json:"last,omitempty"`
}

func (x *YahooFinancePrice) Reset() {
	*x = YahooFinancePrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_realtime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YahooFinancePrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YahooFinancePrice) ProtoMessage() {}

func (x *YahooFinancePrice) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_realtime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YahooFinancePrice.ProtoReflect.Descriptor instead.
func (*YahooFinancePrice) Descriptor() ([]byte, []int) {
	return file_forwarder_realtime_proto_rawDescGZIP(), []int{1}
}

func (x *YahooFinancePrice) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *YahooFinancePrice) GetLast() float64 {
	if x != nil {
		return x.Last
	}
	return 0
}

type VolumeRankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Date  string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *VolumeRankRequest) Reset() {
	*x = VolumeRankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_realtime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeRankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeRankRequest) ProtoMessage() {}

func (x *VolumeRankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_realtime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeRankRequest.ProtoReflect.Descriptor instead.
func (*VolumeRankRequest) Descriptor() ([]byte, []int) {
	return file_forwarder_realtime_proto_rawDescGZIP(), []int{2}
}

func (x *VolumeRankRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *VolumeRankRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type SnapshotMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts              int64   `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Code            string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Exchange        string  `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Open            float64 `protobuf:"fixed64,4,opt,name=open,proto3" json:"open,omitempty"`
	High            float64 `protobuf:"fixed64,5,opt,name=high,proto3" json:"high,omitempty"`
	Low             float64 `protobuf:"fixed64,6,opt,name=low,proto3" json:"low,omitempty"`
	Close           float64 `protobuf:"fixed64,7,opt,name=close,proto3" json:"close,omitempty"`
	TickType        string  `protobuf:"bytes,8,opt,name=tick_type,json=tickType,proto3" json:"tick_type,omitempty"`
	ChangePrice     float64 `protobuf:"fixed64,9,opt,name=change_price,json=changePrice,proto3" json:"change_price,omitempty"`
	ChangeRate      float64 `protobuf:"fixed64,10,opt,name=change_rate,json=changeRate,proto3" json:"change_rate,omitempty"`
	ChangeType      string  `protobuf:"bytes,11,opt,name=change_type,json=changeType,proto3" json:"change_type,omitempty"`
	AveragePrice    float64 `protobuf:"fixed64,12,opt,name=average_price,json=averagePrice,proto3" json:"average_price,omitempty"`
	Volume          int64   `protobuf:"varint,13,opt,name=volume,proto3" json:"volume,omitempty"`
	TotalVolume     int64   `protobuf:"varint,14,opt,name=total_volume,json=totalVolume,proto3" json:"total_volume,omitempty"`
	Amount          int64   `protobuf:"varint,15,opt,name=amount,proto3" json:"amount,omitempty"`
	TotalAmount     int64   `protobuf:"varint,16,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	YesterdayVolume float64 `protobuf:"fixed64,17,opt,name=yesterday_volume,json=yesterdayVolume,proto3" json:"yesterday_volume,omitempty"`
	BuyPrice        float64 `protobuf:"fixed64,18,opt,name=buy_price,json=buyPrice,proto3" json:"buy_price,omitempty"`
	BuyVolume       float64 `protobuf:"fixed64,19,opt,name=buy_volume,json=buyVolume,proto3" json:"buy_volume,omitempty"`
	SellPrice       float64 `protobuf:"fixed64,20,opt,name=sell_price,json=sellPrice,proto3" json:"sell_price,omitempty"`
	SellVolume      int64   `protobuf:"varint,21,opt,name=sell_volume,json=sellVolume,proto3" json:"sell_volume,omitempty"`
	VolumeRatio     float64 `protobuf:"fixed64,22,opt,name=volume_ratio,json=volumeRatio,proto3" json:"volume_ratio,omitempty"`
}

func (x *SnapshotMessage) Reset() {
	*x = SnapshotMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_realtime_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotMessage) ProtoMessage() {}

func (x *SnapshotMessage) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_realtime_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotMessage.ProtoReflect.Descriptor instead.
func (*SnapshotMessage) Descriptor() ([]byte, []int) {
	return file_forwarder_realtime_proto_rawDescGZIP(), []int{3}
}

func (x *SnapshotMessage) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *SnapshotMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SnapshotMessage) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *SnapshotMessage) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *SnapshotMessage) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *SnapshotMessage) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *SnapshotMessage) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *SnapshotMessage) GetTickType() string {
	if x != nil {
		return x.TickType
	}
	return ""
}

func (x *SnapshotMessage) GetChangePrice() float64 {
	if x != nil {
		return x.ChangePrice
	}
	return 0
}

func (x *SnapshotMessage) GetChangeRate() float64 {
	if x != nil {
		return x.ChangeRate
	}
	return 0
}

func (x *SnapshotMessage) GetChangeType() string {
	if x != nil {
		return x.ChangeType
	}
	return ""
}

func (x *SnapshotMessage) GetAveragePrice() float64 {
	if x != nil {
		return x.AveragePrice
	}
	return 0
}

func (x *SnapshotMessage) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *SnapshotMessage) GetTotalVolume() int64 {
	if x != nil {
		return x.TotalVolume
	}
	return 0
}

func (x *SnapshotMessage) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SnapshotMessage) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *SnapshotMessage) GetYesterdayVolume() float64 {
	if x != nil {
		return x.YesterdayVolume
	}
	return 0
}

func (x *SnapshotMessage) GetBuyPrice() float64 {
	if x != nil {
		return x.BuyPrice
	}
	return 0
}

func (x *SnapshotMessage) GetBuyVolume() float64 {
	if x != nil {
		return x.BuyVolume
	}
	return 0
}

func (x *SnapshotMessage) GetSellPrice() float64 {
	if x != nil {
		return x.SellPrice
	}
	return 0
}

func (x *SnapshotMessage) GetSellVolume() int64 {
	if x != nil {
		return x.SellVolume
	}
	return 0
}

func (x *SnapshotMessage) GetVolumeRatio() float64 {
	if x != nil {
		return x.VolumeRatio
	}
	return 0
}

type StockVolumeRankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*StockVolumeRankMessage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *StockVolumeRankResponse) Reset() {
	*x = StockVolumeRankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_realtime_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockVolumeRankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockVolumeRankResponse) ProtoMessage() {}

func (x *StockVolumeRankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_realtime_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockVolumeRankResponse.ProtoReflect.Descriptor instead.
func (*StockVolumeRankResponse) Descriptor() ([]byte, []int) {
	return file_forwarder_realtime_proto_rawDescGZIP(), []int{4}
}

func (x *StockVolumeRankResponse) GetData() []*StockVolumeRankMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type StockVolumeRankMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date            string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Code            string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name            string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Ts              int64   `protobuf:"varint,4,opt,name=ts,proto3" json:"ts,omitempty"`
	Open            float64 `protobuf:"fixed64,5,opt,name=open,proto3" json:"open,omitempty"`
	High            float64 `protobuf:"fixed64,6,opt,name=high,proto3" json:"high,omitempty"`
	Low             float64 `protobuf:"fixed64,7,opt,name=low,proto3" json:"low,omitempty"`
	Close           float64 `protobuf:"fixed64,8,opt,name=close,proto3" json:"close,omitempty"`
	PriceRange      float64 `protobuf:"fixed64,9,opt,name=price_range,json=priceRange,proto3" json:"price_range,omitempty"`
	TickType        int64   `protobuf:"varint,10,opt,name=tick_type,json=tickType,proto3" json:"tick_type,omitempty"`
	ChangePrice     float64 `protobuf:"fixed64,11,opt,name=change_price,json=changePrice,proto3" json:"change_price,omitempty"`
	ChangeType      int64   `protobuf:"varint,12,opt,name=change_type,json=changeType,proto3" json:"change_type,omitempty"`
	AveragePrice    float64 `protobuf:"fixed64,13,opt,name=average_price,json=averagePrice,proto3" json:"average_price,omitempty"`
	Volume          int64   `protobuf:"varint,14,opt,name=volume,proto3" json:"volume,omitempty"`
	TotalVolume     int64   `protobuf:"varint,15,opt,name=total_volume,json=totalVolume,proto3" json:"total_volume,omitempty"`
	Amount          int64   `protobuf:"varint,16,opt,name=amount,proto3" json:"amount,omitempty"`
	TotalAmount     int64   `protobuf:"varint,17,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	YesterdayVolume int64   `protobuf:"varint,18,opt,name=yesterday_volume,json=yesterdayVolume,proto3" json:"yesterday_volume,omitempty"`
	VolumeRatio     float64 `protobuf:"fixed64,19,opt,name=volume_ratio,json=volumeRatio,proto3" json:"volume_ratio,omitempty"`
	BuyPrice        float64 `protobuf:"fixed64,20,opt,name=buy_price,json=buyPrice,proto3" json:"buy_price,omitempty"`
	BuyVolume       int64   `protobuf:"varint,21,opt,name=buy_volume,json=buyVolume,proto3" json:"buy_volume,omitempty"`
	SellPrice       float64 `protobuf:"fixed64,22,opt,name=sell_price,json=sellPrice,proto3" json:"sell_price,omitempty"`
	SellVolume      int64   `protobuf:"varint,23,opt,name=sell_volume,json=sellVolume,proto3" json:"sell_volume,omitempty"`
	BidOrders       int64   `protobuf:"varint,24,opt,name=bid_orders,json=bidOrders,proto3" json:"bid_orders,omitempty"`
	BidVolumes      int64   `protobuf:"varint,25,opt,name=bid_volumes,json=bidVolumes,proto3" json:"bid_volumes,omitempty"`
	AskOrders       int64   `protobuf:"varint,26,opt,name=ask_orders,json=askOrders,proto3" json:"ask_orders,omitempty"`
	AskVolumes      int64   `protobuf:"varint,27,opt,name=ask_volumes,json=askVolumes,proto3" json:"ask_volumes,omitempty"`
}

func (x *StockVolumeRankMessage) Reset() {
	*x = StockVolumeRankMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_realtime_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockVolumeRankMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockVolumeRankMessage) ProtoMessage() {}

func (x *StockVolumeRankMessage) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_realtime_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockVolumeRankMessage.ProtoReflect.Descriptor instead.
func (*StockVolumeRankMessage) Descriptor() ([]byte, []int) {
	return file_forwarder_realtime_proto_rawDescGZIP(), []int{5}
}

func (x *StockVolumeRankMessage) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *StockVolumeRankMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StockVolumeRankMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StockVolumeRankMessage) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *StockVolumeRankMessage) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *StockVolumeRankMessage) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *StockVolumeRankMessage) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *StockVolumeRankMessage) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *StockVolumeRankMessage) GetPriceRange() float64 {
	if x != nil {
		return x.PriceRange
	}
	return 0
}

func (x *StockVolumeRankMessage) GetTickType() int64 {
	if x != nil {
		return x.TickType
	}
	return 0
}

func (x *StockVolumeRankMessage) GetChangePrice() float64 {
	if x != nil {
		return x.ChangePrice
	}
	return 0
}

func (x *StockVolumeRankMessage) GetChangeType() int64 {
	if x != nil {
		return x.ChangeType
	}
	return 0
}

func (x *StockVolumeRankMessage) GetAveragePrice() float64 {
	if x != nil {
		return x.AveragePrice
	}
	return 0
}

func (x *StockVolumeRankMessage) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *StockVolumeRankMessage) GetTotalVolume() int64 {
	if x != nil {
		return x.TotalVolume
	}
	return 0
}

func (x *StockVolumeRankMessage) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *StockVolumeRankMessage) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *StockVolumeRankMessage) GetYesterdayVolume() int64 {
	if x != nil {
		return x.YesterdayVolume
	}
	return 0
}

func (x *StockVolumeRankMessage) GetVolumeRatio() float64 {
	if x != nil {
		return x.VolumeRatio
	}
	return 0
}

func (x *StockVolumeRankMessage) GetBuyPrice() float64 {
	if x != nil {
		return x.BuyPrice
	}
	return 0
}

func (x *StockVolumeRankMessage) GetBuyVolume() int64 {
	if x != nil {
		return x.BuyVolume
	}
	return 0
}

func (x *StockVolumeRankMessage) GetSellPrice() float64 {
	if x != nil {
		return x.SellPrice
	}
	return 0
}

func (x *StockVolumeRankMessage) GetSellVolume() int64 {
	if x != nil {
		return x.SellVolume
	}
	return 0
}

func (x *StockVolumeRankMessage) GetBidOrders() int64 {
	if x != nil {
		return x.BidOrders
	}
	return 0
}

func (x *StockVolumeRankMessage) GetBidVolumes() int64 {
	if x != nil {
		return x.BidVolumes
	}
	return 0
}

func (x *StockVolumeRankMessage) GetAskOrders() int64 {
	if x != nil {
		return x.AskOrders
	}
	return 0
}

func (x *StockVolumeRankMessage) GetAskVolumes() int64 {
	if x != nil {
		return x.AskVolumes
	}
	return 0
}

var File_forwarder_realtime_proto protoreflect.FileDescriptor

var file_forwarder_realtime_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x61, 0x6c,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x10, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d,
	0x0a, 0x11, 0x59, 0x61, 0x68, 0x6f, 0x6f, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x22, 0x3d, 0x0a,
	0x11, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x88, 0x05, 0x0a,
	0x0f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x79, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x79, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x79, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64,
	0x61, 0x79, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x62, 0x75, 0x79,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x5f, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x62, 0x75, 0x79, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6c, 0x6c, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x50, 0x0a, 0x17, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9b, 0x06, 0x0a, 0x16, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x79, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x79, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x79, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64,
	0x61, 0x79, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x75, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x62, 0x75, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x5f,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x75,
	0x79, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x6c, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73, 0x65, 0x6c,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6c,
	0x6c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x69, 0x64, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x69, 0x64,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x6b, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x73, 0x6b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x73, 0x6b,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x32, 0x95, 0x05, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1b, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x41, 0x72, 0x72, 0x12, 0x16, 0x2e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x41, 0x72, 0x72, 0x1a, 0x1b, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x53, 0x45, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1b, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x4f, 0x54, 0x43, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1b, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x73, 0x64, 0x61, 0x71, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x59, 0x61, 0x68, 0x6f, 0x6f, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x73, 0x64, 0x61, 0x71,
	0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c,
	0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x59, 0x61, 0x68, 0x6f, 0x6f,
	0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x42, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x41, 0x72, 0x72, 0x12, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x72,
	0x1a, 0x1b, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_forwarder_realtime_proto_rawDescOnce sync.Once
	file_forwarder_realtime_proto_rawDescData = file_forwarder_realtime_proto_rawDesc
)

func file_forwarder_realtime_proto_rawDescGZIP() []byte {
	file_forwarder_realtime_proto_rawDescOnce.Do(func() {
		file_forwarder_realtime_proto_rawDescData = protoimpl.X.CompressGZIP(file_forwarder_realtime_proto_rawDescData)
	})
	return file_forwarder_realtime_proto_rawDescData
}

var file_forwarder_realtime_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_forwarder_realtime_proto_goTypes = []any{
	(*SnapshotResponse)(nil),        // 0: forwarder.SnapshotResponse
	(*YahooFinancePrice)(nil),       // 1: forwarder.YahooFinancePrice
	(*VolumeRankRequest)(nil),       // 2: forwarder.VolumeRankRequest
	(*SnapshotMessage)(nil),         // 3: forwarder.SnapshotMessage
	(*StockVolumeRankResponse)(nil), // 4: forwarder.StockVolumeRankResponse
	(*StockVolumeRankMessage)(nil),  // 5: forwarder.StockVolumeRankMessage
	(*emptypb.Empty)(nil),           // 6: google.protobuf.Empty
	(*StockNumArr)(nil),             // 7: forwarder.StockNumArr
	(*FutureCodeArr)(nil),           // 8: forwarder.FutureCodeArr
}
var file_forwarder_realtime_proto_depIdxs = []int32{
	3,  // 0: forwarder.SnapshotResponse.data:type_name -> forwarder.SnapshotMessage
	5,  // 1: forwarder.StockVolumeRankResponse.data:type_name -> forwarder.StockVolumeRankMessage
	6,  // 2: forwarder.RealTimeDataInterface.GetAllStockSnapshot:input_type -> google.protobuf.Empty
	7,  // 3: forwarder.RealTimeDataInterface.GetStockSnapshotByNumArr:input_type -> forwarder.StockNumArr
	6,  // 4: forwarder.RealTimeDataInterface.GetStockSnapshotTSE:input_type -> google.protobuf.Empty
	6,  // 5: forwarder.RealTimeDataInterface.GetStockSnapshotOTC:input_type -> google.protobuf.Empty
	6,  // 6: forwarder.RealTimeDataInterface.GetNasdaq:input_type -> google.protobuf.Empty
	6,  // 7: forwarder.RealTimeDataInterface.GetNasdaqFuture:input_type -> google.protobuf.Empty
	2,  // 8: forwarder.RealTimeDataInterface.GetStockVolumeRank:input_type -> forwarder.VolumeRankRequest
	8,  // 9: forwarder.RealTimeDataInterface.GetFutureSnapshotByCodeArr:input_type -> forwarder.FutureCodeArr
	0,  // 10: forwarder.RealTimeDataInterface.GetAllStockSnapshot:output_type -> forwarder.SnapshotResponse
	0,  // 11: forwarder.RealTimeDataInterface.GetStockSnapshotByNumArr:output_type -> forwarder.SnapshotResponse
	0,  // 12: forwarder.RealTimeDataInterface.GetStockSnapshotTSE:output_type -> forwarder.SnapshotResponse
	0,  // 13: forwarder.RealTimeDataInterface.GetStockSnapshotOTC:output_type -> forwarder.SnapshotResponse
	1,  // 14: forwarder.RealTimeDataInterface.GetNasdaq:output_type -> forwarder.YahooFinancePrice
	1,  // 15: forwarder.RealTimeDataInterface.GetNasdaqFuture:output_type -> forwarder.YahooFinancePrice
	4,  // 16: forwarder.RealTimeDataInterface.GetStockVolumeRank:output_type -> forwarder.StockVolumeRankResponse
	0,  // 17: forwarder.RealTimeDataInterface.GetFutureSnapshotByCodeArr:output_type -> forwarder.SnapshotResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_forwarder_realtime_proto_init() }
func file_forwarder_realtime_proto_init() {
	if File_forwarder_realtime_proto != nil {
		return
	}
	file_forwarder_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_forwarder_realtime_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SnapshotResponse); i {
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
		file_forwarder_realtime_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*YahooFinancePrice); i {
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
		file_forwarder_realtime_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VolumeRankRequest); i {
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
		file_forwarder_realtime_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SnapshotMessage); i {
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
		file_forwarder_realtime_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*StockVolumeRankResponse); i {
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
		file_forwarder_realtime_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*StockVolumeRankMessage); i {
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
			RawDescriptor: file_forwarder_realtime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_forwarder_realtime_proto_goTypes,
		DependencyIndexes: file_forwarder_realtime_proto_depIdxs,
		MessageInfos:      file_forwarder_realtime_proto_msgTypes,
	}.Build()
	File_forwarder_realtime_proto = out.File
	file_forwarder_realtime_proto_rawDesc = nil
	file_forwarder_realtime_proto_goTypes = nil
	file_forwarder_realtime_proto_depIdxs = nil
}
