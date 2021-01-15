// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the connection.
// Latest API Version: 2019-06-01.
type Connection struct {
	pulumi.CustomResourceState

	// Gets or sets the connectionType of the connection.
	ConnectionType ConnectionTypeAssociationPropertyResponsePtrOutput `pulumi:"connectionType"`
	// Gets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets the field definition values of the connection.
	FieldDefinitionValues pulumi.StringMapOutput `pulumi:"fieldDefinitionValues"`
	// Gets the last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionName'")
	}
	if args.ConnectionType == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionType'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation/v20151031:Connection"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20190601:Connection"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20200113preview:Connection"),
		},
	})
	opts = append(opts, aliases)
	var resource Connection
	err := ctx.RegisterResource("azure-nextgen:automation/latest:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("azure-nextgen:automation/latest:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Gets or sets the connectionType of the connection.
	ConnectionType *ConnectionTypeAssociationPropertyResponse `pulumi:"connectionType"`
	// Gets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets the field definition values of the connection.
	FieldDefinitionValues map[string]string `pulumi:"fieldDefinitionValues"`
	// Gets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ConnectionState struct {
	// Gets or sets the connectionType of the connection.
	ConnectionType ConnectionTypeAssociationPropertyResponsePtrInput
	// Gets the creation time.
	CreationTime pulumi.StringPtrInput
	// Gets or sets the description.
	Description pulumi.StringPtrInput
	// Gets the field definition values of the connection.
	FieldDefinitionValues pulumi.StringMapInput
	// Gets the last modified time.
	LastModifiedTime pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The parameters supplied to the create or update connection operation.
	ConnectionName string `pulumi:"connectionName"`
	// Gets or sets the connectionType of the connection.
	ConnectionType ConnectionTypeAssociationProperty `pulumi:"connectionType"`
	// Gets or sets the description of the connection.
	Description *string `pulumi:"description"`
	// Gets or sets the field definition properties of the connection.
	FieldDefinitionValues map[string]string `pulumi:"fieldDefinitionValues"`
	// Gets or sets the name of the connection.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The parameters supplied to the create or update connection operation.
	ConnectionName pulumi.StringInput
	// Gets or sets the connectionType of the connection.
	ConnectionType ConnectionTypeAssociationPropertyInput
	// Gets or sets the description of the connection.
	Description pulumi.StringPtrInput
	// Gets or sets the field definition properties of the connection.
	FieldDefinitionValues pulumi.StringMapInput
	// Gets or sets the name of the connection.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct {
	*pulumi.OutputState
}

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
}
