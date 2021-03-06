# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['Policy']


class Policy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_format: Optional[pulumi.Input[Union[str, 'PolicyContentFormat']]] = None,
                 policy_content: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Policy Contract details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[str, 'PolicyContentFormat']] content_format: Format of the policyContent.
        :param pulumi.Input[str] policy_content: Json escaped Xml Encoded contents of the Policy.
        :param pulumi.Input[str] policy_id: The identifier of the Policy.
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

            if content_format is None:
                content_format = 'xml'
            __props__['content_format'] = content_format
            if policy_content is None and not opts.urn:
                raise TypeError("Missing required property 'policy_content'")
            __props__['policy_content'] = policy_content
            __props__['policy_id'] = policy_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:apimanagement/v20180601preview:Policy"), pulumi.Alias(type_="azure-native:apimanagement:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement:Policy"), pulumi.Alias(type_="azure-native:apimanagement/latest:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement/latest:Policy"), pulumi.Alias(type_="azure-native:apimanagement/v20170301:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20170301:Policy"), pulumi.Alias(type_="azure-native:apimanagement/v20180101:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20180101:Policy"), pulumi.Alias(type_="azure-native:apimanagement/v20190101:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20190101:Policy"), pulumi.Alias(type_="azure-native:apimanagement/v20191201:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201:Policy"), pulumi.Alias(type_="azure-native:apimanagement/v20191201preview:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201preview:Policy"), pulumi.Alias(type_="azure-native:apimanagement/v20200601preview:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20200601preview:Policy"), pulumi.Alias(type_="azure-native:apimanagement/v20201201:Policy"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20201201:Policy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Policy, __self__).__init__(
            'azure-native:apimanagement/v20180601preview:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["content_format"] = None
        __props__["name"] = None
        __props__["policy_content"] = None
        __props__["type"] = None
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentFormat")
    def content_format(self) -> pulumi.Output[Optional[str]]:
        """
        Format of the policyContent.
        """
        return pulumi.get(self, "content_format")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyContent")
    def policy_content(self) -> pulumi.Output[str]:
        """
        Json escaped Xml Encoded contents of the Policy.
        """
        return pulumi.get(self, "policy_content")

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

