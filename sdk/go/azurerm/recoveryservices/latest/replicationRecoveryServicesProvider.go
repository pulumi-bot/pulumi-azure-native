// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provider details.
//
// ## Example Usage
// ### Adds a recovery services provider.
//
// ```go
// package main
//
// import (
// 	recoveryservices "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/recoveryservices/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := recoveryservices.NewReplicationRecoveryServicesProvider(ctx, "replicationRecoveryServicesProvider", &recoveryservices.ReplicationRecoveryServicesProviderArgs{
// 			FabricName:        pulumi.String("vmwarefabric1"),
// 			ProviderName:      pulumi.String("vmwareprovider1"),
// 			ResourceGroupName: pulumi.String("resourcegroup1"),
// 			ResourceName:      pulumi.String("migrationvault"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ReplicationRecoveryServicesProvider struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provider properties.
	Properties RecoveryServicesProviderPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationRecoveryServicesProvider registers a new resource with the given unique name, arguments, and options.
func NewReplicationRecoveryServicesProvider(ctx *pulumi.Context,
	name string, args *ReplicationRecoveryServicesProviderArgs, opts ...pulumi.ResourceOption) (*ReplicationRecoveryServicesProvider, error) {
	if args == nil || args.FabricName == nil {
		return nil, errors.New("missing required argument 'FabricName'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ProviderName == nil {
		return nil, errors.New("missing required argument 'ProviderName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &ReplicationRecoveryServicesProviderArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:recoveryservices/v20180110:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20180710:ReplicationRecoveryServicesProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationRecoveryServicesProvider
	err := ctx.RegisterResource("azurerm:recoveryservices/latest:ReplicationRecoveryServicesProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationRecoveryServicesProvider gets an existing ReplicationRecoveryServicesProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationRecoveryServicesProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationRecoveryServicesProviderState, opts ...pulumi.ResourceOption) (*ReplicationRecoveryServicesProvider, error) {
	var resource ReplicationRecoveryServicesProvider
	err := ctx.ReadResource("azurerm:recoveryservices/latest:ReplicationRecoveryServicesProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationRecoveryServicesProvider resources.
type replicationRecoveryServicesProviderState struct {
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Provider properties.
	Properties *RecoveryServicesProviderPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type *string `pulumi:"type"`
}

type ReplicationRecoveryServicesProviderState struct {
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// Provider properties.
	Properties RecoveryServicesProviderPropertiesResponsePtrInput
	// Resource Type
	Type pulumi.StringPtrInput
}

func (ReplicationRecoveryServicesProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryServicesProviderState)(nil)).Elem()
}

type replicationRecoveryServicesProviderArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// The properties of an add provider request.
	Properties AddRecoveryServicesProviderInputProperties `pulumi:"properties"`
	// Recovery services provider name.
	ProviderName string `pulumi:"providerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationRecoveryServicesProvider resource.
type ReplicationRecoveryServicesProviderArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput
	// The properties of an add provider request.
	Properties AddRecoveryServicesProviderInputPropertiesInput
	// Recovery services provider name.
	ProviderName pulumi.StringInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationRecoveryServicesProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryServicesProviderArgs)(nil)).Elem()
}
