// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The application type name resource
type ApplicationType struct {
	pulumi.CustomResourceState

	// Azure resource etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationType registers a new resource with the given unique name, arguments, and options.
func NewApplicationType(ctx *pulumi.Context,
	name string, args *ApplicationTypeArgs, opts ...pulumi.ResourceOption) (*ApplicationType, error) {
	if args == nil || args.ApplicationTypeName == nil {
		return nil, errors.New("missing required argument 'ApplicationTypeName'")
	}
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationTypeArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:servicefabric/latest:ApplicationType"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20170701preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301:ApplicationType"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190601preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20200301:ApplicationType"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationType
	err := ctx.RegisterResource("azurerm:servicefabric/v20191101preview:ApplicationType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationType gets an existing ApplicationType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationTypeState, opts ...pulumi.ResourceOption) (*ApplicationType, error) {
	var resource ApplicationType
	err := ctx.ReadResource("azurerm:servicefabric/v20191101preview:ApplicationType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationType resources.
type applicationTypeState struct {
	// Azure resource etag.
	Etag *string `pulumi:"etag"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// Azure resource name.
	Name *string `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type *string `pulumi:"type"`
}

type ApplicationTypeState struct {
	// Azure resource etag.
	Etag pulumi.StringPtrInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// Azure resource name.
	Name pulumi.StringPtrInput
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// Azure resource type.
	Type pulumi.StringPtrInput
}

func (ApplicationTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeState)(nil)).Elem()
}

type applicationTypeArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationType resource.
type ApplicationTypeArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName pulumi.StringInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
}

func (ApplicationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeArgs)(nil)).Elem()
}
