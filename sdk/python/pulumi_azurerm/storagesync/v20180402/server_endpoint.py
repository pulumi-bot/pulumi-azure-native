# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ServerEndpoint(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Server Endpoint properties.
      * `cloud_tiering` (`str`) - Cloud Tiering.
      * `friendly_name` (`str`) - Friendly Name
      * `last_operation_name` (`str`) - Resource Last Operation Name
      * `last_workflow_id` (`str`) - ServerEndpoint lastWorkflowId
      * `provisioning_state` (`str`) - ServerEndpoint Provisioning State
      * `server_local_path` (`str`) - Server Local path.
      * `server_resource_id` (`str`) - Server Resource Id.
      * `sync_status` (`dict`) - Sync Health Status
      * `volume_free_space_percent` (`float`) - Level of free space to be maintained by Cloud Tiering if it is enabled.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, cloud_tiering=None, friendly_name=None, location=None, name=None, resource_group_name=None, server_local_path=None, server_resource_id=None, storage_sync_service_name=None, sync_group_name=None, tags=None, volume_free_space_percent=None, __props__=None, __name__=None, __opts__=None):
        """
        Server Endpoint object.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_tiering: Cloud Tiering.
        :param pulumi.Input[str] friendly_name: Friendly Name
        :param pulumi.Input[str] location: Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
        :param pulumi.Input[str] name: Name of Server Endpoint object.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] server_local_path: Server Local path.
        :param pulumi.Input[str] server_resource_id: Server Resource Id.
        :param pulumi.Input[str] storage_sync_service_name: Name of Storage Sync Service resource.
        :param pulumi.Input[str] sync_group_name: Name of Sync Group resource.
        :param pulumi.Input[dict] tags: Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
        :param pulumi.Input[float] volume_free_space_percent: Level of free space to be maintained by Cloud Tiering if it is enabled.
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

            __props__['cloud_tiering'] = cloud_tiering
            __props__['friendly_name'] = friendly_name
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['server_local_path'] = server_local_path
            __props__['server_resource_id'] = server_resource_id
            if storage_sync_service_name is None:
                raise TypeError("Missing required property 'storage_sync_service_name'")
            __props__['storage_sync_service_name'] = storage_sync_service_name
            if sync_group_name is None:
                raise TypeError("Missing required property 'sync_group_name'")
            __props__['sync_group_name'] = sync_group_name
            __props__['tags'] = tags
            __props__['volume_free_space_percent'] = volume_free_space_percent
            __props__['properties'] = None
            __props__['type'] = None
        super(ServerEndpoint, __self__).__init__(
            'azurerm:storagesync/v20180402:ServerEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ServerEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ServerEndpoint(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
