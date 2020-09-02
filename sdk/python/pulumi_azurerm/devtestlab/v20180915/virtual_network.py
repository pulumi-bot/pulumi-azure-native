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

__all__ = ['VirtualNetwork']


class VirtualNetwork(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_subnets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SubnetArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_provider_resource_id: Optional[pulumi.Input[str]] = None,
                 lab_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subnet_overrides: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SubnetOverrideArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A virtual network.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SubnetArgs']]]] allowed_subnets: The allowed subnets of the virtual network.
        :param pulumi.Input[str] description: The description of the virtual network.
        :param pulumi.Input[str] external_provider_resource_id: The Microsoft.Network resource identifier of the virtual network.
        :param pulumi.Input[str] lab_name: The name of the lab.
        :param pulumi.Input[str] location: The location of the resource.
        :param pulumi.Input[str] name: The name of the virtual network.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SubnetOverrideArgs']]]] subnet_overrides: The subnet overrides of the virtual network.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
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

            __props__['allowed_subnets'] = allowed_subnets
            __props__['description'] = description
            __props__['external_provider_resource_id'] = external_provider_resource_id
            if lab_name is None:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['subnet_overrides'] = subnet_overrides
            __props__['tags'] = tags
            __props__['created_date'] = None
            __props__['external_subnets'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
            __props__['unique_identifier'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:devtestlab/latest:VirtualNetwork"), pulumi.Alias(type_="azurerm:devtestlab/v20150521preview:VirtualNetwork"), pulumi.Alias(type_="azurerm:devtestlab/v20160515:VirtualNetwork")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualNetwork, __self__).__init__(
            'azurerm:devtestlab/v20180915:VirtualNetwork',
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
    @pulumi.getter(name="allowedSubnets")
    def allowed_subnets(self) -> pulumi.Output[Optional[List['outputs.SubnetResponse']]]:
        """
        The allowed subnets of the virtual network.
        """
        return pulumi.get(self, "allowed_subnets")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> pulumi.Output[str]:
        """
        The creation date of the virtual network.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the virtual network.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalProviderResourceId")
    def external_provider_resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Microsoft.Network resource identifier of the virtual network.
        """
        return pulumi.get(self, "external_provider_resource_id")

    @property
    @pulumi.getter(name="externalSubnets")
    def external_subnets(self) -> pulumi.Output[List['outputs.ExternalSubnetResponse']]:
        """
        The external subnet properties.
        """
        return pulumi.get(self, "external_subnets")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="subnetOverrides")
    def subnet_overrides(self) -> pulumi.Output[Optional[List['outputs.SubnetOverrideResponse']]]:
        """
        The subnet overrides of the virtual network.
        """
        return pulumi.get(self, "subnet_overrides")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> pulumi.Output[str]:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

