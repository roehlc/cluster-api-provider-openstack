/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha8

// OpenStackMachineTemplateResource describes the data needed to create a OpenStackMachine from a template.
type OpenStackMachineTemplateResource struct {
	// Spec is the specification of the desired behavior of the machine.
	Spec OpenStackMachineSpec `json:"spec"`
}

type ImageFilter struct {
	// The ID of the desired image. If this is provided, the other filters will be ignored.
	ID string `json:"id,omitempty"`
	// The name of the desired image. If specified, the combination of name and tags must return a single matching image or an error will be raised.
	Name string `json:"name,omitempty"`
	// The tags associated with the desired image. If specified, the combination of name and tags must return a single matching image or an error will be raised.
	Tags []string `json:"tags,omitempty"`
}

type ExternalRouterIPParam struct {
	// The FixedIP in the corresponding subnet
	FixedIP string `json:"fixedIP,omitempty"`
	// The subnet in which the FixedIP is used for the Gateway of this router
	Subnet SubnetFilter `json:"subnet"`
}

type SecurityGroupFilter struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
	Tags        string `json:"tags,omitempty"`
	TagsAny     string `json:"tagsAny,omitempty"`
	NotTags     string `json:"notTags,omitempty"`
	NotTagsAny  string `json:"notTagsAny,omitempty"`
}

type NetworkFilter struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
	ID          string `json:"id,omitempty"`
	Tags        string `json:"tags,omitempty"`
	TagsAny     string `json:"tagsAny,omitempty"`
	NotTags     string `json:"notTags,omitempty"`
	NotTagsAny  string `json:"notTagsAny,omitempty"`
}

type SubnetFilter struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	ProjectID       string `json:"projectId,omitempty"`
	IPVersion       int    `json:"ipVersion,omitempty"`
	GatewayIP       string `json:"gateway_ip,omitempty"`
	CIDR            string `json:"cidr,omitempty"`
	IPv6AddressMode string `json:"ipv6AddressMode,omitempty"`
	IPv6RAMode      string `json:"ipv6RaMode,omitempty"`
	ID              string `json:"id,omitempty"`
	Tags            string `json:"tags,omitempty"`
	TagsAny         string `json:"tagsAny,omitempty"`
	NotTags         string `json:"notTags,omitempty"`
	NotTagsAny      string `json:"notTagsAny,omitempty"`
}

type RouterFilter struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
	Tags        string `json:"tags,omitempty"`
	TagsAny     string `json:"tagsAny,omitempty"`
	NotTags     string `json:"notTags,omitempty"`
	NotTagsAny  string `json:"notTagsAny,omitempty"`
}

type SubnetSpec struct {
	// CIDR is representing the IP address range used to create the subnet, e.g. 10.0.0.0/24.
	// This field is required when defining a subnet.
	// +required
	CIDR string `json:"cidr"`
	// DNSNameservers holds a list of DNS server addresses that will be provided when creating
	// the subnet. These addresses need to have the same IP version as CIDR.
	DNSNameservers []string `json:"dnsNameservers,omitempty"`
}

type PortOpts struct {
	// Network is a query for an openstack network that the port will be created or discovered on.
	// This will fail if the query returns more than one network.
	Network *NetworkFilter `json:"network,omitempty"`
	// Used to make the name of the port unique. If unspecified, instead the 0-based index of the port in the list is used.
	NameSuffix   string `json:"nameSuffix,omitempty"`
	Description  string `json:"description,omitempty"`
	AdminStateUp *bool  `json:"adminStateUp,omitempty"`
	MACAddress   string `json:"macAddress,omitempty"`
	// Specify pairs of subnet and/or IP address. These should be subnets of the network with the given NetworkID.
	FixedIPs []FixedIP `json:"fixedIPs,omitempty"`
	// The names, uuids, filters or any combination these of the security groups to assign to the instance
	SecurityGroupFilters []SecurityGroupFilter `json:"securityGroupFilters,omitempty"`
	AllowedAddressPairs  []AddressPair         `json:"allowedAddressPairs,omitempty"`
	// Enables and disables trunk at port level. If not provided, openStackMachine.Spec.Trunk is inherited.
	Trunk *bool `json:"trunk,omitempty"`

	// The ID of the host where the port is allocated
	HostID string `json:"hostId,omitempty"`

	// The virtual network interface card (vNIC) type that is bound to the neutron port.
	VNICType string `json:"vnicType,omitempty"`

	// Profile is a set of key-value pairs that are used for binding details.
	// We intentionally don't expose this as a map[string]string because we only want to enable
	// the users to set the values of the keys that are known to work in OpenStack Networking API.
	// See https://docs.openstack.org/api-ref/network/v2/index.html?expanded=create-port-detail#create-port
	Profile BindingProfile `json:"profile,omitempty"`

	// DisablePortSecurity enables or disables the port security when set.
	// When not set, it takes the value of the corresponding field at the network level.
	DisablePortSecurity *bool `json:"disablePortSecurity,omitempty"`

	// PropageteUplinkStatus enables or disables the propagate uplink status on the port.
	PropagateUplinkStatus *bool `json:"propagateUplinkStatus,omitempty"`

	// Tags applied to the port (and corresponding trunk, if a trunk is configured.)
	// These tags are applied in addition to the instance's tags, which will also be applied to the port.
	// +listType=set
	Tags []string `json:"tags,omitempty"`

	// Value specs are extra parameters to include in the API request with OpenStack.
	// This is an extension point for the API, so what they do and if they are supported,
	// depends on the specific OpenStack implementation.
	// +optional
	// +listType=map
	// +listMapKey=name
	ValueSpecs []ValueSpec `json:"valueSpecs,omitempty"`
}

type BindingProfile struct {
	// OVSHWOffload enables or disables the OVS hardware offload feature.
	OVSHWOffload bool `json:"ovsHWOffload,omitempty"`

	// TrustedVF enables or disables the “trusted mode” for the VF.
	TrustedVF bool `json:"trustedVF,omitempty"`
}

type FixedIP struct {
	// Subnet is an openstack subnet query that will return the id of a subnet to create
	// the fixed IP of a port in. This query must not return more than one subnet.
	Subnet    *SubnetFilter `json:"subnet"`
	IPAddress string        `json:"ipAddress,omitempty"`
}

type AddressPair struct {
	IPAddress  string `json:"ipAddress,omitempty"`
	MACAddress string `json:"macAddress,omitempty"`
}

type BastionStatus struct {
	ID                  string                     `json:"id,omitempty"`
	Name                string                     `json:"name,omitempty"`
	SSHKeyName          string                     `json:"sshKeyName,omitempty"`
	State               InstanceState              `json:"state,omitempty"`
	IP                  string                     `json:"ip,omitempty"`
	FloatingIP          string                     `json:"floatingIP,omitempty"`
	ReferencedResources ReferencedMachineResources `json:"referencedResources,omitempty"`
}

type RootVolume struct {
	Size             int    `json:"diskSize,omitempty"`
	VolumeType       string `json:"volumeType,omitempty"`
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// BlockDeviceStorage is the storage type of a block device to create and
// contains additional storage options.
// +union
//
//nolint:godot
type BlockDeviceStorage struct {
	// Type is the type of block device to create.
	// This can be either "Volume" or "Local".
	// +unionDiscriminator
	Type BlockDeviceType `json:"type"`

	// Volume contains additional storage options for a volume block device.
	// +optional
	// +unionMember,optional
	Volume *BlockDeviceVolume `json:"volume,omitempty"`
}

// BlockDeviceVolume contains additional storage options for a volume block device.
type BlockDeviceVolume struct {
	// Type is the Cinder volume type of the volume.
	// If omitted, the default Cinder volume type that is configured in the OpenStack cloud
	// will be used.
	// +optional
	Type string `json:"type,omitempty"`

	// AvailabilityZone is the volume availability zone to create the volume in.
	// If omitted, the availability zone of the server will be used.
	// The availability zone must NOT contain spaces otherwise it will lead to volume that belongs
	// to this availability zone register failure, see kubernetes/cloud-provider-openstack#1379 for
	// further information.
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// AdditionalBlockDevice is a block device to attach to the server.
type AdditionalBlockDevice struct {
	// Name of the block device in the context of a machine.
	// If the block device is a volume, the Cinder volume will be named
	// as a combination of the machine name and this name.
	// Also, this name will be used for tagging the block device.
	// Information about the block device tag can be obtained from the OpenStack
	// metadata API or the config drive.
	Name string `json:"name"`

	// SizeGiB is the size of the block device in gibibytes (GiB).
	SizeGiB int `json:"sizeGiB"`

	// Storage specifies the storage type of the block device and
	// additional storage options.
	Storage BlockDeviceStorage `json:"storage"`
}

type ServerGroupFilter struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// BlockDeviceType defines the type of block device to create.
type BlockDeviceType string

const (
	// LocalBlockDevice is an ephemeral block device attached to the server.
	LocalBlockDevice BlockDeviceType = "Local"

	// VolumeBlockDevice is a volume block device attached to the server.
	VolumeBlockDevice BlockDeviceType = "Volume"
)

// NetworkStatus contains basic information about an existing neutron network.
type NetworkStatus struct {
	Name string `json:"name"`
	ID   string `json:"id"`

	//+optional
	Tags []string `json:"tags,omitempty"`
}

// NetworkStatusWithSubnets represents basic information about an existing neutron network and an associated set of subnets.
type NetworkStatusWithSubnets struct {
	NetworkStatus `json:",inline"`

	// Subnets is a list of subnets associated with the default cluster network. Machines which use the default cluster network will get an address from all of these subnets.
	Subnets []Subnet `json:"subnets,omitempty"`
}

// Subnet represents basic information about the associated OpenStack Neutron Subnet.
type Subnet struct {
	Name string `json:"name"`
	ID   string `json:"id"`

	CIDR string `json:"cidr"`

	//+optional
	Tags []string `json:"tags,omitempty"`
}

// Router represents basic information about the associated OpenStack Neutron Router.
type Router struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	//+optional
	Tags []string `json:"tags,omitempty"`
	//+optional
	IPs []string `json:"ips,omitempty"`
}

// LoadBalancer represents basic information about the associated OpenStack LoadBalancer.
type LoadBalancer struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	IP         string `json:"ip"`
	InternalIP string `json:"internalIP"`
	//+optional
	AllowedCIDRs []string `json:"allowedCIDRs,omitempty"`
	//+optional
	Tags []string `json:"tags,omitempty"`
}

// SecurityGroupStatus represents the basic information of the associated
// OpenStack Neutron Security Group.
type SecurityGroupStatus struct {
	// name of the security group
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// id of the security group
	// +kubebuilder:validation:Required
	ID string `json:"id"`

	// list of security group rules
	// +optional
	Rules []SecurityGroupRuleStatus `json:"rules,omitempty"`
}

// SecurityGroupRuleSpec represent the basic information of the associated OpenStack
// Security Group Role.
// For now this is only used for the allNodesSecurityGroupRules but when we add
// other security groups, we'll need to add a validation because
// Remote* fields are mutually exclusive.
type SecurityGroupRuleSpec struct {
	// name of the security group rule.
	// It's used to identify the rule so it can be patched and will not be sent to the OpenStack API.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// description of the security group rule.
	// +optional
	Description *string `json:"description,omitempty"`

	// direction in which the security group rule is applied. The only values
	// allowed are "ingress" or "egress". For a compute instance, an ingress
	// security group rule is applied to incoming (ingress) traffic for that
	// instance. An egress rule is applied to traffic leaving the instance.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:enum=ingress;egress
	Direction string `json:"direction"`

	// etherType must be IPv4 or IPv6, and addresses represented in CIDR must match the
	// ingress or egress rules.
	// +kubebuilder:validation:enum=IPv4;IPv6
	// +optional
	EtherType *string `json:"etherType,omitempty"`

	// portRangeMin is a number in the range that is matched by the security group
	// rule. If the protocol is TCP or UDP, this value must be less than or equal
	// to the value of the portRangeMax attribute.
	// +optional
	PortRangeMin *int `json:"portRangeMin,omitempty"`

	// portRangeMax is a number in the range that is matched by the security group
	// rule. The portRangeMin attribute constrains the portRangeMax attribute.
	// +optional
	PortRangeMax *int `json:"portRangeMax,omitempty"`

	// protocol is the protocol that is matched by the security group rule.
	// +optional
	Protocol *string `json:"protocol,omitempty"`

	// remoteGroupID is the remote group ID to be associated with this security group rule.
	// You can specify either remoteGroupID or remoteIPPrefix or remoteManagedGroups.
	// +optional
	RemoteGroupID *string `json:"remoteGroupID,omitempty"`

	// remoteIPPrefix is the remote IP prefix to be associated with this security group rule.
	// You can specify either remoteGroupID or remoteIPPrefix or remoteManagedGroups.
	// +optional
	RemoteIPPrefix *string `json:"remoteIPPrefix,omitempty"`

	// remoteManagedGroups is the remote managed groups to be associated with this security group rule.
	// You can specify either remoteGroupID or remoteIPPrefix or remoteManagedGroups.
	// +optional
	RemoteManagedGroups []ManagedSecurityGroupName `json:"remoteManagedGroups,omitempty"`
}

type SecurityGroupRuleStatus struct {
	// id of the security group rule
	// +kubebuilder:validation:Required
	ID string `json:"id"`

	// description of the security group rule.
	// +optional
	Description *string `json:"description,omitempty"`

	// direction in which the security group rule is applied. The only values
	// allowed are "ingress" or "egress". For a compute instance, an ingress
	// security group rule is applied to incoming (ingress) traffic for that
	// instance. An egress rule is applied to traffic leaving the instance.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:enum=ingress;egress
	Direction string `json:"direction"`

	// etherType must be IPv4 or IPv6, and addresses represented in CIDR must match the
	// ingress or egress rules.
	// +kubebuilder:validation:enum=IPv4;IPv6
	// +optional
	EtherType *string `json:"etherType,omitempty"`

	// portRangeMin is a number in the range that is matched by the security group
	// rule. If the protocol is TCP or UDP, this value must be less than or equal
	// to the value of the portRangeMax attribute.
	// +optional
	PortRangeMin *int `json:"portRangeMin,omitempty"`

	// portRangeMax is a number in the range that is matched by the security group
	// rule. The portRangeMin attribute constrains the portRangeMax attribute.
	// +optional
	PortRangeMax *int `json:"portRangeMax,omitempty"`

	// protocol is the protocol that is matched by the security group rule.
	// +optional
	Protocol *string `json:"protocol,omitempty"`

	// remoteGroupID is the remote group ID to be associated with this security group rule.
	// You can specify either remoteGroupID or remoteIPPrefix or remoteManagedGroups.
	// +optional
	RemoteGroupID *string `json:"remoteGroupID,omitempty"`

	// remoteIPPrefix is the remote IP prefix to be associated with this security group rule.
	// You can specify either remoteGroupID or remoteIPPrefix or remoteManagedGroups.
	// +optional
	RemoteIPPrefix *string `json:"remoteIPPrefix,omitempty"`
}

// +kubebuilder:validation:Enum=bastion;controlplane;worker
type ManagedSecurityGroupName string

func (m ManagedSecurityGroupName) String() string {
	return string(m)
}

// InstanceState describes the state of an OpenStack instance.
type InstanceState string

var (
	// InstanceStateBuild is the string representing an instance in a build state.
	InstanceStateBuild = InstanceState("BUILD")

	// InstanceStateActive is the string representing an instance in an active state.
	InstanceStateActive = InstanceState("ACTIVE")

	// InstanceStateError is the string representing an instance in an error state.
	InstanceStateError = InstanceState("ERROR")

	// InstanceStateStopped is the string representing an instance in a stopped state.
	InstanceStateStopped = InstanceState("STOPPED")

	// InstanceStateShutoff is the string representing an instance in a shutoff state.
	InstanceStateShutoff = InstanceState("SHUTOFF")

	// InstanceStateDeleted is the string representing an instance in a deleted state.
	InstanceStateDeleted = InstanceState("DELETED")

	// InstanceStateUndefined is the string representing an undefined instance state.
	InstanceStateUndefined = InstanceState("")
)

// Bastion represents basic information about the bastion node.
type Bastion struct {
	//+optional
	Enabled bool `json:"enabled"`

	// Instance for the bastion itself
	Instance OpenStackMachineSpec `json:"instance,omitempty"`

	//+optional
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// FloatingIP which will be associated to the bastion machine.
	// The floating IP should already exist and should not be associated with a port.
	//+optional
	FloatingIP string `json:"floatingIP,omitempty"`
}

type APIServerLoadBalancer struct {
	// Enabled defines whether a load balancer should be created.
	Enabled bool `json:"enabled,omitempty"`
	// AdditionalPorts adds additional tcp ports to the load balancer.
	AdditionalPorts []int `json:"additionalPorts,omitempty"`
	// AllowedCIDRs restrict access to all API-Server listeners to the given address CIDRs.
	AllowedCIDRs []string `json:"allowedCidrs,omitempty"`
	// Octavia Provider Used to create load balancer
	Provider string `json:"provider,omitempty"`
}

// ReferencedMachineResources contains resolved references to resources required by the machine.
type ReferencedMachineResources struct {
	// ServerGroupID is the ID of the server group the machine should be added to and is calculated based on ServerGroupFilter.
	// +optional
	ServerGroupID string `json:"serverGroupID,omitempty"`

	// ImageID is the ID of the image to use for the machine and is calculated based on ImageFilter.
	// +optional
	ImageID string `json:"imageID,omitempty"`
}

// ValueSpec represents a single value_spec key-value pair.
type ValueSpec struct {
	// Name is the name of the key-value pair.
	// This is just for identifying the pair and will not be sent to the OpenStack API.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Key is the key in the key-value pair.
	// +kubebuilder:validation:Required
	Key string `json:"key"`
	// Value is the value in the key-value pair.
	// +kubebuilder:validation:Required
	Value string `json:"value"`
}
