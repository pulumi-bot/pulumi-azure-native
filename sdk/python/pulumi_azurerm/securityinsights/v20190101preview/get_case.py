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
    'GetCaseResult',
    'AwaitableGetCaseResult',
    'get_case',
]

@pulumi.output_type
class GetCaseResult:
    """
    Represents a case in Azure Security Insights.
    """
    def __init__(__self__, case_number=None, close_reason=None, closed_reason_text=None, created_time_utc=None, description=None, end_time_utc=None, etag=None, labels=None, last_comment=None, last_updated_time_utc=None, name=None, owner=None, related_alert_ids=None, severity=None, start_time_utc=None, status=None, tactics=None, title=None, total_comments=None, type=None):
        if case_number and not isinstance(case_number, float):
            raise TypeError("Expected argument 'case_number' to be a float")
        pulumi.set(__self__, "case_number", case_number)
        if close_reason and not isinstance(close_reason, str):
            raise TypeError("Expected argument 'close_reason' to be a str")
        pulumi.set(__self__, "close_reason", close_reason)
        if closed_reason_text and not isinstance(closed_reason_text, str):
            raise TypeError("Expected argument 'closed_reason_text' to be a str")
        pulumi.set(__self__, "closed_reason_text", closed_reason_text)
        if created_time_utc and not isinstance(created_time_utc, str):
            raise TypeError("Expected argument 'created_time_utc' to be a str")
        pulumi.set(__self__, "created_time_utc", created_time_utc)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if end_time_utc and not isinstance(end_time_utc, str):
            raise TypeError("Expected argument 'end_time_utc' to be a str")
        pulumi.set(__self__, "end_time_utc", end_time_utc)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if last_comment and not isinstance(last_comment, str):
            raise TypeError("Expected argument 'last_comment' to be a str")
        pulumi.set(__self__, "last_comment", last_comment)
        if last_updated_time_utc and not isinstance(last_updated_time_utc, str):
            raise TypeError("Expected argument 'last_updated_time_utc' to be a str")
        pulumi.set(__self__, "last_updated_time_utc", last_updated_time_utc)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, dict):
            raise TypeError("Expected argument 'owner' to be a dict")
        pulumi.set(__self__, "owner", owner)
        if related_alert_ids and not isinstance(related_alert_ids, list):
            raise TypeError("Expected argument 'related_alert_ids' to be a list")
        pulumi.set(__self__, "related_alert_ids", related_alert_ids)
        if severity and not isinstance(severity, str):
            raise TypeError("Expected argument 'severity' to be a str")
        pulumi.set(__self__, "severity", severity)
        if start_time_utc and not isinstance(start_time_utc, str):
            raise TypeError("Expected argument 'start_time_utc' to be a str")
        pulumi.set(__self__, "start_time_utc", start_time_utc)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tactics and not isinstance(tactics, list):
            raise TypeError("Expected argument 'tactics' to be a list")
        pulumi.set(__self__, "tactics", tactics)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if total_comments and not isinstance(total_comments, float):
            raise TypeError("Expected argument 'total_comments' to be a float")
        pulumi.set(__self__, "total_comments", total_comments)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="caseNumber")
    def case_number(self) -> float:
        """
        a sequential number
        """
        return pulumi.get(self, "case_number")

    @property
    @pulumi.getter(name="closeReason")
    def close_reason(self) -> Optional[str]:
        """
        The reason the case was closed
        """
        return pulumi.get(self, "close_reason")

    @property
    @pulumi.getter(name="closedReasonText")
    def closed_reason_text(self) -> Optional[str]:
        """
        the case close reason details
        """
        return pulumi.get(self, "closed_reason_text")

    @property
    @pulumi.getter(name="createdTimeUtc")
    def created_time_utc(self) -> str:
        """
        The time the case was created
        """
        return pulumi.get(self, "created_time_utc")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the case
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endTimeUtc")
    def end_time_utc(self) -> Optional[str]:
        """
        The end time of the case
        """
        return pulumi.get(self, "end_time_utc")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        Etag of the azure resource
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> Optional[List[str]]:
        """
        List of labels relevant to this case
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastComment")
    def last_comment(self) -> str:
        """
        the last comment in the case
        """
        return pulumi.get(self, "last_comment")

    @property
    @pulumi.getter(name="lastUpdatedTimeUtc")
    def last_updated_time_utc(self) -> str:
        """
        The last time the case was updated
        """
        return pulumi.get(self, "last_updated_time_utc")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> Optional['outputs.UserInfoResponse']:
        """
        Describes a user that the case is assigned to
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="relatedAlertIds")
    def related_alert_ids(self) -> List[str]:
        """
        List of related alert identifiers
        """
        return pulumi.get(self, "related_alert_ids")

    @property
    @pulumi.getter
    def severity(self) -> str:
        """
        The severity of the case
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter(name="startTimeUtc")
    def start_time_utc(self) -> str:
        """
        The start time of the case
        """
        return pulumi.get(self, "start_time_utc")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the case
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tactics(self) -> List[str]:
        """
        The tactics associated with case
        """
        return pulumi.get(self, "tactics")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        The title of the case
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="totalComments")
    def total_comments(self) -> float:
        """
        the number of total comments in the case
        """
        return pulumi.get(self, "total_comments")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetCaseResult(GetCaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCaseResult(
            case_number=self.case_number,
            close_reason=self.close_reason,
            closed_reason_text=self.closed_reason_text,
            created_time_utc=self.created_time_utc,
            description=self.description,
            end_time_utc=self.end_time_utc,
            etag=self.etag,
            labels=self.labels,
            last_comment=self.last_comment,
            last_updated_time_utc=self.last_updated_time_utc,
            name=self.name,
            owner=self.owner,
            related_alert_ids=self.related_alert_ids,
            severity=self.severity,
            start_time_utc=self.start_time_utc,
            status=self.status,
            tactics=self.tactics,
            title=self.title,
            total_comments=self.total_comments,
            type=self.type)


def get_case(case_id: Optional[str] = None,
             operational_insights_resource_provider: Optional[str] = None,
             resource_group_name: Optional[str] = None,
             workspace_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCaseResult:
    """
    Use this data source to access information about an existing resource.

    :param str case_id: Case ID
    :param str operational_insights_resource_provider: The namespace of workspaces resource provider- Microsoft.OperationalInsights.
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    :param str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['caseId'] = case_id
    __args__['operationalInsightsResourceProvider'] = operational_insights_resource_provider
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:securityinsights/v20190101preview:getCase', __args__, opts=opts, typ=GetCaseResult).value

    return AwaitableGetCaseResult(
        case_number=__ret__.case_number,
        close_reason=__ret__.close_reason,
        closed_reason_text=__ret__.closed_reason_text,
        created_time_utc=__ret__.created_time_utc,
        description=__ret__.description,
        end_time_utc=__ret__.end_time_utc,
        etag=__ret__.etag,
        labels=__ret__.labels,
        last_comment=__ret__.last_comment,
        last_updated_time_utc=__ret__.last_updated_time_utc,
        name=__ret__.name,
        owner=__ret__.owner,
        related_alert_ids=__ret__.related_alert_ids,
        severity=__ret__.severity,
        start_time_utc=__ret__.start_time_utc,
        status=__ret__.status,
        tactics=__ret__.tactics,
        title=__ret__.title,
        total_comments=__ret__.total_comments,
        type=__ret__.type)
