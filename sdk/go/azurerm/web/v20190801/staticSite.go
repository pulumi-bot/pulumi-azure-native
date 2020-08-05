// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Static Site ARM resource.
type StaticSite struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Core resource properties
	Properties StaticSiteResponseOutput `pulumi:"properties"`
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStaticSite registers a new resource with the given unique name, arguments, and options.
func NewStaticSite(ctx *pulumi.Context,
	name string, args *StaticSiteArgs, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StaticSiteArgs{}
	}
	var resource StaticSite
	err := ctx.RegisterResource("azurerm:web/v20190801:StaticSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticSite gets an existing StaticSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteState, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	var resource StaticSite
	err := ctx.ReadResource("azurerm:web/v20190801:StaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticSite resources.
type staticSiteState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Core resource properties
	Properties *StaticSiteResponse `pulumi:"properties"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescriptionResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type StaticSiteState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Core resource properties
	Properties StaticSiteResponsePtrInput
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (StaticSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteState)(nil)).Elem()
}

type staticSiteArgs struct {
	// The target branch in the repository.
	Branch *string `pulumi:"branch"`
	// Build properties to configure on the repository.
	BuildProperties *StaticSiteBuildProperties `pulumi:"buildProperties"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Name of the static site to create or update.
	Name string `pulumi:"name"`
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken *string `pulumi:"repositoryToken"`
	// URL for the repository of the static site.
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescription `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StaticSite resource.
type StaticSiteArgs struct {
	// The target branch in the repository.
	Branch pulumi.StringPtrInput
	// Build properties to configure on the repository.
	BuildProperties StaticSiteBuildPropertiesPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringInput
	// Name of the static site to create or update.
	Name pulumi.StringInput
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken pulumi.StringPtrInput
	// URL for the repository of the static site.
	RepositoryUrl pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (StaticSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteArgs)(nil)).Elem()
}
