# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class WorkspaceProviderIncident(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Etag of the azure resource
    """
    name: pulumi.Output[str]
    """
    Azure resource name
    """
    properties: pulumi.Output[dict]
    """
    Incident properties
      * `additional_data` (`dict`) - Additional data on the incident
        * `alert_product_names` (`list`) - List of product names of alerts in the incident
        * `alerts_count` (`float`) - The number of alerts in the incident
        * `bookmarks_count` (`float`) - The number of bookmarks in the incident
        * `comments_count` (`float`) - The number of comments in the incident
        * `tactics` (`list`) - The tactics associated with incident

      * `classification` (`str`) - The reason the incident was closed
      * `classification_comment` (`str`) - Describes the reason the incident was closed
      * `classification_reason` (`str`) - The classification reason the incident was closed with
      * `created_time_utc` (`str`) - The time the incident was created
      * `description` (`str`) - The description of the incident
      * `first_activity_time_utc` (`str`) - The time of the first activity in the incident
      * `incident_number` (`float`) - A sequential number
      * `incident_url` (`str`) - The deep-link url to the incident in Azure portal
      * `labels` (`list`) - List of labels relevant to this incident
        * `label_name` (`str`) - The name of the label
        * `label_type` (`str`) - The type of the label

      * `last_activity_time_utc` (`str`) - The time of the last activity in the incident
      * `last_modified_time_utc` (`str`) - The last time the incident was updated
      * `owner` (`dict`) - Describes a user that the incident is assigned to
        * `assigned_to` (`str`) - The name of the user the incident is assigned to.
        * `email` (`str`) - The email of the user the incident is assigned to.
        * `object_id` (`str`) - The object id of the user the incident is assigned to.
        * `user_principal_name` (`str`) - The user principal name of the user the incident is assigned to.

      * `related_analytic_rule_ids` (`list`) - List of resource ids of Analytic rules related to the incident
      * `severity` (`str`) - The severity of the incident
      * `status` (`str`) - The status of the incident
      * `title` (`str`) - The title of the incident
    """
    type: pulumi.Output[str]
    """
    Azure resource type
    """
    def __init__(__self__, resource_name, opts=None, etag=None, name=None, properties=None, resource_group_name=None, workspace_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents an incident in Azure Security Insights.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: Etag of the azure resource
        :param pulumi.Input[str] name: Incident ID
        :param pulumi.Input[dict] properties: Incident properties
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
        :param pulumi.Input[str] workspace_name: The name of the workspace.

        The **properties** object supports the following:

          * `classification` (`pulumi.Input[str]`) - The reason the incident was closed
          * `classification_comment` (`pulumi.Input[str]`) - Describes the reason the incident was closed
          * `classification_reason` (`pulumi.Input[str]`) - The classification reason the incident was closed with
          * `description` (`pulumi.Input[str]`) - The description of the incident
          * `first_activity_time_utc` (`pulumi.Input[str]`) - The time of the first activity in the incident
          * `labels` (`pulumi.Input[list]`) - List of labels relevant to this incident
            * `label_name` (`pulumi.Input[str]`) - The name of the label

          * `last_activity_time_utc` (`pulumi.Input[str]`) - The time of the last activity in the incident
          * `owner` (`pulumi.Input[dict]`) - Describes a user that the incident is assigned to
            * `assigned_to` (`pulumi.Input[str]`) - The name of the user the incident is assigned to.
            * `email` (`pulumi.Input[str]`) - The email of the user the incident is assigned to.
            * `object_id` (`pulumi.Input[str]`) - The object id of the user the incident is assigned to.
            * `user_principal_name` (`pulumi.Input[str]`) - The user principal name of the user the incident is assigned to.

          * `severity` (`pulumi.Input[str]`) - The severity of the incident
          * `status` (`pulumi.Input[str]`) - The status of the incident
          * `title` (`pulumi.Input[str]`) - The title of the incident
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

            __props__['etag'] = etag
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['type'] = None
        super(WorkspaceProviderIncident, __self__).__init__(
            'azurerm:operationalinsights:WorkspaceProviderIncident',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WorkspaceProviderIncident resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WorkspaceProviderIncident(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
