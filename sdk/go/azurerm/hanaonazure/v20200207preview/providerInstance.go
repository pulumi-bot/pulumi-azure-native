// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200207preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A provider instance associated with a SAP monitor.
//
// ## Example Usage
// ### Create a SAP Monitor
//
// ```go
// package main
//
// import (
// 	hanaonazure "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/hanaonazure/v20200207preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := hanaonazure.NewProviderInstance(ctx, "providerInstance", &hanaonazure.ProviderInstanceArgs{
// 			Metadata:             pulumi.String("{\"key\":\"value\"}"),
// 			Properties:           pulumi.String("{\"hostname\":\"10.0.0.10\",\"dbName\":\"SYSTEMDB\",\"sqlPort\":30015,\"dbUsername\":\"SYSTEM\",\"dbPassword\":\"PASSWORD\"}"),
// 			ProviderInstanceName: pulumi.String("myProviderInstance"),
// 			ResourceGroupName:    pulumi.String("myResourceGroup"),
// 			SapMonitorName:       pulumi.String("mySapMonitor"),
// 			Type:                 pulumi.String("hana"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ProviderInstance struct {
	pulumi.CustomResourceState

	// A JSON string containing metadata of the provider instance.
	Metadata pulumi.StringPtrOutput `pulumi:"metadata"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A JSON string containing the properties of the provider instance.
	Properties pulumi.StringOutput `pulumi:"properties"`
	// State of provisioning of the provider instance
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProviderInstance registers a new resource with the given unique name, arguments, and options.
func NewProviderInstance(ctx *pulumi.Context,
	name string, args *ProviderInstanceArgs, opts ...pulumi.ResourceOption) (*ProviderInstance, error) {
	if args == nil || args.ProviderInstanceName == nil {
		return nil, errors.New("missing required argument 'ProviderInstanceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SapMonitorName == nil {
		return nil, errors.New("missing required argument 'SapMonitorName'")
	}
	if args == nil {
		args = &ProviderInstanceArgs{}
	}
	var resource ProviderInstance
	err := ctx.RegisterResource("azurerm:hanaonazure/v20200207preview:ProviderInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderInstance gets an existing ProviderInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderInstanceState, opts ...pulumi.ResourceOption) (*ProviderInstance, error) {
	var resource ProviderInstance
	err := ctx.ReadResource("azurerm:hanaonazure/v20200207preview:ProviderInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderInstance resources.
type providerInstanceState struct {
	// A JSON string containing metadata of the provider instance.
	Metadata *string `pulumi:"metadata"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// A JSON string containing the properties of the provider instance.
	Properties *string `pulumi:"properties"`
	// State of provisioning of the provider instance
	ProvisioningState *string `pulumi:"provisioningState"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ProviderInstanceState struct {
	// A JSON string containing metadata of the provider instance.
	Metadata pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// A JSON string containing the properties of the provider instance.
	Properties pulumi.StringPtrInput
	// State of provisioning of the provider instance
	ProvisioningState pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ProviderInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerInstanceState)(nil)).Elem()
}

type providerInstanceArgs struct {
	// A JSON string containing metadata of the provider instance.
	Metadata *string `pulumi:"metadata"`
	// A JSON string containing the properties of the provider instance.
	Properties *string `pulumi:"properties"`
	// Name of the provider instance.
	ProviderInstanceName string `pulumi:"providerInstanceName"`
	// Name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SAP monitor resource.
	SapMonitorName string `pulumi:"sapMonitorName"`
	// The type of provider instance.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ProviderInstance resource.
type ProviderInstanceArgs struct {
	// A JSON string containing metadata of the provider instance.
	Metadata pulumi.StringPtrInput
	// A JSON string containing the properties of the provider instance.
	Properties pulumi.StringPtrInput
	// Name of the provider instance.
	ProviderInstanceName pulumi.StringInput
	// Name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Name of the SAP monitor resource.
	SapMonitorName pulumi.StringInput
	// The type of provider instance.
	Type pulumi.StringPtrInput
}

func (ProviderInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerInstanceArgs)(nil)).Elem()
}
