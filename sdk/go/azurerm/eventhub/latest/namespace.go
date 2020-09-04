// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Single Namespace item in List or Get Operation
type Namespace struct {
	pulumi.CustomResourceState

	// The time the Namespace was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled pulumi.BoolPtrOutput `pulumi:"isAutoInflateEnabled"`
	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled pulumi.BoolPtrOutput `pulumi:"kafkaEnabled"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits pulumi.IntPtrOutput `pulumi:"maximumThroughputUnits"`
	// Identifier for Azure Insights metrics.
	MetricId pulumi.StringOutput `pulumi:"metricId"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Namespace.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringOutput `pulumi:"serviceBusEndpoint"`
	// Properties of sku resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the Namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
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
			Type: pulumi.String("azurerm:eventhub/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azurerm:eventhub/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azurerm:eventhub/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azurerm:eventhub/v20180101preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azurerm:eventhub/latest:Namespace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:eventhub/latest:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// The time the Namespace was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled *bool `pulumi:"isAutoInflateEnabled"`
	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled *bool `pulumi:"kafkaEnabled"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits *int `pulumi:"maximumThroughputUnits"`
	// Identifier for Azure Insights metrics.
	MetricId *string `pulumi:"metricId"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Provisioning state of the Namespace.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `pulumi:"serviceBusEndpoint"`
	// Properties of sku resource
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The time the Namespace was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type NamespaceState struct {
	// The time the Namespace was created.
	CreatedAt pulumi.StringPtrInput
	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled pulumi.BoolPtrInput
	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits pulumi.IntPtrInput
	// Identifier for Azure Insights metrics.
	MetricId pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Provisioning state of the Namespace.
	ProvisioningState pulumi.StringPtrInput
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringPtrInput
	// Properties of sku resource
	Sku SkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The time the Namespace was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled *bool `pulumi:"isAutoInflateEnabled"`
	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled *bool `pulumi:"kafkaEnabled"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits *int `pulumi:"maximumThroughputUnits"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Properties of sku resource
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled pulumi.BoolPtrInput
	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits pulumi.IntPtrInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// Properties of sku resource
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}
