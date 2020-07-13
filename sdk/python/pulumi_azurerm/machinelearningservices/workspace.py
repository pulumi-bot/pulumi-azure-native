# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Workspace(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the resource.
      * `principal_id` (`str`) - The principal ID of resource identity.
      * `tenant_id` (`str`) - The tenant ID of resource.
      * `type` (`str`) - The identity type.
      * `user_assigned_identities` (`dict`) - The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
    """
    location: pulumi.Output[str]
    """
    Specifies the location of the resource.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the machine learning workspace.
      * `allow_public_access_when_behind_vnet` (`bool`) - The flag to indicate whether to allow public access when behind VNet.
      * `application_insights` (`str`) - ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
      * `container_registry` (`str`) - ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
      * `creation_time` (`str`) - The creation time of the machine learning workspace in ISO8601 format.
      * `description` (`str`) - The description of this workspace.
      * `discovery_url` (`str`) - Url for the discovery service to identify regional endpoints for machine learning experimentation services
      * `encryption` (`dict`) - The encryption settings of Azure ML workspace.
        * `key_vault_properties` (`dict`) - Customer Key vault properties.
          * `identity_client_id` (`str`) - For future use - The client id of the identity which will be used to access key vault.
          * `key_identifier` (`str`) - Key vault uri to access the encryption key.
          * `key_vault_arm_id` (`str`) - The ArmId of the keyVault where the customer owned encryption key is present.

        * `status` (`str`) - Indicates whether or not the encryption is enabled for the workspace.

      * `friendly_name` (`str`) - The friendly name for this workspace. This name in mutable
      * `hbi_workspace` (`bool`) - The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
      * `image_build_compute` (`str`) - The compute name for image build
      * `key_vault` (`str`) - ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
      * `notebook_info` (`dict`) - The notebook info of Azure ML workspace.
        * `fqdn` (`str`)
        * `notebook_preparation_error` (`dict`) - The error that occurs when preparing notebook.
          * `error_message` (`str`)
          * `status_code` (`float`)

        * `resource_id` (`str`) - the data plane resourceId that used to initialize notebook component

      * `private_endpoint_connections` (`list`) - The list of private endpoint connections in the workspace.
        * `id` (`str`) - Specifies the resource ID.
        * `identity` (`dict`) - The identity of the resource.
          * `principal_id` (`str`) - The principal ID of resource identity.
          * `tenant_id` (`str`) - The tenant ID of resource.
          * `type` (`str`) - The identity type.
          * `user_assigned_identities` (`dict`) - The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        * `location` (`str`) - Specifies the location of the resource.
        * `name` (`str`) - Specifies the name of the resource.
        * `properties` (`dict`) - Resource properties.
          * `private_endpoint` (`dict`) - The resource of private end point.
            * `id` (`str`) - The ARM identifier for Private Endpoint

          * `private_link_service_connection_state` (`dict`) - A collection of information about the state of the connection between service consumer and provider.
            * `actions_required` (`str`) - A message indicating if changes on the service provider require any updates on the consumer.
            * `description` (`str`) - The reason for approval/rejection of the connection.
            * `status` (`str`) - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.

          * `provisioning_state` (`str`) - The provisioning state of the private endpoint connection resource.

        * `sku` (`dict`) - The sku of the workspace.
          * `name` (`str`) - Name of the sku
          * `tier` (`str`) - Tier of the sku like Basic or Enterprise

        * `tags` (`dict`) - Contains resource tags defined as key/value pairs.
        * `type` (`str`) - Specifies the type of the resource.

      * `private_link_count` (`float`) - Count of private connections in the workspace
      * `provisioning_state` (`str`) - The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
      * `service_provisioned_resource_group` (`str`) - The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
      * `shared_private_link_resources` (`list`) - The list of shared private link resources in this workspace.
        * `name` (`str`) - Unique name of the private link.
        * `properties` (`dict`) - Resource properties.
          * `group_id` (`str`) - The private link resource group id.
          * `private_link_resource_id` (`str`) - The resource id that private link links to.
          * `request_message` (`str`) - Request message.
          * `status` (`str`) - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.

      * `storage_account` (`str`) - ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
      * `workspace_id` (`str`) - The immutable id associated with this workspace.
    """
    sku: pulumi.Output[dict]
    """
    The sku of the workspace.
      * `name` (`str`) - Name of the sku
      * `tier` (`str`) - Tier of the sku like Basic or Enterprise
    """
    tags: pulumi.Output[dict]
    """
    Contains resource tags defined as key/value pairs.
    """
    type: pulumi.Output[str]
    """
    Specifies the type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, identity=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        An object that represents a machine learning workspace.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: The identity of the resource.
        :param pulumi.Input[str] location: Specifies the location of the resource.
        :param pulumi.Input[str] name: Name of Azure Machine Learning workspace.
        :param pulumi.Input[dict] properties: The properties of the machine learning workspace.
        :param pulumi.Input[str] resource_group_name: Name of the resource group in which workspace is located.
        :param pulumi.Input[dict] sku: The sku of the workspace.
        :param pulumi.Input[dict] tags: Contains resource tags defined as key/value pairs.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.
          * `user_assigned_identities` (`pulumi.Input[dict]`) - The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        The **properties** object supports the following:

          * `allow_public_access_when_behind_vnet` (`pulumi.Input[bool]`) - The flag to indicate whether to allow public access when behind VNet.
          * `application_insights` (`pulumi.Input[str]`) - ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
          * `container_registry` (`pulumi.Input[str]`) - ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
          * `description` (`pulumi.Input[str]`) - The description of this workspace.
          * `discovery_url` (`pulumi.Input[str]`) - Url for the discovery service to identify regional endpoints for machine learning experimentation services
          * `encryption` (`pulumi.Input[dict]`) - The encryption settings of Azure ML workspace.
            * `key_vault_properties` (`pulumi.Input[dict]`) - Customer Key vault properties.
              * `identity_client_id` (`pulumi.Input[str]`) - For future use - The client id of the identity which will be used to access key vault.
              * `key_identifier` (`pulumi.Input[str]`) - Key vault uri to access the encryption key.
              * `key_vault_arm_id` (`pulumi.Input[str]`) - The ArmId of the keyVault where the customer owned encryption key is present.

            * `status` (`pulumi.Input[str]`) - Indicates whether or not the encryption is enabled for the workspace.

          * `friendly_name` (`pulumi.Input[str]`) - The friendly name for this workspace. This name in mutable
          * `hbi_workspace` (`pulumi.Input[bool]`) - The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
          * `image_build_compute` (`pulumi.Input[str]`) - The compute name for image build
          * `key_vault` (`pulumi.Input[str]`) - ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
          * `shared_private_link_resources` (`pulumi.Input[list]`) - The list of shared private link resources in this workspace.
            * `name` (`pulumi.Input[str]`) - Unique name of the private link.
            * `properties` (`pulumi.Input[dict]`) - Resource properties.
              * `group_id` (`pulumi.Input[str]`) - The private link resource group id.
              * `private_link_resource_id` (`pulumi.Input[str]`) - The resource id that private link links to.
              * `request_message` (`pulumi.Input[str]`) - Request message.
              * `status` (`pulumi.Input[str]`) - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.

          * `storage_account` (`pulumi.Input[str]`) - ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of the sku
          * `tier` (`pulumi.Input[str]`) - Tier of the sku like Basic or Enterprise
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['identity'] = identity
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
        super(Workspace, __self__).__init__(
            'azurerm:machinelearningservices:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Workspace(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
