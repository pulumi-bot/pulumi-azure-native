# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Project(pulumi.CustomResource):
    e_tag: pulumi.Output[str]
    """
    For optimistic concurrency control.
    """
    location: pulumi.Output[str]
    """
    Azure location in which project is created.
    """
    name: pulumi.Output[str]
    """
    Name of the project.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the project.
      * `created_timestamp` (`str`) - Time when this project was created. Date-Time represented in ISO-8601 format.
      * `customer_workspace_id` (`str`) - ARM ID of the Service Map workspace created by user.
      * `customer_workspace_location` (`str`) - Location of the Service Map workspace created by user.
      * `discovery_status` (`str`) - Reports whether project is under discovery.
      * `last_assessment_timestamp` (`str`) - Time when last assessment was created. Date-Time represented in ISO-8601 format. This value will be null until assessment is created.
      * `last_discovery_session_id` (`str`) - Session id of the last discovery.
      * `last_discovery_timestamp` (`str`) - Time when this project was created. Date-Time represented in ISO-8601 format. This value will be null until discovery is complete.
      * `number_of_assessments` (`float`) - Number of assessments created in the project.
      * `number_of_groups` (`float`) - Number of groups created in the project.
      * `number_of_machines` (`float`) - Number of machines in the project.
      * `provisioning_state` (`str`) - Provisioning state of the project.
      * `updated_timestamp` (`str`) - Time when this project was last updated. Date-Time represented in ISO-8601 format.
    """
    tags: pulumi.Output[dict]
    """
    Tags provided by Azure Tagging service.
    """
    type: pulumi.Output[str]
    """
    Type of the object = [Microsoft.Migrate/projects].
    """
    def __init__(__self__, resource_name, opts=None, e_tag=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Azure Migrate Project.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] e_tag: For optimistic concurrency control.
        :param pulumi.Input[str] location: Azure location in which project is created.
        :param pulumi.Input[str] name: Name of the Azure Migrate project.
        :param pulumi.Input[dict] properties: Properties of the project.
        :param pulumi.Input[str] resource_group_name: Name of the Azure Resource Group that project is part of.
        :param pulumi.Input[dict] tags: Tags provided by Azure Tagging service.

        The **properties** object supports the following:

          * `customer_workspace_id` (`pulumi.Input[str]`) - ARM ID of the Service Map workspace created by user.
          * `customer_workspace_location` (`pulumi.Input[str]`) - Location of the Service Map workspace created by user.
          * `provisioning_state` (`pulumi.Input[str]`) - Provisioning state of the project.
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

            __props__['e_tag'] = e_tag
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(Project, __self__).__init__(
            'azurerm:migrate/v20180202:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Project(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
