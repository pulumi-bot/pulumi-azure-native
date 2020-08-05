# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WebAppFunction(pulumi.CustomResource):
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
      * `invoke_url_template` (`str`) - The invocation URL
      * `is_disabled` (`bool`) - Value indicating whether the function is disabled
      * `language` (`str`) - The function language
      * `script_href` (`str`) - Script URI.
      * `script_root_path_href` (`str`) - Script root path URI.
      * `secrets_file_href` (`str`) - Secrets file URI.
      * `test_data` (`str`) - Test data used when testing via the Azure Portal.
      * `test_data_href` (`str`) - Test data URI.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, config=None, config_href=None, files=None, function_app_id=None, href=None, invoke_url_template=None, is_disabled=None, kind=None, language=None, name=None, resource_group_name=None, script_href=None, script_root_path_href=None, secrets_file_href=None, test_data=None, test_data_href=None, __props__=None, __name__=None, __opts__=None):
        """
        Function information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] config: Config information.
        :param pulumi.Input[str] config_href: Config URI.
        :param pulumi.Input[dict] files: File list.
        :param pulumi.Input[str] function_app_id: Function App ID.
        :param pulumi.Input[str] href: Function URI.
        :param pulumi.Input[str] invoke_url_template: The invocation URL
        :param pulumi.Input[bool] is_disabled: Value indicating whether the function is disabled
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] language: The function language
        :param pulumi.Input[str] name: Function name.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[str] script_href: Script URI.
        :param pulumi.Input[str] script_root_path_href: Script root path URI.
        :param pulumi.Input[str] secrets_file_href: Secrets file URI.
        :param pulumi.Input[str] test_data: Test data used when testing via the Azure Portal.
        :param pulumi.Input[str] test_data_href: Test data URI.
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

            __props__['config'] = config
            __props__['config_href'] = config_href
            __props__['files'] = files
            __props__['function_app_id'] = function_app_id
            __props__['href'] = href
            __props__['invoke_url_template'] = invoke_url_template
            __props__['is_disabled'] = is_disabled
            __props__['kind'] = kind
            __props__['language'] = language
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['script_href'] = script_href
            __props__['script_root_path_href'] = script_root_path_href
            __props__['secrets_file_href'] = secrets_file_href
            __props__['test_data'] = test_data
            __props__['test_data_href'] = test_data_href
            __props__['properties'] = None
            __props__['type'] = None
        super(WebAppFunction, __self__).__init__(
            'azurerm:web/v20180201:WebAppFunction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WebAppFunction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppFunction(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
