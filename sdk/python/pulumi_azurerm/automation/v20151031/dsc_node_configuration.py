# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class DscNodeConfiguration(pulumi.CustomResource):
    configuration: pulumi.Output[dict]
    """
    Gets or sets the configuration of the node.
      * `name` (`str`) - Gets or sets the name of the Dsc configuration.
    """
    creation_time: pulumi.Output[str]
    """
    Gets or sets creation time.
    """
    last_modified_time: pulumi.Output[str]
    """
    Gets or sets the last modified time.
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, automation_account_name=None, configuration=None, increment_node_configuration_build=None, name=None, resource_group_name=None, source=None, __props__=None, __name__=None, __opts__=None):
        """
        Definition of the dsc node configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account.
        :param pulumi.Input[dict] configuration: Gets or sets the configuration of the node.
        :param pulumi.Input[bool] increment_node_configuration_build: If a new build version of NodeConfiguration is required.
        :param pulumi.Input[str] name: The create or update parameters for configuration.
        :param pulumi.Input[str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[dict] source: Gets or sets the source.

        The **configuration** object supports the following:

          * `name` (`pulumi.Input[str]`) - Gets or sets the name of the Dsc configuration.

        The **source** object supports the following:

          * `hash` (`pulumi.Input[dict]`) - Gets or sets the hash.
            * `algorithm` (`pulumi.Input[str]`) - Gets or sets the content hash algorithm used to hash the content.
            * `value` (`pulumi.Input[str]`) - Gets or sets expected hash value of the content.

          * `type` (`pulumi.Input[str]`) - Gets or sets the content source type.
          * `value` (`pulumi.Input[str]`) - Gets or sets the value of the content. This is based on the content source type.
          * `version` (`pulumi.Input[str]`) - Gets or sets the version of the content.
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

            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            if configuration is None:
                raise TypeError("Missing required property 'configuration'")
            __props__['configuration'] = configuration
            __props__['increment_node_configuration_build'] = increment_node_configuration_build
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source is None:
                raise TypeError("Missing required property 'source'")
            __props__['source'] = source
            __props__['creation_time'] = None
            __props__['last_modified_time'] = None
            __props__['type'] = None
        super(DscNodeConfiguration, __self__).__init__(
            'azurerm:automation/v20151031:DscNodeConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing DscNodeConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DscNodeConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop