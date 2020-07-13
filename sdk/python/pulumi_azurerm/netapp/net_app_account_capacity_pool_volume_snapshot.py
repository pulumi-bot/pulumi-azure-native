# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class NetAppAccountCapacityPoolVolumeSnapshot(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Snapshot Properties
      * `created` (`str`) - The creation date of the snapshot
      * `file_system_id` (`str`) - UUID v4 used to identify the FileSystem
      * `provisioning_state` (`str`) - Azure lifecycle management
      * `snapshot_id` (`str`) - UUID v4 used to identify the Snapshot
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, location=None, name=None, pool_name=None, properties=None, resource_group_name=None, volume_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Snapshot of a Volume

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the mount target
        :param pulumi.Input[str] pool_name: The name of the capacity pool
        :param pulumi.Input[dict] properties: Snapshot Properties
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] volume_name: The name of the volume

        The **properties** object supports the following:

          * `file_system_id` (`pulumi.Input[str]`) - UUID v4 used to identify the FileSystem
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if pool_name is None:
                raise TypeError("Missing required property 'pool_name'")
            __props__['pool_name'] = pool_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if volume_name is None:
                raise TypeError("Missing required property 'volume_name'")
            __props__['volume_name'] = volume_name
            __props__['type'] = None
        super(NetAppAccountCapacityPoolVolumeSnapshot, __self__).__init__(
            'azurerm:netapp:NetAppAccountCapacityPoolVolumeSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing NetAppAccountCapacityPoolVolumeSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return NetAppAccountCapacityPoolVolumeSnapshot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
