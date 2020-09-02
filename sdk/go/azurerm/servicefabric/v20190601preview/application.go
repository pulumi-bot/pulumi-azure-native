// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The application resource.
type Application struct {
	pulumi.CustomResourceState

	// Azure resource etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Describes the managed identities for an Azure resource.
	Identity ManagedIdentityResponsePtrOutput `pulumi:"identity"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// List of user assigned identities for the application, each mapped to a friendly name.
	ManagedIdentities ApplicationUserAssignedIdentityResponseArrayOutput `pulumi:"managedIdentities"`
	// The maximum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. By default, the value of this property is zero and it means that the services can be placed on any node.
	MaximumNodes pulumi.IntPtrOutput `pulumi:"maximumNodes"`
	// List of application capacity metric description.
	Metrics ApplicationMetricDescriptionResponseArrayOutput `pulumi:"metrics"`
	// The minimum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. If this property is set to zero, no capacity will be reserved. The value of this property cannot be more than the value of the MaximumNodes property.
	MinimumNodes pulumi.IntPtrOutput `pulumi:"minimumNodes"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of application parameters with overridden values from their default values specified in the application manifest.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Remove the current application capacity settings.
	RemoveApplicationCapacity pulumi.BoolPtrOutput `pulumi:"removeApplicationCapacity"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The application type name as defined in the application manifest.
	TypeName pulumi.StringPtrOutput `pulumi:"typeName"`
	// The version of the application type as defined in the application manifest.
	TypeVersion pulumi.StringPtrOutput `pulumi:"typeVersion"`
	// Describes the policy for a monitored application upgrade.
	UpgradePolicy ApplicationUpgradePolicyResponsePtrOutput `pulumi:"upgradePolicy"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil || args.ApplicationName == nil {
		return nil, errors.New("missing required argument 'ApplicationName'")
	}
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:servicefabric/latest:Application"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20170701preview:Application"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301:Application"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301preview:Application"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20191101preview:Application"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20200301:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azurerm:servicefabric/v20190601preview:Application", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:servicefabric/v20190601preview:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// Azure resource etag.
	Etag *string `pulumi:"etag"`
	// Describes the managed identities for an Azure resource.
	Identity *ManagedIdentityResponse `pulumi:"identity"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// List of user assigned identities for the application, each mapped to a friendly name.
	ManagedIdentities []ApplicationUserAssignedIdentityResponse `pulumi:"managedIdentities"`
	// The maximum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. By default, the value of this property is zero and it means that the services can be placed on any node.
	MaximumNodes *int `pulumi:"maximumNodes"`
	// List of application capacity metric description.
	Metrics []ApplicationMetricDescriptionResponse `pulumi:"metrics"`
	// The minimum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. If this property is set to zero, no capacity will be reserved. The value of this property cannot be more than the value of the MaximumNodes property.
	MinimumNodes *int `pulumi:"minimumNodes"`
	// Azure resource name.
	Name *string `pulumi:"name"`
	// List of application parameters with overridden values from their default values specified in the application manifest.
	Parameters map[string]string `pulumi:"parameters"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState *string `pulumi:"provisioningState"`
	// Remove the current application capacity settings.
	RemoveApplicationCapacity *bool `pulumi:"removeApplicationCapacity"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type *string `pulumi:"type"`
	// The application type name as defined in the application manifest.
	TypeName *string `pulumi:"typeName"`
	// The version of the application type as defined in the application manifest.
	TypeVersion *string `pulumi:"typeVersion"`
	// Describes the policy for a monitored application upgrade.
	UpgradePolicy *ApplicationUpgradePolicyResponse `pulumi:"upgradePolicy"`
}

type ApplicationState struct {
	// Azure resource etag.
	Etag pulumi.StringPtrInput
	// Describes the managed identities for an Azure resource.
	Identity ManagedIdentityResponsePtrInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// List of user assigned identities for the application, each mapped to a friendly name.
	ManagedIdentities ApplicationUserAssignedIdentityResponseArrayInput
	// The maximum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. By default, the value of this property is zero and it means that the services can be placed on any node.
	MaximumNodes pulumi.IntPtrInput
	// List of application capacity metric description.
	Metrics ApplicationMetricDescriptionResponseArrayInput
	// The minimum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. If this property is set to zero, no capacity will be reserved. The value of this property cannot be more than the value of the MaximumNodes property.
	MinimumNodes pulumi.IntPtrInput
	// Azure resource name.
	Name pulumi.StringPtrInput
	// List of application parameters with overridden values from their default values specified in the application manifest.
	Parameters pulumi.StringMapInput
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringPtrInput
	// Remove the current application capacity settings.
	RemoveApplicationCapacity pulumi.BoolPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// Azure resource type.
	Type pulumi.StringPtrInput
	// The application type name as defined in the application manifest.
	TypeName pulumi.StringPtrInput
	// The version of the application type as defined in the application manifest.
	TypeVersion pulumi.StringPtrInput
	// Describes the policy for a monitored application upgrade.
	UpgradePolicy ApplicationUpgradePolicyResponsePtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The name of the application resource.
	ApplicationName string `pulumi:"applicationName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Describes the managed identities for an Azure resource.
	Identity *ManagedIdentity `pulumi:"identity"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// List of user assigned identities for the application, each mapped to a friendly name.
	ManagedIdentities []ApplicationUserAssignedIdentity `pulumi:"managedIdentities"`
	// The maximum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. By default, the value of this property is zero and it means that the services can be placed on any node.
	MaximumNodes *int `pulumi:"maximumNodes"`
	// List of application capacity metric description.
	Metrics []ApplicationMetricDescription `pulumi:"metrics"`
	// The minimum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. If this property is set to zero, no capacity will be reserved. The value of this property cannot be more than the value of the MaximumNodes property.
	MinimumNodes *int `pulumi:"minimumNodes"`
	// List of application parameters with overridden values from their default values specified in the application manifest.
	Parameters map[string]string `pulumi:"parameters"`
	// Remove the current application capacity settings.
	RemoveApplicationCapacity *bool `pulumi:"removeApplicationCapacity"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The application type name as defined in the application manifest.
	TypeName *string `pulumi:"typeName"`
	// The version of the application type as defined in the application manifest.
	TypeVersion *string `pulumi:"typeVersion"`
	// Describes the policy for a monitored application upgrade.
	UpgradePolicy *ApplicationUpgradePolicy `pulumi:"upgradePolicy"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The name of the application resource.
	ApplicationName pulumi.StringInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Describes the managed identities for an Azure resource.
	Identity ManagedIdentityPtrInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// List of user assigned identities for the application, each mapped to a friendly name.
	ManagedIdentities ApplicationUserAssignedIdentityArrayInput
	// The maximum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. By default, the value of this property is zero and it means that the services can be placed on any node.
	MaximumNodes pulumi.IntPtrInput
	// List of application capacity metric description.
	Metrics ApplicationMetricDescriptionArrayInput
	// The minimum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. If this property is set to zero, no capacity will be reserved. The value of this property cannot be more than the value of the MaximumNodes property.
	MinimumNodes pulumi.IntPtrInput
	// List of application parameters with overridden values from their default values specified in the application manifest.
	Parameters pulumi.StringMapInput
	// Remove the current application capacity settings.
	RemoveApplicationCapacity pulumi.BoolPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// The application type name as defined in the application manifest.
	TypeName pulumi.StringPtrInput
	// The version of the application type as defined in the application manifest.
	TypeVersion pulumi.StringPtrInput
	// Describes the policy for a monitored application upgrade.
	UpgradePolicy ApplicationUpgradePolicyPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
