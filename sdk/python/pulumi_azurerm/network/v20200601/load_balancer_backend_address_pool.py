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

__all__ = ['LoadBalancerBackendAddressPool']


class LoadBalancerBackendAddressPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 load_balancer_backend_addresses: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['LoadBalancerBackendAddressArgs']]]]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Pool of backend IP addresses.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['LoadBalancerBackendAddressArgs']]]] load_balancer_backend_addresses: An array of backend addresses.
        :param pulumi.Input[str] load_balancer_name: The name of the load balancer.
        :param pulumi.Input[str] name: The name of the backend address pool.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
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

            __props__['id'] = id
            __props__['load_balancer_backend_addresses'] = load_balancer_backend_addresses
            if load_balancer_name is None:
                raise TypeError("Missing required property 'load_balancer_name'")
            __props__['load_balancer_name'] = load_balancer_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['backend_ip_configurations'] = None
            __props__['etag'] = None
            __props__['load_balancing_rules'] = None
            __props__['outbound_rule'] = None
            __props__['outbound_rules'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20170601:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20170801:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20170901:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20171001:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20171101:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20180101:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20180201:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20180401:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20180601:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20180701:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20180801:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20181001:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20181101:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20181201:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20190201:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20190401:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20190601:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20190701:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20190801:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20190901:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20191101:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20191201:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20200301:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20200401:LoadBalancerBackendAddressPool"), pulumi.Alias(type_="azurerm:network/v20200501:LoadBalancerBackendAddressPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(LoadBalancerBackendAddressPool, __self__).__init__(
            'azurerm:network/v20200601:LoadBalancerBackendAddressPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LoadBalancerBackendAddressPool':
        """
        Get an existing LoadBalancerBackendAddressPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return LoadBalancerBackendAddressPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendIPConfigurations")
    def backend_ip_configurations(self) -> List['outputs.NetworkInterfaceIPConfigurationResponse']:
        """
        An array of references to IP addresses defined in network interfaces.
        """
        return pulumi.get(self, "backend_ip_configurations")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="loadBalancerBackendAddresses")
    def load_balancer_backend_addresses(self) -> Optional[List['outputs.LoadBalancerBackendAddressResponse']]:
        """
        An array of backend addresses.
        """
        return pulumi.get(self, "load_balancer_backend_addresses")

    @property
    @pulumi.getter(name="loadBalancingRules")
    def load_balancing_rules(self) -> List['outputs.SubResourceResponse']:
        """
        An array of references to load balancing rules that use this backend address pool.
        """
        return pulumi.get(self, "load_balancing_rules")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundRule")
    def outbound_rule(self) -> 'outputs.SubResourceResponse':
        """
        A reference to an outbound rule that uses this backend address pool.
        """
        return pulumi.get(self, "outbound_rule")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> List['outputs.SubResourceResponse']:
        """
        An array of references to outbound rules that use this backend address pool.
        """
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the backend address pool resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

