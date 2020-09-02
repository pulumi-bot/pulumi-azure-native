// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Live Event.
type LiveEvent struct {
	pulumi.CustomResourceState

	// The exact time the Live Event was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// The Live Event access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput `pulumi:"crossSiteAccessPolicies"`
	// The Live Event description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Live Event encoding.
	Encoding LiveEventEncodingResponsePtrOutput `pulumi:"encoding"`
	// The Live Event input.
	Input LiveEventInputResponseOutput `pulumi:"input"`
	// The exact time the Live Event was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Live Event preview.
	Preview LiveEventPreviewResponsePtrOutput `pulumi:"preview"`
	// The provisioning state of the Live Event.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource state of the Live Event.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions pulumi.StringArrayOutput `pulumi:"streamOptions"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
	UseStaticHostname pulumi.BoolPtrOutput `pulumi:"useStaticHostname"`
}

// NewLiveEvent registers a new resource with the given unique name, arguments, and options.
func NewLiveEvent(ctx *pulumi.Context,
	name string, args *LiveEventArgs, opts ...pulumi.ResourceOption) (*LiveEvent, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Input == nil {
		return nil, errors.New("missing required argument 'Input'")
	}
	if args == nil || args.LiveEventName == nil {
		return nil, errors.New("missing required argument 'LiveEventName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LiveEventArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:media/v20180330preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180601preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180701:LiveEvent"),
		},
		{
			Type: pulumi.String("azurerm:media/v20190501preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azurerm:media/v20200501:LiveEvent"),
		},
	})
	opts = append(opts, aliases)
	var resource LiveEvent
	err := ctx.RegisterResource("azurerm:media/latest:LiveEvent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiveEvent gets an existing LiveEvent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiveEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiveEventState, opts ...pulumi.ResourceOption) (*LiveEvent, error) {
	var resource LiveEvent
	err := ctx.ReadResource("azurerm:media/latest:LiveEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiveEvent resources.
type liveEventState struct {
	// The exact time the Live Event was created.
	Created *string `pulumi:"created"`
	// The Live Event access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPoliciesResponse `pulumi:"crossSiteAccessPolicies"`
	// The Live Event description.
	Description *string `pulumi:"description"`
	// The Live Event encoding.
	Encoding *LiveEventEncodingResponse `pulumi:"encoding"`
	// The Live Event input.
	Input *LiveEventInputResponse `pulumi:"input"`
	// The exact time the Live Event was last modified.
	LastModified *string `pulumi:"lastModified"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The Live Event preview.
	Preview *LiveEventPreviewResponse `pulumi:"preview"`
	// The provisioning state of the Live Event.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource state of the Live Event.
	ResourceState *string `pulumi:"resourceState"`
	// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions []string `pulumi:"streamOptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
	UseStaticHostname *bool `pulumi:"useStaticHostname"`
}

type LiveEventState struct {
	// The exact time the Live Event was created.
	Created pulumi.StringPtrInput
	// The Live Event access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrInput
	// The Live Event description.
	Description pulumi.StringPtrInput
	// The Live Event encoding.
	Encoding LiveEventEncodingResponsePtrInput
	// The Live Event input.
	Input LiveEventInputResponsePtrInput
	// The exact time the Live Event was last modified.
	LastModified pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The Live Event preview.
	Preview LiveEventPreviewResponsePtrInput
	// The provisioning state of the Live Event.
	ProvisioningState pulumi.StringPtrInput
	// The resource state of the Live Event.
	ResourceState pulumi.StringPtrInput
	// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
	UseStaticHostname pulumi.BoolPtrInput
}

func (LiveEventState) ElementType() reflect.Type {
	return reflect.TypeOf((*liveEventState)(nil)).Elem()
}

type liveEventArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart *bool `pulumi:"autoStart"`
	// The Live Event access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPolicies `pulumi:"crossSiteAccessPolicies"`
	// The Live Event description.
	Description *string `pulumi:"description"`
	// The Live Event encoding.
	Encoding *LiveEventEncoding `pulumi:"encoding"`
	// The Live Event input.
	Input LiveEventInput `pulumi:"input"`
	// The name of the Live Event.
	LiveEventName string `pulumi:"liveEventName"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The Live Event preview.
	Preview *LiveEventPreview `pulumi:"preview"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions []string `pulumi:"streamOptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
	UseStaticHostname *bool `pulumi:"useStaticHostname"`
}

// The set of arguments for constructing a LiveEvent resource.
type LiveEventArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart pulumi.BoolPtrInput
	// The Live Event access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	// The Live Event description.
	Description pulumi.StringPtrInput
	// The Live Event encoding.
	Encoding LiveEventEncodingPtrInput
	// The Live Event input.
	Input LiveEventInputInput
	// The name of the Live Event.
	LiveEventName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The Live Event preview.
	Preview LiveEventPreviewPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
	UseStaticHostname pulumi.BoolPtrInput
}

func (LiveEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liveEventArgs)(nil)).Elem()
}
