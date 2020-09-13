// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Diagnostic details.
//
// ## Example Usage
// ### ApiManagementCreateApiDiagnostic
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApiDiagnostic(ctx, "apiDiagnostic", &apimanagement.ApiDiagnosticArgs{
// 			AlwaysLog: pulumi.String("allErrors"),
// 			ApiId:     pulumi.String("57d1f7558aa04f15146d9d8a"),
// 			Backend: &apimanagement.PipelineDiagnosticSettingsArgs{
// 				Request: &apimanagement.HttpMessageDiagnosticArgs{
// 					Body: &apimanagement.BodyDiagnosticSettingsArgs{
// 						Bytes: pulumi.Int(512),
// 					},
// 					Headers: pulumi.StringArray{
// 						pulumi.String("Content-type"),
// 					},
// 				},
// 				Response: &apimanagement.HttpMessageDiagnosticArgs{
// 					Body: &apimanagement.BodyDiagnosticSettingsArgs{
// 						Bytes: pulumi.Int(512),
// 					},
// 					Headers: pulumi.StringArray{
// 						pulumi.String("Content-type"),
// 					},
// 				},
// 			},
// 			DiagnosticId: pulumi.String("applicationinsights"),
// 			Frontend: &apimanagement.PipelineDiagnosticSettingsArgs{
// 				Request: &apimanagement.HttpMessageDiagnosticArgs{
// 					Body: &apimanagement.BodyDiagnosticSettingsArgs{
// 						Bytes: pulumi.Int(512),
// 					},
// 					Headers: pulumi.StringArray{
// 						pulumi.String("Content-type"),
// 					},
// 				},
// 				Response: &apimanagement.HttpMessageDiagnosticArgs{
// 					Body: &apimanagement.BodyDiagnosticSettingsArgs{
// 						Bytes: pulumi.Int(512),
// 					},
// 					Headers: pulumi.StringArray{
// 						pulumi.String("Content-type"),
// 					},
// 				},
// 			},
// 			LoggerId:          pulumi.String("/loggers/applicationinsights"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			Sampling: &apimanagement.SamplingSettingsArgs{
// 				Percentage:   pulumi.Float64(50),
// 				SamplingType: pulumi.String("fixed"),
// 			},
// 			ServiceName: pulumi.String("apimService1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ApiDiagnostic struct {
	pulumi.CustomResourceState

	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog pulumi.StringPtrOutput `pulumi:"alwaysLog"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"backend"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"frontend"`
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol pulumi.StringPtrOutput `pulumi:"httpCorrelationProtocol"`
	// Log the ClientIP. Default is false.
	LogClientIp pulumi.BoolPtrOutput `pulumi:"logClientIp"`
	// Resource Id of a target logger.
	LoggerId pulumi.StringOutput `pulumi:"loggerId"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sampling settings for Diagnostic.
	Sampling SamplingSettingsResponsePtrOutput `pulumi:"sampling"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity pulumi.StringPtrOutput `pulumi:"verbosity"`
}

// NewApiDiagnostic registers a new resource with the given unique name, arguments, and options.
func NewApiDiagnostic(ctx *pulumi.Context,
	name string, args *ApiDiagnosticArgs, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.DiagnosticId == nil {
		return nil, errors.New("missing required argument 'DiagnosticId'")
	}
	if args == nil || args.LoggerId == nil {
		return nil, errors.New("missing required argument 'LoggerId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ApiDiagnosticArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:ApiDiagnostic"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiDiagnostic
	err := ctx.RegisterResource("azurerm:apimanagement/latest:ApiDiagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiDiagnostic gets an existing ApiDiagnostic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDiagnosticState, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	var resource ApiDiagnostic
	err := ctx.ReadResource("azurerm:apimanagement/latest:ApiDiagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiDiagnostic resources.
type apiDiagnosticState struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog *string `pulumi:"alwaysLog"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend *PipelineDiagnosticSettingsResponse `pulumi:"backend"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend *PipelineDiagnosticSettingsResponse `pulumi:"frontend"`
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// Log the ClientIP. Default is false.
	LogClientIp *bool `pulumi:"logClientIp"`
	// Resource Id of a target logger.
	LoggerId *string `pulumi:"loggerId"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Sampling settings for Diagnostic.
	Sampling *SamplingSettingsResponse `pulumi:"sampling"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity *string `pulumi:"verbosity"`
}

type ApiDiagnosticState struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog pulumi.StringPtrInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend PipelineDiagnosticSettingsResponsePtrInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend PipelineDiagnosticSettingsResponsePtrInput
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol pulumi.StringPtrInput
	// Log the ClientIP. Default is false.
	LogClientIp pulumi.BoolPtrInput
	// Resource Id of a target logger.
	LoggerId pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Sampling settings for Diagnostic.
	Sampling SamplingSettingsResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity pulumi.StringPtrInput
}

func (ApiDiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticState)(nil)).Elem()
}

type apiDiagnosticArgs struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog *string `pulumi:"alwaysLog"`
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend *PipelineDiagnosticSettings `pulumi:"backend"`
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId string `pulumi:"diagnosticId"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend *PipelineDiagnosticSettings `pulumi:"frontend"`
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// Log the ClientIP. Default is false.
	LogClientIp *bool `pulumi:"logClientIp"`
	// Resource Id of a target logger.
	LoggerId string `pulumi:"loggerId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sampling settings for Diagnostic.
	Sampling *SamplingSettings `pulumi:"sampling"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity *string `pulumi:"verbosity"`
}

// The set of arguments for constructing a ApiDiagnostic resource.
type ApiDiagnosticArgs struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog pulumi.StringPtrInput
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend PipelineDiagnosticSettingsPtrInput
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId pulumi.StringInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend PipelineDiagnosticSettingsPtrInput
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol pulumi.StringPtrInput
	// Log the ClientIP. Default is false.
	LogClientIp pulumi.BoolPtrInput
	// Resource Id of a target logger.
	LoggerId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Sampling settings for Diagnostic.
	Sampling SamplingSettingsPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity pulumi.StringPtrInput
}

func (ApiDiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticArgs)(nil)).Elem()
}
