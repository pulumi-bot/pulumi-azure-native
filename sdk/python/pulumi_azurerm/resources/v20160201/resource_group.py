# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['ResourceGroup']


class ResourceGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Resource group information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the resource group. It cannot be changed after the resource group has been created. Has to be one of the supported Azure Locations, such as West US, East US, West Europe, East Asia, etc.
        :param pulumi.Input[str] name: The name of the resource group to be created or updated.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags attached to the resource group.
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['properties'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:resources/v20151101:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20160701:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20160901:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20170510:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20180201:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20180501:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20190301:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20190501:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20190510:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20190701:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20190801:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20191001:ResourceGroup"), pulumi.Alias(type_="azurerm:resources/v20200601:ResourceGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ResourceGroup, __self__).__init__(
            'azurerm:resources/v20160201:ResourceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ResourceGroup':
        """
        Get an existing ResourceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ResourceGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the resource group. It cannot be changed after the resource group has been created. Has to be one of the supported Azure Locations, such as West US, East US, West Europe, East Asia, etc.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The Name of the resource group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.ResourceGroupPropertiesResponse':
        """
        The resource group properties.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags attached to the resource group.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

