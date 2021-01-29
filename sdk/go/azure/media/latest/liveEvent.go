// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The live event.
// Latest API Version: 2020-05-01.
type LiveEvent struct {
	pulumi.CustomResourceState

	// The creation time for the live event
	Created pulumi.StringOutput `pulumi:"created"`
	// Live event cross site access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput `pulumi:"crossSiteAccessPolicies"`
	// A description for the live event.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding LiveEventEncodingResponsePtrOutput `pulumi:"encoding"`
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix pulumi.StringPtrOutput `pulumi:"hostnamePrefix"`
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input LiveEventInputResponseOutput `pulumi:"input"`
	// The last modified time of the live event.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview LiveEventPreviewResponsePtrOutput `pulumi:"preview"`
	// The provisioning state of the live event.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource state of the live event. See https://go.microsoft.com/fwlink/?linkid=2139012 for more information.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions pulumi.StringArrayOutput `pulumi:"streamOptions"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions LiveEventTranscriptionResponseArrayOutput `pulumi:"transcriptions"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
	UseStaticHostname pulumi.BoolPtrOutput `pulumi:"useStaticHostname"`
}

// NewLiveEvent registers a new resource with the given unique name, arguments, and options.
func NewLiveEvent(ctx *pulumi.Context,
	name string, args *LiveEventArgs, opts ...pulumi.ResourceOption) (*LiveEvent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Input == nil {
		return nil, errors.New("invalid value for required argument 'Input'")
	}
	if args.LiveEventName == nil {
		return nil, errors.New("invalid value for required argument 'LiveEventName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:media/v20180330preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180701:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20190501preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:LiveEvent"),
		},
	})
	opts = append(opts, aliases)
	var resource LiveEvent
	err := ctx.RegisterResource("azure-nextgen:media/latest:LiveEvent", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:media/latest:LiveEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiveEvent resources.
type liveEventState struct {
	// The creation time for the live event
	Created *string `pulumi:"created"`
	// Live event cross site access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPoliciesResponse `pulumi:"crossSiteAccessPolicies"`
	// A description for the live event.
	Description *string `pulumi:"description"`
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding *LiveEventEncodingResponse `pulumi:"encoding"`
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix *string `pulumi:"hostnamePrefix"`
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input *LiveEventInputResponse `pulumi:"input"`
	// The last modified time of the live event.
	LastModified *string `pulumi:"lastModified"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview *LiveEventPreviewResponse `pulumi:"preview"`
	// The provisioning state of the live event.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource state of the live event. See https://go.microsoft.com/fwlink/?linkid=2139012 for more information.
	ResourceState *string `pulumi:"resourceState"`
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions []string `pulumi:"streamOptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions []LiveEventTranscriptionResponse `pulumi:"transcriptions"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
	UseStaticHostname *bool `pulumi:"useStaticHostname"`
}

type LiveEventState struct {
	// The creation time for the live event
	Created pulumi.StringPtrInput
	// Live event cross site access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrInput
	// A description for the live event.
	Description pulumi.StringPtrInput
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding LiveEventEncodingResponsePtrInput
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix pulumi.StringPtrInput
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input LiveEventInputResponsePtrInput
	// The last modified time of the live event.
	LastModified pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview LiveEventPreviewResponsePtrInput
	// The provisioning state of the live event.
	ProvisioningState pulumi.StringPtrInput
	// The resource state of the live event. See https://go.microsoft.com/fwlink/?linkid=2139012 for more information.
	ResourceState pulumi.StringPtrInput
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions LiveEventTranscriptionResponseArrayInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
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
	// Live event cross site access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPolicies `pulumi:"crossSiteAccessPolicies"`
	// A description for the live event.
	Description *string `pulumi:"description"`
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding *LiveEventEncoding `pulumi:"encoding"`
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix *string `pulumi:"hostnamePrefix"`
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input LiveEventInputType `pulumi:"input"`
	// The name of the live event, maximum length is 32.
	LiveEventName string `pulumi:"liveEventName"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview *LiveEventPreview `pulumi:"preview"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions []string `pulumi:"streamOptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions []LiveEventTranscription `pulumi:"transcriptions"`
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
	UseStaticHostname *bool `pulumi:"useStaticHostname"`
}

// The set of arguments for constructing a LiveEvent resource.
type LiveEventArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart pulumi.BoolPtrInput
	// Live event cross site access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	// A description for the live event.
	Description pulumi.StringPtrInput
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding LiveEventEncodingPtrInput
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix pulumi.StringPtrInput
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input LiveEventInputTypeInput
	// The name of the live event, maximum length is 32.
	LiveEventName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview LiveEventPreviewPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions LiveEventTranscriptionArrayInput
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
	UseStaticHostname pulumi.BoolPtrInput
}

func (LiveEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liveEventArgs)(nil)).Elem()
}

type LiveEventInput interface {
	pulumi.Input

	ToLiveEventOutput() LiveEventOutput
	ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput
}

func (*LiveEvent) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEvent)(nil))
}

func (i *LiveEvent) ToLiveEventOutput() LiveEventOutput {
	return i.ToLiveEventOutputWithContext(context.Background())
}

func (i *LiveEvent) ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutput)
}

type LiveEventOutput struct {
	*pulumi.OutputState
}

func (LiveEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEvent)(nil))
}

func (o LiveEventOutput) ToLiveEventOutput() LiveEventOutput {
	return o
}

func (o LiveEventOutput) ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LiveEventOutput{})
}
