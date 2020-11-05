# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['Snapshot']


class Snapshot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 pool_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 volume_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Snapshot of a Volume

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account
        :param pulumi.Input[str] file_system_id: UUID v4 used to identify the FileSystem
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] pool_name: The name of the capacity pool
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] snapshot_name: The name of the mount target
        :param pulumi.Input[str] volume_name: The name of the volume
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
            __props__['file_system_id'] = file_system_id
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if pool_name is None:
                raise TypeError("Missing required property 'pool_name'")
            __props__['pool_name'] = pool_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if snapshot_name is None:
                raise TypeError("Missing required property 'snapshot_name'")
            __props__['snapshot_name'] = snapshot_name
            if volume_name is None:
                raise TypeError("Missing required property 'volume_name'")
            __props__['volume_name'] = volume_name
            __props__['created'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['snapshot_id'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:netapp/latest:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20170815:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20190501:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20190601:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20190801:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20191001:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20191101:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20200201:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20200301:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20200501:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20200601:Snapshot"), pulumi.Alias(type_="azure-nextgen:netapp/v20200701:Snapshot")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Snapshot, __self__).__init__(
            'azure-nextgen:netapp/v20190701:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Snapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        The creation date of the snapshot
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[Optional[str]]:
        """
        UUID v4 used to identify the FileSystem
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
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
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[str]:
        """
        UUID v4 used to identify the Snapshot
        """
        return pulumi.get(self, "snapshot_id")

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

