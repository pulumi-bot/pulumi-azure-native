# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Connector(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Properties of connector.
      * `connector_id` (`float`) - ID of the connector.
      * `connector_name` (`str`) - Name of the connector.
      * `connector_properties` (`dict`) - The connector properties.
      * `connector_type` (`str`) - Type of connector.
      * `created` (`str`) - The created time.
      * `description` (`str`) - Description of the connector.
      * `display_name` (`str`) - Display name of the connector.
      * `is_internal` (`bool`) - If this is an internal connector.
      * `last_modified` (`str`) - The last modified time.
      * `state` (`str`) - State of connector.
      * `tenant_id` (`str`) - The hub name.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, connector_properties=None, connector_type=None, description=None, display_name=None, hub_name=None, is_internal=None, name=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The connector resource format.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] connector_properties: The connector properties.
        :param pulumi.Input[str] connector_type: Type of connector.
        :param pulumi.Input[str] description: Description of the connector.
        :param pulumi.Input[str] display_name: Display name of the connector.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[bool] is_internal: If this is an internal connector.
        :param pulumi.Input[str] name: Name of the connector.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
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

            if connector_properties is None:
                raise TypeError("Missing required property 'connector_properties'")
            __props__['connector_properties'] = connector_properties
            if connector_type is None:
                raise TypeError("Missing required property 'connector_type'")
            __props__['connector_type'] = connector_type
            __props__['description'] = description
            __props__['display_name'] = display_name
            if hub_name is None:
                raise TypeError("Missing required property 'hub_name'")
            __props__['hub_name'] = hub_name
            __props__['is_internal'] = is_internal
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['properties'] = None
            __props__['type'] = None
        super(Connector, __self__).__init__(
            'azurerm:customerinsights/v20170101:Connector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Connector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Connector(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
