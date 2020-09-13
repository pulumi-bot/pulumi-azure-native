// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Backend details.
//
// ## Example Usage
// ### ApiManagementCreateBackendProxyBackend
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
// 		_, err := apimanagement.NewBackend(ctx, "backend", &apimanagement.BackendArgs{
// 			BackendId: pulumi.String("proxybackend"),
// 			Credentials: &apimanagement.BackendCredentialsContractArgs{
// 				Authorization: &apimanagement.BackendAuthorizationHeaderCredentialsArgs{
// 					Parameter: pulumi.String("opensesma"),
// 					Scheme:    pulumi.String("Basic"),
// 				},
// 				Header: pulumi.StringArrayMap{
// 					"x-my-1": pulumi.StringArray{
// 						pulumi.String("val1"),
// 						pulumi.String("val2"),
// 					},
// 				},
// 				Query: pulumi.StringArrayMap{
// 					"sv": pulumi.StringArray{
// 						pulumi.String("xx"),
// 						pulumi.String("bb"),
// 						pulumi.String("cc"),
// 					},
// 				},
// 			},
// 			Description: pulumi.String("description5308"),
// 			Protocol:    pulumi.String("http"),
// 			Proxy: &apimanagement.BackendProxyContractArgs{
// 				Password: pulumi.String("opensesame"),
// 				Url:      pulumi.String("http://192.168.1.1:8080"),
// 				Username: pulumi.String("Contoso\\admin"),
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Tls: &apimanagement.BackendTlsPropertiesArgs{
// 				ValidateCertificateChain: pulumi.Bool(true),
// 				ValidateCertificateName:  pulumi.Bool(true),
// 			},
// 			Url: pulumi.String("https://backendname2644/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateBackendServiceFabric
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
// 		_, err := apimanagement.NewBackend(ctx, "backend", &apimanagement.BackendArgs{
// 			BackendId:   pulumi.String("sfbackend"),
// 			Description: pulumi.String("Service Fabric Test App 1"),
// 			Properties: &apimanagement.BackendPropertiesArgs{
// 				ServiceFabricCluster: &apimanagement.BackendServiceFabricClusterPropertiesArgs{
// 					ClientCertificatethumbprint: pulumi.String("EBA029198AA3E76EF0D70482626E5BCF148594A6"),
// 					ManagementEndpoints: pulumi.StringArray{
// 						pulumi.String("https://somecluster.com"),
// 					},
// 					MaxPartitionResolutionRetries: pulumi.Int(5),
// 					ServerX509Names: apimanagement.X509CertificateNameArray{
// 						&apimanagement.X509CertificateNameArgs{
// 							IssuerCertificateThumbprint: pulumi.String("IssuerCertificateThumbprint1"),
// 							Name:                        pulumi.String("ServerCommonName1"),
// 						},
// 					},
// 				},
// 			},
// 			Protocol:          pulumi.String("http"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Url:               pulumi.String("fabric:/mytestapp/mytestservice"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Backend struct {
	pulumi.CustomResourceState

	// Backend Credentials Contract Properties
	Credentials BackendCredentialsContractResponsePtrOutput `pulumi:"credentials"`
	// Backend Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Backend Properties contract
	Properties BackendPropertiesResponseOutput `pulumi:"properties"`
	// Backend communication protocol.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Backend Proxy Contract Properties
	Proxy BackendProxyContractResponsePtrOutput `pulumi:"proxy"`
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Backend Title.
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// Backend TLS Properties
	Tls BackendTlsPropertiesResponsePtrOutput `pulumi:"tls"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Runtime Url of the Backend.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil || args.BackendId == nil {
		return nil, errors.New("missing required argument 'BackendId'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil {
		args = &BackendArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:Backend"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20160707:Backend"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20161010:Backend"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:Backend"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:Backend"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:Backend"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:Backend"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:Backend"),
		},
	})
	opts = append(opts, aliases)
	var resource Backend
	err := ctx.RegisterResource("azurerm:apimanagement/v20180601preview:Backend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackend gets an existing Backend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendState, opts ...pulumi.ResourceOption) (*Backend, error) {
	var resource Backend
	err := ctx.ReadResource("azurerm:apimanagement/v20180601preview:Backend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backend resources.
type backendState struct {
	// Backend Credentials Contract Properties
	Credentials *BackendCredentialsContractResponse `pulumi:"credentials"`
	// Backend Description.
	Description *string `pulumi:"description"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Backend Properties contract
	Properties *BackendPropertiesResponse `pulumi:"properties"`
	// Backend communication protocol.
	Protocol *string `pulumi:"protocol"`
	// Backend Proxy Contract Properties
	Proxy *BackendProxyContractResponse `pulumi:"proxy"`
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId *string `pulumi:"resourceId"`
	// Backend Title.
	Title *string `pulumi:"title"`
	// Backend TLS Properties
	Tls *BackendTlsPropertiesResponse `pulumi:"tls"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
	// Runtime Url of the Backend.
	Url *string `pulumi:"url"`
}

type BackendState struct {
	// Backend Credentials Contract Properties
	Credentials BackendCredentialsContractResponsePtrInput
	// Backend Description.
	Description pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Backend Properties contract
	Properties BackendPropertiesResponsePtrInput
	// Backend communication protocol.
	Protocol pulumi.StringPtrInput
	// Backend Proxy Contract Properties
	Proxy BackendProxyContractResponsePtrInput
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId pulumi.StringPtrInput
	// Backend Title.
	Title pulumi.StringPtrInput
	// Backend TLS Properties
	Tls BackendTlsPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
	// Runtime Url of the Backend.
	Url pulumi.StringPtrInput
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	// Identifier of the Backend entity. Must be unique in the current API Management service instance.
	BackendId string `pulumi:"backendId"`
	// Backend Credentials Contract Properties
	Credentials *BackendCredentialsContract `pulumi:"credentials"`
	// Backend Description.
	Description *string `pulumi:"description"`
	// Backend Properties contract
	Properties *BackendProperties `pulumi:"properties"`
	// Backend communication protocol.
	Protocol string `pulumi:"protocol"`
	// Backend Proxy Contract Properties
	Proxy *BackendProxyContract `pulumi:"proxy"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId *string `pulumi:"resourceId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Backend Title.
	Title *string `pulumi:"title"`
	// Backend TLS Properties
	Tls *BackendTlsProperties `pulumi:"tls"`
	// Runtime Url of the Backend.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// Identifier of the Backend entity. Must be unique in the current API Management service instance.
	BackendId pulumi.StringInput
	// Backend Credentials Contract Properties
	Credentials BackendCredentialsContractPtrInput
	// Backend Description.
	Description pulumi.StringPtrInput
	// Backend Properties contract
	Properties BackendPropertiesPtrInput
	// Backend communication protocol.
	Protocol pulumi.StringInput
	// Backend Proxy Contract Properties
	Proxy BackendProxyContractPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Backend Title.
	Title pulumi.StringPtrInput
	// Backend TLS Properties
	Tls BackendTlsPropertiesPtrInput
	// Runtime Url of the Backend.
	Url pulumi.StringInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}
