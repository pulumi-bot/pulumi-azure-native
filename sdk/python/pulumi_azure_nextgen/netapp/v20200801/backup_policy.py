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

__all__ = ['BackupPolicy']


class BackupPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 backup_policy_name: Optional[pulumi.Input[str]] = None,
                 daily_backups_to_keep: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 monthly_backups_to_keep: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 volume_backups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeBackupsArgs']]]]] = None,
                 volumes_assigned: Optional[pulumi.Input[int]] = None,
                 weekly_backups_to_keep: Optional[pulumi.Input[int]] = None,
                 yearly_backups_to_keep: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Backup policy information

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account
        :param pulumi.Input[str] backup_policy_name: Backup policy Name which uniquely identify backup policy.
        :param pulumi.Input[int] daily_backups_to_keep: Daily backups count to keep
        :param pulumi.Input[bool] enabled: The property to decide policy is enabled or not
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[int] monthly_backups_to_keep: Monthly backups count to keep
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeBackupsArgs']]]] volume_backups: A list of volumes assigned to this policy
        :param pulumi.Input[int] volumes_assigned: Volumes using current backup policy
        :param pulumi.Input[int] weekly_backups_to_keep: Weekly backups count to keep
        :param pulumi.Input[int] yearly_backups_to_keep: Yearly backups count to keep
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

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if backup_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'backup_policy_name'")
            __props__['backup_policy_name'] = backup_policy_name
            __props__['daily_backups_to_keep'] = daily_backups_to_keep
            __props__['enabled'] = enabled
            __props__['location'] = location
            __props__['monthly_backups_to_keep'] = monthly_backups_to_keep
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['volume_backups'] = volume_backups
            __props__['volumes_assigned'] = volumes_assigned
            __props__['weekly_backups_to_keep'] = weekly_backups_to_keep
            __props__['yearly_backups_to_keep'] = yearly_backups_to_keep
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:netapp/latest:BackupPolicy"), pulumi.Alias(type_="azure-nextgen:netapp/v20200501:BackupPolicy"), pulumi.Alias(type_="azure-nextgen:netapp/v20200601:BackupPolicy"), pulumi.Alias(type_="azure-nextgen:netapp/v20200701:BackupPolicy"), pulumi.Alias(type_="azure-nextgen:netapp/v20200901:BackupPolicy"), pulumi.Alias(type_="azure-nextgen:netapp/v20201101:BackupPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BackupPolicy, __self__).__init__(
            'azure-nextgen:netapp/v20200801:BackupPolicy',
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
    def daily_backups_to_keep(self) -> pulumi.Output[Optional[int]]:
        """
        Daily backups count to keep
        """
        return pulumi.get(self, "daily_backups_to_keep")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        The property to decide policy is enabled or not
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="monthlyBackupsToKeep")
    def monthly_backups_to_keep(self) -> pulumi.Output[Optional[int]]:
        """
        Monthly backups count to keep
        """
        return pulumi.get(self, "monthly_backups_to_keep")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of backup policy
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Azure lifecycle management
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="volumeBackups")
    def volume_backups(self) -> pulumi.Output[Optional[Sequence['outputs.VolumeBackupsResponse']]]:
        """
        A list of volumes assigned to this policy
        """
        return pulumi.get(self, "volume_backups")

    @property
    @pulumi.getter(name="volumesAssigned")
    def volumes_assigned(self) -> pulumi.Output[Optional[int]]:
        """
        Volumes using current backup policy
        """
        return pulumi.get(self, "volumes_assigned")

    @property
    @pulumi.getter(name="weeklyBackupsToKeep")
    def weekly_backups_to_keep(self) -> pulumi.Output[Optional[int]]:
        """
        Weekly backups count to keep
        """
        return pulumi.get(self, "weekly_backups_to_keep")

    @property
    @pulumi.getter(name="yearlyBackupsToKeep")
    def yearly_backups_to_keep(self) -> pulumi.Output[Optional[int]]:
        """
        Yearly backups count to keep
        """
        return pulumi.get(self, "yearly_backups_to_keep")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

