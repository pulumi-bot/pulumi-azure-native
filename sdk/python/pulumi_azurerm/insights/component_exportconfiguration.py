# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ComponentExportconfiguration(pulumi.CustomResource):
    application_name: pulumi.Output[str]
    """
    The name of the Application Insights component.
    """
    container_name: pulumi.Output[str]
    """
    The name of the destination storage container.
    """
    destination_account_id: pulumi.Output[str]
    """
    The name of destination account.
    """
    destination_storage_location_id: pulumi.Output[str]
    """
    The destination account location ID.
    """
    destination_storage_subscription_id: pulumi.Output[str]
    """
    The destination storage account subscription ID.
    """
    destination_type: pulumi.Output[str]
    """
    The destination type.
    """
    export_id: pulumi.Output[str]
    """
    The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.
    """
    export_status: pulumi.Output[str]
    """
    This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
    """
    instrumentation_key: pulumi.Output[str]
    """
    The instrumentation key of the Application Insights component.
    """
    is_user_enabled: pulumi.Output[str]
    """
    This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
    """
    last_gap_time: pulumi.Output[str]
    """
    The last time the Continuous Export configuration started failing.
    """
    last_success_time: pulumi.Output[str]
    """
    The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
    """
    last_user_update: pulumi.Output[str]
    """
    Last time the Continuous Export configuration was updated.
    """
    notification_queue_enabled: pulumi.Output[str]
    """
    Deprecated
    """
    permanent_error_reason: pulumi.Output[str]
    """
    This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
    """
    record_types: pulumi.Output[str]
    """
    This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
    """
    resource_group: pulumi.Output[str]
    """
    The resource group of the Application Insights component.
    """
    storage_name: pulumi.Output[str]
    """
    The name of the destination storage account.
    """
    subscription_id: pulumi.Output[str]
    """
    The subscription of the Application Insights component.
    """
    def __init__(__self__, resource_name, opts=None, destination_account_id=None, destination_address=None, destination_storage_location_id=None, destination_storage_subscription_id=None, destination_type=None, is_enabled=None, notification_queue_enabled=None, notification_queue_uri=None, record_types=None, export_id=None, resource_group_name=None, resource_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Properties that define a Continuous Export configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_account_id: The name of destination storage account.
        :param pulumi.Input[str] destination_address: The SAS URL for the destination storage container. It must grant write permission.
        :param pulumi.Input[str] destination_storage_location_id: The location ID of the destination storage container.
        :param pulumi.Input[str] destination_storage_subscription_id: The subscription ID of the destination storage container.
        :param pulumi.Input[str] destination_type: The Continuous Export destination type. This has to be 'Blob'.
        :param pulumi.Input[str] is_enabled: Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'.
        :param pulumi.Input[str] notification_queue_enabled: Deprecated
        :param pulumi.Input[str] notification_queue_uri: Deprecated
        :param pulumi.Input[str] record_types: The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
        :param pulumi.Input[str] export_id: The Continuous Export configuration ID. This is unique within a Application Insights component.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] resource_name: The name of the Application Insights component resource.
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

            __props__['destination_account_id'] = destination_account_id
            __props__['destination_address'] = destination_address
            __props__['destination_storage_location_id'] = destination_storage_location_id
            __props__['destination_storage_subscription_id'] = destination_storage_subscription_id
            __props__['destination_type'] = destination_type
            __props__['is_enabled'] = is_enabled
            __props__['notification_queue_enabled'] = notification_queue_enabled
            __props__['notification_queue_uri'] = notification_queue_uri
            __props__['record_types'] = record_types
            if export_id is None:
                raise TypeError("Missing required property 'export_id'")
            __props__['export_id'] = export_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name is None:
                raise TypeError("Missing required property 'resource_name'")
            __props__['resource_name'] = resource_name
            __props__['application_name'] = None
            __props__['container_name'] = None
            __props__['export_id'] = None
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
        super(ComponentExportconfiguration, __self__).__init__(
            'azurerm:insights:ComponentExportconfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ComponentExportconfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ComponentExportconfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop