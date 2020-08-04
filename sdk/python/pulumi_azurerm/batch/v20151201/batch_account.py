# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class BatchAccount(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The properties associated with the account.
      * `account_endpoint` (`str`) - The endpoint used by this account to interact with the Batch services.
      * `active_job_and_job_schedule_quota` (`float`) - The active job and job schedule quota for this Batch account.
      * `auto_storage` (`dict`) - The properties and status of any auto storage account associated with the account.
        * `last_key_sync` (`str`) - The UTC time at which storage keys were last synchronized with the Batch account.
        * `storage_account_id` (`str`) - The resource ID of the storage account to be used for auto storage account.

      * `core_quota` (`float`) - The core quota for this Batch account.
      * `pool_quota` (`float`) - The pool quota for this Batch account.
      * `provisioning_state` (`str`) - The provisioned state of the resource
    """
    tags: pulumi.Output[dict]
    """
    The tags of the resource
    """
    type: pulumi.Output[str]
    """
    The type of the resource
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Contains information about an Azure Batch account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The region in which to create the account.
        :param pulumi.Input[str] name: A name for the Batch account which must be unique within the region. Batch account names must be between 3 and 24 characters in length and must use only numbers and lowercase letters. This name is used as part of the DNS name that is used to access the Batch service in the region in which the account is created. For example: http://accountname.region.batch.azure.com/.
        :param pulumi.Input[dict] properties: The properties of the account.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the new Batch account.
        :param pulumi.Input[dict] tags: The user specified tags associated with the account.

        The **properties** object supports the following:

          * `auto_storage` (`pulumi.Input[dict]`) - The properties related to auto storage account.
            * `storage_account_id` (`pulumi.Input[str]`) - The resource ID of the storage account to be used for auto storage account.
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(BatchAccount, __self__).__init__(
            'azurerm:batch/v20151201:BatchAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing BatchAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BatchAccount(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
