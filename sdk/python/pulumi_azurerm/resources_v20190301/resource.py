# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Resource(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the resource.
      * `principal_id` (`str`) - The principal ID of resource identity.
      * `tenant_id` (`str`) - The tenant ID of resource.
      * `type` (`str`) - The identity type.
      * `user_assigned_identities` (`dict`) - The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
    """
    kind: pulumi.Output[str]
    """
    The kind of the resource.
    """
    location: pulumi.Output[str]
    """
    Resource location
    """
    managed_by: pulumi.Output[str]
    """
    ID of the resource that manages this resource.
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    plan: pulumi.Output[dict]
    """
    The plan of the resource.
      * `name` (`str`) - The plan ID.
      * `product` (`str`) - The offer ID.
      * `promotion_code` (`str`) - The promotion code.
      * `publisher` (`str`) - The publisher ID.
      * `version` (`str`) - The plan's version.
    """
    properties: pulumi.Output[dict]
    """
    The resource properties.
    """
    sku: pulumi.Output[dict]
    """
    The SKU of the resource.
      * `capacity` (`float`) - The SKU capacity.
      * `family` (`str`) - The SKU family.
      * `model` (`str`) - The SKU model.
      * `name` (`str`) - The SKU name.
      * `size` (`str`) - The SKU size.
      * `tier` (`str`) - The SKU tier.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, identity=None, kind=None, location=None, managed_by=None, name=None, parent_resource_path=None, plan=None, properties=None, resource_group_name=None, resource_provider_namespace=None, resource_type=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Resource information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: The identity of the resource.
        :param pulumi.Input[str] kind: The kind of the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] managed_by: ID of the resource that manages this resource.
        :param pulumi.Input[str] name: The name of the resource to create.
        :param pulumi.Input[str] parent_resource_path: The parent resource identity.
        :param pulumi.Input[dict] plan: The plan of the resource.
        :param pulumi.Input[dict] properties: The resource properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group for the resource. The name is case insensitive.
        :param pulumi.Input[str] resource_provider_namespace: The namespace of the resource provider.
        :param pulumi.Input[str] resource_type: The resource type of the resource to create.
        :param pulumi.Input[dict] sku: The SKU of the resource.
        :param pulumi.Input[dict] tags: Resource tags

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.
          * `user_assigned_identities` (`pulumi.Input[dict]`) - The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        The **plan** object supports the following:

          * `name` (`pulumi.Input[str]`) - The plan ID.
          * `product` (`pulumi.Input[str]`) - The offer ID.
          * `promotion_code` (`pulumi.Input[str]`) - The promotion code.
          * `publisher` (`pulumi.Input[str]`) - The publisher ID.
          * `version` (`pulumi.Input[str]`) - The plan's version.

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - The SKU capacity.
          * `family` (`pulumi.Input[str]`) - The SKU family.
          * `model` (`pulumi.Input[str]`) - The SKU model.
          * `name` (`pulumi.Input[str]`) - The SKU name.
          * `size` (`pulumi.Input[str]`) - The SKU size.
          * `tier` (`pulumi.Input[str]`) - The SKU tier.
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

            __props__['identity'] = identity
            __props__['kind'] = kind
            __props__['location'] = location
            __props__['managed_by'] = managed_by
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if parent_resource_path is None:
                raise TypeError("Missing required property 'parent_resource_path'")
            __props__['parent_resource_path'] = parent_resource_path
            __props__['plan'] = plan
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_provider_namespace is None:
                raise TypeError("Missing required property 'resource_provider_namespace'")
            __props__['resource_provider_namespace'] = resource_provider_namespace
            if resource_type is None:
                raise TypeError("Missing required property 'resource_type'")
            __props__['resource_type'] = resource_type
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        super(Resource, __self__).__init__(
            'azurerm:resources/v20190301:Resource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Resource(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop