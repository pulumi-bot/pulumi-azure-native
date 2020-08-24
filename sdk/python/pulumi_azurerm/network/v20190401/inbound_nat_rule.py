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

__all__ = ['InboundNatRule']


class InboundNatRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_port: Optional[pulumi.Input[float]] = None,
                 enable_floating_ip: Optional[pulumi.Input[bool]] = None,
                 enable_tcp_reset: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 frontend_ip_configuration: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 frontend_port: Optional[pulumi.Input[float]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 idle_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Inbound NAT rule of the load balancer.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] backend_port: The port used for the internal endpoint. Acceptable values range from 1 to 65535.
        :param pulumi.Input[bool] enable_floating_ip: Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
        :param pulumi.Input[bool] enable_tcp_reset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] frontend_ip_configuration: A reference to frontend IP addresses.
        :param pulumi.Input[float] frontend_port: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[float] idle_timeout_in_minutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
        :param pulumi.Input[str] load_balancer_name: The name of the load balancer.
        :param pulumi.Input[str] name: The name of the inbound nat rule.
        :param pulumi.Input[str] protocol: The reference to the transport protocol used by the load balancing rule.
        :param pulumi.Input[str] provisioning_state: Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
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

            __props__['backend_port'] = backend_port
            __props__['enable_floating_ip'] = enable_floating_ip
            __props__['enable_tcp_reset'] = enable_tcp_reset
            __props__['etag'] = etag
            __props__['frontend_ip_configuration'] = frontend_ip_configuration
            __props__['frontend_port'] = frontend_port
            __props__['id'] = id
            __props__['idle_timeout_in_minutes'] = idle_timeout_in_minutes
            if load_balancer_name is None:
                raise TypeError("Missing required property 'load_balancer_name'")
            __props__['load_balancer_name'] = load_balancer_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['protocol'] = protocol
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['backend_ip_configuration'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20170601:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20170801:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20170901:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20171001:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20171101:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20180101:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20180201:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20180401:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20180601:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20180701:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20180801:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20181001:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20181101:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20181201:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20190201:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20190601:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20190701:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20190801:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20190901:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20191101:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20191201:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20200301:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20200401:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20200501:InboundNatRule"), pulumi.Alias(type_="azurerm:network/v20200601:InboundNatRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(InboundNatRule, __self__).__init__(
            'azurerm:network/v20190401:InboundNatRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'InboundNatRule':
        """
        Get an existing InboundNatRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return InboundNatRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendIPConfiguration")
    def backend_ip_configuration(self) -> 'outputs.NetworkInterfaceIPConfigurationResponse':
        """
        A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
        """
        return pulumi.get(self, "backend_ip_configuration")

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> Optional[float]:
        """
        The port used for the internal endpoint. Acceptable values range from 1 to 65535.
        """
        return pulumi.get(self, "backend_port")

    @property
    @pulumi.getter(name="enableFloatingIP")
    def enable_floating_ip(self) -> Optional[bool]:
        """
        Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
        """
        return pulumi.get(self, "enable_floating_ip")

    @property
    @pulumi.getter(name="enableTcpReset")
    def enable_tcp_reset(self) -> Optional[bool]:
        """
        Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        """
        return pulumi.get(self, "enable_tcp_reset")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="frontendIPConfiguration")
    def frontend_ip_configuration(self) -> Optional['outputs.SubResourceResponse']:
        """
        A reference to frontend IP addresses.
        """
        return pulumi.get(self, "frontend_ip_configuration")

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> Optional[float]:
        """
        The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
        """
        return pulumi.get(self, "frontend_port")

    @property
    @pulumi.getter(name="idleTimeoutInMinutes")
    def idle_timeout_in_minutes(self) -> Optional[float]:
        """
        The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
        """
        return pulumi.get(self, "idle_timeout_in_minutes")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        The reference to the transport protocol used by the load balancing rule.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

