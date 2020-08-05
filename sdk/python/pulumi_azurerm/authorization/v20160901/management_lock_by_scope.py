# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ManagementLockByScope(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the lock.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the lock.
      * `level` (`str`) - The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
      * `notes` (`str`) - Notes about the lock. Maximum of 512 characters.
      * `owners` (`list`) - The owners of the lock.
        * `application_id` (`str`) - The application ID of the lock owner.
    """
    type: pulumi.Output[str]
    """
    The resource type of the lock - Microsoft.Authorization/locks.
    """
    def __init__(__self__, resource_name, opts=None, level=None, name=None, notes=None, owners=None, scope=None, __props__=None, __name__=None, __opts__=None):
        """
        The lock information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] level: The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
        :param pulumi.Input[str] name: The name of lock.
        :param pulumi.Input[str] notes: Notes about the lock. Maximum of 512 characters.
        :param pulumi.Input[list] owners: The owners of the lock.
        :param pulumi.Input[str] scope: The scope for the lock. When providing a scope for the assignment, use '/subscriptions/{subscriptionId}' for subscriptions, '/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}' for resource groups, and '/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePathIfPresent}/{resourceType}/{resourceName}' for resources.

        The **owners** object supports the following:

          * `application_id` (`pulumi.Input[str]`) - The application ID of the lock owner.
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

            if level is None:
                raise TypeError("Missing required property 'level'")
            __props__['level'] = level
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['notes'] = notes
            __props__['owners'] = owners
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['properties'] = None
            __props__['type'] = None
        super(ManagementLockByScope, __self__).__init__(
            'azurerm:authorization/v20160901:ManagementLockByScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ManagementLockByScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagementLockByScope(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
