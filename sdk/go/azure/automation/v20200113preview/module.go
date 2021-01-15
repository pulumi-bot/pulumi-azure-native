// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200113preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the module type.
type Module struct {
	pulumi.CustomResourceState

	// Gets or sets the activity count of the module.
	ActivityCount pulumi.IntPtrOutput `pulumi:"activityCount"`
	// Gets or sets the contentLink of the module.
	ContentLink ContentLinkResponsePtrOutput `pulumi:"contentLink"`
	// Gets or sets the creation time.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets or sets the error info of the module.
	Error ModuleErrorInfoResponsePtrOutput `pulumi:"error"`
	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets or sets type of module, if its composite or not.
	IsComposite pulumi.BoolPtrOutput `pulumi:"isComposite"`
	// Gets or sets the isGlobal flag of the module.
	IsGlobal pulumi.BoolPtrOutput `pulumi:"isGlobal"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the provisioning state of the module.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets the size in bytes of the module.
	SizeInBytes pulumi.Float64PtrOutput `pulumi:"sizeInBytes"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets the version of the module.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewModule registers a new resource with the given unique name, arguments, and options.
func NewModule(ctx *pulumi.Context,
	name string, args *ModuleArgs, opts ...pulumi.ResourceOption) (*Module, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ContentLink == nil {
		return nil, errors.New("invalid value for required argument 'ContentLink'")
	}
	if args.ModuleName == nil {
		return nil, errors.New("invalid value for required argument 'ModuleName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation/latest:Module"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20151031:Module"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20190601:Module"),
		},
	})
	opts = append(opts, aliases)
	var resource Module
	err := ctx.RegisterResource("azure-nextgen:automation/v20200113preview:Module", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModule gets an existing Module resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModuleState, opts ...pulumi.ResourceOption) (*Module, error) {
	var resource Module
	err := ctx.ReadResource("azure-nextgen:automation/v20200113preview:Module", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Module resources.
type moduleState struct {
	// Gets or sets the activity count of the module.
	ActivityCount *int `pulumi:"activityCount"`
	// Gets or sets the contentLink of the module.
	ContentLink *ContentLinkResponse `pulumi:"contentLink"`
	// Gets or sets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets or sets the error info of the module.
	Error *ModuleErrorInfoResponse `pulumi:"error"`
	// Gets or sets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// Gets or sets type of module, if its composite or not.
	IsComposite *bool `pulumi:"isComposite"`
	// Gets or sets the isGlobal flag of the module.
	IsGlobal *bool `pulumi:"isGlobal"`
	// Gets or sets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets the provisioning state of the module.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets the size in bytes of the module.
	SizeInBytes *float64 `pulumi:"sizeInBytes"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// Gets or sets the version of the module.
	Version *string `pulumi:"version"`
}

type ModuleState struct {
	// Gets or sets the activity count of the module.
	ActivityCount pulumi.IntPtrInput
	// Gets or sets the contentLink of the module.
	ContentLink ContentLinkResponsePtrInput
	// Gets or sets the creation time.
	CreationTime pulumi.StringPtrInput
	// Gets or sets the description.
	Description pulumi.StringPtrInput
	// Gets or sets the error info of the module.
	Error ModuleErrorInfoResponsePtrInput
	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrInput
	// Gets or sets type of module, if its composite or not.
	IsComposite pulumi.BoolPtrInput
	// Gets or sets the isGlobal flag of the module.
	IsGlobal pulumi.BoolPtrInput
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets the provisioning state of the module.
	ProvisioningState pulumi.StringPtrInput
	// Gets or sets the size in bytes of the module.
	SizeInBytes pulumi.Float64PtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// Gets or sets the version of the module.
	Version pulumi.StringPtrInput
}

func (ModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleState)(nil)).Elem()
}

type moduleArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the module content link.
	ContentLink ContentLink `pulumi:"contentLink"`
	// Gets or sets the location of the resource.
	Location *string `pulumi:"location"`
	// The name of module.
	ModuleName string `pulumi:"moduleName"`
	// Gets or sets name of the resource.
	Name *string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Module resource.
type ModuleArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the module content link.
	ContentLink ContentLinkInput
	// Gets or sets the location of the resource.
	Location pulumi.StringPtrInput
	// The name of module.
	ModuleName pulumi.StringInput
	// Gets or sets name of the resource.
	Name pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (ModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleArgs)(nil)).Elem()
}

type ModuleInput interface {
	pulumi.Input

	ToModuleOutput() ModuleOutput
	ToModuleOutputWithContext(ctx context.Context) ModuleOutput
}

func (*Module) ElementType() reflect.Type {
	return reflect.TypeOf((*Module)(nil))
}

func (i *Module) ToModuleOutput() ModuleOutput {
	return i.ToModuleOutputWithContext(context.Background())
}

func (i *Module) ToModuleOutputWithContext(ctx context.Context) ModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleOutput)
}

type ModuleOutput struct {
	*pulumi.OutputState
}

func (ModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Module)(nil))
}

func (o ModuleOutput) ToModuleOutput() ModuleOutput {
	return o
}

func (o ModuleOutput) ToModuleOutputWithContext(ctx context.Context) ModuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ModuleOutput{})
}
