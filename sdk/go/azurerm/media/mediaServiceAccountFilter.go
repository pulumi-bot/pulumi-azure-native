// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Account Filter.
type MediaServiceAccountFilter struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Media Filter properties.
	Properties MediaFilterPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMediaServiceAccountFilter registers a new resource with the given unique name, arguments, and options.
func NewMediaServiceAccountFilter(ctx *pulumi.Context,
	name string, args *MediaServiceAccountFilterArgs, opts ...pulumi.ResourceOption) (*MediaServiceAccountFilter, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MediaServiceAccountFilterArgs{}
	}
	var resource MediaServiceAccountFilter
	err := ctx.RegisterResource("azurerm:media:MediaServiceAccountFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaServiceAccountFilter gets an existing MediaServiceAccountFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaServiceAccountFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaServiceAccountFilterState, opts ...pulumi.ResourceOption) (*MediaServiceAccountFilter, error) {
	var resource MediaServiceAccountFilter
	err := ctx.ReadResource("azurerm:media:MediaServiceAccountFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaServiceAccountFilter resources.
type mediaServiceAccountFilterState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// The Media Filter properties.
	Properties *MediaFilterPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type MediaServiceAccountFilterState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// The Media Filter properties.
	Properties MediaFilterPropertiesResponsePtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (MediaServiceAccountFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceAccountFilterState)(nil)).Elem()
}

type mediaServiceAccountFilterArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Account Filter name
	Name string `pulumi:"name"`
	// The Media Filter properties.
	Properties *MediaFilterProperties `pulumi:"properties"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a MediaServiceAccountFilter resource.
type MediaServiceAccountFilterArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Account Filter name
	Name pulumi.StringInput
	// The Media Filter properties.
	Properties MediaFilterPropertiesPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (MediaServiceAccountFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceAccountFilterArgs)(nil)).Elem()
}