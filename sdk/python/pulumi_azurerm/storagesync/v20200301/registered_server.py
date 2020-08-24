# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['RegisteredServer']


class RegisteredServer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_version: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 last_heart_beat: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_certificate: Optional[pulumi.Input[str]] = None,
                 server_os_version: Optional[pulumi.Input[str]] = None,
                 server_role: Optional[pulumi.Input[str]] = None,
                 storage_sync_service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Registered Server resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_version: Registered Server Agent Version
        :param pulumi.Input[str] cluster_id: Registered Server clusterId
        :param pulumi.Input[str] cluster_name: Registered Server clusterName
        :param pulumi.Input[str] friendly_name: Friendly Name
        :param pulumi.Input[str] last_heart_beat: Registered Server last heart beat
        :param pulumi.Input[str] name: Registered Server serverId
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] server_certificate: Registered Server Certificate
        :param pulumi.Input[str] server_os_version: Registered Server OS Version
        :param pulumi.Input[str] server_role: Registered Server serverRole
        :param pulumi.Input[str] storage_sync_service_name: Name of Storage Sync Service resource.
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

            __props__['agent_version'] = agent_version
            __props__['cluster_id'] = cluster_id
            __props__['cluster_name'] = cluster_name
            __props__['friendly_name'] = friendly_name
            __props__['last_heart_beat'] = last_heart_beat
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['server_certificate'] = server_certificate
            __props__['server_os_version'] = server_os_version
            __props__['server_role'] = server_role
            if storage_sync_service_name is None:
                raise TypeError("Missing required property 'storage_sync_service_name'")
            __props__['storage_sync_service_name'] = storage_sync_service_name
            __props__['discovery_endpoint_uri'] = None
            __props__['last_operation_name'] = None
            __props__['last_workflow_id'] = None
            __props__['management_endpoint_uri'] = None
            __props__['monitoring_configuration'] = None
            __props__['monitoring_endpoint_uri'] = None
            __props__['provisioning_state'] = None
            __props__['resource_location'] = None
            __props__['server_id'] = None
            __props__['server_management_error_code'] = None
            __props__['service_location'] = None
            __props__['storage_sync_service_uid'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:storagesync/v20180402:RegisteredServer"), pulumi.Alias(type_="azurerm:storagesync/v20180701:RegisteredServer"), pulumi.Alias(type_="azurerm:storagesync/v20181001:RegisteredServer"), pulumi.Alias(type_="azurerm:storagesync/v20190201:RegisteredServer"), pulumi.Alias(type_="azurerm:storagesync/v20190301:RegisteredServer"), pulumi.Alias(type_="azurerm:storagesync/v20190601:RegisteredServer"), pulumi.Alias(type_="azurerm:storagesync/v20191001:RegisteredServer")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RegisteredServer, __self__).__init__(
            'azurerm:storagesync/v20200301:RegisteredServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegisteredServer':
        """
        Get an existing RegisteredServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RegisteredServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> Optional[str]:
        """
        Registered Server Agent Version
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        """
        Registered Server clusterId
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[str]:
        """
        Registered Server clusterName
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="discoveryEndpointUri")
    def discovery_endpoint_uri(self) -> Optional[str]:
        """
        Resource discoveryEndpointUri
        """
        return pulumi.get(self, "discovery_endpoint_uri")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[str]:
        """
        Friendly Name
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="lastHeartBeat")
    def last_heart_beat(self) -> Optional[str]:
        """
        Registered Server last heart beat
        """
        return pulumi.get(self, "last_heart_beat")

    @property
    @pulumi.getter(name="lastOperationName")
    def last_operation_name(self) -> Optional[str]:
        """
        Resource Last Operation Name
        """
        return pulumi.get(self, "last_operation_name")

    @property
    @pulumi.getter(name="lastWorkflowId")
    def last_workflow_id(self) -> Optional[str]:
        """
        Registered Server lastWorkflowId
        """
        return pulumi.get(self, "last_workflow_id")

    @property
    @pulumi.getter(name="managementEndpointUri")
    def management_endpoint_uri(self) -> Optional[str]:
        """
        Management Endpoint Uri
        """
        return pulumi.get(self, "management_endpoint_uri")

    @property
    @pulumi.getter(name="monitoringConfiguration")
    def monitoring_configuration(self) -> Optional[str]:
        """
        Monitoring Configuration
        """
        return pulumi.get(self, "monitoring_configuration")

    @property
    @pulumi.getter(name="monitoringEndpointUri")
    def monitoring_endpoint_uri(self) -> Optional[str]:
        """
        Telemetry Endpoint Uri
        """
        return pulumi.get(self, "monitoring_endpoint_uri")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Registered Server Provisioning State
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceLocation")
    def resource_location(self) -> Optional[str]:
        """
        Resource Location
        """
        return pulumi.get(self, "resource_location")

    @property
    @pulumi.getter(name="serverCertificate")
    def server_certificate(self) -> Optional[str]:
        """
        Registered Server Certificate
        """
        return pulumi.get(self, "server_certificate")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[str]:
        """
        Registered Server serverId
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="serverManagementErrorCode")
    def server_management_error_code(self) -> Optional[float]:
        """
        Registered Server Management Error Code
        """
        return pulumi.get(self, "server_management_error_code")

    @property
    @pulumi.getter(name="serverOSVersion")
    def server_os_version(self) -> Optional[str]:
        """
        Registered Server OS Version
        """
        return pulumi.get(self, "server_os_version")

    @property
    @pulumi.getter(name="serverRole")
    def server_role(self) -> Optional[str]:
        """
        Registered Server serverRole
        """
        return pulumi.get(self, "server_role")

    @property
    @pulumi.getter(name="serviceLocation")
    def service_location(self) -> Optional[str]:
        """
        Service Location
        """
        return pulumi.get(self, "service_location")

    @property
    @pulumi.getter(name="storageSyncServiceUid")
    def storage_sync_service_uid(self) -> Optional[str]:
        """
        Registered Server storageSyncServiceUid
        """
        return pulumi.get(self, "storage_sync_service_uid")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

