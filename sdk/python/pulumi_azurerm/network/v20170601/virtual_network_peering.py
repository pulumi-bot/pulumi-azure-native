# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VirtualNetworkPeering(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    name: pulumi.Output[str]
    """
    The name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the virtual network peering.
      * `allow_forwarded_traffic` (`bool`) - Whether the forwarded traffic from the VMs in the remote virtual network will be allowed/disallowed.
      * `allow_gateway_transit` (`bool`) - If gateway links can be used in remote virtual networking to link to this virtual network.
      * `allow_virtual_network_access` (`bool`) - Whether the VMs in the linked virtual network space would be able to access all the VMs in local Virtual network space.
      * `peering_state` (`str`) - The status of the virtual network peering. Possible values are 'Initiated', 'Connected', and 'Disconnected'.
      * `provisioning_state` (`str`) - The provisioning state of the resource.
      * `remote_virtual_network` (`dict`) - The reference of the remote virtual network.
        * `id` (`str`) - Resource ID.

      * `use_remote_gateways` (`bool`) - If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
    """
    def __init__(__self__, resource_name, opts=None, etag=None, id=None, name=None, properties=None, resource_group_name=None, virtual_network_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Peerings in a virtual network resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the peering.
        :param pulumi.Input[dict] properties: Properties of the virtual network peering.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] virtual_network_name: The name of the virtual network.

        The **properties** object supports the following:

          * `allow_forwarded_traffic` (`pulumi.Input[bool]`) - Whether the forwarded traffic from the VMs in the remote virtual network will be allowed/disallowed.
          * `allow_gateway_transit` (`pulumi.Input[bool]`) - If gateway links can be used in remote virtual networking to link to this virtual network.
          * `allow_virtual_network_access` (`pulumi.Input[bool]`) - Whether the VMs in the linked virtual network space would be able to access all the VMs in local Virtual network space.
          * `peering_state` (`pulumi.Input[str]`) - The status of the virtual network peering. Possible values are 'Initiated', 'Connected', and 'Disconnected'.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource.
          * `remote_virtual_network` (`pulumi.Input[dict]`) - The reference of the remote virtual network.
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `use_remote_gateways` (`pulumi.Input[bool]`) - If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
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

            __props__['etag'] = etag
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if virtual_network_name is None:
                raise TypeError("Missing required property 'virtual_network_name'")
            __props__['virtual_network_name'] = virtual_network_name
        super(VirtualNetworkPeering, __self__).__init__(
            'azurerm:network/v20170601:VirtualNetworkPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualNetworkPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualNetworkPeering(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
