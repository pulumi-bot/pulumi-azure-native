// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Example Usage
// ### Create JIT network access policy
//
// ```go
// package main
//
// import (
// 	security "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/security/v20200101"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := security.NewJitNetworkAccessPolicy(ctx, "jitNetworkAccessPolicy", &security.JitNetworkAccessPolicyArgs{
// 			AscLocation:                pulumi.String("westeurope"),
// 			JitNetworkAccessPolicyName: pulumi.String("default"),
// 			Kind:                       pulumi.String("Basic"),
// 			Requests: security.JitNetworkAccessRequestArray{
// 				&security.JitNetworkAccessRequestArgs{
// 					Requestor:    pulumi.String("barbara@contoso.com"),
// 					StartTimeUtc: pulumi.String("2018-05-17T08:06:45.5691611Z"),
// 					VirtualMachines: security.JitNetworkAccessRequestVirtualMachineArray{
// 						&security.JitNetworkAccessRequestVirtualMachineArgs{
// 							Id: pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
// 							Ports: security.JitNetworkAccessRequestPortArray{
// 								&security.JitNetworkAccessRequestPortArgs{
// 									AllowedSourceAddressPrefix: pulumi.String("192.127.0.2"),
// 									EndTimeUtc:                 pulumi.String("2018-05-17T09:06:45.5691611Z"),
// 									Number:                     pulumi.Int(3389),
// 									Status:                     pulumi.String("Initiated"),
// 									StatusReason:               pulumi.String("UserRequested"),
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("myRg1"),
// 			VirtualMachines: security.JitNetworkAccessPolicyVirtualMachineArray{
// 				&security.JitNetworkAccessPolicyVirtualMachineArgs{
// 					Id: pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
// 					Ports: security.JitNetworkAccessPortRuleArray{
// 						&security.JitNetworkAccessPortRuleArgs{
// 							AllowedSourceAddressPrefix: pulumi.String("*"),
// 							MaxRequestAccessDuration:   pulumi.String("PT3H"),
// 							Number:                     pulumi.Int(22),
// 							Protocol:                   pulumi.String("*"),
// 						},
// 						&security.JitNetworkAccessPortRuleArgs{
// 							AllowedSourceAddressPrefix: pulumi.String("*"),
// 							MaxRequestAccessDuration:   pulumi.String("PT3H"),
// 							Number:                     pulumi.Int(3389),
// 							Protocol:                   pulumi.String("*"),
// 						},
// 					},
// 				},
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
type JitNetworkAccessPolicy struct {
	pulumi.CustomResourceState

	// Kind of the resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Location where the resource is stored
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state of the Just-in-Time policy.
	ProvisioningState pulumi.StringOutput                        `pulumi:"provisioningState"`
	Requests          JitNetworkAccessRequestResponseArrayOutput `pulumi:"requests"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines JitNetworkAccessPolicyVirtualMachineResponseArrayOutput `pulumi:"virtualMachines"`
}

// NewJitNetworkAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, args *JitNetworkAccessPolicyArgs, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	if args == nil || args.AscLocation == nil {
		return nil, errors.New("missing required argument 'AscLocation'")
	}
	if args == nil || args.JitNetworkAccessPolicyName == nil {
		return nil, errors.New("missing required argument 'JitNetworkAccessPolicyName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualMachines == nil {
		return nil, errors.New("missing required argument 'VirtualMachines'")
	}
	if args == nil {
		args = &JitNetworkAccessPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:security/latest:JitNetworkAccessPolicy"),
		},
		{
			Type: pulumi.String("azurerm:security/v20150601preview:JitNetworkAccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource JitNetworkAccessPolicy
	err := ctx.RegisterResource("azurerm:security/v20200101:JitNetworkAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJitNetworkAccessPolicy gets an existing JitNetworkAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JitNetworkAccessPolicyState, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	var resource JitNetworkAccessPolicy
	err := ctx.ReadResource("azurerm:security/v20200101:JitNetworkAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JitNetworkAccessPolicy resources.
type jitNetworkAccessPolicyState struct {
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Gets the provisioning state of the Just-in-Time policy.
	ProvisioningState *string                           `pulumi:"provisioningState"`
	Requests          []JitNetworkAccessRequestResponse `pulumi:"requests"`
	// Resource type
	Type *string `pulumi:"type"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines []JitNetworkAccessPolicyVirtualMachineResponse `pulumi:"virtualMachines"`
}

type JitNetworkAccessPolicyState struct {
	// Kind of the resource
	Kind pulumi.StringPtrInput
	// Location where the resource is stored
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Gets the provisioning state of the Just-in-Time policy.
	ProvisioningState pulumi.StringPtrInput
	Requests          JitNetworkAccessRequestResponseArrayInput
	// Resource type
	Type pulumi.StringPtrInput
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines JitNetworkAccessPolicyVirtualMachineResponseArrayInput
}

func (JitNetworkAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyState)(nil)).Elem()
}

type jitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation string `pulumi:"ascLocation"`
	// Name of a Just-in-Time access configuration policy.
	JitNetworkAccessPolicyName string `pulumi:"jitNetworkAccessPolicyName"`
	// Kind of the resource
	Kind     *string                   `pulumi:"kind"`
	Requests []JitNetworkAccessRequest `pulumi:"requests"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines []JitNetworkAccessPolicyVirtualMachine `pulumi:"virtualMachines"`
}

// The set of arguments for constructing a JitNetworkAccessPolicy resource.
type JitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation pulumi.StringInput
	// Name of a Just-in-Time access configuration policy.
	JitNetworkAccessPolicyName pulumi.StringInput
	// Kind of the resource
	Kind     pulumi.StringPtrInput
	Requests JitNetworkAccessRequestArrayInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines JitNetworkAccessPolicyVirtualMachineArrayInput
}

func (JitNetworkAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyArgs)(nil)).Elem()
}
