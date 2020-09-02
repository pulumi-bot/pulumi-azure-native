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
    'ConfigurationProfileAssignmentComplianceResponse',
    'ConfigurationProfileAssignmentPropertiesResponse',
]

@pulumi.output_type
class ConfigurationProfileAssignmentComplianceResponse(dict):
    """
    The compliance status for the configuration profile assignment.
    """
    def __init__(__self__, *,
                 update_status: str):
        """
        The compliance status for the configuration profile assignment.
        :param str update_status: The state of compliance, which only appears in the response.
        """
        pulumi.set(__self__, "update_status", update_status)

    @property
    @pulumi.getter(name="updateStatus")
    def update_status(self) -> str:
        """
        The state of compliance, which only appears in the response.
        """
        return pulumi.get(self, "update_status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConfigurationProfileAssignmentPropertiesResponse(dict):
    """
    Automanage configuration profile assignment properties.
    """
    def __init__(__self__, *,
                 provisioning_status: str,
                 account_id: Optional[str] = None,
                 compliance: Optional['outputs.ConfigurationProfileAssignmentComplianceResponse'] = None,
                 configuration_profile: Optional[str] = None,
                 configuration_profile_preference_id: Optional[str] = None,
                 target_id: Optional[str] = None):
        """
        Automanage configuration profile assignment properties.
        :param str provisioning_status: The state of onboarding, which only appears in the response.
        :param str account_id: The Automanage account ARM Resource URI
        :param 'ConfigurationProfileAssignmentComplianceResponseArgs' compliance: The configuration setting for the configuration profile.
        :param str configuration_profile: A value indicating configuration profile.
        :param str configuration_profile_preference_id: The configuration profile custom preferences ARM resource URI
        :param str target_id: The target VM resource URI
        """
        pulumi.set(__self__, "provisioning_status", provisioning_status)
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if compliance is not None:
            pulumi.set(__self__, "compliance", compliance)
        if configuration_profile is not None:
            pulumi.set(__self__, "configuration_profile", configuration_profile)
        if configuration_profile_preference_id is not None:
            pulumi.set(__self__, "configuration_profile_preference_id", configuration_profile_preference_id)
        if target_id is not None:
            pulumi.set(__self__, "target_id", target_id)

    @property
    @pulumi.getter(name="provisioningStatus")
    def provisioning_status(self) -> str:
        """
        The state of onboarding, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_status")

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[str]:
        """
        The Automanage account ARM Resource URI
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def compliance(self) -> Optional['outputs.ConfigurationProfileAssignmentComplianceResponse']:
        """
        The configuration setting for the configuration profile.
        """
        return pulumi.get(self, "compliance")

    @property
    @pulumi.getter(name="configurationProfile")
    def configuration_profile(self) -> Optional[str]:
        """
        A value indicating configuration profile.
        """
        return pulumi.get(self, "configuration_profile")

    @property
    @pulumi.getter(name="configurationProfilePreferenceId")
    def configuration_profile_preference_id(self) -> Optional[str]:
        """
        The configuration profile custom preferences ARM resource URI
        """
        return pulumi.get(self, "configuration_profile_preference_id")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[str]:
        """
        The target VM resource URI
        """
        return pulumi.get(self, "target_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


