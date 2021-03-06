# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['ApiRelease']


class ApiRelease(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 release_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ApiRelease details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: Identifier of the API the release belongs to.
        :param pulumi.Input[str] notes: Release Notes
        :param pulumi.Input[str] release_id: Release identifier within an API. Must be unique in the current API Management service instance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
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

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__['api_id'] = api_id
            __props__['notes'] = notes
            __props__['release_id'] = release_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['created_date_time'] = None
            __props__['name'] = None
            __props__['type'] = None
            __props__['updated_date_time'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:apimanagement/v20200601preview:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement/latest:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement/latest:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement/v20170301:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20170301:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement/v20180101:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20180101:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement/v20180601preview:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20180601preview:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement/v20190101:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20190101:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement/v20191201:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement/v20191201preview:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201preview:ApiRelease"), pulumi.Alias(type_="azure-native:apimanagement/v20201201:ApiRelease"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20201201:ApiRelease")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApiRelease, __self__).__init__(
            'azure-native:apimanagement/v20200601preview:ApiRelease',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApiRelease':
        """
        Get an existing ApiRelease resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_id"] = None
        __props__["created_date_time"] = None
        __props__["name"] = None
        __props__["notes"] = None
        __props__["type"] = None
        __props__["updated_date_time"] = None
        return ApiRelease(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of the API the release belongs to.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="createdDateTime")
    def created_date_time(self) -> pulumi.Output[str]:
        """
        The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
        """
        return pulumi.get(self, "created_date_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        Release Notes
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedDateTime")
    def updated_date_time(self) -> pulumi.Output[str]:
        """
        The time the API release was updated.
        """
        return pulumi.get(self, "updated_date_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

