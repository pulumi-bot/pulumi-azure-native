# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetThreatIntelligenceIndicatorResult',
    'AwaitableGetThreatIntelligenceIndicatorResult',
    'get_threat_intelligence_indicator',
]

@pulumi.output_type
class GetThreatIntelligenceIndicatorResult:
    """
    Threat intelligence resource.
    """
    def __init__(__self__, etag=None, kind=None, name=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        Etag of the azure resource
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The kind of the entity.
        """
        return pulumi.get(self, "kind")

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


class AwaitableGetThreatIntelligenceIndicatorResult(GetThreatIntelligenceIndicatorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetThreatIntelligenceIndicatorResult(
            etag=self.etag,
            kind=self.kind,
            name=self.name,
            type=self.type)


def get_threat_intelligence_indicator(name: Optional[str] = None,
                                      operational_insights_resource_provider: Optional[str] = None,
                                      resource_group_name: Optional[str] = None,
                                      workspace_name: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetThreatIntelligenceIndicatorResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Threat Intelligence Identifier
    :param str operational_insights_resource_provider: The namespace of workspaces resource provider- Microsoft.OperationalInsights.
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    :param str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['operationalInsightsResourceProvider'] = operational_insights_resource_provider
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:securityinsights/v20190101preview:getThreatIntelligenceIndicator', __args__, opts=opts, typ=GetThreatIntelligenceIndicatorResult).value

    return AwaitableGetThreatIntelligenceIndicatorResult(
        etag=__ret__.etag,
        kind=__ret__.kind,
        name=__ret__.name,
        type=__ret__.type)
