# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Assessment(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Describes properties of an assessment.
      * `additional_data` (`dict`) - Additional data regarding the assessment
      * `display_name` (`str`) - User friendly display name of the assessment
      * `links` (`dict`) - Links relevant to the assessment
        * `azure_portal_uri` (`str`) - Link to assessment in Azure Portal

      * `metadata` (`dict`) - Describes properties of an assessment metadata.
        * `assessment_type` (`str`) - BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
        * `category` (`list`)
        * `description` (`str`) - Human readable description of the assessment
        * `display_name` (`str`) - User friendly display name of the assessment
        * `implementation_effort` (`str`) - The implementation effort required to remediate this assessment
        * `partner_data` (`dict`) - Describes the partner that created the assessment
          * `partner_name` (`str`) - Name of the company of the partner
          * `product_name` (`str`) - Name of the product of the partner that created the assessment
          * `secret` (`str`) - Secret to authenticate the partner and verify it created the assessment - write only

        * `policy_definition_id` (`str`) - Azure resource ID of the policy definition that turns this assessment calculation on
        * `preview` (`bool`) - True if this assessment is in preview release status
        * `remediation_description` (`str`) - Human readable description of what you should do to mitigate this security issue
        * `severity` (`str`) - The severity level of the assessment
        * `threats` (`list`)
        * `user_impact` (`str`) - The user impact of the assessment

      * `partners_data` (`dict`) - Data regarding 3rd party partner integration
        * `partner_name` (`str`) - Name of the company of the partner
        * `secret` (`str`) - secret to authenticate the partner - write only

      * `resource_details` (`dict`) - Details of the resource that was assessed
        * `source` (`str`) - The platform where the assessed resource resides

      * `status` (`dict`) - The result of the assessment
        * `cause` (`str`) - Programmatic code for the cause of the assessment status
        * `code` (`str`) - Programmatic code for the status of the assessment
        * `description` (`str`) - Human readable description of the assessment status
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Security assessment on a resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The Assessment Key - Unique key for the assessment type
        :param pulumi.Input[dict] properties: Describes properties of an assessment.
        :param pulumi.Input[str] resource_id: The identifier of the resource.

        The **properties** object supports the following:

          * `additional_data` (`pulumi.Input[dict]`) - Additional data regarding the assessment
          * `links` (`pulumi.Input[dict]`) - Links relevant to the assessment
          * `metadata` (`pulumi.Input[dict]`) - Describes properties of an assessment metadata.
            * `assessment_type` (`pulumi.Input[str]`) - BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
            * `category` (`pulumi.Input[list]`)
            * `description` (`pulumi.Input[str]`) - Human readable description of the assessment
            * `display_name` (`pulumi.Input[str]`) - User friendly display name of the assessment
            * `implementation_effort` (`pulumi.Input[str]`) - The implementation effort required to remediate this assessment
            * `partner_data` (`pulumi.Input[dict]`) - Describes the partner that created the assessment
              * `partner_name` (`pulumi.Input[str]`) - Name of the company of the partner
              * `product_name` (`pulumi.Input[str]`) - Name of the product of the partner that created the assessment
              * `secret` (`pulumi.Input[str]`) - Secret to authenticate the partner and verify it created the assessment - write only

            * `preview` (`pulumi.Input[bool]`) - True if this assessment is in preview release status
            * `remediation_description` (`pulumi.Input[str]`) - Human readable description of what you should do to mitigate this security issue
            * `severity` (`pulumi.Input[str]`) - The severity level of the assessment
            * `threats` (`pulumi.Input[list]`)
            * `user_impact` (`pulumi.Input[str]`) - The user impact of the assessment

          * `partners_data` (`pulumi.Input[dict]`) - Data regarding 3rd party partner integration
            * `partner_name` (`pulumi.Input[str]`) - Name of the company of the partner
            * `secret` (`pulumi.Input[str]`) - secret to authenticate the partner - write only

          * `resource_details` (`pulumi.Input[dict]`) - Details of the resource that was assessed
            * `source` (`pulumi.Input[str]`) - The platform where the assessed resource resides

          * `status` (`pulumi.Input[dict]`) - The result of the assessment
            * `cause` (`pulumi.Input[str]`) - Programmatic code for the cause of the assessment status
            * `code` (`pulumi.Input[str]`) - Programmatic code for the status of the assessment
            * `description` (`pulumi.Input[str]`) - Human readable description of the assessment status
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
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['type'] = None
        super(Assessment, __self__).__init__(
            'azurerm:security/v20200101:Assessment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Assessment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Assessment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
