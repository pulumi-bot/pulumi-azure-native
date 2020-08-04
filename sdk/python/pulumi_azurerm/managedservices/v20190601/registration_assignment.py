# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class RegistrationAssignment(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Name of the registration assignment.
    """
    properties: pulumi.Output[dict]
    """
    Properties of a registration assignment.
      * `provisioning_state` (`str`) - Current state of the registration assignment.
      * `registration_definition` (`dict`) - Registration definition inside registration assignment.
        * `id` (`str`) - Fully qualified path of the registration definition.
        * `name` (`str`) - Name of the registration definition.
        * `plan` (`dict`) - Plan details for the managed services.
          * `name` (`str`) - The plan name.
          * `product` (`str`) - The product code.
          * `publisher` (`str`) - The publisher ID.
          * `version` (`str`) - The plan's version.

        * `properties` (`dict`) - Properties of registration definition inside registration assignment.
        * `type` (`str`) - Type of the resource (Microsoft.ManagedServices/registrationDefinitions).

      * `registration_definition_id` (`str`) - Fully qualified path of the registration definition.
    """
    type: pulumi.Output[str]
    """
    Type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, scope=None, __props__=None, __name__=None, __opts__=None):
        """
        Registration assignment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Guid of the registration assignment.
        :param pulumi.Input[dict] properties: Properties of a registration assignment.
        :param pulumi.Input[str] scope: Scope of the resource.

        The **properties** object supports the following:

          * `registration_definition_id` (`pulumi.Input[str]`) - Fully qualified path of the registration definition.
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
            __props__['properties'] = properties
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['type'] = None
        super(RegistrationAssignment, __self__).__init__(
            'azurerm:managedservices/v20190601:RegistrationAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RegistrationAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RegistrationAssignment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
