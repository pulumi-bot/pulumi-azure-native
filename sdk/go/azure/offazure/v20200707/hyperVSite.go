// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200707

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Site REST Resource.
type HyperVSite struct {
	pulumi.CustomResourceState

	// eTag for concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the Hyper-V site.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Nested properties of Hyper-V site.
	Properties SitePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput       `pulumi:"tags"`
	// Type of resource. Type = Microsoft.OffAzure/HyperVSites.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHyperVSite registers a new resource with the given unique name, arguments, and options.
func NewHyperVSite(ctx *pulumi.Context,
	name string, args *HyperVSiteArgs, opts ...pulumi.ResourceOption) (*HyperVSite, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SiteName == nil {
		return nil, errors.New("missing required argument 'SiteName'")
	}
	if args == nil {
		args = &HyperVSiteArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:offazure/latest:HyperVSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:offazure/v20200101:HyperVSite"),
		},
	})
	opts = append(opts, aliases)
	var resource HyperVSite
	err := ctx.RegisterResource("azure-nextgen:offazure/v20200707:HyperVSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHyperVSite gets an existing HyperVSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHyperVSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HyperVSiteState, opts ...pulumi.ResourceOption) (*HyperVSite, error) {
	var resource HyperVSite
	err := ctx.ReadResource("azure-nextgen:offazure/v20200707:HyperVSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HyperVSite resources.
type hyperVSiteState struct {
	// eTag for concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location *string `pulumi:"location"`
	// Name of the Hyper-V site.
	Name *string `pulumi:"name"`
	// Nested properties of Hyper-V site.
	Properties *SitePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string       `pulumi:"tags"`
	// Type of resource. Type = Microsoft.OffAzure/HyperVSites.
	Type *string `pulumi:"type"`
}

type HyperVSiteState struct {
	// eTag for concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which Sites is created.
	Location pulumi.StringPtrInput
	// Name of the Hyper-V site.
	Name pulumi.StringPtrInput
	// Nested properties of Hyper-V site.
	Properties SitePropertiesResponsePtrInput
	Tags       pulumi.StringMapInput
	// Type of resource. Type = Microsoft.OffAzure/HyperVSites.
	Type pulumi.StringPtrInput
}

func (HyperVSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVSiteState)(nil)).Elem()
}

type hyperVSiteArgs struct {
	// eTag for concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location *string `pulumi:"location"`
	// Name of the Hyper-V site.
	Name *string `pulumi:"name"`
	// Nested properties of Hyper-V site.
	Properties *SiteProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name.
	SiteName string            `pulumi:"siteName"`
	Tags     map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a HyperVSite resource.
type HyperVSiteArgs struct {
	// eTag for concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which Sites is created.
	Location pulumi.StringPtrInput
	// Name of the Hyper-V site.
	Name pulumi.StringPtrInput
	// Nested properties of Hyper-V site.
	Properties SitePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Site name.
	SiteName pulumi.StringInput
	Tags     pulumi.StringMapInput
}

func (HyperVSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVSiteArgs)(nil)).Elem()
}

type HyperVSiteInput interface {
	pulumi.Input

	ToHyperVSiteOutput() HyperVSiteOutput
	ToHyperVSiteOutputWithContext(ctx context.Context) HyperVSiteOutput
}

func (HyperVSite) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVSite)(nil)).Elem()
}

func (i HyperVSite) ToHyperVSiteOutput() HyperVSiteOutput {
	return i.ToHyperVSiteOutputWithContext(context.Background())
}

func (i HyperVSite) ToHyperVSiteOutputWithContext(ctx context.Context) HyperVSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVSiteOutput)
}

type HyperVSiteOutput struct {
	*pulumi.OutputState
}

func (HyperVSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVSiteOutput)(nil)).Elem()
}

func (o HyperVSiteOutput) ToHyperVSiteOutput() HyperVSiteOutput {
	return o
}

func (o HyperVSiteOutput) ToHyperVSiteOutputWithContext(ctx context.Context) HyperVSiteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HyperVSiteOutput{})
}
