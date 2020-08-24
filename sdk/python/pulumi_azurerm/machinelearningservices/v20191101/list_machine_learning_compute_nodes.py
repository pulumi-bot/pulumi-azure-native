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
    'ListMachineLearningComputeNodesResult',
    'AwaitableListMachineLearningComputeNodesResult',
    'list_machine_learning_compute_nodes',
]

@pulumi.output_type
class ListMachineLearningComputeNodesResult:
    """
    Compute node information related to a AmlCompute.
    """
    def __init__(__self__, compute_type=None, next_link=None, nodes=None):
        if compute_type and not isinstance(compute_type, str):
            raise TypeError("Expected argument 'compute_type' to be a str")
        pulumi.set(__self__, "compute_type", compute_type)
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)

    @property
    @pulumi.getter(name="computeType")
    def compute_type(self) -> str:
        """
        The type of compute
        """
        return pulumi.get(self, "compute_type")

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> str:
        """
        The continuation token.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def nodes(self) -> List['outputs.AmlComputeNodeInformationResponseResult']:
        """
        The collection of returned AmlCompute nodes details.
        """
        return pulumi.get(self, "nodes")


class AwaitableListMachineLearningComputeNodesResult(ListMachineLearningComputeNodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListMachineLearningComputeNodesResult(
            compute_type=self.compute_type,
            next_link=self.next_link,
            nodes=self.nodes)


def list_machine_learning_compute_nodes(compute_name: Optional[str] = None,
                                        resource_group_name: Optional[str] = None,
                                        workspace_name: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListMachineLearningComputeNodesResult:
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
    __ret__ = pulumi.runtime.invoke('azurerm:machinelearningservices/v20191101:listMachineLearningComputeNodes', __args__, opts=opts, typ=ListMachineLearningComputeNodesResult).value

    return AwaitableListMachineLearningComputeNodesResult(
        compute_type=__ret__.compute_type,
        next_link=__ret__.next_link,
        nodes=__ret__.nodes)
