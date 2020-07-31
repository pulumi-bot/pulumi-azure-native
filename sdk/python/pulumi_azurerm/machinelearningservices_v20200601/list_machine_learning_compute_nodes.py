# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ListMachineLearningComputeNodesResult:
    """
    Compute node information related to a AmlCompute.
    """
    def __init__(__self__, compute_type=None, next_link=None, nodes=None):
        if compute_type and not isinstance(compute_type, str):
            raise TypeError("Expected argument 'compute_type' to be a str")
        __self__.compute_type = compute_type
        """
        The type of compute
        """
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        __self__.next_link = next_link
        """
        The continuation token.
        """
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        __self__.nodes = nodes
        """
        The collection of returned AmlCompute nodes details.
        """


class AwaitableListMachineLearningComputeNodesResult(ListMachineLearningComputeNodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListMachineLearningComputeNodesResult(
            compute_type=self.compute_type,
            next_link=self.next_link,
            nodes=self.nodes)


def list_machine_learning_compute_nodes(compute_name=None, resource_group_name=None, workspace_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str compute_name: Name of the Azure Machine Learning compute.
    :param str resource_group_name: Name of the resource group in which workspace is located.
    :param str workspace_name: Name of Azure Machine Learning workspace.
    """
    __args__ = dict()
    __args__['computeName'] = compute_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:machinelearningservices/v20200601:listMachineLearningComputeNodes', __args__, opts=opts).value

    return AwaitableListMachineLearningComputeNodesResult(
        compute_type=__ret__.get('computeType'),
        next_link=__ret__.get('nextLink'),
        nodes=__ret__.get('nodes'))