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

__all__ = ['AvailabilityGroupListener']


class AvailabilityGroupListener(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_group_listener_name: Optional[pulumi.Input[str]] = None,
                 availability_group_name: Optional[pulumi.Input[str]] = None,
                 create_default_availability_group_if_not_exist: Optional[pulumi.Input[bool]] = None,
                 load_balancer_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['LoadBalancerConfigurationArgs']]]]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sql_virtual_machine_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A SQL Server availability group listener.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_group_listener_name: Name of the availability group listener.
        :param pulumi.Input[str] availability_group_name: Name of the availability group.
        :param pulumi.Input[bool] create_default_availability_group_if_not_exist: Create a default availability group if it does not exist.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['LoadBalancerConfigurationArgs']]]] load_balancer_configurations: List of load balancer configurations for an availability group listener.
        :param pulumi.Input[float] port: Listener port.
        :param pulumi.Input[str] resource_group_name: Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[str] sql_virtual_machine_group_name: Name of the SQL virtual machine group.
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

            if availability_group_listener_name is None:
                raise TypeError("Missing required property 'availability_group_listener_name'")
            __props__['availability_group_listener_name'] = availability_group_listener_name
            __props__['availability_group_name'] = availability_group_name
            __props__['create_default_availability_group_if_not_exist'] = create_default_availability_group_if_not_exist
            __props__['load_balancer_configurations'] = load_balancer_configurations
            __props__['port'] = port
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sql_virtual_machine_group_name is None:
                raise TypeError("Missing required property 'sql_virtual_machine_group_name'")
            __props__['sql_virtual_machine_group_name'] = sql_virtual_machine_group_name
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:sqlvirtualmachine/latest:AvailabilityGroupListener")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AvailabilityGroupListener, __self__).__init__(
            'azurerm:sqlvirtualmachine/v20170301preview:AvailabilityGroupListener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AvailabilityGroupListener':
        """
        Get an existing AvailabilityGroupListener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AvailabilityGroupListener(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityGroupName")
    def availability_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the availability group.
        """
        return pulumi.get(self, "availability_group_name")

    @property
    @pulumi.getter(name="createDefaultAvailabilityGroupIfNotExist")
    def create_default_availability_group_if_not_exist(self) -> pulumi.Output[Optional[bool]]:
        """
        Create a default availability group if it does not exist.
        """
        return pulumi.get(self, "create_default_availability_group_if_not_exist")

    @property
    @pulumi.getter(name="loadBalancerConfigurations")
    def load_balancer_configurations(self) -> pulumi.Output[Optional[List['outputs.LoadBalancerConfigurationResponse']]]:
        """
        List of load balancer configurations for an availability group listener.
        """
        return pulumi.get(self, "load_balancer_configurations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[float]]:
        """
        Listener port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Provisioning state to track the async operation status.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

