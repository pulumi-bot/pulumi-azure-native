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

__all__ = ['NetworkInterface']


class NetworkInterface(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_settings: Optional[pulumi.Input[pulumi.InputType['NetworkInterfaceDnsSettingsArgs']]] = None,
                 enable_accelerated_networking: Optional[pulumi.Input[bool]] = None,
                 enable_ip_forwarding: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkInterfaceIPConfigurationArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_security_group: Optional[pulumi.Input[pulumi.InputType['NetworkSecurityGroupArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A network interface in a resource group.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NetworkInterfaceDnsSettingsArgs']] dns_settings: The DNS settings in network interface.
        :param pulumi.Input[bool] enable_accelerated_networking: If the network interface is accelerated networking enabled.
        :param pulumi.Input[bool] enable_ip_forwarding: Indicates whether IP forwarding is enabled on this network interface.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkInterfaceIPConfigurationArgs']]]] ip_configurations: A list of IPConfigurations of the network interface.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the network interface.
        :param pulumi.Input[pulumi.InputType['NetworkSecurityGroupArgs']] network_security_group: The reference to the NetworkSecurityGroup resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
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

            __props__['dns_settings'] = dns_settings
            __props__['enable_accelerated_networking'] = enable_accelerated_networking
            __props__['enable_ip_forwarding'] = enable_ip_forwarding
            __props__['id'] = id
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_security_group'] = network_security_group
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['hosted_workloads'] = None
            __props__['mac_address'] = None
            __props__['primary'] = None
            __props__['private_endpoint'] = None
            __props__['provisioning_state'] = None
            __props__['resource_guid'] = None
            __props__['tap_configurations'] = None
            __props__['type'] = None
            __props__['virtual_machine'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20150615:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20160330:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20160601:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20160901:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20161201:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20170301:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20170601:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20170801:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20170901:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20171001:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20171101:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20180101:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20180201:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20180401:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20180601:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20180701:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20180801:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20181001:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20181101:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20181201:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20190201:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20190401:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20190601:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20190701:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20190801:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20190901:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20191101:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20191201:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20200301:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20200401:NetworkInterface"), pulumi.Alias(type_="azurerm:network/v20200601:NetworkInterface")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NetworkInterface, __self__).__init__(
            'azurerm:network/v20200501:NetworkInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NetworkInterface':
        """
        Get an existing NetworkInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return NetworkInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsSettings")
    def dns_settings(self) -> Optional['outputs.NetworkInterfaceDnsSettingsResponse']:
        """
        The DNS settings in network interface.
        """
        return pulumi.get(self, "dns_settings")

    @property
    @pulumi.getter(name="enableAcceleratedNetworking")
    def enable_accelerated_networking(self) -> Optional[bool]:
        """
        If the network interface is accelerated networking enabled.
        """
        return pulumi.get(self, "enable_accelerated_networking")

    @property
    @pulumi.getter(name="enableIPForwarding")
    def enable_ip_forwarding(self) -> Optional[bool]:
        """
        Indicates whether IP forwarding is enabled on this network interface.
        """
        return pulumi.get(self, "enable_ip_forwarding")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="hostedWorkloads")
    def hosted_workloads(self) -> List[str]:
        """
        A list of references to linked BareMetal resources.
        """
        return pulumi.get(self, "hosted_workloads")

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> Optional[List['outputs.NetworkInterfaceIPConfigurationResponse']]:
        """
        A list of IPConfigurations of the network interface.
        """
        return pulumi.get(self, "ip_configurations")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> str:
        """
        The MAC address of the network interface.
        """
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkSecurityGroup")
    def network_security_group(self) -> Optional['outputs.NetworkSecurityGroupResponse']:
        """
        The reference to the NetworkSecurityGroup resource.
        """
        return pulumi.get(self, "network_security_group")

    @property
    @pulumi.getter
    def primary(self) -> bool:
        """
        Whether this is a primary network interface on a virtual machine.
        """
        return pulumi.get(self, "primary")

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> 'outputs.PrivateEndpointResponse':
        """
        A reference to the private endpoint to which the network interface is linked.
        """
        return pulumi.get(self, "private_endpoint")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the network interface resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> str:
        """
        The resource GUID property of the network interface resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tapConfigurations")
    def tap_configurations(self) -> List['outputs.NetworkInterfaceTapConfigurationResponse']:
        """
        A list of TapConfigurations of the network interface.
        """
        return pulumi.get(self, "tap_configurations")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualMachine")
    def virtual_machine(self) -> 'outputs.SubResourceResponse':
        """
        The reference to a virtual machine.
        """
        return pulumi.get(self, "virtual_machine")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

