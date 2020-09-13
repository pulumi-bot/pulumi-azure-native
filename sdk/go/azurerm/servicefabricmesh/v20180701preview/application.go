// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This type describes an application resource.
//
// ## Example Usage
// ### ApplicationCreateOrUpdate
//
// ```go
// package main
//
// import (
// 	servicefabricmesh "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/servicefabricmesh/v20180701preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicefabricmesh.NewApplication(ctx, "application", &servicefabricmesh.ApplicationArgs{
// 			ApplicationName:   pulumi.String("helloWorldApp"),
// 			Description:       pulumi.String("SeaBreeze HelloWorld Application!"),
// 			Location:          pulumi.String("EastUS"),
// 			ResourceGroupName: pulumi.String("sbz_demo"),
// 			Services: servicefabricmesh.ServiceResourceDescriptionArray{
// 				&servicefabricmesh.ServiceResourceDescriptionArgs{
// 					Name: pulumi.String("helloWorldService"),
// 				},
// 			},
// 			Tags: nil,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Application struct {
	pulumi.CustomResourceState

	// Internal use.
	DebugParams pulumi.StringPtrOutput `pulumi:"debugParams"`
	// User readable description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics DiagnosticsDescriptionResponsePtrOutput `pulumi:"diagnostics"`
	// Describes the health state of an application resource.
	HealthState pulumi.StringOutput `pulumi:"healthState"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Names of the services in the application.
	ServiceNames pulumi.StringArrayOutput `pulumi:"serviceNames"`
	// describes the services in the application.
	Services ServiceResourceDescriptionResponseArrayOutput `pulumi:"services"`
	// Status of the application resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Gives additional information about the current status of the application deployment.
	StatusDetails pulumi.StringOutput `pulumi:"statusDetails"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
	UnhealthyEvaluation pulumi.StringOutput `pulumi:"unhealthyEvaluation"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil || args.ApplicationName == nil {
		return nil, errors.New("missing required argument 'ApplicationName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:servicefabricmesh/v20180901preview:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azurerm:servicefabricmesh/v20180701preview:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azurerm:servicefabricmesh/v20180701preview:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// Internal use.
	DebugParams *string `pulumi:"debugParams"`
	// User readable description of the application.
	Description *string `pulumi:"description"`
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics *DiagnosticsDescriptionResponse `pulumi:"diagnostics"`
	// Describes the health state of an application resource.
	HealthState *string `pulumi:"healthState"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// State of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Names of the services in the application.
	ServiceNames []string `pulumi:"serviceNames"`
	// describes the services in the application.
	Services []ServiceResourceDescriptionResponse `pulumi:"services"`
	// Status of the application resource.
	Status *string `pulumi:"status"`
	// Gives additional information about the current status of the application deployment.
	StatusDetails *string `pulumi:"statusDetails"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
	UnhealthyEvaluation *string `pulumi:"unhealthyEvaluation"`
}

type ApplicationState struct {
	// Internal use.
	DebugParams pulumi.StringPtrInput
	// User readable description of the application.
	Description pulumi.StringPtrInput
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics DiagnosticsDescriptionResponsePtrInput
	// Describes the health state of an application resource.
	HealthState pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// State of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Names of the services in the application.
	ServiceNames pulumi.StringArrayInput
	// describes the services in the application.
	Services ServiceResourceDescriptionResponseArrayInput
	// Status of the application resource.
	Status pulumi.StringPtrInput
	// Gives additional information about the current status of the application deployment.
	StatusDetails pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
	UnhealthyEvaluation pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The identity of the application.
	ApplicationName string `pulumi:"applicationName"`
	// Internal use.
	DebugParams *string `pulumi:"debugParams"`
	// User readable description of the application.
	Description *string `pulumi:"description"`
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics *DiagnosticsDescription `pulumi:"diagnostics"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// describes the services in the application.
	Services []ServiceResourceDescription `pulumi:"services"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The identity of the application.
	ApplicationName pulumi.StringInput
	// Internal use.
	DebugParams pulumi.StringPtrInput
	// User readable description of the application.
	Description pulumi.StringPtrInput
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics DiagnosticsDescriptionPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Azure resource group name
	ResourceGroupName pulumi.StringInput
	// describes the services in the application.
	Services ServiceResourceDescriptionArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
