# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['ApiIssueComment']


class ApiIssueComment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 comment_id: Optional[pulumi.Input[str]] = None,
                 created_date: Optional[pulumi.Input[str]] = None,
                 issue_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 text: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Issue Comment Contract details.

        ## ApiManagementCreateApiIssueComment

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        api_issue_comment = azurerm.apimanagement.latest.ApiIssueComment("apiIssueComment",
            api_id="57d1f7558aa04f15146d9d8a",
            comment_id="599e29ab193c3c0bd0b3e2fb",
            created_date="2018-02-01T22:21:20.467Z",
            issue_id="57d2ef278aa04f0ad01d6cdc",
            resource_group_name="rg1",
            service_name="apimService1",
            text="Issue comment.",
            user_id="/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: API identifier. Must be unique in the current API Management service instance.
        :param pulumi.Input[str] comment_id: Comment identifier within an Issue. Must be unique in the current Issue.
        :param pulumi.Input[str] created_date: Date and time when the comment was created.
        :param pulumi.Input[str] issue_id: Issue identifier. Must be unique in the current API Management service instance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.
        :param pulumi.Input[str] text: Comment text.
        :param pulumi.Input[str] user_id: A resource identifier for the user who left the comment.
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
            if comment_id is None:
                raise TypeError("Missing required property 'comment_id'")
            __props__['comment_id'] = comment_id
            __props__['created_date'] = created_date
            if issue_id is None:
                raise TypeError("Missing required property 'issue_id'")
            __props__['issue_id'] = issue_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            if text is None:
                raise TypeError("Missing required property 'text'")
            __props__['text'] = text
            if user_id is None:
                raise TypeError("Missing required property 'user_id'")
            __props__['user_id'] = user_id
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:apimanagement/v20170301:ApiIssueComment"), pulumi.Alias(type_="azurerm:apimanagement/v20180101:ApiIssueComment"), pulumi.Alias(type_="azurerm:apimanagement/v20190101:ApiIssueComment"), pulumi.Alias(type_="azurerm:apimanagement/v20191201:ApiIssueComment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApiIssueComment, __self__).__init__(
            'azurerm:apimanagement/latest:ApiIssueComment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApiIssueComment':
        """
        Get an existing ApiIssueComment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ApiIssueComment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> pulumi.Output[Optional[str]]:
        """
        Date and time when the comment was created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def text(self) -> pulumi.Output[str]:
        """
        Comment text.
        """
        return pulumi.get(self, "text")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        A resource identifier for the user who left the comment.
        """
        return pulumi.get(self, "user_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
