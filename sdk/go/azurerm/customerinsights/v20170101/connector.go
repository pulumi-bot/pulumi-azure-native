// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The connector resource format.
type Connector struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of connector.
	Properties ConnectorResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnector registers a new resource with the given unique name, arguments, and options.
func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil || args.HubName == nil {
		return nil, errors.New("missing required argument 'HubName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ConnectorArgs{}
	}
	var resource Connector
	err := ctx.RegisterResource("azurerm:customerinsights/v20170101:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnector gets an existing Connector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("azurerm:customerinsights/v20170101:Connector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connector resources.
type connectorState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of connector.
	Properties *ConnectorResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ConnectorState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of connector.
	Properties ConnectorResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorState)(nil)).Elem()
}

type connectorArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the connector.
	Name string `pulumi:"name"`
	// Properties of connector.
	Properties *ConnectorDefinition `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Connector resource.
type ConnectorArgs struct {
	// The name of the hub.
	HubName pulumi.StringInput
	// The name of the connector.
	Name pulumi.StringInput
	// Properties of connector.
	Properties ConnectorDefinitionPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorArgs)(nil)).Elem()
}