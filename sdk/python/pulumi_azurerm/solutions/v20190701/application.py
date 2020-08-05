# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Application(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the resource.
      * `principal_id` (`str`) - The principal ID of resource identity.
      * `tenant_id` (`str`) - The tenant ID of resource.
      * `type` (`str`) - The identity type.
      * `user_assigned_identities` (`dict`) - The list of user identities associated with the resource. The user identity dictionary key references will be resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
    """
    kind: pulumi.Output[str]
    """
    The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
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
    The plan information.
      * `name` (`str`) - The plan name.
      * `product` (`str`) - The product code.
      * `promotion_code` (`str`) - The promotion code.
      * `publisher` (`str`) - The publisher ID.
      * `version` (`str`) - The plan's version.
    """
    properties: pulumi.Output[dict]
    """
    The managed application properties.
      * `application_definition_id` (`str`) - The fully qualified path of managed application definition Id.
      * `artifacts` (`list`) - The collection of managed application artifacts.
        * `name` (`str`) - The managed application artifact name.
        * `type` (`str`) - The managed application artifact type.
        * `uri` (`str`) - The managed application artifact blob uri.

      * `authorizations` (`list`) - The  read-only authorizations property that is retrieved from the application package.
        * `principal_id` (`str`) - The provider's principal identifier. This is the identity that the provider will use to call ARM to manage the managed application resources.
        * `role_definition_id` (`str`) - The provider's role definition identifier. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.

      * `billing_details` (`dict`) - The managed application billing details.
        * `resource_usage_id` (`str`) - The managed application resource usage Id.

      * `created_by` (`dict`) - The client entity that created the JIT request.
        * `application_id` (`str`) - The client application Id.
        * `oid` (`str`) - The client Oid.
        * `puid` (`str`) - The client Puid

      * `customer_support` (`dict`) - The read-only customer support property that is retrieved from the application package.
        * `contact_name` (`str`) - The contact name.
        * `email` (`str`) - The contact email.
        * `phone` (`str`) - The contact phone number.

      * `jit_access_policy` (`dict`) - The managed application Jit access policy.
        * `jit_access_enabled` (`bool`) - Whether the JIT access is enabled.
        * `jit_approval_mode` (`str`) - JIT approval mode.
        * `jit_approvers` (`list`) - The JIT approvers
          * `display_name` (`str`) - The approver display name.
          * `id` (`str`) - The approver service principal Id.
          * `type` (`str`) - The approver type.

        * `maximum_jit_access_duration` (`str`) - The maximum duration JIT access is granted. This is an ISO8601 time period value.

      * `managed_resource_group_id` (`str`) - The managed resource group Id.
      * `management_mode` (`str`) - The managed application management mode.
      * `outputs` (`dict`) - Name and value pairs that define the managed application outputs.
      * `parameters` (`dict`) - Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
      * `provisioning_state` (`str`) - The managed application provisioning state.
      * `publisher_tenant_id` (`str`) - The publisher tenant Id.
      * `support_urls` (`dict`) - The read-only support URLs property that is retrieved from the application package.
        * `government_cloud` (`str`) - The government cloud support URL.
        * `public_azure` (`str`) - The public azure support URL.

      * `updated_by` (`dict`) - The client entity that last updated the JIT request.
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
    def __init__(__self__, resource_name, opts=None, application_definition_id=None, identity=None, jit_access_policy=None, kind=None, location=None, managed_by=None, managed_resource_group_id=None, name=None, parameters=None, plan=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Information about managed application.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_definition_id: The fully qualified path of managed application definition Id.
        :param pulumi.Input[dict] identity: The identity of the resource.
        :param pulumi.Input[dict] jit_access_policy: The managed application Jit access policy.
        :param pulumi.Input[str] kind: The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] managed_by: ID of the resource that manages this resource.
        :param pulumi.Input[str] managed_resource_group_id: The managed resource group Id.
        :param pulumi.Input[str] name: The name of the managed application.
        :param pulumi.Input[dict] parameters: Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
        :param pulumi.Input[dict] plan: The plan information.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[dict] sku: The SKU of the resource.
        :param pulumi.Input[dict] tags: Resource tags

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.
          * `user_assigned_identities` (`pulumi.Input[dict]`) - The list of user identities associated with the resource. The user identity dictionary key references will be resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        The **jit_access_policy** object supports the following:

          * `jit_access_enabled` (`pulumi.Input[bool]`) - Whether the JIT access is enabled.
          * `jit_approval_mode` (`pulumi.Input[str]`) - JIT approval mode.
          * `jit_approvers` (`pulumi.Input[list]`) - The JIT approvers
            * `display_name` (`pulumi.Input[str]`) - The approver display name.
            * `id` (`pulumi.Input[str]`) - The approver service principal Id.
            * `type` (`pulumi.Input[str]`) - The approver type.

          * `maximum_jit_access_duration` (`pulumi.Input[str]`) - The maximum duration JIT access is granted. This is an ISO8601 time period value.

        The **plan** object supports the following:

          * `name` (`pulumi.Input[str]`) - The plan name.
          * `product` (`pulumi.Input[str]`) - The product code.
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

            __props__['application_definition_id'] = application_definition_id
            __props__['identity'] = identity
            __props__['jit_access_policy'] = jit_access_policy
            if kind is None:
                raise TypeError("Missing required property 'kind'")
            __props__['kind'] = kind
            __props__['location'] = location
            __props__['managed_by'] = managed_by
            __props__['managed_resource_group_id'] = managed_resource_group_id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['plan'] = plan
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(Application, __self__).__init__(
            'azurerm:solutions/v20190701:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Application(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
