// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200215

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing a cluster principal assignment.
type ClusterPrincipalAssignment struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The principal name
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// Principal type.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Cluster principal role.
	Role pulumi.StringOutput `pulumi:"role"`
	// The tenant id of the principal
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The tenant name of the principal
	TenantName pulumi.StringOutput `pulumi:"tenantName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewClusterPrincipalAssignment registers a new resource with the given unique name, arguments, and options.
func NewClusterPrincipalAssignment(ctx *pulumi.Context,
	name string, args *ClusterPrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*ClusterPrincipalAssignment, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.PrincipalAssignmentName == nil {
		return nil, errors.New("missing required argument 'PrincipalAssignmentName'")
	}
	if args == nil || args.PrincipalId == nil {
		return nil, errors.New("missing required argument 'PrincipalId'")
	}
	if args == nil || args.PrincipalType == nil {
		return nil, errors.New("missing required argument 'PrincipalType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &ClusterPrincipalAssignmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:kusto/latest:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20191109:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20200614:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20200918:ClusterPrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource ClusterPrincipalAssignment
	err := ctx.RegisterResource("azure-nextgen:kusto/v20200215:ClusterPrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterPrincipalAssignment gets an existing ClusterPrincipalAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterPrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterPrincipalAssignmentState, opts ...pulumi.ResourceOption) (*ClusterPrincipalAssignment, error) {
	var resource ClusterPrincipalAssignment
	err := ctx.ReadResource("azure-nextgen:kusto/v20200215:ClusterPrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterPrincipalAssignment resources.
type clusterPrincipalAssignmentState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId *string `pulumi:"principalId"`
	// The principal name
	PrincipalName *string `pulumi:"principalName"`
	// Principal type.
	PrincipalType *string `pulumi:"principalType"`
	// The provisioned state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Cluster principal role.
	Role *string `pulumi:"role"`
	// The tenant id of the principal
	TenantId *string `pulumi:"tenantId"`
	// The tenant name of the principal
	TenantName *string `pulumi:"tenantName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type ClusterPrincipalAssignmentState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringPtrInput
	// The principal name
	PrincipalName pulumi.StringPtrInput
	// Principal type.
	PrincipalType pulumi.StringPtrInput
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Cluster principal role.
	Role pulumi.StringPtrInput
	// The tenant id of the principal
	TenantId pulumi.StringPtrInput
	// The tenant name of the principal
	TenantName pulumi.StringPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (ClusterPrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPrincipalAssignmentState)(nil)).Elem()
}

type clusterPrincipalAssignmentArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId string `pulumi:"principalId"`
	// Principal type.
	PrincipalType string `pulumi:"principalType"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cluster principal role.
	Role string `pulumi:"role"`
	// The tenant id of the principal
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ClusterPrincipalAssignment resource.
type ClusterPrincipalAssignmentArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName pulumi.StringInput
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringInput
	// Principal type.
	PrincipalType pulumi.StringInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// Cluster principal role.
	Role pulumi.StringInput
	// The tenant id of the principal
	TenantId pulumi.StringPtrInput
}

func (ClusterPrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPrincipalAssignmentArgs)(nil)).Elem()
}

type ClusterPrincipalAssignmentInput interface {
	pulumi.Input

	ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput
	ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput
}

func (ClusterPrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPrincipalAssignment)(nil)).Elem()
}

func (i ClusterPrincipalAssignment) ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput {
	return i.ToClusterPrincipalAssignmentOutputWithContext(context.Background())
}

func (i ClusterPrincipalAssignment) ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPrincipalAssignmentOutput)
}

type ClusterPrincipalAssignmentOutput struct {
	*pulumi.OutputState
}

func (ClusterPrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPrincipalAssignmentOutput)(nil)).Elem()
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput {
	return o
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterPrincipalAssignmentOutput{})
}
