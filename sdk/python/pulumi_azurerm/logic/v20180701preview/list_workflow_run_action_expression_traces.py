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
    'ListWorkflowRunActionExpressionTracesResult',
    'AwaitableListWorkflowRunActionExpressionTracesResult',
    'list_workflow_run_action_expression_traces',
]

@pulumi.output_type
class ListWorkflowRunActionExpressionTracesResult:
    def __init__(__self__, inputs=None):
        if inputs and not isinstance(inputs, list):
            raise TypeError("Expected argument 'inputs' to be a list")
        pulumi.set(__self__, "inputs", inputs)

    @property
    @pulumi.getter
    def inputs(self) -> Optional[List['outputs.ExpressionRootResponseResult']]:
        return pulumi.get(self, "inputs")


class AwaitableListWorkflowRunActionExpressionTracesResult(ListWorkflowRunActionExpressionTracesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListWorkflowRunActionExpressionTracesResult(
            inputs=self.inputs)


def list_workflow_run_action_expression_traces(action_name: Optional[str] = None,
                                               resource_group_name: Optional[str] = None,
                                               run_name: Optional[str] = None,
                                               workflow_name: Optional[str] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListWorkflowRunActionExpressionTracesResult:
    """
    Use this data source to access information about an existing resource.

    :param str action_name: The workflow action name.
    :param str resource_group_name: The resource group name.
    :param str run_name: The workflow run name.
    :param str workflow_name: The workflow name.
    """
    __args__ = dict()
    __args__['actionName'] = action_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['runName'] = run_name
    __args__['workflowName'] = workflow_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:logic/v20180701preview:listWorkflowRunActionExpressionTraces', __args__, opts=opts, typ=ListWorkflowRunActionExpressionTracesResult).value

    return AwaitableListWorkflowRunActionExpressionTracesResult(
        inputs=__ret__.inputs)
