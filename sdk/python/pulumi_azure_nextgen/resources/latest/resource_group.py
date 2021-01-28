# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['ResourceGroup']


class ResourceGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_by: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Resource group information.
        Latest API Version: 2020-10-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
        :param pulumi.Input[str] managed_by: The ID of the resource that manages this resource group.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to create or update. Can include alphanumeric, underscore, parentheses, hyphen, period (except at end), and Unicode characters that match the allowed characters.
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

            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['managed_by'] = managed_by
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['properties'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:resources/v20151101:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20160201:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20160701:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20160901:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20170510:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20180201:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20180501:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20190301:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20190501:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20190510:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20190701:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20190801:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20191001:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20200601:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20200801:ResourceGroup"), pulumi.Alias(type_="azure-nextgen:resources/v20201001:ResourceGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ResourceGroup, __self__).__init__(
            'azure-nextgen:resources/latest:ResourceGroup',
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
    def location(self) -> pulumi.Output[str]:
        """
        The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the resource that manages this resource group.
        """
        return pulumi.get(self, "managed_by")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.ResourceGroupPropertiesResponse']:
        """
        The resource group properties.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags attached to the resource group.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource group.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

