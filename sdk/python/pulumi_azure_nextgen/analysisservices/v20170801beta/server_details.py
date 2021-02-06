# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ServerDetails']


class ServerDetails(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 as_administrators: Optional[pulumi.Input[pulumi.InputType['ServerAdministratorsArgs']]] = None,
                 backup_blob_container_uri: Optional[pulumi.Input[str]] = None,
                 gateway_details: Optional[pulumi.Input[pulumi.InputType['GatewayDetailsArgs']]] = None,
                 ip_v4_firewall_settings: Optional[pulumi.Input[pulumi.InputType['IPv4FirewallSettingsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_mode: Optional[pulumi.Input[int]] = None,
                 querypool_connection_mode: Optional[pulumi.Input['ConnectionMode']] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_monitor_mode: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ResourceSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents an instance of an Analysis Services resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServerAdministratorsArgs']] as_administrators: A collection of AS server administrators
        :param pulumi.Input[str] backup_blob_container_uri: The SAS container URI to the backup container.
        :param pulumi.Input[pulumi.InputType['GatewayDetailsArgs']] gateway_details: The gateway details configured for the AS server.
        :param pulumi.Input[pulumi.InputType['IPv4FirewallSettingsArgs']] ip_v4_firewall_settings: The firewall settings for the AS server.
        :param pulumi.Input[str] location: Location of the Analysis Services resource.
        :param pulumi.Input[int] managed_mode: The managed mode of the server (0 = not managed, 1 = managed).
        :param pulumi.Input['ConnectionMode'] querypool_connection_mode: How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
        :param pulumi.Input[str] resource_group_name: The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
        :param pulumi.Input[int] server_monitor_mode: The server monitor mode for AS server
        :param pulumi.Input[str] server_name: The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
        :param pulumi.Input[pulumi.InputType['ResourceSkuArgs']] sku: The SKU of the Analysis Services resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pairs of additional resource provisioning properties.
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

            __props__['as_administrators'] = as_administrators
            __props__['backup_blob_container_uri'] = backup_blob_container_uri
            __props__['gateway_details'] = gateway_details
            __props__['ip_v4_firewall_settings'] = ip_v4_firewall_settings
            __props__['location'] = location
            if managed_mode is None:
                managed_mode = 1
            __props__['managed_mode'] = managed_mode
            if querypool_connection_mode is None:
                querypool_connection_mode = 'All'
            __props__['querypool_connection_mode'] = querypool_connection_mode
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_monitor_mode is None:
                server_monitor_mode = 1
            __props__['server_monitor_mode'] = server_monitor_mode
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['server_full_name'] = None
            __props__['state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:analysisservices/latest:ServerDetails"), pulumi.Alias(type_="azure-nextgen:analysisservices/v20160516:ServerDetails"), pulumi.Alias(type_="azure-nextgen:analysisservices/v20170714:ServerDetails"), pulumi.Alias(type_="azure-nextgen:analysisservices/v20170801:ServerDetails")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ServerDetails, __self__).__init__(
            'azure-nextgen:analysisservices/v20170801beta:ServerDetails',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServerDetails':
        """
        Get an existing ServerDetails resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ServerDetails(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="asAdministrators")
    def as_administrators(self) -> pulumi.Output[Optional['outputs.ServerAdministratorsResponse']]:
        """
        A collection of AS server administrators
        """
        return pulumi.get(self, "as_administrators")

    @property
    @pulumi.getter(name="backupBlobContainerUri")
    def backup_blob_container_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The SAS container URI to the backup container.
        """
        return pulumi.get(self, "backup_blob_container_uri")

    @property
    @pulumi.getter(name="gatewayDetails")
    def gateway_details(self) -> pulumi.Output[Optional['outputs.GatewayDetailsResponse']]:
        """
        The gateway details configured for the AS server.
        """
        return pulumi.get(self, "gateway_details")

    @property
    @pulumi.getter(name="ipV4FirewallSettings")
    def ip_v4_firewall_settings(self) -> pulumi.Output[Optional['outputs.IPv4FirewallSettingsResponse']]:
        """
        The firewall settings for the AS server.
        """
        return pulumi.get(self, "ip_v4_firewall_settings")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Location of the Analysis Services resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedMode")
    def managed_mode(self) -> pulumi.Output[Optional[int]]:
        """
        The managed mode of the server (0 = not managed, 1 = managed).
        """
        return pulumi.get(self, "managed_mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Analysis Services resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="querypoolConnectionMode")
    def querypool_connection_mode(self) -> pulumi.Output[Optional[str]]:
        """
        How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
        """
        return pulumi.get(self, "querypool_connection_mode")

    @property
    @pulumi.getter(name="serverFullName")
    def server_full_name(self) -> pulumi.Output[str]:
        """
        The full name of the Analysis Services resource.
        """
        return pulumi.get(self, "server_full_name")

    @property
    @pulumi.getter(name="serverMonitorMode")
    def server_monitor_mode(self) -> pulumi.Output[Optional[int]]:
        """
        The server monitor mode for AS server
        """
        return pulumi.get(self, "server_monitor_mode")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.ResourceSkuResponse']:
        """
        The SKU of the Analysis Services resource.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value pairs of additional resource provisioning properties.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the Analysis Services resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

