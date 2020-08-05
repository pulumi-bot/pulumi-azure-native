# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class BlobContainer(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Resource Etag.
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Properties of the blob container.
      * `default_encryption_scope` (`str`) - Default the container to use specified encryption scope for all writes.
      * `deleted` (`bool`) - Indicates whether the blob container was deleted.
      * `deleted_time` (`str`) - Blob container deletion time.
      * `deny_encryption_scope_override` (`bool`) - Block override of encryption scope from the container default.
      * `has_immutability_policy` (`bool`) - The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.
      * `has_legal_hold` (`bool`) - The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
      * `immutability_policy` (`dict`) - The ImmutabilityPolicy property of the container.
        * `etag` (`str`) - ImmutabilityPolicy Etag.
        * `properties` (`dict`) - The properties of an ImmutabilityPolicy of a blob container.
          * `allow_protected_append_writes` (`bool`) - This property can only be changed for unlocked time-based retention policies. When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with ExtendImmutabilityPolicy API
          * `immutability_period_since_creation_in_days` (`float`) - The immutability period for the blobs in the container since the policy creation, in days.
          * `state` (`str`) - The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.

        * `update_history` (`list`) - The ImmutabilityPolicy update history of the blob container.
          * `immutability_period_since_creation_in_days` (`float`) - The immutability period for the blobs in the container since the policy creation, in days.
          * `object_identifier` (`str`) - Returns the Object ID of the user who updated the ImmutabilityPolicy.
          * `tenant_id` (`str`) - Returns the Tenant ID that issued the token for the user who updated the ImmutabilityPolicy.
          * `timestamp` (`str`) - Returns the date and time the ImmutabilityPolicy was updated.
          * `update` (`str`) - The ImmutabilityPolicy update type of a blob container, possible values include: put, lock and extend.
          * `upn` (`str`) - Returns the User Principal Name of the user who updated the ImmutabilityPolicy.

      * `last_modified_time` (`str`) - Returns the date and time the container was last modified.
      * `lease_duration` (`str`) - Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
      * `lease_state` (`str`) - Lease state of the container.
      * `lease_status` (`str`) - The lease status of the container.
      * `legal_hold` (`dict`) - The LegalHold property of the container.
        * `has_legal_hold` (`bool`) - The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
        * `tags` (`list`) - The list of LegalHold tags of a blob container.
          * `object_identifier` (`str`) - Returns the Object ID of the user who added the tag.
          * `tag` (`str`) - The tag value.
          * `tenant_id` (`str`) - Returns the Tenant ID that issued the token for the user who added the tag.
          * `timestamp` (`str`) - Returns the date and time the tag was added.
          * `upn` (`str`) - Returns the User Principal Name of the user who added the tag.

      * `metadata` (`dict`) - A name-value pair to associate with the container as metadata.
      * `public_access` (`str`) - Specifies whether data in the container may be accessed publicly and the level of access.
      * `remaining_retention_days` (`float`) - Remaining retention days for soft deleted blob container.
      * `version` (`str`) - The version of the deleted blob container.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, default_encryption_scope=None, deny_encryption_scope_override=None, metadata=None, name=None, public_access=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Properties of the blob container, including Id, resource name, resource type, Etag.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        :param pulumi.Input[str] default_encryption_scope: Default the container to use specified encryption scope for all writes.
        :param pulumi.Input[bool] deny_encryption_scope_override: Block override of encryption scope from the container default.
        :param pulumi.Input[dict] metadata: A name-value pair to associate with the container as metadata.
        :param pulumi.Input[str] name: The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
        :param pulumi.Input[str] public_access: Specifies whether data in the container may be accessed publicly and the level of access.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['default_encryption_scope'] = default_encryption_scope
            __props__['deny_encryption_scope_override'] = deny_encryption_scope_override
            __props__['metadata'] = metadata
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['public_access'] = public_access
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(BlobContainer, __self__).__init__(
            'azurerm:storage/v20190601:BlobContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing BlobContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BlobContainer(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
