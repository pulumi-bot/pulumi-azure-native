# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['BackupPolicy']

warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:storsimple:BackupPolicy'.""", DeprecationWarning)


class BackupPolicy(pulumi.CustomResource):
    warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:storsimple:BackupPolicy'.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_policy_name: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input['Kind']] = None,
                 manager_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 volume_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The backup policy.
        Latest API Version: 2017-06-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_policy_name: The name of the backup policy to be created/updated.
        :param pulumi.Input[str] device_name: The device name
        :param pulumi.Input['Kind'] kind: The Kind of the object. Currently only Series8000 is supported
        :param pulumi.Input[str] manager_name: The manager name
        :param pulumi.Input[str] resource_group_name: The resource group name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] volume_ids: The path IDs of the volumes which are part of the backup policy.
        """
        pulumi.log.warn("""BackupPolicy is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:storsimple:BackupPolicy'.""")
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

            __props__['backup_policy_name'] = backup_policy_name
            if device_name is None and not opts.urn:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            __props__['kind'] = kind
            if manager_name is None and not opts.urn:
                raise TypeError("Missing required property 'manager_name'")
            __props__['manager_name'] = manager_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if volume_ids is None and not opts.urn:
                raise TypeError("Missing required property 'volume_ids'")
            __props__['volume_ids'] = volume_ids
            __props__['backup_policy_creation_type'] = None
            __props__['last_backup_time'] = None
            __props__['name'] = None
            __props__['next_backup_time'] = None
            __props__['scheduled_backup_status'] = None
            __props__['schedules_count'] = None
            __props__['ssm_host_name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:storsimple/latest:BackupPolicy"), pulumi.Alias(type_="azure-native:storsimple:BackupPolicy"), pulumi.Alias(type_="azure-nextgen:storsimple:BackupPolicy"), pulumi.Alias(type_="azure-native:storsimple/v20170601:BackupPolicy"), pulumi.Alias(type_="azure-nextgen:storsimple/v20170601:BackupPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BackupPolicy, __self__).__init__(
            'azure-native:storsimple/latest:BackupPolicy',
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

        __props__["backup_policy_creation_type"] = None
        __props__["kind"] = None
        __props__["last_backup_time"] = None
        __props__["name"] = None
        __props__["next_backup_time"] = None
        __props__["scheduled_backup_status"] = None
        __props__["schedules_count"] = None
        __props__["ssm_host_name"] = None
        __props__["type"] = None
        __props__["volume_ids"] = None
        return BackupPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupPolicyCreationType")
    def backup_policy_creation_type(self) -> pulumi.Output[str]:
        """
        The backup policy creation type. Indicates whether this was created through SaaS or through StorSimple Snapshot Manager.
        """
        return pulumi.get(self, "backup_policy_creation_type")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        The Kind of the object. Currently only Series8000 is supported
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastBackupTime")
    def last_backup_time(self) -> pulumi.Output[str]:
        """
        The time of the last backup for the backup policy.
        """
        return pulumi.get(self, "last_backup_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the object.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextBackupTime")
    def next_backup_time(self) -> pulumi.Output[str]:
        """
        The time of the next backup for the backup policy.
        """
        return pulumi.get(self, "next_backup_time")

    @property
    @pulumi.getter(name="scheduledBackupStatus")
    def scheduled_backup_status(self) -> pulumi.Output[str]:
        """
        Indicates whether at least one of the schedules in the backup policy is active or not.
        """
        return pulumi.get(self, "scheduled_backup_status")

    @property
    @pulumi.getter(name="schedulesCount")
    def schedules_count(self) -> pulumi.Output[float]:
        """
        The count of schedules the backup policy contains.
        """
        return pulumi.get(self, "schedules_count")

    @property
    @pulumi.getter(name="ssmHostName")
    def ssm_host_name(self) -> pulumi.Output[str]:
        """
        If the backup policy was created by StorSimple Snapshot Manager, then this field indicates the hostname of the StorSimple Snapshot Manager.
        """
        return pulumi.get(self, "ssm_host_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="volumeIds")
    def volume_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The path IDs of the volumes which are part of the backup policy.
        """
        return pulumi.get(self, "volume_ids")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

