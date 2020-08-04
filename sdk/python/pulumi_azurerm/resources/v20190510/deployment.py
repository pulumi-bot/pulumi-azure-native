# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Deployment(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    the location of the deployment.
    """
    name: pulumi.Output[str]
    """
    The name of the deployment.
    """
    properties: pulumi.Output[dict]
    """
    Deployment properties.
      * `correlation_id` (`str`) - The correlation ID of the deployment.
      * `debug_setting` (`dict`) - The debug setting of the deployment.
        * `detail_level` (`str`) - Specifies the type of information to log for debugging. The permitted values are none, requestContent, responseContent, or both requestContent and responseContent separated by a comma. The default is none. When setting this value, carefully consider the type of information you are passing in during deployment. By logging information about the request or response, you could potentially expose sensitive data that is retrieved through the deployment operations.

      * `dependencies` (`list`) - The list of deployment dependencies.
        * `depends_on` (`list`) - The list of dependencies.
          * `id` (`str`) - The ID of the dependency.
          * `resource_name` (`str`) - The dependency resource name.
          * `resource_type` (`str`) - The dependency resource type.

        * `id` (`str`) - The ID of the dependency.
        * `resource_name` (`str`) - The dependency resource name.
        * `resource_type` (`str`) - The dependency resource type.

      * `duration` (`str`) - The duration of the template deployment.
      * `mode` (`str`) - The deployment mode. Possible values are Incremental and Complete.
      * `on_error_deployment` (`dict`) - The deployment on error behavior.
        * `deployment_name` (`str`) - The deployment to be used on error case.
        * `provisioning_state` (`str`) - The state of the provisioning for the on error deployment.
        * `type` (`str`) - The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.

      * `outputs` (`dict`) - Key/value pairs that represent deployment output.
      * `parameters` (`dict`) - Deployment parameters. Use only one of Parameters or ParametersLink.
      * `parameters_link` (`dict`) - The URI referencing the parameters. Use only one of Parameters or ParametersLink.
        * `content_version` (`str`) - If included, must match the ContentVersion in the template.
        * `uri` (`str`) - The URI of the parameters file.

      * `providers` (`list`) - The list of resource providers needed for the deployment.
        * `id` (`str`) - The provider ID.
        * `namespace` (`str`) - The namespace of the resource provider.
        * `registration_policy` (`str`) - The registration policy of the resource provider.
        * `registration_state` (`str`) - The registration state of the resource provider.
        * `resource_types` (`list`) - The collection of provider resource types.
          * `aliases` (`list`) - The aliases that are supported by this resource type.
            * `name` (`str`) - The alias name.
            * `paths` (`list`) - The paths for an alias.
              * `api_versions` (`list`) - The API versions.
              * `path` (`str`) - The path of an alias.

          * `api_versions` (`list`) - The API version.
          * `capabilities` (`str`) - The additional capabilities offered by this resource type.
          * `locations` (`list`) - The collection of locations where this resource type can be created.
          * `properties` (`dict`) - The properties.
          * `resource_type` (`str`) - The resource type.

      * `provisioning_state` (`str`) - The state of the provisioning.
      * `template` (`dict`) - The template content. Use only one of Template or TemplateLink.
      * `template_link` (`dict`) - The URI referencing the template. Use only one of Template or TemplateLink.
        * `content_version` (`str`) - If included, must match the ContentVersion in the template.
        * `uri` (`str`) - The URI of the template to deploy.

      * `timestamp` (`str`) - The timestamp of the template deployment.
    """
    type: pulumi.Output[str]
    """
    The type of the deployment.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Deployment information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location to store the deployment data.
        :param pulumi.Input[str] name: The name of the deployment.
        :param pulumi.Input[dict] properties: The deployment properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to deploy the resources to. The name is case insensitive. The resource group must already exist.

        The **properties** object supports the following:

          * `debug_setting` (`pulumi.Input[dict]`) - The debug setting of the deployment.
            * `detail_level` (`pulumi.Input[str]`) - Specifies the type of information to log for debugging. The permitted values are none, requestContent, responseContent, or both requestContent and responseContent separated by a comma. The default is none. When setting this value, carefully consider the type of information you are passing in during deployment. By logging information about the request or response, you could potentially expose sensitive data that is retrieved through the deployment operations.

          * `mode` (`pulumi.Input[str]`) - The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode, resources are deployed without deleting existing resources that are not included in the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included in the template are deleted. Be careful when using Complete mode as you may unintentionally delete resources.
          * `on_error_deployment` (`pulumi.Input[dict]`) - The deployment on error behavior.
            * `deployment_name` (`pulumi.Input[str]`) - The deployment to be used on error case.
            * `type` (`pulumi.Input[str]`) - The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.

          * `parameters` (`pulumi.Input[dict]`) - Name and value pairs that define the deployment parameters for the template. You use this element when you want to provide the parameter values directly in the request rather than link to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
          * `parameters_link` (`pulumi.Input[dict]`) - The URI of parameters file. You use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
            * `content_version` (`pulumi.Input[str]`) - If included, must match the ContentVersion in the template.
            * `uri` (`pulumi.Input[str]`) - The URI of the parameters file.

          * `template` (`pulumi.Input[dict]`) - The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
          * `template_link` (`pulumi.Input[dict]`) - The URI of the template. Use either the templateLink property or the template property, but not both.
            * `content_version` (`pulumi.Input[str]`) - If included, must match the ContentVersion in the template.
            * `uri` (`pulumi.Input[str]`) - The URI of the template to deploy.
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

            __props__['location'] = location
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
        super(Deployment, __self__).__init__(
            'azurerm:resources/v20190510:Deployment',
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
