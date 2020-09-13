// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701privatepreview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a server.
//
// ## Example Usage
// ### Create a new server
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	dbformysql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/dbformysql/v20200701privatepreview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dbformysql.NewServer(ctx, "server", &dbformysql.ServerArgs{
// 			AdministratorLogin:         pulumi.String("cloudsa"),
// 			AdministratorLoginPassword: pulumi.String(fmt.Sprintf("%v%v%v", "pass", "$", "w0rd")),
// 			CreateMode:                 pulumi.String("Default"),
// 			Location:                   pulumi.String("westus"),
// 			ResourceGroupName:          pulumi.String("testrg"),
// 			ServerName:                 pulumi.String("mysqltestsvc4"),
// 			Sku: &dbformysql.SkuArgs{
// 				Name: pulumi.String("Standard_D14_v2"),
// 				Tier: pulumi.String("GeneralPurpose"),
// 			},
// 			SslEnforcement: pulumi.String("Enabled"),
// 			StorageProfile: &dbformysql.StorageProfileArgs{
// 				BackupRetentionDays: pulumi.Int(7),
// 				StorageIops:         pulumi.Int(200),
// 				StorageMB:           pulumi.Int(128000),
// 			},
// 			Tags: pulumi.StringMap{
// 				"ElasticServer": pulumi.String("1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a replica server
//
// ```go
// package main
//
// import (
// 	dbformysql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/dbformysql/v20200701privatepreview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dbformysql.NewServer(ctx, "server", &dbformysql.ServerArgs{
// 			CreateMode:        pulumi.String("Replica"),
// 			Location:          pulumi.String("westus"),
// 			ResourceGroupName: pulumi.String("TargetResourceGroup"),
// 			ServerName:        pulumi.String("targetserver"),
// 			SourceServerId:    pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/PrimaryResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/primaryserver"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a server as a point in time restore
//
// ```go
// package main
//
// import (
// 	dbformysql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/dbformysql/v20200701privatepreview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dbformysql.NewServer(ctx, "server", &dbformysql.ServerArgs{
// 			CreateMode:         pulumi.String("PointInTimeRestore"),
// 			Location:           pulumi.String("brazilsouth"),
// 			ResourceGroupName:  pulumi.String("TargetResourceGroup"),
// 			RestorePointInTime: pulumi.String("2017-12-14T00:00:37.467Z"),
// 			ServerName:         pulumi.String("targetserver"),
// 			Sku: &dbformysql.SkuArgs{
// 				Name: pulumi.String("Standard_D14_v2"),
// 				Tier: pulumi.String("GeneralPurpose"),
// 			},
// 			SourceServerId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/sourceserver"),
// 			Tags: pulumi.StringMap{
// 				"ElasticServer": pulumi.String("1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Server struct {
	pulumi.CustomResourceState

	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrOutput `pulumi:"administratorLogin"`
	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword pulumi.StringPtrOutput `pulumi:"administratorLoginPassword"`
	// availability Zone information of the server.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// Status showing whether the data encryption is enabled with customer-managed keys.
	ByokEnforcement pulumi.StringOutput `pulumi:"byokEnforcement"`
	// The mode to create a new MySQL server.
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// Delegated subnet arguments.
	DelegatedSubnetArguments DelegatedSubnetArgumentsResponsePtrOutput `pulumi:"delegatedSubnetArguments"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate pulumi.StringOutput `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// Enable HA or not for a server.
	HaEnabled pulumi.StringPtrOutput `pulumi:"haEnabled"`
	// The state of a HA server.
	HaState pulumi.StringOutput `pulumi:"haState"`
	// The Azure Active Directory identity of the server.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption pulumi.StringPtrOutput `pulumi:"infrastructureEncryption"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowResponsePtrOutput `pulumi:"maintenanceWindow"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringOutput `pulumi:"publicNetworkAccess"`
	// The maximum number of replicas that a primary server can have.
	ReplicaCapacity pulumi.IntOutput `pulumi:"replicaCapacity"`
	// The replication role.
	ReplicationRole pulumi.StringPtrOutput `pulumi:"replicationRole"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime pulumi.StringPtrOutput `pulumi:"restorePointInTime"`
	// The SKU (pricing tier) of the server.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The source MySQL server id.
	SourceServerId pulumi.StringPtrOutput `pulumi:"sourceServerId"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement pulumi.StringPtrOutput `pulumi:"sslEnforcement"`
	// availability Zone information of the server.
	StandbyAvailabilityZone pulumi.StringOutput `pulumi:"standbyAvailabilityZone"`
	// The state of a server.
	State pulumi.StringOutput `pulumi:"state"`
	// Storage profile of a server.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// Server version.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
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
	var resource Server
	err := ctx.RegisterResource("azurerm:dbformysql/v20200701privatepreview:Server", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:dbformysql/v20200701privatepreview:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// availability Zone information of the server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Status showing whether the data encryption is enabled with customer-managed keys.
	ByokEnforcement *string `pulumi:"byokEnforcement"`
	// The mode to create a new MySQL server.
	CreateMode *string `pulumi:"createMode"`
	// Delegated subnet arguments.
	DelegatedSubnetArguments *DelegatedSubnetArgumentsResponse `pulumi:"delegatedSubnetArguments"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `pulumi:"fullyQualifiedDomainName"`
	// Enable HA or not for a server.
	HaEnabled *string `pulumi:"haEnabled"`
	// The state of a HA server.
	HaState *string `pulumi:"haState"`
	// The Azure Active Directory identity of the server.
	Identity *IdentityResponse `pulumi:"identity"`
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption *string `pulumi:"infrastructureEncryption"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindowResponse `pulumi:"maintenanceWindow"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The maximum number of replicas that a primary server can have.
	ReplicaCapacity *int `pulumi:"replicaCapacity"`
	// The replication role.
	ReplicationRole *string `pulumi:"replicationRole"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The SKU (pricing tier) of the server.
	Sku *SkuResponse `pulumi:"sku"`
	// The source MySQL server id.
	SourceServerId *string `pulumi:"sourceServerId"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// availability Zone information of the server.
	StandbyAvailabilityZone *string `pulumi:"standbyAvailabilityZone"`
	// The state of a server.
	State *string `pulumi:"state"`
	// Storage profile of a server.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// Server version.
	Version *string `pulumi:"version"`
}

type ServerState struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrInput
	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword pulumi.StringPtrInput
	// availability Zone information of the server.
	AvailabilityZone pulumi.StringPtrInput
	// Status showing whether the data encryption is enabled with customer-managed keys.
	ByokEnforcement pulumi.StringPtrInput
	// The mode to create a new MySQL server.
	CreateMode pulumi.StringPtrInput
	// Delegated subnet arguments.
	DelegatedSubnetArguments DelegatedSubnetArgumentsResponsePtrInput
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate pulumi.StringPtrInput
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringPtrInput
	// Enable HA or not for a server.
	HaEnabled pulumi.StringPtrInput
	// The state of a HA server.
	HaState pulumi.StringPtrInput
	// The Azure Active Directory identity of the server.
	Identity IdentityResponsePtrInput
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowResponsePtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringPtrInput
	// The maximum number of replicas that a primary server can have.
	ReplicaCapacity pulumi.IntPtrInput
	// The replication role.
	ReplicationRole pulumi.StringPtrInput
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime pulumi.StringPtrInput
	// The SKU (pricing tier) of the server.
	Sku SkuResponsePtrInput
	// The source MySQL server id.
	SourceServerId pulumi.StringPtrInput
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement pulumi.StringPtrInput
	// availability Zone information of the server.
	StandbyAvailabilityZone pulumi.StringPtrInput
	// The state of a server.
	State pulumi.StringPtrInput
	// Storage profile of a server.
	StorageProfile StorageProfileResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// Server version.
	Version pulumi.StringPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// availability Zone information of the server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The mode to create a new MySQL server.
	CreateMode *string `pulumi:"createMode"`
	// Delegated subnet arguments.
	DelegatedSubnetArguments *DelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	// Enable HA or not for a server.
	HaEnabled *string `pulumi:"haEnabled"`
	// The Azure Active Directory identity of the server.
	Identity *Identity `pulumi:"identity"`
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption *string `pulumi:"infrastructureEncryption"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// The replication role.
	ReplicationRole *string `pulumi:"replicationRole"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The SKU (pricing tier) of the server.
	Sku *Sku `pulumi:"sku"`
	// The source MySQL server id.
	SourceServerId *string `pulumi:"sourceServerId"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Storage profile of a server.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Server version.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrInput
	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword pulumi.StringPtrInput
	// availability Zone information of the server.
	AvailabilityZone pulumi.StringPtrInput
	// The mode to create a new MySQL server.
	CreateMode pulumi.StringPtrInput
	// Delegated subnet arguments.
	DelegatedSubnetArguments DelegatedSubnetArgumentsPtrInput
	// Enable HA or not for a server.
	HaEnabled pulumi.StringPtrInput
	// The Azure Active Directory identity of the server.
	Identity IdentityPtrInput
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Maintenance window of a server.
	MaintenanceWindow MaintenanceWindowPtrInput
	// The replication role.
	ReplicationRole pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The SKU (pricing tier) of the server.
	Sku SkuPtrInput
	// The source MySQL server id.
	SourceServerId pulumi.StringPtrInput
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement pulumi.StringPtrInput
	// Storage profile of a server.
	StorageProfile StorageProfilePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Server version.
	Version pulumi.StringPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}
