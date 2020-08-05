# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Deployment(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the deployment.
    """
    properties: pulumi.Output[dict]
    """
    Deployment properties.
      * `correlation_id` (`str`) - The correlation ID of the deployment.
      * `debug_setting` (`dict`) - The debug setting of the deployment.
        * `detail_level` (`str`) - The debug detail level.

      * `dependencies` (`list`) - The list of deployment dependencies.
        * `depends_on` (`list`) - The list of dependencies.
          * `id` (`str`) - The ID of the dependency.
          * `resource_name` (`str`) - The dependency resource name.
          * `resource_type` (`str`) - The dependency resource type.

        * `id` (`str`) - The ID of the dependency.
        * `resource_name` (`str`) - The dependency resource name.
        * `resource_type` (`str`) - The dependency resource type.

      * `mode` (`str`) - The deployment mode.
      * `outputs` (`dict`) - Key/value pairs that represent deployment output.
      * `parameters` (`dict`) - Deployment parameters. Use only one of Parameters or ParametersLink.
      * `parameters_link` (`dict`) - The URI referencing the parameters. Use only one of Parameters or ParametersLink.
        * `content_version` (`str`) - If included it must match the ContentVersion in the template.
        * `uri` (`str`) - URI referencing the template.

      * `providers` (`list`) - The list of resource providers needed for the deployment.
        * `id` (`str`) - The provider id.
        * `namespace` (`str`) - The namespace of the provider.
        * `registration_state` (`str`) - The registration state of the provider.
        * `resource_types` (`list`) - The collection of provider resource types.
          * `aliases` (`list`) - The aliases that are supported by this resource type.
            * `name` (`str`) - The alias name.
            * `paths` (`list`) - The paths for an alias.
              * `api_versions` (`list`) - The api versions.
              * `path` (`str`) - The path of an alias.

          * `api_versions` (`list`) - The api version.
          * `locations` (`list`) - The collection of locations where this resource type can be created in.
          * `properties` (`dict`) - The properties.
          * `resource_type` (`str`) - The resource type.

      * `provisioning_state` (`str`) - The state of the provisioning.
      * `template` (`dict`) - The template content. Use only one of Template or TemplateLink.
      * `template_link` (`dict`) - The URI referencing the template. Use only one of Template or TemplateLink.
        * `content_version` (`str`) - If included it must match the ContentVersion in the template.
        * `uri` (`str`) - URI referencing the template.

      * `timestamp` (`str`) - The timestamp of the template deployment.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Deployment information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the deployment.
        :param pulumi.Input[dict] properties: The deployment properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.

        The **properties** object supports the following:

          * `debug_setting` (`pulumi.Input[dict]`) - The debug setting of the deployment.
          * `mode` (`pulumi.Input[str]`) - The deployment mode.
          * `parameters` (`pulumi.Input[dict]`) - Deployment parameters. It can be a JObject or a well formed JSON string. Use only one of Parameters or ParametersLink.
          * `parameters_link` (`pulumi.Input[dict]`) - The parameters URI. Use only one of Parameters or ParametersLink.
            * `content_version` (`pulumi.Input[str]`) - If included it must match the ContentVersion in the template.
            * `uri` (`pulumi.Input[str]`) - URI referencing the template.

          * `template` (`pulumi.Input[dict]`) - The template content. It can be a JObject or a well formed JSON string. Use only one of Template or TemplateLink.
          * `template_link` (`pulumi.Input[dict]`) - The template URI. Use only one of Template or TemplateLink.
            * `content_version` (`pulumi.Input[str]`) - If included it must match the ContentVersion in the template.
            * `uri` (`pulumi.Input[str]`) - URI referencing the template.
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
        super(Deployment, __self__).__init__(
            'azurerm:resources/v20160201:Deployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Deployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Deployment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
