# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['SyncMember']


class SyncMember(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 sql_server_database_id: Optional[pulumi.Input[str]] = None,
                 sync_agent_id: Optional[pulumi.Input[str]] = None,
                 sync_direction: Optional[pulumi.Input[str]] = None,
                 sync_group_name: Optional[pulumi.Input[str]] = None,
                 sync_member_azure_database_resource_id: Optional[pulumi.Input[str]] = None,
                 sync_member_name: Optional[pulumi.Input[str]] = None,
                 use_private_link_connection: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An Azure SQL Database sync member.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Database name of the member database in the sync member.
        :param pulumi.Input[str] database_type: Database type of the sync member.
        :param pulumi.Input[str] password: Password of the member database in the sync member.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[str] server_name: Server name of the member database in the sync member
        :param pulumi.Input[str] sql_server_database_id: SQL Server database id of the sync member.
        :param pulumi.Input[str] sync_agent_id: ARM resource id of the sync agent in the sync member.
        :param pulumi.Input[str] sync_direction: Sync direction of the sync member.
        :param pulumi.Input[str] sync_group_name: The name of the sync group on which the sync member is hosted.
        :param pulumi.Input[str] sync_member_azure_database_resource_id: ARM resource id of the sync member logical database, for sync members in Azure.
        :param pulumi.Input[str] sync_member_name: The name of the sync member.
        :param pulumi.Input[bool] use_private_link_connection: Whether to use private link connection.
        :param pulumi.Input[str] user_name: User name of the member database in the sync member.
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

            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['database_type'] = database_type
            __props__['password'] = password
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            __props__['sql_server_database_id'] = sql_server_database_id
            __props__['sync_agent_id'] = sync_agent_id
            __props__['sync_direction'] = sync_direction
            if sync_group_name is None:
                raise TypeError("Missing required property 'sync_group_name'")
            __props__['sync_group_name'] = sync_group_name
            __props__['sync_member_azure_database_resource_id'] = sync_member_azure_database_resource_id
            if sync_member_name is None:
                raise TypeError("Missing required property 'sync_member_name'")
            __props__['sync_member_name'] = sync_member_name
            __props__['use_private_link_connection'] = use_private_link_connection
            __props__['user_name'] = user_name
            __props__['name'] = None
            __props__['sync_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:sql/latest:SyncMember"), pulumi.Alias(type_="azurerm:sql/v20150501preview:SyncMember")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SyncMember, __self__).__init__(
            'azurerm:sql/v20190601preview:SyncMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SyncMember':
        """
        Get an existing SyncMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SyncMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[Optional[str]]:
        """
        Database name of the member database in the sync member.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="databaseType")
    def database_type(self) -> pulumi.Output[Optional[str]]:
        """
        Database type of the sync member.
        """
        return pulumi.get(self, "database_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password of the member database in the sync member.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[Optional[str]]:
        """
        Server name of the member database in the sync member
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="sqlServerDatabaseId")
    def sql_server_database_id(self) -> pulumi.Output[Optional[str]]:
        """
        SQL Server database id of the sync member.
        """
        return pulumi.get(self, "sql_server_database_id")

    @property
    @pulumi.getter(name="syncAgentId")
    def sync_agent_id(self) -> pulumi.Output[Optional[str]]:
        """
        ARM resource id of the sync agent in the sync member.
        """
        return pulumi.get(self, "sync_agent_id")

    @property
    @pulumi.getter(name="syncDirection")
    def sync_direction(self) -> pulumi.Output[Optional[str]]:
        """
        Sync direction of the sync member.
        """
        return pulumi.get(self, "sync_direction")

    @property
    @pulumi.getter(name="syncMemberAzureDatabaseResourceId")
    def sync_member_azure_database_resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        ARM resource id of the sync member logical database, for sync members in Azure.
        """
        return pulumi.get(self, "sync_member_azure_database_resource_id")

    @property
    @pulumi.getter(name="syncState")
    def sync_state(self) -> pulumi.Output[str]:
        """
        Sync state of the sync member.
        """
        return pulumi.get(self, "sync_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usePrivateLinkConnection")
    def use_private_link_connection(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to use private link connection.
        """
        return pulumi.get(self, "use_private_link_connection")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[Optional[str]]:
        """
        User name of the member database in the sync member.
        """
        return pulumi.get(self, "user_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

