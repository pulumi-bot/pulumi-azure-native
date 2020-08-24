# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VirtualNetworkPeering']


class VirtualNetworkPeering(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_forwarded_traffic: Optional[pulumi.Input[bool]] = None,
                 allow_gateway_transit: Optional[pulumi.Input[bool]] = None,
                 allow_virtual_network_access: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peering_state: Optional[pulumi.Input[str]] = None,
                 remote_address_space: Optional[pulumi.Input[pulumi.InputType['AddressSpaceArgs']]] = None,
                 remote_virtual_network: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 use_remote_gateways: Optional[pulumi.Input[bool]] = None,
                 virtual_network_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Peerings in a virtual network resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_forwarded_traffic: Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
        :param pulumi.Input[bool] allow_gateway_transit: If gateway links can be used in remote virtual networking to link to this virtual network.
        :param pulumi.Input[bool] allow_virtual_network_access: Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the peering.
        :param pulumi.Input[str] peering_state: The status of the virtual network peering.
        :param pulumi.Input[pulumi.InputType['AddressSpaceArgs']] remote_address_space: The reference to the remote virtual network address space.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] remote_virtual_network: The reference to the remote virtual network. The remote virtual network can be in the same or different region (preview). See here to register for the preview and learn more (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[bool] use_remote_gateways: If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
        :param pulumi.Input[str] virtual_network_name: The name of the virtual network.
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
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['peering_state'] = peering_state
            __props__['remote_address_space'] = remote_address_space
            __props__['remote_virtual_network'] = remote_virtual_network
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['use_remote_gateways'] = use_remote_gateways
            if virtual_network_name is None:
                raise TypeError("Missing required property 'virtual_network_name'")
            __props__['virtual_network_name'] = virtual_network_name
            __props__['etag'] = None
            __props__['provisioning_state'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20160601:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20160901:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20161201:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20170301:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20170601:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20170801:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20170901:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20171001:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20171101:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20180101:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20180201:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20180401:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20180601:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20180701:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20180801:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20181001:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20181101:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20181201:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20190201:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20190401:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20190601:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20190701:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20190801:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20190901:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20191101:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20191201:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20200301:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20200501:VirtualNetworkPeering"), pulumi.Alias(type_="azurerm:network/v20200601:VirtualNetworkPeering")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualNetworkPeering, __self__).__init__(
            'azurerm:network/v20200401:VirtualNetworkPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualNetworkPeering':
        """
        Get an existing VirtualNetworkPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualNetworkPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowForwardedTraffic")
    def allow_forwarded_traffic(self) -> Optional[bool]:
        """
        Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
        """
        return pulumi.get(self, "allow_forwarded_traffic")

    @property
    @pulumi.getter(name="allowGatewayTransit")
    def allow_gateway_transit(self) -> Optional[bool]:
        """
        If gateway links can be used in remote virtual networking to link to this virtual network.
        """
        return pulumi.get(self, "allow_gateway_transit")

    @property
    @pulumi.getter(name="allowVirtualNetworkAccess")
    def allow_virtual_network_access(self) -> Optional[bool]:
        """
        Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
        """
        return pulumi.get(self, "allow_virtual_network_access")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peeringState")
    def peering_state(self) -> Optional[str]:
        """
        The status of the virtual network peering.
        """
        return pulumi.get(self, "peering_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the virtual network peering resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="remoteAddressSpace")
    def remote_address_space(self) -> Optional['outputs.AddressSpaceResponse']:
        """
        The reference to the remote virtual network address space.
        """
        return pulumi.get(self, "remote_address_space")

    @property
    @pulumi.getter(name="remoteVirtualNetwork")
    def remote_virtual_network(self) -> Optional['outputs.SubResourceResponse']:
        """
        The reference to the remote virtual network. The remote virtual network can be in the same or different region (preview). See here to register for the preview and learn more (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
        """
        return pulumi.get(self, "remote_virtual_network")

    @property
    @pulumi.getter(name="useRemoteGateways")
    def use_remote_gateways(self) -> Optional[bool]:
        """
        If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
        """
        return pulumi.get(self, "use_remote_gateways")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

