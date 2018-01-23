// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.ext.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WantTicketReq struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	FilmId int64 `protobuf:"varint,2,opt,name=filmId" json:"filmId,omitempty"`
}

func (m *WantTicketReq) Reset()                    { *m = WantTicketReq{} }
func (m *WantTicketReq) String() string            { return proto.CompactTextString(m) }
func (*WantTicketReq) ProtoMessage()               {}
func (*WantTicketReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *WantTicketReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *WantTicketReq) GetFilmId() int64 {
	if m != nil {
		return m.FilmId
	}
	return 0
}

type WantTicketRsp struct {
}

func (m *WantTicketRsp) Reset()                    { *m = WantTicketRsp{} }
func (m *WantTicketRsp) String() string            { return proto.CompactTextString(m) }
func (*WantTicketRsp) ProtoMessage()               {}
func (*WantTicketRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type TicketReq struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	FilmId int64 `protobuf:"varint,2,opt,name=filmId" json:"filmId,omitempty"`
	MhId   int64 `protobuf:"varint,3,opt,name=mhId" json:"mhId,omitempty"`
	SitId  int64 `protobuf:"varint,4,opt,name=sitId" json:"sitId,omitempty"`
}

func (m *TicketReq) Reset()                    { *m = TicketReq{} }
func (m *TicketReq) String() string            { return proto.CompactTextString(m) }
func (*TicketReq) ProtoMessage()               {}
func (*TicketReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *TicketReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TicketReq) GetFilmId() int64 {
	if m != nil {
		return m.FilmId
	}
	return 0
}

func (m *TicketReq) GetMhId() int64 {
	if m != nil {
		return m.MhId
	}
	return 0
}

func (m *TicketReq) GetSitId() int64 {
	if m != nil {
		return m.SitId
	}
	return 0
}

type TicketRsp struct {
}

func (m *TicketRsp) Reset()                    { *m = TicketRsp{} }
func (m *TicketRsp) String() string            { return proto.CompactTextString(m) }
func (*TicketRsp) ProtoMessage()               {}
func (*TicketRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type Order struct {
	OrderNum   string `protobuf:"bytes,1,opt,name=orderNum" json:"orderNum,omitempty"`
	OrderPrice string `protobuf:"bytes,2,opt,name=orderPrice" json:"orderPrice,omitempty"`
	CreateAt   string `protobuf:"bytes,3,opt,name=createAt" json:"createAt,omitempty"`
	PayAt      string `protobuf:"bytes,4,opt,name=payAt" json:"payAt,omitempty"`
	MhId       int64  `protobuf:"varint,5,opt,name=mhId" json:"mhId,omitempty"`
	UserId     int64  `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	MovieId    int64  `protobuf:"varint,7,opt,name=movieId" json:"movieId,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Order) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *Order) GetOrderPrice() string {
	if m != nil {
		return m.OrderPrice
	}
	return ""
}

func (m *Order) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *Order) GetPayAt() string {
	if m != nil {
		return m.PayAt
	}
	return ""
}

func (m *Order) GetMhId() int64 {
	if m != nil {
		return m.MhId
	}
	return 0
}

func (m *Order) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

type LookOrdersReq struct {
}

func (m *LookOrdersReq) Reset()                    { *m = LookOrdersReq{} }
func (m *LookOrdersReq) String() string            { return proto.CompactTextString(m) }
func (*LookOrdersReq) ProtoMessage()               {}
func (*LookOrdersReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type LookOrdersRsp struct {
	Order []*Order `protobuf:"bytes,1,rep,name=order" json:"order,omitempty"`
}

func (m *LookOrdersRsp) Reset()                    { *m = LookOrdersRsp{} }
func (m *LookOrdersRsp) String() string            { return proto.CompactTextString(m) }
func (*LookOrdersRsp) ProtoMessage()               {}
func (*LookOrdersRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *LookOrdersRsp) GetOrder() []*Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type PayOrderReq struct {
	OrderId int64 `protobuf:"varint,1,opt,name=orderId" json:"orderId,omitempty"`
}

func (m *PayOrderReq) Reset()                    { *m = PayOrderReq{} }
func (m *PayOrderReq) String() string            { return proto.CompactTextString(m) }
func (*PayOrderReq) ProtoMessage()               {}
func (*PayOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *PayOrderReq) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type PayOrderRsp struct {
}

func (m *PayOrderRsp) Reset()                    { *m = PayOrderRsp{} }
func (m *PayOrderRsp) String() string            { return proto.CompactTextString(m) }
func (*PayOrderRsp) ProtoMessage()               {}
func (*PayOrderRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type UndoOrderReq struct {
	OrderId int64 `protobuf:"varint,1,opt,name=orderId" json:"orderId,omitempty"`
}

func (m *UndoOrderReq) Reset()                    { *m = UndoOrderReq{} }
func (m *UndoOrderReq) String() string            { return proto.CompactTextString(m) }
func (*UndoOrderReq) ProtoMessage()               {}
func (*UndoOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *UndoOrderReq) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type UndoOrderRsp struct {
}

func (m *UndoOrderRsp) Reset()                    { *m = UndoOrderRsp{} }
func (m *UndoOrderRsp) String() string            { return proto.CompactTextString(m) }
func (*UndoOrderRsp) ProtoMessage()               {}
func (*UndoOrderRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func init() {
	proto.RegisterType((*WantTicketReq)(nil), "pb.WantTicketReq")
	proto.RegisterType((*WantTicketRsp)(nil), "pb.WantTicketRsp")
	proto.RegisterType((*TicketReq)(nil), "pb.TicketReq")
	proto.RegisterType((*TicketRsp)(nil), "pb.TicketRsp")
	proto.RegisterType((*Order)(nil), "pb.Order")
	proto.RegisterType((*LookOrdersReq)(nil), "pb.LookOrdersReq")
	proto.RegisterType((*LookOrdersRsp)(nil), "pb.LookOrdersRsp")
	proto.RegisterType((*PayOrderReq)(nil), "pb.PayOrderReq")
	proto.RegisterType((*PayOrderRsp)(nil), "pb.PayOrderRsp")
	proto.RegisterType((*UndoOrderReq)(nil), "pb.UndoOrderReq")
	proto.RegisterType((*UndoOrderRsp)(nil), "pb.UndoOrderRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for OrderServiceExt service

type OrderServiceExtClient interface {
	// 想看
	WantTicket(ctx context.Context, in *WantTicketReq, opts ...client.CallOption) (*WantTicketRsp, error)
	// 下单
	Ticket(ctx context.Context, in *TicketReq, opts ...client.CallOption) (*TicketRsp, error)
	// 查看所有订单
	LookOrders(ctx context.Context, in *LookOrdersReq, opts ...client.CallOption) (*LookOrdersRsp, error)
	// 支付
	PayOrder(ctx context.Context, in *PayOrderReq, opts ...client.CallOption) (*PayOrderRsp, error)
	// 取消订单
	UndoOrder(ctx context.Context, in *UndoOrderReq, opts ...client.CallOption) (*UndoOrderRsp, error)
}

type orderServiceExtClient struct {
	c           client.Client
	serviceName string
}

func NewOrderServiceExtClient(serviceName string, c client.Client) OrderServiceExtClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &orderServiceExtClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *orderServiceExtClient) WantTicket(ctx context.Context, in *WantTicketReq, opts ...client.CallOption) (*WantTicketRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.WantTicket", in)
	out := new(WantTicketRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) Ticket(ctx context.Context, in *TicketReq, opts ...client.CallOption) (*TicketRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.Ticket", in)
	out := new(TicketRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) LookOrders(ctx context.Context, in *LookOrdersReq, opts ...client.CallOption) (*LookOrdersRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.LookOrders", in)
	out := new(LookOrdersRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) PayOrder(ctx context.Context, in *PayOrderReq, opts ...client.CallOption) (*PayOrderRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.PayOrder", in)
	out := new(PayOrderRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) UndoOrder(ctx context.Context, in *UndoOrderReq, opts ...client.CallOption) (*UndoOrderRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.UndoOrder", in)
	out := new(UndoOrderRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderServiceExt service

type OrderServiceExtHandler interface {
	// 想看
	WantTicket(context.Context, *WantTicketReq, *WantTicketRsp) error
	// 下单
	Ticket(context.Context, *TicketReq, *TicketRsp) error
	// 查看所有订单
	LookOrders(context.Context, *LookOrdersReq, *LookOrdersRsp) error
	// 支付
	PayOrder(context.Context, *PayOrderReq, *PayOrderRsp) error
	// 取消订单
	UndoOrder(context.Context, *UndoOrderReq, *UndoOrderRsp) error
}

func RegisterOrderServiceExtHandler(s server.Server, hdlr OrderServiceExtHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&OrderServiceExt{hdlr}, opts...))
}

type OrderServiceExt struct {
	OrderServiceExtHandler
}

func (h *OrderServiceExt) WantTicket(ctx context.Context, in *WantTicketReq, out *WantTicketRsp) error {
	return h.OrderServiceExtHandler.WantTicket(ctx, in, out)
}

func (h *OrderServiceExt) Ticket(ctx context.Context, in *TicketReq, out *TicketRsp) error {
	return h.OrderServiceExtHandler.Ticket(ctx, in, out)
}

func (h *OrderServiceExt) LookOrders(ctx context.Context, in *LookOrdersReq, out *LookOrdersRsp) error {
	return h.OrderServiceExtHandler.LookOrders(ctx, in, out)
}

func (h *OrderServiceExt) PayOrder(ctx context.Context, in *PayOrderReq, out *PayOrderRsp) error {
	return h.OrderServiceExtHandler.PayOrder(ctx, in, out)
}

func (h *OrderServiceExt) UndoOrder(ctx context.Context, in *UndoOrderReq, out *UndoOrderRsp) error {
	return h.OrderServiceExtHandler.UndoOrder(ctx, in, out)
}

func init() { proto.RegisterFile("order.ext.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xe2, 0x30,
	0x10, 0xc6, 0x37, 0x40, 0x02, 0x19, 0x96, 0xcd, 0xae, 0xb5, 0x5a, 0x45, 0x39, 0x6c, 0x91, 0x2f,
	0xcd, 0x29, 0x6a, 0x29, 0xf7, 0x8a, 0x43, 0x0f, 0x91, 0xaa, 0x16, 0xb9, 0xad, 0x7a, 0x0e, 0x89,
	0xab, 0x46, 0x34, 0xc4, 0xc4, 0x06, 0xc1, 0x2b, 0xf4, 0x95, 0xfa, 0x72, 0x95, 0x27, 0x90, 0x3f,
	0x9c, 0xaa, 0xde, 0xf2, 0xfd, 0x3c, 0xe3, 0xf9, 0xbe, 0x91, 0x03, 0x4e, 0x5e, 0x24, 0xbc, 0x08,
	0xf8, 0x4e, 0x05, 0xa2, 0xc8, 0x55, 0x4e, 0x3a, 0x62, 0x41, 0xaf, 0x61, 0xf4, 0x1c, 0xad, 0xd4,
	0x63, 0x1a, 0x2f, 0xb9, 0x62, 0x7c, 0x4d, 0xfe, 0x81, 0xb5, 0x91, 0xbc, 0x08, 0x13, 0xd7, 0x18,
	0x1b, 0x7e, 0x97, 0x1d, 0x94, 0xe6, 0x2f, 0xe9, 0x5b, 0x16, 0x26, 0x6e, 0xa7, 0xe4, 0xa5, 0xa2,
	0x4e, 0xeb, 0x02, 0x29, 0x28, 0x07, 0xfb, 0xdb, 0xb7, 0x11, 0x02, 0xbd, 0xec, 0x35, 0x4c, 0xdc,
	0x2e, 0x52, 0xfc, 0x26, 0x7f, 0xc1, 0x94, 0xa9, 0x0a, 0x13, 0xb7, 0x87, 0xb0, 0x14, 0x74, 0x58,
	0x8d, 0x91, 0x82, 0x7e, 0x18, 0x60, 0xde, 0xeb, 0x74, 0xc4, 0x83, 0x01, 0xc6, 0xbc, 0xdb, 0x64,
	0x38, 0xd2, 0x66, 0x95, 0x26, 0xff, 0x01, 0xf0, 0x7b, 0x5e, 0xa4, 0x31, 0xc7, 0xc1, 0x36, 0x6b,
	0x10, 0xdd, 0x1b, 0x17, 0x3c, 0x52, 0x7c, 0xa6, 0xd0, 0x80, 0xcd, 0x2a, 0xad, 0x4d, 0x88, 0x68,
	0x3f, 0x53, 0x68, 0xc2, 0x66, 0xa5, 0xa8, 0xec, 0x9a, 0x0d, 0xbb, 0x75, 0x64, 0xab, 0x15, 0xd9,
	0x85, 0x7e, 0x96, 0x6f, 0x53, 0x1e, 0x26, 0x6e, 0x1f, 0x0f, 0x8e, 0x52, 0xaf, 0xf0, 0x36, 0xcf,
	0x97, 0x18, 0x40, 0x32, 0xbe, 0xa6, 0x17, 0x2d, 0x20, 0x05, 0x39, 0x03, 0x13, 0x7d, 0xba, 0xc6,
	0xb8, 0xeb, 0x0f, 0x27, 0x76, 0x20, 0x16, 0x01, 0x9e, 0xb2, 0x92, 0xd3, 0x73, 0x18, 0xce, 0xa3,
	0x7d, 0x89, 0xf8, 0x5a, 0xcf, 0x42, 0x5e, 0xed, 0xfd, 0x28, 0xe9, 0xa8, 0x51, 0x28, 0x05, 0xf5,
	0xe1, 0xe7, 0xd3, 0x2a, 0xc9, 0xbf, 0xd0, 0xf8, 0xab, 0x59, 0x29, 0xc5, 0xe4, 0xbd, 0x03, 0x0e,
	0x8a, 0x07, 0x5e, 0x6c, 0xd3, 0x98, 0xdf, 0xec, 0x14, 0x99, 0x02, 0xd4, 0x6f, 0x81, 0xfc, 0xd1,
	0x2e, 0x5b, 0x8f, 0xcb, 0x3b, 0x45, 0x52, 0xd0, 0x1f, 0xc4, 0x07, 0xeb, 0xd0, 0x31, 0xd2, 0xc7,
	0x75, 0x75, 0x53, 0x62, 0xe5, 0x14, 0xa0, 0xde, 0x4b, 0x79, 0x7f, 0x6b, 0x71, 0xde, 0x29, 0xc2,
	0xae, 0x00, 0x06, 0xc7, 0xc8, 0xc4, 0xd1, 0x05, 0x8d, 0x4d, 0x79, 0x6d, 0x80, 0xf5, 0x97, 0x60,
	0x57, 0x49, 0xc9, 0x6f, 0x7d, 0xde, 0x5c, 0x91, 0x77, 0x42, 0x74, 0xcb, 0xc2, 0xc2, 0x1f, 0xea,
	0xea, 0x33, 0x00, 0x00, 0xff, 0xff, 0x90, 0x99, 0x7b, 0x2b, 0x63, 0x03, 0x00, 0x00,
}
