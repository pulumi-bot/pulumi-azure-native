# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['WorkloadNetworkDnsService']


class WorkloadNetworkDnsService(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_dns_zone: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 dns_service_id: Optional[pulumi.Input[str]] = None,
                 dns_service_ip: Optional[pulumi.Input[str]] = None,
                 fqdn_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_level: Optional[pulumi.Input[Union[str, 'DnsServiceLogLevelEnum']]] = None,
                 private_cloud_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 revision: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        NSX DNS Service

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_dns_zone: Default DNS zone of the DNS Service.
        :param pulumi.Input[str] display_name: Display name of the DNS Service.
        :param pulumi.Input[str] dns_service_id: NSX DNS Service identifier. Generally the same as the DNS Service's display name
        :param pulumi.Input[str] dns_service_ip: DNS service IP of the DNS Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fqdn_zones: FQDN zones of the DNS Service.
        :param pulumi.Input[Union[str, 'DnsServiceLogLevelEnum']] log_level: DNS Service log level.
        :param pulumi.Input[str] private_cloud_name: Name of the private cloud
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[float] revision: NSX revision number.
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

            __props__['default_dns_zone'] = default_dns_zone
            __props__['display_name'] = display_name
            __props__['dns_service_id'] = dns_service_id
            __props__['dns_service_ip'] = dns_service_ip
            __props__['fqdn_zones'] = fqdn_zones
            __props__['log_level'] = log_level
            if private_cloud_name is None and not opts.urn:
                raise TypeError("Missing required property 'private_cloud_name'")
            __props__['private_cloud_name'] = private_cloud_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['revision'] = revision
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['status'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:avs/v20210101preview:WorkloadNetworkDnsService"), pulumi.Alias(type_="azure-native:avs:WorkloadNetworkDnsService"), pulumi.Alias(type_="azure-nextgen:avs:WorkloadNetworkDnsService"), pulumi.Alias(type_="azure-native:avs/v20200717preview:WorkloadNetworkDnsService"), pulumi.Alias(type_="azure-nextgen:avs/v20200717preview:WorkloadNetworkDnsService")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WorkloadNetworkDnsService, __self__).__init__(
            'azure-native:avs/v20210101preview:WorkloadNetworkDnsService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkloadNetworkDnsService':
        """
        Get an existing WorkloadNetworkDnsService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["default_dns_zone"] = None
        __props__["display_name"] = None
        __props__["dns_service_ip"] = None
        __props__["fqdn_zones"] = None
        __props__["log_level"] = None
        __props__["name"] = None
        __props__["provisioning_state"] = None
        __props__["revision"] = None
        __props__["status"] = None
        __props__["type"] = None
        return WorkloadNetworkDnsService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultDnsZone")
    def default_dns_zone(self) -> pulumi.Output[Optional[str]]:
        """
        Default DNS zone of the DNS Service.
        """
        return pulumi.get(self, "default_dns_zone")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Display name of the DNS Service.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="dnsServiceIp")
    def dns_service_ip(self) -> pulumi.Output[Optional[str]]:
        """
        DNS service IP of the DNS Service.
        """
        return pulumi.get(self, "dns_service_ip")

    @property
    @pulumi.getter(name="fqdnZones")
    def fqdn_zones(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        FQDN zones of the DNS Service.
        """
        return pulumi.get(self, "fqdn_zones")

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> pulumi.Output[Optional[str]]:
        """
        DNS Service log level.
        """
        return pulumi.get(self, "log_level")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Output[Optional[float]]:
        """
        NSX revision number.
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        DNS Service status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

