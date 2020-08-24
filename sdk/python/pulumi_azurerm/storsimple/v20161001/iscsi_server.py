# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['IscsiServer']


class IscsiServer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_schedule_group_id: Optional[pulumi.Input[str]] = None,
                 chap_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 manager_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 reverse_chap_id: Optional[pulumi.Input[str]] = None,
                 storage_domain_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The iSCSI server.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_schedule_group_id: The backup policy id.
        :param pulumi.Input[str] chap_id: The chap id.
        :param pulumi.Input[str] description: The description.
        :param pulumi.Input[str] device_name: The device name.
        :param pulumi.Input[str] manager_name: The manager name
        :param pulumi.Input[str] name: The iSCSI server name.
        :param pulumi.Input[str] resource_group_name: The resource group name
        :param pulumi.Input[str] reverse_chap_id: The reverse chap id.
        :param pulumi.Input[str] storage_domain_id: The storage domain id.
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

            if backup_schedule_group_id is None:
                raise TypeError("Missing required property 'backup_schedule_group_id'")
            __props__['backup_schedule_group_id'] = backup_schedule_group_id
            __props__['chap_id'] = chap_id
            __props__['description'] = description
            if device_name is None:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            if manager_name is None:
                raise TypeError("Missing required property 'manager_name'")
            __props__['manager_name'] = manager_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['reverse_chap_id'] = reverse_chap_id
            if storage_domain_id is None:
                raise TypeError("Missing required property 'storage_domain_id'")
            __props__['storage_domain_id'] = storage_domain_id
            __props__['type'] = None
        super(IscsiServer, __self__).__init__(
            'azurerm:storsimple/v20161001:IscsiServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IscsiServer':
        """
        Get an existing IscsiServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IscsiServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupScheduleGroupId")
    def backup_schedule_group_id(self) -> str:
        """
        The backup policy id.
        """
        return pulumi.get(self, "backup_schedule_group_id")

    @property
    @pulumi.getter(name="chapId")
    def chap_id(self) -> Optional[str]:
        """
        The chap id.
        """
        return pulumi.get(self, "chap_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="reverseChapId")
    def reverse_chap_id(self) -> Optional[str]:
        """
        The reverse chap id.
        """
        return pulumi.get(self, "reverse_chap_id")

    @property
    @pulumi.getter(name="storageDomainId")
    def storage_domain_id(self) -> str:
        """
        The storage domain id.
        """
        return pulumi.get(self, "storage_domain_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

