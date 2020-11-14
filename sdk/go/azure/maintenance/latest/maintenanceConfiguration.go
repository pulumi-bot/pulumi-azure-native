// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Maintenance configuration record type
type MaintenanceConfiguration struct {
	pulumi.CustomResourceState

	// Gets or sets extensionProperties of the maintenanceConfiguration. This is for future use only and would be a set of key value pairs for additional information e.g. whether to follow SDP etc.
	ExtensionProperties pulumi.StringMapOutput `pulumi:"extensionProperties"`
	// Gets or sets location of the resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets or sets maintenanceScope of the configuration. It represent the impact area of the maintenance
	MaintenanceScope pulumi.StringPtrOutput `pulumi:"maintenanceScope"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets namespace of the resource e.g. Microsoft.Maintenance or Microsoft.Sql
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Gets or sets tags of the resource
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMaintenanceConfiguration registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceConfiguration(ctx *pulumi.Context,
	name string, args *MaintenanceConfigurationArgs, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &MaintenanceConfigurationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20180601preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20200401:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20200701preview:MaintenanceConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource MaintenanceConfiguration
	err := ctx.RegisterResource("azure-nextgen:maintenance/latest:MaintenanceConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceConfiguration gets an existing MaintenanceConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceConfigurationState, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	var resource MaintenanceConfiguration
	err := ctx.ReadResource("azure-nextgen:maintenance/latest:MaintenanceConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceConfiguration resources.
type maintenanceConfigurationState struct {
	// Gets or sets extensionProperties of the maintenanceConfiguration. This is for future use only and would be a set of key value pairs for additional information e.g. whether to follow SDP etc.
	ExtensionProperties map[string]string `pulumi:"extensionProperties"`
	// Gets or sets location of the resource
	Location *string `pulumi:"location"`
	// Gets or sets maintenanceScope of the configuration. It represent the impact area of the maintenance
	MaintenanceScope *string `pulumi:"maintenanceScope"`
	// Name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets namespace of the resource e.g. Microsoft.Maintenance or Microsoft.Sql
	Namespace *string `pulumi:"namespace"`
	// Gets or sets tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type *string `pulumi:"type"`
}

type MaintenanceConfigurationState struct {
	// Gets or sets extensionProperties of the maintenanceConfiguration. This is for future use only and would be a set of key value pairs for additional information e.g. whether to follow SDP etc.
	ExtensionProperties pulumi.StringMapInput
	// Gets or sets location of the resource
	Location pulumi.StringPtrInput
	// Gets or sets maintenanceScope of the configuration. It represent the impact area of the maintenance
	MaintenanceScope pulumi.StringPtrInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets namespace of the resource e.g. Microsoft.Maintenance or Microsoft.Sql
	Namespace pulumi.StringPtrInput
	// Gets or sets tags of the resource
	Tags pulumi.StringMapInput
	// Type of the resource
	Type pulumi.StringPtrInput
}

func (MaintenanceConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceConfigurationState)(nil)).Elem()
}

type maintenanceConfigurationArgs struct {
	// Gets or sets extensionProperties of the maintenanceConfiguration. This is for future use only and would be a set of key value pairs for additional information e.g. whether to follow SDP etc.
	ExtensionProperties map[string]string `pulumi:"extensionProperties"`
	// Gets or sets location of the resource
	Location *string `pulumi:"location"`
	// Gets or sets maintenanceScope of the configuration. It represent the impact area of the maintenance
	MaintenanceScope *string `pulumi:"maintenanceScope"`
	// Gets or sets namespace of the resource e.g. Microsoft.Maintenance or Microsoft.Sql
	Namespace *string `pulumi:"namespace"`
	// Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource Identifier
	ResourceName string `pulumi:"resourceName"`
	// Gets or sets tags of the resource
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MaintenanceConfiguration resource.
type MaintenanceConfigurationArgs struct {
	// Gets or sets extensionProperties of the maintenanceConfiguration. This is for future use only and would be a set of key value pairs for additional information e.g. whether to follow SDP etc.
	ExtensionProperties pulumi.StringMapInput
	// Gets or sets location of the resource
	Location pulumi.StringPtrInput
	// Gets or sets maintenanceScope of the configuration. It represent the impact area of the maintenance
	MaintenanceScope pulumi.StringPtrInput
	// Gets or sets namespace of the resource e.g. Microsoft.Maintenance or Microsoft.Sql
	Namespace pulumi.StringPtrInput
	// Resource Group Name
	ResourceGroupName pulumi.StringInput
	// Resource Identifier
	ResourceName pulumi.StringInput
	// Gets or sets tags of the resource
	Tags pulumi.StringMapInput
}

func (MaintenanceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceConfigurationArgs)(nil)).Elem()
}

type MaintenanceConfigurationInput interface {
	pulumi.Input

	ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput
	ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput
}

func (MaintenanceConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceConfiguration)(nil)).Elem()
}

func (i MaintenanceConfiguration) ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput {
	return i.ToMaintenanceConfigurationOutputWithContext(context.Background())
}

func (i MaintenanceConfiguration) ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceConfigurationOutput)
}

type MaintenanceConfigurationOutput struct {
	*pulumi.OutputState
}

func (MaintenanceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceConfigurationOutput)(nil)).Elem()
}

func (o MaintenanceConfigurationOutput) ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput {
	return o
}

func (o MaintenanceConfigurationOutput) ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MaintenanceConfigurationOutput{})
}
