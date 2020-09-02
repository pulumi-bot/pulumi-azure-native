# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['Product']


class Product(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approval_required: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 subscription_required: Optional[pulumi.Input[bool]] = None,
                 subscriptions_limit: Optional[pulumi.Input[float]] = None,
                 terms: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Product details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] approval_required: whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of true.
        :param pulumi.Input[str] description: Product description. May include HTML formatting tags.
        :param pulumi.Input[str] display_name: Product name.
        :param pulumi.Input[str] product_id: Product identifier. Must be unique in the current API Management service instance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.
        :param pulumi.Input[str] state: whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
        :param pulumi.Input[bool] subscription_required: Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
        :param pulumi.Input[float] subscriptions_limit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of true.
        :param pulumi.Input[str] terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
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

            __props__['approval_required'] = approval_required
            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if product_id is None:
                raise TypeError("Missing required property 'product_id'")
            __props__['product_id'] = product_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['state'] = state
            __props__['subscription_required'] = subscription_required
            __props__['subscriptions_limit'] = subscriptions_limit
            __props__['terms'] = terms
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:apimanagement/latest:Product"), pulumi.Alias(type_="azurerm:apimanagement/v20160707:Product"), pulumi.Alias(type_="azurerm:apimanagement/v20161010:Product"), pulumi.Alias(type_="azurerm:apimanagement/v20170301:Product"), pulumi.Alias(type_="azurerm:apimanagement/v20180101:Product"), pulumi.Alias(type_="azurerm:apimanagement/v20180601preview:Product"), pulumi.Alias(type_="azurerm:apimanagement/v20190101:Product"), pulumi.Alias(type_="azurerm:apimanagement/v20191201preview:Product")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Product, __self__).__init__(
            'azurerm:apimanagement/v20191201:Product',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Product':
        """
        Get an existing Product resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Product(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="approvalRequired")
    def approval_required(self) -> pulumi.Output[Optional[bool]]:
        """
        whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of true.
        """
        return pulumi.get(self, "approval_required")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Product description. May include HTML formatting tags.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Product name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subscriptionRequired")
    def subscription_required(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
        """
        return pulumi.get(self, "subscription_required")

    @property
    @pulumi.getter(name="subscriptionsLimit")
    def subscriptions_limit(self) -> pulumi.Output[Optional[float]]:
        """
        Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of true.
        """
        return pulumi.get(self, "subscriptions_limit")

    @property
    @pulumi.getter
    def terms(self) -> pulumi.Output[Optional[str]]:
        """
        Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
        """
        return pulumi.get(self, "terms")

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

