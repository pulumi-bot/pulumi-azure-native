// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A single API Management service resource in List or Get response.
//
// ## Example Usage
// ### ApiManagementCreateMultiRegionServiceWithCustomHostname
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20180601preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
// 			AdditionalLocations: apimanagement.AdditionalLocationArray{
// 				&apimanagement.AdditionalLocationArgs{
// 					Location: pulumi.String("West US"),
// 					Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
// 						Capacity: pulumi.Int(1),
// 						Name:     pulumi.String("Premium"),
// 					},
// 					VirtualNetworkConfiguration: &apimanagement.VirtualNetworkConfigurationArgs{
// 						SubnetResourceId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/westUsVirtualNetwork/subnets/apimSubnet"),
// 					},
// 				},
// 			},
// 			HostnameConfigurations: apimanagement.HostnameConfigurationArray{
// 				&apimanagement.HostnameConfigurationArgs{
// 					CertificatePassword: pulumi.String("**************Password of the Certificate************************************************"),
// 					EncodedCertificate:  pulumi.String("************Base 64 Encoded Pfx Certificate************************"),
// 					HostName:            pulumi.String("proxyhostname1.contoso.com"),
// 					Type:                pulumi.String("Proxy"),
// 				},
// 				&apimanagement.HostnameConfigurationArgs{
// 					CertificatePassword:        pulumi.String("**************Password of the Certificate************************************************"),
// 					EncodedCertificate:         pulumi.String("************Base 64 Encoded Pfx Certificate************************"),
// 					HostName:                   pulumi.String("proxyhostname2.contoso.com"),
// 					NegotiateClientCertificate: pulumi.Bool(true),
// 					Type:                       pulumi.String("Proxy"),
// 				},
// 				&apimanagement.HostnameConfigurationArgs{
// 					CertificatePassword: pulumi.String("**************Password of the Certificate************************************************"),
// 					EncodedCertificate:  pulumi.String("************Base 64 Encoded Pfx Certificate************************"),
// 					HostName:            pulumi.String("portalhostname1.contoso.com"),
// 					Type:                pulumi.String("Portal"),
// 				},
// 			},
// 			Location:          pulumi.String("Central US"),
// 			PublisherEmail:    pulumi.String("admin@live.com"),
// 			PublisherName:     pulumi.String("contoso"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
// 				Capacity: pulumi.Int(1),
// 				Name:     pulumi.String("Premium"),
// 			},
// 			VirtualNetworkConfiguration: &apimanagement.VirtualNetworkConfigurationArgs{
// 				SubnetResourceId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/centralUsVirtualNetwork/subnets/apimSubnet"),
// 			},
// 			VirtualNetworkType: pulumi.String("External"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateService
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20180601preview"
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
// ### ApiManagementCreateServiceHavingMsi
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20180601preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
// 			Identity: &apimanagement.ApiManagementServiceIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 			Location:          pulumi.String("Japan East"),
// 			PublisherEmail:    pulumi.String("admin@contoso.com"),
// 			PublisherName:     pulumi.String("Contoso"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
// 				Name: pulumi.String("Developer"),
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
// ### ApiManagementCreateServiceWithSystemCertificates
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20180601preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
// 			Certificates: apimanagement.CertificateConfigurationArray{
// 				&apimanagement.CertificateConfigurationArgs{
// 					CertificatePassword: pulumi.String("**************Password of the Certificate************************************************"),
// 					EncodedCertificate:  pulumi.String("************Base 64 Encoded Pfx Certificate************************"),
// 					StoreName:           pulumi.String("CertificateAuthority"),
// 				},
// 			},
// 			Location:          pulumi.String("Central US"),
// 			PublisherEmail:    pulumi.String("apim@autorestsdk.com"),
// 			PublisherName:     pulumi.String("autorestsdk"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
// 				Capacity: pulumi.Int(1),
// 				Name:     pulumi.String("Basic"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("value1"),
// 				"tag2": pulumi.String("value2"),
// 				"tag3": pulumi.String("value3"),
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
	AdditionalLocations AdditionalLocationResponseArrayOutput `pulumi:"additionalLocations"`
	// List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
	Certificates CertificateConfigurationResponseArrayOutput `pulumi:"certificates"`
	// Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedAtUtc pulumi.StringOutput `pulumi:"createdAtUtc"`
	// Custom properties of the API Management service. Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2). Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1 and setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.
	CustomProperties pulumi.StringMapOutput `pulumi:"customProperties"`
	// ETag of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Gateway URL of the API Management service in the Default Region.
	GatewayRegionalUrl pulumi.StringOutput `pulumi:"gatewayRegionalUrl"`
	// Gateway URL of the API Management service.
	GatewayUrl pulumi.StringOutput `pulumi:"gatewayUrl"`
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations HostnameConfigurationResponseArrayOutput `pulumi:"hostnameConfigurations"`
	// Managed service identity of the Api Management service.
	Identity ApiManagementServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Management API endpoint URL of the API Management service.
	ManagementApiUrl pulumi.StringOutput `pulumi:"managementApiUrl"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Email address from which the notification will be sent.
	NotificationSenderEmail pulumi.StringPtrOutput `pulumi:"notificationSenderEmail"`
	// Publisher portal endpoint Url of the API Management service.
	PortalUrl pulumi.StringOutput `pulumi:"portalUrl"`
	// Private Static Load Balanced IP addresses of the API Management service in Primary region which is deployed in an Internal Virtual Network. Available only for Basic, Standard and Premium SKU.
	PrivateIPAddresses pulumi.StringArrayOutput `pulumi:"privateIPAddresses"`
	// The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Public Static Load Balanced IP addresses of the API Management service in Primary region. Available only for Basic, Standard and Premium SKU.
	PublicIPAddresses pulumi.StringArrayOutput `pulumi:"publicIPAddresses"`
	// Publisher email.
	PublisherEmail pulumi.StringOutput `pulumi:"publisherEmail"`
	// Publisher name.
	PublisherName pulumi.StringOutput `pulumi:"publisherName"`
	// SCM endpoint URL of the API Management service.
	ScmUrl pulumi.StringOutput `pulumi:"scmUrl"`
	// SKU properties of the API Management service.
	Sku ApiManagementServiceSkuPropertiesResponseOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	TargetProvisioningState pulumi.StringOutput `pulumi:"targetProvisioningState"`
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type pulumi.StringOutput `pulumi:"type"`
	// Virtual network configuration of the API Management service.
	VirtualNetworkConfiguration VirtualNetworkConfigurationResponsePtrOutput `pulumi:"virtualNetworkConfiguration"`
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VirtualNetworkType pulumi.StringPtrOutput `pulumi:"virtualNetworkType"`
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
			Type: pulumi.String("azurerm:apimanagement/v20161010:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:ApiManagementService"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ApiManagementService"),
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
	err := ctx.RegisterResource("azurerm:apimanagement/v20180601preview:ApiManagementService", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:apimanagement/v20180601preview:ApiManagementService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiManagementService resources.
type apiManagementServiceState struct {
	// Additional datacenter locations of the API Management service.
	AdditionalLocations []AdditionalLocationResponse `pulumi:"additionalLocations"`
	// List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
	Certificates []CertificateConfigurationResponse `pulumi:"certificates"`
	// Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedAtUtc *string `pulumi:"createdAtUtc"`
	// Custom properties of the API Management service. Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2). Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1 and setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.
	CustomProperties map[string]string `pulumi:"customProperties"`
	// ETag of the resource.
	Etag *string `pulumi:"etag"`
	// Gateway URL of the API Management service in the Default Region.
	GatewayRegionalUrl *string `pulumi:"gatewayRegionalUrl"`
	// Gateway URL of the API Management service.
	GatewayUrl *string `pulumi:"gatewayUrl"`
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations []HostnameConfigurationResponse `pulumi:"hostnameConfigurations"`
	// Managed service identity of the Api Management service.
	Identity *ApiManagementServiceIdentityResponse `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Management API endpoint URL of the API Management service.
	ManagementApiUrl *string `pulumi:"managementApiUrl"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Email address from which the notification will be sent.
	NotificationSenderEmail *string `pulumi:"notificationSenderEmail"`
	// Publisher portal endpoint Url of the API Management service.
	PortalUrl *string `pulumi:"portalUrl"`
	// Private Static Load Balanced IP addresses of the API Management service in Primary region which is deployed in an Internal Virtual Network. Available only for Basic, Standard and Premium SKU.
	PrivateIPAddresses []string `pulumi:"privateIPAddresses"`
	// The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Public Static Load Balanced IP addresses of the API Management service in Primary region. Available only for Basic, Standard and Premium SKU.
	PublicIPAddresses []string `pulumi:"publicIPAddresses"`
	// Publisher email.
	PublisherEmail *string `pulumi:"publisherEmail"`
	// Publisher name.
	PublisherName *string `pulumi:"publisherName"`
	// SCM endpoint URL of the API Management service.
	ScmUrl *string `pulumi:"scmUrl"`
	// SKU properties of the API Management service.
	Sku *ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	TargetProvisioningState *string `pulumi:"targetProvisioningState"`
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type *string `pulumi:"type"`
	// Virtual network configuration of the API Management service.
	VirtualNetworkConfiguration *VirtualNetworkConfigurationResponse `pulumi:"virtualNetworkConfiguration"`
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VirtualNetworkType *string `pulumi:"virtualNetworkType"`
}

type ApiManagementServiceState struct {
	// Additional datacenter locations of the API Management service.
	AdditionalLocations AdditionalLocationResponseArrayInput
	// List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
	Certificates CertificateConfigurationResponseArrayInput
	// Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedAtUtc pulumi.StringPtrInput
	// Custom properties of the API Management service. Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2). Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1 and setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.
	CustomProperties pulumi.StringMapInput
	// ETag of the resource.
	Etag pulumi.StringPtrInput
	// Gateway URL of the API Management service in the Default Region.
	GatewayRegionalUrl pulumi.StringPtrInput
	// Gateway URL of the API Management service.
	GatewayUrl pulumi.StringPtrInput
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations HostnameConfigurationResponseArrayInput
	// Managed service identity of the Api Management service.
	Identity ApiManagementServiceIdentityResponsePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Management API endpoint URL of the API Management service.
	ManagementApiUrl pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Email address from which the notification will be sent.
	NotificationSenderEmail pulumi.StringPtrInput
	// Publisher portal endpoint Url of the API Management service.
	PortalUrl pulumi.StringPtrInput
	// Private Static Load Balanced IP addresses of the API Management service in Primary region which is deployed in an Internal Virtual Network. Available only for Basic, Standard and Premium SKU.
	PrivateIPAddresses pulumi.StringArrayInput
	// The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	ProvisioningState pulumi.StringPtrInput
	// Public Static Load Balanced IP addresses of the API Management service in Primary region. Available only for Basic, Standard and Premium SKU.
	PublicIPAddresses pulumi.StringArrayInput
	// Publisher email.
	PublisherEmail pulumi.StringPtrInput
	// Publisher name.
	PublisherName pulumi.StringPtrInput
	// SCM endpoint URL of the API Management service.
	ScmUrl pulumi.StringPtrInput
	// SKU properties of the API Management service.
	Sku ApiManagementServiceSkuPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	TargetProvisioningState pulumi.StringPtrInput
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type pulumi.StringPtrInput
	// Virtual network configuration of the API Management service.
	VirtualNetworkConfiguration VirtualNetworkConfigurationResponsePtrInput
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VirtualNetworkType pulumi.StringPtrInput
}

func (ApiManagementServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiManagementServiceState)(nil)).Elem()
}

type apiManagementServiceArgs struct {
	// Additional datacenter locations of the API Management service.
	AdditionalLocations []AdditionalLocation `pulumi:"additionalLocations"`
	// List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
	Certificates []CertificateConfiguration `pulumi:"certificates"`
	// Custom properties of the API Management service. Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2). Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1 and setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.
	CustomProperties map[string]string `pulumi:"customProperties"`
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations []HostnameConfiguration `pulumi:"hostnameConfigurations"`
	// Managed service identity of the Api Management service.
	Identity *ApiManagementServiceIdentity `pulumi:"identity"`
	// Resource location.
	Location string `pulumi:"location"`
	// Email address from which the notification will be sent.
	NotificationSenderEmail *string `pulumi:"notificationSenderEmail"`
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
	// Virtual network configuration of the API Management service.
	VirtualNetworkConfiguration *VirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VirtualNetworkType *string `pulumi:"virtualNetworkType"`
}

// The set of arguments for constructing a ApiManagementService resource.
type ApiManagementServiceArgs struct {
	// Additional datacenter locations of the API Management service.
	AdditionalLocations AdditionalLocationArrayInput
	// List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
	Certificates CertificateConfigurationArrayInput
	// Custom properties of the API Management service. Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2). Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1 and setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.
	CustomProperties pulumi.StringMapInput
	// Custom hostname configuration of the API Management service.
	HostnameConfigurations HostnameConfigurationArrayInput
	// Managed service identity of the Api Management service.
	Identity ApiManagementServiceIdentityPtrInput
	// Resource location.
	Location pulumi.StringInput
	// Email address from which the notification will be sent.
	NotificationSenderEmail pulumi.StringPtrInput
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
	// Virtual network configuration of the API Management service.
	VirtualNetworkConfiguration VirtualNetworkConfigurationPtrInput
	// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VirtualNetworkType pulumi.StringPtrInput
}

func (ApiManagementServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiManagementServiceArgs)(nil)).Elem()
}
