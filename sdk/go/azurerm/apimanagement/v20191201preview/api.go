// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Api details.
//
// ## Example Usage
// ### ApiManagementCreateApi
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId: pulumi.String("tempgroup"),
// 			AuthenticationSettings: &apimanagement.AuthenticationSettingsContractArgs{
// 				OAuth2: &apimanagement.OAuth2AuthenticationSettingsContractArgs{
// 					AuthorizationServerId: pulumi.String("authorizationServerId2283"),
// 					Scope:                 pulumi.String("oauth2scope2580"),
// 				},
// 			},
// 			Description: pulumi.String("apidescription5200"),
// 			DisplayName: pulumi.String("apiname1463"),
// 			Path:        pulumi.String("newapiPath"),
// 			Protocols: pulumi.StringArray{
// 				pulumi.String("https"),
// 				pulumi.String("http"),
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			ServiceUrl:        pulumi.String("http://newechoapi.cloudapp.net/api"),
// 			SubscriptionKeyParameterNames: &apimanagement.SubscriptionKeyParameterNamesContractArgs{
// 				Header: pulumi.String("header4520"),
// 				Query:  pulumi.String("query3037"),
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
// ### ApiManagementCreateApiClone
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:       pulumi.String("echo-api2"),
// 			Description: pulumi.String("Copy of Existing Echo Api including Operations."),
// 			DisplayName: pulumi.String("Echo API2"),
// 			IsCurrent:   pulumi.Bool(true),
// 			Path:        pulumi.String("echo2"),
// 			Protocols: pulumi.StringArray{
// 				pulumi.String("http"),
// 				pulumi.String("https"),
// 			},
// 			ResourceGroupName:    pulumi.String("rg1"),
// 			ServiceName:          pulumi.String("apimService1"),
// 			ServiceUrl:           pulumi.String("http://echoapi.cloudapp.net/api"),
// 			SourceApiId:          pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/58a4aeac497000007d040001"),
// 			SubscriptionRequired: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateApiNewVersionUsingExistingApi
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:           pulumi.String("echoapiv3"),
// 			ApiVersion:      pulumi.String("v4"),
// 			ApiVersionSetId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458"),
// 			Description:     pulumi.String("Create Echo API into a new Version using Existing Version Set and Copy all Operations."),
// 			DisplayName:     pulumi.String("Echo API2"),
// 			IsCurrent:       pulumi.Bool(true),
// 			Path:            pulumi.String("echo2"),
// 			Protocols: pulumi.StringArray{
// 				pulumi.String("http"),
// 				pulumi.String("https"),
// 			},
// 			ResourceGroupName:    pulumi.String("rg1"),
// 			ServiceName:          pulumi.String("apimService1"),
// 			ServiceUrl:           pulumi.String("http://echoapi.cloudapp.net/api"),
// 			SourceApiId:          pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echoPath"),
// 			SubscriptionRequired: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateApiRevisionFromExistingApi
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:                  pulumi.String("echo-api;rev=3"),
// 			ApiRevisionDescription: pulumi.String("Creating a Revision of an existing API"),
// 			Path:                   pulumi.String("echo"),
// 			ResourceGroupName:      pulumi.String("rg1"),
// 			ServiceName:            pulumi.String("apimService1"),
// 			ServiceUrl:             pulumi.String("http://echoapi.cloudapp.net/apiv3"),
// 			SourceApiId:            pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateApiUsingImportOverrideServiceUrl
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:             pulumi.String("apidocs"),
// 			Format:            pulumi.String("swagger-link"),
// 			Path:              pulumi.String("petstoreapi123"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			ServiceUrl:        pulumi.String("http://petstore.swagger.wordnik.com/api"),
// 			Value:             pulumi.String("http://apimpimportviaurl.azurewebsites.net/api/apidocs/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateApiUsingOai3Import
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:             pulumi.String("petstore"),
// 			Format:            pulumi.String("openapi-link"),
// 			Path:              pulumi.String("petstore"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Value:             pulumi.String("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v3.0/petstore.yaml"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateApiUsingSwaggerImport
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:             pulumi.String("petstore"),
// 			Format:            pulumi.String("swagger-link-json"),
// 			Path:              pulumi.String("petstore"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Value:             pulumi.String("http://petstore.swagger.io/v2/swagger.json"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateApiUsingWadlImport
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:             pulumi.String("petstore"),
// 			Format:            pulumi.String("wadl-link-json"),
// 			Path:              pulumi.String("collector"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Value:             pulumi.String("https://developer.cisco.com/media/wae-release-6-2-api-reference/wae-collector-rest-api/application.wadl"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### ApiManagementCreateApiWithOpenIdConnect
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId: pulumi.String("tempgroup"),
// 			AuthenticationSettings: &apimanagement.AuthenticationSettingsContractArgs{
// 				Openid: &apimanagement.OpenIdAuthenticationSettingsContractArgs{
// 					BearerTokenSendingMethods: pulumi.StringArray{
// 						pulumi.String("authorizationHeader"),
// 					},
// 					OpenidProviderId: pulumi.String("testopenid"),
// 				},
// 			},
// 			Description: pulumi.String("This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters."),
// 			DisplayName: pulumi.String("Swagger Petstore"),
// 			Path:        pulumi.String("petstore"),
// 			Protocols: pulumi.StringArray{
// 				pulumi.String("https"),
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			ServiceUrl:        pulumi.String("http://petstore.swagger.io/v2"),
// 			SubscriptionKeyParameterNames: &apimanagement.SubscriptionKeyParameterNamesContractArgs{
// 				Header: pulumi.String("Ocp-Apim-Subscription-Key"),
// 				Query:  pulumi.String("subscription-key"),
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
// ### ApiManagementCreateSoapPassThroughApiUsingWsdlImport
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:             pulumi.String("soapApi"),
// 			Format:            pulumi.String("wsdl-link"),
// 			Path:              pulumi.String("currency"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			SoapApiType:       pulumi.String("soap"),
// 			Value:             pulumi.String("http://www.webservicex.net/CurrencyConvertor.asmx?WSDL"),
// 			WsdlSelector: &apimanagement.ApiCreateOrUpdatePropertiesWsdlSelectorArgs{
// 				WsdlEndpointName: pulumi.String("CurrencyConvertorSoap"),
// 				WsdlServiceName:  pulumi.String("CurrencyConvertor"),
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
// ### ApiManagementCreateSoapToRestApiUsingWsdlImport
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
// 			ApiId:             pulumi.String("soapApi"),
// 			Format:            pulumi.String("wsdl-link"),
// 			Path:              pulumi.String("currency"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Value:             pulumi.String("http://www.webservicex.net/CurrencyConvertor.asmx?WSDL"),
// 			WsdlSelector: &apimanagement.ApiCreateOrUpdatePropertiesWsdlSelectorArgs{
// 				WsdlEndpointName: pulumi.String("CurrencyConvertorSoap"),
// 				WsdlServiceName:  pulumi.String("CurrencyConvertor"),
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
type Api struct {
	pulumi.CustomResourceState

	// Describes the Revision of the Api. If no value is provided, default revision 1 is created
	ApiRevision pulumi.StringPtrOutput `pulumi:"apiRevision"`
	// Description of the Api Revision.
	ApiRevisionDescription pulumi.StringPtrOutput `pulumi:"apiRevisionDescription"`
	// Type of API.
	ApiType pulumi.StringPtrOutput `pulumi:"apiType"`
	// Indicates the Version identifier of the API if the API is versioned
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Description of the Api Version.
	ApiVersionDescription pulumi.StringPtrOutput `pulumi:"apiVersionDescription"`
	// Version set details
	ApiVersionSet ApiVersionSetContractDetailsResponsePtrOutput `pulumi:"apiVersionSet"`
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId pulumi.StringPtrOutput `pulumi:"apiVersionSetId"`
	// Collection of authentication settings included into this API.
	AuthenticationSettings AuthenticationSettingsContractResponsePtrOutput `pulumi:"authenticationSettings"`
	// Description of the API. May include HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// API name. Must be 1 to 300 characters long.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Indicates if API revision is current api revision.
	IsCurrent pulumi.BoolPtrOutput `pulumi:"isCurrent"`
	// Indicates if API revision is accessible via the gateway.
	IsOnline pulumi.BoolOutput `pulumi:"isOnline"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path pulumi.StringOutput `pulumi:"path"`
	// Describes on which protocols the operations in this API can be invoked.
	Protocols pulumi.StringArrayOutput `pulumi:"protocols"`
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl pulumi.StringPtrOutput `pulumi:"serviceUrl"`
	// API identifier of the source API.
	SourceApiId pulumi.StringPtrOutput `pulumi:"sourceApiId"`
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractResponsePtrOutput `pulumi:"subscriptionKeyParameterNames"`
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired pulumi.BoolPtrOutput `pulumi:"subscriptionRequired"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ApiArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:Api"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20160707:Api"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20161010:Api"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:Api"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:Api"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:Api"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:Api"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:Api"),
		},
	})
	opts = append(opts, aliases)
	var resource Api
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201preview:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("azurerm:apimanagement/v20191201preview:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// Describes the Revision of the Api. If no value is provided, default revision 1 is created
	ApiRevision *string `pulumi:"apiRevision"`
	// Description of the Api Revision.
	ApiRevisionDescription *string `pulumi:"apiRevisionDescription"`
	// Type of API.
	ApiType *string `pulumi:"apiType"`
	// Indicates the Version identifier of the API if the API is versioned
	ApiVersion *string `pulumi:"apiVersion"`
	// Description of the Api Version.
	ApiVersionDescription *string `pulumi:"apiVersionDescription"`
	// Version set details
	ApiVersionSet *ApiVersionSetContractDetailsResponse `pulumi:"apiVersionSet"`
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId *string `pulumi:"apiVersionSetId"`
	// Collection of authentication settings included into this API.
	AuthenticationSettings *AuthenticationSettingsContractResponse `pulumi:"authenticationSettings"`
	// Description of the API. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// API name. Must be 1 to 300 characters long.
	DisplayName *string `pulumi:"displayName"`
	// Indicates if API revision is current api revision.
	IsCurrent *bool `pulumi:"isCurrent"`
	// Indicates if API revision is accessible via the gateway.
	IsOnline *bool `pulumi:"isOnline"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path *string `pulumi:"path"`
	// Describes on which protocols the operations in this API can be invoked.
	Protocols []string `pulumi:"protocols"`
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl *string `pulumi:"serviceUrl"`
	// API identifier of the source API.
	SourceApiId *string `pulumi:"sourceApiId"`
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContractResponse `pulumi:"subscriptionKeyParameterNames"`
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type ApiState struct {
	// Describes the Revision of the Api. If no value is provided, default revision 1 is created
	ApiRevision pulumi.StringPtrInput
	// Description of the Api Revision.
	ApiRevisionDescription pulumi.StringPtrInput
	// Type of API.
	ApiType pulumi.StringPtrInput
	// Indicates the Version identifier of the API if the API is versioned
	ApiVersion pulumi.StringPtrInput
	// Description of the Api Version.
	ApiVersionDescription pulumi.StringPtrInput
	// Version set details
	ApiVersionSet ApiVersionSetContractDetailsResponsePtrInput
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId pulumi.StringPtrInput
	// Collection of authentication settings included into this API.
	AuthenticationSettings AuthenticationSettingsContractResponsePtrInput
	// Description of the API. May include HTML formatting tags.
	Description pulumi.StringPtrInput
	// API name. Must be 1 to 300 characters long.
	DisplayName pulumi.StringPtrInput
	// Indicates if API revision is current api revision.
	IsCurrent pulumi.BoolPtrInput
	// Indicates if API revision is accessible via the gateway.
	IsOnline pulumi.BoolPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path pulumi.StringPtrInput
	// Describes on which protocols the operations in this API can be invoked.
	Protocols pulumi.StringArrayInput
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl pulumi.StringPtrInput
	// API identifier of the source API.
	SourceApiId pulumi.StringPtrInput
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractResponsePtrInput
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired pulumi.BoolPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Describes the Revision of the Api. If no value is provided, default revision 1 is created
	ApiRevision *string `pulumi:"apiRevision"`
	// Description of the Api Revision.
	ApiRevisionDescription *string `pulumi:"apiRevisionDescription"`
	// Type of API.
	ApiType *string `pulumi:"apiType"`
	// Indicates the Version identifier of the API if the API is versioned
	ApiVersion *string `pulumi:"apiVersion"`
	// Description of the Api Version.
	ApiVersionDescription *string `pulumi:"apiVersionDescription"`
	// Version set details
	ApiVersionSet *ApiVersionSetContractDetails `pulumi:"apiVersionSet"`
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId *string `pulumi:"apiVersionSetId"`
	// Collection of authentication settings included into this API.
	AuthenticationSettings *AuthenticationSettingsContract `pulumi:"authenticationSettings"`
	// Description of the API. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// API name. Must be 1 to 300 characters long.
	DisplayName *string `pulumi:"displayName"`
	// Format of the Content in which the API is getting imported.
	Format *string `pulumi:"format"`
	// Indicates if API revision is current api revision.
	IsCurrent *bool `pulumi:"isCurrent"`
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path string `pulumi:"path"`
	// Describes on which protocols the operations in this API can be invoked.
	Protocols []string `pulumi:"protocols"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl *string `pulumi:"serviceUrl"`
	// Type of Api to create.
	//  * `http` creates a SOAP to REST API
	//  * `soap` creates a SOAP pass-through API .
	SoapApiType *string `pulumi:"soapApiType"`
	// API identifier of the source API.
	SourceApiId *string `pulumi:"sourceApiId"`
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract `pulumi:"subscriptionKeyParameterNames"`
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	// Content value when Importing an API.
	Value *string `pulumi:"value"`
	// Criteria to limit import of WSDL to a subset of the document.
	WsdlSelector *ApiCreateOrUpdatePropertiesWsdlSelector `pulumi:"wsdlSelector"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Describes the Revision of the Api. If no value is provided, default revision 1 is created
	ApiRevision pulumi.StringPtrInput
	// Description of the Api Revision.
	ApiRevisionDescription pulumi.StringPtrInput
	// Type of API.
	ApiType pulumi.StringPtrInput
	// Indicates the Version identifier of the API if the API is versioned
	ApiVersion pulumi.StringPtrInput
	// Description of the Api Version.
	ApiVersionDescription pulumi.StringPtrInput
	// Version set details
	ApiVersionSet ApiVersionSetContractDetailsPtrInput
	// A resource identifier for the related ApiVersionSet.
	ApiVersionSetId pulumi.StringPtrInput
	// Collection of authentication settings included into this API.
	AuthenticationSettings AuthenticationSettingsContractPtrInput
	// Description of the API. May include HTML formatting tags.
	Description pulumi.StringPtrInput
	// API name. Must be 1 to 300 characters long.
	DisplayName pulumi.StringPtrInput
	// Format of the Content in which the API is getting imported.
	Format pulumi.StringPtrInput
	// Indicates if API revision is current api revision.
	IsCurrent pulumi.BoolPtrInput
	// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path pulumi.StringInput
	// Describes on which protocols the operations in this API can be invoked.
	Protocols pulumi.StringArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl pulumi.StringPtrInput
	// Type of Api to create.
	//  * `http` creates a SOAP to REST API
	//  * `soap` creates a SOAP pass-through API .
	SoapApiType pulumi.StringPtrInput
	// API identifier of the source API.
	SourceApiId pulumi.StringPtrInput
	// Protocols over which API is made available.
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractPtrInput
	// Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired pulumi.BoolPtrInput
	// Content value when Importing an API.
	Value pulumi.StringPtrInput
	// Criteria to limit import of WSDL to a subset of the document.
	WsdlSelector ApiCreateOrUpdatePropertiesWsdlSelectorPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}
