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
    'GetBudgetResult',
    'AwaitableGetBudgetResult',
    'get_budget',
]

@pulumi.output_type
class GetBudgetResult:
    """
    A budget resource.
    """
    def __init__(__self__, amount=None, category=None, current_spend=None, e_tag=None, filter=None, name=None, notifications=None, time_grain=None, time_period=None, type=None):
        if amount and not isinstance(amount, float):
            raise TypeError("Expected argument 'amount' to be a float")
        pulumi.set(__self__, "amount", amount)
        if category and not isinstance(category, str):
            raise TypeError("Expected argument 'category' to be a str")
        pulumi.set(__self__, "category", category)
        if current_spend and not isinstance(current_spend, dict):
            raise TypeError("Expected argument 'current_spend' to be a dict")
        pulumi.set(__self__, "current_spend", current_spend)
        if e_tag and not isinstance(e_tag, str):
            raise TypeError("Expected argument 'e_tag' to be a str")
        pulumi.set(__self__, "e_tag", e_tag)
        if filter and not isinstance(filter, dict):
            raise TypeError("Expected argument 'filter' to be a dict")
        pulumi.set(__self__, "filter", filter)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notifications and not isinstance(notifications, dict):
            raise TypeError("Expected argument 'notifications' to be a dict")
        pulumi.set(__self__, "notifications", notifications)
        if time_grain and not isinstance(time_grain, str):
            raise TypeError("Expected argument 'time_grain' to be a str")
        pulumi.set(__self__, "time_grain", time_grain)
        if time_period and not isinstance(time_period, dict):
            raise TypeError("Expected argument 'time_period' to be a dict")
        pulumi.set(__self__, "time_period", time_period)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def amount(self) -> float:
        """
        The total amount of cost to track with the budget
        """
        return pulumi.get(self, "amount")

    @property
    @pulumi.getter
    def category(self) -> str:
        """
        The category of the budget, whether the budget tracks cost or usage.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="currentSpend")
    def current_spend(self) -> 'outputs.CurrentSpendResponse':
        """
        The current amount of cost which is being tracked for a budget.
        """
        return pulumi.get(self, "current_spend")

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> Optional[str]:
        """
        eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter
    def filter(self) -> Optional['outputs.BudgetFilterResponse']:
        """
        May be used to filter budgets by resource group, resource, or meter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> Optional[Mapping[str, 'outputs.NotificationResponse']]:
        """
        Dictionary of notifications associated with the budget. Budget can have up to five notifications.
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter(name="timeGrain")
    def time_grain(self) -> str:
        """
        The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
        """
        return pulumi.get(self, "time_grain")

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> 'outputs.BudgetTimePeriodResponse':
        """
        Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
        """
        return pulumi.get(self, "time_period")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetBudgetResult(GetBudgetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBudgetResult(
            amount=self.amount,
            category=self.category,
            current_spend=self.current_spend,
            e_tag=self.e_tag,
            filter=self.filter,
            name=self.name,
            notifications=self.notifications,
            time_grain=self.time_grain,
            time_period=self.time_period,
            type=self.type)


def get_budget(name: Optional[str] = None,
               scope: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBudgetResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Budget Name.
    :param str scope: The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['scope'] = scope
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:consumption/v20191001:getBudget', __args__, opts=opts, typ=GetBudgetResult).value

    return AwaitableGetBudgetResult(
        amount=__ret__.amount,
        category=__ret__.category,
        current_spend=__ret__.current_spend,
        e_tag=__ret__.e_tag,
        filter=__ret__.filter,
        name=__ret__.name,
        notifications=__ret__.notifications,
        time_grain=__ret__.time_grain,
        time_period=__ret__.time_period,
        type=__ret__.type)
