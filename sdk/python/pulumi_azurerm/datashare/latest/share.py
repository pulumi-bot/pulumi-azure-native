# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['Share']


class Share(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 share_kind: Optional[pulumi.Input[str]] = None,
                 share_name: Optional[pulumi.Input[str]] = None,
                 terms: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A share data transfer object.

        ## Example Usage
        ### Shares_Create

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        share = azurerm.datashare.latest.Share("share",
            account_name="Account1",
            description="share description",
            resource_group_name="SampleResourceGroup",
            share_kind="CopyBased",
            share_name="Share1",
            terms="Confidential")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the share account.
        :param pulumi.Input[str] description: Share description.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[str] share_kind: Share kind.
        :param pulumi.Input[str] share_name: The name of the share.
        :param pulumi.Input[str] terms: Share terms.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['description'] = description
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['share_kind'] = share_kind
            if share_name is None:
                raise TypeError("Missing required property 'share_name'")
            __props__['share_name'] = share_name
            __props__['terms'] = terms
            __props__['created_at'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
            __props__['user_email'] = None
            __props__['user_name'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:datashare/v20181101preview:Share"), pulumi.Alias(type_="azurerm:datashare/v20191101:Share")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Share, __self__).__init__(
            'azurerm:datashare/latest:Share',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Share':
        """
        Get an existing Share resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Share(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Time at which the share was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Share description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the azure resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Gets or sets the provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="shareKind")
    def share_kind(self) -> pulumi.Output[Optional[str]]:
        """
        Share kind.
        """
        return pulumi.get(self, "share_kind")

    @property
    @pulumi.getter
    def terms(self) -> pulumi.Output[Optional[str]]:
        """
        Share terms.
        """
        return pulumi.get(self, "terms")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the azure resource
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userEmail")
    def user_email(self) -> pulumi.Output[str]:
        """
        Email of the user who created the resource
        """
        return pulumi.get(self, "user_email")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Name of the user who created the resource
        """
        return pulumi.get(self, "user_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

