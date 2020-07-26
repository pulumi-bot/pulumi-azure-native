# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Vault(pulumi.CustomResource):
    e_tag: pulumi.Output[str]
    """
    Optional ETag.
    """
    identity: pulumi.Output[dict]
    """
    Identity for the resource.
      * `principal_id` (`str`) - The principal ID of resource identity.
      * `tenant_id` (`str`) - The tenant ID of resource.
      * `type` (`str`) - The identity type.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name associated with the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the vault.
      * `private_endpoint_connections` (`list`) - List of private endpoint connection.
        * `id` (`str`) - Format of id subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.[Service]/{resource}/{resourceName}/privateEndpointConnections/{connectionName}.
        * `properties` (`dict`) - Private Endpoint Connection Response Properties.
          * `private_endpoint` (`dict`) - Gets or sets private endpoint associated with the private endpoint connection
            * `id` (`str`) - Gets or sets id

          * `private_link_service_connection_state` (`dict`) - Gets or sets private link service connection state
            * `action_required` (`str`) - Gets or sets actions required
            * `description` (`str`) - Gets or sets description
            * `status` (`str`) - Gets or sets the status

          * `provisioning_state` (`str`) - Gets or sets provisioning state of the private endpoint connection

      * `private_endpoint_state_for_backup` (`str`) - Private endpoint state for backup.
      * `private_endpoint_state_for_site_recovery` (`str`) - Private endpoint state for site recovery.
      * `provisioning_state` (`str`) - Provisioning State.
      * `upgrade_details` (`dict`) - Details for upgrading vault.
        * `end_time_utc` (`str`) - UTC time at which the upgrade operation has ended.
        * `last_updated_time_utc` (`str`) - UTC time at which the upgrade operation status was last updated.
        * `message` (`str`) - Message to the user containing information about the upgrade operation.
        * `operation_id` (`str`) - ID of the vault upgrade operation.
        * `previous_resource_id` (`str`) - Resource ID of the vault before the upgrade.
        * `start_time_utc` (`str`) - UTC time at which the upgrade operation has started.
        * `status` (`str`) - Status of the vault upgrade operation.
        * `trigger_type` (`str`) - The way the vault upgrade was triggered.
        * `upgraded_resource_id` (`str`) - Resource ID of the upgraded vault.
    """
    sku: pulumi.Output[dict]
    """
    Identifies the unique system identifier for each Azure resource.
      * `name` (`str`) - The Sku name.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
    """
    def __init__(__self__, resource_name, opts=None, e_tag=None, identity=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Resource information, as returned by the resource provider.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] e_tag: Optional ETag.
        :param pulumi.Input[dict] identity: Identity for the resource.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the recovery services vault.
        :param pulumi.Input[dict] properties: Properties of the vault.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[dict] sku: Identifies the unique system identifier for each Azure resource.
        :param pulumi.Input[dict] tags: Resource tags.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.

        The **properties** object supports the following:

          * `upgrade_details` (`pulumi.Input[dict]`) - Details for upgrading vault.

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - The Sku name.
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

            __props__['e_tag'] = e_tag
            __props__['identity'] = identity
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
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        super(Vault, __self__).__init__(
            'azurerm:recoveryservices:Vault',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Vault resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Vault(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
