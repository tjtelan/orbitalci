// Code generated by protoc-gen-go. DO NOT EDIT.
// source: respositories.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PaginatedRepository struct {
	Pagelen float64                                 `protobuf:"fixed64,1,opt,name=pagelen" json:"pagelen,omitempty"`
	Size    float64                                 `protobuf:"fixed64,2,opt,name=size" json:"size,omitempty"`
	Values  []*PaginatedRepository_RepositoryValues `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	Page    float64                                 `protobuf:"fixed64,4,opt,name=page" json:"page,omitempty"`
	Next    string                                  `protobuf:"bytes,5,opt,name=next" json:"next,omitempty"`
}

func (m *PaginatedRepository) Reset()                    { *m = PaginatedRepository{} }
func (m *PaginatedRepository) String() string            { return proto.CompactTextString(m) }
func (*PaginatedRepository) ProtoMessage()               {}
func (*PaginatedRepository) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *PaginatedRepository) GetPagelen() float64 {
	if m != nil {
		return m.Pagelen
	}
	return 0
}

func (m *PaginatedRepository) GetSize() float64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PaginatedRepository) GetValues() []*PaginatedRepository_RepositoryValues {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *PaginatedRepository) GetPage() float64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PaginatedRepository) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

type PaginatedRepository_RepositoryValues struct {
	Name        string                                                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Links       *PaginatedRepository_RepositoryValues_RepositoryLinks `protobuf:"bytes,2,opt,name=links" json:"links,omitempty"`
	Project     *PaginatedRepository_RepositoryValues_Project         `protobuf:"bytes,3,opt,name=project" json:"project,omitempty"`
	CreatedOn   *google_protobuf.Timestamp                            `protobuf:"bytes,4,opt,name=created_on,json=createdOn" json:"created_on,omitempty"`
	Mainbranch  *PaginatedRepository_RepositoryValues_MainBranch      `protobuf:"bytes,5,opt,name=mainbranch" json:"mainbranch,omitempty"`
	FullName    string                                                `protobuf:"bytes,6,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	UpdatedOn   *google_protobuf.Timestamp                            `protobuf:"bytes,7,opt,name=updated_on,json=updatedOn" json:"updated_on,omitempty"`
	Size        float64                                               `protobuf:"fixed64,8,opt,name=size" json:"size,omitempty"`
	Type        string                                                `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
	Slug        string                                                `protobuf:"bytes,10,opt,name=slug" json:"slug,omitempty"`
	IsPrivate   bool                                                  `protobuf:"varint,11,opt,name=is_private,json=isPrivate" json:"is_private,omitempty"`
	Description string                                                `protobuf:"bytes,12,opt,name=description" json:"description,omitempty"`
}

func (m *PaginatedRepository_RepositoryValues) Reset()         { *m = PaginatedRepository_RepositoryValues{} }
func (m *PaginatedRepository_RepositoryValues) String() string { return proto.CompactTextString(m) }
func (*PaginatedRepository_RepositoryValues) ProtoMessage()    {}
func (*PaginatedRepository_RepositoryValues) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0}
}

func (m *PaginatedRepository_RepositoryValues) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues) GetLinks() *PaginatedRepository_RepositoryValues_RepositoryLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetProject() *PaginatedRepository_RepositoryValues_Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetCreatedOn() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedOn
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetMainbranch() *PaginatedRepository_RepositoryValues_MainBranch {
	if m != nil {
		return m.Mainbranch
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues) GetUpdatedOn() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedOn
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetSize() float64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PaginatedRepository_RepositoryValues) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func (m *PaginatedRepository_RepositoryValues) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type PaginatedRepository_RepositoryValues_RepositoryLinks struct {
	Watchers     *LinkUrl       `protobuf:"bytes,1,opt,name=watchers" json:"watchers,omitempty"`
	Branches     *LinkUrl       `protobuf:"bytes,2,opt,name=branches" json:"branches,omitempty"`
	Tags         *LinkUrl       `protobuf:"bytes,3,opt,name=tags" json:"tags,omitempty"`
	Commits      *LinkUrl       `protobuf:"bytes,4,opt,name=commits" json:"commits,omitempty"`
	Clone        []*LinkAndName `protobuf:"bytes,5,rep,name=clone" json:"clone,omitempty"`
	Self         *LinkUrl       `protobuf:"bytes,6,opt,name=self" json:"self,omitempty"`
	Source       *LinkUrl       `protobuf:"bytes,7,opt,name=source" json:"source,omitempty"`
	Hooks        *LinkUrl       `protobuf:"bytes,8,opt,name=hooks" json:"hooks,omitempty"`
	Forks        *LinkUrl       `protobuf:"bytes,9,opt,name=forks" json:"forks,omitempty"`
	Downloads    *LinkUrl       `protobuf:"bytes,10,opt,name=downloads" json:"downloads,omitempty"`
	Pullrequests *LinkUrl       `protobuf:"bytes,11,opt,name=pullrequests" json:"pullrequests,omitempty"`
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) Reset() {
	*m = PaginatedRepository_RepositoryValues_RepositoryLinks{}
}
func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) String() string {
	return proto.CompactTextString(m)
}
func (*PaginatedRepository_RepositoryValues_RepositoryLinks) ProtoMessage() {}
func (*PaginatedRepository_RepositoryValues_RepositoryLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0, 0}
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetWatchers() *LinkUrl {
	if m != nil {
		return m.Watchers
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetBranches() *LinkUrl {
	if m != nil {
		return m.Branches
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetTags() *LinkUrl {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetCommits() *LinkUrl {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetClone() []*LinkAndName {
	if m != nil {
		return m.Clone
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetSelf() *LinkUrl {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetSource() *LinkUrl {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetHooks() *LinkUrl {
	if m != nil {
		return m.Hooks
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetForks() *LinkUrl {
	if m != nil {
		return m.Forks
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetDownloads() *LinkUrl {
	if m != nil {
		return m.Downloads
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetPullrequests() *LinkUrl {
	if m != nil {
		return m.Pullrequests
	}
	return nil
}

type PaginatedRepository_RepositoryValues_Project struct {
	Key   string                                                `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Type  string                                                `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Uuid  string                                                `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	Links *PaginatedRepository_RepositoryValues_RepositoryLinks `protobuf:"bytes,4,opt,name=links" json:"links,omitempty"`
	Name  string                                                `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
}

func (m *PaginatedRepository_RepositoryValues_Project) Reset() {
	*m = PaginatedRepository_RepositoryValues_Project{}
}
func (m *PaginatedRepository_RepositoryValues_Project) String() string {
	return proto.CompactTextString(m)
}
func (*PaginatedRepository_RepositoryValues_Project) ProtoMessage() {}
func (*PaginatedRepository_RepositoryValues_Project) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0, 1}
}

func (m *PaginatedRepository_RepositoryValues_Project) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues_Project) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues_Project) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues_Project) GetLinks() *PaginatedRepository_RepositoryValues_RepositoryLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PaginatedRepository_RepositoryValues_MainBranch struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PaginatedRepository_RepositoryValues_MainBranch) Reset() {
	*m = PaginatedRepository_RepositoryValues_MainBranch{}
}
func (m *PaginatedRepository_RepositoryValues_MainBranch) String() string {
	return proto.CompactTextString(m)
}
func (*PaginatedRepository_RepositoryValues_MainBranch) ProtoMessage() {}
func (*PaginatedRepository_RepositoryValues_MainBranch) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0, 2}
}

func (m *PaginatedRepository_RepositoryValues_MainBranch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues_MainBranch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*PaginatedRepository)(nil), "protos.PaginatedRepository")
	proto.RegisterType((*PaginatedRepository_RepositoryValues)(nil), "protos.PaginatedRepository.RepositoryValues")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_RepositoryLinks)(nil), "protos.PaginatedRepository.RepositoryValues.RepositoryLinks")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_Project)(nil), "protos.PaginatedRepository.RepositoryValues.Project")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_MainBranch)(nil), "protos.PaginatedRepository.RepositoryValues.MainBranch")
}

func init() { proto.RegisterFile("respositories.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5f, 0x6b, 0xd4, 0x40,
	0x10, 0x27, 0xbd, 0xbf, 0x99, 0x14, 0x5a, 0xb6, 0x2f, 0x21, 0x22, 0x1e, 0x8a, 0x78, 0x45, 0xbd,
	0xc2, 0xb5, 0x20, 0x05, 0x5f, 0x14, 0x1f, 0xb5, 0x2d, 0xc1, 0x3f, 0x8f, 0x65, 0x9b, 0xcc, 0x5d,
	0xd7, 0x4b, 0x76, 0x63, 0x76, 0xd3, 0x5a, 0x9f, 0xfc, 0x3e, 0xbe, 0xfa, 0x99, 0xfc, 0x1c, 0xb2,
	0xb3, 0x49, 0x73, 0xed, 0x1d, 0xc8, 0x81, 0x4f, 0x99, 0x9d, 0xf9, 0xcd, 0x6f, 0x27, 0xf3, 0x9b,
	0x1d, 0xd8, 0x2b, 0x51, 0x17, 0x4a, 0x0b, 0xa3, 0x4a, 0x81, 0x7a, 0x52, 0x94, 0xca, 0x28, 0xd6,
	0xa7, 0x8f, 0x8e, 0x1e, 0xcd, 0x95, 0x9a, 0x67, 0x78, 0x40, 0xc7, 0x8b, 0x6a, 0x76, 0x60, 0x44,
	0x8e, 0xda, 0xf0, 0xbc, 0x70, 0xc0, 0x68, 0x3b, 0x51, 0x79, 0xae, 0xa4, 0x3b, 0x3d, 0xfe, 0x1d,
	0xc0, 0xde, 0x19, 0x9f, 0x0b, 0xc9, 0x0d, 0xa6, 0x31, 0xd6, 0xb4, 0x37, 0x2c, 0x84, 0x41, 0xc1,
	0xe7, 0x98, 0xa1, 0x0c, 0xbd, 0x91, 0x37, 0xf6, 0xe2, 0xe6, 0xc8, 0x18, 0x74, 0xb5, 0xf8, 0x81,
	0xe1, 0x16, 0xb9, 0xc9, 0x66, 0xef, 0xa0, 0x7f, 0xc5, 0xb3, 0x0a, 0x75, 0xd8, 0x19, 0x75, 0xc6,
	0xc1, 0xf4, 0x85, 0x63, 0xd7, 0x93, 0x35, 0xd4, 0x93, 0xd6, 0xfc, 0x4c, 0x39, 0x71, 0x9d, 0x6b,
	0x99, 0xed, 0x25, 0x61, 0xd7, 0x31, 0x5b, 0xdb, 0xfa, 0x24, 0x7e, 0x37, 0x61, 0x6f, 0xe4, 0x8d,
	0xfd, 0x98, 0xec, 0xe8, 0x27, 0xc0, 0xee, 0x7d, 0x12, 0x02, 0xf2, 0x1c, 0xa9, 0x5a, 0x0b, 0xe4,
	0x39, 0xb2, 0x18, 0x7a, 0x99, 0x90, 0x0b, 0x4d, 0xb5, 0x06, 0xd3, 0xd7, 0x9b, 0x54, 0xb5, 0xe4,
	0x78, 0x6f, 0x39, 0x62, 0x47, 0xc5, 0x4e, 0x60, 0x50, 0x94, 0xea, 0x2b, 0x26, 0x26, 0xec, 0x10,
	0xeb, 0xd1, 0x46, 0xac, 0x67, 0x2e, 0x37, 0x6e, 0x48, 0xd8, 0x31, 0x40, 0x52, 0xa2, 0x4d, 0x3b,
	0x57, 0x92, 0x7e, 0x3d, 0x98, 0x46, 0x13, 0x27, 0xe2, 0xa4, 0x11, 0x71, 0xf2, 0xb1, 0x11, 0x31,
	0xf6, 0x6b, 0xf4, 0xa9, 0x64, 0x5f, 0x00, 0x72, 0x2e, 0xe4, 0x45, 0xc9, 0x65, 0x72, 0x49, 0x1d,
	0x0a, 0xa6, 0xaf, 0x36, 0xaa, 0xe6, 0x03, 0x17, 0xf2, 0x2d, 0xa5, 0xc7, 0x4b, 0x54, 0xec, 0x01,
	0xf8, 0xb3, 0x2a, 0xcb, 0xce, 0xa9, 0xa1, 0x7d, 0x6a, 0xe8, 0xd0, 0x3a, 0x4e, 0x6c, 0x53, 0x8f,
	0x01, 0xaa, 0x22, 0x6d, 0x0a, 0x1e, 0xfc, 0xbb, 0xe0, 0x1a, 0x7d, 0xda, 0x8e, 0xce, 0x70, 0x69,
	0x74, 0x18, 0x74, 0xcd, 0x4d, 0x81, 0xa1, 0xef, 0x74, 0xb3, 0x36, 0xe1, 0xb2, 0x6a, 0x1e, 0x82,
	0xf3, 0x59, 0x9b, 0x3d, 0x04, 0x10, 0xfa, 0xbc, 0x28, 0xc5, 0x15, 0x37, 0x18, 0x06, 0x23, 0x6f,
	0x3c, 0x8c, 0x7d, 0xa1, 0xcf, 0x9c, 0x83, 0x8d, 0x20, 0x48, 0x51, 0x27, 0xa5, 0x28, 0x8c, 0x50,
	0x32, 0xdc, 0xa6, 0xcc, 0x65, 0x57, 0xf4, 0xa7, 0x03, 0x3b, 0xf7, 0x34, 0x65, 0xcf, 0x61, 0x78,
	0xcd, 0x4d, 0x72, 0x89, 0xa5, 0xa6, 0xc1, 0x09, 0xa6, 0x3b, 0x4d, 0xff, 0x2c, 0xe0, 0x53, 0x99,
	0xc5, 0xb7, 0x00, 0x0b, 0x76, 0xfd, 0xc1, 0x66, 0xa0, 0x56, 0xc1, 0x0d, 0x80, 0x3d, 0x81, 0xae,
	0xe1, 0x73, 0x5d, 0xcf, 0xc8, 0x0a, 0x90, 0x82, 0x6c, 0x1f, 0x06, 0xf6, 0x31, 0x0a, 0xa3, 0x6b,
	0xe1, 0x57, 0x70, 0x4d, 0x9c, 0xed, 0x43, 0x2f, 0xc9, 0x94, 0xc4, 0xb0, 0x47, 0x0f, 0x6c, 0x6f,
	0x19, 0xf8, 0x46, 0xa6, 0x56, 0x99, 0xd8, 0x21, 0xec, 0xd5, 0x1a, 0xb3, 0x19, 0x09, 0xb7, 0xee,
	0x6a, 0x1b, 0x64, 0xcf, 0xa0, 0xaf, 0x55, 0x55, 0x26, 0x58, 0x2b, 0xb8, 0x02, 0xab, 0xc3, 0xec,
	0x29, 0xf4, 0x2e, 0x95, 0x5a, 0x68, 0x12, 0x6d, 0x0d, 0xce, 0x45, 0x2d, 0x6c, 0xa6, 0xca, 0x85,
	0x26, 0x1d, 0xd7, 0xc1, 0x28, 0xca, 0x5e, 0x82, 0x9f, 0xaa, 0x6b, 0x99, 0x29, 0x9e, 0x6a, 0x92,
	0x77, 0x0d, 0xb4, 0x45, 0xb0, 0x43, 0xd8, 0x2e, 0xaa, 0x2c, 0x2b, 0xf1, 0x5b, 0x85, 0xda, 0x68,
	0x92, 0x7d, 0x4d, 0xc6, 0x1d, 0x50, 0xf4, 0xcb, 0x83, 0x41, 0xfd, 0xcc, 0xd8, 0x2e, 0x74, 0x16,
	0x78, 0x53, 0x2f, 0x05, 0x6b, 0xde, 0xce, 0xdb, 0xd6, 0xdd, 0x79, 0xab, 0x2a, 0x91, 0x92, 0x58,
	0x7e, 0x4c, 0x76, 0xbb, 0x3b, 0xba, 0xff, 0x6f, 0x77, 0x34, 0x3b, 0xaa, 0xd7, 0xee, 0xa8, 0xe8,
	0x08, 0xa0, 0x7d, 0x85, 0xb7, 0xd5, 0x79, 0x77, 0xab, 0xa3, 0xac, 0xad, 0x36, 0xeb, 0xc2, 0x6d,
	0xfb, 0xc3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0x69, 0x2b, 0xc6, 0x0b, 0x06, 0x00, 0x00,
}