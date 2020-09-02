# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['RemediationAtManagementGroup']


class RemediationAtManagementGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filters: Optional[pulumi.Input[pulumi.InputType['RemediationFiltersArgs']]] = None,
                 management_group_id: Optional[pulumi.Input[str]] = None,
                 management_groups_namespace: Optional[pulumi.Input[str]] = None,
                 policy_assignment_id: Optional[pulumi.Input[str]] = None,
                 policy_definition_reference_id: Optional[pulumi.Input[str]] = None,
                 remediation_name: Optional[pulumi.Input[str]] = None,
                 resource_discovery_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The remediation definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RemediationFiltersArgs']] filters: The filters that will be applied to determine which resources to remediate.
        :param pulumi.Input[str] management_group_id: Management group ID.
        :param pulumi.Input[str] management_groups_namespace: The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
        :param pulumi.Input[str] policy_assignment_id: The resource ID of the policy assignment that should be remediated.
        :param pulumi.Input[str] policy_definition_reference_id: The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        :param pulumi.Input[str] remediation_name: The name of the remediation.
        :param pulumi.Input[str] resource_discovery_mode: The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
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

            __props__['filters'] = filters
            if management_group_id is None:
                raise TypeError("Missing required property 'management_group_id'")
            __props__['management_group_id'] = management_group_id
            if management_groups_namespace is None:
                raise TypeError("Missing required property 'management_groups_namespace'")
            __props__['management_groups_namespace'] = management_groups_namespace
            __props__['policy_assignment_id'] = policy_assignment_id
            __props__['policy_definition_reference_id'] = policy_definition_reference_id
            if remediation_name is None:
                raise TypeError("Missing required property 'remediation_name'")
            __props__['remediation_name'] = remediation_name
            __props__['resource_discovery_mode'] = resource_discovery_mode
            __props__['created_on'] = None
            __props__['deployment_status'] = None
            __props__['last_updated_on'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:policyinsights/latest:RemediationAtManagementGroup"), pulumi.Alias(type_="azurerm:policyinsights/v20180701preview:RemediationAtManagementGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RemediationAtManagementGroup, __self__).__init__(
            'azurerm:policyinsights/v20190701:RemediationAtManagementGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RemediationAtManagementGroup':
        """
        Get an existing RemediationAtManagementGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RemediationAtManagementGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> pulumi.Output[str]:
        """
        The time at which the remediation was created.
        """
        return pulumi.get(self, "created_on")

    @property
    @pulumi.getter(name="deploymentStatus")
    def deployment_status(self) -> pulumi.Output['outputs.RemediationDeploymentSummaryResponse']:
        """
        The deployment status summary for all deployments created by the remediation.
        """
        return pulumi.get(self, "deployment_status")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional['outputs.RemediationFiltersResponse']]:
        """
        The filters that will be applied to determine which resources to remediate.
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="lastUpdatedOn")
    def last_updated_on(self) -> pulumi.Output[str]:
        """
        The time at which the remediation was last updated.
        """
        return pulumi.get(self, "last_updated_on")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the remediation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyAssignmentId")
    def policy_assignment_id(self) -> pulumi.Output[Optional[str]]:
        """
        The resource ID of the policy assignment that should be remediated.
        """
        return pulumi.get(self, "policy_assignment_id")

    @property
    @pulumi.getter(name="policyDefinitionReferenceId")
    def policy_definition_reference_id(self) -> pulumi.Output[Optional[str]]:
        """
        The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        """
        return pulumi.get(self, "policy_definition_reference_id")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The status of the remediation.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceDiscoveryMode")
    def resource_discovery_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
        """
        return pulumi.get(self, "resource_discovery_mode")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the remediation.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

