// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of a namespace resource.
type Namespace struct {
	pulumi.CustomResourceState

	// The time the namespace was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Identifier for Azure Insights metrics.
	MetricId pulumi.StringOutput `pulumi:"metricId"`
	// Resource name.
	Name              pulumi.StringOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringOutput `pulumi:"serviceBusEndpoint"`
	// SKU of the namespace.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NamespaceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:relay/v20160701:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:relay/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:relay/v20180101preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-nextgen:relay/latest:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-nextgen:relay/latest:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// The time the namespace was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Identifier for Azure Insights metrics.
	MetricId *string `pulumi:"metricId"`
	// Resource name.
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `pulumi:"serviceBusEndpoint"`
	// SKU of the namespace.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type NamespaceState struct {
	// The time the namespace was created.
	CreatedAt pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Identifier for Azure Insights metrics.
	MetricId pulumi.StringPtrInput
	// Resource name.
	Name              pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringPtrInput
	// SKU of the namespace.
	Sku SkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The time the namespace was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Resource location.
	Location string `pulumi:"location"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU of the namespace.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Resource location.
	Location pulumi.StringInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// SKU of the namespace.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil)).Elem()
}

func (i Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct {
	*pulumi.OutputState
}

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceOutput)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
