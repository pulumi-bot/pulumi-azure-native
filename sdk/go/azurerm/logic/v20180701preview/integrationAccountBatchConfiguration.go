// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The batch configuration resource definition.
type IntegrationAccountBatchConfiguration struct {
	pulumi.CustomResourceState

	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The batch configuration properties.
	Properties BatchConfigurationPropertiesResponseOutput `pulumi:"properties"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccountBatchConfiguration registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountBatchConfiguration(ctx *pulumi.Context,
	name string, args *IntegrationAccountBatchConfigurationArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountBatchConfiguration, error) {
	if args == nil || args.BatchConfigurationName == nil {
		return nil, errors.New("missing required argument 'BatchConfigurationName'")
	}
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &IntegrationAccountBatchConfigurationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/latest:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20160601:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20190501:IntegrationAccountBatchConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountBatchConfiguration
	err := ctx.RegisterResource("azurerm:logic/v20180701preview:IntegrationAccountBatchConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountBatchConfiguration gets an existing IntegrationAccountBatchConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountBatchConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountBatchConfigurationState, opts ...pulumi.ResourceOption) (*IntegrationAccountBatchConfiguration, error) {
	var resource IntegrationAccountBatchConfiguration
	err := ctx.ReadResource("azurerm:logic/v20180701preview:IntegrationAccountBatchConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountBatchConfiguration resources.
type integrationAccountBatchConfigurationState struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The batch configuration properties.
	Properties *BatchConfigurationPropertiesResponse `pulumi:"properties"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type IntegrationAccountBatchConfigurationState struct {
	// The resource location.
	Location pulumi.StringPtrInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The batch configuration properties.
	Properties BatchConfigurationPropertiesResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountBatchConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountBatchConfigurationState)(nil)).Elem()
}

type integrationAccountBatchConfigurationArgs struct {
	// The batch configuration name.
	BatchConfigurationName string `pulumi:"batchConfigurationName"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The batch configuration properties.
	Properties BatchConfigurationProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccountBatchConfiguration resource.
type IntegrationAccountBatchConfigurationArgs struct {
	// The batch configuration name.
	BatchConfigurationName pulumi.StringInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The batch configuration properties.
	Properties BatchConfigurationPropertiesInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountBatchConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountBatchConfigurationArgs)(nil)).Elem()
}
