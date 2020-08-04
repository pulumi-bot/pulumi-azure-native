# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class EmailTemplate(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Email Template entity contract properties.
      * `body` (`str`) - Email Template Body. This should be a valid XDocument
      * `description` (`str`) - Description of the Email Template.
      * `is_default` (`bool`) - Whether the template is the default template provided by Api Management or has been edited.
      * `parameters` (`list`) - Email Template Parameter values.
        * `description` (`str`) - Template parameter description.
        * `name` (`str`) - Template parameter name.
        * `title` (`str`) - Template parameter title.

      * `subject` (`str`) - Subject of the Template.
      * `title` (`str`) - Title of the Template.
    """
    type: pulumi.Output[str]
    """
    Resource type for API Management resource.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, service_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Email Template details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Email Template Name Identifier.
        :param pulumi.Input[dict] properties: Email Template Update contract properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.

        The **properties** object supports the following:

          * `body` (`pulumi.Input[str]`) - Email Template Body. This should be a valid XDocument
          * `description` (`pulumi.Input[str]`) - Description of the Email Template.
          * `parameters` (`pulumi.Input[list]`) - Email Template Parameter values.
            * `description` (`pulumi.Input[str]`) - Template parameter description.
            * `name` (`pulumi.Input[str]`) - Template parameter name.
            * `title` (`pulumi.Input[str]`) - Template parameter title.

          * `subject` (`pulumi.Input[str]`) - Subject of the Template.
          * `title` (`pulumi.Input[str]`) - Title of the Template.
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
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['type'] = None
        super(EmailTemplate, __self__).__init__(
            'azurerm:apimanagement/v20190101:EmailTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing EmailTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return EmailTemplate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
