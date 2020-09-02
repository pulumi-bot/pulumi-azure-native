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

__all__ = ['JitNetworkAccessPolicy']


class JitNetworkAccessPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asc_location: Optional[pulumi.Input[str]] = None,
                 jit_network_access_policy_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 requests: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['JitNetworkAccessRequestArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 virtual_machines: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['JitNetworkAccessPolicyVirtualMachineArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a JitNetworkAccessPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asc_location: The location where ASC stores the data of the subscription. can be retrieved from Get locations
        :param pulumi.Input[str] jit_network_access_policy_name: Name of a Just-in-Time access configuration policy.
        :param pulumi.Input[str] kind: Kind of the resource
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['JitNetworkAccessPolicyVirtualMachineArgs']]]] virtual_machines: Configurations for Microsoft.Compute/virtualMachines resource type.
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

            if asc_location is None:
                raise TypeError("Missing required property 'asc_location'")
            __props__['asc_location'] = asc_location
            if jit_network_access_policy_name is None:
                raise TypeError("Missing required property 'jit_network_access_policy_name'")
            __props__['jit_network_access_policy_name'] = jit_network_access_policy_name
            __props__['kind'] = kind
            __props__['requests'] = requests
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if virtual_machines is None:
                raise TypeError("Missing required property 'virtual_machines'")
            __props__['virtual_machines'] = virtual_machines
            __props__['location'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:security/latest:JitNetworkAccessPolicy"), pulumi.Alias(type_="azurerm:security/v20200101:JitNetworkAccessPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(JitNetworkAccessPolicy, __self__).__init__(
            'azurerm:security/v20150601preview:JitNetworkAccessPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'JitNetworkAccessPolicy':
        """
        Get an existing JitNetworkAccessPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return JitNetworkAccessPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Kind of the resource
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Location where the resource is stored
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Gets the provisioning state of the Just-in-Time policy.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def requests(self) -> pulumi.Output[Optional[List['outputs.JitNetworkAccessRequestResponse']]]:
        return pulumi.get(self, "requests")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualMachines")
    def virtual_machines(self) -> pulumi.Output[List['outputs.JitNetworkAccessPolicyVirtualMachineResponse']]:
        """
        Configurations for Microsoft.Compute/virtualMachines resource type.
        """
        return pulumi.get(self, "virtual_machines")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

