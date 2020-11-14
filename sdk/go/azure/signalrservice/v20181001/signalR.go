// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A class represent a SignalR service resource.
type SignalR struct {
	pulumi.CustomResourceState

	// Cross-Origin Resource Sharing (CORS) settings.
	Cors SignalRCorsSettingsResponsePtrOutput `pulumi:"cors"`
	// The publicly accessible IP of the SignalR service.
	ExternalIP pulumi.StringOutput `pulumi:"externalIP"`
	// List of SignalR featureFlags. e.g. ServiceMode.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, SignalR service will use its globally default value.
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features SignalRFeatureResponseArrayOutput `pulumi:"features"`
	// FQDN of the SignalR service instance. Format: xxx.service.signalr.net
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Prefix for the hostName of the SignalR service. Retained for future use.
	// The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
	HostNamePrefix pulumi.StringPtrOutput `pulumi:"hostNamePrefix"`
	// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The publicly accessible port of the SignalR service which is designed for browser/client side usage.
	PublicPort pulumi.IntOutput `pulumi:"publicPort"`
	// The publicly accessible port of the SignalR service which is designed for customer server side usage.
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// SKU of the service.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the service - e.g. "Microsoft.SignalRService/SignalR"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewSignalR registers a new resource with the given unique name, arguments, and options.
func NewSignalR(ctx *pulumi.Context,
	name string, args *SignalRArgs, opts ...pulumi.ResourceOption) (*SignalR, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &SignalRArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:signalrservice/latest:SignalR"),
		},
		{
			Type: pulumi.String("azure-nextgen:signalrservice/v20180301preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-nextgen:signalrservice/v20200501:SignalR"),
		},
		{
			Type: pulumi.String("azure-nextgen:signalrservice/v20200701preview:SignalR"),
		},
	})
	opts = append(opts, aliases)
	var resource SignalR
	err := ctx.RegisterResource("azure-nextgen:signalrservice/v20181001:SignalR", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignalR gets an existing SignalR resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignalR(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRState, opts ...pulumi.ResourceOption) (*SignalR, error) {
	var resource SignalR
	err := ctx.ReadResource("azure-nextgen:signalrservice/v20181001:SignalR", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SignalR resources.
type signalRState struct {
	// Cross-Origin Resource Sharing (CORS) settings.
	Cors *SignalRCorsSettingsResponse `pulumi:"cors"`
	// The publicly accessible IP of the SignalR service.
	ExternalIP *string `pulumi:"externalIP"`
	// List of SignalR featureFlags. e.g. ServiceMode.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, SignalR service will use its globally default value.
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features []SignalRFeatureResponse `pulumi:"features"`
	// FQDN of the SignalR service instance. Format: xxx.service.signalr.net
	HostName *string `pulumi:"hostName"`
	// Prefix for the hostName of the SignalR service. Retained for future use.
	// The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
	HostNamePrefix *string `pulumi:"hostNamePrefix"`
	// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The publicly accessible port of the SignalR service which is designed for browser/client side usage.
	PublicPort *int `pulumi:"publicPort"`
	// The publicly accessible port of the SignalR service which is designed for customer server side usage.
	ServerPort *int `pulumi:"serverPort"`
	// SKU of the service.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the service - e.g. "Microsoft.SignalRService/SignalR"
	Type *string `pulumi:"type"`
	// Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
	Version *string `pulumi:"version"`
}

type SignalRState struct {
	// Cross-Origin Resource Sharing (CORS) settings.
	Cors SignalRCorsSettingsResponsePtrInput
	// The publicly accessible IP of the SignalR service.
	ExternalIP pulumi.StringPtrInput
	// List of SignalR featureFlags. e.g. ServiceMode.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, SignalR service will use its globally default value.
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features SignalRFeatureResponseArrayInput
	// FQDN of the SignalR service instance. Format: xxx.service.signalr.net
	HostName pulumi.StringPtrInput
	// Prefix for the hostName of the SignalR service. Retained for future use.
	// The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
	HostNamePrefix pulumi.StringPtrInput
	// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The publicly accessible port of the SignalR service which is designed for browser/client side usage.
	PublicPort pulumi.IntPtrInput
	// The publicly accessible port of the SignalR service which is designed for customer server side usage.
	ServerPort pulumi.IntPtrInput
	// SKU of the service.
	Sku ResourceSkuResponsePtrInput
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
	// The type of the service - e.g. "Microsoft.SignalRService/SignalR"
	Type pulumi.StringPtrInput
	// Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
	Version pulumi.StringPtrInput
}

func (SignalRState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRState)(nil)).Elem()
}

type signalRArgs struct {
	// Azure GEO region: e.g. West US | East US | North Central US | South Central US | West Europe | North Europe | East Asia | Southeast Asia | etc.
	// The geo region of a resource never changes after it is created.
	Location string `pulumi:"location"`
	// Settings used to provision or configure the resource
	Properties *SignalRCreateOrUpdateProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SignalR resource.
	ResourceName string `pulumi:"resourceName"`
	// The billing information of the resource.(e.g. basic vs. standard)
	Sku *ResourceSku `pulumi:"sku"`
	// A list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SignalR resource.
type SignalRArgs struct {
	// Azure GEO region: e.g. West US | East US | North Central US | South Central US | West Europe | North Europe | East Asia | Southeast Asia | etc.
	// The geo region of a resource never changes after it is created.
	Location pulumi.StringInput
	// Settings used to provision or configure the resource
	Properties SignalRCreateOrUpdatePropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the SignalR resource.
	ResourceName pulumi.StringInput
	// The billing information of the resource.(e.g. basic vs. standard)
	Sku ResourceSkuPtrInput
	// A list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
}

func (SignalRArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRArgs)(nil)).Elem()
}

type SignalRInput interface {
	pulumi.Input

	ToSignalROutput() SignalROutput
	ToSignalROutputWithContext(ctx context.Context) SignalROutput
}

func (SignalR) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalR)(nil)).Elem()
}

func (i SignalR) ToSignalROutput() SignalROutput {
	return i.ToSignalROutputWithContext(context.Background())
}

func (i SignalR) ToSignalROutputWithContext(ctx context.Context) SignalROutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalROutput)
}

type SignalROutput struct {
	*pulumi.OutputState
}

func (SignalROutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalROutput)(nil)).Elem()
}

func (o SignalROutput) ToSignalROutput() SignalROutput {
	return o
}

func (o SignalROutput) ToSignalROutputWithContext(ctx context.Context) SignalROutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SignalROutput{})
}
