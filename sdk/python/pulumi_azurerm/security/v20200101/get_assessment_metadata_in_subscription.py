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
    'GetAssessmentMetadataInSubscriptionResult',
    'AwaitableGetAssessmentMetadataInSubscriptionResult',
    'get_assessment_metadata_in_subscription',
]

@pulumi.output_type
class GetAssessmentMetadataInSubscriptionResult:
    """
    Security assessment metadata
    """
    def __init__(__self__, assessment_type=None, category=None, description=None, display_name=None, implementation_effort=None, name=None, partner_data=None, policy_definition_id=None, preview=None, remediation_description=None, severity=None, threats=None, type=None, user_impact=None):
        if assessment_type and not isinstance(assessment_type, str):
            raise TypeError("Expected argument 'assessment_type' to be a str")
        pulumi.set(__self__, "assessment_type", assessment_type)
        if category and not isinstance(category, list):
            raise TypeError("Expected argument 'category' to be a list")
        pulumi.set(__self__, "category", category)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if implementation_effort and not isinstance(implementation_effort, str):
            raise TypeError("Expected argument 'implementation_effort' to be a str")
        pulumi.set(__self__, "implementation_effort", implementation_effort)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partner_data and not isinstance(partner_data, dict):
            raise TypeError("Expected argument 'partner_data' to be a dict")
        pulumi.set(__self__, "partner_data", partner_data)
        if policy_definition_id and not isinstance(policy_definition_id, str):
            raise TypeError("Expected argument 'policy_definition_id' to be a str")
        pulumi.set(__self__, "policy_definition_id", policy_definition_id)
        if preview and not isinstance(preview, bool):
            raise TypeError("Expected argument 'preview' to be a bool")
        pulumi.set(__self__, "preview", preview)
        if remediation_description and not isinstance(remediation_description, str):
            raise TypeError("Expected argument 'remediation_description' to be a str")
        pulumi.set(__self__, "remediation_description", remediation_description)
        if severity and not isinstance(severity, str):
            raise TypeError("Expected argument 'severity' to be a str")
        pulumi.set(__self__, "severity", severity)
        if threats and not isinstance(threats, list):
            raise TypeError("Expected argument 'threats' to be a list")
        pulumi.set(__self__, "threats", threats)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user_impact and not isinstance(user_impact, str):
            raise TypeError("Expected argument 'user_impact' to be a str")
        pulumi.set(__self__, "user_impact", user_impact)

    @property
    @pulumi.getter(name="assessmentType")
    def assessment_type(self) -> str:
        """
        BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
        """
        return pulumi.get(self, "assessment_type")

    @property
    @pulumi.getter
    def category(self) -> Optional[List[str]]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Human readable description of the assessment
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        User friendly display name of the assessment
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="implementationEffort")
    def implementation_effort(self) -> Optional[str]:
        """
        The implementation effort required to remediate this assessment
        """
        return pulumi.get(self, "implementation_effort")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partnerData")
    def partner_data(self) -> Optional['outputs.SecurityAssessmentMetadataPartnerDataResponse']:
        """
        Describes the partner that created the assessment
        """
        return pulumi.get(self, "partner_data")

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> str:
        """
        Azure resource ID of the policy definition that turns this assessment calculation on
        """
        return pulumi.get(self, "policy_definition_id")

    @property
    @pulumi.getter
    def preview(self) -> Optional[bool]:
        """
        True if this assessment is in preview release status
        """
        return pulumi.get(self, "preview")

    @property
    @pulumi.getter(name="remediationDescription")
    def remediation_description(self) -> Optional[str]:
        """
        Human readable description of what you should do to mitigate this security issue
        """
        return pulumi.get(self, "remediation_description")

    @property
    @pulumi.getter
    def severity(self) -> str:
        """
        The severity level of the assessment
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def threats(self) -> Optional[List[str]]:
        return pulumi.get(self, "threats")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userImpact")
    def user_impact(self) -> Optional[str]:
        """
        The user impact of the assessment
        """
        return pulumi.get(self, "user_impact")


class AwaitableGetAssessmentMetadataInSubscriptionResult(GetAssessmentMetadataInSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssessmentMetadataInSubscriptionResult(
            assessment_type=self.assessment_type,
            category=self.category,
            description=self.description,
            display_name=self.display_name,
            implementation_effort=self.implementation_effort,
            name=self.name,
            partner_data=self.partner_data,
            policy_definition_id=self.policy_definition_id,
            preview=self.preview,
            remediation_description=self.remediation_description,
            severity=self.severity,
            threats=self.threats,
            type=self.type,
            user_impact=self.user_impact)


def get_assessment_metadata_in_subscription(name: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAssessmentMetadataInSubscriptionResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The Assessment Key - Unique key for the assessment type
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:security/v20200101:getAssessmentMetadataInSubscription', __args__, opts=opts, typ=GetAssessmentMetadataInSubscriptionResult).value

    return AwaitableGetAssessmentMetadataInSubscriptionResult(
        assessment_type=__ret__.assessment_type,
        category=__ret__.category,
        description=__ret__.description,
        display_name=__ret__.display_name,
        implementation_effort=__ret__.implementation_effort,
        name=__ret__.name,
        partner_data=__ret__.partner_data,
        policy_definition_id=__ret__.policy_definition_id,
        preview=__ret__.preview,
        remediation_description=__ret__.remediation_description,
        severity=__ret__.severity,
        threats=__ret__.threats,
        type=__ret__.type,
        user_impact=__ret__.user_impact)
