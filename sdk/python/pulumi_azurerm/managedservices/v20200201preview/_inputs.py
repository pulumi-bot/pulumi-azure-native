# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'AuthorizationArgs',
    'EligibleAuthorizationArgs',
    'JustInTimeAccessPolicyArgs',
    'PlanArgs',
    'RegistrationAssignmentPropertiesArgs',
    'RegistrationDefinitionPropertiesArgs',
]

@pulumi.input_type
class AuthorizationArgs:
    def __init__(__self__, *,
                 principal_id: pulumi.Input[str],
                 role_definition_id: pulumi.Input[str],
                 delegated_role_definition_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 principal_id_display_name: Optional[pulumi.Input[str]] = None):
        """
        Authorization tuple containing principal Id (of user/service principal/security group) and role definition id.
        :param pulumi.Input[str] principal_id: Principal Id of the security group/service principal/user that would be assigned permissions to the projected subscription
        :param pulumi.Input[str] role_definition_id: The role definition identifier. This role will define all the permissions that the security group/service principal/user must have on the projected subscription. This role cannot be an owner role.
        :param pulumi.Input[List[pulumi.Input[str]]] delegated_role_definition_ids: The delegatedRoleDefinitionIds field is required when the roleDefinitionId refers to the User Access Administrator Role. It is the list of role definition ids which define all the permissions that the user in the authorization can assign to other security groups/service principals/users.
        :param pulumi.Input[str] principal_id_display_name: Display name of the principal Id.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "role_definition_id", role_definition_id)
        if delegated_role_definition_ids is not None:
            pulumi.set(__self__, "delegated_role_definition_ids", delegated_role_definition_ids)
        if principal_id_display_name is not None:
            pulumi.set(__self__, "principal_id_display_name", principal_id_display_name)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Input[str]:
        """
        Principal Id of the security group/service principal/user that would be assigned permissions to the projected subscription
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> pulumi.Input[str]:
        """
        The role definition identifier. This role will define all the permissions that the security group/service principal/user must have on the projected subscription. This role cannot be an owner role.
        """
        return pulumi.get(self, "role_definition_id")

    @role_definition_id.setter
    def role_definition_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_definition_id", value)

    @property
    @pulumi.getter(name="delegatedRoleDefinitionIds")
    def delegated_role_definition_ids(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The delegatedRoleDefinitionIds field is required when the roleDefinitionId refers to the User Access Administrator Role. It is the list of role definition ids which define all the permissions that the user in the authorization can assign to other security groups/service principals/users.
        """
        return pulumi.get(self, "delegated_role_definition_ids")

    @delegated_role_definition_ids.setter
    def delegated_role_definition_ids(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "delegated_role_definition_ids", value)

    @property
    @pulumi.getter(name="principalIdDisplayName")
    def principal_id_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the principal Id.
        """
        return pulumi.get(self, "principal_id_display_name")

    @principal_id_display_name.setter
    def principal_id_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id_display_name", value)


@pulumi.input_type
class EligibleAuthorizationArgs:
    def __init__(__self__, *,
                 principal_id: pulumi.Input[str],
                 role_definition_id: pulumi.Input[str],
                 just_in_time_access_policy: Optional[pulumi.Input['JustInTimeAccessPolicyArgs']] = None,
                 principal_id_display_name: Optional[pulumi.Input[str]] = None):
        """
        Eligible authorization tuple containing principle Id (of user/service principal/security group), role definition id, and the just-in-time access setting.
        :param pulumi.Input[str] principal_id: Principal Id of the security group/service principal/user that would be delegated permissions to the projected subscription
        :param pulumi.Input[str] role_definition_id: The role definition identifier. This role will delegate all the permissions that the security group/service principal/user must have on the projected subscription. This role cannot be an owner role.
        :param pulumi.Input['JustInTimeAccessPolicyArgs'] just_in_time_access_policy: Just-in-time access policy setting.
        :param pulumi.Input[str] principal_id_display_name: Display name of the principal Id.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "role_definition_id", role_definition_id)
        if just_in_time_access_policy is not None:
            pulumi.set(__self__, "just_in_time_access_policy", just_in_time_access_policy)
        if principal_id_display_name is not None:
            pulumi.set(__self__, "principal_id_display_name", principal_id_display_name)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Input[str]:
        """
        Principal Id of the security group/service principal/user that would be delegated permissions to the projected subscription
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> pulumi.Input[str]:
        """
        The role definition identifier. This role will delegate all the permissions that the security group/service principal/user must have on the projected subscription. This role cannot be an owner role.
        """
        return pulumi.get(self, "role_definition_id")

    @role_definition_id.setter
    def role_definition_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_definition_id", value)

    @property
    @pulumi.getter(name="justInTimeAccessPolicy")
    def just_in_time_access_policy(self) -> Optional[pulumi.Input['JustInTimeAccessPolicyArgs']]:
        """
        Just-in-time access policy setting.
        """
        return pulumi.get(self, "just_in_time_access_policy")

    @just_in_time_access_policy.setter
    def just_in_time_access_policy(self, value: Optional[pulumi.Input['JustInTimeAccessPolicyArgs']]):
        pulumi.set(self, "just_in_time_access_policy", value)

    @property
    @pulumi.getter(name="principalIdDisplayName")
    def principal_id_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the principal Id.
        """
        return pulumi.get(self, "principal_id_display_name")

    @principal_id_display_name.setter
    def principal_id_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id_display_name", value)


@pulumi.input_type
class JustInTimeAccessPolicyArgs:
    def __init__(__self__, *,
                 multi_factor_auth_provider: pulumi.Input[str],
                 maximum_activation_duration: Optional[pulumi.Input[str]] = None):
        """
        Just-in-time access policy setting.
        :param pulumi.Input[str] multi_factor_auth_provider: MFA provider.
        :param pulumi.Input[str] maximum_activation_duration: Maximum access duration in ISO 8601 format.  The default value is "PT8H".
        """
        pulumi.set(__self__, "multi_factor_auth_provider", multi_factor_auth_provider)
        if maximum_activation_duration is not None:
            pulumi.set(__self__, "maximum_activation_duration", maximum_activation_duration)

    @property
    @pulumi.getter(name="multiFactorAuthProvider")
    def multi_factor_auth_provider(self) -> pulumi.Input[str]:
        """
        MFA provider.
        """
        return pulumi.get(self, "multi_factor_auth_provider")

    @multi_factor_auth_provider.setter
    def multi_factor_auth_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "multi_factor_auth_provider", value)

    @property
    @pulumi.getter(name="maximumActivationDuration")
    def maximum_activation_duration(self) -> Optional[pulumi.Input[str]]:
        """
        Maximum access duration in ISO 8601 format.  The default value is "PT8H".
        """
        return pulumi.get(self, "maximum_activation_duration")

    @maximum_activation_duration.setter
    def maximum_activation_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maximum_activation_duration", value)


@pulumi.input_type
class PlanArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 product: pulumi.Input[str],
                 publisher: pulumi.Input[str],
                 version: pulumi.Input[str]):
        """
        Plan details for the managed services.
        :param pulumi.Input[str] name: The plan name.
        :param pulumi.Input[str] product: The product code.
        :param pulumi.Input[str] publisher: The publisher ID.
        :param pulumi.Input[str] version: The plan's version.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "product", product)
        pulumi.set(__self__, "publisher", publisher)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The plan name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def product(self) -> pulumi.Input[str]:
        """
        The product code.
        """
        return pulumi.get(self, "product")

    @product.setter
    def product(self, value: pulumi.Input[str]):
        pulumi.set(self, "product", value)

    @property
    @pulumi.getter
    def publisher(self) -> pulumi.Input[str]:
        """
        The publisher ID.
        """
        return pulumi.get(self, "publisher")

    @publisher.setter
    def publisher(self, value: pulumi.Input[str]):
        pulumi.set(self, "publisher", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        The plan's version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class RegistrationAssignmentPropertiesArgs:
    def __init__(__self__, *,
                 registration_definition_id: pulumi.Input[str]):
        """
        Properties of a registration assignment.
        :param pulumi.Input[str] registration_definition_id: Fully qualified path of the registration definition.
        """
        pulumi.set(__self__, "registration_definition_id", registration_definition_id)

    @property
    @pulumi.getter(name="registrationDefinitionId")
    def registration_definition_id(self) -> pulumi.Input[str]:
        """
        Fully qualified path of the registration definition.
        """
        return pulumi.get(self, "registration_definition_id")

    @registration_definition_id.setter
    def registration_definition_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "registration_definition_id", value)


@pulumi.input_type
class RegistrationDefinitionPropertiesArgs:
    def __init__(__self__, *,
                 authorizations: pulumi.Input[List[pulumi.Input['AuthorizationArgs']]],
                 managed_by_tenant_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 eligible_authorizations: Optional[pulumi.Input[List[pulumi.Input['EligibleAuthorizationArgs']]]] = None,
                 registration_definition_name: Optional[pulumi.Input[str]] = None):
        """
        Properties of a registration definition.
        :param pulumi.Input[List[pulumi.Input['AuthorizationArgs']]] authorizations: Authorization tuple containing principal id of the user/security group or service principal and id of the build-in role.
        :param pulumi.Input[str] managed_by_tenant_id: Id of the managedBy tenant.
        :param pulumi.Input[str] description: Description of the registration definition.
        :param pulumi.Input[List[pulumi.Input['EligibleAuthorizationArgs']]] eligible_authorizations: Eligible PIM authorization tuple containing principal id of the user/security group or service principal, id of the built-in role, and just-in-time access policy setting
        :param pulumi.Input[str] registration_definition_name: Name of the registration definition.
        """
        pulumi.set(__self__, "authorizations", authorizations)
        pulumi.set(__self__, "managed_by_tenant_id", managed_by_tenant_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if eligible_authorizations is not None:
            pulumi.set(__self__, "eligible_authorizations", eligible_authorizations)
        if registration_definition_name is not None:
            pulumi.set(__self__, "registration_definition_name", registration_definition_name)

    @property
    @pulumi.getter
    def authorizations(self) -> pulumi.Input[List[pulumi.Input['AuthorizationArgs']]]:
        """
        Authorization tuple containing principal id of the user/security group or service principal and id of the build-in role.
        """
        return pulumi.get(self, "authorizations")

    @authorizations.setter
    def authorizations(self, value: pulumi.Input[List[pulumi.Input['AuthorizationArgs']]]):
        pulumi.set(self, "authorizations", value)

    @property
    @pulumi.getter(name="managedByTenantId")
    def managed_by_tenant_id(self) -> pulumi.Input[str]:
        """
        Id of the managedBy tenant.
        """
        return pulumi.get(self, "managed_by_tenant_id")

    @managed_by_tenant_id.setter
    def managed_by_tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "managed_by_tenant_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the registration definition.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="eligibleAuthorizations")
    def eligible_authorizations(self) -> Optional[pulumi.Input[List[pulumi.Input['EligibleAuthorizationArgs']]]]:
        """
        Eligible PIM authorization tuple containing principal id of the user/security group or service principal, id of the built-in role, and just-in-time access policy setting
        """
        return pulumi.get(self, "eligible_authorizations")

    @eligible_authorizations.setter
    def eligible_authorizations(self, value: Optional[pulumi.Input[List[pulumi.Input['EligibleAuthorizationArgs']]]]):
        pulumi.set(self, "eligible_authorizations", value)

    @property
    @pulumi.getter(name="registrationDefinitionName")
    def registration_definition_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the registration definition.
        """
        return pulumi.get(self, "registration_definition_name")

    @registration_definition_name.setter
    def registration_definition_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registration_definition_name", value)


