# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class AutomationAccountRunbook(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Gets or sets the etag of the resource.
    """
    location: pulumi.Output[str]
    """
    The Azure Region where the resource lives
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Gets or sets the runbook properties.
      * `creation_time` (`str`) - Gets or sets the creation time.
      * `description` (`str`) - Gets or sets the description.
      * `draft` (`dict`) - Gets or sets the draft runbook properties.
        * `creation_time` (`str`) - Gets or sets the creation time of the runbook draft.
        * `draft_content_link` (`dict`) - Gets or sets the draft runbook content link.
          * `content_hash` (`dict`) - Gets or sets the hash.
            * `algorithm` (`str`) - Gets or sets the content hash algorithm used to hash the content.
            * `value` (`str`) - Gets or sets expected hash value of the content.

          * `uri` (`str`) - Gets or sets the uri of the runbook content.
          * `version` (`str`) - Gets or sets the version of the content.

        * `in_edit` (`bool`) - Gets or sets whether runbook is in edit mode.
        * `last_modified_time` (`str`) - Gets or sets the last modified time of the runbook draft.
        * `output_types` (`list`) - Gets or sets the runbook output types.
        * `parameters` (`dict`) - Gets or sets the runbook draft parameters.

      * `job_count` (`float`) - Gets or sets the job count of the runbook.
      * `last_modified_by` (`str`) - Gets or sets the last modified by.
      * `last_modified_time` (`str`) - Gets or sets the last modified time.
      * `log_activity_trace` (`float`) - Gets or sets the option to log activity trace of the runbook.
      * `log_progress` (`bool`) - Gets or sets progress log option.
      * `log_verbose` (`bool`) - Gets or sets verbose log option.
      * `output_types` (`list`) - Gets or sets the runbook output types.
      * `parameters` (`dict`) - Gets or sets the runbook parameters.
      * `provisioning_state` (`str`) - Gets or sets the provisioning state of the runbook.
      * `publish_content_link` (`dict`) - Gets or sets the published runbook content link.
      * `runbook_type` (`str`) - Gets or sets the type of the runbook.
      * `state` (`str`) - Gets or sets the state of the runbook.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, automation_account_name=None, location=None, name=None, properties=None, resource_group_name=None, runbook_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Definition of the runbook type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account.
        :param pulumi.Input[str] location: Gets or sets the location of the resource.
        :param pulumi.Input[str] name: Gets or sets the name of the resource.
        :param pulumi.Input[dict] properties: Gets or sets runbook create or update properties.
        :param pulumi.Input[str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[str] runbook_name: The runbook name.
        :param pulumi.Input[dict] tags: Gets or sets the tags attached to the resource.

        The **properties** object supports the following:

          * `description` (`pulumi.Input[str]`) - Gets or sets the description of the runbook.
          * `draft` (`pulumi.Input[dict]`) - Gets or sets the draft runbook properties.
            * `creation_time` (`pulumi.Input[str]`) - Gets or sets the creation time of the runbook draft.
            * `draft_content_link` (`pulumi.Input[dict]`) - Gets or sets the draft runbook content link.
              * `content_hash` (`pulumi.Input[dict]`) - Gets or sets the hash.
                * `algorithm` (`pulumi.Input[str]`) - Gets or sets the content hash algorithm used to hash the content.
                * `value` (`pulumi.Input[str]`) - Gets or sets expected hash value of the content.

              * `uri` (`pulumi.Input[str]`) - Gets or sets the uri of the runbook content.
              * `version` (`pulumi.Input[str]`) - Gets or sets the version of the content.

            * `in_edit` (`pulumi.Input[bool]`) - Gets or sets whether runbook is in edit mode.
            * `last_modified_time` (`pulumi.Input[str]`) - Gets or sets the last modified time of the runbook draft.
            * `output_types` (`pulumi.Input[list]`) - Gets or sets the runbook output types.
            * `parameters` (`pulumi.Input[dict]`) - Gets or sets the runbook draft parameters.

          * `log_activity_trace` (`pulumi.Input[float]`) - Gets or sets the activity-level tracing options of the runbook.
          * `log_progress` (`pulumi.Input[bool]`) - Gets or sets progress log option.
          * `log_verbose` (`pulumi.Input[bool]`) - Gets or sets verbose log option.
          * `publish_content_link` (`pulumi.Input[dict]`) - Gets or sets the published runbook content link.
          * `runbook_type` (`pulumi.Input[str]`) - Gets or sets the type of the runbook.
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
            __props__['location'] = location
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if runbook_name is None:
                raise TypeError("Missing required property 'runbook_name'")
            __props__['runbook_name'] = runbook_name
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['type'] = None
        super(AutomationAccountRunbook, __self__).__init__(
            'azurerm:automation:AutomationAccountRunbook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AutomationAccountRunbook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AutomationAccountRunbook(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
