// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Static Site ARM resource.
// Latest API Version: 2020-09-01.
type StaticSite struct {
	pulumi.CustomResourceState

	// The target branch in the repository.
	Branch pulumi.StringPtrOutput `pulumi:"branch"`
	// Build properties to configure on the repository.
	BuildProperties StaticSiteBuildPropertiesResponsePtrOutput `pulumi:"buildProperties"`
	// The custom domains associated with this static site.
	CustomDomains pulumi.StringArrayOutput `pulumi:"customDomains"`
	// The default autogenerated hostname for the static site.
	DefaultHostname pulumi.StringOutput `pulumi:"defaultHostname"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken pulumi.StringPtrOutput `pulumi:"repositoryToken"`
	// URL for the repository of the static site.
	RepositoryUrl pulumi.StringPtrOutput `pulumi:"repositoryUrl"`
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrOutput `pulumi:"sku"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStaticSite registers a new resource with the given unique name, arguments, and options.
func NewStaticSite(ctx *pulumi.Context,
	name string, args *StaticSiteArgs, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:StaticSite"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticSite
	err := ctx.RegisterResource("azure-nextgen:web/latest:StaticSite", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:web/latest:StaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticSite resources.
type staticSiteState struct {
	// The target branch in the repository.
	Branch *string `pulumi:"branch"`
	// Build properties to configure on the repository.
	BuildProperties *StaticSiteBuildPropertiesResponse `pulumi:"buildProperties"`
	// The custom domains associated with this static site.
	CustomDomains []string `pulumi:"customDomains"`
	// The default autogenerated hostname for the static site.
	DefaultHostname *string `pulumi:"defaultHostname"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken *string `pulumi:"repositoryToken"`
	// URL for the repository of the static site.
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescriptionResponse `pulumi:"sku"`
	// The system metadata relating to this resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type StaticSiteState struct {
	// The target branch in the repository.
	Branch pulumi.StringPtrInput
	// Build properties to configure on the repository.
	BuildProperties StaticSiteBuildPropertiesResponsePtrInput
	// The custom domains associated with this static site.
	CustomDomains pulumi.StringArrayInput
	// The default autogenerated hostname for the static site.
	DefaultHostname pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken pulumi.StringPtrInput
	// URL for the repository of the static site.
	RepositoryUrl pulumi.StringPtrInput
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrInput
	// The system metadata relating to this resource.
	SystemData SystemDataResponsePtrInput
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
	Location *string `pulumi:"location"`
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
	Location pulumi.StringPtrInput
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

type StaticSiteInput interface {
	pulumi.Input

	ToStaticSiteOutput() StaticSiteOutput
	ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput
}

func (*StaticSite) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSite)(nil))
}

func (i *StaticSite) ToStaticSiteOutput() StaticSiteOutput {
	return i.ToStaticSiteOutputWithContext(context.Background())
}

func (i *StaticSite) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteOutput)
}

type StaticSiteOutput struct {
	*pulumi.OutputState
}

func (StaticSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSite)(nil))
}

func (o StaticSiteOutput) ToStaticSiteOutput() StaticSiteOutput {
	return o
}

func (o StaticSiteOutput) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StaticSiteOutput{})
}
