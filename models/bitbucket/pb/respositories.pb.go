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

// some fields are omitted from the response, but if you need anything else,
// see https://confluence.atlassian.com/bitbucket/changesets-resource-296095208.html#changesetsResource-GETanindividualchangeset
// for a list of avilable data on changesets
type ChangeSetV1 struct {
	Node      string `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	RawAuthor string `protobuf:"bytes,2,opt,name=raw_author,json=rawAuthor" json:"raw_author,omitempty"`
	Author    string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	RawNode   string `protobuf:"bytes,4,opt,name=raw_node,json=rawNode" json:"raw_node,omitempty"`
	Branch    string `protobuf:"bytes,5,opt,name=branch" json:"branch,omitempty"`
}

func (m *ChangeSetV1) Reset()                    { *m = ChangeSetV1{} }
func (m *ChangeSetV1) String() string            { return proto.CompactTextString(m) }
func (*ChangeSetV1) ProtoMessage()               {}
func (*ChangeSetV1) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ChangeSetV1) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *ChangeSetV1) GetRawAuthor() string {
	if m != nil {
		return m.RawAuthor
	}
	return ""
}

func (m *ChangeSetV1) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ChangeSetV1) GetRawNode() string {
	if m != nil {
		return m.RawNode
	}
	return ""
}

func (m *ChangeSetV1) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func init() {
	proto.RegisterType((*PaginatedRepository)(nil), "protos.PaginatedRepository")
	proto.RegisterType((*PaginatedRepository_RepositoryValues)(nil), "protos.PaginatedRepository.RepositoryValues")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_RepositoryLinks)(nil), "protos.PaginatedRepository.RepositoryValues.RepositoryLinks")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_Project)(nil), "protos.PaginatedRepository.RepositoryValues.Project")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_MainBranch)(nil), "protos.PaginatedRepository.RepositoryValues.MainBranch")
	proto.RegisterType((*ChangeSetV1)(nil), "protos.ChangeSetV1")
}

func init() { proto.RegisterFile("respositories.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x6f, 0xd3, 0x4a,
	0x10, 0x97, 0x9b, 0xbf, 0x1e, 0x57, 0x6a, 0xb5, 0x95, 0x2a, 0xbf, 0x3c, 0x3d, 0xbd, 0xa8, 0x08,
	0x91, 0x0a, 0x48, 0x45, 0x5a, 0x09, 0x55, 0xe2, 0x52, 0xe0, 0x08, 0x6d, 0x65, 0xa0, 0x1c, 0xa3,
	0xad, 0xbd, 0x49, 0x96, 0xd8, 0xbb, 0x66, 0x77, 0xdd, 0x50, 0x4e, 0xdc, 0xf9, 0x28, 0x5c, 0xf9,
	0x4c, 0x7c, 0x0e, 0xb4, 0xb3, 0x76, 0x9d, 0x36, 0x91, 0x50, 0x25, 0x4e, 0x99, 0x9d, 0xf9, 0xcd,
	0x6f, 0xc7, 0xf3, 0x9b, 0x9d, 0xc0, 0x8e, 0x62, 0x3a, 0x97, 0x9a, 0x1b, 0xa9, 0x38, 0xd3, 0xc3,
	0x5c, 0x49, 0x23, 0x49, 0x1b, 0x7f, 0x74, 0xef, 0xff, 0xa9, 0x94, 0xd3, 0x94, 0x1d, 0xe0, 0xf1,
	0xb2, 0x98, 0x1c, 0x18, 0x9e, 0x31, 0x6d, 0x68, 0x96, 0x3b, 0x60, 0x6f, 0x33, 0x96, 0x59, 0x26,
	0x85, 0x3b, 0xed, 0xfd, 0x0c, 0x60, 0xe7, 0x9c, 0x4e, 0xb9, 0xa0, 0x86, 0x25, 0x11, 0x2b, 0x69,
	0xaf, 0x49, 0x08, 0x9d, 0x9c, 0x4e, 0x59, 0xca, 0x44, 0xe8, 0xf5, 0xbd, 0x81, 0x17, 0x55, 0x47,
	0x42, 0xa0, 0xa9, 0xf9, 0x57, 0x16, 0x6e, 0xa0, 0x1b, 0x6d, 0xf2, 0x1a, 0xda, 0x57, 0x34, 0x2d,
	0x98, 0x0e, 0x1b, 0xfd, 0xc6, 0x20, 0x18, 0x3d, 0x71, 0xec, 0x7a, 0xb8, 0x86, 0x7a, 0x58, 0x9b,
	0x17, 0x98, 0x13, 0x95, 0xb9, 0x96, 0xd9, 0x5e, 0x12, 0x36, 0x1d, 0xb3, 0xb5, 0xad, 0x4f, 0xb0,
	0x2f, 0x26, 0x6c, 0xf5, 0xbd, 0x81, 0x1f, 0xa1, 0xdd, 0xfb, 0x06, 0xb0, 0x7d, 0x97, 0x04, 0x81,
	0x34, 0x63, 0x58, 0xad, 0x05, 0xd2, 0x8c, 0x91, 0x08, 0x5a, 0x29, 0x17, 0x73, 0x8d, 0xb5, 0x06,
	0xa3, 0x17, 0xf7, 0xa9, 0x6a, 0xc9, 0xf1, 0xc6, 0x72, 0x44, 0x8e, 0x8a, 0x9c, 0x42, 0x27, 0x57,
	0xf2, 0x13, 0x8b, 0x4d, 0xd8, 0x40, 0xd6, 0xa3, 0x7b, 0xb1, 0x9e, 0xbb, 0xdc, 0xa8, 0x22, 0x21,
	0xc7, 0x00, 0xb1, 0x62, 0x36, 0x6d, 0x2c, 0x05, 0x7e, 0x7a, 0x30, 0xea, 0x0d, 0x9d, 0x88, 0xc3,
	0x4a, 0xc4, 0xe1, 0xfb, 0x4a, 0xc4, 0xc8, 0x2f, 0xd1, 0x67, 0x82, 0x7c, 0x04, 0xc8, 0x28, 0x17,
	0x97, 0x8a, 0x8a, 0x78, 0x86, 0x1d, 0x0a, 0x46, 0xcf, 0xef, 0x55, 0xcd, 0x5b, 0xca, 0xc5, 0x4b,
	0x4c, 0x8f, 0x96, 0xa8, 0xc8, 0xbf, 0xe0, 0x4f, 0x8a, 0x34, 0x1d, 0x63, 0x43, 0xdb, 0xd8, 0xd0,
	0xae, 0x75, 0x9c, 0xda, 0xa6, 0x1e, 0x03, 0x14, 0x79, 0x52, 0x15, 0xdc, 0xf9, 0x73, 0xc1, 0x25,
	0xfa, 0xac, 0x1e, 0x9d, 0xee, 0xd2, 0xe8, 0x10, 0x68, 0x9a, 0xeb, 0x9c, 0x85, 0xbe, 0xd3, 0xcd,
	0xda, 0x88, 0x4b, 0x8b, 0x69, 0x08, 0xce, 0x67, 0x6d, 0xf2, 0x1f, 0x00, 0xd7, 0xe3, 0x5c, 0xf1,
	0x2b, 0x6a, 0x58, 0x18, 0xf4, 0xbd, 0x41, 0x37, 0xf2, 0xb9, 0x3e, 0x77, 0x0e, 0xd2, 0x87, 0x20,
	0x61, 0x3a, 0x56, 0x3c, 0x37, 0x5c, 0x8a, 0x70, 0x13, 0x33, 0x97, 0x5d, 0xbd, 0x5f, 0x0d, 0xd8,
	0xba, 0xa3, 0x29, 0x79, 0x0c, 0xdd, 0x05, 0x35, 0xf1, 0x8c, 0x29, 0x8d, 0x83, 0x13, 0x8c, 0xb6,
	0xaa, 0xfe, 0x59, 0xc0, 0x07, 0x95, 0x46, 0x37, 0x00, 0x0b, 0x76, 0xfd, 0x61, 0xd5, 0x40, 0xad,
	0x82, 0x2b, 0x00, 0x79, 0x00, 0x4d, 0x43, 0xa7, 0xba, 0x9c, 0x91, 0x15, 0x20, 0x06, 0xc9, 0x3e,
	0x74, 0xec, 0x63, 0xe4, 0x46, 0x97, 0xc2, 0xaf, 0xe0, 0xaa, 0x38, 0xd9, 0x87, 0x56, 0x9c, 0x4a,
	0xc1, 0xc2, 0x16, 0x3e, 0xb0, 0x9d, 0x65, 0xe0, 0x89, 0x48, 0xac, 0x32, 0x91, 0x43, 0xd8, 0xab,
	0x35, 0x4b, 0x27, 0x28, 0xdc, 0xba, 0xab, 0x6d, 0x90, 0x3c, 0x82, 0xb6, 0x96, 0x85, 0x8a, 0x59,
	0xa9, 0xe0, 0x0a, 0xac, 0x0c, 0x93, 0x87, 0xd0, 0x9a, 0x49, 0x39, 0xd7, 0x28, 0xda, 0x1a, 0x9c,
	0x8b, 0x5a, 0xd8, 0x44, 0xaa, 0xb9, 0x46, 0x1d, 0xd7, 0xc1, 0x30, 0x4a, 0x9e, 0x82, 0x9f, 0xc8,
	0x85, 0x48, 0x25, 0x4d, 0x34, 0xca, 0xbb, 0x06, 0x5a, 0x23, 0xc8, 0x21, 0x6c, 0xe6, 0x45, 0x9a,
	0x2a, 0xf6, 0xb9, 0x60, 0xda, 0x68, 0x94, 0x7d, 0x4d, 0xc6, 0x2d, 0x50, 0xef, 0x87, 0x07, 0x9d,
	0xf2, 0x99, 0x91, 0x6d, 0x68, 0xcc, 0xd9, 0x75, 0xb9, 0x14, 0xac, 0x79, 0x33, 0x6f, 0x1b, 0xb7,
	0xe7, 0xad, 0x28, 0x78, 0x82, 0x62, 0xf9, 0x11, 0xda, 0xf5, 0xee, 0x68, 0xfe, 0xbd, 0xdd, 0x51,
	0xed, 0xa8, 0x56, 0xbd, 0xa3, 0x7a, 0x47, 0x00, 0xf5, 0x2b, 0xbc, 0xa9, 0xce, 0xbb, 0x5d, 0x1d,
	0x66, 0x6d, 0xd4, 0x59, 0x7b, 0xdf, 0x3d, 0x08, 0x5e, 0xcd, 0xa8, 0x98, 0xb2, 0x77, 0xcc, 0x5c,
	0x3c, 0x43, 0x8c, 0x4c, 0xea, 0xed, 0x27, 0x13, 0x66, 0x5f, 0x8c, 0xa2, 0x8b, 0x31, 0x2d, 0xcc,
	0x4c, 0xaa, 0x32, 0xdb, 0x57, 0x74, 0x71, 0x82, 0x0e, 0xb2, 0x0b, 0xed, 0x32, 0xe4, 0x3e, 0xbb,
	0x3c, 0x91, 0x7f, 0xa0, 0x6b, 0xd3, 0x90, 0xae, 0x89, 0x91, 0x8e, 0xa2, 0x8b, 0x53, 0xcb, 0xb8,
	0x0b, 0xed, 0xa5, 0x65, 0xe3, 0x47, 0xe5, 0xe9, 0xd2, 0xfd, 0xf7, 0x1c, 0xfe, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xef, 0x1c, 0x16, 0x3e, 0x99, 0x06, 0x00, 0x00,
}