// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161010

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A single API Management service resource in List or Get response.
//
// ## Example Usage
// ### ApiManagementCreateService
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20161010"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
// 			Location:          pulumi.String("West US"),
// 			PublisherEmail:    pulumi.String("admin@live.com"),
// 			PublisherName:     pulumi.String("contoso"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
// 				Capacity: pulumi.Int(1),
// 				Name:     pulumi.String("Premium"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ApiManagementService struct {
	pulumi.CustomResourceState

	// Additional datacenter locations of the API Management service.
	AdditionalLocations AdditionalRegionResponseArrayOutput `pulumi:"additionalLocations"`
	// Addresser email.
	AddresserEmail pulumi.StringPtrOutput `pulumi:"addresserEmail"`
	// Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedAtUtc pulumi.StringOutput `pulumi:"createdAtUtc"`
	// Custom properties of the API Management service, like disabling TLS 1.0.
	CustomProperties pulumi.StringMapOutput `pulumi:"customProperties"`
	// ETag of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations HostnameConfigurationResponseArrayOutput `pulumi:"hostnameConfigurations"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Management API endpoint URL of the API Management service.
	ManagementApiUrl pulumi.StringOutput `pulumi:"managementApiUrl"`
	// Resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Publisher portal endpoint Url of the API Management service.
	PortalUrl pulumi.StringOutput `pulumi:"portalUrl"`
	// The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Publisher email.
	PublisherEmail pulumi.StringOutput `pulumi:"publisherEmail"`
	// Publisher name.
	PublisherName pulumi.StringOutput `pulumi:"publisherName"`
	// Proxy endpoint URL of the API Management service.
	RuntimeUrl pulumi.StringOutput `pulumi:"runtimeUrl"`
	// SCM endpoint URL of the API Management service.
	ScmUrl pulumi.StringOutput `pulumi:"scmUrl"`
	// SKU properties of the API Management service.
	Sku ApiManagementServiceSkuPropertiesResponseOutput `pulumi:"sku"`
	// Static IP addresses of the API Management service virtual machines. Available only for Standard and Premium SKU.
	StaticIPs pulumi.StringArrayOutput `pulumi:"staticIPs"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	TargetProvisioningState pulumi.StringOutput `pulumi:"targetProvisioningState"`
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type pulumi.StringOutput `pulumi:"type"`
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VpnType pulumi.StringPtrOutput `pulumi:"vpnType"`
	// Virtual network configuration of the API Management service.
	Vpnconfiguration VirtualNetworkConfigurationResponsePtrOutput `pulumi:"vpnconfiguration"`
}

// NewApiManagementService registers a new resource with the given unique name, arguments, and options.
func NewApiManagementService(ctx *pulumi.Context,
	name string, args *ApiManagementServiceArgs, opts ...pulumi.ResourceOption) (*ApiManagementService, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PublisherEmail == nil {
		return nil, errors.New("missing required argument 'PublisherEmail'")
	}
	if args == nil || args.PublisherName == nil {
		return nil, errors.New("missing required argument 'PublisherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ApiManagementServiceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20160707:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:ApiManagementService"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiManagementService
	err := ctx.RegisterResource("azurerm:apimanagement/v20161010:ApiManagementService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiManagementService gets an existing ApiManagementService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiManagementService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiManagementServiceState, opts ...pulumi.ResourceOption) (*ApiManagementService, error) {
	var resource ApiManagementService
	err := ctx.ReadResource("azurerm:apimanagement/v20161010:ApiManagementService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiManagementService resources.
type apiManagementServiceState struct {
	// Additional datacenter locations of the API Management service.
	AdditionalLocations []AdditionalRegionResponse `pulumi:"additionalLocations"`
	// Addresser email.
	AddresserEmail *string `pulumi:"addresserEmail"`
	// Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedAtUtc *string `pulumi:"createdAtUtc"`
	// Custom properties of the API Management service, like disabling TLS 1.0.
	CustomProperties map[string]string `pulumi:"customProperties"`
	// ETag of the resource.
	Etag *string `pulumi:"etag"`
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations []HostnameConfigurationResponse `pulumi:"hostnameConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Management API endpoint URL of the API Management service.
	ManagementApiUrl *string `pulumi:"managementApiUrl"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Publisher portal endpoint Url of the API Management service.
	PortalUrl *string `pulumi:"portalUrl"`
	// The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Publisher email.
	PublisherEmail *string `pulumi:"publisherEmail"`
	// Publisher name.
	PublisherName *string `pulumi:"publisherName"`
	// Proxy endpoint URL of the API Management service.
	RuntimeUrl *string `pulumi:"runtimeUrl"`
	// SCM endpoint URL of the API Management service.
	ScmUrl *string `pulumi:"scmUrl"`
	// SKU properties of the API Management service.
	Sku *ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	// Static IP addresses of the API Management service virtual machines. Available only for Standard and Premium SKU.
	StaticIPs []string `pulumi:"staticIPs"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	TargetProvisioningState *string `pulumi:"targetProvisioningState"`
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type *string `pulumi:"type"`
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VpnType *string `pulumi:"vpnType"`
	// Virtual network configuration of the API Management service.
	Vpnconfiguration *VirtualNetworkConfigurationResponse `pulumi:"vpnconfiguration"`
}

type ApiManagementServiceState struct {
	// Additional datacenter locations of the API Management service.
	AdditionalLocations AdditionalRegionResponseArrayInput
	// Addresser email.
	AddresserEmail pulumi.StringPtrInput
	// Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedAtUtc pulumi.StringPtrInput
	// Custom properties of the API Management service, like disabling TLS 1.0.
	CustomProperties pulumi.StringMapInput
	// ETag of the resource.
	Etag pulumi.StringPtrInput
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations HostnameConfigurationResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Management API endpoint URL of the API Management service.
	ManagementApiUrl pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Publisher portal endpoint Url of the API Management service.
	PortalUrl pulumi.StringPtrInput
	// The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	ProvisioningState pulumi.StringPtrInput
	// Publisher email.
	PublisherEmail pulumi.StringPtrInput
	// Publisher name.
	PublisherName pulumi.StringPtrInput
	// Proxy endpoint URL of the API Management service.
	RuntimeUrl pulumi.StringPtrInput
	// SCM endpoint URL of the API Management service.
	ScmUrl pulumi.StringPtrInput
	// SKU properties of the API Management service.
	Sku ApiManagementServiceSkuPropertiesResponsePtrInput
	// Static IP addresses of the API Management service virtual machines. Available only for Standard and Premium SKU.
	StaticIPs pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	TargetProvisioningState pulumi.StringPtrInput
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type pulumi.StringPtrInput
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VpnType pulumi.StringPtrInput
	// Virtual network configuration of the API Management service.
	Vpnconfiguration VirtualNetworkConfigurationResponsePtrInput
}

func (ApiManagementServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiManagementServiceState)(nil)).Elem()
}

type apiManagementServiceArgs struct {
	// Additional datacenter locations of the API Management service.
	AdditionalLocations []AdditionalRegion `pulumi:"additionalLocations"`
	// Addresser email.
	AddresserEmail *string `pulumi:"addresserEmail"`
	// Custom properties of the API Management service, like disabling TLS 1.0.
	CustomProperties map[string]string `pulumi:"customProperties"`
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations []HostnameConfiguration `pulumi:"hostnameConfigurations"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Publisher email.
	PublisherEmail string `pulumi:"publisherEmail"`
	// Publisher name.
	PublisherName string `pulumi:"publisherName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// SKU properties of the API Management service.
	Sku ApiManagementServiceSkuProperties `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VpnType *string `pulumi:"vpnType"`
	// Virtual network configuration of the API Management service.
	Vpnconfiguration *VirtualNetworkConfiguration `pulumi:"vpnconfiguration"`
}

// The set of arguments for constructing a ApiManagementService resource.
type ApiManagementServiceArgs struct {
	// Additional datacenter locations of the API Management service.
	AdditionalLocations AdditionalRegionArrayInput
	// Addresser email.
	AddresserEmail pulumi.StringPtrInput
	// Custom properties of the API Management service, like disabling TLS 1.0.
	CustomProperties pulumi.StringMapInput
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations HostnameConfigurationArrayInput
	// Resource location.
	Location pulumi.StringInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Publisher email.
	PublisherEmail pulumi.StringInput
	// Publisher name.
	PublisherName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// SKU properties of the API Management service.
	Sku ApiManagementServiceSkuPropertiesInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VpnType pulumi.StringPtrInput
	// Virtual network configuration of the API Management service.
	Vpnconfiguration VirtualNetworkConfigurationPtrInput
}

func (ApiManagementServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiManagementServiceArgs)(nil)).Elem()
}
