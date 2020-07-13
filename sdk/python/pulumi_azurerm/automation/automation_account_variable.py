# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class AutomationAccountVariable(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Gets or sets the properties of the variable.
      * `creation_time` (`str`) - Gets or sets the creation time.
      * `description` (`str`) - Gets or sets the description.
      * `is_encrypted` (`bool`) - Gets or sets the encrypted flag of the variable.
      * `last_modified_time` (`str`) - Gets or sets the last modified time.
      * `value` (`str`) - Gets or sets the value of the variable.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, automation_account_name=None, name=None, properties=None, resource_group_name=None, variable_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Definition of the variable.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account.
        :param pulumi.Input[str] name: Gets or sets the name of the variable.
        :param pulumi.Input[dict] properties: Gets or sets the properties of the variable.
        :param pulumi.Input[str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[str] variable_name: The variable name.

        The **properties** object supports the following:

          * `description` (`pulumi.Input[str]`) - Gets or sets the description of the variable.
          * `is_encrypted` (`pulumi.Input[bool]`) - Gets or sets the encrypted flag of the variable.
          * `value` (`pulumi.Input[str]`) - Gets or sets the value of the variable.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if variable_name is None:
                raise TypeError("Missing required property 'variable_name'")
            __props__['variable_name'] = variable_name
            __props__['type'] = None
        super(AutomationAccountVariable, __self__).__init__(
            'azurerm:automation:AutomationAccountVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AutomationAccountVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AutomationAccountVariable(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
