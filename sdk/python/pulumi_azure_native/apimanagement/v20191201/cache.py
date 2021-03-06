# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['Cache']


class Cache(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_id: Optional[pulumi.Input[str]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Cache details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cache_id: Identifier of the Cache entity. Cache identifier (should be either 'default' or valid Azure region identifier).
        :param pulumi.Input[str] connection_string: Runtime connection string to cache
        :param pulumi.Input[str] description: Cache description
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_id: Original uri of entity in external system cache points to
        :param pulumi.Input[str] service_name: The name of the API Management service.
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

            __props__['cache_id'] = cache_id
            if connection_string is None and not opts.urn:
                raise TypeError("Missing required property 'connection_string'")
            __props__['connection_string'] = connection_string
            __props__['description'] = description
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_id'] = resource_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201:Cache"), pulumi.Alias(type_="azure-native:apimanagement:Cache"), pulumi.Alias(type_="azure-nextgen:apimanagement:Cache"), pulumi.Alias(type_="azure-native:apimanagement/latest:Cache"), pulumi.Alias(type_="azure-nextgen:apimanagement/latest:Cache"), pulumi.Alias(type_="azure-native:apimanagement/v20180601preview:Cache"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20180601preview:Cache"), pulumi.Alias(type_="azure-native:apimanagement/v20190101:Cache"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20190101:Cache"), pulumi.Alias(type_="azure-native:apimanagement/v20191201preview:Cache"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201preview:Cache"), pulumi.Alias(type_="azure-native:apimanagement/v20200601preview:Cache"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20200601preview:Cache"), pulumi.Alias(type_="azure-native:apimanagement/v20201201:Cache"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20201201:Cache")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Cache, __self__).__init__(
            'azure-native:apimanagement/v20191201:Cache',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Cache':
        """
        Get an existing Cache resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["connection_string"] = None
        __props__["description"] = None
        __props__["name"] = None
        __props__["resource_id"] = None
        __props__["type"] = None
        return Cache(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        Runtime connection string to cache
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Cache description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        Original uri of entity in external system cache points to
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

