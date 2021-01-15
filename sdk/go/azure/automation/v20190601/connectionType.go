// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the connection type.
type ConnectionType struct {
	pulumi.CustomResourceState

	// Gets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets the field definitions of the connection type.
	FieldDefinitions FieldDefinitionResponseMapOutput `pulumi:"fieldDefinitions"`
	// Gets or sets a Boolean value to indicate if the connection type is global.
	IsGlobal pulumi.BoolPtrOutput `pulumi:"isGlobal"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// Gets the name of the connection type.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectionType registers a new resource with the given unique name, arguments, and options.
func NewConnectionType(ctx *pulumi.Context,
	name string, args *ConnectionTypeArgs, opts ...pulumi.ResourceOption) (*ConnectionType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ConnectionTypeName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionTypeName'")
	}
	if args.FieldDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'FieldDefinitions'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation/latest:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20151031:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20200113preview:ConnectionType"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectionType
	err := ctx.RegisterResource("azure-nextgen:automation/v20190601:ConnectionType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionType gets an existing ConnectionType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionTypeState, opts ...pulumi.ResourceOption) (*ConnectionType, error) {
	var resource ConnectionType
	err := ctx.ReadResource("azure-nextgen:automation/v20190601:ConnectionType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionType resources.
type connectionTypeState struct {
	// Gets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets the field definitions of the connection type.
	FieldDefinitions map[string]FieldDefinitionResponse `pulumi:"fieldDefinitions"`
	// Gets or sets a Boolean value to indicate if the connection type is global.
	IsGlobal *bool `pulumi:"isGlobal"`
	// Gets or sets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// Gets the name of the connection type.
	Name *string `pulumi:"name"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ConnectionTypeState struct {
	// Gets the creation time.
	CreationTime pulumi.StringPtrInput
	// Gets or sets the description.
	Description pulumi.StringPtrInput
	// Gets the field definitions of the connection type.
	FieldDefinitions FieldDefinitionResponseMapInput
	// Gets or sets a Boolean value to indicate if the connection type is global.
	IsGlobal pulumi.BoolPtrInput
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrInput
	// Gets the name of the connection type.
	Name pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ConnectionTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionTypeState)(nil)).Elem()
}

type connectionTypeArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The parameters supplied to the create or update connection type operation.
	ConnectionTypeName string `pulumi:"connectionTypeName"`
	// Gets or sets the field definitions of the connection type.
	FieldDefinitions map[string]FieldDefinition `pulumi:"fieldDefinitions"`
	// Gets or sets a Boolean value to indicate if the connection type is global.
	IsGlobal *bool `pulumi:"isGlobal"`
	// Gets or sets the name of the connection type.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ConnectionType resource.
type ConnectionTypeArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The parameters supplied to the create or update connection type operation.
	ConnectionTypeName pulumi.StringInput
	// Gets or sets the field definitions of the connection type.
	FieldDefinitions FieldDefinitionMapInput
	// Gets or sets a Boolean value to indicate if the connection type is global.
	IsGlobal pulumi.BoolPtrInput
	// Gets or sets the name of the connection type.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
}

func (ConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionTypeArgs)(nil)).Elem()
}

type ConnectionTypeInput interface {
	pulumi.Input

	ToConnectionTypeOutput() ConnectionTypeOutput
	ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput
}

func (*ConnectionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionType)(nil))
}

func (i *ConnectionType) ToConnectionTypeOutput() ConnectionTypeOutput {
	return i.ToConnectionTypeOutputWithContext(context.Background())
}

func (i *ConnectionType) ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeOutput)
}

type ConnectionTypeOutput struct {
	*pulumi.OutputState
}

func (ConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionType)(nil))
}

func (o ConnectionTypeOutput) ToConnectionTypeOutput() ConnectionTypeOutput {
	return o
}

func (o ConnectionTypeOutput) ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionTypeOutput{})
}
