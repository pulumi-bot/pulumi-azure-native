# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Server(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The Azure Active Directory identity of the server.
      * `principal_id` (`str`) - The Azure Active Directory principal id.
      * `tenant_id` (`str`) - The Azure Active Directory tenant id.
      * `type` (`str`) - The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource.
    """
    location: pulumi.Output[str]
    """
    The location the resource resides in.
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Properties of the server.
      * `administrator_login` (`str`) - The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
      * `byok_enforcement` (`str`) - Status showing whether the server data encryption is enabled with customer-managed keys.
      * `earliest_restore_date` (`str`) - Earliest restore point creation time (ISO8601 format)
      * `fully_qualified_domain_name` (`str`) - The fully qualified domain name of a server.
      * `infrastructure_encryption` (`str`) - Status showing whether the server enabled infrastructure encryption.
      * `master_server_id` (`str`) - The master server id of a replica server.
      * `minimal_tls_version` (`str`) - Enforce a minimal Tls version for the server.
      * `private_endpoint_connections` (`list`) - List of private endpoint connections on a server
        * `id` (`str`) - Resource ID of the Private Endpoint Connection.
        * `properties` (`dict`) - Private endpoint connection properties
          * `private_endpoint` (`dict`) - Private endpoint which the connection belongs to.
            * `id` (`str`) - Resource id of the private endpoint.

          * `private_link_service_connection_state` (`dict`) - Connection state of the private endpoint connection.
            * `actions_required` (`str`) - The actions required for private link service connection.
            * `description` (`str`) - The private link service connection description.
            * `status` (`str`) - The private link service connection status.

          * `provisioning_state` (`str`) - State of the private endpoint connection.

      * `public_network_access` (`str`) - Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
      * `replica_capacity` (`float`) - The maximum number of replicas that a master server can have.
      * `replication_role` (`str`) - The replication role of the server.
      * `ssl_enforcement` (`str`) - Enable ssl enforcement or not when connect to server.
      * `storage_profile` (`dict`) - Storage profile of a server.
        * `backup_retention_days` (`float`) - Backup retention days for the server.
        * `geo_redundant_backup` (`str`) - Enable Geo-redundant or not for server backup.
        * `storage_autogrow` (`str`) - Enable Storage Auto Grow.
        * `storage_mb` (`float`) - Max storage allowed for a server.

      * `user_visible_state` (`str`) - A state of a server that is visible to user.
      * `version` (`str`) - Server version.
    """
    sku: pulumi.Output[dict]
    """
    The SKU (pricing tier) of the server.
      * `capacity` (`float`) - The scale up/out capacity, representing server's compute units.
      * `family` (`str`) - The family of hardware.
      * `name` (`str`) - The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
      * `size` (`str`) - The size code, to be interpreted by resource as appropriate.
      * `tier` (`str`) - The tier of the particular SKU, e.g. Basic.
    """
    tags: pulumi.Output[dict]
    """
    Application-specific metadata in the form of key-value pairs.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, identity=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a server.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: The Azure Active Directory identity of the server.
        :param pulumi.Input[str] location: The location the resource resides in.
        :param pulumi.Input[str] name: The name of the server.
        :param pulumi.Input[dict] properties: Properties of the server.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[dict] sku: The SKU (pricing tier) of the server.
        :param pulumi.Input[dict] tags: Application-specific metadata in the form of key-value pairs.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource.

        The **properties** object supports the following:

          * `create_mode` (`pulumi.Input[str]`) - The mode to create a new server.
          * `infrastructure_encryption` (`pulumi.Input[str]`) - Status showing whether the server enabled infrastructure encryption.
          * `minimal_tls_version` (`pulumi.Input[str]`) - Enforce a minimal Tls version for the server.
          * `public_network_access` (`pulumi.Input[str]`) - Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
          * `ssl_enforcement` (`pulumi.Input[str]`) - Enable ssl enforcement or not when connect to server.
          * `storage_profile` (`pulumi.Input[dict]`) - Storage profile of a server.
            * `backup_retention_days` (`pulumi.Input[float]`) - Backup retention days for the server.
            * `geo_redundant_backup` (`pulumi.Input[str]`) - Enable Geo-redundant or not for server backup.
            * `storage_autogrow` (`pulumi.Input[str]`) - Enable Storage Auto Grow.
            * `storage_mb` (`pulumi.Input[float]`) - Max storage allowed for a server.

          * `version` (`pulumi.Input[str]`) - Server version.

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - The scale up/out capacity, representing server's compute units.
          * `family` (`pulumi.Input[str]`) - The family of hardware.
          * `name` (`pulumi.Input[str]`) - The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
          * `size` (`pulumi.Input[str]`) - The size code, to be interpreted by resource as appropriate.
          * `tier` (`pulumi.Input[str]`) - The tier of the particular SKU, e.g. Basic.
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

            __props__['identity'] = identity
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        super(Server, __self__).__init__(
            'azurerm:dbforpostgresql/v20171201:Server',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Server resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Server(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
