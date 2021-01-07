# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['VirtualNetwork']


class VirtualNetwork(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_space: Optional[pulumi.Input[pulumi.InputType['AddressSpaceArgs']]] = None,
                 ddos_protection_plan: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 dhcp_options: Optional[pulumi.Input[pulumi.InputType['DhcpOptionsArgs']]] = None,
                 enable_ddos_protection: Optional[pulumi.Input[bool]] = None,
                 enable_vm_protection: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_guid: Optional[pulumi.Input[str]] = None,
                 subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_network_name: Optional[pulumi.Input[str]] = None,
                 virtual_network_peerings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkPeeringArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Virtual Network resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AddressSpaceArgs']] address_space: The AddressSpace that contains an array of IP address ranges that can be used by subnets.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] ddos_protection_plan: The DDoS protection plan associated with the virtual network.
        :param pulumi.Input[pulumi.InputType['DhcpOptionsArgs']] dhcp_options: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
        :param pulumi.Input[bool] enable_ddos_protection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
        :param pulumi.Input[bool] enable_vm_protection: Indicates if VM protection is enabled for all the subnets in the virtual network.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resourceGuid property of the Virtual Network resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetArgs']]]] subnets: A list of subnets in a Virtual Network.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] virtual_network_name: The name of the virtual network.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkPeeringArgs']]]] virtual_network_peerings: A list of peerings in a Virtual Network.
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

            __props__['address_space'] = address_space
            __props__['ddos_protection_plan'] = ddos_protection_plan
            __props__['dhcp_options'] = dhcp_options
            __props__['enable_ddos_protection'] = enable_ddos_protection
            __props__['enable_vm_protection'] = enable_vm_protection
            __props__['etag'] = etag
            __props__['id'] = id
            __props__['location'] = location
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['subnets'] = subnets
            __props__['tags'] = tags
            if virtual_network_name is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_network_name'")
            __props__['virtual_network_name'] = virtual_network_name
            __props__['virtual_network_peerings'] = virtual_network_peerings
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/latest:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20150501preview:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20150615:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20160330:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20160601:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20160901:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20161201:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20170301:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20170601:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20170801:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20170901:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20171001:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20171101:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20180101:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20180401:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20180601:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20180701:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20180801:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20181001:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20181101:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20181201:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20190201:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20190401:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20190601:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20190701:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20190801:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20190901:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20191101:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20191201:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20200301:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20200401:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20200501:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20200601:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20200701:VirtualNetwork"), pulumi.Alias(type_="azure-nextgen:network/v20200801:VirtualNetwork")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualNetwork, __self__).__init__(
            'azure-nextgen:network/v20180201:VirtualNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualNetwork':
        """
        Get an existing VirtualNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualNetwork(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressSpace")
    def address_space(self) -> pulumi.Output[Optional['outputs.AddressSpaceResponse']]:
        """
        The AddressSpace that contains an array of IP address ranges that can be used by subnets.
        """
        return pulumi.get(self, "address_space")

    @property
    @pulumi.getter(name="ddosProtectionPlan")
    def ddos_protection_plan(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        The DDoS protection plan associated with the virtual network.
        """
        return pulumi.get(self, "ddos_protection_plan")

    @property
    @pulumi.getter(name="dhcpOptions")
    def dhcp_options(self) -> pulumi.Output[Optional['outputs.DhcpOptionsResponse']]:
        """
        The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
        """
        return pulumi.get(self, "dhcp_options")

    @property
    @pulumi.getter(name="enableDdosProtection")
    def enable_ddos_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
        """
        return pulumi.get(self, "enable_ddos_protection")

    @property
    @pulumi.getter(name="enableVmProtection")
    def enable_vm_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if VM protection is enabled for all the subnets in the virtual network.
        """
        return pulumi.get(self, "enable_vm_protection")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        Gets a unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> pulumi.Output[Optional[str]]:
        """
        The resourceGuid property of the Virtual Network resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Output[Optional[Sequence['outputs.SubnetResponse']]]:
        """
        A list of subnets in a Virtual Network.
        """
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkPeerings")
    def virtual_network_peerings(self) -> pulumi.Output[Optional[Sequence['outputs.VirtualNetworkPeeringResponse']]]:
        """
        A list of peerings in a Virtual Network.
        """
        return pulumi.get(self, "virtual_network_peerings")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

