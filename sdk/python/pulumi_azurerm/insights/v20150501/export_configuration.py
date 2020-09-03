# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['ExportConfiguration']


class ExportConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_account_id: Optional[pulumi.Input[str]] = None,
                 destination_address: Optional[pulumi.Input[str]] = None,
                 destination_storage_location_id: Optional[pulumi.Input[str]] = None,
                 destination_storage_subscription_id: Optional[pulumi.Input[str]] = None,
                 destination_type: Optional[pulumi.Input[str]] = None,
                 export_id: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[str]] = None,
                 notification_queue_enabled: Optional[pulumi.Input[str]] = None,
                 notification_queue_uri: Optional[pulumi.Input[str]] = None,
                 record_types: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Properties that define a Continuous Export configuration.

        ## Example Usage
        ### ExportConfigurationUpdate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        export_configuration = azurerm.insights.v20150501.ExportConfiguration("exportConfiguration",
            destination_account_id="/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.ClassicStorage/storageAccounts/mystorageblob",
            destination_address="https://mystorageblob.blob.core.windows.net/fchentest?sv=2015-04-05&sr=c&sig=token",
            destination_storage_location_id="eastus",
            destination_storage_subscription_id="subid",
            destination_type="Blob",
            export_id="uGOoki0jQsyEs3IdQ83Q4QsNr4=",
            is_enabled="true",
            notification_queue_enabled="false",
            notification_queue_uri="",
            record_types="Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability",
            resource_group_name="my-resource-group",
            resource_name="my-component")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_account_id: The name of destination storage account.
        :param pulumi.Input[str] destination_address: The SAS URL for the destination storage container. It must grant write permission.
        :param pulumi.Input[str] destination_storage_location_id: The location ID of the destination storage container.
        :param pulumi.Input[str] destination_storage_subscription_id: The subscription ID of the destination storage container.
        :param pulumi.Input[str] destination_type: The Continuous Export destination type. This has to be 'Blob'.
        :param pulumi.Input[str] export_id: The Continuous Export configuration ID. This is unique within a Application Insights component.
        :param pulumi.Input[str] is_enabled: Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'.
        :param pulumi.Input[str] notification_queue_enabled: Deprecated
        :param pulumi.Input[str] notification_queue_uri: Deprecated
        :param pulumi.Input[str] record_types: The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] resource_name_: The name of the Application Insights component resource.
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

            __props__['destination_account_id'] = destination_account_id
            __props__['destination_address'] = destination_address
            __props__['destination_storage_location_id'] = destination_storage_location_id
            __props__['destination_storage_subscription_id'] = destination_storage_subscription_id
            __props__['destination_type'] = destination_type
            if export_id is None:
                raise TypeError("Missing required property 'export_id'")
            __props__['export_id'] = export_id
            __props__['is_enabled'] = is_enabled
            __props__['notification_queue_enabled'] = notification_queue_enabled
            __props__['notification_queue_uri'] = notification_queue_uri
            __props__['record_types'] = record_types
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['application_name'] = None
            __props__['container_name'] = None
            __props__['export_status'] = None
            __props__['instrumentation_key'] = None
            __props__['is_user_enabled'] = None
            __props__['last_gap_time'] = None
            __props__['last_success_time'] = None
            __props__['last_user_update'] = None
            __props__['permanent_error_reason'] = None
            __props__['resource_group'] = None
            __props__['storage_name'] = None
            __props__['subscription_id'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:insights/latest:ExportConfiguration")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ExportConfiguration, __self__).__init__(
            'azurerm:insights/v20150501:ExportConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExportConfiguration':
        """
        Get an existing ExportConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExportConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Output[str]:
        """
        The name of the Application Insights component.
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Output[str]:
        """
        The name of the destination storage container.
        """
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter(name="destinationAccountId")
    def destination_account_id(self) -> pulumi.Output[str]:
        """
        The name of destination account.
        """
        return pulumi.get(self, "destination_account_id")

    @property
    @pulumi.getter(name="destinationStorageLocationId")
    def destination_storage_location_id(self) -> pulumi.Output[str]:
        """
        The destination account location ID.
        """
        return pulumi.get(self, "destination_storage_location_id")

    @property
    @pulumi.getter(name="destinationStorageSubscriptionId")
    def destination_storage_subscription_id(self) -> pulumi.Output[str]:
        """
        The destination storage account subscription ID.
        """
        return pulumi.get(self, "destination_storage_subscription_id")

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> pulumi.Output[str]:
        """
        The destination type.
        """
        return pulumi.get(self, "destination_type")

    @property
    @pulumi.getter(name="exportId")
    def export_id(self) -> pulumi.Output[str]:
        """
        The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.
        """
        return pulumi.get(self, "export_id")

    @property
    @pulumi.getter(name="exportStatus")
    def export_status(self) -> pulumi.Output[str]:
        """
        This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
        """
        return pulumi.get(self, "export_status")

    @property
    @pulumi.getter(name="instrumentationKey")
    def instrumentation_key(self) -> pulumi.Output[str]:
        """
        The instrumentation key of the Application Insights component.
        """
        return pulumi.get(self, "instrumentation_key")

    @property
    @pulumi.getter(name="isUserEnabled")
    def is_user_enabled(self) -> pulumi.Output[str]:
        """
        This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
        """
        return pulumi.get(self, "is_user_enabled")

    @property
    @pulumi.getter(name="lastGapTime")
    def last_gap_time(self) -> pulumi.Output[str]:
        """
        The last time the Continuous Export configuration started failing.
        """
        return pulumi.get(self, "last_gap_time")

    @property
    @pulumi.getter(name="lastSuccessTime")
    def last_success_time(self) -> pulumi.Output[str]:
        """
        The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
        """
        return pulumi.get(self, "last_success_time")

    @property
    @pulumi.getter(name="lastUserUpdate")
    def last_user_update(self) -> pulumi.Output[str]:
        """
        Last time the Continuous Export configuration was updated.
        """
        return pulumi.get(self, "last_user_update")

    @property
    @pulumi.getter(name="notificationQueueEnabled")
    def notification_queue_enabled(self) -> pulumi.Output[Optional[str]]:
        """
        Deprecated
        """
        return pulumi.get(self, "notification_queue_enabled")

    @property
    @pulumi.getter(name="permanentErrorReason")
    def permanent_error_reason(self) -> pulumi.Output[str]:
        """
        This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
        """
        return pulumi.get(self, "permanent_error_reason")

    @property
    @pulumi.getter(name="recordTypes")
    def record_types(self) -> pulumi.Output[Optional[str]]:
        """
        This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
        """
        return pulumi.get(self, "record_types")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> pulumi.Output[str]:
        """
        The resource group of the Application Insights component.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="storageName")
    def storage_name(self) -> pulumi.Output[str]:
        """
        The name of the destination storage account.
        """
        return pulumi.get(self, "storage_name")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Output[str]:
        """
        The subscription of the Application Insights component.
        """
        return pulumi.get(self, "subscription_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

