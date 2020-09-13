// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The policy set definition.
//
// ## Example Usage
// ### Create or update a policy set definition at management group level
//
// ```go
// package main
//
// import (
// 	management "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/management/v20180501"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := management.NewPolicySetDefinitionAtManagementGroup(ctx, "policySetDefinitionAtManagementGroup", &management.PolicySetDefinitionAtManagementGroupArgs{
// 			Description:       pulumi.String("Policies to enforce low cost storage SKUs"),
// 			DisplayName:       pulumi.String("Cost Management"),
// 			ManagementGroupId: pulumi.String("MyManagementGroup"),
// 			Metadata: pulumi.StringMap{
// 				"category": pulumi.String("Cost Management"),
// 			},
// 			PolicyDefinitions: management.PolicyDefinitionReferenceArray{
// 				&management.PolicyDefinitionReferenceArgs{
// 					Parameters: pulumi.StringArrayMapMap{
// 						"listOfAllowedSKUs": pulumi.StringArrayMap{
// 							"value": pulumi.StringArray{
// 								pulumi.String("Standard_GRS"),
// 								pulumi.String("Standard_LRS"),
// 							},
// 						},
// 					},
// 					PolicyDefinitionId: pulumi.String("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
// 				},
// 				&management.PolicyDefinitionReferenceArgs{
// 					Parameters: pulumi.StringMapMap{
// 						"prefix": pulumi.StringMap{
// 							"value": pulumi.String("DeptA"),
// 						},
// 						"suffix": pulumi.StringMap{
// 							"value": pulumi.String("-LC"),
// 						},
// 					},
// 					PolicyDefinitionId: pulumi.String("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
// 				},
// 			},
// 			PolicySetDefinitionName: pulumi.String("CostManagement"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type PolicySetDefinitionAtManagementGroup struct {
	pulumi.CustomResourceState

	// The policy set definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The policy set definition metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// The name of the policy set definition.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy set definition parameters that can be used in policy definition references.
	Parameters pulumi.MapOutput `pulumi:"parameters"`
	// An array of policy definition references.
	PolicyDefinitions PolicyDefinitionReferenceResponseArrayOutput `pulumi:"policyDefinitions"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
	PolicyType pulumi.StringPtrOutput `pulumi:"policyType"`
	// The type of the resource (Microsoft.Authorization/policySetDefinitions).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicySetDefinitionAtManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewPolicySetDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, args *PolicySetDefinitionAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*PolicySetDefinitionAtManagementGroup, error) {
	if args == nil || args.ManagementGroupId == nil {
		return nil, errors.New("missing required argument 'ManagementGroupId'")
	}
	if args == nil || args.PolicyDefinitions == nil {
		return nil, errors.New("missing required argument 'PolicyDefinitions'")
	}
	if args == nil || args.PolicySetDefinitionName == nil {
		return nil, errors.New("missing required argument 'PolicySetDefinitionName'")
	}
	if args == nil {
		args = &PolicySetDefinitionAtManagementGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:management/latest:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20170601preview:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20180301:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20190101:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20190601:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20190901:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20200301:PolicySetDefinitionAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicySetDefinitionAtManagementGroup
	err := ctx.RegisterResource("azurerm:management/v20180501:PolicySetDefinitionAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicySetDefinitionAtManagementGroup gets an existing PolicySetDefinitionAtManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicySetDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicySetDefinitionAtManagementGroupState, opts ...pulumi.ResourceOption) (*PolicySetDefinitionAtManagementGroup, error) {
	var resource PolicySetDefinitionAtManagementGroup
	err := ctx.ReadResource("azurerm:management/v20180501:PolicySetDefinitionAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicySetDefinitionAtManagementGroup resources.
type policySetDefinitionAtManagementGroupState struct {
	// The policy set definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName *string `pulumi:"displayName"`
	// The policy set definition metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The name of the policy set definition.
	Name *string `pulumi:"name"`
	// The policy set definition parameters that can be used in policy definition references.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// An array of policy definition references.
	PolicyDefinitions []PolicyDefinitionReferenceResponse `pulumi:"policyDefinitions"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
	PolicyType *string `pulumi:"policyType"`
	// The type of the resource (Microsoft.Authorization/policySetDefinitions).
	Type *string `pulumi:"type"`
}

type PolicySetDefinitionAtManagementGroupState struct {
	// The policy set definition description.
	Description pulumi.StringPtrInput
	// The display name of the policy set definition.
	DisplayName pulumi.StringPtrInput
	// The policy set definition metadata.
	Metadata pulumi.MapInput
	// The name of the policy set definition.
	Name pulumi.StringPtrInput
	// The policy set definition parameters that can be used in policy definition references.
	Parameters pulumi.MapInput
	// An array of policy definition references.
	PolicyDefinitions PolicyDefinitionReferenceResponseArrayInput
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
	PolicyType pulumi.StringPtrInput
	// The type of the resource (Microsoft.Authorization/policySetDefinitions).
	Type pulumi.StringPtrInput
}

func (PolicySetDefinitionAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionAtManagementGroupState)(nil)).Elem()
}

type policySetDefinitionAtManagementGroupArgs struct {
	// The policy set definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the management group.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The policy set definition metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The policy set definition parameters that can be used in policy definition references.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// An array of policy definition references.
	PolicyDefinitions []PolicyDefinitionReference `pulumi:"policyDefinitions"`
	// The name of the policy set definition to create.
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
	PolicyType *string `pulumi:"policyType"`
}

// The set of arguments for constructing a PolicySetDefinitionAtManagementGroup resource.
type PolicySetDefinitionAtManagementGroupArgs struct {
	// The policy set definition description.
	Description pulumi.StringPtrInput
	// The display name of the policy set definition.
	DisplayName pulumi.StringPtrInput
	// The ID of the management group.
	ManagementGroupId pulumi.StringInput
	// The policy set definition metadata.
	Metadata pulumi.MapInput
	// The policy set definition parameters that can be used in policy definition references.
	Parameters pulumi.MapInput
	// An array of policy definition references.
	PolicyDefinitions PolicyDefinitionReferenceArrayInput
	// The name of the policy set definition to create.
	PolicySetDefinitionName pulumi.StringInput
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
	PolicyType pulumi.StringPtrInput
}

func (PolicySetDefinitionAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionAtManagementGroupArgs)(nil)).Elem()
}
