# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['Component']


class Component(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_type: Optional[pulumi.Input[str]] = None,
                 disable_ip_masking: Optional[pulumi.Input[bool]] = None,
                 flow_type: Optional[pulumi.Input[str]] = None,
                 hockey_app_id: Optional[pulumi.Input[str]] = None,
                 immediate_purge_data_on30_days: Optional[pulumi.Input[bool]] = None,
                 ingestion_mode: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 request_source: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 retention_in_days: Optional[pulumi.Input[float]] = None,
                 sampling_percentage: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An Application Insights component definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_type: Type of application being monitored.
        :param pulumi.Input[bool] disable_ip_masking: Disable IP masking.
        :param pulumi.Input[str] flow_type: Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
        :param pulumi.Input[str] hockey_app_id: The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
        :param pulumi.Input[bool] immediate_purge_data_on30_days: Purge data immediately after 30 days.
        :param pulumi.Input[str] ingestion_mode: Indicates the flow of the ingestion.
        :param pulumi.Input[str] kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] request_source: Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] resource_name_: The name of the Application Insights component resource.
        :param pulumi.Input[float] retention_in_days: Retention period in days.
        :param pulumi.Input[float] sampling_percentage: Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
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

            if application_type is None:
                raise TypeError("Missing required property 'application_type'")
            __props__['application_type'] = application_type
            __props__['disable_ip_masking'] = disable_ip_masking
            __props__['flow_type'] = flow_type
            __props__['hockey_app_id'] = hockey_app_id
            __props__['immediate_purge_data_on30_days'] = immediate_purge_data_on30_days
            __props__['ingestion_mode'] = ingestion_mode
            if kind is None:
                raise TypeError("Missing required property 'kind'")
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['request_source'] = request_source
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['retention_in_days'] = retention_in_days
            __props__['sampling_percentage'] = sampling_percentage
            __props__['tags'] = tags
            __props__['app_id'] = None
            __props__['application_id'] = None
            __props__['connection_string'] = None
            __props__['creation_date'] = None
            __props__['hockey_app_token'] = None
            __props__['instrumentation_key'] = None
            __props__['name'] = None
            __props__['private_link_scoped_resources'] = None
            __props__['provisioning_state'] = None
            __props__['tenant_id'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:insights/v20150501:Component"), pulumi.Alias(type_="azurerm:insights/v20180501preview:Component"), pulumi.Alias(type_="azurerm:insights/v20200202preview:Component")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Component, __self__).__init__(
            'azurerm:insights/latest:Component',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Component':
        """
        Get an existing Component resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Component(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        Application Insights Unique ID for your Application.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The unique ID of your application. This field mirrors the 'Name' field and cannot be changed.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="applicationType")
    def application_type(self) -> pulumi.Output[str]:
        """
        Type of application being monitored.
        """
        return pulumi.get(self, "application_type")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        Application Insights component connection string.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        Creation Date for the Application Insights component, in ISO 8601 format.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="disableIpMasking")
    def disable_ip_masking(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable IP masking.
        """
        return pulumi.get(self, "disable_ip_masking")

    @property
    @pulumi.getter(name="flowType")
    def flow_type(self) -> pulumi.Output[Optional[str]]:
        """
        Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
        """
        return pulumi.get(self, "flow_type")

    @property
    @pulumi.getter(name="hockeyAppId")
    def hockey_app_id(self) -> pulumi.Output[Optional[str]]:
        """
        The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
        """
        return pulumi.get(self, "hockey_app_id")

    @property
    @pulumi.getter(name="hockeyAppToken")
    def hockey_app_token(self) -> pulumi.Output[str]:
        """
        Token used to authenticate communications with between Application Insights and HockeyApp.
        """
        return pulumi.get(self, "hockey_app_token")

    @property
    @pulumi.getter(name="immediatePurgeDataOn30Days")
    def immediate_purge_data_on30_days(self) -> pulumi.Output[Optional[bool]]:
        """
        Purge data immediately after 30 days.
        """
        return pulumi.get(self, "immediate_purge_data_on30_days")

    @property
    @pulumi.getter(name="ingestionMode")
    def ingestion_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates the flow of the ingestion.
        """
        return pulumi.get(self, "ingestion_mode")

    @property
    @pulumi.getter(name="instrumentationKey")
    def instrumentation_key(self) -> pulumi.Output[str]:
        """
        Application Insights Instrumentation key. A read-only value that applications can use to identify the destination for all telemetry sent to Azure Application Insights. This value will be supplied upon construction of each new Application Insights component.
        """
        return pulumi.get(self, "instrumentation_key")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateLinkScopedResources")
    def private_link_scoped_resources(self) -> pulumi.Output[List['outputs.PrivateLinkScopedResourceResponse']]:
        """
        List of linked private link scope resources.
        """
        return pulumi.get(self, "private_link_scoped_resources")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Current state of this component: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="requestSource")
    def request_source(self) -> pulumi.Output[Optional[str]]:
        """
        Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
        """
        return pulumi.get(self, "request_source")

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> pulumi.Output[Optional[float]]:
        """
        Retention period in days.
        """
        return pulumi.get(self, "retention_in_days")

    @property
    @pulumi.getter(name="samplingPercentage")
    def sampling_percentage(self) -> pulumi.Output[Optional[float]]:
        """
        Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
        """
        return pulumi.get(self, "sampling_percentage")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        Azure Tenant Id.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

