# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['SqlPool']


class SqlPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collation: Optional[pulumi.Input[str]] = None,
                 create_mode: Optional[pulumi.Input[str]] = None,
                 creation_date: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_size_bytes: Optional[pulumi.Input[float]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 recoverable_database_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 restore_point_in_time: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 source_database_id: Optional[pulumi.Input[str]] = None,
                 sql_pool_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 storage_account_type: Optional[pulumi.Input[Union[str, 'StorageAccountType']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A SQL Analytics pool

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collation: Collation mode
        :param pulumi.Input[str] create_mode: What is this?
        :param pulumi.Input[str] creation_date: Date the SQL pool was created
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[float] max_size_bytes: Maximum size in bytes
        :param pulumi.Input[str] provisioning_state: Resource state
        :param pulumi.Input[str] recoverable_database_id: Backup database to restore from
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] restore_point_in_time: Snapshot time to restore
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: SQL pool SKU
        :param pulumi.Input[str] source_database_id: Source database to create from
        :param pulumi.Input[str] sql_pool_name: SQL pool name
        :param pulumi.Input[str] status: Resource status
        :param pulumi.Input[Union[str, 'StorageAccountType']] storage_account_type: The storage account type used to store backups for this sql pool.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] workspace_name: The name of the workspace
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

            __props__['collation'] = collation
            __props__['create_mode'] = create_mode
            __props__['creation_date'] = creation_date
            __props__['location'] = location
            __props__['max_size_bytes'] = max_size_bytes
            __props__['provisioning_state'] = provisioning_state
            __props__['recoverable_database_id'] = recoverable_database_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['restore_point_in_time'] = restore_point_in_time
            __props__['sku'] = sku
            __props__['source_database_id'] = source_database_id
            __props__['sql_pool_name'] = sql_pool_name
            __props__['status'] = status
            __props__['storage_account_type'] = storage_account_type
            __props__['tags'] = tags
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:synapse/v20201201:SqlPool"), pulumi.Alias(type_="azure-native:synapse:SqlPool"), pulumi.Alias(type_="azure-nextgen:synapse:SqlPool"), pulumi.Alias(type_="azure-native:synapse/latest:SqlPool"), pulumi.Alias(type_="azure-nextgen:synapse/latest:SqlPool"), pulumi.Alias(type_="azure-native:synapse/v20190601preview:SqlPool"), pulumi.Alias(type_="azure-nextgen:synapse/v20190601preview:SqlPool"), pulumi.Alias(type_="azure-native:synapse/v20200401preview:SqlPool"), pulumi.Alias(type_="azure-nextgen:synapse/v20200401preview:SqlPool"), pulumi.Alias(type_="azure-native:synapse/v20210301:SqlPool"), pulumi.Alias(type_="azure-nextgen:synapse/v20210301:SqlPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SqlPool, __self__).__init__(
            'azure-native:synapse/v20201201:SqlPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SqlPool':
        """
        Get an existing SqlPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["collation"] = None
        __props__["create_mode"] = None
        __props__["creation_date"] = None
        __props__["location"] = None
        __props__["max_size_bytes"] = None
        __props__["name"] = None
        __props__["provisioning_state"] = None
        __props__["recoverable_database_id"] = None
        __props__["restore_point_in_time"] = None
        __props__["sku"] = None
        __props__["source_database_id"] = None
        __props__["status"] = None
        __props__["storage_account_type"] = None
        __props__["tags"] = None
        __props__["type"] = None
        return SqlPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def collation(self) -> pulumi.Output[Optional[str]]:
        """
        Collation mode
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> pulumi.Output[Optional[str]]:
        """
        What is this?
        """
        return pulumi.get(self, "create_mode")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[Optional[str]]:
        """
        Date the SQL pool was created
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxSizeBytes")
    def max_size_bytes(self) -> pulumi.Output[Optional[float]]:
        """
        Maximum size in bytes
        """
        return pulumi.get(self, "max_size_bytes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        Resource state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="recoverableDatabaseId")
    def recoverable_database_id(self) -> pulumi.Output[Optional[str]]:
        """
        Backup database to restore from
        """
        return pulumi.get(self, "recoverable_database_id")

    @property
    @pulumi.getter(name="restorePointInTime")
    def restore_point_in_time(self) -> pulumi.Output[Optional[str]]:
        """
        Snapshot time to restore
        """
        return pulumi.get(self, "restore_point_in_time")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        SQL pool SKU
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="sourceDatabaseId")
    def source_database_id(self) -> pulumi.Output[Optional[str]]:
        """
        Source database to create from
        """
        return pulumi.get(self, "source_database_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Resource status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageAccountType")
    def storage_account_type(self) -> pulumi.Output[Optional[str]]:
        """
        The storage account type used to store backups for this sql pool.
        """
        return pulumi.get(self, "storage_account_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

