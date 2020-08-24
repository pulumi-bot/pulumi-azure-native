# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'CloudToDevicePropertiesResponse',
    'EventHubPropertiesResponse',
    'FeedbackPropertiesResponse',
    'IotHubPropertiesResponse',
    'IotHubSkuInfoResponse',
    'IpFilterRuleResponse',
    'MessagingEndpointPropertiesResponse',
    'OperationsMonitoringPropertiesResponse',
    'SharedAccessSignatureAuthorizationRuleResponse',
    'StorageEndpointPropertiesResponse',
]

@pulumi.output_type
class CloudToDevicePropertiesResponse(dict):
    """
    The IoT hub cloud-to-device messaging properties.
    """
    def __init__(__self__, *,
                 default_ttl_as_iso8601: Optional[str] = None,
                 feedback: Optional['outputs.FeedbackPropertiesResponse'] = None,
                 max_delivery_count: Optional[float] = None):
        """
        The IoT hub cloud-to-device messaging properties.
        :param str default_ttl_as_iso8601: The default time to live for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        :param 'FeedbackPropertiesResponseArgs' feedback: The properties of the feedback queue for cloud-to-device messages.
        :param float max_delivery_count: The max delivery count for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        """
        if default_ttl_as_iso8601 is not None:
            pulumi.set(__self__, "default_ttl_as_iso8601", default_ttl_as_iso8601)
        if feedback is not None:
            pulumi.set(__self__, "feedback", feedback)
        if max_delivery_count is not None:
            pulumi.set(__self__, "max_delivery_count", max_delivery_count)

    @property
    @pulumi.getter(name="defaultTtlAsIso8601")
    def default_ttl_as_iso8601(self) -> Optional[str]:
        """
        The default time to live for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        """
        return pulumi.get(self, "default_ttl_as_iso8601")

    @property
    @pulumi.getter
    def feedback(self) -> Optional['outputs.FeedbackPropertiesResponse']:
        """
        The properties of the feedback queue for cloud-to-device messages.
        """
        return pulumi.get(self, "feedback")

    @property
    @pulumi.getter(name="maxDeliveryCount")
    def max_delivery_count(self) -> Optional[float]:
        """
        The max delivery count for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        """
        return pulumi.get(self, "max_delivery_count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EventHubPropertiesResponse(dict):
    """
    The properties of the provisioned Event Hub-compatible endpoint used by the IoT hub.
    """
    def __init__(__self__, *,
                 endpoint: str,
                 partition_ids: List[str],
                 path: str,
                 partition_count: Optional[float] = None,
                 retention_time_in_days: Optional[float] = None):
        """
        The properties of the provisioned Event Hub-compatible endpoint used by the IoT hub.
        :param str endpoint: The Event Hub-compatible endpoint.
        :param List[str] partition_ids: The partition ids in the Event Hub-compatible endpoint.
        :param str path: The Event Hub-compatible name.
        :param float partition_count: The number of partitions for receiving device-to-cloud messages in the Event Hub-compatible endpoint. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#device-to-cloud-messages.
        :param float retention_time_in_days: The retention time for device-to-cloud messages in days. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#device-to-cloud-messages
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "partition_ids", partition_ids)
        pulumi.set(__self__, "path", path)
        if partition_count is not None:
            pulumi.set(__self__, "partition_count", partition_count)
        if retention_time_in_days is not None:
            pulumi.set(__self__, "retention_time_in_days", retention_time_in_days)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The Event Hub-compatible endpoint.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="partitionIds")
    def partition_ids(self) -> List[str]:
        """
        The partition ids in the Event Hub-compatible endpoint.
        """
        return pulumi.get(self, "partition_ids")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The Event Hub-compatible name.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> Optional[float]:
        """
        The number of partitions for receiving device-to-cloud messages in the Event Hub-compatible endpoint. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#device-to-cloud-messages.
        """
        return pulumi.get(self, "partition_count")

    @property
    @pulumi.getter(name="retentionTimeInDays")
    def retention_time_in_days(self) -> Optional[float]:
        """
        The retention time for device-to-cloud messages in days. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#device-to-cloud-messages
        """
        return pulumi.get(self, "retention_time_in_days")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FeedbackPropertiesResponse(dict):
    """
    The properties of the feedback queue for cloud-to-device messages.
    """
    def __init__(__self__, *,
                 lock_duration_as_iso8601: Optional[str] = None,
                 max_delivery_count: Optional[float] = None,
                 ttl_as_iso8601: Optional[str] = None):
        """
        The properties of the feedback queue for cloud-to-device messages.
        :param str lock_duration_as_iso8601: The lock duration for the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        :param float max_delivery_count: The number of times the IoT hub attempts to deliver a message on the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        :param str ttl_as_iso8601: The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        """
        if lock_duration_as_iso8601 is not None:
            pulumi.set(__self__, "lock_duration_as_iso8601", lock_duration_as_iso8601)
        if max_delivery_count is not None:
            pulumi.set(__self__, "max_delivery_count", max_delivery_count)
        if ttl_as_iso8601 is not None:
            pulumi.set(__self__, "ttl_as_iso8601", ttl_as_iso8601)

    @property
    @pulumi.getter(name="lockDurationAsIso8601")
    def lock_duration_as_iso8601(self) -> Optional[str]:
        """
        The lock duration for the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        """
        return pulumi.get(self, "lock_duration_as_iso8601")

    @property
    @pulumi.getter(name="maxDeliveryCount")
    def max_delivery_count(self) -> Optional[float]:
        """
        The number of times the IoT hub attempts to deliver a message on the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        """
        return pulumi.get(self, "max_delivery_count")

    @property
    @pulumi.getter(name="ttlAsIso8601")
    def ttl_as_iso8601(self) -> Optional[str]:
        """
        The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        """
        return pulumi.get(self, "ttl_as_iso8601")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotHubPropertiesResponse(dict):
    """
    The properties of an IoT hub.
    """
    def __init__(__self__, *,
                 host_name: str,
                 provisioning_state: str,
                 authorization_policies: Optional[List['outputs.SharedAccessSignatureAuthorizationRuleResponse']] = None,
                 cloud_to_device: Optional['outputs.CloudToDevicePropertiesResponse'] = None,
                 comments: Optional[str] = None,
                 enable_file_upload_notifications: Optional[bool] = None,
                 event_hub_endpoints: Optional[Mapping[str, 'outputs.EventHubPropertiesResponse']] = None,
                 features: Optional[str] = None,
                 ip_filter_rules: Optional[List['outputs.IpFilterRuleResponse']] = None,
                 messaging_endpoints: Optional[Mapping[str, 'outputs.MessagingEndpointPropertiesResponse']] = None,
                 operations_monitoring_properties: Optional['outputs.OperationsMonitoringPropertiesResponse'] = None,
                 storage_endpoints: Optional[Mapping[str, 'outputs.StorageEndpointPropertiesResponse']] = None):
        """
        The properties of an IoT hub.
        :param str host_name: The name of the host.
        :param str provisioning_state: The provisioning state.
        :param List['SharedAccessSignatureAuthorizationRuleResponseArgs'] authorization_policies: The shared access policies you can use to secure a connection to the IoT hub.
        :param 'CloudToDevicePropertiesResponseArgs' cloud_to_device: The IoT hub cloud-to-device messaging properties.
        :param str comments: Comments.
        :param bool enable_file_upload_notifications: If True, file upload notifications are enabled.
        :param Mapping[str, 'EventHubPropertiesResponseArgs'] event_hub_endpoints: The Event Hub-compatible endpoint properties. The possible keys to this dictionary are events and operationsMonitoringEvents. Both of these keys have to be present in the dictionary while making create or update calls for the IoT hub.
        :param str features: The capabilities and features enabled for the IoT hub.
        :param List['IpFilterRuleResponseArgs'] ip_filter_rules: The IP filter rules.
        :param Mapping[str, 'MessagingEndpointPropertiesResponseArgs'] messaging_endpoints: The messaging endpoint properties for the file upload notification queue.
        :param 'OperationsMonitoringPropertiesResponseArgs' operations_monitoring_properties: The operations monitoring properties for the IoT hub. The possible keys to the dictionary are Connections, DeviceTelemetry, C2DCommands, DeviceIdentityOperations, FileUploadOperations.
        :param Mapping[str, 'StorageEndpointPropertiesResponseArgs'] storage_endpoints: The list of Azure Storage endpoints where you can upload files. Currently you can configure only one Azure Storage account and that MUST have its key as $default. Specifying more than one storage account causes an error to be thrown. Not specifying a value for this property when the enableFileUploadNotifications property is set to True, causes an error to be thrown.
        """
        pulumi.set(__self__, "host_name", host_name)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if authorization_policies is not None:
            pulumi.set(__self__, "authorization_policies", authorization_policies)
        if cloud_to_device is not None:
            pulumi.set(__self__, "cloud_to_device", cloud_to_device)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if enable_file_upload_notifications is not None:
            pulumi.set(__self__, "enable_file_upload_notifications", enable_file_upload_notifications)
        if event_hub_endpoints is not None:
            pulumi.set(__self__, "event_hub_endpoints", event_hub_endpoints)
        if features is not None:
            pulumi.set(__self__, "features", features)
        if ip_filter_rules is not None:
            pulumi.set(__self__, "ip_filter_rules", ip_filter_rules)
        if messaging_endpoints is not None:
            pulumi.set(__self__, "messaging_endpoints", messaging_endpoints)
        if operations_monitoring_properties is not None:
            pulumi.set(__self__, "operations_monitoring_properties", operations_monitoring_properties)
        if storage_endpoints is not None:
            pulumi.set(__self__, "storage_endpoints", storage_endpoints)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> str:
        """
        The name of the host.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="authorizationPolicies")
    def authorization_policies(self) -> Optional[List['outputs.SharedAccessSignatureAuthorizationRuleResponse']]:
        """
        The shared access policies you can use to secure a connection to the IoT hub.
        """
        return pulumi.get(self, "authorization_policies")

    @property
    @pulumi.getter(name="cloudToDevice")
    def cloud_to_device(self) -> Optional['outputs.CloudToDevicePropertiesResponse']:
        """
        The IoT hub cloud-to-device messaging properties.
        """
        return pulumi.get(self, "cloud_to_device")

    @property
    @pulumi.getter
    def comments(self) -> Optional[str]:
        """
        Comments.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="enableFileUploadNotifications")
    def enable_file_upload_notifications(self) -> Optional[bool]:
        """
        If True, file upload notifications are enabled.
        """
        return pulumi.get(self, "enable_file_upload_notifications")

    @property
    @pulumi.getter(name="eventHubEndpoints")
    def event_hub_endpoints(self) -> Optional[Mapping[str, 'outputs.EventHubPropertiesResponse']]:
        """
        The Event Hub-compatible endpoint properties. The possible keys to this dictionary are events and operationsMonitoringEvents. Both of these keys have to be present in the dictionary while making create or update calls for the IoT hub.
        """
        return pulumi.get(self, "event_hub_endpoints")

    @property
    @pulumi.getter
    def features(self) -> Optional[str]:
        """
        The capabilities and features enabled for the IoT hub.
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter(name="ipFilterRules")
    def ip_filter_rules(self) -> Optional[List['outputs.IpFilterRuleResponse']]:
        """
        The IP filter rules.
        """
        return pulumi.get(self, "ip_filter_rules")

    @property
    @pulumi.getter(name="messagingEndpoints")
    def messaging_endpoints(self) -> Optional[Mapping[str, 'outputs.MessagingEndpointPropertiesResponse']]:
        """
        The messaging endpoint properties for the file upload notification queue.
        """
        return pulumi.get(self, "messaging_endpoints")

    @property
    @pulumi.getter(name="operationsMonitoringProperties")
    def operations_monitoring_properties(self) -> Optional['outputs.OperationsMonitoringPropertiesResponse']:
        """
        The operations monitoring properties for the IoT hub. The possible keys to the dictionary are Connections, DeviceTelemetry, C2DCommands, DeviceIdentityOperations, FileUploadOperations.
        """
        return pulumi.get(self, "operations_monitoring_properties")

    @property
    @pulumi.getter(name="storageEndpoints")
    def storage_endpoints(self) -> Optional[Mapping[str, 'outputs.StorageEndpointPropertiesResponse']]:
        """
        The list of Azure Storage endpoints where you can upload files. Currently you can configure only one Azure Storage account and that MUST have its key as $default. Specifying more than one storage account causes an error to be thrown. Not specifying a value for this property when the enableFileUploadNotifications property is set to True, causes an error to be thrown.
        """
        return pulumi.get(self, "storage_endpoints")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotHubSkuInfoResponse(dict):
    """
    Information about the SKU of the IoT hub.
    """
    def __init__(__self__, *,
                 capacity: float,
                 name: str,
                 tier: str):
        """
        Information about the SKU of the IoT hub.
        :param float capacity: The number of provisioned IoT Hub units. See: https://docs.microsoft.com/azure/azure-subscription-service-limits#iot-hub-limits.
        :param str name: The name of the SKU.
        :param str tier: The billing tier for the IoT hub.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def capacity(self) -> float:
        """
        The number of provisioned IoT Hub units. See: https://docs.microsoft.com/azure/azure-subscription-service-limits#iot-hub-limits.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the SKU.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        The billing tier for the IoT hub.
        """
        return pulumi.get(self, "tier")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IpFilterRuleResponse(dict):
    """
    The IP filter rules for the IoT hub.
    """
    def __init__(__self__, *,
                 action: str,
                 filter_name: str,
                 ip_mask: str):
        """
        The IP filter rules for the IoT hub.
        :param str action: The desired action for requests captured by this rule.
        :param str filter_name: The name of the IP filter rule.
        :param str ip_mask: A string that contains the IP address range in CIDR notation for the rule.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "filter_name", filter_name)
        pulumi.set(__self__, "ip_mask", ip_mask)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        The desired action for requests captured by this rule.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="filterName")
    def filter_name(self) -> str:
        """
        The name of the IP filter rule.
        """
        return pulumi.get(self, "filter_name")

    @property
    @pulumi.getter(name="ipMask")
    def ip_mask(self) -> str:
        """
        A string that contains the IP address range in CIDR notation for the rule.
        """
        return pulumi.get(self, "ip_mask")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MessagingEndpointPropertiesResponse(dict):
    """
    The properties of the messaging endpoints used by this IoT hub.
    """
    def __init__(__self__, *,
                 lock_duration_as_iso8601: Optional[str] = None,
                 max_delivery_count: Optional[float] = None,
                 ttl_as_iso8601: Optional[str] = None):
        """
        The properties of the messaging endpoints used by this IoT hub.
        :param str lock_duration_as_iso8601: The lock duration. See: https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-devguide-file-upload.
        :param float max_delivery_count: The number of times the IoT hub attempts to deliver a message. See: https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-devguide-file-upload.
        :param str ttl_as_iso8601: The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-devguide-file-upload.
        """
        if lock_duration_as_iso8601 is not None:
            pulumi.set(__self__, "lock_duration_as_iso8601", lock_duration_as_iso8601)
        if max_delivery_count is not None:
            pulumi.set(__self__, "max_delivery_count", max_delivery_count)
        if ttl_as_iso8601 is not None:
            pulumi.set(__self__, "ttl_as_iso8601", ttl_as_iso8601)

    @property
    @pulumi.getter(name="lockDurationAsIso8601")
    def lock_duration_as_iso8601(self) -> Optional[str]:
        """
        The lock duration. See: https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-devguide-file-upload.
        """
        return pulumi.get(self, "lock_duration_as_iso8601")

    @property
    @pulumi.getter(name="maxDeliveryCount")
    def max_delivery_count(self) -> Optional[float]:
        """
        The number of times the IoT hub attempts to deliver a message. See: https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-devguide-file-upload.
        """
        return pulumi.get(self, "max_delivery_count")

    @property
    @pulumi.getter(name="ttlAsIso8601")
    def ttl_as_iso8601(self) -> Optional[str]:
        """
        The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-devguide-file-upload.
        """
        return pulumi.get(self, "ttl_as_iso8601")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OperationsMonitoringPropertiesResponse(dict):
    """
    The operations monitoring properties for the IoT hub. The possible keys to the dictionary are Connections, DeviceTelemetry, C2DCommands, DeviceIdentityOperations, FileUploadOperations.
    """
    def __init__(__self__, *,
                 events: Optional[Mapping[str, str]] = None):
        """
        The operations monitoring properties for the IoT hub. The possible keys to the dictionary are Connections, DeviceTelemetry, C2DCommands, DeviceIdentityOperations, FileUploadOperations.
        """
        if events is not None:
            pulumi.set(__self__, "events", events)

    @property
    @pulumi.getter
    def events(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "events")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SharedAccessSignatureAuthorizationRuleResponse(dict):
    """
    The properties of an IoT hub shared access policy.
    """
    def __init__(__self__, *,
                 key_name: str,
                 rights: str,
                 primary_key: Optional[str] = None,
                 secondary_key: Optional[str] = None):
        """
        The properties of an IoT hub shared access policy.
        :param str key_name: The name of the shared access policy.
        :param str rights: The permissions assigned to the shared access policy.
        :param str primary_key: The primary key.
        :param str secondary_key: The secondary key.
        """
        pulumi.set(__self__, "key_name", key_name)
        pulumi.set(__self__, "rights", rights)
        if primary_key is not None:
            pulumi.set(__self__, "primary_key", primary_key)
        if secondary_key is not None:
            pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> str:
        """
        The name of the shared access policy.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter
    def rights(self) -> str:
        """
        The permissions assigned to the shared access policy.
        """
        return pulumi.get(self, "rights")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> Optional[str]:
        """
        The primary key.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> Optional[str]:
        """
        The secondary key.
        """
        return pulumi.get(self, "secondary_key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StorageEndpointPropertiesResponse(dict):
    """
    The properties of the Azure Storage endpoint for file upload.
    """
    def __init__(__self__, *,
                 connection_string: str,
                 container_name: str,
                 sas_ttl_as_iso8601: Optional[str] = None):
        """
        The properties of the Azure Storage endpoint for file upload.
        :param str connection_string: The connection string for the Azure Storage account to which files are uploaded.
        :param str container_name: The name of the root container where you upload files. The container need not exist but should be creatable using the connectionString specified.
        :param str sas_ttl_as_iso8601: The period of time for which the SAS URI generated by IoT Hub for file upload is valid. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-file-upload#file-upload-notification-configuration-options.
        """
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "container_name", container_name)
        if sas_ttl_as_iso8601 is not None:
            pulumi.set(__self__, "sas_ttl_as_iso8601", sas_ttl_as_iso8601)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        The connection string for the Azure Storage account to which files are uploaded.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> str:
        """
        The name of the root container where you upload files. The container need not exist but should be creatable using the connectionString specified.
        """
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter(name="sasTtlAsIso8601")
    def sas_ttl_as_iso8601(self) -> Optional[str]:
        """
        The period of time for which the SAS URI generated by IoT Hub for file upload is valid. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-file-upload#file-upload-notification-configuration-options.
        """
        return pulumi.get(self, "sas_ttl_as_iso8601")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


