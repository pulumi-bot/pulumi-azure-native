// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes the source control configuration for web app
type SiteSourceControlSlot struct {
	pulumi.CustomResourceState

	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name       pulumi.StringPtrOutput                    `pulumi:"name"`
	Properties SiteSourceControlResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSiteSourceControlSlot registers a new resource with the given unique name, arguments, and options.
func NewSiteSourceControlSlot(ctx *pulumi.Context,
	name string, args *SiteSourceControlSlotArgs, opts ...pulumi.ResourceOption) (*SiteSourceControlSlot, error) {
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
		args = &SiteSourceControlSlotArgs{}
	}
	var resource SiteSourceControlSlot
	err := ctx.RegisterResource("azurerm:web/v20150801:SiteSourceControlSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteSourceControlSlot gets an existing SiteSourceControlSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteSourceControlSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteSourceControlSlotState, opts ...pulumi.ResourceOption) (*SiteSourceControlSlot, error) {
	var resource SiteSourceControlSlot
	err := ctx.ReadResource("azurerm:web/v20150801:SiteSourceControlSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteSourceControlSlot resources.
type siteSourceControlSlotState struct {
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name       *string                              `pulumi:"name"`
	Properties *SiteSourceControlResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SiteSourceControlSlotState struct {
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name       pulumi.StringPtrInput
	Properties SiteSourceControlResponsePropertiesPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteSourceControlSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSourceControlSlotState)(nil)).Elem()
}

type siteSourceControlSlotArgs struct {
	// Name of branch to use for deployment
	Branch *string `pulumi:"branch"`
	// Whether to manual or continuous integration
	DeploymentRollbackEnabled *bool `pulumi:"deploymentRollbackEnabled"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Whether to manual or continuous integration
	IsManualIntegration *bool `pulumi:"isManualIntegration"`
	// Mercurial or Git repository type
	IsMercurial *bool `pulumi:"isMercurial"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Name of web app slot. If not specified then will default to production slot.
	Name string `pulumi:"name"`
	// Repository or source control url
	RepoUrl *string `pulumi:"repoUrl"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SiteSourceControlSlot resource.
type SiteSourceControlSlotArgs struct {
	// Name of branch to use for deployment
	Branch pulumi.StringPtrInput
	// Whether to manual or continuous integration
	DeploymentRollbackEnabled pulumi.BoolPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Whether to manual or continuous integration
	IsManualIntegration pulumi.BoolPtrInput
	// Mercurial or Git repository type
	IsMercurial pulumi.BoolPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringInput
	// Name of web app slot. If not specified then will default to production slot.
	Name pulumi.StringInput
	// Repository or source control url
	RepoUrl pulumi.StringPtrInput
	// Name of resource group
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteSourceControlSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSourceControlSlotArgs)(nil)).Elem()
}
