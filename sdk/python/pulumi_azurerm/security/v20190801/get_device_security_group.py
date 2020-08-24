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
    'GetDeviceSecurityGroupResult',
    'AwaitableGetDeviceSecurityGroupResult',
    'get_device_security_group',
]

@pulumi.output_type
class GetDeviceSecurityGroupResult:
    """
    The device security group resource
    """
    def __init__(__self__, allowlist_rules=None, denylist_rules=None, name=None, threshold_rules=None, time_window_rules=None, type=None):
        if allowlist_rules and not isinstance(allowlist_rules, list):
            raise TypeError("Expected argument 'allowlist_rules' to be a list")
        pulumi.set(__self__, "allowlist_rules", allowlist_rules)
        if denylist_rules and not isinstance(denylist_rules, list):
            raise TypeError("Expected argument 'denylist_rules' to be a list")
        pulumi.set(__self__, "denylist_rules", denylist_rules)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if threshold_rules and not isinstance(threshold_rules, list):
            raise TypeError("Expected argument 'threshold_rules' to be a list")
        pulumi.set(__self__, "threshold_rules", threshold_rules)
        if time_window_rules and not isinstance(time_window_rules, list):
            raise TypeError("Expected argument 'time_window_rules' to be a list")
        pulumi.set(__self__, "time_window_rules", time_window_rules)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="allowlistRules")
    def allowlist_rules(self) -> Optional[List['outputs.AllowlistCustomAlertRuleResponse']]:
        """
        The allow-list custom alert rules.
        """
        return pulumi.get(self, "allowlist_rules")

    @property
    @pulumi.getter(name="denylistRules")
    def denylist_rules(self) -> Optional[List['outputs.DenylistCustomAlertRuleResponse']]:
        """
        The deny-list custom alert rules.
        """
        return pulumi.get(self, "denylist_rules")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="thresholdRules")
    def threshold_rules(self) -> Optional[List['outputs.ThresholdCustomAlertRuleResponse']]:
        """
        The list of custom alert threshold rules.
        """
        return pulumi.get(self, "threshold_rules")

    @property
    @pulumi.getter(name="timeWindowRules")
    def time_window_rules(self) -> Optional[List['outputs.TimeWindowCustomAlertRuleResponse']]:
        """
        The list of custom alert time-window rules.
        """
        return pulumi.get(self, "time_window_rules")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetDeviceSecurityGroupResult(GetDeviceSecurityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceSecurityGroupResult(
            allowlist_rules=self.allowlist_rules,
            denylist_rules=self.denylist_rules,
            name=self.name,
            threshold_rules=self.threshold_rules,
            time_window_rules=self.time_window_rules,
            type=self.type)


def get_device_security_group(name: Optional[str] = None,
                              resource_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceSecurityGroupResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the device security group. Note that the name of the device security group is case insensitive.
    :param str resource_id: The identifier of the resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceId'] = resource_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:security/v20190801:getDeviceSecurityGroup', __args__, opts=opts, typ=GetDeviceSecurityGroupResult).value

    return AwaitableGetDeviceSecurityGroupResult(
        allowlist_rules=__ret__.allowlist_rules,
        denylist_rules=__ret__.denylist_rules,
        name=__ret__.name,
        threshold_rules=__ret__.threshold_rules,
        time_window_rules=__ret__.time_window_rules,
        type=__ret__.type)
