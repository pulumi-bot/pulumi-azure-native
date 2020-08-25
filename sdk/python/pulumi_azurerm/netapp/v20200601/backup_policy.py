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

__all__ = ['BackupPolicy']


class BackupPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 daily_backups_to_keep: Optional[pulumi.Input[float]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 monthly_backups_to_keep: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 volume_backups: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VolumeBackupsArgs']]]]] = None,
                 volumes_assigned: Optional[pulumi.Input[float]] = None,
                 weekly_backups_to_keep: Optional[pulumi.Input[float]] = None,
                 yearly_backups_to_keep: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Backup policy information

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account
        :param pulumi.Input[float] daily_backups_to_keep: Daily backups count to keep
        :param pulumi.Input[bool] enabled: The property to decide policy is enabled or not
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[float] monthly_backups_to_keep: Monthly backups count to keep
        :param pulumi.Input[str] name: Backup policy Name which uniquely identify backup policy.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VolumeBackupsArgs']]]] volume_backups: A list of volumes assigned to this policy
        :param pulumi.Input[float] volumes_assigned: Volumes using current backup policy
        :param pulumi.Input[float] weekly_backups_to_keep: Weekly backups count to keep
        :param pulumi.Input[float] yearly_backups_to_keep: Yearly backups count to keep
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
            __props__['daily_backups_to_keep'] = daily_backups_to_keep
            __props__['enabled'] = enabled
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['monthly_backups_to_keep'] = monthly_backups_to_keep
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['volume_backups'] = volume_backups
            __props__['volumes_assigned'] = volumes_assigned
            __props__['weekly_backups_to_keep'] = weekly_backups_to_keep
            __props__['yearly_backups_to_keep'] = yearly_backups_to_keep
            __props__['provisioning_state'] = None
            __props__['type'] = None
        super(BackupPolicy, __self__).__init__(
            'azurerm:netapp/v20200601:backupPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BackupPolicy':
        """
        Get an existing BackupPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BackupPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dailyBackupsToKeep")
    def daily_backups_to_keep(self) -> Optional[float]:
        """
        Daily backups count to keep
        """
        return pulumi.get(self, "daily_backups_to_keep")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        The property to decide policy is enabled or not
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="monthlyBackupsToKeep")
    def monthly_backups_to_keep(self) -> Optional[float]:
        """
        Monthly backups count to keep
        """
        return pulumi.get(self, "monthly_backups_to_keep")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of backup policy
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Azure lifecycle management
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="volumeBackups")
    def volume_backups(self) -> Optional[List['outputs.VolumeBackupsResponse']]:
        """
        A list of volumes assigned to this policy
        """
        return pulumi.get(self, "volume_backups")

    @property
    @pulumi.getter(name="volumesAssigned")
    def volumes_assigned(self) -> Optional[float]:
        """
        Volumes using current backup policy
        """
        return pulumi.get(self, "volumes_assigned")

    @property
    @pulumi.getter(name="weeklyBackupsToKeep")
    def weekly_backups_to_keep(self) -> Optional[float]:
        """
        Weekly backups count to keep
        """
        return pulumi.get(self, "weekly_backups_to_keep")

    @property
    @pulumi.getter(name="yearlyBackupsToKeep")
    def yearly_backups_to_keep(self) -> Optional[float]:
        """
        Yearly backups count to keep
        """
        return pulumi.get(self, "yearly_backups_to_keep")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

