// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/display_upload_product_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enumerates display upload product types.
type DisplayUploadProductTypeEnum_DisplayUploadProductType int32

const (
	// Not specified.
	DisplayUploadProductTypeEnum_UNSPECIFIED DisplayUploadProductTypeEnum_DisplayUploadProductType = 0
	// The value is unknown in this version.
	DisplayUploadProductTypeEnum_UNKNOWN DisplayUploadProductTypeEnum_DisplayUploadProductType = 1
	// HTML5 upload ad. This product type requires the upload_media_bundle
	// field in DisplayUploadAdInfo to be set.
	DisplayUploadProductTypeEnum_HTML5_UPLOAD_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 2
	// Dynamic HTML5 education ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in an education campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_EDUCATION_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 3
	// Dynamic HTML5 flight ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a flight campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_FLIGHT_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 4
	// Dynamic HTML5 hotel and rental ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a hotel campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_HOTEL_RENTAL_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 5
	// Dynamic HTML5 job ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a job campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_JOB_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 6
	// Dynamic HTML5 local ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a local campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_LOCAL_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 7
	// Dynamic HTML5 real estate ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a real estate campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_REAL_ESTATE_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 8
	// Dynamic HTML5 custom ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a custom campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_CUSTOM_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 9
	// Dynamic HTML5 travel ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a travel campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_TRAVEL_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 10
	// Dynamic HTML5 hotel ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a hotel campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_HOTEL_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 11
)

var DisplayUploadProductTypeEnum_DisplayUploadProductType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "HTML5_UPLOAD_AD",
	3:  "DYNAMIC_HTML5_EDUCATION_AD",
	4:  "DYNAMIC_HTML5_FLIGHT_AD",
	5:  "DYNAMIC_HTML5_HOTEL_RENTAL_AD",
	6:  "DYNAMIC_HTML5_JOB_AD",
	7:  "DYNAMIC_HTML5_LOCAL_AD",
	8:  "DYNAMIC_HTML5_REAL_ESTATE_AD",
	9:  "DYNAMIC_HTML5_CUSTOM_AD",
	10: "DYNAMIC_HTML5_TRAVEL_AD",
	11: "DYNAMIC_HTML5_HOTEL_AD",
}

var DisplayUploadProductTypeEnum_DisplayUploadProductType_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"HTML5_UPLOAD_AD":               2,
	"DYNAMIC_HTML5_EDUCATION_AD":    3,
	"DYNAMIC_HTML5_FLIGHT_AD":       4,
	"DYNAMIC_HTML5_HOTEL_RENTAL_AD": 5,
	"DYNAMIC_HTML5_JOB_AD":          6,
	"DYNAMIC_HTML5_LOCAL_AD":        7,
	"DYNAMIC_HTML5_REAL_ESTATE_AD":  8,
	"DYNAMIC_HTML5_CUSTOM_AD":       9,
	"DYNAMIC_HTML5_TRAVEL_AD":       10,
	"DYNAMIC_HTML5_HOTEL_AD":        11,
}

func (x DisplayUploadProductTypeEnum_DisplayUploadProductType) String() string {
	return proto.EnumName(DisplayUploadProductTypeEnum_DisplayUploadProductType_name, int32(x))
}

func (DisplayUploadProductTypeEnum_DisplayUploadProductType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_93737b3f73f1d8eb, []int{0, 0}
}

// Container for display upload product types. Product types that have the word
// "DYNAMIC" in them must be associated with a campaign that has a dynamic
// remarketing feed. See https://support.google.com/google-ads/answer/6053288
// for more info about dynamic remarketing. Other product types are regarded
// as "static" and do not have this requirement.
type DisplayUploadProductTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisplayUploadProductTypeEnum) Reset()         { *m = DisplayUploadProductTypeEnum{} }
func (m *DisplayUploadProductTypeEnum) String() string { return proto.CompactTextString(m) }
func (*DisplayUploadProductTypeEnum) ProtoMessage()    {}
func (*DisplayUploadProductTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_93737b3f73f1d8eb, []int{0}
}

func (m *DisplayUploadProductTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisplayUploadProductTypeEnum.Unmarshal(m, b)
}
func (m *DisplayUploadProductTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisplayUploadProductTypeEnum.Marshal(b, m, deterministic)
}
func (m *DisplayUploadProductTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisplayUploadProductTypeEnum.Merge(m, src)
}
func (m *DisplayUploadProductTypeEnum) XXX_Size() int {
	return xxx_messageInfo_DisplayUploadProductTypeEnum.Size(m)
}
func (m *DisplayUploadProductTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DisplayUploadProductTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DisplayUploadProductTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.DisplayUploadProductTypeEnum_DisplayUploadProductType", DisplayUploadProductTypeEnum_DisplayUploadProductType_name, DisplayUploadProductTypeEnum_DisplayUploadProductType_value)
	proto.RegisterType((*DisplayUploadProductTypeEnum)(nil), "google.ads.googleads.v2.enums.DisplayUploadProductTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/display_upload_product_type.proto", fileDescriptor_93737b3f73f1d8eb)
}

var fileDescriptor_93737b3f73f1d8eb = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x6d, 0xae, 0xde, 0xab, 0xd3, 0xc5, 0x2d, 0xa3, 0xe8, 0xa5, 0xb6, 0x62, 0xef, 0x03,
	0x4c, 0x20, 0xe2, 0x26, 0x2e, 0x64, 0x9a, 0xcc, 0x6d, 0xa3, 0x69, 0x12, 0xda, 0x49, 0x45, 0x29,
	0x84, 0x78, 0x27, 0x84, 0x42, 0x9b, 0x19, 0x3a, 0x69, 0xa1, 0xaf, 0xe3, 0xd2, 0x47, 0xf1, 0x15,
	0x7c, 0x03, 0x97, 0xae, 0x5d, 0xc8, 0xcc, 0xd8, 0x42, 0x28, 0x75, 0x33, 0x1c, 0xe6, 0xfb, 0xcf,
	0x7f, 0x0e, 0xfc, 0x07, 0xbc, 0x2f, 0x39, 0x2f, 0x57, 0x85, 0x9d, 0x33, 0x69, 0x9b, 0x52, 0x55,
	0x3b, 0xc7, 0x2e, 0xaa, 0xed, 0x5a, 0xda, 0x6c, 0x29, 0xc5, 0x2a, 0xdf, 0x67, 0x5b, 0xb1, 0xe2,
	0x39, 0xcb, 0xc4, 0x86, 0xb3, 0xed, 0x7d, 0x9d, 0xd5, 0x7b, 0x51, 0x20, 0xb1, 0xe1, 0x35, 0x87,
	0x7d, 0xd3, 0x85, 0x72, 0x26, 0xd1, 0xd1, 0x00, 0xed, 0x1c, 0xa4, 0x0d, 0xba, 0xbd, 0x83, 0xbf,
	0x58, 0xda, 0x79, 0x55, 0xf1, 0x3a, 0xaf, 0x97, 0xbc, 0x92, 0xa6, 0xf9, 0xf6, 0x8f, 0x05, 0x7a,
	0xbe, 0x19, 0x91, 0xea, 0x09, 0x89, 0x19, 0x40, 0xf7, 0xa2, 0x20, 0xd5, 0x76, 0x7d, 0xfb, 0xd3,
	0x02, 0x37, 0xe7, 0x04, 0xf0, 0x1a, 0xb4, 0xd3, 0x68, 0x96, 0x10, 0x2f, 0xb8, 0x0b, 0x88, 0xdf,
	0x79, 0x00, 0xdb, 0xe0, 0x2a, 0x8d, 0x3e, 0x46, 0xf1, 0xa7, 0xa8, 0xd3, 0x82, 0x4f, 0xc1, 0xf5,
	0x98, 0x4e, 0xc2, 0xb7, 0x59, 0x9a, 0x84, 0x31, 0xf6, 0x33, 0xec, 0x77, 0x2c, 0xf8, 0x0a, 0x74,
	0xfd, 0xcf, 0x11, 0x9e, 0x04, 0x5e, 0x66, 0x20, 0xf1, 0x53, 0x0f, 0xd3, 0x20, 0x8e, 0x14, 0xbf,
	0x80, 0x2f, 0xc1, 0x8b, 0x26, 0xbf, 0x0b, 0x83, 0xd1, 0x98, 0x2a, 0xf8, 0x10, 0x0e, 0x40, 0xbf,
	0x09, 0xc7, 0x31, 0x25, 0x61, 0x36, 0x25, 0x11, 0xc5, 0xa1, 0x92, 0x3c, 0x82, 0x37, 0xe0, 0x59,
	0x53, 0xf2, 0x21, 0x1e, 0x2a, 0x72, 0x09, 0xbb, 0xe0, 0x79, 0x93, 0x84, 0xb1, 0x67, 0xba, 0xae,
	0xe0, 0x6b, 0xd0, 0x6b, 0xb2, 0x29, 0xc1, 0x61, 0x46, 0x66, 0x14, 0x53, 0xa2, 0x14, 0x8f, 0x4f,
	0xf7, 0xf2, 0xd2, 0x19, 0x8d, 0x27, 0x0a, 0x3e, 0x39, 0x85, 0x74, 0x8a, 0xe7, 0x44, 0x7b, 0x83,
	0xd3, 0xb9, 0x66, 0x69, 0xec, 0x77, 0xda, 0xc3, 0xdf, 0x2d, 0x30, 0xb8, 0xe7, 0x6b, 0xf4, 0xdf,
	0x08, 0x87, 0xfd, 0x73, 0x01, 0x24, 0x2a, 0xc3, 0xa4, 0xf5, 0x65, 0xf8, 0xaf, 0xbf, 0xe4, 0xab,
	0xbc, 0x2a, 0x11, 0xdf, 0x94, 0x76, 0x59, 0x54, 0x3a, 0xe1, 0xc3, 0x4d, 0x89, 0xa5, 0x3c, 0x73,
	0x62, 0xef, 0xf4, 0xfb, 0xcd, 0xba, 0x18, 0x61, 0xfc, 0xdd, 0xea, 0x8f, 0x8c, 0x15, 0x66, 0x12,
	0x99, 0x52, 0x55, 0x73, 0x07, 0xa9, 0x6b, 0x90, 0x3f, 0x0e, 0x7c, 0x81, 0x99, 0x5c, 0x1c, 0xf9,
	0x62, 0xee, 0x2c, 0x34, 0xff, 0x65, 0x0d, 0xcc, 0xa7, 0xeb, 0x62, 0x26, 0x5d, 0xf7, 0xa8, 0x70,
	0xdd, 0xb9, 0xe3, 0xba, 0x5a, 0xf3, 0xf5, 0x52, 0x2f, 0xf6, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe2, 0xaf, 0x76, 0x61, 0xfa, 0x02, 0x00, 0x00,
}
