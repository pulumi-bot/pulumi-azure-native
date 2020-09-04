# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['PrivateLinkForAzureAd']


class PrivateLinkForAzureAd(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_tenants: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_tenant_id: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 resource_group: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tenants: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        PrivateLink Policy configuration object.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_tenants: Flag indicating whether all tenants are allowed
        :param pulumi.Input[str] name: Name of this resource.
        :param pulumi.Input[str] owner_tenant_id: Guid of the owner tenant
        :param pulumi.Input[str] policy_name: The name of the private link policy in Azure AD.
        :param pulumi.Input[str] resource_group: Name of the resource group
        :param pulumi.Input[str] resource_group_name: Name of an Azure resource group.
        :param pulumi.Input[str] resource_name_: Name of the private link policy resource
        :param pulumi.Input[str] subscription_id: Subscription Identifier
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[List[pulumi.Input[str]]] tenants: The list of tenantIds.
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

            __props__['all_tenants'] = all_tenants
            __props__['name'] = name
            __props__['owner_tenant_id'] = owner_tenant_id
            if policy_name is None:
                raise TypeError("Missing required property 'policy_name'")
            __props__['policy_name'] = policy_name
            __props__['resource_group'] = resource_group
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_name'] = resource_name_
            __props__['subscription_id'] = subscription_id
            __props__['tags'] = tags
            __props__['tenants'] = tenants
            __props__['type'] = None
        super(PrivateLinkForAzureAd, __self__).__init__(
            'azurerm:aadiam/v20200301preview:privateLinkForAzureAd',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateLinkForAzureAd':
        """
        Get an existing PrivateLinkForAzureAd resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PrivateLinkForAzureAd(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allTenants")
    def all_tenants(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag indicating whether all tenants are allowed
        """
        return pulumi.get(self, "all_tenants")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerTenantId")
    def owner_tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        Guid of the owner tenant
        """
        return pulumi.get(self, "owner_tenant_id")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the resource group
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the private link policy resource
        """
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Output[Optional[str]]:
        """
        Subscription Identifier
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tenants(self) -> pulumi.Output[Optional[List[str]]]:
        """
        The list of tenantIds.
        """
        return pulumi.get(self, "tenants")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of this resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

