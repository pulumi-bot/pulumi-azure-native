# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'CloudTieringCachePerformanceResponse',
    'CloudTieringDatePolicyStatusResponse',
    'CloudTieringFilesNotTieringResponse',
    'CloudTieringSpaceSavingsResponse',
    'CloudTieringVolumeFreeSpacePolicyStatusResponse',
    'FilesNotTieringErrorResponse',
    'ServerEndpointCloudTieringStatusResponse',
    'ServerEndpointFilesNotSyncingErrorResponse',
    'ServerEndpointRecallErrorResponse',
    'ServerEndpointRecallStatusResponse',
    'ServerEndpointSyncStatusResponse',
    'SyncActivityStatusResponse',
    'SyncSessionStatusResponse',
]

@pulumi.output_type
class CloudTieringCachePerformanceResponse(dict):
    """
    Server endpoint cloud tiering status object.
    """
    def __init__(__self__, *,
                 cache_hit_bytes: float,
                 cache_hit_bytes_percent: int,
                 cache_miss_bytes: float,
                 last_updated_timestamp: str):
        """
        Server endpoint cloud tiering status object.
        :param float cache_hit_bytes: Count of bytes that were served from the local server
        :param int cache_hit_bytes_percent: Percentage of total bytes (hit + miss) that were served from the local server
        :param float cache_miss_bytes: Count of bytes that were served from the cloud
        :param str last_updated_timestamp: Last updated timestamp
        """
        pulumi.set(__self__, "cache_hit_bytes", cache_hit_bytes)
        pulumi.set(__self__, "cache_hit_bytes_percent", cache_hit_bytes_percent)
        pulumi.set(__self__, "cache_miss_bytes", cache_miss_bytes)
        pulumi.set(__self__, "last_updated_timestamp", last_updated_timestamp)

    @property
    @pulumi.getter(name="cacheHitBytes")
    def cache_hit_bytes(self) -> float:
        """
        Count of bytes that were served from the local server
        """
        return pulumi.get(self, "cache_hit_bytes")

    @property
    @pulumi.getter(name="cacheHitBytesPercent")
    def cache_hit_bytes_percent(self) -> int:
        """
        Percentage of total bytes (hit + miss) that were served from the local server
        """
        return pulumi.get(self, "cache_hit_bytes_percent")

    @property
    @pulumi.getter(name="cacheMissBytes")
    def cache_miss_bytes(self) -> float:
        """
        Count of bytes that were served from the cloud
        """
        return pulumi.get(self, "cache_miss_bytes")

    @property
    @pulumi.getter(name="lastUpdatedTimestamp")
    def last_updated_timestamp(self) -> str:
        """
        Last updated timestamp
        """
        return pulumi.get(self, "last_updated_timestamp")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CloudTieringDatePolicyStatusResponse(dict):
    """
    Status of the date policy
    """
    def __init__(__self__, *,
                 last_updated_timestamp: str,
                 tiered_files_most_recent_access_timestamp: str):
        """
        Status of the date policy
        :param str last_updated_timestamp: Last updated timestamp
        :param str tiered_files_most_recent_access_timestamp: Most recent access time of tiered files
        """
        pulumi.set(__self__, "last_updated_timestamp", last_updated_timestamp)
        pulumi.set(__self__, "tiered_files_most_recent_access_timestamp", tiered_files_most_recent_access_timestamp)

    @property
    @pulumi.getter(name="lastUpdatedTimestamp")
    def last_updated_timestamp(self) -> str:
        """
        Last updated timestamp
        """
        return pulumi.get(self, "last_updated_timestamp")

    @property
    @pulumi.getter(name="tieredFilesMostRecentAccessTimestamp")
    def tiered_files_most_recent_access_timestamp(self) -> str:
        """
        Most recent access time of tiered files
        """
        return pulumi.get(self, "tiered_files_most_recent_access_timestamp")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CloudTieringFilesNotTieringResponse(dict):
    """
    Server endpoint cloud tiering status object.
    """
    def __init__(__self__, *,
                 errors: Sequence['outputs.FilesNotTieringErrorResponse'],
                 last_updated_timestamp: str,
                 total_file_count: float):
        """
        Server endpoint cloud tiering status object.
        :param Sequence['FilesNotTieringErrorResponseArgs'] errors: Array of tiering errors
        :param str last_updated_timestamp: Last updated timestamp
        :param float total_file_count: Last cloud tiering result (HResult)
        """
        pulumi.set(__self__, "errors", errors)
        pulumi.set(__self__, "last_updated_timestamp", last_updated_timestamp)
        pulumi.set(__self__, "total_file_count", total_file_count)

    @property
    @pulumi.getter
    def errors(self) -> Sequence['outputs.FilesNotTieringErrorResponse']:
        """
        Array of tiering errors
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="lastUpdatedTimestamp")
    def last_updated_timestamp(self) -> str:
        """
        Last updated timestamp
        """
        return pulumi.get(self, "last_updated_timestamp")

    @property
    @pulumi.getter(name="totalFileCount")
    def total_file_count(self) -> float:
        """
        Last cloud tiering result (HResult)
        """
        return pulumi.get(self, "total_file_count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CloudTieringSpaceSavingsResponse(dict):
    """
    Server endpoint cloud tiering status object.
    """
    def __init__(__self__, *,
                 cached_size_bytes: float,
                 last_updated_timestamp: str,
                 space_savings_bytes: float,
                 space_savings_percent: int,
                 total_size_cloud_bytes: float,
                 volume_size_bytes: float):
        """
        Server endpoint cloud tiering status object.
        :param float cached_size_bytes: Cached content size on the server
        :param str last_updated_timestamp: Last updated timestamp
        :param float space_savings_bytes: Count of bytes saved on the server
        :param int space_savings_percent: Percentage of cached size over total size
        :param float total_size_cloud_bytes: Total size of content in the azure file share
        :param float volume_size_bytes: Volume size
        """
        pulumi.set(__self__, "cached_size_bytes", cached_size_bytes)
        pulumi.set(__self__, "last_updated_timestamp", last_updated_timestamp)
        pulumi.set(__self__, "space_savings_bytes", space_savings_bytes)
        pulumi.set(__self__, "space_savings_percent", space_savings_percent)
        pulumi.set(__self__, "total_size_cloud_bytes", total_size_cloud_bytes)
        pulumi.set(__self__, "volume_size_bytes", volume_size_bytes)

    @property
    @pulumi.getter(name="cachedSizeBytes")
    def cached_size_bytes(self) -> float:
        """
        Cached content size on the server
        """
        return pulumi.get(self, "cached_size_bytes")

    @property
    @pulumi.getter(name="lastUpdatedTimestamp")
    def last_updated_timestamp(self) -> str:
        """
        Last updated timestamp
        """
        return pulumi.get(self, "last_updated_timestamp")

    @property
    @pulumi.getter(name="spaceSavingsBytes")
    def space_savings_bytes(self) -> float:
        """
        Count of bytes saved on the server
        """
        return pulumi.get(self, "space_savings_bytes")

    @property
    @pulumi.getter(name="spaceSavingsPercent")
    def space_savings_percent(self) -> int:
        """
        Percentage of cached size over total size
        """
        return pulumi.get(self, "space_savings_percent")

    @property
    @pulumi.getter(name="totalSizeCloudBytes")
    def total_size_cloud_bytes(self) -> float:
        """
        Total size of content in the azure file share
        """
        return pulumi.get(self, "total_size_cloud_bytes")

    @property
    @pulumi.getter(name="volumeSizeBytes")
    def volume_size_bytes(self) -> float:
        """
        Volume size
        """
        return pulumi.get(self, "volume_size_bytes")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CloudTieringVolumeFreeSpacePolicyStatusResponse(dict):
    """
    Status of the volume free space policy
    """
    def __init__(__self__, *,
                 current_volume_free_space_percent: int,
                 effective_volume_free_space_policy: int,
                 last_updated_timestamp: str):
        """
        Status of the volume free space policy
        :param int current_volume_free_space_percent: Current volume free space percentage.
        :param int effective_volume_free_space_policy: In the case where multiple server endpoints are present in a volume, an effective free space policy is applied.
        :param str last_updated_timestamp: Last updated timestamp
        """
        pulumi.set(__self__, "current_volume_free_space_percent", current_volume_free_space_percent)
        pulumi.set(__self__, "effective_volume_free_space_policy", effective_volume_free_space_policy)
        pulumi.set(__self__, "last_updated_timestamp", last_updated_timestamp)

    @property
    @pulumi.getter(name="currentVolumeFreeSpacePercent")
    def current_volume_free_space_percent(self) -> int:
        """
        Current volume free space percentage.
        """
        return pulumi.get(self, "current_volume_free_space_percent")

    @property
    @pulumi.getter(name="effectiveVolumeFreeSpacePolicy")
    def effective_volume_free_space_policy(self) -> int:
        """
        In the case where multiple server endpoints are present in a volume, an effective free space policy is applied.
        """
        return pulumi.get(self, "effective_volume_free_space_policy")

    @property
    @pulumi.getter(name="lastUpdatedTimestamp")
    def last_updated_timestamp(self) -> str:
        """
        Last updated timestamp
        """
        return pulumi.get(self, "last_updated_timestamp")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FilesNotTieringErrorResponse(dict):
    """
    Files not tiering error object
    """
    def __init__(__self__, *,
                 error_code: int,
                 file_count: float):
        """
        Files not tiering error object
        :param int error_code: Error code (HResult)
        :param float file_count: Count of files with this error
        """
        pulumi.set(__self__, "error_code", error_code)
        pulumi.set(__self__, "file_count", file_count)

    @property
    @pulumi.getter(name="errorCode")
    def error_code(self) -> int:
        """
        Error code (HResult)
        """
        return pulumi.get(self, "error_code")

    @property
    @pulumi.getter(name="fileCount")
    def file_count(self) -> float:
        """
        Count of files with this error
        """
        return pulumi.get(self, "file_count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerEndpointCloudTieringStatusResponse(dict):
    """
    Server endpoint cloud tiering status object.
    """
    def __init__(__self__, *,
                 cache_performance: 'outputs.CloudTieringCachePerformanceResponse',
                 date_policy_status: 'outputs.CloudTieringDatePolicyStatusResponse',
                 files_not_tiering: 'outputs.CloudTieringFilesNotTieringResponse',
                 health: str,
                 health_last_updated_timestamp: str,
                 last_cloud_tiering_result: int,
                 last_success_timestamp: str,
                 last_updated_timestamp: str,
                 space_savings: 'outputs.CloudTieringSpaceSavingsResponse',
                 volume_free_space_policy_status: 'outputs.CloudTieringVolumeFreeSpacePolicyStatusResponse'):
        """
        Server endpoint cloud tiering status object.
        :param 'CloudTieringCachePerformanceResponseArgs' cache_performance: Information regarding how well the local cache on the server is performing.
        :param 'CloudTieringDatePolicyStatusResponseArgs' date_policy_status: Status of the date policy
        :param 'CloudTieringFilesNotTieringResponseArgs' files_not_tiering: Information regarding files that failed to be tiered
        :param str health: Cloud tiering health state.
        :param str health_last_updated_timestamp: The last updated timestamp of health state
        :param int last_cloud_tiering_result: Last cloud tiering result (HResult)
        :param str last_success_timestamp: Last cloud tiering success timestamp
        :param str last_updated_timestamp: Last updated timestamp
        :param 'CloudTieringSpaceSavingsResponseArgs' space_savings: Information regarding how much local space cloud tiering is saving.
        :param 'CloudTieringVolumeFreeSpacePolicyStatusResponseArgs' volume_free_space_policy_status: Status of the volume free space policy
        """
        pulumi.set(__self__, "cache_performance", cache_performance)
        pulumi.set(__self__, "date_policy_status", date_policy_status)
        pulumi.set(__self__, "files_not_tiering", files_not_tiering)
        pulumi.set(__self__, "health", health)
        pulumi.set(__self__, "health_last_updated_timestamp", health_last_updated_timestamp)
        pulumi.set(__self__, "last_cloud_tiering_result", last_cloud_tiering_result)
        pulumi.set(__self__, "last_success_timestamp", last_success_timestamp)
        pulumi.set(__self__, "last_updated_timestamp", last_updated_timestamp)
        pulumi.set(__self__, "space_savings", space_savings)
        pulumi.set(__self__, "volume_free_space_policy_status", volume_free_space_policy_status)

    @property
    @pulumi.getter(name="cachePerformance")
    def cache_performance(self) -> 'outputs.CloudTieringCachePerformanceResponse':
        """
        Information regarding how well the local cache on the server is performing.
        """
        return pulumi.get(self, "cache_performance")

    @property
    @pulumi.getter(name="datePolicyStatus")
    def date_policy_status(self) -> 'outputs.CloudTieringDatePolicyStatusResponse':
        """
        Status of the date policy
        """
        return pulumi.get(self, "date_policy_status")

    @property
    @pulumi.getter(name="filesNotTiering")
    def files_not_tiering(self) -> 'outputs.CloudTieringFilesNotTieringResponse':
        """
        Information regarding files that failed to be tiered
        """
        return pulumi.get(self, "files_not_tiering")

    @property
    @pulumi.getter
    def health(self) -> str:
        """
        Cloud tiering health state.
        """
        return pulumi.get(self, "health")

    @property
    @pulumi.getter(name="healthLastUpdatedTimestamp")
    def health_last_updated_timestamp(self) -> str:
        """
        The last updated timestamp of health state
        """
        return pulumi.get(self, "health_last_updated_timestamp")

    @property
    @pulumi.getter(name="lastCloudTieringResult")
    def last_cloud_tiering_result(self) -> int:
        """
        Last cloud tiering result (HResult)
        """
        return pulumi.get(self, "last_cloud_tiering_result")

    @property
    @pulumi.getter(name="lastSuccessTimestamp")
    def last_success_timestamp(self) -> str:
        """
        Last cloud tiering success timestamp
        """
        return pulumi.get(self, "last_success_timestamp")

    @property
    @pulumi.getter(name="lastUpdatedTimestamp")
    def last_updated_timestamp(self) -> str:
        """
        Last updated timestamp
        """
        return pulumi.get(self, "last_updated_timestamp")

    @property
    @pulumi.getter(name="spaceSavings")
    def space_savings(self) -> 'outputs.CloudTieringSpaceSavingsResponse':
        """
        Information regarding how much local space cloud tiering is saving.
        """
        return pulumi.get(self, "space_savings")

    @property
    @pulumi.getter(name="volumeFreeSpacePolicyStatus")
    def volume_free_space_policy_status(self) -> 'outputs.CloudTieringVolumeFreeSpacePolicyStatusResponse':
        """
        Status of the volume free space policy
        """
        return pulumi.get(self, "volume_free_space_policy_status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerEndpointFilesNotSyncingErrorResponse(dict):
    """
    Files not syncing error object
    """
    def __init__(__self__, *,
                 error_code: int,
                 persistent_count: float,
                 transient_count: float):
        """
        Files not syncing error object
        :param int error_code: Error code (HResult)
        :param float persistent_count: Count of persistent files not syncing with the specified error code
        :param float transient_count: Count of transient files not syncing with the specified error code
        """
        pulumi.set(__self__, "error_code", error_code)
        pulumi.set(__self__, "persistent_count", persistent_count)
        pulumi.set(__self__, "transient_count", transient_count)

    @property
    @pulumi.getter(name="errorCode")
    def error_code(self) -> int:
        """
        Error code (HResult)
        """
        return pulumi.get(self, "error_code")

    @property
    @pulumi.getter(name="persistentCount")
    def persistent_count(self) -> float:
        """
        Count of persistent files not syncing with the specified error code
        """
        return pulumi.get(self, "persistent_count")

    @property
    @pulumi.getter(name="transientCount")
    def transient_count(self) -> float:
        """
        Count of transient files not syncing with the specified error code
        """
        return pulumi.get(self, "transient_count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerEndpointRecallErrorResponse(dict):
    """
    Server endpoint recall error object
    """
    def __init__(__self__, *,
                 count: float,
                 error_code: int):
        """
        Server endpoint recall error object
        :param float count: Count of occurences of the error
        :param int error_code: Error code (HResult)
        """
        pulumi.set(__self__, "count", count)
        pulumi.set(__self__, "error_code", error_code)

    @property
    @pulumi.getter
    def count(self) -> float:
        """
        Count of occurences of the error
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="errorCode")
    def error_code(self) -> int:
        """
        Error code (HResult)
        """
        return pulumi.get(self, "error_code")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerEndpointRecallStatusResponse(dict):
    """
    Server endpoint recall status object.
    """
    def __init__(__self__, *,
                 last_updated_timestamp: str,
                 recall_errors: Sequence['outputs.ServerEndpointRecallErrorResponse'],
                 total_recall_errors_count: float):
        """
        Server endpoint recall status object.
        :param str last_updated_timestamp: Last updated timestamp
        :param Sequence['ServerEndpointRecallErrorResponseArgs'] recall_errors: Array of recall errors
        :param float total_recall_errors_count: Total count of recall errors.
        """
        pulumi.set(__self__, "last_updated_timestamp", last_updated_timestamp)
        pulumi.set(__self__, "recall_errors", recall_errors)
        pulumi.set(__self__, "total_recall_errors_count", total_recall_errors_count)

    @property
    @pulumi.getter(name="lastUpdatedTimestamp")
    def last_updated_timestamp(self) -> str:
        """
        Last updated timestamp
        """
        return pulumi.get(self, "last_updated_timestamp")

    @property
    @pulumi.getter(name="recallErrors")
    def recall_errors(self) -> Sequence['outputs.ServerEndpointRecallErrorResponse']:
        """
        Array of recall errors
        """
        return pulumi.get(self, "recall_errors")

    @property
    @pulumi.getter(name="totalRecallErrorsCount")
    def total_recall_errors_count(self) -> float:
        """
        Total count of recall errors.
        """
        return pulumi.get(self, "total_recall_errors_count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerEndpointSyncStatusResponse(dict):
    """
    Server Endpoint sync status
    """
    def __init__(__self__, *,
                 combined_health: str,
                 download_activity: 'outputs.SyncActivityStatusResponse',
                 download_health: str,
                 download_status: 'outputs.SyncSessionStatusResponse',
                 last_updated_timestamp: str,
                 offline_data_transfer_status: str,
                 sync_activity: str,
                 total_persistent_files_not_syncing_count: float,
                 upload_activity: 'outputs.SyncActivityStatusResponse',
                 upload_health: str,
                 upload_status: 'outputs.SyncSessionStatusResponse'):
        """
        Server Endpoint sync status
        :param str combined_health: Combined Health Status.
        :param 'SyncActivityStatusResponseArgs' download_activity: Download sync activity
        :param str download_health: Download Health Status.
        :param 'SyncSessionStatusResponseArgs' download_status: Download Status
        :param str last_updated_timestamp: Last Updated Timestamp
        :param str offline_data_transfer_status: Offline Data Transfer State
        :param str sync_activity: Sync activity
        :param float total_persistent_files_not_syncing_count: Total count of persistent files not syncing (combined upload + download).
        :param 'SyncActivityStatusResponseArgs' upload_activity: Upload sync activity
        :param str upload_health: Upload Health Status.
        :param 'SyncSessionStatusResponseArgs' upload_status: Upload Status
        """
        pulumi.set(__self__, "combined_health", combined_health)
        pulumi.set(__self__, "download_activity", download_activity)
        pulumi.set(__self__, "download_health", download_health)
        pulumi.set(__self__, "download_status", download_status)
        pulumi.set(__self__, "last_updated_timestamp", last_updated_timestamp)
        pulumi.set(__self__, "offline_data_transfer_status", offline_data_transfer_status)
        pulumi.set(__self__, "sync_activity", sync_activity)
        pulumi.set(__self__, "total_persistent_files_not_syncing_count", total_persistent_files_not_syncing_count)
        pulumi.set(__self__, "upload_activity", upload_activity)
        pulumi.set(__self__, "upload_health", upload_health)
        pulumi.set(__self__, "upload_status", upload_status)

    @property
    @pulumi.getter(name="combinedHealth")
    def combined_health(self) -> str:
        """
        Combined Health Status.
        """
        return pulumi.get(self, "combined_health")

    @property
    @pulumi.getter(name="downloadActivity")
    def download_activity(self) -> 'outputs.SyncActivityStatusResponse':
        """
        Download sync activity
        """
        return pulumi.get(self, "download_activity")

    @property
    @pulumi.getter(name="downloadHealth")
    def download_health(self) -> str:
        """
        Download Health Status.
        """
        return pulumi.get(self, "download_health")

    @property
    @pulumi.getter(name="downloadStatus")
    def download_status(self) -> 'outputs.SyncSessionStatusResponse':
        """
        Download Status
        """
        return pulumi.get(self, "download_status")

    @property
    @pulumi.getter(name="lastUpdatedTimestamp")
    def last_updated_timestamp(self) -> str:
        """
        Last Updated Timestamp
        """
        return pulumi.get(self, "last_updated_timestamp")

    @property
    @pulumi.getter(name="offlineDataTransferStatus")
    def offline_data_transfer_status(self) -> str:
        """
        Offline Data Transfer State
        """
        return pulumi.get(self, "offline_data_transfer_status")

    @property
    @pulumi.getter(name="syncActivity")
    def sync_activity(self) -> str:
        """
        Sync activity
        """
        return pulumi.get(self, "sync_activity")

    @property
    @pulumi.getter(name="totalPersistentFilesNotSyncingCount")
    def total_persistent_files_not_syncing_count(self) -> float:
        """
        Total count of persistent files not syncing (combined upload + download).
        """
        return pulumi.get(self, "total_persistent_files_not_syncing_count")

    @property
    @pulumi.getter(name="uploadActivity")
    def upload_activity(self) -> 'outputs.SyncActivityStatusResponse':
        """
        Upload sync activity
        """
        return pulumi.get(self, "upload_activity")

    @property
    @pulumi.getter(name="uploadHealth")
    def upload_health(self) -> str:
        """
        Upload Health Status.
        """
        return pulumi.get(self, "upload_health")

    @property
    @pulumi.getter(name="uploadStatus")
    def upload_status(self) -> 'outputs.SyncSessionStatusResponse':
        """
        Upload Status
        """
        return pulumi.get(self, "upload_status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SyncActivityStatusResponse(dict):
    """
    Sync Session status object.
    """
    def __init__(__self__, *,
                 applied_bytes: float,
                 applied_item_count: float,
                 per_item_error_count: float,
                 timestamp: str,
                 total_bytes: float,
                 total_item_count: float):
        """
        Sync Session status object.
        :param float applied_bytes: Applied bytes
        :param float applied_item_count: Applied item count.
        :param float per_item_error_count: Per item error count
        :param str timestamp: Timestamp when properties were updated
        :param float total_bytes: Total bytes (if available)
        :param float total_item_count: Total item count (if available)
        """
        pulumi.set(__self__, "applied_bytes", applied_bytes)
        pulumi.set(__self__, "applied_item_count", applied_item_count)
        pulumi.set(__self__, "per_item_error_count", per_item_error_count)
        pulumi.set(__self__, "timestamp", timestamp)
        pulumi.set(__self__, "total_bytes", total_bytes)
        pulumi.set(__self__, "total_item_count", total_item_count)

    @property
    @pulumi.getter(name="appliedBytes")
    def applied_bytes(self) -> float:
        """
        Applied bytes
        """
        return pulumi.get(self, "applied_bytes")

    @property
    @pulumi.getter(name="appliedItemCount")
    def applied_item_count(self) -> float:
        """
        Applied item count.
        """
        return pulumi.get(self, "applied_item_count")

    @property
    @pulumi.getter(name="perItemErrorCount")
    def per_item_error_count(self) -> float:
        """
        Per item error count
        """
        return pulumi.get(self, "per_item_error_count")

    @property
    @pulumi.getter
    def timestamp(self) -> str:
        """
        Timestamp when properties were updated
        """
        return pulumi.get(self, "timestamp")

    @property
    @pulumi.getter(name="totalBytes")
    def total_bytes(self) -> float:
        """
        Total bytes (if available)
        """
        return pulumi.get(self, "total_bytes")

    @property
    @pulumi.getter(name="totalItemCount")
    def total_item_count(self) -> float:
        """
        Total item count (if available)
        """
        return pulumi.get(self, "total_item_count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SyncSessionStatusResponse(dict):
    """
    Sync Session status object.
    """
    def __init__(__self__, *,
                 files_not_syncing_errors: Sequence['outputs.ServerEndpointFilesNotSyncingErrorResponse'],
                 last_sync_per_item_error_count: float,
                 last_sync_result: int,
                 last_sync_success_timestamp: str,
                 last_sync_timestamp: str,
                 persistent_files_not_syncing_count: float,
                 transient_files_not_syncing_count: float):
        """
        Sync Session status object.
        :param Sequence['ServerEndpointFilesNotSyncingErrorResponseArgs'] files_not_syncing_errors: Array of per-item errors coming from the last sync session.
        :param float last_sync_per_item_error_count: Last sync per item error count.
        :param int last_sync_result: Last sync result (HResult)
        :param str last_sync_success_timestamp: Last sync success timestamp
        :param str last_sync_timestamp: Last sync timestamp
        :param float persistent_files_not_syncing_count: Count of persistent files not syncing.
        :param float transient_files_not_syncing_count: Count of transient files not syncing.
        """
        pulumi.set(__self__, "files_not_syncing_errors", files_not_syncing_errors)
        pulumi.set(__self__, "last_sync_per_item_error_count", last_sync_per_item_error_count)
        pulumi.set(__self__, "last_sync_result", last_sync_result)
        pulumi.set(__self__, "last_sync_success_timestamp", last_sync_success_timestamp)
        pulumi.set(__self__, "last_sync_timestamp", last_sync_timestamp)
        pulumi.set(__self__, "persistent_files_not_syncing_count", persistent_files_not_syncing_count)
        pulumi.set(__self__, "transient_files_not_syncing_count", transient_files_not_syncing_count)

    @property
    @pulumi.getter(name="filesNotSyncingErrors")
    def files_not_syncing_errors(self) -> Sequence['outputs.ServerEndpointFilesNotSyncingErrorResponse']:
        """
        Array of per-item errors coming from the last sync session.
        """
        return pulumi.get(self, "files_not_syncing_errors")

    @property
    @pulumi.getter(name="lastSyncPerItemErrorCount")
    def last_sync_per_item_error_count(self) -> float:
        """
        Last sync per item error count.
        """
        return pulumi.get(self, "last_sync_per_item_error_count")

    @property
    @pulumi.getter(name="lastSyncResult")
    def last_sync_result(self) -> int:
        """
        Last sync result (HResult)
        """
        return pulumi.get(self, "last_sync_result")

    @property
    @pulumi.getter(name="lastSyncSuccessTimestamp")
    def last_sync_success_timestamp(self) -> str:
        """
        Last sync success timestamp
        """
        return pulumi.get(self, "last_sync_success_timestamp")

    @property
    @pulumi.getter(name="lastSyncTimestamp")
    def last_sync_timestamp(self) -> str:
        """
        Last sync timestamp
        """
        return pulumi.get(self, "last_sync_timestamp")

    @property
    @pulumi.getter(name="persistentFilesNotSyncingCount")
    def persistent_files_not_syncing_count(self) -> float:
        """
        Count of persistent files not syncing.
        """
        return pulumi.get(self, "persistent_files_not_syncing_count")

    @property
    @pulumi.getter(name="transientFilesNotSyncingCount")
    def transient_files_not_syncing_count(self) -> float:
        """
        Count of transient files not syncing.
        """
        return pulumi.get(self, "transient_files_not_syncing_count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


