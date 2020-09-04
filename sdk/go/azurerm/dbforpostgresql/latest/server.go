// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a server.
type Server struct {
	pulumi.CustomResourceState

	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrOutput `pulumi:"administratorLogin"`
	// Status showing whether the server data encryption is enabled with customer-managed keys.
	ByokEnforcement pulumi.StringOutput `pulumi:"byokEnforcement"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate pulumi.StringPtrOutput `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringPtrOutput `pulumi:"fullyQualifiedDomainName"`
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption pulumi.StringPtrOutput `pulumi:"infrastructureEncryption"`
	// The location the resource resides in.
	Location pulumi.StringOutput `pulumi:"location"`
	// The master server id of a replica server.
	MasterServerId pulumi.StringPtrOutput `pulumi:"masterServerId"`
	// Enforce a minimal Tls version for the server.
	MinimalTlsVersion pulumi.StringPtrOutput `pulumi:"minimalTlsVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections on a server
	PrivateEndpointConnections ServerPrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The maximum number of replicas that a master server can have.
	ReplicaCapacity pulumi.IntPtrOutput `pulumi:"replicaCapacity"`
	// The replication role of the server.
	ReplicationRole pulumi.StringPtrOutput `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement pulumi.StringPtrOutput `pulumi:"sslEnforcement"`
	// Storage profile of a server.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Application-specific metadata in the form of key-value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// A state of a server that is visible to user.
	UserVisibleState pulumi.StringPtrOutput `pulumi:"userVisibleState"`
	// Server version.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil {
		args = &ServerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:dbforpostgresql/v20171201:Server"),
		},
		{
			Type: pulumi.String("azurerm:dbforpostgresql/v20171201preview:Server"),
		},
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azurerm:dbforpostgresql/latest:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azurerm:dbforpostgresql/latest:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// Status showing whether the server data encryption is enabled with customer-managed keys.
	ByokEnforcement *string `pulumi:"byokEnforcement"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `pulumi:"fullyQualifiedDomainName"`
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption *string `pulumi:"infrastructureEncryption"`
	// The location the resource resides in.
	Location *string `pulumi:"location"`
	// The master server id of a replica server.
	MasterServerId *string `pulumi:"masterServerId"`
	// Enforce a minimal Tls version for the server.
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The maximum number of replicas that a master server can have.
	ReplicaCapacity *int `pulumi:"replicaCapacity"`
	// The replication role of the server.
	ReplicationRole *string `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku *SkuResponse `pulumi:"sku"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Storage profile of a server.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// A state of a server that is visible to user.
	UserVisibleState *string `pulumi:"userVisibleState"`
	// Server version.
	Version *string `pulumi:"version"`
}

type ServerState struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrInput
	// Status showing whether the server data encryption is enabled with customer-managed keys.
	ByokEnforcement pulumi.StringPtrInput
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate pulumi.StringPtrInput
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringPtrInput
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityResponsePtrInput
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption pulumi.StringPtrInput
	// The location the resource resides in.
	Location pulumi.StringPtrInput
	// The master server id of a replica server.
	MasterServerId pulumi.StringPtrInput
	// Enforce a minimal Tls version for the server.
	MinimalTlsVersion pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// List of private endpoint connections on a server
	PrivateEndpointConnections ServerPrivateEndpointConnectionResponseArrayInput
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringPtrInput
	// The maximum number of replicas that a master server can have.
	ReplicaCapacity pulumi.IntPtrInput
	// The replication role of the server.
	ReplicationRole pulumi.StringPtrInput
	// The SKU (pricing tier) of the server.
	Sku SkuResponsePtrInput
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement pulumi.StringPtrInput
	// Storage profile of a server.
	StorageProfile StorageProfileResponsePtrInput
	// Application-specific metadata in the form of key-value pairs.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// A state of a server that is visible to user.
	UserVisibleState pulumi.StringPtrInput
	// Server version.
	Version pulumi.StringPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentity `pulumi:"identity"`
	// The location the resource resides in.
	Location string `pulumi:"location"`
	// Properties of the server.
	Properties ServerPropertiesForCreate `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The SKU (pricing tier) of the server.
	Sku *Sku `pulumi:"sku"`
	// Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityPtrInput
	// The location the resource resides in.
	Location pulumi.StringInput
	// Properties of the server.
	Properties ServerPropertiesForCreateInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The SKU (pricing tier) of the server.
	Sku SkuPtrInput
	// Application-specific metadata in the form of key-value pairs.
	Tags pulumi.StringMapInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}
