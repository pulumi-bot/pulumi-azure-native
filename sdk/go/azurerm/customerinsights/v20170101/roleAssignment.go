// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Role Assignment resource format.
//
// ## Example Usage
// ### RoleAssignments_CreateOrUpdate
//
// ```go
// package main
//
// import (
// 	customerinsights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/customerinsights/v20170101"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := customerinsights.NewRoleAssignment(ctx, "roleAssignment", &customerinsights.RoleAssignmentArgs{
// 			AssignmentName: pulumi.String("assignmentName8976"),
// 			HubName:        pulumi.String("sdkTestHub"),
// 			Principals: customerinsights.AssignmentPrincipalArray{
// 				&customerinsights.AssignmentPrincipalArgs{
// 					PrincipalId:   pulumi.String("4c54c38ffa9b416ba5a6d6c8a20cbe7e"),
// 					PrincipalType: pulumi.String("User"),
// 				},
// 				&customerinsights.AssignmentPrincipalArgs{
// 					PrincipalId:   pulumi.String("93061d15a5054f2b9948ae25724cf9d5"),
// 					PrincipalType: pulumi.String("User"),
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("TestHubRG"),
// 			Role:              pulumi.String("Admin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type RoleAssignment struct {
	pulumi.CustomResourceState

	// The name of the metadata object.
	AssignmentName pulumi.StringOutput `pulumi:"assignmentName"`
	// Widget types set for the assignment.
	ConflationPolicies ResourceSetDescriptionResponsePtrOutput `pulumi:"conflationPolicies"`
	// Connectors set for the assignment.
	Connectors ResourceSetDescriptionResponsePtrOutput `pulumi:"connectors"`
	// Localized description for the metadata.
	Description pulumi.StringMapOutput `pulumi:"description"`
	// Localized display names for the metadata.
	DisplayName pulumi.StringMapOutput `pulumi:"displayName"`
	// Interactions set for the assignment.
	Interactions ResourceSetDescriptionResponsePtrOutput `pulumi:"interactions"`
	// Kpis set for the assignment.
	Kpis ResourceSetDescriptionResponsePtrOutput `pulumi:"kpis"`
	// Links set for the assignment.
	Links ResourceSetDescriptionResponsePtrOutput `pulumi:"links"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The principals being assigned to.
	Principals AssignmentPrincipalResponseArrayOutput `pulumi:"principals"`
	// Profiles set for the assignment.
	Profiles ResourceSetDescriptionResponsePtrOutput `pulumi:"profiles"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The Role assignments set for the relationship links.
	RelationshipLinks ResourceSetDescriptionResponsePtrOutput `pulumi:"relationshipLinks"`
	// The Role assignments set for the relationships.
	Relationships ResourceSetDescriptionResponsePtrOutput `pulumi:"relationships"`
	// Type of roles.
	Role pulumi.StringOutput `pulumi:"role"`
	// The Role assignments set for the assignment.
	RoleAssignments ResourceSetDescriptionResponsePtrOutput `pulumi:"roleAssignments"`
	// Sas Policies set for the assignment.
	SasPolicies ResourceSetDescriptionResponsePtrOutput `pulumi:"sasPolicies"`
	// The Role assignments set for the assignment.
	Segments ResourceSetDescriptionResponsePtrOutput `pulumi:"segments"`
	// The hub name.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Views set for the assignment.
	Views ResourceSetDescriptionResponsePtrOutput `pulumi:"views"`
	// Widget types set for the assignment.
	WidgetTypes ResourceSetDescriptionResponsePtrOutput `pulumi:"widgetTypes"`
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil || args.AssignmentName == nil {
		return nil, errors.New("missing required argument 'AssignmentName'")
	}
	if args == nil || args.HubName == nil {
		return nil, errors.New("missing required argument 'HubName'")
	}
	if args == nil || args.Principals == nil {
		return nil, errors.New("missing required argument 'Principals'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &RoleAssignmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:customerinsights/latest:RoleAssignment"),
		},
		{
			Type: pulumi.String("azurerm:customerinsights/v20170426:RoleAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleAssignment
	err := ctx.RegisterResource("azurerm:customerinsights/v20170101:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAssignment gets an existing RoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azurerm:customerinsights/v20170101:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssignment resources.
type roleAssignmentState struct {
	// The name of the metadata object.
	AssignmentName *string `pulumi:"assignmentName"`
	// Widget types set for the assignment.
	ConflationPolicies *ResourceSetDescriptionResponse `pulumi:"conflationPolicies"`
	// Connectors set for the assignment.
	Connectors *ResourceSetDescriptionResponse `pulumi:"connectors"`
	// Localized description for the metadata.
	Description map[string]string `pulumi:"description"`
	// Localized display names for the metadata.
	DisplayName map[string]string `pulumi:"displayName"`
	// Interactions set for the assignment.
	Interactions *ResourceSetDescriptionResponse `pulumi:"interactions"`
	// Kpis set for the assignment.
	Kpis *ResourceSetDescriptionResponse `pulumi:"kpis"`
	// Links set for the assignment.
	Links *ResourceSetDescriptionResponse `pulumi:"links"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The principals being assigned to.
	Principals []AssignmentPrincipalResponse `pulumi:"principals"`
	// Profiles set for the assignment.
	Profiles *ResourceSetDescriptionResponse `pulumi:"profiles"`
	// Provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The Role assignments set for the relationship links.
	RelationshipLinks *ResourceSetDescriptionResponse `pulumi:"relationshipLinks"`
	// The Role assignments set for the relationships.
	Relationships *ResourceSetDescriptionResponse `pulumi:"relationships"`
	// Type of roles.
	Role *string `pulumi:"role"`
	// The Role assignments set for the assignment.
	RoleAssignments *ResourceSetDescriptionResponse `pulumi:"roleAssignments"`
	// Sas Policies set for the assignment.
	SasPolicies *ResourceSetDescriptionResponse `pulumi:"sasPolicies"`
	// The Role assignments set for the assignment.
	Segments *ResourceSetDescriptionResponse `pulumi:"segments"`
	// The hub name.
	TenantId *string `pulumi:"tenantId"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Views set for the assignment.
	Views *ResourceSetDescriptionResponse `pulumi:"views"`
	// Widget types set for the assignment.
	WidgetTypes *ResourceSetDescriptionResponse `pulumi:"widgetTypes"`
}

type RoleAssignmentState struct {
	// The name of the metadata object.
	AssignmentName pulumi.StringPtrInput
	// Widget types set for the assignment.
	ConflationPolicies ResourceSetDescriptionResponsePtrInput
	// Connectors set for the assignment.
	Connectors ResourceSetDescriptionResponsePtrInput
	// Localized description for the metadata.
	Description pulumi.StringMapInput
	// Localized display names for the metadata.
	DisplayName pulumi.StringMapInput
	// Interactions set for the assignment.
	Interactions ResourceSetDescriptionResponsePtrInput
	// Kpis set for the assignment.
	Kpis ResourceSetDescriptionResponsePtrInput
	// Links set for the assignment.
	Links ResourceSetDescriptionResponsePtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The principals being assigned to.
	Principals AssignmentPrincipalResponseArrayInput
	// Profiles set for the assignment.
	Profiles ResourceSetDescriptionResponsePtrInput
	// Provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The Role assignments set for the relationship links.
	RelationshipLinks ResourceSetDescriptionResponsePtrInput
	// The Role assignments set for the relationships.
	Relationships ResourceSetDescriptionResponsePtrInput
	// Type of roles.
	Role pulumi.StringPtrInput
	// The Role assignments set for the assignment.
	RoleAssignments ResourceSetDescriptionResponsePtrInput
	// Sas Policies set for the assignment.
	SasPolicies ResourceSetDescriptionResponsePtrInput
	// The Role assignments set for the assignment.
	Segments ResourceSetDescriptionResponsePtrInput
	// The hub name.
	TenantId pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Views set for the assignment.
	Views ResourceSetDescriptionResponsePtrInput
	// Widget types set for the assignment.
	WidgetTypes ResourceSetDescriptionResponsePtrInput
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	// The assignment name
	AssignmentName string `pulumi:"assignmentName"`
	// Widget types set for the assignment.
	ConflationPolicies *ResourceSetDescription `pulumi:"conflationPolicies"`
	// Connectors set for the assignment.
	Connectors *ResourceSetDescription `pulumi:"connectors"`
	// Localized description for the metadata.
	Description map[string]string `pulumi:"description"`
	// Localized display names for the metadata.
	DisplayName map[string]string `pulumi:"displayName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// Interactions set for the assignment.
	Interactions *ResourceSetDescription `pulumi:"interactions"`
	// Kpis set for the assignment.
	Kpis *ResourceSetDescription `pulumi:"kpis"`
	// Links set for the assignment.
	Links *ResourceSetDescription `pulumi:"links"`
	// The principals being assigned to.
	Principals []AssignmentPrincipal `pulumi:"principals"`
	// Profiles set for the assignment.
	Profiles *ResourceSetDescription `pulumi:"profiles"`
	// The Role assignments set for the relationship links.
	RelationshipLinks *ResourceSetDescription `pulumi:"relationshipLinks"`
	// The Role assignments set for the relationships.
	Relationships *ResourceSetDescription `pulumi:"relationships"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Type of roles.
	Role string `pulumi:"role"`
	// The Role assignments set for the assignment.
	RoleAssignments *ResourceSetDescription `pulumi:"roleAssignments"`
	// Sas Policies set for the assignment.
	SasPolicies *ResourceSetDescription `pulumi:"sasPolicies"`
	// The Role assignments set for the assignment.
	Segments *ResourceSetDescription `pulumi:"segments"`
	// Views set for the assignment.
	Views *ResourceSetDescription `pulumi:"views"`
	// Widget types set for the assignment.
	WidgetTypes *ResourceSetDescription `pulumi:"widgetTypes"`
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// The assignment name
	AssignmentName pulumi.StringInput
	// Widget types set for the assignment.
	ConflationPolicies ResourceSetDescriptionPtrInput
	// Connectors set for the assignment.
	Connectors ResourceSetDescriptionPtrInput
	// Localized description for the metadata.
	Description pulumi.StringMapInput
	// Localized display names for the metadata.
	DisplayName pulumi.StringMapInput
	// The name of the hub.
	HubName pulumi.StringInput
	// Interactions set for the assignment.
	Interactions ResourceSetDescriptionPtrInput
	// Kpis set for the assignment.
	Kpis ResourceSetDescriptionPtrInput
	// Links set for the assignment.
	Links ResourceSetDescriptionPtrInput
	// The principals being assigned to.
	Principals AssignmentPrincipalArrayInput
	// Profiles set for the assignment.
	Profiles ResourceSetDescriptionPtrInput
	// The Role assignments set for the relationship links.
	RelationshipLinks ResourceSetDescriptionPtrInput
	// The Role assignments set for the relationships.
	Relationships ResourceSetDescriptionPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Type of roles.
	Role pulumi.StringInput
	// The Role assignments set for the assignment.
	RoleAssignments ResourceSetDescriptionPtrInput
	// Sas Policies set for the assignment.
	SasPolicies ResourceSetDescriptionPtrInput
	// The Role assignments set for the assignment.
	Segments ResourceSetDescriptionPtrInput
	// Views set for the assignment.
	Views ResourceSetDescriptionPtrInput
	// Widget types set for the assignment.
	WidgetTypes ResourceSetDescriptionPtrInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}
