// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The service resource.
//
// ## Example Usage
// ### Put a service with maximum parameters
//
// ```go
//
// ```
// ### Put a service with minimum parameters
//
// ```go
//
// ```
type Service struct {
	pulumi.CustomResourceState

	// A list that describes the correlation of the service with other services.
	CorrelationScheme ServiceCorrelationDescriptionResponseArrayOutput `pulumi:"correlationScheme"`
	// Specifies the move cost for the service.
	DefaultMoveCost pulumi.StringPtrOutput `pulumi:"defaultMoveCost"`
	// Azure resource etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes how the service is partitioned.
	PartitionDescription pulumi.AnyOutput `pulumi:"partitionDescription"`
	// The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
	PlacementConstraints pulumi.StringPtrOutput `pulumi:"placementConstraints"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The kind of service (Stateless or Stateful).
	ServiceKind pulumi.StringOutput `pulumi:"serviceKind"`
	// The service load metrics is given as an array of ServiceLoadMetricDescription objects.
	ServiceLoadMetrics ServiceLoadMetricDescriptionResponseArrayOutput `pulumi:"serviceLoadMetrics"`
	// The activation Mode of the service package
	ServicePackageActivationMode pulumi.StringPtrOutput `pulumi:"servicePackageActivationMode"`
	// A list that describes the correlation of the service with other services.
	ServicePlacementPolicies ServicePlacementPolicyDescriptionResponseArrayOutput `pulumi:"servicePlacementPolicies"`
	// The name of the service type
	ServiceTypeName pulumi.StringPtrOutput `pulumi:"serviceTypeName"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.ApplicationName == nil {
		return nil, errors.New("missing required argument 'ApplicationName'")
	}
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceKind == nil {
		return nil, errors.New("missing required argument 'ServiceKind'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:servicefabric/latest:Service"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20170701preview:Service"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301:Service"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301preview:Service"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20191101preview:Service"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20200301:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azurerm:servicefabric/v20190601preview:Service", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:servicefabric/v20190601preview:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// A list that describes the correlation of the service with other services.
	CorrelationScheme []ServiceCorrelationDescriptionResponse `pulumi:"correlationScheme"`
	// Specifies the move cost for the service.
	DefaultMoveCost *string `pulumi:"defaultMoveCost"`
	// Azure resource etag.
	Etag *string `pulumi:"etag"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// Azure resource name.
	Name *string `pulumi:"name"`
	// Describes how the service is partitioned.
	PartitionDescription interface{} `pulumi:"partitionDescription"`
	// The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
	PlacementConstraints *string `pulumi:"placementConstraints"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState *string `pulumi:"provisioningState"`
	// The kind of service (Stateless or Stateful).
	ServiceKind *string `pulumi:"serviceKind"`
	// The service load metrics is given as an array of ServiceLoadMetricDescription objects.
	ServiceLoadMetrics []ServiceLoadMetricDescriptionResponse `pulumi:"serviceLoadMetrics"`
	// The activation Mode of the service package
	ServicePackageActivationMode *string `pulumi:"servicePackageActivationMode"`
	// A list that describes the correlation of the service with other services.
	ServicePlacementPolicies []ServicePlacementPolicyDescriptionResponse `pulumi:"servicePlacementPolicies"`
	// The name of the service type
	ServiceTypeName *string `pulumi:"serviceTypeName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type *string `pulumi:"type"`
}

type ServiceState struct {
	// A list that describes the correlation of the service with other services.
	CorrelationScheme ServiceCorrelationDescriptionResponseArrayInput
	// Specifies the move cost for the service.
	DefaultMoveCost pulumi.StringPtrInput
	// Azure resource etag.
	Etag pulumi.StringPtrInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// Azure resource name.
	Name pulumi.StringPtrInput
	// Describes how the service is partitioned.
	PartitionDescription pulumi.Input
	// The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
	PlacementConstraints pulumi.StringPtrInput
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringPtrInput
	// The kind of service (Stateless or Stateful).
	ServiceKind pulumi.StringPtrInput
	// The service load metrics is given as an array of ServiceLoadMetricDescription objects.
	ServiceLoadMetrics ServiceLoadMetricDescriptionResponseArrayInput
	// The activation Mode of the service package
	ServicePackageActivationMode pulumi.StringPtrInput
	// A list that describes the correlation of the service with other services.
	ServicePlacementPolicies ServicePlacementPolicyDescriptionResponseArrayInput
	// The name of the service type
	ServiceTypeName pulumi.StringPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// Azure resource type.
	Type pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The name of the application resource.
	ApplicationName string `pulumi:"applicationName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// A list that describes the correlation of the service with other services.
	CorrelationScheme []ServiceCorrelationDescription `pulumi:"correlationScheme"`
	// Specifies the move cost for the service.
	DefaultMoveCost *string `pulumi:"defaultMoveCost"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// Describes how the service is partitioned.
	PartitionDescription interface{} `pulumi:"partitionDescription"`
	// The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
	PlacementConstraints *string `pulumi:"placementConstraints"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The kind of service (Stateless or Stateful).
	ServiceKind string `pulumi:"serviceKind"`
	// The service load metrics is given as an array of ServiceLoadMetricDescription objects.
	ServiceLoadMetrics []ServiceLoadMetricDescription `pulumi:"serviceLoadMetrics"`
	// The name of the service resource in the format of {applicationName}~{serviceName}.
	ServiceName string `pulumi:"serviceName"`
	// The activation Mode of the service package
	ServicePackageActivationMode *string `pulumi:"servicePackageActivationMode"`
	// A list that describes the correlation of the service with other services.
	ServicePlacementPolicies []ServicePlacementPolicyDescription `pulumi:"servicePlacementPolicies"`
	// The name of the service type
	ServiceTypeName *string `pulumi:"serviceTypeName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The name of the application resource.
	ApplicationName pulumi.StringInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// A list that describes the correlation of the service with other services.
	CorrelationScheme ServiceCorrelationDescriptionArrayInput
	// Specifies the move cost for the service.
	DefaultMoveCost pulumi.StringPtrInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// Describes how the service is partitioned.
	PartitionDescription pulumi.Input
	// The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
	PlacementConstraints pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The kind of service (Stateless or Stateful).
	ServiceKind pulumi.StringInput
	// The service load metrics is given as an array of ServiceLoadMetricDescription objects.
	ServiceLoadMetrics ServiceLoadMetricDescriptionArrayInput
	// The name of the service resource in the format of {applicationName}~{serviceName}.
	ServiceName pulumi.StringInput
	// The activation Mode of the service package
	ServicePackageActivationMode pulumi.StringPtrInput
	// A list that describes the correlation of the service with other services.
	ServicePlacementPolicies ServicePlacementPolicyDescriptionArrayInput
	// The name of the service type
	ServiceTypeName pulumi.StringPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
