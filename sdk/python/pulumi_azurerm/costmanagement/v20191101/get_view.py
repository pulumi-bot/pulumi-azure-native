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
    'GetViewResult',
    'AwaitableGetViewResult',
    'get_view',
]

@pulumi.output_type
class GetViewResult:
    """
    States and configurations of Cost Analysis.
    """
    def __init__(__self__, accumulated=None, chart=None, created_on=None, dataset=None, display_name=None, e_tag=None, kpis=None, metric=None, modified_on=None, name=None, pivots=None, scope=None, time_period=None, timeframe=None, type=None):
        if accumulated and not isinstance(accumulated, str):
            raise TypeError("Expected argument 'accumulated' to be a str")
        pulumi.set(__self__, "accumulated", accumulated)
        if chart and not isinstance(chart, str):
            raise TypeError("Expected argument 'chart' to be a str")
        pulumi.set(__self__, "chart", chart)
        if created_on and not isinstance(created_on, str):
            raise TypeError("Expected argument 'created_on' to be a str")
        pulumi.set(__self__, "created_on", created_on)
        if dataset and not isinstance(dataset, dict):
            raise TypeError("Expected argument 'dataset' to be a dict")
        pulumi.set(__self__, "dataset", dataset)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if e_tag and not isinstance(e_tag, str):
            raise TypeError("Expected argument 'e_tag' to be a str")
        pulumi.set(__self__, "e_tag", e_tag)
        if kpis and not isinstance(kpis, list):
            raise TypeError("Expected argument 'kpis' to be a list")
        pulumi.set(__self__, "kpis", kpis)
        if metric and not isinstance(metric, str):
            raise TypeError("Expected argument 'metric' to be a str")
        pulumi.set(__self__, "metric", metric)
        if modified_on and not isinstance(modified_on, str):
            raise TypeError("Expected argument 'modified_on' to be a str")
        pulumi.set(__self__, "modified_on", modified_on)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pivots and not isinstance(pivots, list):
            raise TypeError("Expected argument 'pivots' to be a list")
        pulumi.set(__self__, "pivots", pivots)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
        if time_period and not isinstance(time_period, dict):
            raise TypeError("Expected argument 'time_period' to be a dict")
        pulumi.set(__self__, "time_period", time_period)
        if timeframe and not isinstance(timeframe, str):
            raise TypeError("Expected argument 'timeframe' to be a str")
        pulumi.set(__self__, "timeframe", timeframe)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def accumulated(self) -> Optional[str]:
        """
        Show costs accumulated over time.
        """
        return pulumi.get(self, "accumulated")

    @property
    @pulumi.getter
    def chart(self) -> Optional[str]:
        """
        Chart type of the main view in Cost Analysis. Required.
        """
        return pulumi.get(self, "chart")

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> str:
        """
        Date the user created this view.
        """
        return pulumi.get(self, "created_on")

    @property
    @pulumi.getter
    def dataset(self) -> Optional['outputs.ReportConfigDatasetResponse']:
        """
        Has definition for data in this report config.
        """
        return pulumi.get(self, "dataset")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        User input name of the view. Required.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> Optional[str]:
        """
        eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter
    def kpis(self) -> Optional[List['outputs.KpiPropertiesResponse']]:
        """
        List of KPIs to show in Cost Analysis UI.
        """
        return pulumi.get(self, "kpis")

    @property
    @pulumi.getter
    def metric(self) -> Optional[str]:
        """
        Metric to use when displaying costs.
        """
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter(name="modifiedOn")
    def modified_on(self) -> str:
        """
        Date when the user last modified this view.
        """
        return pulumi.get(self, "modified_on")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pivots(self) -> Optional[List['outputs.PivotPropertiesResponse']]:
        """
        Configuration of 3 sub-views in the Cost Analysis UI.
        """
        return pulumi.get(self, "pivots")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> Optional['outputs.ReportConfigTimePeriodResponse']:
        """
        Has time period for pulling data for the report.
        """
        return pulumi.get(self, "time_period")

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
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetViewResult(GetViewResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetViewResult(
            accumulated=self.accumulated,
            chart=self.chart,
            created_on=self.created_on,
            dataset=self.dataset,
            display_name=self.display_name,
            e_tag=self.e_tag,
            kpis=self.kpis,
            metric=self.metric,
            modified_on=self.modified_on,
            name=self.name,
            pivots=self.pivots,
            scope=self.scope,
            time_period=self.time_period,
            timeframe=self.timeframe,
            type=self.type)


def get_view(name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetViewResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: View name
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:costmanagement/v20191101:getView', __args__, opts=opts, typ=GetViewResult).value

    return AwaitableGetViewResult(
        accumulated=__ret__.accumulated,
        chart=__ret__.chart,
        created_on=__ret__.created_on,
        dataset=__ret__.dataset,
        display_name=__ret__.display_name,
        e_tag=__ret__.e_tag,
        kpis=__ret__.kpis,
        metric=__ret__.metric,
        modified_on=__ret__.modified_on,
        name=__ret__.name,
        pivots=__ret__.pivots,
        scope=__ret__.scope,
        time_period=__ret__.time_period,
        timeframe=__ret__.timeframe,
        type=__ret__.type)
