# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class View(pulumi.CustomResource):
    e_tag: pulumi.Output[str]
    """
    eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the view.
      * `accumulated` (`str`) - Show costs accumulated over time.
      * `chart` (`str`) - Chart type of the main view in Cost Analysis. Required.
      * `created_on` (`str`) - Date the user created this view.
      * `display_name` (`str`) - User input name of the view. Required.
      * `kpis` (`list`) - List of KPIs to show in Cost Analysis UI.
        * `enabled` (`bool`) - show the KPI in the UI?
        * `id` (`str`) - ID of resource related to metric (budget).
        * `type` (`str`) - KPI type (Forecast, Budget).

      * `metric` (`str`) - Metric to use when displaying costs.
      * `modified_on` (`str`) - Date when the user last modified this view.
      * `pivots` (`list`) - Configuration of 3 sub-views in the Cost Analysis UI.
        * `name` (`str`) - Data field to show in view.
        * `type` (`str`) - Data type to show in view.

      * `query` (`dict`) - Query body configuration. Required.
        * `dataset` (`dict`) - Has definition for data in this report config.
          * `aggregation` (`dict`) - Dictionary of aggregation expression to use in the report. The key of each item in the dictionary is the alias for the aggregated column. Report can have up to 2 aggregation clauses.
          * `configuration` (`dict`) - Has configuration information for the data in the report. The configuration will be ignored if aggregation and grouping are provided.
            * `columns` (`list`) - Array of column names to be included in the report. Any valid report column name is allowed. If not provided, then report includes all columns.

          * `filter` (`dict`) - Has filter expression to use in the report.
            * `and` (`list`) - The logical "AND" expression. Must have at least 2 items.
            * `dimension` (`dict`) - Has comparison expression for a dimension
              * `name` (`str`) - The name of the column to use in comparison.
              * `operator` (`str`) - The operator to use for comparison.
              * `values` (`list`) - Array of values to use for comparison

            * `not` (`dict`) - The logical "NOT" expression.
            * `or` (`list`) - The logical "OR" expression. Must have at least 2 items.
            * `tag` (`dict`) - Has comparison expression for a tag

          * `granularity` (`str`) - The granularity of rows in the report.
          * `grouping` (`list`) - Array of group by expression to use in the report. Report can have up to 2 group by clauses.
            * `name` (`str`) - The name of the column to group. This version supports subscription lowest possible grain.
            * `type` (`str`) - Has type of the column to group.

          * `sorting` (`list`) - Array of order by expression to use in the report.
            * `direction` (`str`) - Direction of sort.
            * `name` (`str`) - The name of the column to sort.

        * `time_period` (`dict`) - Has time period for pulling data for the report.
          * `from` (`str`) - The start date to pull data from.
          * `to` (`str`) - The end date to pull data to.

        * `timeframe` (`str`) - The time frame for pulling data for the report. If custom, then a specific time period must be provided.
        * `type` (`str`) - The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.

      * `scope` (`str`) - Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, accumulated=None, chart=None, dataset=None, display_name=None, e_tag=None, kpis=None, metric=None, name=None, pivots=None, scope=None, time_period=None, timeframe=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        States and configurations of Cost Analysis.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accumulated: Show costs accumulated over time.
        :param pulumi.Input[str] chart: Chart type of the main view in Cost Analysis. Required.
        :param pulumi.Input[dict] dataset: Has definition for data in this report config.
        :param pulumi.Input[str] display_name: User input name of the view. Required.
        :param pulumi.Input[str] e_tag: eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        :param pulumi.Input[list] kpis: List of KPIs to show in Cost Analysis UI.
        :param pulumi.Input[str] metric: Metric to use when displaying costs.
        :param pulumi.Input[str] name: View name
        :param pulumi.Input[list] pivots: Configuration of 3 sub-views in the Cost Analysis UI.
        :param pulumi.Input[str] scope: Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
        :param pulumi.Input[dict] time_period: Has time period for pulling data for the report.
        :param pulumi.Input[str] timeframe: The time frame for pulling data for the report. If custom, then a specific time period must be provided.
        :param pulumi.Input[str] type: The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.

        The **dataset** object supports the following:

          * `aggregation` (`pulumi.Input[dict]`) - Dictionary of aggregation expression to use in the report. The key of each item in the dictionary is the alias for the aggregated column. Report can have up to 2 aggregation clauses.
          * `configuration` (`pulumi.Input[dict]`) - Has configuration information for the data in the report. The configuration will be ignored if aggregation and grouping are provided.
            * `columns` (`pulumi.Input[list]`) - Array of column names to be included in the report. Any valid report column name is allowed. If not provided, then report includes all columns.

          * `filter` (`pulumi.Input[dict]`) - Has filter expression to use in the report.
            * `and` (`pulumi.Input[list]`) - The logical "AND" expression. Must have at least 2 items.
            * `dimension` (`pulumi.Input[dict]`) - Has comparison expression for a dimension
              * `name` (`pulumi.Input[str]`) - The name of the column to use in comparison.
              * `operator` (`pulumi.Input[str]`) - The operator to use for comparison.
              * `values` (`pulumi.Input[list]`) - Array of values to use for comparison

            * `not` (`pulumi.Input[dict]`) - The logical "NOT" expression.
            * `or` (`pulumi.Input[list]`) - The logical "OR" expression. Must have at least 2 items.
            * `tag` (`pulumi.Input[dict]`) - Has comparison expression for a tag

          * `granularity` (`pulumi.Input[str]`) - The granularity of rows in the report.
          * `grouping` (`pulumi.Input[list]`) - Array of group by expression to use in the report. Report can have up to 2 group by clauses.
            * `name` (`pulumi.Input[str]`) - The name of the column to group. This version supports subscription lowest possible grain.
            * `type` (`pulumi.Input[str]`) - Has type of the column to group.

          * `sorting` (`pulumi.Input[list]`) - Array of order by expression to use in the report.
            * `direction` (`pulumi.Input[str]`) - Direction of sort.
            * `name` (`pulumi.Input[str]`) - The name of the column to sort.

        The **kpis** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - show the KPI in the UI?
          * `id` (`pulumi.Input[str]`) - ID of resource related to metric (budget).
          * `type` (`pulumi.Input[str]`) - KPI type (Forecast, Budget).

        The **pivots** object supports the following:

          * `name` (`pulumi.Input[str]`) - Data field to show in view.
          * `type` (`pulumi.Input[str]`) - Data type to show in view.

        The **time_period** object supports the following:

          * `from` (`pulumi.Input[str]`) - The start date to pull data from.
          * `to` (`pulumi.Input[str]`) - The end date to pull data to.
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

            __props__['accumulated'] = accumulated
            __props__['chart'] = chart
            __props__['dataset'] = dataset
            __props__['display_name'] = display_name
            __props__['e_tag'] = e_tag
            __props__['kpis'] = kpis
            __props__['metric'] = metric
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['pivots'] = pivots
            __props__['scope'] = scope
            __props__['time_period'] = time_period
            if timeframe is None:
                raise TypeError("Missing required property 'timeframe'")
            __props__['timeframe'] = timeframe
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['properties'] = None
        super(View, __self__).__init__(
            'azurerm:costmanagement/v20191101:View',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing View resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return View(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
