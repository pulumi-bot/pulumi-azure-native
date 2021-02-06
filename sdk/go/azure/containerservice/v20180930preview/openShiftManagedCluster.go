// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180930preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// OpenShift Managed cluster.
type OpenShiftManagedCluster struct {
	pulumi.CustomResourceState

	// Configuration of OpenShift cluster VMs.
	AgentPoolProfiles OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput `pulumi:"agentPoolProfiles"`
	// Configures OpenShift authentication.
	AuthProfile OpenShiftManagedClusterAuthProfileResponsePtrOutput `pulumi:"authProfile"`
	// User-specified FQDN for OpenShift API server loadbalancer internal hostname.
	Fqdn pulumi.StringPtrOutput `pulumi:"fqdn"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Configuration for OpenShift master VMs.
	MasterPoolProfile OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput `pulumi:"masterPoolProfile"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for OpenShift networking.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Version of OpenShift specified when creating the cluster.
	OpenShiftVersion pulumi.StringOutput `pulumi:"openShiftVersion"`
	// Define the resource plan as required by ARM for billing purposes
	Plan PurchasePlanResponsePtrOutput `pulumi:"plan"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Optional user-specified FQDN for OpenShift API server.
	PublicHostname pulumi.StringPtrOutput `pulumi:"publicHostname"`
	// Configuration for OpenShift router(s).
	RouterProfiles OpenShiftRouterProfileResponseArrayOutput `pulumi:"routerProfiles"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOpenShiftManagedCluster registers a new resource with the given unique name, arguments, and options.
func NewOpenShiftManagedCluster(ctx *pulumi.Context,
	name string, args *OpenShiftManagedClusterArgs, opts ...pulumi.ResourceOption) (*OpenShiftManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OpenShiftVersion == nil {
		return nil, errors.New("invalid value for required argument 'OpenShiftVersion'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerservice/latest:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190430:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190930preview:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20191027preview:OpenShiftManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource OpenShiftManagedCluster
	err := ctx.RegisterResource("azure-nextgen:containerservice/v20180930preview:OpenShiftManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenShiftManagedCluster gets an existing OpenShiftManagedCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenShiftManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenShiftManagedClusterState, opts ...pulumi.ResourceOption) (*OpenShiftManagedCluster, error) {
	var resource OpenShiftManagedCluster
	err := ctx.ReadResource("azure-nextgen:containerservice/v20180930preview:OpenShiftManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenShiftManagedCluster resources.
type openShiftManagedClusterState struct {
	// Configuration of OpenShift cluster VMs.
	AgentPoolProfiles []OpenShiftManagedClusterAgentPoolProfileResponse `pulumi:"agentPoolProfiles"`
	// Configures OpenShift authentication.
	AuthProfile *OpenShiftManagedClusterAuthProfileResponse `pulumi:"authProfile"`
	// User-specified FQDN for OpenShift API server loadbalancer internal hostname.
	Fqdn *string `pulumi:"fqdn"`
	// Resource location
	Location *string `pulumi:"location"`
	// Configuration for OpenShift master VMs.
	MasterPoolProfile *OpenShiftManagedClusterMasterPoolProfileResponse `pulumi:"masterPoolProfile"`
	// Resource name
	Name *string `pulumi:"name"`
	// Configuration for OpenShift networking.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// Version of OpenShift specified when creating the cluster.
	OpenShiftVersion *string `pulumi:"openShiftVersion"`
	// Define the resource plan as required by ARM for billing purposes
	Plan *PurchasePlanResponse `pulumi:"plan"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Optional user-specified FQDN for OpenShift API server.
	PublicHostname *string `pulumi:"publicHostname"`
	// Configuration for OpenShift router(s).
	RouterProfiles []OpenShiftRouterProfileResponse `pulumi:"routerProfiles"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type OpenShiftManagedClusterState struct {
	// Configuration of OpenShift cluster VMs.
	AgentPoolProfiles OpenShiftManagedClusterAgentPoolProfileResponseArrayInput
	// Configures OpenShift authentication.
	AuthProfile OpenShiftManagedClusterAuthProfileResponsePtrInput
	// User-specified FQDN for OpenShift API server loadbalancer internal hostname.
	Fqdn pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Configuration for OpenShift master VMs.
	MasterPoolProfile OpenShiftManagedClusterMasterPoolProfileResponsePtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Configuration for OpenShift networking.
	NetworkProfile NetworkProfileResponsePtrInput
	// Version of OpenShift specified when creating the cluster.
	OpenShiftVersion pulumi.StringPtrInput
	// Define the resource plan as required by ARM for billing purposes
	Plan PurchasePlanResponsePtrInput
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// Optional user-specified FQDN for OpenShift API server.
	PublicHostname pulumi.StringPtrInput
	// Configuration for OpenShift router(s).
	RouterProfiles OpenShiftRouterProfileResponseArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (OpenShiftManagedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftManagedClusterState)(nil)).Elem()
}

type openShiftManagedClusterArgs struct {
	// Configuration of OpenShift cluster VMs.
	AgentPoolProfiles []OpenShiftManagedClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	// Configures OpenShift authentication.
	AuthProfile *OpenShiftManagedClusterAuthProfile `pulumi:"authProfile"`
	// User-specified FQDN for OpenShift API server loadbalancer internal hostname.
	Fqdn *string `pulumi:"fqdn"`
	// Resource location
	Location *string `pulumi:"location"`
	// Configuration for OpenShift master VMs.
	MasterPoolProfile *OpenShiftManagedClusterMasterPoolProfile `pulumi:"masterPoolProfile"`
	// Configuration for OpenShift networking.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// Version of OpenShift specified when creating the cluster.
	OpenShiftVersion string `pulumi:"openShiftVersion"`
	// Define the resource plan as required by ARM for billing purposes
	Plan *PurchasePlan `pulumi:"plan"`
	// Optional user-specified FQDN for OpenShift API server.
	PublicHostname *string `pulumi:"publicHostname"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// Configuration for OpenShift router(s).
	RouterProfiles []OpenShiftRouterProfile `pulumi:"routerProfiles"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a OpenShiftManagedCluster resource.
type OpenShiftManagedClusterArgs struct {
	// Configuration of OpenShift cluster VMs.
	AgentPoolProfiles OpenShiftManagedClusterAgentPoolProfileArrayInput
	// Configures OpenShift authentication.
	AuthProfile OpenShiftManagedClusterAuthProfilePtrInput
	// User-specified FQDN for OpenShift API server loadbalancer internal hostname.
	Fqdn pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Configuration for OpenShift master VMs.
	MasterPoolProfile OpenShiftManagedClusterMasterPoolProfilePtrInput
	// Configuration for OpenShift networking.
	NetworkProfile NetworkProfilePtrInput
	// Version of OpenShift specified when creating the cluster.
	OpenShiftVersion pulumi.StringInput
	// Define the resource plan as required by ARM for billing purposes
	Plan PurchasePlanPtrInput
	// Optional user-specified FQDN for OpenShift API server.
	PublicHostname pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the OpenShift managed cluster resource.
	ResourceName pulumi.StringInput
	// Configuration for OpenShift router(s).
	RouterProfiles OpenShiftRouterProfileArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (OpenShiftManagedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftManagedClusterArgs)(nil)).Elem()
}

type OpenShiftManagedClusterInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput
	ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput
}

func (*OpenShiftManagedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedCluster)(nil))
}

func (i *OpenShiftManagedCluster) ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput {
	return i.ToOpenShiftManagedClusterOutputWithContext(context.Background())
}

func (i *OpenShiftManagedCluster) ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterOutput)
}

type OpenShiftManagedClusterOutput struct {
	*pulumi.OutputState
}

func (OpenShiftManagedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedCluster)(nil))
}

func (o OpenShiftManagedClusterOutput) ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput {
	return o
}

func (o OpenShiftManagedClusterOutput) ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OpenShiftManagedClusterOutput{})
}
