# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AccessPolicy(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, environment_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        An access policy is used to grant users and applications access to the environment. Roles are assigned to service principals in Azure Active Directory. These roles define the actions the principal can perform through the Time Series Insights data plane APIs.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] environment_name: The name of the Time Series Insights environment associated with the specified resource group.
        :param pulumi.Input[str] name: Name of the access policy.
        :param pulumi.Input[str] resource_group_name: Name of an Azure Resource group.

        The **properties** object supports the following:

          * `description` (`pulumi.Input[str]`) - An description of the access policy.
          * `principal_object_id` (`pulumi.Input[str]`) - The objectId of the principal in Azure Active Directory.
          * `roles` (`pulumi.Input[list]`) - The list of roles the principal is assigned on the environment.
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

            if environment_name is None:
                raise TypeError("Missing required property 'environment_name'")
            __props__['environment_name'] = environment_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(AccessPolicy, __self__).__init__(
            'azurerm:timeseriesinsights/v20171115:AccessPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AccessPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AccessPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
