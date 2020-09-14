# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'AzurePlanResponseResult',
    'CostAllocationProportionResponse',
    'CostAllocationRuleDetailsResponse',
    'CostAllocationRulePropertiesResponse',
    'InvoiceSectionWithCreateSubPermissionResponseResult',
    'ReportAggregationResponse',
    'ReportComparisonExpressionResponse',
    'ReportDatasetConfigurationResponse',
    'ReportDatasetResponse',
    'ReportDefinitionResponse',
    'ReportDeliveryDestinationResponse',
    'ReportDeliveryInfoResponse',
    'ReportFilterResponse',
    'ReportGroupingResponse',
    'ReportRecurrencePeriodResponse',
    'ReportScheduleResponse',
    'ReportTimePeriodResponse',
    'SourceCostAllocationResourceResponse',
    'TargetCostAllocationResourceResponse',
]

@pulumi.output_type
class AzurePlanResponseResult(dict):
    """
    Details of the Azure plan.
    """
    def __init__(__self__, *,
                 sku_description: str,
                 sku_id: Optional[str] = None):
        """
        Details of the Azure plan.
        :param str sku_description: The sku description.
        :param str sku_id: The sku id.
        """
        pulumi.set(__self__, "sku_description", sku_description)
        if sku_id is not None:
            pulumi.set(__self__, "sku_id", sku_id)

    @property
    @pulumi.getter(name="skuDescription")
    def sku_description(self) -> str:
        """
        The sku description.
        """
        return pulumi.get(self, "sku_description")

    @property
    @pulumi.getter(name="skuId")
    def sku_id(self) -> Optional[str]:
        """
        The sku id.
        """
        return pulumi.get(self, "sku_id")


@pulumi.output_type
class CostAllocationProportionResponse(dict):
    """
    Target resources and allocation
    """
    def __init__(__self__, *,
                 name: str,
                 percentage: float):
        """
        Target resources and allocation
        :param str name: Target resource for cost allocation
        :param float percentage: Percentage of source cost to allocate to this resource. This value can be specified to two decimal places and the total percentage of all resources in this rule must sum to 100.00.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "percentage", percentage)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Target resource for cost allocation
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def percentage(self) -> float:
        """
        Percentage of source cost to allocate to this resource. This value can be specified to two decimal places and the total percentage of all resources in this rule must sum to 100.00.
        """
        return pulumi.get(self, "percentage")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CostAllocationRuleDetailsResponse(dict):
    """
    Resource details of the cost allocation rule
    """
    def __init__(__self__, *,
                 source_resources: Optional[List['outputs.SourceCostAllocationResourceResponse']] = None,
                 target_resources: Optional[List['outputs.TargetCostAllocationResourceResponse']] = None):
        """
        Resource details of the cost allocation rule
        :param List['SourceCostAllocationResourceResponseArgs'] source_resources: Source resources for cost allocation. At this time, this list can contain no more than one element.
        :param List['TargetCostAllocationResourceResponseArgs'] target_resources: Target resources for cost allocation. At this time, this list can contain no more than one element.
        """
        if source_resources is not None:
            pulumi.set(__self__, "source_resources", source_resources)
        if target_resources is not None:
            pulumi.set(__self__, "target_resources", target_resources)

    @property
    @pulumi.getter(name="sourceResources")
    def source_resources(self) -> Optional[List['outputs.SourceCostAllocationResourceResponse']]:
        """
        Source resources for cost allocation. At this time, this list can contain no more than one element.
        """
        return pulumi.get(self, "source_resources")

    @property
    @pulumi.getter(name="targetResources")
    def target_resources(self) -> Optional[List['outputs.TargetCostAllocationResourceResponse']]:
        """
        Target resources for cost allocation. At this time, this list can contain no more than one element.
        """
        return pulumi.get(self, "target_resources")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CostAllocationRulePropertiesResponse(dict):
    """
    The properties of a cost allocation rule
    """
    def __init__(__self__, *,
                 created_date: str,
                 details: 'outputs.CostAllocationRuleDetailsResponse',
                 status: str,
                 updated_date: str,
                 description: Optional[str] = None):
        """
        The properties of a cost allocation rule
        :param str created_date: Time at which the rule was created. Rules that change cost for the same resource are applied in order of creation.
        :param 'CostAllocationRuleDetailsResponseArgs' details: Resource information for the cost allocation rule
        :param str status: Status of the rule
        :param str updated_date: Time at which the rule was last updated.
        :param str description: Description of a cost allocation rule.
        """
        pulumi.set(__self__, "created_date", created_date)
        pulumi.set(__self__, "details", details)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "updated_date", updated_date)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        Time at which the rule was created. Rules that change cost for the same resource are applied in order of creation.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def details(self) -> 'outputs.CostAllocationRuleDetailsResponse':
        """
        Resource information for the cost allocation rule
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the rule
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedDate")
    def updated_date(self) -> str:
        """
        Time at which the rule was last updated.
        """
        return pulumi.get(self, "updated_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of a cost allocation rule.
        """
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InvoiceSectionWithCreateSubPermissionResponseResult(dict):
    """
    Invoice section properties with create subscription permission.
    """
    def __init__(__self__, *,
                 billing_profile_display_name: str,
                 billing_profile_id: str,
                 billing_profile_spending_limit: str,
                 billing_profile_status: str,
                 billing_profile_status_reason_code: str,
                 billing_profile_system_id: str,
                 invoice_section_display_name: str,
                 invoice_section_id: str,
                 invoice_section_system_id: str,
                 enabled_azure_plans: Optional[List['outputs.AzurePlanResponseResult']] = None):
        """
        Invoice section properties with create subscription permission.
        :param str billing_profile_display_name: The name of the billing profile for the invoice section.
        :param str billing_profile_id: The ID of the billing profile for the invoice section.
        :param str billing_profile_spending_limit: The billing profile spending limit.
        :param str billing_profile_status: The status of the billing profile.
        :param str billing_profile_status_reason_code: Reason for the specified billing profile status.
        :param str billing_profile_system_id: The system generated unique identifier for a billing profile.
        :param str invoice_section_display_name: The name of the invoice section.
        :param str invoice_section_id: The ID of the invoice section.
        :param str invoice_section_system_id: The system generated unique identifier for an invoice section.
        :param List['AzurePlanResponseArgs'] enabled_azure_plans: Enabled azure plans for the associated billing profile.
        """
        pulumi.set(__self__, "billing_profile_display_name", billing_profile_display_name)
        pulumi.set(__self__, "billing_profile_id", billing_profile_id)
        pulumi.set(__self__, "billing_profile_spending_limit", billing_profile_spending_limit)
        pulumi.set(__self__, "billing_profile_status", billing_profile_status)
        pulumi.set(__self__, "billing_profile_status_reason_code", billing_profile_status_reason_code)
        pulumi.set(__self__, "billing_profile_system_id", billing_profile_system_id)
        pulumi.set(__self__, "invoice_section_display_name", invoice_section_display_name)
        pulumi.set(__self__, "invoice_section_id", invoice_section_id)
        pulumi.set(__self__, "invoice_section_system_id", invoice_section_system_id)
        if enabled_azure_plans is not None:
            pulumi.set(__self__, "enabled_azure_plans", enabled_azure_plans)

    @property
    @pulumi.getter(name="billingProfileDisplayName")
    def billing_profile_display_name(self) -> str:
        """
        The name of the billing profile for the invoice section.
        """
        return pulumi.get(self, "billing_profile_display_name")

    @property
    @pulumi.getter(name="billingProfileId")
    def billing_profile_id(self) -> str:
        """
        The ID of the billing profile for the invoice section.
        """
        return pulumi.get(self, "billing_profile_id")

    @property
    @pulumi.getter(name="billingProfileSpendingLimit")
    def billing_profile_spending_limit(self) -> str:
        """
        The billing profile spending limit.
        """
        return pulumi.get(self, "billing_profile_spending_limit")

    @property
    @pulumi.getter(name="billingProfileStatus")
    def billing_profile_status(self) -> str:
        """
        The status of the billing profile.
        """
        return pulumi.get(self, "billing_profile_status")

    @property
    @pulumi.getter(name="billingProfileStatusReasonCode")
    def billing_profile_status_reason_code(self) -> str:
        """
        Reason for the specified billing profile status.
        """
        return pulumi.get(self, "billing_profile_status_reason_code")

    @property
    @pulumi.getter(name="billingProfileSystemId")
    def billing_profile_system_id(self) -> str:
        """
        The system generated unique identifier for a billing profile.
        """
        return pulumi.get(self, "billing_profile_system_id")

    @property
    @pulumi.getter(name="invoiceSectionDisplayName")
    def invoice_section_display_name(self) -> str:
        """
        The name of the invoice section.
        """
        return pulumi.get(self, "invoice_section_display_name")

    @property
    @pulumi.getter(name="invoiceSectionId")
    def invoice_section_id(self) -> str:
        """
        The ID of the invoice section.
        """
        return pulumi.get(self, "invoice_section_id")

    @property
    @pulumi.getter(name="invoiceSectionSystemId")
    def invoice_section_system_id(self) -> str:
        """
        The system generated unique identifier for an invoice section.
        """
        return pulumi.get(self, "invoice_section_system_id")

    @property
    @pulumi.getter(name="enabledAzurePlans")
    def enabled_azure_plans(self) -> Optional[List['outputs.AzurePlanResponseResult']]:
        """
        Enabled azure plans for the associated billing profile.
        """
        return pulumi.get(self, "enabled_azure_plans")


@pulumi.output_type
class ReportAggregationResponse(dict):
    """
    The aggregation expression to be used in the report.
    """
    def __init__(__self__, *,
                 function: str,
                 name: str):
        """
        The aggregation expression to be used in the report.
        :param str function: The name of the aggregation function to use.
        :param str name: The name of the column to aggregate.
        """
        pulumi.set(__self__, "function", function)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def function(self) -> str:
        """
        The name of the aggregation function to use.
        """
        return pulumi.get(self, "function")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the column to aggregate.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportComparisonExpressionResponse(dict):
    """
    The comparison expression to be used in the report.
    """
    def __init__(__self__, *,
                 name: str,
                 operator: str,
                 values: List[str]):
        """
        The comparison expression to be used in the report.
        :param str name: The name of the column to use in comparison.
        :param str operator: The operator to use for comparison.
        :param List[str] values: Array of values to use for comparison
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the column to use in comparison.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def operator(self) -> str:
        """
        The operator to use for comparison.
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        Array of values to use for comparison
        """
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportDatasetConfigurationResponse(dict):
    """
    The configuration of dataset in the report.
    """
    def __init__(__self__, *,
                 columns: Optional[List[str]] = None):
        """
        The configuration of dataset in the report.
        :param List[str] columns: Array of column names to be included in the report. Any valid report column name is allowed. If not provided, then report includes all columns.
        """
        if columns is not None:
            pulumi.set(__self__, "columns", columns)

    @property
    @pulumi.getter
    def columns(self) -> Optional[List[str]]:
        """
        Array of column names to be included in the report. Any valid report column name is allowed. If not provided, then report includes all columns.
        """
        return pulumi.get(self, "columns")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportDatasetResponse(dict):
    """
    The definition of data present in the report.
    """
    def __init__(__self__, *,
                 aggregation: Optional[Mapping[str, 'outputs.ReportAggregationResponse']] = None,
                 configuration: Optional['outputs.ReportDatasetConfigurationResponse'] = None,
                 filter: Optional['outputs.ReportFilterResponse'] = None,
                 granularity: Optional[str] = None,
                 grouping: Optional[List['outputs.ReportGroupingResponse']] = None):
        """
        The definition of data present in the report.
        :param Mapping[str, 'ReportAggregationResponseArgs'] aggregation: Dictionary of aggregation expression to use in the report. The key of each item in the dictionary is the alias for the aggregated column. Report can have up to 2 aggregation clauses.
        :param 'ReportDatasetConfigurationResponseArgs' configuration: Has configuration information for the data in the report. The configuration will be ignored if aggregation and grouping are provided.
        :param 'ReportFilterResponseArgs' filter: Has filter expression to use in the report.
        :param str granularity: The granularity of rows in the report.
        :param List['ReportGroupingResponseArgs'] grouping: Array of group by expression to use in the report. Report can have up to 2 group by clauses.
        """
        if aggregation is not None:
            pulumi.set(__self__, "aggregation", aggregation)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if granularity is not None:
            pulumi.set(__self__, "granularity", granularity)
        if grouping is not None:
            pulumi.set(__self__, "grouping", grouping)

    @property
    @pulumi.getter
    def aggregation(self) -> Optional[Mapping[str, 'outputs.ReportAggregationResponse']]:
        """
        Dictionary of aggregation expression to use in the report. The key of each item in the dictionary is the alias for the aggregated column. Report can have up to 2 aggregation clauses.
        """
        return pulumi.get(self, "aggregation")

    @property
    @pulumi.getter
    def configuration(self) -> Optional['outputs.ReportDatasetConfigurationResponse']:
        """
        Has configuration information for the data in the report. The configuration will be ignored if aggregation and grouping are provided.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def filter(self) -> Optional['outputs.ReportFilterResponse']:
        """
        Has filter expression to use in the report.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def granularity(self) -> Optional[str]:
        """
        The granularity of rows in the report.
        """
        return pulumi.get(self, "granularity")

    @property
    @pulumi.getter
    def grouping(self) -> Optional[List['outputs.ReportGroupingResponse']]:
        """
        Array of group by expression to use in the report. Report can have up to 2 group by clauses.
        """
        return pulumi.get(self, "grouping")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportDefinitionResponse(dict):
    """
    The definition of a report.
    """
    def __init__(__self__, *,
                 timeframe: str,
                 type: str,
                 dataset: Optional['outputs.ReportDatasetResponse'] = None,
                 time_period: Optional['outputs.ReportTimePeriodResponse'] = None):
        """
        The definition of a report.
        :param str timeframe: The time frame for pulling data for the report. If custom, then a specific time period must be provided.
        :param str type: The type of the report.
        :param 'ReportDatasetResponseArgs' dataset: Has definition for data in this report.
        :param 'ReportTimePeriodResponseArgs' time_period: Has time period for pulling data for the report.
        """
        pulumi.set(__self__, "timeframe", timeframe)
        pulumi.set(__self__, "type", type)
        if dataset is not None:
            pulumi.set(__self__, "dataset", dataset)
        if time_period is not None:
            pulumi.set(__self__, "time_period", time_period)

    @property
    @pulumi.getter
    def timeframe(self) -> str:
        """
        The time frame for pulling data for the report. If custom, then a specific time period must be provided.
        """
        return pulumi.get(self, "timeframe")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the report.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def dataset(self) -> Optional['outputs.ReportDatasetResponse']:
        """
        Has definition for data in this report.
        """
        return pulumi.get(self, "dataset")

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> Optional['outputs.ReportTimePeriodResponse']:
        """
        Has time period for pulling data for the report.
        """
        return pulumi.get(self, "time_period")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportDeliveryDestinationResponse(dict):
    """
    The destination information for the delivery of the report.
    """
    def __init__(__self__, *,
                 container: str,
                 resource_id: str,
                 root_folder_path: Optional[str] = None):
        """
        The destination information for the delivery of the report.
        :param str container: The name of the container where reports will be uploaded.
        :param str resource_id: The resource id of the storage account where reports will be delivered.
        :param str root_folder_path: The name of the directory where reports will be uploaded.
        """
        pulumi.set(__self__, "container", container)
        pulumi.set(__self__, "resource_id", resource_id)
        if root_folder_path is not None:
            pulumi.set(__self__, "root_folder_path", root_folder_path)

    @property
    @pulumi.getter
    def container(self) -> str:
        """
        The name of the container where reports will be uploaded.
        """
        return pulumi.get(self, "container")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The resource id of the storage account where reports will be delivered.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="rootFolderPath")
    def root_folder_path(self) -> Optional[str]:
        """
        The name of the directory where reports will be uploaded.
        """
        return pulumi.get(self, "root_folder_path")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportDeliveryInfoResponse(dict):
    """
    The delivery information associated with a report.
    """
    def __init__(__self__, *,
                 destination: 'outputs.ReportDeliveryDestinationResponse'):
        """
        The delivery information associated with a report.
        :param 'ReportDeliveryDestinationResponseArgs' destination: Has destination for the report being delivered.
        """
        pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> 'outputs.ReportDeliveryDestinationResponse':
        """
        Has destination for the report being delivered.
        """
        return pulumi.get(self, "destination")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportFilterResponse(dict):
    """
    The filter expression to be used in the report.
    """
    def __init__(__self__, *,
                 and_: Optional[List['outputs.ReportFilterResponse']] = None,
                 dimension: Optional['outputs.ReportComparisonExpressionResponse'] = None,
                 not_: Optional['outputs.ReportFilterResponse'] = None,
                 or_: Optional[List['outputs.ReportFilterResponse']] = None,
                 tag: Optional['outputs.ReportComparisonExpressionResponse'] = None):
        """
        The filter expression to be used in the report.
        :param List['ReportFilterResponseArgs'] and_: The logical "AND" expression. Must have at least 2 items.
        :param 'ReportComparisonExpressionResponseArgs' dimension: Has comparison expression for a dimension
        :param 'ReportFilterResponseArgs' not_: The logical "NOT" expression.
        :param List['ReportFilterResponseArgs'] or_: The logical "OR" expression. Must have at least 2 items.
        :param 'ReportComparisonExpressionResponseArgs' tag: Has comparison expression for a tag
        """
        if and_ is not None:
            pulumi.set(__self__, "and_", and_)
        if dimension is not None:
            pulumi.set(__self__, "dimension", dimension)
        if not_ is not None:
            pulumi.set(__self__, "not_", not_)
        if or_ is not None:
            pulumi.set(__self__, "or_", or_)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter(name="and")
    def and_(self) -> Optional[List['outputs.ReportFilterResponse']]:
        """
        The logical "AND" expression. Must have at least 2 items.
        """
        return pulumi.get(self, "and_")

    @property
    @pulumi.getter
    def dimension(self) -> Optional['outputs.ReportComparisonExpressionResponse']:
        """
        Has comparison expression for a dimension
        """
        return pulumi.get(self, "dimension")

    @property
    @pulumi.getter(name="not")
    def not_(self) -> Optional['outputs.ReportFilterResponse']:
        """
        The logical "NOT" expression.
        """
        return pulumi.get(self, "not_")

    @property
    @pulumi.getter(name="or")
    def or_(self) -> Optional[List['outputs.ReportFilterResponse']]:
        """
        The logical "OR" expression. Must have at least 2 items.
        """
        return pulumi.get(self, "or_")

    @property
    @pulumi.getter
    def tag(self) -> Optional['outputs.ReportComparisonExpressionResponse']:
        """
        Has comparison expression for a tag
        """
        return pulumi.get(self, "tag")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportGroupingResponse(dict):
    """
    The group by expression to be used in the report.
    """
    def __init__(__self__, *,
                 name: str,
                 type: str):
        """
        The group by expression to be used in the report.
        :param str name: The name of the column to group.
        :param str type: Has type of the column to group.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the column to group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Has type of the column to group.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportRecurrencePeriodResponse(dict):
    """
    The start and end date for recurrence schedule.
    """
    def __init__(__self__, *,
                 from_: str,
                 to: Optional[str] = None):
        """
        The start and end date for recurrence schedule.
        :param str from_: The start date of recurrence.
        :param str to: The end date of recurrence.
        """
        pulumi.set(__self__, "from_", from_)
        if to is not None:
            pulumi.set(__self__, "to", to)

    @property
    @pulumi.getter(name="from")
    def from_(self) -> str:
        """
        The start date of recurrence.
        """
        return pulumi.get(self, "from_")

    @property
    @pulumi.getter
    def to(self) -> Optional[str]:
        """
        The end date of recurrence.
        """
        return pulumi.get(self, "to")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportScheduleResponse(dict):
    """
    The schedule associated with a report.
    """
    def __init__(__self__, *,
                 recurrence: str,
                 recurrence_period: Optional['outputs.ReportRecurrencePeriodResponse'] = None,
                 status: Optional[str] = None):
        """
        The schedule associated with a report.
        :param str recurrence: The schedule recurrence.
        :param 'ReportRecurrencePeriodResponseArgs' recurrence_period: Has start and end date of the recurrence. The start date must be in future. If present, the end date must be greater than start date.
        :param str status: The status of the schedule. Whether active or not. If inactive, the report's scheduled execution is paused.
        """
        pulumi.set(__self__, "recurrence", recurrence)
        if recurrence_period is not None:
            pulumi.set(__self__, "recurrence_period", recurrence_period)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def recurrence(self) -> str:
        """
        The schedule recurrence.
        """
        return pulumi.get(self, "recurrence")

    @property
    @pulumi.getter(name="recurrencePeriod")
    def recurrence_period(self) -> Optional['outputs.ReportRecurrencePeriodResponse']:
        """
        Has start and end date of the recurrence. The start date must be in future. If present, the end date must be greater than start date.
        """
        return pulumi.get(self, "recurrence_period")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the schedule. Whether active or not. If inactive, the report's scheduled execution is paused.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReportTimePeriodResponse(dict):
    """
    The start and end date for pulling data for the report.
    """
    def __init__(__self__, *,
                 from_: str,
                 to: str):
        """
        The start and end date for pulling data for the report.
        :param str from_: The start date to pull data from.
        :param str to: The end date to pull data to.
        """
        pulumi.set(__self__, "from_", from_)
        pulumi.set(__self__, "to", to)

    @property
    @pulumi.getter(name="from")
    def from_(self) -> str:
        """
        The start date to pull data from.
        """
        return pulumi.get(self, "from_")

    @property
    @pulumi.getter
    def to(self) -> str:
        """
        The end date to pull data to.
        """
        return pulumi.get(self, "to")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SourceCostAllocationResourceResponse(dict):
    """
    Source resources for cost allocation
    """
    def __init__(__self__, *,
                 name: str,
                 resource_type: str,
                 values: List[str]):
        """
        Source resources for cost allocation
        :param str name: If resource type is dimension, this must be either ResourceGroupName or SubscriptionId. If resource type is tag, this must be a valid Azure tag
        :param str resource_type: Type of resources contained in this cost allocation rule
        :param List[str] values: Source Resources for cost allocation. This list cannot contain more than 25 values.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        If resource type is dimension, this must be either ResourceGroupName or SubscriptionId. If resource type is tag, this must be a valid Azure tag
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        Type of resources contained in this cost allocation rule
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        Source Resources for cost allocation. This list cannot contain more than 25 values.
        """
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TargetCostAllocationResourceResponse(dict):
    """
    Target resources for cost allocation.
    """
    def __init__(__self__, *,
                 name: str,
                 policy_type: str,
                 resource_type: str,
                 values: List['outputs.CostAllocationProportionResponse']):
        """
        Target resources for cost allocation.
        :param str name: If resource type is dimension, this must be either ResourceGroupName or SubscriptionId. If resource type is tag, this must be a valid Azure tag
        :param str policy_type: Method of cost allocation for the rule
        :param str resource_type: Type of resources contained in this cost allocation rule
        :param List['CostAllocationProportionResponseArgs'] values: Target resources for cost allocation. This list cannot contain more than 25 values.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "policy_type", policy_type)
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        If resource type is dimension, this must be either ResourceGroupName or SubscriptionId. If resource type is tag, this must be a valid Azure tag
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> str:
        """
        Method of cost allocation for the rule
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        Type of resources contained in this cost allocation rule
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def values(self) -> List['outputs.CostAllocationProportionResponse']:
        """
        Target resources for cost allocation. This list cannot contain more than 25 values.
        """
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


