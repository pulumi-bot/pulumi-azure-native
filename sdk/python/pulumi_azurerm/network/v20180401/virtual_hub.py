# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VirtualHub(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Gets a unique read-only string that changes whenever the resource is updated.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Parameters for VirtualHub
      * `address_prefix` (`str`) - Address-prefix for this VirtualHub.
      * `hub_virtual_network_connections` (`list`) - list of all vnet connections with this VirtualHub.
        * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `location` (`str`) - Resource location.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Parameters for HubVirtualNetworkConnection
          * `allow_hub_to_remote_vnet_transit` (`bool`) - VirtualHub to RemoteVnet transit to enabled or not.
          * `allow_remote_vnet_to_use_hub_vnet_gateways` (`bool`) - Allow RemoteVnet to use Virtual Hub's gateways.
          * `provisioning_state` (`str`) - The provisioning state of the resource.
          * `remote_virtual_network` (`dict`) - Reference to the remote virtual network.
            * `id` (`str`) - Resource ID.

        * `tags` (`dict`) - Resource tags.
        * `type` (`str`) - Resource type.

      * `provisioning_state` (`str`) - The provisioning state of the resource.
      * `virtual_wan` (`dict`) - The VirtualWAN to which the VirtualHub belongs
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, address_prefix=None, hub_virtual_network_connections=None, id=None, location=None, name=None, provisioning_state=None, resource_group_name=None, tags=None, virtual_wan=None, __props__=None, __name__=None, __opts__=None):
        """
        VirtualHub Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: Address-prefix for this VirtualHub.
        :param pulumi.Input[list] hub_virtual_network_connections: list of all vnet connections with this VirtualHub.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the VirtualHub.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the resource.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VirtualHub.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[dict] virtual_wan: The VirtualWAN to which the VirtualHub belongs

        The **hub_virtual_network_connections** object supports the following:

          * `allow_hub_to_remote_vnet_transit` (`pulumi.Input[bool]`) - VirtualHub to RemoteVnet transit to enabled or not.
          * `allow_remote_vnet_to_use_hub_vnet_gateways` (`pulumi.Input[bool]`) - Allow RemoteVnet to use Virtual Hub's gateways.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `location` (`pulumi.Input[str]`) - Resource location.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource.
          * `remote_virtual_network` (`pulumi.Input[dict]`) - Reference to the remote virtual network.
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `tags` (`pulumi.Input[dict]`) - Resource tags.
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

            __props__['address_prefix'] = address_prefix
            __props__['hub_virtual_network_connections'] = hub_virtual_network_connections
            __props__['id'] = id
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_wan'] = virtual_wan
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(VirtualHub, __self__).__init__(
            'azurerm:network/v20180401:VirtualHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualHub(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
