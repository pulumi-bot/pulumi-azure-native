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

__all__ = ['Schedule']


class Schedule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 daily_recurrence: Optional[pulumi.Input[pulumi.InputType['DayDetailsArgs']]] = None,
                 hourly_recurrence: Optional[pulumi.Input[pulumi.InputType['HourDetailsArgs']]] = None,
                 lab_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_settings: Optional[pulumi.Input[pulumi.InputType['NotificationSettingsArgs']]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 task_type: Optional[pulumi.Input[str]] = None,
                 time_zone_id: Optional[pulumi.Input[str]] = None,
                 unique_identifier: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input[pulumi.InputType['WeekDetailsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A schedule.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DayDetailsArgs']] daily_recurrence: If the schedule will occur once each day of the week, specify the daily recurrence.
        :param pulumi.Input[pulumi.InputType['HourDetailsArgs']] hourly_recurrence: If the schedule will occur multiple times a day, specify the hourly recurrence.
        :param pulumi.Input[str] lab_name: The name of the lab.
        :param pulumi.Input[str] location: The location of the resource.
        :param pulumi.Input[str] name: The name of the schedule.
        :param pulumi.Input[pulumi.InputType['NotificationSettingsArgs']] notification_settings: Notification settings.
        :param pulumi.Input[str] provisioning_state: The provisioning status of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] status: The status of the schedule (i.e. Enabled, Disabled)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        :param pulumi.Input[str] target_resource_id: The resource ID to which the schedule belongs
        :param pulumi.Input[str] task_type: The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
        :param pulumi.Input[str] time_zone_id: The time zone ID (e.g. Pacific Standard time).
        :param pulumi.Input[str] unique_identifier: The unique immutable identifier of a resource (Guid).
        :param pulumi.Input[pulumi.InputType['WeekDetailsArgs']] weekly_recurrence: If the schedule will occur only some days of the week, specify the weekly recurrence.
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

            __props__['daily_recurrence'] = daily_recurrence
            __props__['hourly_recurrence'] = hourly_recurrence
            if lab_name is None:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['notification_settings'] = notification_settings
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['status'] = status
            __props__['tags'] = tags
            __props__['target_resource_id'] = target_resource_id
            __props__['task_type'] = task_type
            __props__['time_zone_id'] = time_zone_id
            __props__['unique_identifier'] = unique_identifier
            __props__['weekly_recurrence'] = weekly_recurrence
            __props__['created_date'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:devtestlab/v20180915:Schedule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Schedule, __self__).__init__(
            'azurerm:devtestlab/v20160515:Schedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Schedule':
        """
        Get an existing Schedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Schedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        The creation date of the schedule.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter(name="dailyRecurrence")
    def daily_recurrence(self) -> Optional['outputs.DayDetailsResponse']:
        """
        If the schedule will occur once each day of the week, specify the daily recurrence.
        """
        return pulumi.get(self, "daily_recurrence")

    @property
    @pulumi.getter(name="hourlyRecurrence")
    def hourly_recurrence(self) -> Optional['outputs.HourDetailsResponse']:
        """
        If the schedule will occur multiple times a day, specify the hourly recurrence.
        """
        return pulumi.get(self, "hourly_recurrence")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationSettings")
    def notification_settings(self) -> Optional['outputs.NotificationSettingsResponse']:
        """
        Notification settings.
        """
        return pulumi.get(self, "notification_settings")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the schedule (i.e. Enabled, Disabled)
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> Optional[str]:
        """
        The resource ID to which the schedule belongs
        """
        return pulumi.get(self, "target_resource_id")

    @property
    @pulumi.getter(name="taskType")
    def task_type(self) -> Optional[str]:
        """
        The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
        """
        return pulumi.get(self, "task_type")

    @property
    @pulumi.getter(name="timeZoneId")
    def time_zone_id(self) -> Optional[str]:
        """
        The time zone ID (e.g. Pacific Standard time).
        """
        return pulumi.get(self, "time_zone_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> Optional[str]:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")

    @property
    @pulumi.getter(name="weeklyRecurrence")
    def weekly_recurrence(self) -> Optional['outputs.WeekDetailsResponse']:
        """
        If the schedule will occur only some days of the week, specify the weekly recurrence.
        """
        return pulumi.get(self, "weekly_recurrence")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

