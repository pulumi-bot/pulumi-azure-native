# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['AccessControlRecord']


class AccessControlRecord(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_control_record_name: Optional[pulumi.Input[str]] = None,
                 initiator_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 manager_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The access control record.

        ## AccessControlRecordsCreateOrUpdate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        access_control_record = azurerm.storsimple.latest.AccessControlRecord("accessControlRecord",
            access_control_record_name="ACRForTest",
            initiator_name="iqn.2017-06.com.contoso:ForTest",
            manager_name="ManagerForSDKTest1",
            resource_group_name="ResourceGroupForSDKTest")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_control_record_name: The name of the access control record.
        :param pulumi.Input[str] initiator_name: The iSCSI initiator name (IQN).
        :param pulumi.Input[str] kind: The Kind of the object. Currently only Series8000 is supported
        :param pulumi.Input[str] manager_name: The manager name
        :param pulumi.Input[str] resource_group_name: The resource group name
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

            if access_control_record_name is None:
                raise TypeError("Missing required property 'access_control_record_name'")
            __props__['access_control_record_name'] = access_control_record_name
            if initiator_name is None:
                raise TypeError("Missing required property 'initiator_name'")
            __props__['initiator_name'] = initiator_name
            __props__['kind'] = kind
            if manager_name is None:
                raise TypeError("Missing required property 'manager_name'")
            __props__['manager_name'] = manager_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['name'] = None
            __props__['type'] = None
            __props__['volume_count'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:storsimple/v20161001:AccessControlRecord"), pulumi.Alias(type_="azurerm:storsimple/v20170601:AccessControlRecord")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AccessControlRecord, __self__).__init__(
            'azurerm:storsimple/latest:AccessControlRecord',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AccessControlRecord':
        """
        Get an existing AccessControlRecord resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AccessControlRecord(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="initiatorName")
    def initiator_name(self) -> pulumi.Output[str]:
        """
        The iSCSI initiator name (IQN).
        """
        return pulumi.get(self, "initiator_name")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        The Kind of the object. Currently only Series8000 is supported
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the object.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="volumeCount")
    def volume_count(self) -> pulumi.Output[float]:
        """
        The number of volumes using the access control record.
        """
        return pulumi.get(self, "volume_count")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
