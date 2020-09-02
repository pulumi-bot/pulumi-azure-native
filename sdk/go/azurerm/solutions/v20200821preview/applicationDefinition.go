// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200821preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about managed application definition.
type ApplicationDefinition struct {
	pulumi.CustomResourceState

	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts ApplicationDefinitionArtifactResponseArrayOutput `pulumi:"artifacts"`
	// The managed application provider authorizations.
	Authorizations ApplicationAuthorizationResponseArrayOutput `pulumi:"authorizations"`
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition pulumi.MapOutput `pulumi:"createUiDefinition"`
	// The managed application deployment policy.
	DeploymentPolicy ApplicationDeploymentPolicyResponsePtrOutput `pulumi:"deploymentPolicy"`
	// The managed application definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The managed application definition display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// A value indicating whether the package is enabled or not.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The managed application lock level.
	LockLevel pulumi.StringOutput `pulumi:"lockLevel"`
	// The managed application locking policy.
	LockingPolicy ApplicationPackageLockingPolicyDefinitionResponsePtrOutput `pulumi:"lockingPolicy"`
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate pulumi.MapOutput `pulumi:"mainTemplate"`
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// The managed application management policy that determines publisher's access to the managed resource group.
	ManagementPolicy ApplicationManagementPolicyResponsePtrOutput `pulumi:"managementPolicy"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The managed application notification policy.
	NotificationPolicy ApplicationNotificationPolicyResponsePtrOutput `pulumi:"notificationPolicy"`
	// The managed application definition package file Uri. Use this element
	PackageFileUri pulumi.StringPtrOutput `pulumi:"packageFileUri"`
	// The managed application provider policies.
	Policies ApplicationPolicyResponseArrayOutput `pulumi:"policies"`
	// The SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The storage account id for bring your own storage scenario.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationDefinition registers a new resource with the given unique name, arguments, and options.
func NewApplicationDefinition(ctx *pulumi.Context,
	name string, args *ApplicationDefinitionArgs, opts ...pulumi.ResourceOption) (*ApplicationDefinition, error) {
	if args == nil || args.ApplicationDefinitionName == nil {
		return nil, errors.New("missing required argument 'ApplicationDefinitionName'")
	}
	if args == nil || args.LockLevel == nil {
		return nil, errors.New("missing required argument 'LockLevel'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationDefinitionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:solutions/latest:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azurerm:solutions/v20170901:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azurerm:solutions/v20180601:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azurerm:solutions/v20190701:ApplicationDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationDefinition
	err := ctx.RegisterResource("azurerm:solutions/v20200821preview:ApplicationDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationDefinition gets an existing ApplicationDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationDefinitionState, opts ...pulumi.ResourceOption) (*ApplicationDefinition, error) {
	var resource ApplicationDefinition
	err := ctx.ReadResource("azurerm:solutions/v20200821preview:ApplicationDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationDefinition resources.
type applicationDefinitionState struct {
	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts []ApplicationDefinitionArtifactResponse `pulumi:"artifacts"`
	// The managed application provider authorizations.
	Authorizations []ApplicationAuthorizationResponse `pulumi:"authorizations"`
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition map[string]interface{} `pulumi:"createUiDefinition"`
	// The managed application deployment policy.
	DeploymentPolicy *ApplicationDeploymentPolicyResponse `pulumi:"deploymentPolicy"`
	// The managed application definition description.
	Description *string `pulumi:"description"`
	// The managed application definition display name.
	DisplayName *string `pulumi:"displayName"`
	// A value indicating whether the package is enabled or not.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Resource location
	Location *string `pulumi:"location"`
	// The managed application lock level.
	LockLevel *string `pulumi:"lockLevel"`
	// The managed application locking policy.
	LockingPolicy *ApplicationPackageLockingPolicyDefinitionResponse `pulumi:"lockingPolicy"`
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate map[string]interface{} `pulumi:"mainTemplate"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The managed application management policy that determines publisher's access to the managed resource group.
	ManagementPolicy *ApplicationManagementPolicyResponse `pulumi:"managementPolicy"`
	// Resource name
	Name *string `pulumi:"name"`
	// The managed application notification policy.
	NotificationPolicy *ApplicationNotificationPolicyResponse `pulumi:"notificationPolicy"`
	// The managed application definition package file Uri. Use this element
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The managed application provider policies.
	Policies []ApplicationPolicyResponse `pulumi:"policies"`
	// The SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// The storage account id for bring your own storage scenario.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ApplicationDefinitionState struct {
	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts ApplicationDefinitionArtifactResponseArrayInput
	// The managed application provider authorizations.
	Authorizations ApplicationAuthorizationResponseArrayInput
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition pulumi.MapInput
	// The managed application deployment policy.
	DeploymentPolicy ApplicationDeploymentPolicyResponsePtrInput
	// The managed application definition description.
	Description pulumi.StringPtrInput
	// The managed application definition display name.
	DisplayName pulumi.StringPtrInput
	// A value indicating whether the package is enabled or not.
	IsEnabled pulumi.BoolPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The managed application lock level.
	LockLevel pulumi.StringPtrInput
	// The managed application locking policy.
	LockingPolicy ApplicationPackageLockingPolicyDefinitionResponsePtrInput
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate pulumi.MapInput
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrInput
	// The managed application management policy that determines publisher's access to the managed resource group.
	ManagementPolicy ApplicationManagementPolicyResponsePtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The managed application notification policy.
	NotificationPolicy ApplicationNotificationPolicyResponsePtrInput
	// The managed application definition package file Uri. Use this element
	PackageFileUri pulumi.StringPtrInput
	// The managed application provider policies.
	Policies ApplicationPolicyResponseArrayInput
	// The SKU of the resource.
	Sku SkuResponsePtrInput
	// The storage account id for bring your own storage scenario.
	StorageAccountId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ApplicationDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationDefinitionState)(nil)).Elem()
}

type applicationDefinitionArgs struct {
	// The name of the managed application definition.
	ApplicationDefinitionName string `pulumi:"applicationDefinitionName"`
	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts []ApplicationDefinitionArtifact `pulumi:"artifacts"`
	// The managed application provider authorizations.
	Authorizations []ApplicationAuthorization `pulumi:"authorizations"`
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition map[string]interface{} `pulumi:"createUiDefinition"`
	// The managed application deployment policy.
	DeploymentPolicy *ApplicationDeploymentPolicy `pulumi:"deploymentPolicy"`
	// The managed application definition description.
	Description *string `pulumi:"description"`
	// The managed application definition display name.
	DisplayName *string `pulumi:"displayName"`
	// A value indicating whether the package is enabled or not.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Resource location
	Location *string `pulumi:"location"`
	// The managed application lock level.
	LockLevel string `pulumi:"lockLevel"`
	// The managed application locking policy.
	LockingPolicy *ApplicationPackageLockingPolicyDefinition `pulumi:"lockingPolicy"`
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate map[string]interface{} `pulumi:"mainTemplate"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The managed application management policy that determines publisher's access to the managed resource group.
	ManagementPolicy *ApplicationManagementPolicy `pulumi:"managementPolicy"`
	// The managed application notification policy.
	NotificationPolicy *ApplicationNotificationPolicy `pulumi:"notificationPolicy"`
	// The managed application definition package file Uri. Use this element
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The managed application provider policies.
	Policies []ApplicationPolicy `pulumi:"policies"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// The storage account id for bring your own storage scenario.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationDefinition resource.
type ApplicationDefinitionArgs struct {
	// The name of the managed application definition.
	ApplicationDefinitionName pulumi.StringInput
	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts ApplicationDefinitionArtifactArrayInput
	// The managed application provider authorizations.
	Authorizations ApplicationAuthorizationArrayInput
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition pulumi.MapInput
	// The managed application deployment policy.
	DeploymentPolicy ApplicationDeploymentPolicyPtrInput
	// The managed application definition description.
	Description pulumi.StringPtrInput
	// The managed application definition display name.
	DisplayName pulumi.StringPtrInput
	// A value indicating whether the package is enabled or not.
	IsEnabled pulumi.BoolPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The managed application lock level.
	LockLevel pulumi.StringInput
	// The managed application locking policy.
	LockingPolicy ApplicationPackageLockingPolicyDefinitionPtrInput
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate pulumi.MapInput
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrInput
	// The managed application management policy that determines publisher's access to the managed resource group.
	ManagementPolicy ApplicationManagementPolicyPtrInput
	// The managed application notification policy.
	NotificationPolicy ApplicationNotificationPolicyPtrInput
	// The managed application definition package file Uri. Use this element
	PackageFileUri pulumi.StringPtrInput
	// The managed application provider policies.
	Policies ApplicationPolicyArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU of the resource.
	Sku SkuPtrInput
	// The storage account id for bring your own storage scenario.
	StorageAccountId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ApplicationDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationDefinitionArgs)(nil)).Elem()
}
