// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/operator.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Status describes the current state of a component.
type IstioOperatorSpec_Status int32

const (
	// Component is not present.
	IstioOperatorSpec_NONE IstioOperatorSpec_Status = 0
	// Component is being updated to a different version.
	IstioOperatorSpec_UPDATING IstioOperatorSpec_Status = 1
	// Controller has started but not yet completed reconciliation loop for the component.
	IstioOperatorSpec_RECONCILING IstioOperatorSpec_Status = 2
	// Component is healthy.
	IstioOperatorSpec_HEALTHY IstioOperatorSpec_Status = 3
	// Component is in an error state.
	IstioOperatorSpec_ERROR IstioOperatorSpec_Status = 4
)

var IstioOperatorSpec_Status_name = map[int32]string{
	0: "NONE",
	1: "UPDATING",
	2: "RECONCILING",
	3: "HEALTHY",
	4: "ERROR",
}

var IstioOperatorSpec_Status_value = map[string]int32{
	"NONE":        0,
	"UPDATING":    1,
	"RECONCILING": 2,
	"HEALTHY":     3,
	"ERROR":       4,
}

func (x IstioOperatorSpec_Status) String() string {
	return proto.EnumName(IstioOperatorSpec_Status_name, int32(x))
}

func (IstioOperatorSpec_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_03e04c906468a7ff, []int{0, 0}
}

// IstioOperatorSpec defines the desired installed state of Istio components.
// The spec is a used to define a customization of the default profile values that are supplied with each Istio release.
// Because the spec is a customization API, specifying an empty IstioOperatorSpec results in a default Istio
// component values.
type IstioOperatorSpec struct {
	// Path or name for the profile e.g.
	//     - minimal (looks in profiles dir for a file called minimal.yaml)
	//     - /tmp/istio/install/values/custom/custom-install.yaml (local file path)
	// default profile is used if this field is unset.
	Profile string `protobuf:"bytes,10,opt,name=profile,proto3" json:"profile,omitempty"`
	// Path for the install package. e.g.
	//     - /tmp/istio-installer/nightly (local file path)
	InstallPackagePath string `protobuf:"bytes,11,opt,name=install_package_path,json=installPackagePath,proto3" json:"install_package_path,omitempty"`
	// Root for docker image paths e.g. docker.io/istio
	Hub string `protobuf:"bytes,12,opt,name=hub,proto3" json:"hub,omitempty"`
	// Version tag for docker images e.g. 1.0.6
	Tag string `protobuf:"bytes,13,opt,name=tag,proto3" json:"tag,omitempty"`
	// Resource suffix is appended to all resources installed by each component. Used in upgrade scenarios where two
	// Istio control planes must exist in the same namespace.
	ResourceSuffix string `protobuf:"bytes,14,opt,name=resource_suffix,json=resourceSuffix,proto3" json:"resource_suffix,omitempty"`
	// Config used by control plane components internally.
	MeshConfig *MeshConfig `protobuf:"bytes,40,opt,name=mesh_config,json=meshConfig,proto3" json:"mesh_config,omitempty"`
	// Kubernetes resource settings, enablement and component-specific settings that are not internal to the
	// component.
	Components *IstioComponentSetSpec `protobuf:"bytes,50,opt,name=components,proto3" json:"components,omitempty"`
	// Overrides for default values.yaml. This is a validated pass-through to Helm templates.
	// See the Helm installation options for schema details: https://istio.io/docs/reference/config/installation-options/.
	// Anything that is available in IstioOperatorSpec should be set above rather than using the passthrough. This
	// includes Kubernetes resource settings for components in KubernetesResourcesSpec.
	Values *TypeMapStringInterface `protobuf:"bytes,100,opt,name=values,proto3" json:"values,omitempty"`
	// Unvalidated overrides for default values.yaml. Used for custom templates where new parameters are added.
	UnvalidatedValues *TypeMapStringInterface `protobuf:"bytes,101,opt,name=unvalidated_values,json=unvalidatedValues,proto3" json:"unvalidated_values,omitempty"`
	// Overall status of all components controlled by the operator.
	// - If all components have status NONE, overall status is NONE.
	// - If all components are HEALTHY, overall status is HEALTHY.
	// - If one or more components are RECONCILING and others are HEALTHY, overall status is RECONCILING.
	// - If one or more components are UPDATING and others are HEALTHY, overall status is UPDATING.
	// - If components are a mix of RECONCILING, UPDATING and HEALTHY, overall status is UPDATING.
	// - If any component is in ERROR state, overall status is ERROR.
	Status IstioOperatorSpec_Status `protobuf:"varint,200,opt,name=status,proto3,enum=istio.mesh.v1alpha1.IstioOperatorSpec_Status" json:"status,omitempty"`
	// Individual status of each component controlled by the operator. The map key is the name of the component.
	ComponentStatus      map[string]*IstioOperatorSpec_VersionStatus `protobuf:"bytes,201,rep,name=component_status,json=componentStatus,proto3" json:"component_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *IstioOperatorSpec) Reset()         { *m = IstioOperatorSpec{} }
func (m *IstioOperatorSpec) String() string { return proto.CompactTextString(m) }
func (*IstioOperatorSpec) ProtoMessage()    {}
func (*IstioOperatorSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_03e04c906468a7ff, []int{0}
}
func (m *IstioOperatorSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioOperatorSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioOperatorSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioOperatorSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioOperatorSpec.Merge(m, src)
}
func (m *IstioOperatorSpec) XXX_Size() int {
	return m.Size()
}
func (m *IstioOperatorSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioOperatorSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioOperatorSpec proto.InternalMessageInfo

func (m *IstioOperatorSpec) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *IstioOperatorSpec) GetInstallPackagePath() string {
	if m != nil {
		return m.InstallPackagePath
	}
	return ""
}

func (m *IstioOperatorSpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *IstioOperatorSpec) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *IstioOperatorSpec) GetResourceSuffix() string {
	if m != nil {
		return m.ResourceSuffix
	}
	return ""
}

func (m *IstioOperatorSpec) GetMeshConfig() *MeshConfig {
	if m != nil {
		return m.MeshConfig
	}
	return nil
}

func (m *IstioOperatorSpec) GetComponents() *IstioComponentSetSpec {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *IstioOperatorSpec) GetValues() *TypeMapStringInterface {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *IstioOperatorSpec) GetUnvalidatedValues() *TypeMapStringInterface {
	if m != nil {
		return m.UnvalidatedValues
	}
	return nil
}

func (m *IstioOperatorSpec) GetStatus() IstioOperatorSpec_Status {
	if m != nil {
		return m.Status
	}
	return IstioOperatorSpec_NONE
}

func (m *IstioOperatorSpec) GetComponentStatus() map[string]*IstioOperatorSpec_VersionStatus {
	if m != nil {
		return m.ComponentStatus
	}
	return nil
}

// VersionStatus is the status and version of a component.
type IstioOperatorSpec_VersionStatus struct {
	Version              string                   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Status               IstioOperatorSpec_Status `protobuf:"varint,2,opt,name=status,proto3,enum=istio.mesh.v1alpha1.IstioOperatorSpec_Status" json:"status,omitempty"`
	StatusString         string                   `protobuf:"bytes,3,opt,name=status_string,json=statusString,proto3" json:"status_string,omitempty"`
	Error                string                   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IstioOperatorSpec_VersionStatus) Reset()         { *m = IstioOperatorSpec_VersionStatus{} }
func (m *IstioOperatorSpec_VersionStatus) String() string { return proto.CompactTextString(m) }
func (*IstioOperatorSpec_VersionStatus) ProtoMessage()    {}
func (*IstioOperatorSpec_VersionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_03e04c906468a7ff, []int{0, 0}
}
func (m *IstioOperatorSpec_VersionStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioOperatorSpec_VersionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioOperatorSpec_VersionStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioOperatorSpec_VersionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioOperatorSpec_VersionStatus.Merge(m, src)
}
func (m *IstioOperatorSpec_VersionStatus) XXX_Size() int {
	return m.Size()
}
func (m *IstioOperatorSpec_VersionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioOperatorSpec_VersionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IstioOperatorSpec_VersionStatus proto.InternalMessageInfo

func (m *IstioOperatorSpec_VersionStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *IstioOperatorSpec_VersionStatus) GetStatus() IstioOperatorSpec_Status {
	if m != nil {
		return m.Status
	}
	return IstioOperatorSpec_NONE
}

func (m *IstioOperatorSpec_VersionStatus) GetStatusString() string {
	if m != nil {
		return m.StatusString
	}
	return ""
}

func (m *IstioOperatorSpec_VersionStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("istio.mesh.v1alpha1.IstioOperatorSpec_Status", IstioOperatorSpec_Status_name, IstioOperatorSpec_Status_value)
	proto.RegisterType((*IstioOperatorSpec)(nil), "istio.mesh.v1alpha1.IstioOperatorSpec")
	proto.RegisterMapType((map[string]*IstioOperatorSpec_VersionStatus)(nil), "istio.mesh.v1alpha1.IstioOperatorSpec.ComponentStatusEntry")
	proto.RegisterType((*IstioOperatorSpec_VersionStatus)(nil), "istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus")
}

func init() { proto.RegisterFile("mesh/v1alpha1/operator.proto", fileDescriptor_03e04c906468a7ff) }

var fileDescriptor_03e04c906468a7ff = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0xeb, 0xd6, 0x6d, 0x37, 0xfb, 0xe8, 0xcc, 0x1e, 0xac, 0x0a, 0x46, 0x35, 0x1e, 0x88,
	0x40, 0xa4, 0xac, 0xf0, 0x80, 0xe0, 0x85, 0x51, 0x02, 0xeb, 0xb4, 0xb5, 0x53, 0x3a, 0x26, 0xb1,
	0x97, 0xc8, 0xcb, 0xdc, 0x24, 0x5a, 0x16, 0x47, 0xb6, 0x53, 0x6d, 0xbf, 0x89, 0x3f, 0x32, 0xde,
	0xf8, 0x09, 0x68, 0x7f, 0x04, 0x14, 0x3b, 0x2d, 0x2d, 0x2a, 0xd2, 0x10, 0x6f, 0xf6, 0xb9, 0xf7,
	0x9c, 0xdc, 0x73, 0x72, 0x65, 0x78, 0x70, 0xc9, 0x64, 0xd4, 0x1c, 0xee, 0xd0, 0x24, 0x8b, 0xe8,
	0x4e, 0x93, 0x67, 0x4c, 0x50, 0xc5, 0x85, 0x93, 0x09, 0xae, 0x38, 0xbe, 0x1f, 0x4b, 0x15, 0x73,
	0xa7, 0xe8, 0x71, 0x46, 0x3d, 0xf5, 0xfa, 0x34, 0x25, 0xe0, 0xe9, 0x20, 0x0e, 0x0d, 0xa1, 0xfe,
	0xf0, 0xcf, 0xda, 0x65, 0xc6, 0x53, 0x96, 0x2a, 0x53, 0xde, 0xfe, 0xb9, 0x08, 0x1b, 0x9d, 0x42,
	0xb2, 0x57, 0x7e, 0xa7, 0x9f, 0xb1, 0x00, 0x13, 0x58, 0xcc, 0x04, 0x1f, 0xc4, 0x09, 0x23, 0xd0,
	0x40, 0xf6, 0xb2, 0x37, 0xba, 0xe2, 0x17, 0xb0, 0x19, 0xa7, 0x52, 0xd1, 0x24, 0xf1, 0x33, 0x1a,
	0x5c, 0xd0, 0x90, 0xf9, 0x19, 0x55, 0x11, 0xb1, 0x74, 0x1b, 0x2e, 0x6b, 0x47, 0xa6, 0x74, 0x44,
	0x55, 0x84, 0x6b, 0x50, 0x89, 0xf2, 0x33, 0xb2, 0xa2, 0x1b, 0x8a, 0x63, 0x81, 0x28, 0x1a, 0x92,
	0x55, 0x83, 0x28, 0x1a, 0xe2, 0x27, 0xb0, 0x2e, 0x98, 0xe4, 0xb9, 0x08, 0x98, 0x2f, 0xf3, 0xc1,
	0x20, 0xbe, 0x22, 0x6b, 0xba, 0xba, 0x36, 0x82, 0xfb, 0x1a, 0xc5, 0xef, 0xc0, 0x2a, 0xfc, 0xf8,
	0xc6, 0x22, 0xb1, 0x1b, 0xc8, 0xb6, 0x5a, 0x8f, 0x9c, 0x19, 0xa1, 0x38, 0x87, 0x4c, 0x46, 0x6d,
	0xdd, 0xe6, 0xc1, 0xe5, 0xf8, 0x8c, 0xf7, 0x01, 0xc6, 0x19, 0x48, 0xd2, 0xd2, 0x02, 0x4f, 0x67,
	0x0a, 0xe8, 0x58, 0xda, 0xa3, 0xde, 0x3e, 0x53, 0x45, 0x34, 0xde, 0x04, 0x1b, 0xb7, 0xa1, 0x3a,
	0xa4, 0x49, 0xce, 0x24, 0x39, 0xd7, 0x3a, 0xcf, 0x66, 0xea, 0x1c, 0x5f, 0x67, 0xec, 0x90, 0x66,
	0x7d, 0x25, 0xe2, 0x34, 0xec, 0xa4, 0x8a, 0x89, 0x01, 0x0d, 0x98, 0x57, 0x52, 0xf1, 0x29, 0xe0,
	0x3c, 0x1d, 0xd2, 0x24, 0x3e, 0xa7, 0x8a, 0x9d, 0xfb, 0xa5, 0x20, 0xfb, 0x77, 0xc1, 0x8d, 0x09,
	0x99, 0x13, 0xa3, 0xfd, 0x11, 0xaa, 0x52, 0x51, 0x95, 0x4b, 0x72, 0x83, 0x1a, 0xc8, 0x5e, 0x6b,
	0x3d, 0xff, 0xbb, 0xd3, 0xc9, 0x05, 0x70, 0xfa, 0x9a, 0xe5, 0x95, 0x6c, 0x1c, 0x42, 0x6d, 0x6c,
	0xdb, 0x2f, 0x15, 0xbf, 0xa1, 0x46, 0xc5, 0xb6, 0x5a, 0x6f, 0xef, 0xa8, 0xf8, 0x3b, 0x48, 0x4d,
	0x77, 0x53, 0x25, 0xae, 0xbd, 0xf5, 0x60, 0x1a, 0xad, 0x7f, 0x45, 0xb0, 0x7a, 0xc2, 0x84, 0x8c,
	0x79, 0x6a, 0x90, 0x62, 0x15, 0x87, 0x06, 0x20, 0xc8, 0xac, 0x62, 0x79, 0xc5, 0xee, 0xd8, 0xdc,
	0xdc, 0xff, 0x78, 0x7b, 0x0c, 0xab, 0xe6, 0xe4, 0x4b, 0x1d, 0x28, 0xa9, 0xe8, 0xcf, 0xac, 0x18,
	0xd0, 0x84, 0x8c, 0x37, 0x61, 0x81, 0x09, 0xc1, 0x05, 0x99, 0xd7, 0x45, 0x73, 0xa9, 0x5f, 0xc1,
	0xe6, 0x2c, 0x5b, 0xc5, 0x82, 0x5f, 0xb0, 0xeb, 0x72, 0xde, 0xe2, 0x88, 0xf7, 0x61, 0x41, 0xff,
	0x58, 0x3d, 0xaa, 0xd5, 0x7a, 0x75, 0xc7, 0x51, 0xa7, 0xa2, 0xf0, 0x8c, 0xc4, 0x9b, 0xb9, 0xd7,
	0x68, 0xbb, 0x03, 0xd5, 0x32, 0x9f, 0x25, 0x98, 0xef, 0xf6, 0xba, 0x6e, 0xed, 0x1e, 0x5e, 0x81,
	0xa5, 0xcf, 0x47, 0x1f, 0x76, 0x8f, 0x3b, 0xdd, 0x4f, 0x35, 0x84, 0xd7, 0xc1, 0xf2, 0xdc, 0x76,
	0xaf, 0xdb, 0xee, 0x1c, 0x14, 0xc0, 0x1c, 0xb6, 0x60, 0x71, 0xcf, 0xdd, 0x3d, 0x38, 0xde, 0xfb,
	0x52, 0xab, 0xe0, 0x65, 0x58, 0x70, 0x3d, 0xaf, 0xe7, 0xd5, 0xe6, 0xdf, 0xdb, 0x37, 0xb7, 0x5b,
	0xe8, 0xfb, 0xed, 0x16, 0xfa, 0x71, 0xbb, 0x85, 0x4e, 0xeb, 0x66, 0xb0, 0x98, 0x37, 0x69, 0x16,
	0x37, 0xa7, 0xde, 0x8e, 0xb3, 0xaa, 0x7e, 0x32, 0x5e, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x40,
	0x57, 0x4c, 0x0f, 0xa2, 0x04, 0x00, 0x00,
}

func (m *IstioOperatorSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioOperatorSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioOperatorSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ComponentStatus) > 0 {
		for k := range m.ComponentStatus {
			v := m.ComponentStatus[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintOperator(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintOperator(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintOperator(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xc
			i--
			dAtA[i] = 0xca
		}
	}
	if m.Status != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0xc
		i--
		dAtA[i] = 0xc0
	}
	if m.UnvalidatedValues != nil {
		{
			size, err := m.UnvalidatedValues.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xaa
	}
	if m.Values != nil {
		{
			size, err := m.Values.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xa2
	}
	if m.Components != nil {
		{
			size, err := m.Components.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3
		i--
		dAtA[i] = 0x92
	}
	if m.MeshConfig != nil {
		{
			size, err := m.MeshConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2
		i--
		dAtA[i] = 0xc2
	}
	if len(m.ResourceSuffix) > 0 {
		i -= len(m.ResourceSuffix)
		copy(dAtA[i:], m.ResourceSuffix)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.ResourceSuffix)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Hub) > 0 {
		i -= len(m.Hub)
		copy(dAtA[i:], m.Hub)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.Hub)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.InstallPackagePath) > 0 {
		i -= len(m.InstallPackagePath)
		copy(dAtA[i:], m.InstallPackagePath)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.InstallPackagePath)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Profile) > 0 {
		i -= len(m.Profile)
		copy(dAtA[i:], m.Profile)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.Profile)))
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}

func (m *IstioOperatorSpec_VersionStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioOperatorSpec_VersionStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioOperatorSpec_VersionStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StatusString) > 0 {
		i -= len(m.StatusString)
		copy(dAtA[i:], m.StatusString)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.StatusString)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Status != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOperator(dAtA []byte, offset int, v uint64) int {
	offset -= sovOperator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IstioOperatorSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Profile)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	l = len(m.InstallPackagePath)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	l = len(m.Hub)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	l = len(m.ResourceSuffix)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	if m.MeshConfig != nil {
		l = m.MeshConfig.Size()
		n += 2 + l + sovOperator(uint64(l))
	}
	if m.Components != nil {
		l = m.Components.Size()
		n += 2 + l + sovOperator(uint64(l))
	}
	if m.Values != nil {
		l = m.Values.Size()
		n += 2 + l + sovOperator(uint64(l))
	}
	if m.UnvalidatedValues != nil {
		l = m.UnvalidatedValues.Size()
		n += 2 + l + sovOperator(uint64(l))
	}
	if m.Status != 0 {
		n += 2 + sovOperator(uint64(m.Status))
	}
	if len(m.ComponentStatus) > 0 {
		for k, v := range m.ComponentStatus {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovOperator(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovOperator(uint64(len(k))) + l
			n += mapEntrySize + 2 + sovOperator(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IstioOperatorSpec_VersionStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovOperator(uint64(m.Status))
	}
	l = len(m.StatusString)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOperator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOperator(x uint64) (n int) {
	return sovOperator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IstioOperatorSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperator
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IstioOperatorSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioOperatorSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstallPackagePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InstallPackagePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hub", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hub = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceSuffix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceSuffix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 40:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeshConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MeshConfig == nil {
				m.MeshConfig = &MeshConfig{}
			}
			if err := m.MeshConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 50:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Components", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Components == nil {
				m.Components = &IstioComponentSetSpec{}
			}
			if err := m.Components.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Values == nil {
				m.Values = &TypeMapStringInterface{}
			}
			if err := m.Values.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 101:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnvalidatedValues", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UnvalidatedValues == nil {
				m.UnvalidatedValues = &TypeMapStringInterface{}
			}
			if err := m.UnvalidatedValues.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 200:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= IstioOperatorSpec_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 201:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComponentStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ComponentStatus == nil {
				m.ComponentStatus = make(map[string]*IstioOperatorSpec_VersionStatus)
			}
			var mapkey string
			var mapvalue *IstioOperatorSpec_VersionStatus
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOperator
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOperator
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthOperator
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthOperator
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOperator
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthOperator
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthOperator
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &IstioOperatorSpec_VersionStatus{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipOperator(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthOperator
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ComponentStatus[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOperator
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IstioOperatorSpec_VersionStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperator
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VersionStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= IstioOperatorSpec_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOperator
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOperator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperator
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOperator
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthOperator
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOperator
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipOperator(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthOperator
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthOperator = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperator   = fmt.Errorf("proto: integer overflow")
)