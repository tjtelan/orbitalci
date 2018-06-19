// Code generated by protoc-gen-go. DO NOT EDIT.
// source: slack.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WebhookMsg struct {
	// required
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	// if not entered, will be default channel of url
	Channel string `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
	// if not entered, will be user that created webhook
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	// obv not required
	IconUrl string `protobuf:"bytes,4,opt,name=icon_url,json=iconUrl" json:"icon_url,omitempty"`
	// obv not required
	IconEmoji string `protobuf:"bytes,5,opt,name=icon_emoji,json=iconEmoji" json:"icon_emoji,omitempty"`
	// if you wanna get fancy
	Attachments          []*Attachment `protobuf:"bytes,6,rep,name=attachments" json:"attachments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *WebhookMsg) Reset()         { *m = WebhookMsg{} }
func (m *WebhookMsg) String() string { return proto.CompactTextString(m) }
func (*WebhookMsg) ProtoMessage()    {}
func (*WebhookMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_slack_f60cdea5ddeeb863, []int{0}
}
func (m *WebhookMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookMsg.Unmarshal(m, b)
}
func (m *WebhookMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookMsg.Marshal(b, m, deterministic)
}
func (dst *WebhookMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookMsg.Merge(dst, src)
}
func (m *WebhookMsg) XXX_Size() int {
	return xxx_messageInfo_WebhookMsg.Size(m)
}
func (m *WebhookMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookMsg.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookMsg proto.InternalMessageInfo

func (m *WebhookMsg) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *WebhookMsg) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *WebhookMsg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *WebhookMsg) GetIconUrl() string {
	if m != nil {
		return m.IconUrl
	}
	return ""
}

func (m *WebhookMsg) GetIconEmoji() string {
	if m != nil {
		return m.IconEmoji
	}
	return ""
}

func (m *WebhookMsg) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

type Field struct {
	Title                string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Short                bool     `protobuf:"varint,3,opt,name=short" json:"short,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_slack_f60cdea5ddeeb863, []int{1}
}
func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (dst *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(dst, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Field) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Field) GetShort() bool {
	if m != nil {
		return m.Short
	}
	return false
}

type Attachment struct {
	Fallback             string   `protobuf:"bytes,1,opt,name=fallback" json:"fallback,omitempty"`
	Color                string   `protobuf:"bytes,2,opt,name=color" json:"color,omitempty"`
	Pretext              string   `protobuf:"bytes,3,opt,name=pretext" json:"pretext,omitempty"`
	AuthorName           string   `protobuf:"bytes,4,opt,name=author_name,json=authorName" json:"author_name,omitempty"`
	AuthorLink           string   `protobuf:"bytes,5,opt,name=author_link,json=authorLink" json:"author_link,omitempty"`
	AuthorIcon           string   `protobuf:"bytes,6,opt,name=author_icon,json=authorIcon" json:"author_icon,omitempty"`
	Title                string   `protobuf:"bytes,7,opt,name=title" json:"title,omitempty"`
	TitleLink            string   `protobuf:"bytes,8,opt,name=title_link,json=titleLink" json:"title_link,omitempty"`
	Text                 string   `protobuf:"bytes,9,opt,name=text" json:"text,omitempty"`
	Fields               []*Field `protobuf:"bytes,10,rep,name=fields" json:"fields,omitempty"`
	ImageUrl             string   `protobuf:"bytes,11,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
	ThumbUrl             string   `protobuf:"bytes,12,opt,name=thumb_url,json=thumbUrl" json:"thumb_url,omitempty"`
	Footer               string   `protobuf:"bytes,13,opt,name=footer" json:"footer,omitempty"`
	FooterIcon           string   `protobuf:"bytes,14,opt,name=footer_icon,json=footerIcon" json:"footer_icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attachment) Reset()         { *m = Attachment{} }
func (m *Attachment) String() string { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()    {}
func (*Attachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_slack_f60cdea5ddeeb863, []int{2}
}
func (m *Attachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attachment.Unmarshal(m, b)
}
func (m *Attachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attachment.Marshal(b, m, deterministic)
}
func (dst *Attachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attachment.Merge(dst, src)
}
func (m *Attachment) XXX_Size() int {
	return xxx_messageInfo_Attachment.Size(m)
}
func (m *Attachment) XXX_DiscardUnknown() {
	xxx_messageInfo_Attachment.DiscardUnknown(m)
}

var xxx_messageInfo_Attachment proto.InternalMessageInfo

func (m *Attachment) GetFallback() string {
	if m != nil {
		return m.Fallback
	}
	return ""
}

func (m *Attachment) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Attachment) GetPretext() string {
	if m != nil {
		return m.Pretext
	}
	return ""
}

func (m *Attachment) GetAuthorName() string {
	if m != nil {
		return m.AuthorName
	}
	return ""
}

func (m *Attachment) GetAuthorLink() string {
	if m != nil {
		return m.AuthorLink
	}
	return ""
}

func (m *Attachment) GetAuthorIcon() string {
	if m != nil {
		return m.AuthorIcon
	}
	return ""
}

func (m *Attachment) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Attachment) GetTitleLink() string {
	if m != nil {
		return m.TitleLink
	}
	return ""
}

func (m *Attachment) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Attachment) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Attachment) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *Attachment) GetThumbUrl() string {
	if m != nil {
		return m.ThumbUrl
	}
	return ""
}

func (m *Attachment) GetFooter() string {
	if m != nil {
		return m.Footer
	}
	return ""
}

func (m *Attachment) GetFooterIcon() string {
	if m != nil {
		return m.FooterIcon
	}
	return ""
}

func init() {
	proto.RegisterType((*WebhookMsg)(nil), "slack.WebhookMsg")
	proto.RegisterType((*Field)(nil), "slack.Field")
	proto.RegisterType((*Attachment)(nil), "slack.Attachment")
}

func init() { proto.RegisterFile("slack.proto", fileDescriptor_slack_f60cdea5ddeeb863) }

var fileDescriptor_slack_f60cdea5ddeeb863 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x52, 0x4d, 0xab, 0xdb, 0x30,
	0x10, 0x24, 0x1f, 0x76, 0xec, 0x75, 0x5a, 0xa8, 0x28, 0x45, 0x6d, 0x08, 0x0d, 0xa1, 0x87, 0x9c,
	0x72, 0x68, 0x7e, 0x41, 0x0b, 0x2d, 0x04, 0xda, 0x1e, 0x0c, 0xa5, 0xd0, 0x4b, 0x90, 0x5d, 0x25,
	0x76, 0x2d, 0x4b, 0x41, 0x96, 0x1f, 0xef, 0xcf, 0xbc, 0x9f, 0xf3, 0xfe, 0xd7, 0x43, 0xbb, 0x72,
	0xe2, 0xdb, 0xce, 0xcc, 0x66, 0xa3, 0x99, 0x31, 0x64, 0x9d, 0x12, 0x65, 0xb3, 0xbf, 0x5a, 0xe3,
	0x0c, 0x8b, 0x10, 0x6c, 0x9f, 0x27, 0x00, 0x7f, 0x64, 0x51, 0x19, 0xd3, 0xfc, 0xec, 0x2e, 0x8c,
	0xc1, 0xdc, 0xc9, 0x47, 0xc7, 0x27, 0x9b, 0xc9, 0x2e, 0xcd, 0x71, 0x66, 0x1c, 0x16, 0x65, 0x25,
	0xb4, 0x96, 0x8a, 0x4f, 0x91, 0x1e, 0x20, 0xfb, 0x00, 0x49, 0xdf, 0x49, 0xab, 0x45, 0x2b, 0xf9,
	0x0c, 0xa5, 0x1b, 0x66, 0xef, 0x21, 0xa9, 0x4b, 0xa3, 0x4f, 0xbd, 0x55, 0x7c, 0x4e, 0x3f, 0xf3,
	0xf8, 0xb7, 0x55, 0x6c, 0x0d, 0x80, 0x92, 0x6c, 0xcd, 0xff, 0x9a, 0x47, 0x28, 0xa6, 0x9e, 0xf9,
	0xe6, 0x09, 0x76, 0x80, 0x4c, 0x38, 0x27, 0xca, 0xaa, 0x95, 0xda, 0x75, 0x3c, 0xde, 0xcc, 0x76,
	0xd9, 0xe7, 0x37, 0x7b, 0x7a, 0xfc, 0x97, 0x9b, 0x92, 0x8f, 0xb7, 0xb6, 0x47, 0x88, 0xbe, 0xd7,
	0x52, 0xfd, 0x63, 0x6f, 0x21, 0x72, 0xb5, 0x53, 0x32, 0x58, 0x20, 0xe0, 0xd9, 0x07, 0xa1, 0x7a,
	0x19, 0x1c, 0x10, 0xf0, 0x6c, 0x57, 0x19, 0xeb, 0xf0, 0xf1, 0x49, 0x4e, 0x60, 0xfb, 0x34, 0x03,
	0xb8, 0xff, 0x8d, 0x37, 0x79, 0x16, 0x4a, 0x15, 0xa2, 0x6c, 0xc2, 0xcd, 0x1b, 0xf6, 0x07, 0x4a,
	0xa3, 0x8c, 0x1d, 0xce, 0x22, 0xf0, 0x81, 0x5d, 0xad, 0xc4, 0x1c, 0x29, 0x95, 0x01, 0xb2, 0x8f,
	0x90, 0x89, 0xde, 0x55, 0xc6, 0x9e, 0x30, 0x33, 0xca, 0x05, 0x88, 0xfa, 0xe5, 0x53, 0xbb, 0x2f,
	0xa8, 0x5a, 0x37, 0x21, 0x9b, 0xb0, 0xf0, 0xa3, 0xd6, 0xcd, 0x68, 0xc1, 0x07, 0xc6, 0xe3, 0xf1,
	0xc2, 0xb1, 0x34, 0xfa, 0xee, 0x7f, 0x31, 0xf6, 0xbf, 0x06, 0xc0, 0x81, 0xce, 0x26, 0x14, 0x39,
	0x32, 0x78, 0x75, 0xa8, 0x3d, 0x1d, 0xd5, 0xfe, 0x09, 0xe2, 0xb3, 0x4f, 0xb4, 0xe3, 0x80, 0x0d,
	0x2c, 0x43, 0x03, 0x18, 0x73, 0x1e, 0x34, 0xb6, 0x82, 0xb4, 0x6e, 0xc5, 0x45, 0x62, 0xcf, 0x19,
	0xc5, 0x83, 0x84, 0x2f, 0x7a, 0x05, 0xa9, 0xab, 0xfa, 0xb6, 0x40, 0x71, 0x49, 0x22, 0x12, 0x5e,
	0x7c, 0x07, 0xf1, 0xd9, 0x18, 0x27, 0x2d, 0x7f, 0x85, 0x4a, 0x40, 0xde, 0x21, 0x4d, 0xe4, 0xf0,
	0x35, 0x39, 0x24, 0xca, 0x3b, 0xfc, 0x3a, 0xff, 0x3b, 0xbd, 0x16, 0x45, 0x8c, 0x9f, 0xf1, 0xe1,
	0x25, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x18, 0xba, 0x94, 0xd5, 0x02, 0x00, 0x00,
}