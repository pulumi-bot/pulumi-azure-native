// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A single API Management service resource in List or Get response.
type ApiManagementService struct {
	pulumi.CustomResourceState

	// ETag of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Managed service identity of the Api Management service.
	Identity ApiManagementServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the API Management service.
	Properties ApiManagementServicePropertiesResponseOutput `pulumi:"properties"`
	// SKU properties of the API Management service.
	Sku ApiManagementServiceSkuPropertiesResponseOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApiManagementService registers a new resource with the given unique name, arguments, and options.
func NewApiManagementService(ctx *pulumi.Context,
	name string, args *ApiManagementServiceArgs, opts ...pulumi.ResourceOption) (*ApiManagementService, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
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
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ApiManagementServiceArgs{}
	}
	var resource ApiManagementService
	err := ctx.RegisterResource("azurerm:apimanagement/v20170301:ApiManagementService", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:apimanagement/v20170301:ApiManagementService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiManagementService resources.
type apiManagementServiceState struct {
	// ETag of the resource.
	Etag *string `pulumi:"etag"`
	// Managed service identity of the Api Management service.
	Identity *ApiManagementServiceIdentityResponse `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the API Management service.
	Properties *ApiManagementServicePropertiesResponse `pulumi:"properties"`
	// SKU properties of the API Management service.
	Sku *ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type *string `pulumi:"type"`
}

type ApiManagementServiceState struct {
	// ETag of the resource.
	Etag pulumi.StringPtrInput
	// Managed service identity of the Api Management service.
	Identity ApiManagementServiceIdentityResponsePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the API Management service.
	Properties ApiManagementServicePropertiesResponsePtrInput
	// SKU properties of the API Management service.
	Sku ApiManagementServiceSkuPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type pulumi.StringPtrInput
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
	// The name of the API Management service.
	Name string `pulumi:"name"`
	// Email address from which the notification will be sent.
	NotificationSenderEmail *string `pulumi:"notificationSenderEmail"`
	// Publisher email.
	PublisherEmail string `pulumi:"publisherEmail"`
	// Publisher name.
	PublisherName string `pulumi:"publisherName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
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
	// The name of the API Management service.
	Name pulumi.StringInput
	// Email address from which the notification will be sent.
	NotificationSenderEmail pulumi.StringPtrInput
	// Publisher email.
	PublisherEmail pulumi.StringInput
	// Publisher name.
	PublisherName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
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
