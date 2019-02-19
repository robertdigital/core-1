// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geoip.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GeoIPCountry struct {
	IsoCode string `protobuf:"bytes,1,opt,name=IsoCode" json:"IsoCode,omitempty"`
}

func (m *GeoIPCountry) Reset()                    { *m = GeoIPCountry{} }
func (m *GeoIPCountry) String() string            { return proto.CompactTextString(m) }
func (*GeoIPCountry) ProtoMessage()               {}
func (*GeoIPCountry) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *GeoIPCountry) GetIsoCode() string {
	if m != nil {
		return m.IsoCode
	}
	return ""
}

type GeoIP struct {
	Country *GeoIPCountry `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
}

func (m *GeoIP) Reset()                    { *m = GeoIP{} }
func (m *GeoIP) String() string            { return proto.CompactTextString(m) }
func (*GeoIP) ProtoMessage()               {}
func (*GeoIP) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *GeoIP) GetCountry() *GeoIPCountry {
	if m != nil {
		return m.Country
	}
	return nil
}

func init() {
	proto.RegisterType((*GeoIPCountry)(nil), "sonm.GeoIPCountry")
	proto.RegisterType((*GeoIP)(nil), "sonm.GeoIP")
}

func init() { proto.RegisterFile("geoip.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4f, 0xcd, 0xcf,
	0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xce, 0xcf, 0xcb, 0x55, 0xd2, 0xe0,
	0xe2, 0x71, 0x4f, 0xcd, 0xf7, 0x0c, 0x70, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xaa, 0x14, 0x92, 0xe0,
	0x62, 0xf7, 0x2c, 0xce, 0x77, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0x71, 0x95, 0x4c, 0xb9, 0x58, 0xc1, 0x2a, 0x85, 0x74, 0xb8, 0xd8, 0x93, 0x21, 0xaa, 0xc1, 0x4a,
	0xb8, 0x8d, 0x84, 0xf4, 0x40, 0x46, 0xe9, 0x21, 0x9b, 0x13, 0x04, 0x53, 0xe2, 0xa4, 0x12, 0xa5,
	0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x52, 0xa8, 0x9b, 0x99,
	0xaf, 0x9f, 0x9c, 0x5f, 0x94, 0xaa, 0x0f, 0x76, 0x87, 0x35, 0x48, 0x28, 0x89, 0x0d, 0xcc, 0x36,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x34, 0x2a, 0xe4, 0xee, 0xa2, 0x00, 0x00, 0x00,
}
