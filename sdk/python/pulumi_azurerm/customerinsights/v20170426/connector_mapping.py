# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ConnectorMapping(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    The connector mapping definition.
      * `connector_mapping_name` (`str`) - The connector mapping name
      * `connector_name` (`str`) - The connector name.
      * `connector_type` (`str`) - Type of connector.
      * `created` (`str`) - The created time.
      * `data_format_id` (`str`) - The DataFormat ID.
      * `description` (`str`) - The description of the connector mapping.
      * `display_name` (`str`) - Display name for the connector mapping.
      * `entity_type` (`str`) - Defines which entity type the file should map to.
      * `entity_type_name` (`str`) - The mapping entity name.
      * `last_modified` (`str`) - The last modified time.
      * `mapping_properties` (`dict`) - The properties of the mapping.
        * `availability` (`dict`) - The availability of mapping property.
          * `frequency` (`str`) - The frequency to update.
          * `interval` (`float`) - The interval of the given frequency to use.

        * `complete_operation` (`dict`) - The operation after import is done.
          * `completion_operation_type` (`str`) - The type of completion operation.
          * `destination_folder` (`str`) - The destination folder where files will be moved to once the import is done.

        * `error_management` (`dict`) - The error management setting for the mapping.
          * `error_limit` (`float`) - The error limit allowed while importing data.
          * `error_management_type` (`str`) - The type of error management to use for the mapping.

        * `file_filter` (`str`) - The file filter for the mapping.
        * `folder_path` (`str`) - The folder path for the mapping.
        * `format` (`dict`) - The format of mapping property.
          * `accept_language` (`str`) - The oData language.
          * `array_separator` (`str`) - Character separating array elements.
          * `column_delimiter` (`str`) - The character that signifies a break between columns.
          * `format_type` (`str`) - The type mapping format.
          * `quote_character` (`str`) - Quote character, used to indicate enquoted fields.
          * `quote_escape_character` (`str`) - Escape character for quotes, can be the same as the quoteCharacter.

        * `has_header` (`bool`) - If the file contains a header or not.
        * `structure` (`list`) - Ingestion mapping information at property level.
          * `column_name` (`str`) - The column name of the import file.
          * `custom_format_specifier` (`str`) - Custom format specifier for input parsing.
          * `is_encrypted` (`bool`) - Indicates if the column is encrypted.
          * `property_name` (`str`) - The property name of the mapping entity.

      * `next_run_time` (`str`) - The next run time based on customer's settings.
      * `run_id` (`str`) - The RunId.
      * `state` (`str`) - State of connector mapping.
      * `tenant_id` (`str`) - The hub name.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, connector_name=None, hub_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The connector mapping resource format.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connector_name: The name of the connector.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[str] name: The name of the connector mapping.
        :param pulumi.Input[dict] properties: The connector mapping definition.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **properties** object supports the following:

          * `connector_type` (`pulumi.Input[str]`) - Type of connector.
          * `description` (`pulumi.Input[str]`) - The description of the connector mapping.
          * `display_name` (`pulumi.Input[str]`) - Display name for the connector mapping.
          * `entity_type` (`pulumi.Input[str]`) - Defines which entity type the file should map to.
          * `entity_type_name` (`pulumi.Input[str]`) - The mapping entity name.
          * `mapping_properties` (`pulumi.Input[dict]`) - The properties of the mapping.
            * `availability` (`pulumi.Input[dict]`) - The availability of mapping property.
              * `frequency` (`pulumi.Input[str]`) - The frequency to update.
              * `interval` (`pulumi.Input[float]`) - The interval of the given frequency to use.

            * `complete_operation` (`pulumi.Input[dict]`) - The operation after import is done.
              * `completion_operation_type` (`pulumi.Input[str]`) - The type of completion operation.
              * `destination_folder` (`pulumi.Input[str]`) - The destination folder where files will be moved to once the import is done.

            * `error_management` (`pulumi.Input[dict]`) - The error management setting for the mapping.
              * `error_limit` (`pulumi.Input[float]`) - The error limit allowed while importing data.
              * `error_management_type` (`pulumi.Input[str]`) - The type of error management to use for the mapping.

            * `file_filter` (`pulumi.Input[str]`) - The file filter for the mapping.
            * `folder_path` (`pulumi.Input[str]`) - The folder path for the mapping.
            * `format` (`pulumi.Input[dict]`) - The format of mapping property.
              * `accept_language` (`pulumi.Input[str]`) - The oData language.
              * `array_separator` (`pulumi.Input[str]`) - Character separating array elements.
              * `column_delimiter` (`pulumi.Input[str]`) - The character that signifies a break between columns.
              * `format_type` (`pulumi.Input[str]`) - The type mapping format.
              * `quote_character` (`pulumi.Input[str]`) - Quote character, used to indicate enquoted fields.
              * `quote_escape_character` (`pulumi.Input[str]`) - Escape character for quotes, can be the same as the quoteCharacter.

            * `has_header` (`pulumi.Input[bool]`) - If the file contains a header or not.
            * `structure` (`pulumi.Input[list]`) - Ingestion mapping information at property level.
              * `column_name` (`pulumi.Input[str]`) - The column name of the import file.
              * `custom_format_specifier` (`pulumi.Input[str]`) - Custom format specifier for input parsing.
              * `is_encrypted` (`pulumi.Input[bool]`) - Indicates if the column is encrypted.
              * `property_name` (`pulumi.Input[str]`) - The property name of the mapping entity.
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

            if connector_name is None:
                raise TypeError("Missing required property 'connector_name'")
            __props__['connector_name'] = connector_name
            if hub_name is None:
                raise TypeError("Missing required property 'hub_name'")
            __props__['hub_name'] = hub_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(ConnectorMapping, __self__).__init__(
            'azurerm:customerinsights/v20170426:ConnectorMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ConnectorMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ConnectorMapping(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
