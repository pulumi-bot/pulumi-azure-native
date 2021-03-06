# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['BillingRoleAssignmentByDepartment']


class BillingRoleAssignmentByDepartment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account_name: Optional[pulumi.Input[str]] = None,
                 billing_role_assignment_name: Optional[pulumi.Input[str]] = None,
                 department_name: Optional[pulumi.Input[str]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 principal_tenant_id: Optional[pulumi.Input[str]] = None,
                 role_definition_id: Optional[pulumi.Input[str]] = None,
                 user_authentication_type: Optional[pulumi.Input[str]] = None,
                 user_email_address: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The role assignment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account_name: The ID that uniquely identifies a billing account.
        :param pulumi.Input[str] billing_role_assignment_name: The ID that uniquely identifies a role assignment.
        :param pulumi.Input[str] department_name: The ID that uniquely identifies a department.
        :param pulumi.Input[str] principal_id: The principal id of the user to whom the role was assigned.
        :param pulumi.Input[str] principal_tenant_id: The principal tenant id of the user to whom the role was assigned.
        :param pulumi.Input[str] role_definition_id: The ID of the role definition.
        :param pulumi.Input[str] user_authentication_type: The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        :param pulumi.Input[str] user_email_address: The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if billing_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'billing_account_name'")
            __props__['billing_account_name'] = billing_account_name
            __props__['billing_role_assignment_name'] = billing_role_assignment_name
            if department_name is None and not opts.urn:
                raise TypeError("Missing required property 'department_name'")
            __props__['department_name'] = department_name
            __props__['principal_id'] = principal_id
            __props__['principal_tenant_id'] = principal_tenant_id
            __props__['role_definition_id'] = role_definition_id
            __props__['user_authentication_type'] = user_authentication_type
            __props__['user_email_address'] = user_email_address
            __props__['created_by_principal_id'] = None
            __props__['created_by_principal_tenant_id'] = None
            __props__['created_by_user_email_address'] = None
            __props__['created_on'] = None
            __props__['name'] = None
            __props__['scope'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:billing/v20191001preview:BillingRoleAssignmentByDepartment"), pulumi.Alias(type_="azure-native:billing:BillingRoleAssignmentByDepartment"), pulumi.Alias(type_="azure-nextgen:billing:BillingRoleAssignmentByDepartment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BillingRoleAssignmentByDepartment, __self__).__init__(
            'azure-native:billing/v20191001preview:BillingRoleAssignmentByDepartment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BillingRoleAssignmentByDepartment':
        """
        Get an existing BillingRoleAssignmentByDepartment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["created_by_principal_id"] = None
        __props__["created_by_principal_tenant_id"] = None
        __props__["created_by_user_email_address"] = None
        __props__["created_on"] = None
        __props__["name"] = None
        __props__["principal_id"] = None
        __props__["principal_tenant_id"] = None
        __props__["role_definition_id"] = None
        __props__["scope"] = None
        __props__["type"] = None
        __props__["user_authentication_type"] = None
        __props__["user_email_address"] = None
        return BillingRoleAssignmentByDepartment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdByPrincipalId")
    def created_by_principal_id(self) -> pulumi.Output[str]:
        """
        The principal Id of the user who created the role assignment.
        """
        return pulumi.get(self, "created_by_principal_id")

    @property
    @pulumi.getter(name="createdByPrincipalTenantId")
    def created_by_principal_tenant_id(self) -> pulumi.Output[str]:
        """
        The tenant Id of the user who created the role assignment.
        """
        return pulumi.get(self, "created_by_principal_tenant_id")

    @property
    @pulumi.getter(name="createdByUserEmailAddress")
    def created_by_user_email_address(self) -> pulumi.Output[str]:
        """
        The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
        """
        return pulumi.get(self, "created_by_user_email_address")

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> pulumi.Output[str]:
        """
        The date the role assignment was created.
        """
        return pulumi.get(self, "created_on")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Output[Optional[str]]:
        """
        The principal id of the user to whom the role was assigned.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalTenantId")
    def principal_tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The principal tenant id of the user to whom the role was assigned.
        """
        return pulumi.get(self, "principal_tenant_id")

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the role definition.
        """
        return pulumi.get(self, "role_definition_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        The scope at which the role was assigned.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAuthenticationType")
    def user_authentication_type(self) -> pulumi.Output[Optional[str]]:
        """
        The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        """
        return pulumi.get(self, "user_authentication_type")

    @property
    @pulumi.getter(name="userEmailAddress")
    def user_email_address(self) -> pulumi.Output[Optional[str]]:
        """
        The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        """
        return pulumi.get(self, "user_email_address")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

