# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['BatchAccount']


class BatchAccount(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 auto_storage: Optional[pulumi.Input[pulumi.InputType['AutoStorageBasePropertiesArgs']]] = None,
                 encryption: Optional[pulumi.Input[pulumi.InputType['EncryptionPropertiesArgs']]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['BatchAccountIdentityArgs']]] = None,
                 key_vault_reference: Optional[pulumi.Input[pulumi.InputType['KeyVaultReferenceArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 pool_allocation_mode: Optional[pulumi.Input[str]] = None,
                 public_network_access: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Contains information about an Azure Batch account.

        ## Example Usage
        ### BatchAccountCreate_BYOS

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        batch_account = azurerm.batch.v20200501.BatchAccount("batchAccount",
            account_name="sampleacct",
            auto_storage={
                "storageAccountId": "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage",
            },
            key_vault_reference={
                "id": "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.KeyVault/vaults/sample",
                "url": "http://sample.vault.azure.net/",
            },
            location="japaneast",
            pool_allocation_mode="UserSubscription",
            resource_group_name="default-azurebatch-japaneast")

        ```
        ### BatchAccountCreate_Default

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        batch_account = azurerm.batch.v20200501.BatchAccount("batchAccount",
            account_name="sampleacct",
            auto_storage={
                "storageAccountId": "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage",
            },
            location="japaneast",
            resource_group_name="default-azurebatch-japaneast")

        ```
        ### BatchAccountCreate_SystemAssignedIdentity

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        batch_account = azurerm.batch.v20200501.BatchAccount("batchAccount",
            account_name="sampleacct",
            auto_storage={
                "storageAccountId": "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage",
            },
            identity={
                "type": "SystemAssigned",
            },
            location="japaneast",
            resource_group_name="default-azurebatch-japaneast")

        ```
        ### PrivateBatchAccountCreate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        batch_account = azurerm.batch.v20200501.BatchAccount("batchAccount",
            account_name="sampleacct",
            auto_storage={
                "storageAccountId": "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage",
            },
            key_vault_reference={
                "id": "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.KeyVault/vaults/sample",
                "url": "http://sample.vault.azure.net/",
            },
            location="japaneast",
            public_network_access="Disabled",
            resource_group_name="default-azurebatch-japaneast")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: A name for the Batch account which must be unique within the region. Batch account names must be between 3 and 24 characters in length and must use only numbers and lowercase letters. This name is used as part of the DNS name that is used to access the Batch service in the region in which the account is created. For example: http://accountname.region.batch.azure.com/.
        :param pulumi.Input[pulumi.InputType['AutoStorageBasePropertiesArgs']] auto_storage: The properties related to the auto-storage account.
        :param pulumi.Input[pulumi.InputType['EncryptionPropertiesArgs']] encryption: Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using a Microsoft managed key. For additional control, a customer-managed key can be used instead.
        :param pulumi.Input[pulumi.InputType['BatchAccountIdentityArgs']] identity: The identity of the Batch account.
        :param pulumi.Input[pulumi.InputType['KeyVaultReferenceArgs']] key_vault_reference: A reference to the Azure key vault associated with the Batch account.
        :param pulumi.Input[str] location: The region in which to create the account.
        :param pulumi.Input[str] pool_allocation_mode: The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is UserSubscription, clients must use Azure Active Directory. The default is BatchService.
        :param pulumi.Input[str] public_network_access: If not specified, the default value is 'enabled'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the Batch account.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The user-specified tags associated with the account.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['auto_storage'] = auto_storage
            __props__['encryption'] = encryption
            __props__['identity'] = identity
            __props__['key_vault_reference'] = key_vault_reference
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['pool_allocation_mode'] = pool_allocation_mode
            __props__['public_network_access'] = public_network_access
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['account_endpoint'] = None
            __props__['active_job_and_job_schedule_quota'] = None
            __props__['dedicated_core_quota'] = None
            __props__['dedicated_core_quota_per_vm_family'] = None
            __props__['dedicated_core_quota_per_vm_family_enforced'] = None
            __props__['low_priority_core_quota'] = None
            __props__['name'] = None
            __props__['pool_quota'] = None
            __props__['private_endpoint_connections'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:batch/latest:BatchAccount"), pulumi.Alias(type_="azurerm:batch/v20151201:BatchAccount"), pulumi.Alias(type_="azurerm:batch/v20170101:BatchAccount"), pulumi.Alias(type_="azurerm:batch/v20170501:BatchAccount"), pulumi.Alias(type_="azurerm:batch/v20170901:BatchAccount"), pulumi.Alias(type_="azurerm:batch/v20181201:BatchAccount"), pulumi.Alias(type_="azurerm:batch/v20190401:BatchAccount"), pulumi.Alias(type_="azurerm:batch/v20190801:BatchAccount"), pulumi.Alias(type_="azurerm:batch/v20200301:BatchAccount")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BatchAccount, __self__).__init__(
            'azurerm:batch/v20200501:BatchAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BatchAccount':
        """
        Get an existing BatchAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BatchAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountEndpoint")
    def account_endpoint(self) -> pulumi.Output[str]:
        """
        The account endpoint used to interact with the Batch service.
        """
        return pulumi.get(self, "account_endpoint")

    @property
    @pulumi.getter(name="activeJobAndJobScheduleQuota")
    def active_job_and_job_schedule_quota(self) -> pulumi.Output[float]:
        return pulumi.get(self, "active_job_and_job_schedule_quota")

    @property
    @pulumi.getter(name="autoStorage")
    def auto_storage(self) -> pulumi.Output['outputs.AutoStoragePropertiesResponse']:
        """
        Contains information about the auto-storage account associated with a Batch account.
        """
        return pulumi.get(self, "auto_storage")

    @property
    @pulumi.getter(name="dedicatedCoreQuota")
    def dedicated_core_quota(self) -> pulumi.Output[float]:
        """
        For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
        """
        return pulumi.get(self, "dedicated_core_quota")

    @property
    @pulumi.getter(name="dedicatedCoreQuotaPerVMFamily")
    def dedicated_core_quota_per_vm_family(self) -> pulumi.Output[List['outputs.VirtualMachineFamilyCoreQuotaResponse']]:
        """
        A list of the dedicated core quota per Virtual Machine family for the Batch account. For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
        """
        return pulumi.get(self, "dedicated_core_quota_per_vm_family")

    @property
    @pulumi.getter(name="dedicatedCoreQuotaPerVMFamilyEnforced")
    def dedicated_core_quota_per_vm_family_enforced(self) -> pulumi.Output[bool]:
        """
        Batch is transitioning its core quota system for dedicated cores to be enforced per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not yet be enforced. If this flag is false, dedicated core quota is enforced via the old dedicatedCoreQuota property on the account and does not consider Virtual Machine family. If this flag is true, dedicated core quota is enforced via the dedicatedCoreQuotaPerVMFamily property on the account, and the old dedicatedCoreQuota does not apply.
        """
        return pulumi.get(self, "dedicated_core_quota_per_vm_family_enforced")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output['outputs.EncryptionPropertiesResponse']:
        """
        Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using a Microsoft managed key. For additional control, a customer-managed key can be used instead.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.BatchAccountIdentityResponse']]:
        """
        The identity of the Batch account.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="keyVaultReference")
    def key_vault_reference(self) -> pulumi.Output['outputs.KeyVaultReferenceResponse']:
        """
        Identifies the Azure key vault associated with a Batch account.
        """
        return pulumi.get(self, "key_vault_reference")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="lowPriorityCoreQuota")
    def low_priority_core_quota(self) -> pulumi.Output[float]:
        """
        For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
        """
        return pulumi.get(self, "low_priority_core_quota")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolAllocationMode")
    def pool_allocation_mode(self) -> pulumi.Output[str]:
        """
        The allocation mode for creating pools in the Batch account.
        """
        return pulumi.get(self, "pool_allocation_mode")

    @property
    @pulumi.getter(name="poolQuota")
    def pool_quota(self) -> pulumi.Output[float]:
        return pulumi.get(self, "pool_quota")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> pulumi.Output[List['outputs.PrivateEndpointConnectionResponse']]:
        """
        List of private endpoint connections associated with the Batch account
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioned state of the resource
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[str]:
        """
        If not specified, the default value is 'enabled'.
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

