// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200707

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Site REST Resource.
//
// ## Example Usage
// ### Create VMware site
//
// ```go
// package main
//
// import (
// 	offazure "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/offazure/v20200707"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := offazure.NewSite(ctx, "site", &offazure.SiteArgs{
// 			Location:          pulumi.String("eastus"),
// 			ResourceGroupName: pulumi.String("pajindTest"),
// 			SiteName:          pulumi.String("appliance1e39site"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Site struct {
	pulumi.CustomResourceState

	// eTag for concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the VMware site.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Nested properties of VMWare site.
	Properties SitePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput       `pulumi:"tags"`
	// Type of resource. Type = Microsoft.OffAzure/VMWareSites.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSite registers a new resource with the given unique name, arguments, and options.
func NewSite(ctx *pulumi.Context,
	name string, args *SiteArgs, opts ...pulumi.ResourceOption) (*Site, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SiteName == nil {
		return nil, errors.New("missing required argument 'SiteName'")
	}
	if args == nil {
		args = &SiteArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:offazure/latest:Site"),
		},
		{
			Type: pulumi.String("azurerm:offazure/v20200101:Site"),
		},
	})
	opts = append(opts, aliases)
	var resource Site
	err := ctx.RegisterResource("azurerm:offazure/v20200707:Site", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSite gets an existing Site resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteState, opts ...pulumi.ResourceOption) (*Site, error) {
	var resource Site
	err := ctx.ReadResource("azurerm:offazure/v20200707:Site", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Site resources.
type siteState struct {
	// eTag for concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location *string `pulumi:"location"`
	// Name of the VMware site.
	Name *string `pulumi:"name"`
	// Nested properties of VMWare site.
	Properties *SitePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string       `pulumi:"tags"`
	// Type of resource. Type = Microsoft.OffAzure/VMWareSites.
	Type *string `pulumi:"type"`
}

type SiteState struct {
	// eTag for concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which Sites is created.
	Location pulumi.StringPtrInput
	// Name of the VMware site.
	Name pulumi.StringPtrInput
	// Nested properties of VMWare site.
	Properties SitePropertiesResponsePtrInput
	Tags       pulumi.StringMapInput
	// Type of resource. Type = Microsoft.OffAzure/VMWareSites.
	Type pulumi.StringPtrInput
}

func (SiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteState)(nil)).Elem()
}

type siteArgs struct {
	// eTag for concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location *string `pulumi:"location"`
	// Name of the VMware site.
	Name *string `pulumi:"name"`
	// Nested properties of VMWare site.
	Properties *SiteProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name.
	SiteName string            `pulumi:"siteName"`
	Tags     map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Site resource.
type SiteArgs struct {
	// eTag for concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which Sites is created.
	Location pulumi.StringPtrInput
	// Name of the VMware site.
	Name pulumi.StringPtrInput
	// Nested properties of VMWare site.
	Properties SitePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Site name.
	SiteName pulumi.StringInput
	Tags     pulumi.StringMapInput
}

func (SiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteArgs)(nil)).Elem()
}
