# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class RemediationAtResourceGroup(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the remediation.
    """
    properties: pulumi.Output[dict]
    """
    Properties for the remediation.
      * `created_on` (`str`) - The time at which the remediation was created.
      * `deployment_status` (`dict`) - The deployment status summary for all deployments created by the remediation.
        * `failed_deployments` (`float`) - The number of deployments required by the remediation that have failed.
        * `successful_deployments` (`float`) - The number of deployments required by the remediation that have succeeded.
        * `total_deployments` (`float`) - The number of deployments required by the remediation.

      * `filters` (`dict`) - The filters that will be applied to determine which resources to remediate.
        * `locations` (`list`) - The resource locations that will be remediated.

      * `last_updated_on` (`str`) - The time at which the remediation was last updated.
      * `policy_assignment_id` (`str`) - The resource ID of the policy assignment that should be remediated.
      * `policy_definition_reference_id` (`str`) - The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
      * `provisioning_state` (`str`) - The status of the remediation.
      * `resource_discovery_mode` (`str`) - The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
    """
    type: pulumi.Output[str]
    """
    The type of the remediation.
    """
    def __init__(__self__, resource_name, opts=None, filters=None, name=None, policy_assignment_id=None, policy_definition_reference_id=None, resource_discovery_mode=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The remediation definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] filters: The filters that will be applied to determine which resources to remediate.
        :param pulumi.Input[str] name: The name of the remediation.
        :param pulumi.Input[str] policy_assignment_id: The resource ID of the policy assignment that should be remediated.
        :param pulumi.Input[str] policy_definition_reference_id: The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        :param pulumi.Input[str] resource_discovery_mode: The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
        :param pulumi.Input[str] resource_group_name: Resource group name.

        The **filters** object supports the following:

          * `locations` (`pulumi.Input[list]`) - The resource locations that will be remediated.
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
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['policy_assignment_id'] = policy_assignment_id
            __props__['policy_definition_reference_id'] = policy_definition_reference_id
            __props__['resource_discovery_mode'] = resource_discovery_mode
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['properties'] = None
            __props__['type'] = None
        super(RemediationAtResourceGroup, __self__).__init__(
            'azurerm:policyinsights/v20190701:RemediationAtResourceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RemediationAtResourceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RemediationAtResourceGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
