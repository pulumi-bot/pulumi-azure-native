// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
type AccessPolicyEntry struct {
	//  Application ID of the client making request on behalf of a principal
	ApplicationId *string `pulumi:"applicationId"`
	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
	ObjectId string `pulumi:"objectId"`
	// Permissions the identity has for keys, secrets and certificates.
	Permissions Permissions `pulumi:"permissions"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId string `pulumi:"tenantId"`
}

// AccessPolicyEntryInput is an input type that accepts AccessPolicyEntryArgs and AccessPolicyEntryOutput values.
// You can construct a concrete instance of `AccessPolicyEntryInput` via:
//
//          AccessPolicyEntryArgs{...}
type AccessPolicyEntryInput interface {
	pulumi.Input

	ToAccessPolicyEntryOutput() AccessPolicyEntryOutput
	ToAccessPolicyEntryOutputWithContext(context.Context) AccessPolicyEntryOutput
}

// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
type AccessPolicyEntryArgs struct {
	//  Application ID of the client making request on behalf of a principal
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
	ObjectId pulumi.StringInput `pulumi:"objectId"`
	// Permissions the identity has for keys, secrets and certificates.
	Permissions PermissionsInput `pulumi:"permissions"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (AccessPolicyEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyEntry)(nil)).Elem()
}

func (i AccessPolicyEntryArgs) ToAccessPolicyEntryOutput() AccessPolicyEntryOutput {
	return i.ToAccessPolicyEntryOutputWithContext(context.Background())
}

func (i AccessPolicyEntryArgs) ToAccessPolicyEntryOutputWithContext(ctx context.Context) AccessPolicyEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyEntryOutput)
}

// AccessPolicyEntryArrayInput is an input type that accepts AccessPolicyEntryArray and AccessPolicyEntryArrayOutput values.
// You can construct a concrete instance of `AccessPolicyEntryArrayInput` via:
//
//          AccessPolicyEntryArray{ AccessPolicyEntryArgs{...} }
type AccessPolicyEntryArrayInput interface {
	pulumi.Input

	ToAccessPolicyEntryArrayOutput() AccessPolicyEntryArrayOutput
	ToAccessPolicyEntryArrayOutputWithContext(context.Context) AccessPolicyEntryArrayOutput
}

type AccessPolicyEntryArray []AccessPolicyEntryInput

func (AccessPolicyEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyEntry)(nil)).Elem()
}

func (i AccessPolicyEntryArray) ToAccessPolicyEntryArrayOutput() AccessPolicyEntryArrayOutput {
	return i.ToAccessPolicyEntryArrayOutputWithContext(context.Background())
}

func (i AccessPolicyEntryArray) ToAccessPolicyEntryArrayOutputWithContext(ctx context.Context) AccessPolicyEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyEntryArrayOutput)
}

// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
type AccessPolicyEntryResponse struct {
	//  Application ID of the client making request on behalf of a principal
	ApplicationId *string `pulumi:"applicationId"`
	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
	ObjectId string `pulumi:"objectId"`
	// Permissions the identity has for keys, secrets and certificates.
	Permissions PermissionsResponse `pulumi:"permissions"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId string `pulumi:"tenantId"`
}

// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
type AccessPolicyEntryResponseOutput struct{ *pulumi.OutputState }

func (AccessPolicyEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyEntryResponse)(nil)).Elem()
}

func (o AccessPolicyEntryResponseOutput) ToAccessPolicyEntryResponseOutput() AccessPolicyEntryResponseOutput {
	return o
}

func (o AccessPolicyEntryResponseOutput) ToAccessPolicyEntryResponseOutputWithContext(ctx context.Context) AccessPolicyEntryResponseOutput {
	return o
}

//  Application ID of the client making request on behalf of a principal
func (o AccessPolicyEntryResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessPolicyEntryResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
func (o AccessPolicyEntryResponseOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyEntryResponse) string { return v.ObjectId }).(pulumi.StringOutput)
}

// Permissions the identity has for keys, secrets and certificates.
func (o AccessPolicyEntryResponseOutput) Permissions() PermissionsResponseOutput {
	return o.ApplyT(func(v AccessPolicyEntryResponse) PermissionsResponse { return v.Permissions }).(PermissionsResponseOutput)
}

// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
func (o AccessPolicyEntryResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyEntryResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type AccessPolicyEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessPolicyEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyEntryResponse)(nil)).Elem()
}

func (o AccessPolicyEntryResponseArrayOutput) ToAccessPolicyEntryResponseArrayOutput() AccessPolicyEntryResponseArrayOutput {
	return o
}

func (o AccessPolicyEntryResponseArrayOutput) ToAccessPolicyEntryResponseArrayOutputWithContext(ctx context.Context) AccessPolicyEntryResponseArrayOutput {
	return o
}

func (o AccessPolicyEntryResponseArrayOutput) Index(i pulumi.IntInput) AccessPolicyEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessPolicyEntryResponse {
		return vs[0].([]AccessPolicyEntryResponse)[vs[1].(int)]
	}).(AccessPolicyEntryResponseOutput)
}

// Permissions the identity has for keys, secrets, certificates and storage.
type Permissions struct {
	// Permissions to certificates
	Certificates []string `pulumi:"certificates"`
	// Permissions to keys
	Keys []string `pulumi:"keys"`
	// Permissions to secrets
	Secrets []string `pulumi:"secrets"`
	// Permissions to storage accounts
	Storage []string `pulumi:"storage"`
}

// PermissionsInput is an input type that accepts PermissionsArgs and PermissionsOutput values.
// You can construct a concrete instance of `PermissionsInput` via:
//
//          PermissionsArgs{...}
type PermissionsInput interface {
	pulumi.Input

	ToPermissionsOutput() PermissionsOutput
	ToPermissionsOutputWithContext(context.Context) PermissionsOutput
}

// Permissions the identity has for keys, secrets, certificates and storage.
type PermissionsArgs struct {
	// Permissions to certificates
	Certificates pulumi.StringArrayInput `pulumi:"certificates"`
	// Permissions to keys
	Keys pulumi.StringArrayInput `pulumi:"keys"`
	// Permissions to secrets
	Secrets pulumi.StringArrayInput `pulumi:"secrets"`
	// Permissions to storage accounts
	Storage pulumi.StringArrayInput `pulumi:"storage"`
}

func (PermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Permissions)(nil)).Elem()
}

func (i PermissionsArgs) ToPermissionsOutput() PermissionsOutput {
	return i.ToPermissionsOutputWithContext(context.Background())
}

func (i PermissionsArgs) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsOutput)
}

// Permissions the identity has for keys, secrets, certificates and storage.
type PermissionsResponse struct {
	// Permissions to certificates
	Certificates []string `pulumi:"certificates"`
	// Permissions to keys
	Keys []string `pulumi:"keys"`
	// Permissions to secrets
	Secrets []string `pulumi:"secrets"`
	// Permissions to storage accounts
	Storage []string `pulumi:"storage"`
}

// Permissions the identity has for keys, secrets, certificates and storage.
type PermissionsResponseOutput struct{ *pulumi.OutputState }

func (PermissionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponse)(nil)).Elem()
}

func (o PermissionsResponseOutput) ToPermissionsResponseOutput() PermissionsResponseOutput {
	return o
}

func (o PermissionsResponseOutput) ToPermissionsResponseOutputWithContext(ctx context.Context) PermissionsResponseOutput {
	return o
}

// Permissions to certificates
func (o PermissionsResponseOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

// Permissions to keys
func (o PermissionsResponseOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

// Permissions to secrets
func (o PermissionsResponseOutput) Secrets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []string { return v.Secrets }).(pulumi.StringArrayOutput)
}

// Permissions to storage accounts
func (o PermissionsResponseOutput) Storage() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []string { return v.Storage }).(pulumi.StringArrayOutput)
}

// SKU details
type Sku struct {
	// SKU family name
	Family string `pulumi:"family"`
	// SKU name to specify whether the key vault is a standard vault or a premium vault.
	Name string `pulumi:"name"`
}

// SkuInput is an input type that accepts SkuArgs and SkuOutput values.
// You can construct a concrete instance of `SkuInput` via:
//
//          SkuArgs{...}
type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

// SKU details
type SkuArgs struct {
	// SKU family name
	Family pulumi.StringInput `pulumi:"family"`
	// SKU name to specify whether the key vault is a standard vault or a premium vault.
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}

// SkuPtrInput is an input type that accepts SkuArgs, SkuPtr and SkuPtrOutput values.
// You can construct a concrete instance of `SkuPtrInput` via:
//
//          SkuArgs{...}
//
//  or:
//
//          nil
type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

// SKU details
type SkuResponse struct {
	// SKU family name
	Family string `pulumi:"family"`
	// SKU name to specify whether the key vault is a standard vault or a premium vault.
	Name string `pulumi:"name"`
}

// SKU details
type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyT(func(v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

// SKU family name
func (o SkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

// SKU name to specify whether the key vault is a standard vault or a premium vault.
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse { return *v }).(SkuResponseOutput)
}

// SKU family name
func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

// SKU name to specify whether the key vault is a standard vault or a premium vault.
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// Properties of the vault
type VaultProperties struct {
	// An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not required. Otherwise, access policies are required.
	AccessPolicies []AccessPolicyEntry `pulumi:"accessPolicies"`
	// The vault's create mode to indicate whether the vault need to be recovered or not.
	CreateMode *string `pulumi:"createMode"`
	// Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
	EnablePurgeProtection *bool `pulumi:"enablePurgeProtection"`
	// Property specifying whether recoverable deletion is enabled for this key vault. Setting this property to true activates the soft delete feature, whereby vaults or vault entities can be recovered after deletion. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
	EnableSoftDelete *bool `pulumi:"enableSoftDelete"`
	// Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
	EnabledForDeployment *bool `pulumi:"enabledForDeployment"`
	// Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
	EnabledForDiskEncryption *bool `pulumi:"enabledForDiskEncryption"`
	// Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
	EnabledForTemplateDeployment *bool `pulumi:"enabledForTemplateDeployment"`
	// SKU details
	Sku Sku `pulumi:"sku"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId string `pulumi:"tenantId"`
	// The URI of the vault for performing operations on keys and secrets.
	VaultUri *string `pulumi:"vaultUri"`
}

// VaultPropertiesInput is an input type that accepts VaultPropertiesArgs and VaultPropertiesOutput values.
// You can construct a concrete instance of `VaultPropertiesInput` via:
//
//          VaultPropertiesArgs{...}
type VaultPropertiesInput interface {
	pulumi.Input

	ToVaultPropertiesOutput() VaultPropertiesOutput
	ToVaultPropertiesOutputWithContext(context.Context) VaultPropertiesOutput
}

// Properties of the vault
type VaultPropertiesArgs struct {
	// An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not required. Otherwise, access policies are required.
	AccessPolicies AccessPolicyEntryArrayInput `pulumi:"accessPolicies"`
	// The vault's create mode to indicate whether the vault need to be recovered or not.
	CreateMode pulumi.StringPtrInput `pulumi:"createMode"`
	// Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
	EnablePurgeProtection pulumi.BoolPtrInput `pulumi:"enablePurgeProtection"`
	// Property specifying whether recoverable deletion is enabled for this key vault. Setting this property to true activates the soft delete feature, whereby vaults or vault entities can be recovered after deletion. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
	EnableSoftDelete pulumi.BoolPtrInput `pulumi:"enableSoftDelete"`
	// Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
	EnabledForDeployment pulumi.BoolPtrInput `pulumi:"enabledForDeployment"`
	// Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
	EnabledForDiskEncryption pulumi.BoolPtrInput `pulumi:"enabledForDiskEncryption"`
	// Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
	EnabledForTemplateDeployment pulumi.BoolPtrInput `pulumi:"enabledForTemplateDeployment"`
	// SKU details
	Sku SkuInput `pulumi:"sku"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
	// The URI of the vault for performing operations on keys and secrets.
	VaultUri pulumi.StringPtrInput `pulumi:"vaultUri"`
}

func (VaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultProperties)(nil)).Elem()
}

func (i VaultPropertiesArgs) ToVaultPropertiesOutput() VaultPropertiesOutput {
	return i.ToVaultPropertiesOutputWithContext(context.Background())
}

func (i VaultPropertiesArgs) ToVaultPropertiesOutputWithContext(ctx context.Context) VaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesOutput)
}

func (i VaultPropertiesArgs) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return i.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i VaultPropertiesArgs) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesOutput).ToVaultPropertiesPtrOutputWithContext(ctx)
}

// VaultPropertiesPtrInput is an input type that accepts VaultPropertiesArgs, VaultPropertiesPtr and VaultPropertiesPtrOutput values.
// You can construct a concrete instance of `VaultPropertiesPtrInput` via:
//
//          VaultPropertiesArgs{...}
//
//  or:
//
//          nil
type VaultPropertiesPtrInput interface {
	pulumi.Input

	ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput
	ToVaultPropertiesPtrOutputWithContext(context.Context) VaultPropertiesPtrOutput
}

type vaultPropertiesPtrType VaultPropertiesArgs

func VaultPropertiesPtr(v *VaultPropertiesArgs) VaultPropertiesPtrInput {
	return (*vaultPropertiesPtrType)(v)
}

func (*vaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultProperties)(nil)).Elem()
}

func (i *vaultPropertiesPtrType) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return i.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *vaultPropertiesPtrType) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesPtrOutput)
}

// Properties of the vault
type VaultPropertiesResponse struct {
	// An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not required. Otherwise, access policies are required.
	AccessPolicies []AccessPolicyEntryResponse `pulumi:"accessPolicies"`
	// The vault's create mode to indicate whether the vault need to be recovered or not.
	CreateMode *string `pulumi:"createMode"`
	// Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
	EnablePurgeProtection *bool `pulumi:"enablePurgeProtection"`
	// Property specifying whether recoverable deletion is enabled for this key vault. Setting this property to true activates the soft delete feature, whereby vaults or vault entities can be recovered after deletion. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
	EnableSoftDelete *bool `pulumi:"enableSoftDelete"`
	// Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
	EnabledForDeployment *bool `pulumi:"enabledForDeployment"`
	// Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
	EnabledForDiskEncryption *bool `pulumi:"enabledForDiskEncryption"`
	// Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
	EnabledForTemplateDeployment *bool `pulumi:"enabledForTemplateDeployment"`
	// SKU details
	Sku SkuResponse `pulumi:"sku"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId string `pulumi:"tenantId"`
	// The URI of the vault for performing operations on keys and secrets.
	VaultUri *string `pulumi:"vaultUri"`
}

// Properties of the vault
type VaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponse)(nil)).Elem()
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutputWithContext(ctx context.Context) VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return o.ToVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *VaultPropertiesResponse {
		return &v
	}).(VaultPropertiesResponsePtrOutput)
}

// An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not required. Otherwise, access policies are required.
func (o VaultPropertiesResponseOutput) AccessPolicies() AccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) []AccessPolicyEntryResponse { return v.AccessPolicies }).(AccessPolicyEntryResponseArrayOutput)
}

// The vault's create mode to indicate whether the vault need to be recovered or not.
func (o VaultPropertiesResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

// Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
func (o VaultPropertiesResponseOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnablePurgeProtection }).(pulumi.BoolPtrOutput)
}

// Property specifying whether recoverable deletion is enabled for this key vault. Setting this property to true activates the soft delete feature, whereby vaults or vault entities can be recovered after deletion. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
func (o VaultPropertiesResponseOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnableSoftDelete }).(pulumi.BoolPtrOutput)
}

// Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
func (o VaultPropertiesResponseOutput) EnabledForDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnabledForDeployment }).(pulumi.BoolPtrOutput)
}

// Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
func (o VaultPropertiesResponseOutput) EnabledForDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnabledForDiskEncryption }).(pulumi.BoolPtrOutput)
}

// Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
func (o VaultPropertiesResponseOutput) EnabledForTemplateDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnabledForTemplateDeployment }).(pulumi.BoolPtrOutput)
}

// SKU details
func (o VaultPropertiesResponseOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
func (o VaultPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

// The URI of the vault for performing operations on keys and secrets.
func (o VaultPropertiesResponseOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.VaultUri }).(pulumi.StringPtrOutput)
}

type VaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponse)(nil)).Elem()
}

func (o VaultPropertiesResponsePtrOutput) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return o
}

func (o VaultPropertiesResponsePtrOutput) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return o
}

func (o VaultPropertiesResponsePtrOutput) Elem() VaultPropertiesResponseOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) VaultPropertiesResponse { return *v }).(VaultPropertiesResponseOutput)
}

// An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not required. Otherwise, access policies are required.
func (o VaultPropertiesResponsePtrOutput) AccessPolicies() AccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) []AccessPolicyEntryResponse {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(AccessPolicyEntryResponseArrayOutput)
}

// The vault's create mode to indicate whether the vault need to be recovered or not.
func (o VaultPropertiesResponsePtrOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreateMode
	}).(pulumi.StringPtrOutput)
}

// Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
func (o VaultPropertiesResponsePtrOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePurgeProtection
	}).(pulumi.BoolPtrOutput)
}

// Property specifying whether recoverable deletion is enabled for this key vault. Setting this property to true activates the soft delete feature, whereby vaults or vault entities can be recovered after deletion. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
func (o VaultPropertiesResponsePtrOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableSoftDelete
	}).(pulumi.BoolPtrOutput)
}

// Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
func (o VaultPropertiesResponsePtrOutput) EnabledForDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForDeployment
	}).(pulumi.BoolPtrOutput)
}

// Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
func (o VaultPropertiesResponsePtrOutput) EnabledForDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForDiskEncryption
	}).(pulumi.BoolPtrOutput)
}

// Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
func (o VaultPropertiesResponsePtrOutput) EnabledForTemplateDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForTemplateDeployment
	}).(pulumi.BoolPtrOutput)
}

// SKU details
func (o VaultPropertiesResponsePtrOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *SkuResponse {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(SkuResponsePtrOutput)
}

// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
func (o VaultPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The URI of the vault for performing operations on keys and secrets.
func (o VaultPropertiesResponsePtrOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VaultUri
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(PermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesOutput{})
	pulumi.RegisterOutputType(VaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponsePtrOutput{})
}
