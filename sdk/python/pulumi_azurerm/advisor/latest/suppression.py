# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['Suppression']


class Suppression(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recommendation_id: Optional[pulumi.Input[str]] = None,
                 resource_uri: Optional[pulumi.Input[str]] = None,
                 suppression_id: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.

        ## Example Usage
        ### CreateSuppression

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        suppression = azurerm.advisor.latest.Suppression("suppression",
            name="suppressionName1",
            recommendation_id="recommendationId",
            resource_uri="resourceUri",
            ttl="07:00:00:00")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the suppression.
        :param pulumi.Input[str] recommendation_id: The recommendation ID.
        :param pulumi.Input[str] resource_uri: The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
        :param pulumi.Input[str] suppression_id: The GUID of the suppression.
        :param pulumi.Input[str] ttl: The duration for which the suppression is valid.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if recommendation_id is None:
                raise TypeError("Missing required property 'recommendation_id'")
            __props__['recommendation_id'] = recommendation_id
            if resource_uri is None:
                raise TypeError("Missing required property 'resource_uri'")
            __props__['resource_uri'] = resource_uri
            __props__['suppression_id'] = suppression_id
            __props__['ttl'] = ttl
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:advisor/v20160712preview:Suppression"), pulumi.Alias(type_="azurerm:advisor/v20170331:Suppression"), pulumi.Alias(type_="azurerm:advisor/v20170419:Suppression"), pulumi.Alias(type_="azurerm:advisor/v20200101:Suppression")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Suppression, __self__).__init__(
            'azurerm:advisor/latest:Suppression',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Suppression':
        """
        Get an existing Suppression resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Suppression(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="suppressionId")
    def suppression_id(self) -> pulumi.Output[Optional[str]]:
        """
        The GUID of the suppression.
        """
        return pulumi.get(self, "suppression_id")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The duration for which the suppression is valid.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

