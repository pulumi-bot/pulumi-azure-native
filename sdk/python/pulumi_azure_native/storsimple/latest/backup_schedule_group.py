# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['BackupScheduleGroup']

warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:storsimple:BackupScheduleGroup'.""", DeprecationWarning)


class BackupScheduleGroup(pulumi.CustomResource):
    warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:storsimple:BackupScheduleGroup'.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 manager_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 schedule_group_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[pulumi.InputType['TimeArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The Backup Schedule Group
        Latest API Version: 2016-10-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: The name of the device.
        :param pulumi.Input[str] manager_name: The manager name
        :param pulumi.Input[str] resource_group_name: The resource group name
        :param pulumi.Input[str] schedule_group_name: The name of the schedule group.
        :param pulumi.Input[pulumi.InputType['TimeArgs']] start_time: The start time. When this field is specified we will generate Default GrandFather Father Son Backup Schedules.
        """
        pulumi.log.warn("""BackupScheduleGroup is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:storsimple:BackupScheduleGroup'.""")
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

            if device_name is None and not opts.urn:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            if manager_name is None and not opts.urn:
                raise TypeError("Missing required property 'manager_name'")
            __props__['manager_name'] = manager_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['schedule_group_name'] = schedule_group_name
            if start_time is None and not opts.urn:
                raise TypeError("Missing required property 'start_time'")
            __props__['start_time'] = start_time
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:storsimple/latest:BackupScheduleGroup"), pulumi.Alias(type_="azure-native:storsimple:BackupScheduleGroup"), pulumi.Alias(type_="azure-nextgen:storsimple:BackupScheduleGroup"), pulumi.Alias(type_="azure-native:storsimple/v20161001:BackupScheduleGroup"), pulumi.Alias(type_="azure-nextgen:storsimple/v20161001:BackupScheduleGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BackupScheduleGroup, __self__).__init__(
            'azure-native:storsimple/latest:BackupScheduleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BackupScheduleGroup':
        """
        Get an existing BackupScheduleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = None
        __props__["start_time"] = None
        __props__["type"] = None
        return BackupScheduleGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output['outputs.TimeResponse']:
        """
        The start time. When this field is specified we will generate Default GrandFather Father Son Backup Schedules.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

