# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ReplicationMigrationItem(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource Location
    """
    name: pulumi.Output[str]
    """
    Resource Name
    """
    properties: pulumi.Output[dict]
    """
    The migration item properties.
      * `allowed_operations` (`list`) - The allowed operations on the migration item, based on the current migration state of the item.
      * `current_job` (`dict`) - The current job details.
        * `job_id` (`str`) - The ARM Id of the job being executed.
        * `job_name` (`str`) - The job name.
        * `start_time` (`str`) - The start time of the job.

      * `health` (`str`) - The consolidated health.
      * `health_errors` (`list`) - The list of health errors.
        * `creation_time_utc` (`str`) - Error creation time (UTC)
        * `customer_resolvability` (`str`) - Value indicating whether the health error is customer resolvable.
        * `entity_id` (`str`) - ID of the entity.
        * `error_category` (`str`) - Category of error.
        * `error_code` (`str`) - Error code.
        * `error_id` (`str`) - The health error unique id.
        * `error_level` (`str`) - Level of error.
        * `error_message` (`str`) - Error message.
        * `error_source` (`str`) - Source of error.
        * `error_type` (`str`) - Type of error.
        * `inner_health_errors` (`list`) - The inner health errors. HealthError having a list of HealthError as child errors is problematic. InnerHealthError is used because this will prevent an infinite loop of structures when Hydra tries to auto-generate the contract. We are exposing the related health errors as inner health errors and all API consumers can utilize this in the same fashion as Exception -&gt; InnerException.
          * `creation_time_utc` (`str`) - Error creation time (UTC)
          * `entity_id` (`str`) - ID of the entity.
          * `error_category` (`str`) - Category of error.
          * `error_code` (`str`) - Error code.
          * `error_level` (`str`) - Level of error.
          * `error_message` (`str`) - Error message.
          * `error_source` (`str`) - Source of error.
          * `error_type` (`str`) - Type of error.
          * `possible_causes` (`str`) - Possible causes of error.
          * `recommended_action` (`str`) - Recommended action to resolve error.
          * `recovery_provider_error_message` (`str`) - DRA error message.
          * `summary_message` (`str`) - Summary message of the entity.

        * `possible_causes` (`str`) - Possible causes of error.
        * `recommended_action` (`str`) - Recommended action to resolve error.
        * `recovery_provider_error_message` (`str`) - DRA error message.
        * `summary_message` (`str`) - Summary message of the entity.

      * `machine_name` (`str`) - The on-premise virtual machine name.
      * `migration_state` (`str`) - The migration status.
      * `migration_state_description` (`str`) - The migration state description.
      * `policy_friendly_name` (`str`) - The name of policy governing this item.
      * `policy_id` (`str`) - The ARM Id of policy governing this item.
      * `provider_specific_details` (`dict`) - The migration provider custom settings.
        * `instance_type` (`str`) - Gets the instance type.

      * `recovery_services_provider_id` (`str`) - The recovery services provider ARM Id.
      * `test_migrate_state` (`str`) - The test migrate state.
      * `test_migrate_state_description` (`str`) - The test migrate state description.
    """
    type: pulumi.Output[str]
    """
    Resource Type
    """
    def __init__(__self__, resource_name, opts=None, fabric_name=None, name=None, properties=None, protection_container_name=None, resource_group_name=None, resource_name_=None, __props__=None, __name__=None, __opts__=None):
        """
        Migration item.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fabric_name: Fabric name.
        :param pulumi.Input[str] name: Migration item name.
        :param pulumi.Input[dict] properties: Enable migration input properties.
        :param pulumi.Input[str] protection_container_name: Protection container name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[str] resource_name_: The name of the recovery services vault.

        The **properties** object supports the following:

          * `policy_id` (`pulumi.Input[str]`) - The policy Id.
          * `provider_specific_details` (`pulumi.Input[dict]`) - The provider specific details.
            * `instance_type` (`pulumi.Input[str]`) - The class type.
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

            if fabric_name is None:
                raise TypeError("Missing required property 'fabric_name'")
            __props__['fabric_name'] = fabric_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if protection_container_name is None:
                raise TypeError("Missing required property 'protection_container_name'")
            __props__['protection_container_name'] = protection_container_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['location'] = None
            __props__['type'] = None
        super(ReplicationMigrationItem, __self__).__init__(
            'azurerm:recoveryservices/v20180710:ReplicationMigrationItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ReplicationMigrationItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ReplicationMigrationItem(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
