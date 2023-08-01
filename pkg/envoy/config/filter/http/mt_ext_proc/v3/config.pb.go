// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/mt_ext_proc/v3/config.proto

package v3

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
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

// [#next-free-field: 6]
type MtExtProc struct {
	// External processing http service configuration.
	HttpService *HttpService `protobuf:"bytes,1,opt,name=http_service,json=httpService,proto3" json:"http_service,omitempty"`
	// Sets the HTTP status that is returned to the client when there is a network error between the
	// filter and the processing server. The default status is HTTP 500 Internal Server Error.
	StatusOnError uint64 `protobuf:"varint,2,opt,name=status_on_error,json=statusOnError,proto3" json:"status_on_error,omitempty"`
	// Optional additional prefix to use when emitting statistics. This allows to distinguish
	// emitted statistics between configured *mt_ext_proc* filters in an HTTP filter chain. For example:
	//
	// .. code-block:: yaml
	//
	//   http_filters:
	//     - name: envoy.filters.http.mt_ext_proc
	//       typed_config:
	//         "@type": type.googleapis.com/envoy.extensions.filters.http.mt_ext_proc.v3.MtExtProc
	//         stat_prefix: waf # This emits mt_ext_proc.waf.ok, mt_ext_proc.waf.denied, etc.
	//     - name: envoy.filters.http.mt_ext_proc
	//       typed_config:
	//         "@type": type.googleapis.com/envoy.extensions.filters.http.mt_ext_proc.v3.MtExtProc
	//         stat_prefix: blocker # This emits mt_ext_proc.blocker.ok, mt_ext_proc.blocker.denied, etc.
	//
	StatPrefix string `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Config for decode filter
	Request *RequestConfig `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	// Config for encode filter
	Response             *ResponseConfig `protobuf:"bytes,5,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MtExtProc) Reset()         { *m = MtExtProc{} }
func (m *MtExtProc) String() string { return proto.CompactTextString(m) }
func (*MtExtProc) ProtoMessage()    {}
func (*MtExtProc) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{0}
}

func (m *MtExtProc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MtExtProc.Unmarshal(m, b)
}
func (m *MtExtProc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MtExtProc.Marshal(b, m, deterministic)
}
func (m *MtExtProc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MtExtProc.Merge(m, src)
}
func (m *MtExtProc) XXX_Size() int {
	return xxx_messageInfo_MtExtProc.Size(m)
}
func (m *MtExtProc) XXX_DiscardUnknown() {
	xxx_messageInfo_MtExtProc.DiscardUnknown(m)
}

var xxx_messageInfo_MtExtProc proto.InternalMessageInfo

func (m *MtExtProc) GetHttpService() *HttpService {
	if m != nil {
		return m.HttpService
	}
	return nil
}

func (m *MtExtProc) GetStatusOnError() uint64 {
	if m != nil {
		return m.StatusOnError
	}
	return 0
}

func (m *MtExtProc) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MtExtProc) GetRequest() *RequestConfig {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *MtExtProc) GetResponse() *ResponseConfig {
	if m != nil {
		return m.Response
	}
	return nil
}

// HttpService is used for raw HTTP communication between the filter and the processing service.
// When configured, the filter will parse the client request and use these attributes to call the
// processing server. Depending on the response, the filter may reject or accept the client
// request. Note that in any of these events, metadata can be added, removed or overridden by the
// filter:
//
// [#next-free-field: 2]
type HttpService struct {
	// Sets the HTTP server URI which the processing requests must be sent to.
	ServerUri            *HttpUri `protobuf:"bytes,1,opt,name=server_uri,json=serverUri,proto3" json:"server_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{1}
}

func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (m *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(m, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

// Copied from @envoy/api/envoy/api/v2/core/http_uri.proto
// Envoy external URI descriptor
type HttpUri struct {
	// The HTTP server URI. It should be a full FQDN with protocol, host and path.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//    uri: https://www.googleapis.com/oauth2/v1/certs
	//
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Specify how `uri` is to be fetched. Today, this requires an explicit
	// cluster, but in the future we may support dynamic cluster creation or
	// inline DNS resolution. See `issue
	// <https://github.com/envoyproxy/envoy/issues/1606>`_.
	//
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	// Sets the maximum duration in milliseconds that a response can take to arrive upon request.
	Timeout              *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{2}
}

func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpUri.Unmarshal(m, b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
}
func (m *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(m, src)
}
func (m *HttpUri) XXX_Size() int {
	return xxx_messageInfo_HttpUri.Size(m)
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpUri) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

type RequestConfig struct {
	// Config for request signature
	Signature *SignatureConfig `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Config for request decrypt
	Decrypt *DecryptConfig `protobuf:"bytes,2,opt,name=decrypt,proto3" json:"decrypt,omitempty"`
	// config for buffering body for failover
	Buffer               *BufferSettings `protobuf:"bytes,3,opt,name=buffer,proto3" json:"buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RequestConfig) Reset()         { *m = RequestConfig{} }
func (m *RequestConfig) String() string { return proto.CompactTextString(m) }
func (*RequestConfig) ProtoMessage()    {}
func (*RequestConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{3}
}

func (m *RequestConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestConfig.Unmarshal(m, b)
}
func (m *RequestConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestConfig.Marshal(b, m, deterministic)
}
func (m *RequestConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestConfig.Merge(m, src)
}
func (m *RequestConfig) XXX_Size() int {
	return xxx_messageInfo_RequestConfig.Size(m)
}
func (m *RequestConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RequestConfig proto.InternalMessageInfo

func (m *RequestConfig) GetSignature() *SignatureConfig {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *RequestConfig) GetDecrypt() *DecryptConfig {
	if m != nil {
		return m.Decrypt
	}
	return nil
}

func (m *RequestConfig) GetBuffer() *BufferSettings {
	if m != nil {
		return m.Buffer
	}
	return nil
}

type SignatureConfig struct {
	// Enable the filter for signature verifying.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	//  Changes filter's behaviour on errors:
	//
	//  1. When set to true, the filter will *accept* client request even if the communication with
	//  the processing service has failed, or if the processing service has returned a HTTP 5xx
	//  error.
	//
	//  2. When set to false, ext-proc will *reject* client requests and return a *Forbidden*
	//  response if the communication with the processing service has failed, or if the
	//  processing service has returned a HTTP 5xx error.
	//
	// Note that errors can be *always* tracked in the :ref:`stats
	// <config_http_filters_mt_ext_proc_stats>`.
	FailureModeAllow bool `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	// Sets the maximum duration in milliseconds that a response can take to arrive upon request.
	// default: 500ms
	Timeout              *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SignatureConfig) Reset()         { *m = SignatureConfig{} }
func (m *SignatureConfig) String() string { return proto.CompactTextString(m) }
func (*SignatureConfig) ProtoMessage()    {}
func (*SignatureConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{4}
}

func (m *SignatureConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureConfig.Unmarshal(m, b)
}
func (m *SignatureConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureConfig.Marshal(b, m, deterministic)
}
func (m *SignatureConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureConfig.Merge(m, src)
}
func (m *SignatureConfig) XXX_Size() int {
	return xxx_messageInfo_SignatureConfig.Size(m)
}
func (m *SignatureConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureConfig proto.InternalMessageInfo

func (m *SignatureConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *SignatureConfig) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

func (m *SignatureConfig) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type DecryptConfig struct {
	// Sets the maximum duration in milliseconds that a response can take to arrive upon request.
	// default: 2000ms
	Timeout              *duration.Duration `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DecryptConfig) Reset()         { *m = DecryptConfig{} }
func (m *DecryptConfig) String() string { return proto.CompactTextString(m) }
func (*DecryptConfig) ProtoMessage()    {}
func (*DecryptConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{5}
}

func (m *DecryptConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptConfig.Unmarshal(m, b)
}
func (m *DecryptConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptConfig.Marshal(b, m, deterministic)
}
func (m *DecryptConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptConfig.Merge(m, src)
}
func (m *DecryptConfig) XXX_Size() int {
	return xxx_messageInfo_DecryptConfig.Size(m)
}
func (m *DecryptConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptConfig proto.InternalMessageInfo

func (m *DecryptConfig) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// Configuration for buffering the request data.
type BufferSettings struct {
	// Sets the maximum size of a message body that the filter will hold in memory.
	// Envoy will *not* initiate the external process when buffer reaches the number
	// set in this field if condition is ok, and will proxy the request to upstream directly.
	MaxRequestBytes      uint32   `protobuf:"varint,2,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BufferSettings) Reset()         { *m = BufferSettings{} }
func (m *BufferSettings) String() string { return proto.CompactTextString(m) }
func (*BufferSettings) ProtoMessage()    {}
func (*BufferSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{6}
}

func (m *BufferSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferSettings.Unmarshal(m, b)
}
func (m *BufferSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferSettings.Marshal(b, m, deterministic)
}
func (m *BufferSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferSettings.Merge(m, src)
}
func (m *BufferSettings) XXX_Size() int {
	return xxx_messageInfo_BufferSettings.Size(m)
}
func (m *BufferSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferSettings.DiscardUnknown(m)
}

var xxx_messageInfo_BufferSettings proto.InternalMessageInfo

func (m *BufferSettings) GetMaxRequestBytes() uint32 {
	if m != nil {
		return m.MaxRequestBytes
	}
	return 0
}

type ResponseConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseConfig) Reset()         { *m = ResponseConfig{} }
func (m *ResponseConfig) String() string { return proto.CompactTextString(m) }
func (*ResponseConfig) ProtoMessage()    {}
func (*ResponseConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{7}
}

func (m *ResponseConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseConfig.Unmarshal(m, b)
}
func (m *ResponseConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseConfig.Marshal(b, m, deterministic)
}
func (m *ResponseConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseConfig.Merge(m, src)
}
func (m *ResponseConfig) XXX_Size() int {
	return xxx_messageInfo_ResponseConfig.Size(m)
}
func (m *ResponseConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseConfig proto.InternalMessageInfo

// Extra settings on a per virtualhost/route/weighted-cluster level.
type MtExtProcPerRoute struct {
	Request              *RequestConfigPerRoute `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MtExtProcPerRoute) Reset()         { *m = MtExtProcPerRoute{} }
func (m *MtExtProcPerRoute) String() string { return proto.CompactTextString(m) }
func (*MtExtProcPerRoute) ProtoMessage()    {}
func (*MtExtProcPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{8}
}

func (m *MtExtProcPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MtExtProcPerRoute.Unmarshal(m, b)
}
func (m *MtExtProcPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MtExtProcPerRoute.Marshal(b, m, deterministic)
}
func (m *MtExtProcPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MtExtProcPerRoute.Merge(m, src)
}
func (m *MtExtProcPerRoute) XXX_Size() int {
	return xxx_messageInfo_MtExtProcPerRoute.Size(m)
}
func (m *MtExtProcPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_MtExtProcPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_MtExtProcPerRoute proto.InternalMessageInfo

func (m *MtExtProcPerRoute) GetRequest() *RequestConfigPerRoute {
	if m != nil {
		return m.Request
	}
	return nil
}

type RequestConfigPerRoute struct {
	Signature            *SignatureConfigPerRoute `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RequestConfigPerRoute) Reset()         { *m = RequestConfigPerRoute{} }
func (m *RequestConfigPerRoute) String() string { return proto.CompactTextString(m) }
func (*RequestConfigPerRoute) ProtoMessage()    {}
func (*RequestConfigPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{9}
}

func (m *RequestConfigPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestConfigPerRoute.Unmarshal(m, b)
}
func (m *RequestConfigPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestConfigPerRoute.Marshal(b, m, deterministic)
}
func (m *RequestConfigPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestConfigPerRoute.Merge(m, src)
}
func (m *RequestConfigPerRoute) XXX_Size() int {
	return xxx_messageInfo_RequestConfigPerRoute.Size(m)
}
func (m *RequestConfigPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestConfigPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RequestConfigPerRoute proto.InternalMessageInfo

func (m *RequestConfigPerRoute) GetSignature() *SignatureConfigPerRoute {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignatureConfigPerRoute struct {
	// Disable the ext proc filter for this particular vhost or route.
	// If disabled is specified in multiple per-filter-configs, the most specific one will be used.
	Disabled             bool     `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureConfigPerRoute) Reset()         { *m = SignatureConfigPerRoute{} }
func (m *SignatureConfigPerRoute) String() string { return proto.CompactTextString(m) }
func (*SignatureConfigPerRoute) ProtoMessage()    {}
func (*SignatureConfigPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_6729377b0d6b2f44, []int{10}
}

func (m *SignatureConfigPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureConfigPerRoute.Unmarshal(m, b)
}
func (m *SignatureConfigPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureConfigPerRoute.Marshal(b, m, deterministic)
}
func (m *SignatureConfigPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureConfigPerRoute.Merge(m, src)
}
func (m *SignatureConfigPerRoute) XXX_Size() int {
	return xxx_messageInfo_SignatureConfigPerRoute.Size(m)
}
func (m *SignatureConfigPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureConfigPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureConfigPerRoute proto.InternalMessageInfo

func (m *SignatureConfigPerRoute) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func init() {
	proto.RegisterType((*MtExtProc)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.MtExtProc")
	proto.RegisterType((*HttpService)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.HttpService")
	proto.RegisterType((*HttpUri)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.HttpUri")
	proto.RegisterType((*RequestConfig)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.RequestConfig")
	proto.RegisterType((*SignatureConfig)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.SignatureConfig")
	proto.RegisterType((*DecryptConfig)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.DecryptConfig")
	proto.RegisterType((*BufferSettings)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.BufferSettings")
	proto.RegisterType((*ResponseConfig)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.ResponseConfig")
	proto.RegisterType((*MtExtProcPerRoute)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.MtExtProcPerRoute")
	proto.RegisterType((*RequestConfigPerRoute)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.RequestConfigPerRoute")
	proto.RegisterType((*SignatureConfigPerRoute)(nil), "istio.envoy.config.filter.http.mt_ext_proc.v3.SignatureConfigPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/mt_ext_proc/v3/config.proto", fileDescriptor_6729377b0d6b2f44)
}

var fileDescriptor_6729377b0d6b2f44 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xc5, 0x69, 0x69, 0x92, 0x1b, 0xd2, 0xc7, 0x08, 0x44, 0xc8, 0x82, 0x56, 0x5e, 0xa0, 0x0a,
	0x81, 0x2d, 0x35, 0xa5, 0x0b, 0x54, 0x90, 0x1a, 0x52, 0xd4, 0x4d, 0x45, 0xe5, 0xaa, 0x48, 0x54,
	0x08, 0xcb, 0x8e, 0xaf, 0xd3, 0x91, 0x1c, 0x8f, 0x99, 0x19, 0x87, 0x64, 0xc1, 0x86, 0x0f, 0xe0,
	0x0b, 0xf8, 0x38, 0xf6, 0xfc, 0x04, 0xf2, 0xcc, 0x38, 0x0f, 0x1e, 0x52, 0xd3, 0xee, 0xe2, 0x33,
	0xf7, 0x9c, 0xb9, 0xf7, 0xcc, 0x99, 0x09, 0xec, 0x63, 0x3a, 0x62, 0x13, 0xb7, 0xcf, 0xd2, 0x98,
	0x0e, 0xdc, 0x98, 0x26, 0x12, 0xb9, 0x7b, 0x25, 0x65, 0xe6, 0x0e, 0xa5, 0x8f, 0x63, 0xe9, 0x67,
	0x9c, 0xf5, 0xdd, 0x51, 0xc7, 0x54, 0x38, 0x19, 0x67, 0x92, 0x91, 0xe7, 0x54, 0x48, 0xca, 0x1c,
	0xc5, 0x75, 0xcc, 0x8a, 0xe6, 0x3a, 0x05, 0xd7, 0x99, 0xe3, 0x3a, 0xa3, 0x4e, 0xfb, 0xf1, 0x80,
	0xb1, 0x41, 0x82, 0xae, 0x22, 0x87, 0x79, 0xec, 0x46, 0x39, 0x0f, 0x24, 0x65, 0xa9, 0x96, 0x6b,
	0x6f, 0x9b, 0xf5, 0x20, 0xa3, 0x6e, 0x4c, 0x31, 0x89, 0xfc, 0x10, 0xaf, 0x82, 0x11, 0x65, 0x5c,
	0x17, 0xd8, 0xbf, 0x2a, 0x50, 0x3f, 0x95, 0xc7, 0x63, 0x79, 0xc6, 0x59, 0x9f, 0x84, 0x70, 0xaf,
	0xd8, 0xc5, 0x17, 0xc8, 0x47, 0xb4, 0x8f, 0x2d, 0x6b, 0xc7, 0xda, 0x6d, 0xec, 0xbd, 0x74, 0x96,
	0x6a, 0xca, 0x39, 0x91, 0x32, 0x3b, 0xd7, 0x0a, 0xdd, 0x95, 0x9f, 0x47, 0x15, 0xaf, 0x71, 0x35,
	0x43, 0xc8, 0x13, 0xd8, 0x10, 0x32, 0x90, 0xb9, 0xf0, 0x59, 0xea, 0x23, 0xe7, 0x8c, 0xb7, 0x2a,
	0x3b, 0xd6, 0xee, 0xaa, 0xd7, 0xd4, 0xf0, 0xbb, 0xf4, 0xb8, 0x00, 0xc9, 0x36, 0x34, 0x0a, 0xc0,
	0xcf, 0x38, 0xc6, 0x74, 0xdc, 0x5a, 0xd9, 0xb1, 0x76, 0xeb, 0x1e, 0x14, 0xd0, 0x99, 0x42, 0xc8,
	0x7b, 0xa8, 0x72, 0xfc, 0x9c, 0xa3, 0x90, 0xad, 0x55, 0xd5, 0xe7, 0xe1, 0x92, 0x7d, 0x7a, 0x9a,
	0xfd, 0x46, 0x15, 0x7a, 0xa5, 0x18, 0xf9, 0x00, 0x35, 0x8e, 0x22, 0x63, 0xa9, 0xc0, 0xd6, 0x5d,
	0x25, 0xfc, 0x6a, 0x69, 0x61, 0x4d, 0x37, 0xca, 0x53, 0x39, 0x9b, 0x42, 0x63, 0xce, 0x1c, 0x72,
	0x09, 0x50, 0x38, 0x8d, 0xdc, 0xcf, 0x39, 0x35, 0x66, 0x1f, 0xdc, 0xc0, 0xec, 0x0b, 0x4e, 0xb5,
	0xd1, 0x75, 0x2d, 0x77, 0xc1, 0xa9, 0xfd, 0xcd, 0x82, 0xaa, 0x59, 0x23, 0x9b, 0xb0, 0x52, 0x6e,
	0x50, 0xf7, 0x8a, 0x9f, 0xa4, 0x0d, 0xd5, 0x7e, 0x92, 0x0b, 0x89, 0xda, 0xfc, 0xfa, 0xc9, 0x1d,
	0xaf, 0x04, 0x48, 0x07, 0xaa, 0x92, 0x0e, 0x91, 0xe5, 0x52, 0x99, 0xde, 0xd8, 0x7b, 0xe4, 0xe8,
	0x14, 0x39, 0x65, 0xca, 0x9c, 0x9e, 0x49, 0x99, 0x57, 0x56, 0x76, 0xef, 0x03, 0x51, 0xc9, 0xc9,
	0x33, 0x21, 0x39, 0x06, 0x43, 0x5f, 0x4e, 0x32, 0xb4, 0x7f, 0x54, 0xa0, 0xb9, 0xe0, 0x32, 0xf9,
	0x08, 0x75, 0x41, 0x07, 0x69, 0x20, 0x73, 0x5e, 0xc6, 0xeb, 0xf5, 0x92, 0x13, 0x9f, 0x97, 0x7c,
	0x63, 0xef, 0x4c, 0xb0, 0x88, 0x44, 0x84, 0x7d, 0x3e, 0xc9, 0xa4, 0x1a, 0x6b, 0xf9, 0x48, 0xf4,
	0x34, 0xbb, 0x8c, 0x84, 0x11, 0x23, 0x17, 0xb0, 0x16, 0xe6, 0x71, 0x8c, 0xdc, 0x38, 0xb2, 0x6c,
	0x20, 0xba, 0x8a, 0x7c, 0x8e, 0x52, 0xd2, 0x74, 0x20, 0x3c, 0x23, 0x66, 0x7f, 0xb7, 0x60, 0xe3,
	0x8f, 0x69, 0x48, 0x0b, 0xaa, 0x98, 0x06, 0x61, 0x82, 0x91, 0xb2, 0xa7, 0xe6, 0x95, 0x9f, 0xe4,
	0x19, 0x90, 0x38, 0xa0, 0x49, 0xce, 0xd1, 0x1f, 0xb2, 0x08, 0xfd, 0x20, 0x49, 0xd8, 0x17, 0x35,
	0x67, 0xcd, 0xdb, 0x34, 0x2b, 0xa7, 0x2c, 0xc2, 0xa3, 0x02, 0xbf, 0xd1, 0x29, 0xda, 0x3d, 0x68,
	0x2e, 0x38, 0x30, 0xaf, 0x62, 0x5d, 0x5b, 0xe5, 0x10, 0xd6, 0x17, 0x07, 0x26, 0x4f, 0x61, 0x6b,
	0x18, 0x8c, 0x7d, 0x73, 0xc3, 0xfc, 0x70, 0x22, 0x51, 0xa8, 0xce, 0x9b, 0xde, 0xc6, 0x30, 0x18,
	0x9b, 0x88, 0x74, 0x0b, 0xd8, 0xde, 0x84, 0xf5, 0xc5, 0xfb, 0x63, 0x0b, 0xd8, 0x9a, 0x3e, 0x51,
	0x67, 0xc8, 0x3d, 0x96, 0x4b, 0x24, 0x9f, 0x66, 0xb7, 0x5f, 0x77, 0xd6, 0xbb, 0xcd, 0xed, 0x2f,
	0x65, 0xa7, 0xaf, 0x80, 0xfd, 0x15, 0x1e, 0xfc, 0xb3, 0x82, 0x44, 0x7f, 0x27, 0xf8, 0xed, 0xed,
	0x12, 0x3c, 0xdd, 0x7c, 0x26, 0x6c, 0xbf, 0x80, 0x87, 0xff, 0xa9, 0x22, 0x6d, 0xa8, 0x45, 0x54,
	0xcc, 0x47, 0x64, 0xfa, 0xdd, 0x3d, 0xb8, 0xdc, 0xd7, 0xad, 0x50, 0xa6, 0xde, 0xfc, 0x6b, 0xfe,
	0x07, 0x85, 0x6b, 0xea, 0x38, 0x3b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x60, 0xfa, 0xd9, 0x3d,
	0xb5, 0x06, 0x00, 0x00,
}
