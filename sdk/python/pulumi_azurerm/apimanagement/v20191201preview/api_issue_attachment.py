# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['ApiIssueAttachment']


class ApiIssueAttachment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 attachment_id: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_format: Optional[pulumi.Input[str]] = None,
                 issue_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Issue Attachment Contract details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: API identifier. Must be unique in the current API Management service instance.
        :param pulumi.Input[str] attachment_id: Attachment identifier within an Issue. Must be unique in the current Issue.
        :param pulumi.Input[str] content: An HTTP link or Base64-encoded binary data.
        :param pulumi.Input[str] content_format: Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
        :param pulumi.Input[str] issue_id: Issue identifier. Must be unique in the current API Management service instance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.
        :param pulumi.Input[str] title: Filename by which the binary data will be saved.
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

            if api_id is None:
                raise TypeError("Missing required property 'api_id'")
            __props__['api_id'] = api_id
            if attachment_id is None:
                raise TypeError("Missing required property 'attachment_id'")
            __props__['attachment_id'] = attachment_id
            if content is None:
                raise TypeError("Missing required property 'content'")
            __props__['content'] = content
            if content_format is None:
                raise TypeError("Missing required property 'content_format'")
            __props__['content_format'] = content_format
            if issue_id is None:
                raise TypeError("Missing required property 'issue_id'")
            __props__['issue_id'] = issue_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:apimanagement/latest:ApiIssueAttachment"), pulumi.Alias(type_="azurerm:apimanagement/v20170301:ApiIssueAttachment"), pulumi.Alias(type_="azurerm:apimanagement/v20180101:ApiIssueAttachment"), pulumi.Alias(type_="azurerm:apimanagement/v20180601preview:ApiIssueAttachment"), pulumi.Alias(type_="azurerm:apimanagement/v20190101:ApiIssueAttachment"), pulumi.Alias(type_="azurerm:apimanagement/v20191201:ApiIssueAttachment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApiIssueAttachment, __self__).__init__(
            'azurerm:apimanagement/v20191201preview:ApiIssueAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApiIssueAttachment':
        """
        Get an existing ApiIssueAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ApiIssueAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        An HTTP link or Base64-encoded binary data.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentFormat")
    def content_format(self) -> pulumi.Output[str]:
        """
        Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
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
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        Filename by which the binary data will be saved.
        """
        return pulumi.get(self, "title")

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

