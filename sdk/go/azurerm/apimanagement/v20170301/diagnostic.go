// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Diagnostic details.
//
// ## Example Usage
// ### ApiManagementCreateDiagnostic
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20170301"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewDiagnostic(ctx, "diagnostic", &apimanagement.DiagnosticArgs{
// 			DiagnosticId:      pulumi.String("default"),
// 			Enabled:           pulumi.Bool(true),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Diagnostic struct {
	pulumi.CustomResourceState

	// Indicates whether a diagnostic should receive data or not.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDiagnostic registers a new resource with the given unique name, arguments, and options.
func NewDiagnostic(ctx *pulumi.Context,
	name string, args *DiagnosticArgs, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	if args == nil || args.DiagnosticId == nil {
		return nil, errors.New("missing required argument 'DiagnosticId'")
	}
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &DiagnosticArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:Diagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:Diagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:Diagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:Diagnostic"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:Diagnostic"),
		},
	})
	opts = append(opts, aliases)
	var resource Diagnostic
	err := ctx.RegisterResource("azurerm:apimanagement/v20170301:Diagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiagnostic gets an existing Diagnostic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticState, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	var resource Diagnostic
	err := ctx.ReadResource("azurerm:apimanagement/v20170301:Diagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Diagnostic resources.
type diagnosticState struct {
	// Indicates whether a diagnostic should receive data or not.
	Enabled *bool `pulumi:"enabled"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type DiagnosticState struct {
	// Indicates whether a diagnostic should receive data or not.
	Enabled pulumi.BoolPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (DiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticState)(nil)).Elem()
}

type diagnosticArgs struct {
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId string `pulumi:"diagnosticId"`
	// Indicates whether a diagnostic should receive data or not.
	Enabled bool `pulumi:"enabled"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Diagnostic resource.
type DiagnosticArgs struct {
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId pulumi.StringInput
	// Indicates whether a diagnostic should receive data or not.
	Enabled pulumi.BoolInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (DiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticArgs)(nil)).Elem()
}
