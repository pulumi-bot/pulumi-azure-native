# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['IoMAMPolicyByName']


class IoMAMPolicyByName(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_recheck_offline_timeout: Optional[pulumi.Input[str]] = None,
                 access_recheck_online_timeout: Optional[pulumi.Input[str]] = None,
                 app_sharing_from_level: Optional[pulumi.Input[str]] = None,
                 app_sharing_to_level: Optional[pulumi.Input[str]] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 clipboard_sharing_level: Optional[pulumi.Input[str]] = None,
                 data_backup: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_compliance: Optional[pulumi.Input[str]] = None,
                 file_encryption_level: Optional[pulumi.Input[str]] = None,
                 file_sharing_save_as: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_browser: Optional[pulumi.Input[str]] = None,
                 offline_wipe_timeout: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 pin_num_retry: Optional[pulumi.Input[float]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 touch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        iOS Policy entity for Intune MAM.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host_name: Location hostName for the tenant
        :param pulumi.Input[str] location: Resource Location
        :param pulumi.Input[str] policy_name: Unique name for the policy
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource Tags
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

            __props__['access_recheck_offline_timeout'] = access_recheck_offline_timeout
            __props__['access_recheck_online_timeout'] = access_recheck_online_timeout
            __props__['app_sharing_from_level'] = app_sharing_from_level
            __props__['app_sharing_to_level'] = app_sharing_to_level
            __props__['authentication'] = authentication
            __props__['clipboard_sharing_level'] = clipboard_sharing_level
            __props__['data_backup'] = data_backup
            __props__['description'] = description
            __props__['device_compliance'] = device_compliance
            __props__['file_encryption_level'] = file_encryption_level
            __props__['file_sharing_save_as'] = file_sharing_save_as
            if friendly_name is None:
                raise TypeError("Missing required property 'friendly_name'")
            __props__['friendly_name'] = friendly_name
            if host_name is None:
                raise TypeError("Missing required property 'host_name'")
            __props__['host_name'] = host_name
            __props__['location'] = location
            __props__['managed_browser'] = managed_browser
            __props__['offline_wipe_timeout'] = offline_wipe_timeout
            __props__['pin'] = pin
            __props__['pin_num_retry'] = pin_num_retry
            if policy_name is None:
                raise TypeError("Missing required property 'policy_name'")
            __props__['policy_name'] = policy_name
            __props__['tags'] = tags
            __props__['touch_id'] = touch_id
            __props__['group_status'] = None
            __props__['last_modified_time'] = None
            __props__['name'] = None
            __props__['num_of_apps'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:intune/v20150114privatepreview:IoMAMPolicyByName")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(IoMAMPolicyByName, __self__).__init__(
            'azurerm:intune/v20150114preview:IoMAMPolicyByName',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IoMAMPolicyByName':
        """
        Get an existing IoMAMPolicyByName resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IoMAMPolicyByName(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessRecheckOfflineTimeout")
    def access_recheck_offline_timeout(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "access_recheck_offline_timeout")

    @property
    @pulumi.getter(name="accessRecheckOnlineTimeout")
    def access_recheck_online_timeout(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "access_recheck_online_timeout")

    @property
    @pulumi.getter(name="appSharingFromLevel")
    def app_sharing_from_level(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "app_sharing_from_level")

    @property
    @pulumi.getter(name="appSharingToLevel")
    def app_sharing_to_level(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "app_sharing_to_level")

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="clipboardSharingLevel")
    def clipboard_sharing_level(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "clipboard_sharing_level")

    @property
    @pulumi.getter(name="dataBackup")
    def data_backup(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "data_backup")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceCompliance")
    def device_compliance(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "device_compliance")

    @property
    @pulumi.getter(name="fileEncryptionLevel")
    def file_encryption_level(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "file_encryption_level")

    @property
    @pulumi.getter(name="fileSharingSaveAs")
    def file_sharing_save_as(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "file_sharing_save_as")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="groupStatus")
    def group_status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "group_status")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource Location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedBrowser")
    def managed_browser(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "managed_browser")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numOfApps")
    def num_of_apps(self) -> pulumi.Output[float]:
        return pulumi.get(self, "num_of_apps")

    @property
    @pulumi.getter(name="offlineWipeTimeout")
    def offline_wipe_timeout(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "offline_wipe_timeout")

    @property
    @pulumi.getter
    def pin(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "pin")

    @property
    @pulumi.getter(name="pinNumRetry")
    def pin_num_retry(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "pin_num_retry")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource Tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="touchId")
    def touch_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "touch_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

