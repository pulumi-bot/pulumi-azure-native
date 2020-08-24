# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetAlertRuleActionResult',
    'AwaitableGetAlertRuleActionResult',
    'get_alert_rule_action',
]

@pulumi.output_type
class GetAlertRuleActionResult:
    """
    Action for alert rule.
    """
    def __init__(__self__, etag=None, logic_app_resource_id=None, name=None, type=None, workflow_id=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if logic_app_resource_id and not isinstance(logic_app_resource_id, str):
            raise TypeError("Expected argument 'logic_app_resource_id' to be a str")
        pulumi.set(__self__, "logic_app_resource_id", logic_app_resource_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if workflow_id and not isinstance(workflow_id, str):
            raise TypeError("Expected argument 'workflow_id' to be a str")
        pulumi.set(__self__, "workflow_id", workflow_id)

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        Etag of the action.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="logicAppResourceId")
    def logic_app_resource_id(self) -> str:
        """
        Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
        """
        return pulumi.get(self, "logic_app_resource_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="workflowId")
    def workflow_id(self) -> Optional[str]:
        """
        The name of the logic app's workflow.
        """
        return pulumi.get(self, "workflow_id")


class AwaitableGetAlertRuleActionResult(GetAlertRuleActionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlertRuleActionResult(
            etag=self.etag,
            logic_app_resource_id=self.logic_app_resource_id,
            name=self.name,
            type=self.type,
            workflow_id=self.workflow_id)


def get_alert_rule_action(name: Optional[str] = None,
                          resource_group_name: Optional[str] = None,
                          rule_id: Optional[str] = None,
                          workspace_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlertRuleActionResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Action ID
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    :param str rule_id: Alert rule ID
    :param str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleId'] = rule_id
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:operationalinsights/v20200101:getAlertRuleAction', __args__, opts=opts, typ=GetAlertRuleActionResult).value

    return AwaitableGetAlertRuleActionResult(
        etag=__ret__.etag,
        logic_app_resource_id=__ret__.logic_app_resource_id,
        name=__ret__.name,
        type=__ret__.type,
        workflow_id=__ret__.workflow_id)
