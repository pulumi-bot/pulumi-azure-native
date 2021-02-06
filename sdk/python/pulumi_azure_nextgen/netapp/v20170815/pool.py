# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['Pool']


class Pool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 pool_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_level: Optional[pulumi.Input[Union[str, 'ServiceLevel']]] = None,
                 size: Optional[pulumi.Input[float]] = None,
                 tags: Optional[Any] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Capacity pool resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] pool_name: The name of the capacity pool
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Union[str, 'ServiceLevel']] service_level: The service level of the file system
        :param pulumi.Input[float] size: Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
        :param Any tags: Resource tags
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

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['location'] = location
            if pool_name is None and not opts.urn:
                raise TypeError("Missing required property 'pool_name'")
            __props__['pool_name'] = pool_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_level is None:
                service_level = 'Premium'
            __props__['service_level'] = service_level
            if size is None:
                size = 4398046511104
            __props__['size'] = size
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['pool_id'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:netapp/latest:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20190501:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20190601:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20190701:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20190801:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20191001:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20191101:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20200201:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20200301:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20200501:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20200601:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20200701:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20200801:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20200901:Pool"), pulumi.Alias(type_="azure-nextgen:netapp/v20201101:Pool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Pool, __self__).__init__(
            'azure-nextgen:netapp/v20170815:Pool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Pool':
        """
        Get an existing Pool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Pool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
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
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Output[str]:
        """
        UUID v4 used to identify the Pool
        """
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Azure lifecycle management
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> pulumi.Output[Optional[str]]:
        """
        The service level of the file system
        """
        return pulumi.get(self, "service_level")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[Optional[float]]:
        """
        Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Any]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

