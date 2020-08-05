# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Kpi(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Defines the KPI Threshold limits.
      * `aliases` (`list`) - The aliases.
        * `alias_name` (`str`) - KPI alias name.
        * `expression` (`str`) - The expression.

      * `calculation_window` (`str`) - The calculation window.
      * `calculation_window_field_name` (`str`) - Name of calculation window field.
      * `description` (`dict`) - Localized description for the KPI.
      * `display_name` (`dict`) - Localized display name for the KPI.
      * `entity_type` (`str`) - The mapping entity type.
      * `entity_type_name` (`str`) - The mapping entity name.
      * `expression` (`str`) - The computation expression for the KPI.
      * `extracts` (`list`) - The KPI extracts.
        * `expression` (`str`) - The expression.
        * `extract_name` (`str`) - KPI extract name.

      * `filter` (`str`) - The filter expression for the KPI.
      * `function` (`str`) - The computation function for the KPI.
      * `group_by` (`list`) - the group by properties for the KPI.
      * `group_by_metadata` (`list`) - The KPI GroupByMetadata.
        * `display_name` (`dict`) - The display name.
        * `field_name` (`str`) - The name of the field.
        * `field_type` (`str`) - The type of the field.

      * `kpi_name` (`str`) - The KPI name.
      * `participant_profiles_metadata` (`list`) - The participant profiles.
        * `type_name` (`str`) - Name of the type.

      * `provisioning_state` (`str`) - Provisioning state.
      * `tenant_id` (`str`) - The hub name.
      * `thres_holds` (`dict`) - The KPI thresholds.
        * `increasing_kpi` (`bool`) - Whether or not the KPI is an increasing KPI.
        * `lower_limit` (`float`) - The lower threshold limit.
        * `upper_limit` (`float`) - The upper threshold limit.

      * `unit` (`str`) - The unit of measurement for the KPI.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, aliases=None, calculation_window=None, calculation_window_field_name=None, description=None, display_name=None, entity_type=None, entity_type_name=None, expression=None, extracts=None, filter=None, function=None, group_by=None, hub_name=None, name=None, resource_group_name=None, thres_holds=None, unit=None, __props__=None, __name__=None, __opts__=None):
        """
        The KPI resource format.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] aliases: The aliases.
        :param pulumi.Input[str] calculation_window: The calculation window.
        :param pulumi.Input[str] calculation_window_field_name: Name of calculation window field.
        :param pulumi.Input[dict] description: Localized description for the KPI.
        :param pulumi.Input[dict] display_name: Localized display name for the KPI.
        :param pulumi.Input[str] entity_type: The mapping entity type.
        :param pulumi.Input[str] entity_type_name: The mapping entity name.
        :param pulumi.Input[str] expression: The computation expression for the KPI.
        :param pulumi.Input[list] extracts: The KPI extracts.
        :param pulumi.Input[str] filter: The filter expression for the KPI.
        :param pulumi.Input[str] function: The computation function for the KPI.
        :param pulumi.Input[list] group_by: the group by properties for the KPI.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[str] name: The name of the KPI.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] thres_holds: The KPI thresholds.
        :param pulumi.Input[str] unit: The unit of measurement for the KPI.

        The **aliases** object supports the following:

          * `alias_name` (`pulumi.Input[str]`) - KPI alias name.
          * `expression` (`pulumi.Input[str]`) - The expression.

        The **extracts** object supports the following:

          * `expression` (`pulumi.Input[str]`) - The expression.
          * `extract_name` (`pulumi.Input[str]`) - KPI extract name.

        The **thres_holds** object supports the following:

          * `increasing_kpi` (`pulumi.Input[bool]`) - Whether or not the KPI is an increasing KPI.
          * `lower_limit` (`pulumi.Input[float]`) - The lower threshold limit.
          * `upper_limit` (`pulumi.Input[float]`) - The upper threshold limit.
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

            __props__['aliases'] = aliases
            if calculation_window is None:
                raise TypeError("Missing required property 'calculation_window'")
            __props__['calculation_window'] = calculation_window
            __props__['calculation_window_field_name'] = calculation_window_field_name
            __props__['description'] = description
            __props__['display_name'] = display_name
            if entity_type is None:
                raise TypeError("Missing required property 'entity_type'")
            __props__['entity_type'] = entity_type
            if entity_type_name is None:
                raise TypeError("Missing required property 'entity_type_name'")
            __props__['entity_type_name'] = entity_type_name
            if expression is None:
                raise TypeError("Missing required property 'expression'")
            __props__['expression'] = expression
            __props__['extracts'] = extracts
            __props__['filter'] = filter
            if function is None:
                raise TypeError("Missing required property 'function'")
            __props__['function'] = function
            __props__['group_by'] = group_by
            if hub_name is None:
                raise TypeError("Missing required property 'hub_name'")
            __props__['hub_name'] = hub_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['thres_holds'] = thres_holds
            __props__['unit'] = unit
            __props__['properties'] = None
            __props__['type'] = None
        super(Kpi, __self__).__init__(
            'azurerm:customerinsights/v20170101:Kpi',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Kpi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Kpi(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
