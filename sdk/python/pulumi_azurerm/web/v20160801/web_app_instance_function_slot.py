# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WebAppInstanceFunctionSlot(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    FunctionEnvelope resource specific properties
      * `config` (`dict`) - Config information.
      * `config_href` (`str`) - Config URI.
      * `files` (`dict`) - File list.
      * `function_app_id` (`str`) - Function App ID.
      * `href` (`str`) - Function URI.
      * `name` (`str`) - Function name.
      * `script_href` (`str`) - Script URI.
      * `script_root_path_href` (`str`) - Script root path URI.
      * `secrets_file_href` (`str`) - Secrets file URI.
      * `test_data` (`str`) - Test data used when testing via the Azure Portal.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, kind=None, name=None, properties=None, resource_group_name=None, slot=None, __props__=None, __name__=None, __opts__=None):
        """
        Web Job Information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] name: Function name.
        :param pulumi.Input[dict] properties: FunctionEnvelope resource specific properties
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[str] slot: Name of the deployment slot. If a slot is not specified, the API deletes a deployment for the production slot.

        The **properties** object supports the following:

          * `config` (`pulumi.Input[dict]`) - Config information.
          * `config_href` (`pulumi.Input[str]`) - Config URI.
          * `files` (`pulumi.Input[dict]`) - File list.
          * `href` (`pulumi.Input[str]`) - Function URI.
          * `script_href` (`pulumi.Input[str]`) - Script URI.
          * `script_root_path_href` (`pulumi.Input[str]`) - Script root path URI.
          * `secrets_file_href` (`pulumi.Input[str]`) - Secrets file URI.
          * `test_data` (`pulumi.Input[str]`) - Test data used when testing via the Azure Portal.
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

            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if slot is None:
                raise TypeError("Missing required property 'slot'")
            __props__['slot'] = slot
            __props__['type'] = None
        super(WebAppInstanceFunctionSlot, __self__).__init__(
            'azurerm:web/v20160801:WebAppInstanceFunctionSlot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WebAppInstanceFunctionSlot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppInstanceFunctionSlot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
