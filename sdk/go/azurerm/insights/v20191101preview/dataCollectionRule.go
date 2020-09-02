// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of ARM tracked top level resource.
type DataCollectionRule struct {
	pulumi.CustomResourceState

	// The specification of data flows.
	DataFlows DataFlowResponseArrayOutput `pulumi:"dataFlows"`
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources DataCollectionRuleResponseDataSourcesPtrOutput `pulumi:"dataSources"`
	// Description of the data collection rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The specification of destinations.
	Destinations DataCollectionRuleResponseDestinationsOutput `pulumi:"destinations"`
	// Resource entity tag (ETag).
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataCollectionRule registers a new resource with the given unique name, arguments, and options.
func NewDataCollectionRule(ctx *pulumi.Context,
	name string, args *DataCollectionRuleArgs, opts ...pulumi.ResourceOption) (*DataCollectionRule, error) {
	if args == nil || args.DataCollectionRuleName == nil {
		return nil, errors.New("missing required argument 'DataCollectionRuleName'")
	}
	if args == nil || args.DataFlows == nil {
		return nil, errors.New("missing required argument 'DataFlows'")
	}
	if args == nil || args.Destinations == nil {
		return nil, errors.New("missing required argument 'Destinations'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DataCollectionRuleArgs{}
	}
	var resource DataCollectionRule
	err := ctx.RegisterResource("azurerm:insights/v20191101preview:DataCollectionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCollectionRule gets an existing DataCollectionRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCollectionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCollectionRuleState, opts ...pulumi.ResourceOption) (*DataCollectionRule, error) {
	var resource DataCollectionRule
	err := ctx.ReadResource("azurerm:insights/v20191101preview:DataCollectionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCollectionRule resources.
type dataCollectionRuleState struct {
	// The specification of data flows.
	DataFlows []DataFlowResponse `pulumi:"dataFlows"`
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources *DataCollectionRuleResponseDataSources `pulumi:"dataSources"`
	// Description of the data collection rule.
	Description *string `pulumi:"description"`
	// The specification of destinations.
	Destinations *DataCollectionRuleResponseDestinations `pulumi:"destinations"`
	// Resource entity tag (ETag).
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The resource provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type DataCollectionRuleState struct {
	// The specification of data flows.
	DataFlows DataFlowResponseArrayInput
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources DataCollectionRuleResponseDataSourcesPtrInput
	// Description of the data collection rule.
	Description pulumi.StringPtrInput
	// The specification of destinations.
	Destinations DataCollectionRuleResponseDestinationsPtrInput
	// Resource entity tag (ETag).
	Etag pulumi.StringPtrInput
	// The geo-location where the resource lives.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The resource provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (DataCollectionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleState)(nil)).Elem()
}

type dataCollectionRuleArgs struct {
	// The name of the data collection rule. The name is case insensitive.
	DataCollectionRuleName string `pulumi:"dataCollectionRuleName"`
	// The specification of data flows.
	DataFlows []DataFlow `pulumi:"dataFlows"`
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources *DataCollectionRuleDataSources `pulumi:"dataSources"`
	// Description of the data collection rule.
	Description *string `pulumi:"description"`
	// The specification of destinations.
	Destinations DataCollectionRuleDestinations `pulumi:"destinations"`
	// The geo-location where the resource lives.
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataCollectionRule resource.
type DataCollectionRuleArgs struct {
	// The name of the data collection rule. The name is case insensitive.
	DataCollectionRuleName pulumi.StringInput
	// The specification of data flows.
	DataFlows DataFlowArrayInput
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources DataCollectionRuleDataSourcesPtrInput
	// Description of the data collection rule.
	Description pulumi.StringPtrInput
	// The specification of destinations.
	Destinations DataCollectionRuleDestinationsInput
	// The geo-location where the resource lives.
	Location pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DataCollectionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleArgs)(nil)).Elem()
}
