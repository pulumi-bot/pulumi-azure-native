# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class RosettaNetProcessConfiguration(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The resource location.
    """
    name: pulumi.Output[str]
    """
    Gets the resource name.
    """
    properties: pulumi.Output[dict]
    """
    The integration account RosettaNet process configuration properties.
      * `activity_settings` (`dict`) - The RosettaNet process configuration activity settings.
        * `acknowledgment_of_receipt_settings` (`dict`) - The RosettaNet ProcessConfiguration acknowledgement settings.
          * `is_non_repudiation_required` (`bool`) - The non-repudiation is required or not.
          * `time_to_acknowledge_in_seconds` (`float`) - The time to acknowledge in seconds.

        * `activity_behavior` (`dict`) - The RosettaNet ProcessConfiguration activity behavior.
          * `action_type` (`str`) - The value indicating whether the RosettaNet PIP is used for a single action.
          * `is_authorization_required` (`bool`) - The value indicating whether authorization is required.
          * `is_secured_transport_required` (`bool`) - The value indicating whether secured transport is required.
          * `non_repudiation_of_origin_and_content` (`bool`) - The value indicating whether non-repudiation is for origin and content.
          * `persistent_confidentiality_scope` (`str`) - The persistent confidentiality encryption scope.
          * `response_type` (`str`) - The value indicating whether the RosettaNet PIP communication is synchronous.
          * `retry_count` (`float`) - The value indicating retry count.
          * `time_to_perform_in_seconds` (`float`) - The time to perform in seconds.

        * `activity_type` (`str`) - The RosettaNet ProcessConfiguration activity type.

      * `changed_time` (`str`) - The changed time.
      * `created_time` (`str`) - The created time.
      * `description` (`str`) - The integration account RosettaNet ProcessConfiguration properties.
      * `initiator_role_settings` (`dict`) - The RosettaNet initiator role settings.
        * `action` (`str`) - The action name.
        * `business_document` (`dict`) - The RosettaNet ProcessConfiguration business document.
          * `description` (`str`) - The business document description.
          * `name` (`str`) - The business document name.
          * `version` (`str`) - The business document version.

        * `description` (`str`) - The description.
        * `role` (`str`) - The role name.
        * `role_type` (`str`) - The RosettaNet ProcessConfiguration role type.
        * `service` (`str`) - The service name.
        * `service_classification` (`str`) - The service classification name.

      * `metadata` (`dict`) - The metadata.
      * `process_code` (`str`) - The integration account RosettaNet process code.
      * `process_name` (`str`) - The integration account RosettaNet process name.
      * `process_version` (`str`) - The integration account RosettaNet process version.
      * `responder_role_settings` (`dict`) - The RosettaNet responder role settings.
    """
    tags: pulumi.Output[dict]
    """
    The resource tags.
    """
    type: pulumi.Output[str]
    """
    Gets the resource type.
    """
    def __init__(__self__, resource_name, opts=None, integration_account_name=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The integration account RosettaNet process configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] integration_account_name: The integration account name.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] name: The integration account RosettaNet ProcessConfiguration name.
        :param pulumi.Input[dict] properties: The integration account RosettaNet process configuration properties.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[dict] tags: The resource tags.

        The **properties** object supports the following:

          * `activity_settings` (`pulumi.Input[dict]`) - The RosettaNet process configuration activity settings.
            * `acknowledgment_of_receipt_settings` (`pulumi.Input[dict]`) - The RosettaNet ProcessConfiguration acknowledgement settings.
              * `is_non_repudiation_required` (`pulumi.Input[bool]`) - The non-repudiation is required or not.
              * `time_to_acknowledge_in_seconds` (`pulumi.Input[float]`) - The time to acknowledge in seconds.

            * `activity_behavior` (`pulumi.Input[dict]`) - The RosettaNet ProcessConfiguration activity behavior.
              * `action_type` (`pulumi.Input[str]`) - The value indicating whether the RosettaNet PIP is used for a single action.
              * `is_authorization_required` (`pulumi.Input[bool]`) - The value indicating whether authorization is required.
              * `is_secured_transport_required` (`pulumi.Input[bool]`) - The value indicating whether secured transport is required.
              * `non_repudiation_of_origin_and_content` (`pulumi.Input[bool]`) - The value indicating whether non-repudiation is for origin and content.
              * `persistent_confidentiality_scope` (`pulumi.Input[str]`) - The persistent confidentiality encryption scope.
              * `response_type` (`pulumi.Input[str]`) - The value indicating whether the RosettaNet PIP communication is synchronous.
              * `retry_count` (`pulumi.Input[float]`) - The value indicating retry count.
              * `time_to_perform_in_seconds` (`pulumi.Input[float]`) - The time to perform in seconds.

            * `activity_type` (`pulumi.Input[str]`) - The RosettaNet ProcessConfiguration activity type.

          * `description` (`pulumi.Input[str]`) - The integration account RosettaNet ProcessConfiguration properties.
          * `initiator_role_settings` (`pulumi.Input[dict]`) - The RosettaNet initiator role settings.
            * `action` (`pulumi.Input[str]`) - The action name.
            * `business_document` (`pulumi.Input[dict]`) - The RosettaNet ProcessConfiguration business document.
              * `description` (`pulumi.Input[str]`) - The business document description.
              * `name` (`pulumi.Input[str]`) - The business document name.
              * `version` (`pulumi.Input[str]`) - The business document version.

            * `description` (`pulumi.Input[str]`) - The description.
            * `role` (`pulumi.Input[str]`) - The role name.
            * `role_type` (`pulumi.Input[str]`) - The RosettaNet ProcessConfiguration role type.
            * `service` (`pulumi.Input[str]`) - The service name.
            * `service_classification` (`pulumi.Input[str]`) - The service classification name.

          * `metadata` (`pulumi.Input[dict]`) - The metadata.
          * `process_code` (`pulumi.Input[str]`) - The integration account RosettaNet process code.
          * `process_name` (`pulumi.Input[str]`) - The integration account RosettaNet process name.
          * `process_version` (`pulumi.Input[str]`) - The integration account RosettaNet process version.
          * `responder_role_settings` (`pulumi.Input[dict]`) - The RosettaNet responder role settings.
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

            if integration_account_name is None:
                raise TypeError("Missing required property 'integration_account_name'")
            __props__['integration_account_name'] = integration_account_name
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
            __props__['tags'] = tags
            __props__['type'] = None
        super(RosettaNetProcessConfiguration, __self__).__init__(
            'azurerm:logic/v20160601:RosettaNetProcessConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RosettaNetProcessConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RosettaNetProcessConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop