# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VNetPeering(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Name of the virtual network peering resource
    """
    properties: pulumi.Output[dict]
    """
    List of properties for vNet Peering
      * `allow_forwarded_traffic` (`bool`) - Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
      * `allow_gateway_transit` (`bool`) - If gateway links can be used in remote virtual networking to link to this virtual network.
      * `allow_virtual_network_access` (`bool`) - Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
      * `databricks_address_space` (`dict`) - The reference to the databricks virtual network address space.
        * `address_prefixes` (`list`) - A list of address blocks reserved for this virtual network in CIDR notation.

      * `databricks_virtual_network` (`dict`) -  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
        * `id` (`str`) - The Id of the databricks virtual network.

      * `peering_state` (`str`) - The status of the virtual network peering.
      * `provisioning_state` (`str`) - The provisioning state of the virtual network peering resource.
      * `remote_address_space` (`dict`) - The reference to the remote virtual network address space.
      * `remote_virtual_network` (`dict`) -  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
      * `use_remote_gateways` (`bool`) - If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
    """
    type: pulumi.Output[str]
    """
    type of the virtual network peering resource
    """
    def __init__(__self__, resource_name, opts=None, allow_forwarded_traffic=None, allow_gateway_transit=None, allow_virtual_network_access=None, databricks_address_space=None, databricks_virtual_network=None, name=None, remote_address_space=None, remote_virtual_network=None, resource_group_name=None, use_remote_gateways=None, workspace_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Peerings in a VirtualNetwork resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_forwarded_traffic: Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
        :param pulumi.Input[bool] allow_gateway_transit: If gateway links can be used in remote virtual networking to link to this virtual network.
        :param pulumi.Input[bool] allow_virtual_network_access: Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
        :param pulumi.Input[dict] databricks_address_space: The reference to the databricks virtual network address space.
        :param pulumi.Input[dict] databricks_virtual_network:  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
        :param pulumi.Input[str] name: The name of the workspace vNet peering.
        :param pulumi.Input[dict] remote_address_space: The reference to the remote virtual network address space.
        :param pulumi.Input[dict] remote_virtual_network:  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[bool] use_remote_gateways: If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
        :param pulumi.Input[str] workspace_name: The name of the workspace.

        The **databricks_address_space** object supports the following:

          * `address_prefixes` (`pulumi.Input[list]`) - A list of address blocks reserved for this virtual network in CIDR notation.

        The **databricks_virtual_network** object supports the following:

          * `id` (`pulumi.Input[str]`) - The Id of the databricks virtual network.
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

            __props__['allow_forwarded_traffic'] = allow_forwarded_traffic
            __props__['allow_gateway_transit'] = allow_gateway_transit
            __props__['allow_virtual_network_access'] = allow_virtual_network_access
            __props__['databricks_address_space'] = databricks_address_space
            __props__['databricks_virtual_network'] = databricks_virtual_network
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['remote_address_space'] = remote_address_space
            if remote_virtual_network is None:
                raise TypeError("Missing required property 'remote_virtual_network'")
            __props__['remote_virtual_network'] = remote_virtual_network
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['use_remote_gateways'] = use_remote_gateways
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['properties'] = None
            __props__['type'] = None
        super(VNetPeering, __self__).__init__(
            'azurerm:databricks/v20180401:vNetPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VNetPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VNetPeering(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
