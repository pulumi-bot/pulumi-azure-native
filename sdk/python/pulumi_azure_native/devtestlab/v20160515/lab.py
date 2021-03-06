# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['Lab']


class Lab(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lab_storage_type: Optional[pulumi.Input[Union[str, 'StorageType']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 premium_data_disks: Optional[pulumi.Input[Union[str, 'PremiumDataDisk']]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 unique_identifier: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A lab.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[str, 'StorageType']] lab_storage_type: Type of storage used by the lab. It can be either Premium or Standard. Default is Premium.
        :param pulumi.Input[str] location: The location of the resource.
        :param pulumi.Input[str] name: The name of the lab.
        :param pulumi.Input[Union[str, 'PremiumDataDisk']] premium_data_disks: The setting to enable usage of premium data disks.
               When its value is 'Enabled', creation of standard or premium data disks is allowed.
               When its value is 'Disabled', only creation of standard data disks is allowed.
        :param pulumi.Input[str] provisioning_state: The provisioning status of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        :param pulumi.Input[str] unique_identifier: The unique immutable identifier of a resource (Guid).
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

            __props__['lab_storage_type'] = lab_storage_type
            __props__['location'] = location
            __props__['name'] = name
            __props__['premium_data_disks'] = premium_data_disks
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['unique_identifier'] = unique_identifier
            __props__['artifacts_storage_account'] = None
            __props__['created_date'] = None
            __props__['default_premium_storage_account'] = None
            __props__['default_storage_account'] = None
            __props__['premium_data_disk_storage_account'] = None
            __props__['type'] = None
            __props__['vault_name'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:devtestlab/v20160515:Lab"), pulumi.Alias(type_="azure-native:devtestlab:Lab"), pulumi.Alias(type_="azure-nextgen:devtestlab:Lab"), pulumi.Alias(type_="azure-native:devtestlab/latest:Lab"), pulumi.Alias(type_="azure-nextgen:devtestlab/latest:Lab"), pulumi.Alias(type_="azure-native:devtestlab/v20150521preview:Lab"), pulumi.Alias(type_="azure-nextgen:devtestlab/v20150521preview:Lab"), pulumi.Alias(type_="azure-native:devtestlab/v20180915:Lab"), pulumi.Alias(type_="azure-nextgen:devtestlab/v20180915:Lab")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Lab, __self__).__init__(
            'azure-native:devtestlab/v20160515:Lab',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Lab':
        """
        Get an existing Lab resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["artifacts_storage_account"] = None
        __props__["created_date"] = None
        __props__["default_premium_storage_account"] = None
        __props__["default_storage_account"] = None
        __props__["lab_storage_type"] = None
        __props__["location"] = None
        __props__["name"] = None
        __props__["premium_data_disk_storage_account"] = None
        __props__["premium_data_disks"] = None
        __props__["provisioning_state"] = None
        __props__["tags"] = None
        __props__["type"] = None
        __props__["unique_identifier"] = None
        __props__["vault_name"] = None
        return Lab(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="artifactsStorageAccount")
    def artifacts_storage_account(self) -> pulumi.Output[str]:
        """
        The lab's artifact storage account.
        """
        return pulumi.get(self, "artifacts_storage_account")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> pulumi.Output[str]:
        """
        The creation date of the lab.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter(name="defaultPremiumStorageAccount")
    def default_premium_storage_account(self) -> pulumi.Output[str]:
        """
        The lab's default premium storage account.
        """
        return pulumi.get(self, "default_premium_storage_account")

    @property
    @pulumi.getter(name="defaultStorageAccount")
    def default_storage_account(self) -> pulumi.Output[str]:
        """
        The lab's default storage account.
        """
        return pulumi.get(self, "default_storage_account")

    @property
    @pulumi.getter(name="labStorageType")
    def lab_storage_type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of storage used by the lab. It can be either Premium or Standard. Default is Premium.
        """
        return pulumi.get(self, "lab_storage_type")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="premiumDataDiskStorageAccount")
    def premium_data_disk_storage_account(self) -> pulumi.Output[str]:
        """
        The lab's premium data disk storage account.
        """
        return pulumi.get(self, "premium_data_disk_storage_account")

    @property
    @pulumi.getter(name="premiumDataDisks")
    def premium_data_disks(self) -> pulumi.Output[Optional[str]]:
        """
        The setting to enable usage of premium data disks.
        When its value is 'Enabled', creation of standard or premium data disks is allowed.
        When its value is 'Disabled', only creation of standard data disks is allowed.
        """
        return pulumi.get(self, "premium_data_disks")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
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

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")

    @property
    @pulumi.getter(name="vaultName")
    def vault_name(self) -> pulumi.Output[str]:
        """
        The lab's Key vault.
        """
        return pulumi.get(self, "vault_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

