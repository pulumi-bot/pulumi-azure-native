# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetCaseRelationRelationResult',
    'AwaitableGetCaseRelationRelationResult',
    'get_case_relation_relation',
]

@pulumi.output_type
class GetCaseRelationRelationResult:
    """
    Represents a case relation
    """
    def __init__(__self__, bookmark_id=None, bookmark_name=None, case_identifier=None, etag=None, kind=None, name=None, relation_name=None, type=None):
        if bookmark_id and not isinstance(bookmark_id, str):
            raise TypeError("Expected argument 'bookmark_id' to be a str")
        pulumi.set(__self__, "bookmark_id", bookmark_id)
        if bookmark_name and not isinstance(bookmark_name, str):
            raise TypeError("Expected argument 'bookmark_name' to be a str")
        pulumi.set(__self__, "bookmark_name", bookmark_name)
        if case_identifier and not isinstance(case_identifier, str):
            raise TypeError("Expected argument 'case_identifier' to be a str")
        pulumi.set(__self__, "case_identifier", case_identifier)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if relation_name and not isinstance(relation_name, str):
            raise TypeError("Expected argument 'relation_name' to be a str")
        pulumi.set(__self__, "relation_name", relation_name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="bookmarkId")
    def bookmark_id(self) -> str:
        """
        The case related bookmark id
        """
        return pulumi.get(self, "bookmark_id")

    @property
    @pulumi.getter(name="bookmarkName")
    def bookmark_name(self) -> Optional[str]:
        """
        The case related bookmark name
        """
        return pulumi.get(self, "bookmark_name")

    @property
    @pulumi.getter(name="caseIdentifier")
    def case_identifier(self) -> str:
        """
        The case identifier
        """
        return pulumi.get(self, "case_identifier")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        ETag for relation
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The type of relation node
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
    @pulumi.getter(name="relationName")
    def relation_name(self) -> str:
        """
        Name of relation
        """
        return pulumi.get(self, "relation_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetCaseRelationRelationResult(GetCaseRelationRelationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCaseRelationRelationResult(
            bookmark_id=self.bookmark_id,
            bookmark_name=self.bookmark_name,
            case_identifier=self.case_identifier,
            etag=self.etag,
            kind=self.kind,
            name=self.name,
            relation_name=self.relation_name,
            type=self.type)


def get_case_relation_relation(case_id: Optional[str] = None,
                               operational_insights_resource_provider: Optional[str] = None,
                               relation_name: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               workspace_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCaseRelationRelationResult:
    """
    Use this data source to access information about an existing resource.

    :param str case_id: Case ID
    :param str operational_insights_resource_provider: The namespace of workspaces resource provider- Microsoft.OperationalInsights.
    :param str relation_name: Relation Name
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    :param str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['caseId'] = case_id
    __args__['operationalInsightsResourceProvider'] = operational_insights_resource_provider
    __args__['relationName'] = relation_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:securityinsights/v20190101preview:getCaseRelationRelation', __args__, opts=opts, typ=GetCaseRelationRelationResult).value

    return AwaitableGetCaseRelationRelationResult(
        bookmark_id=__ret__.bookmark_id,
        bookmark_name=__ret__.bookmark_name,
        case_identifier=__ret__.case_identifier,
        etag=__ret__.etag,
        kind=__ret__.kind,
        name=__ret__.name,
        relation_name=__ret__.relation_name,
        type=__ret__.type)
