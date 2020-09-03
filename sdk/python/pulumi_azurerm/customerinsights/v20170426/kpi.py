# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Kpi']


class Kpi(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aliases: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['KpiAliasArgs']]]]] = None,
                 calculation_window: Optional[pulumi.Input[str]] = None,
                 calculation_window_field_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 entity_type_name: Optional[pulumi.Input[str]] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 extracts: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['KpiExtractArgs']]]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 function: Optional[pulumi.Input[str]] = None,
                 group_by: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 hub_name: Optional[pulumi.Input[str]] = None,
                 kpi_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 thres_holds: Optional[pulumi.Input[pulumi.InputType['KpiThresholdsArgs']]] = None,
                 unit: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The KPI resource format.

        ## Example Usage
        ### Kpi_CreateOrUpdate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        kpi = azurerm.customerinsights.v20170426.Kpi("kpi",
            aliases=[{
                "aliasName": "alias",
                "expression": "Id+4",
            }],
            calculation_window="Day",
            description={
                "en-us": "Kpi Description",
            },
            display_name={
                "en-us": "Kpi DisplayName",
            },
            entity_type="Profile",
            entity_type_name="testProfile2327128",
            expression="SavingAccountBalance",
            function="Sum",
            group_by=["SavingAccountBalance"],
            hub_name="sdkTestHub",
            kpi_name="kpiTest45453647",
            resource_group_name="TestHubRG",
            thres_holds={
                "increasingKpi": True,
                "lowerLimit": 5,
                "upperLimit": 50,
            },
            unit="unit")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['KpiAliasArgs']]]] aliases: The aliases.
        :param pulumi.Input[str] calculation_window: The calculation window.
        :param pulumi.Input[str] calculation_window_field_name: Name of calculation window field.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] description: Localized description for the KPI.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] display_name: Localized display name for the KPI.
        :param pulumi.Input[str] entity_type: The mapping entity type.
        :param pulumi.Input[str] entity_type_name: The mapping entity name.
        :param pulumi.Input[str] expression: The computation expression for the KPI.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['KpiExtractArgs']]]] extracts: The KPI extracts.
        :param pulumi.Input[str] filter: The filter expression for the KPI.
        :param pulumi.Input[str] function: The computation function for the KPI.
        :param pulumi.Input[List[pulumi.Input[str]]] group_by: the group by properties for the KPI.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[str] kpi_name: The name of the KPI.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['KpiThresholdsArgs']] thres_holds: The KPI thresholds.
        :param pulumi.Input[str] unit: The unit of measurement for the KPI.
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
            if kpi_name is None:
                raise TypeError("Missing required property 'kpi_name'")
            __props__['kpi_name'] = kpi_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['thres_holds'] = thres_holds
            __props__['unit'] = unit
            __props__['group_by_metadata'] = None
            __props__['name'] = None
            __props__['participant_profiles_metadata'] = None
            __props__['provisioning_state'] = None
            __props__['tenant_id'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:customerinsights/latest:Kpi"), pulumi.Alias(type_="azurerm:customerinsights/v20170101:Kpi")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Kpi, __self__).__init__(
            'azurerm:customerinsights/v20170426:Kpi',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Kpi':
        """
        Get an existing Kpi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Kpi(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def aliases(self) -> pulumi.Output[Optional[List['outputs.KpiAliasResponse']]]:
        """
        The aliases.
        """
        return pulumi.get(self, "aliases")

    @property
    @pulumi.getter(name="calculationWindow")
    def calculation_window(self) -> pulumi.Output[str]:
        """
        The calculation window.
        """
        return pulumi.get(self, "calculation_window")

    @property
    @pulumi.getter(name="calculationWindowFieldName")
    def calculation_window_field_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of calculation window field.
        """
        return pulumi.get(self, "calculation_window_field_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Localized description for the KPI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Localized display name for the KPI.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> pulumi.Output[str]:
        """
        The mapping entity type.
        """
        return pulumi.get(self, "entity_type")

    @property
    @pulumi.getter(name="entityTypeName")
    def entity_type_name(self) -> pulumi.Output[str]:
        """
        The mapping entity name.
        """
        return pulumi.get(self, "entity_type_name")

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Output[str]:
        """
        The computation expression for the KPI.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def extracts(self) -> pulumi.Output[Optional[List['outputs.KpiExtractResponse']]]:
        """
        The KPI extracts.
        """
        return pulumi.get(self, "extracts")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional[str]]:
        """
        The filter expression for the KPI.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def function(self) -> pulumi.Output[str]:
        """
        The computation function for the KPI.
        """
        return pulumi.get(self, "function")

    @property
    @pulumi.getter(name="groupBy")
    def group_by(self) -> pulumi.Output[Optional[List[str]]]:
        """
        the group by properties for the KPI.
        """
        return pulumi.get(self, "group_by")

    @property
    @pulumi.getter(name="groupByMetadata")
    def group_by_metadata(self) -> pulumi.Output[List['outputs.KpiGroupByMetadataResponse']]:
        """
        The KPI GroupByMetadata.
        """
        return pulumi.get(self, "group_by_metadata")

    @property
    @pulumi.getter(name="kpiName")
    def kpi_name(self) -> pulumi.Output[str]:
        """
        The KPI name.
        """
        return pulumi.get(self, "kpi_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="participantProfilesMetadata")
    def participant_profiles_metadata(self) -> pulumi.Output[List['outputs.KpiParticipantProfilesMetadataResponse']]:
        """
        The participant profiles.
        """
        return pulumi.get(self, "participant_profiles_metadata")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The hub name.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="thresHolds")
    def thres_holds(self) -> pulumi.Output[Optional['outputs.KpiThresholdsResponse']]:
        """
        The KPI thresholds.
        """
        return pulumi.get(self, "thres_holds")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def unit(self) -> pulumi.Output[Optional[str]]:
        """
        The unit of measurement for the KPI.
        """
        return pulumi.get(self, "unit")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

