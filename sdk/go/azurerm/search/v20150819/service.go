// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150819

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes an Azure Cognitive Search service and its current state.
type Service struct {
	pulumi.CustomResourceState

	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Search service.
	Properties SearchServicePropertiesResponseOutput `pulumi:"properties"`
	// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Tags to help categorize the resource in the Azure portal.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("azurerm:search/v20150819:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azurerm:search/v20150819:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Properties of the Search service.
	Properties *SearchServicePropertiesResponse `pulumi:"properties"`
	// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
	Sku *SkuResponse `pulumi:"sku"`
	// Tags to help categorize the resource in the Azure portal.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type ServiceState struct {
	// The identity of the resource.
	Identity IdentityResponsePtrInput
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Properties of the Search service.
	Properties SearchServicePropertiesResponsePtrInput
	// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
	Sku SkuResponsePtrInput
	// Tags to help categorize the resource in the Azure portal.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
	HostingMode *string `pulumi:"hostingMode"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
	Location *string `pulumi:"location"`
	// The name of the Azure Cognitive Search service to create or update. Search service names must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot contain consecutive dashes, and must be between 2 and 60 characters in length. Search service names must be globally unique since they are part of the service URI (https://<name>.search.windows.net). You cannot change the service name after the service is created.
	Name string `pulumi:"name"`
	// The number of partitions in the Search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
	PartitionCount *int `pulumi:"partitionCount"`
	// The number of replicas in the Search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
	ReplicaCount *int `pulumi:"replicaCount"`
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
	Sku *Sku `pulumi:"sku"`
	// Tags to help categorize the resource in the Azure portal.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
	HostingMode pulumi.StringPtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
	Location pulumi.StringPtrInput
	// The name of the Azure Cognitive Search service to create or update. Search service names must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot contain consecutive dashes, and must be between 2 and 60 characters in length. Search service names must be globally unique since they are part of the service URI (https://<name>.search.windows.net). You cannot change the service name after the service is created.
	Name pulumi.StringInput
	// The number of partitions in the Search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
	PartitionCount pulumi.IntPtrInput
	// The number of replicas in the Search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
	ReplicaCount pulumi.IntPtrInput
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
	Sku SkuPtrInput
	// Tags to help categorize the resource in the Azure portal.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
