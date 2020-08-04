# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ReplicationFabric(pulumi.CustomResource):
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
    Fabric related data.
      * `bcdr_state` (`str`) - BCDR state of the fabric.
      * `custom_details` (`dict`) - Fabric specific settings.
        * `instance_type` (`str`) - Gets the class type. Overridden in derived classes.

      * `encryption_details` (`dict`) - Encryption details for the fabric.
        * `kek_cert_expiry_date` (`str`) - The key encryption key certificate expiry date.
        * `kek_cert_thumbprint` (`str`) - The key encryption key certificate thumbprint.
        * `kek_state` (`str`) - The key encryption key state for the Vmm.

      * `friendly_name` (`str`) - Friendly name of the fabric.
      * `health` (`str`) - Health of fabric.
      * `health_error_details` (`list`) - Fabric health error details.
        * `child_errors` (`list`) - The child health errors.
        * `creation_time_utc` (`str`) - Error creation time (UTC)
        * `entity_id` (`str`) - ID of the entity.
        * `error_code` (`str`) - Error code.
        * `error_level` (`str`) - Level of error.
        * `error_message` (`str`) - Error message.
        * `error_source` (`str`) - Source of error.
        * `error_type` (`str`) - Type of error.
        * `possible_causes` (`str`) - Possible causes of error.
        * `recommended_action` (`str`) - Recommended action to resolve error.
        * `recovery_provider_error_message` (`str`) - DRA error message.

      * `internal_identifier` (`str`) - Dra Registration Id.
      * `rollover_encryption_details` (`dict`) - Rollover encryption details for the fabric.
    """
    type: pulumi.Output[str]
    """
    Resource Type
    """
    def __init__(__self__, resource_name, opts=None, name=None, resource_group_name=None, resource_name_=None, __props__=None, __name__=None, __opts__=None):
        """
        Fabric definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the ASR fabric.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[str] resource_name_: The name of the recovery services vault.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['location'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(ReplicationFabric, __self__).__init__(
            'azurerm:recoveryservices/v20160810:ReplicationFabric',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ReplicationFabric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ReplicationFabric(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop