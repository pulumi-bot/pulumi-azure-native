# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = [
    'DriveStatusArgs',
    'ExportArgs',
    'JobDetailsArgs',
    'PackageInfomationArgs',
    'ReturnAddressArgs',
    'ReturnShippingArgs',
    'ShippingInformationArgs',
]

@pulumi.input_type
class DriveStatusArgs:
    def __init__(__self__, *,
                 bit_locker_key: Optional[pulumi.Input[str]] = None,
                 bytes_succeeded: Optional[pulumi.Input[float]] = None,
                 copy_status: Optional[pulumi.Input[str]] = None,
                 drive_header_hash: Optional[pulumi.Input[str]] = None,
                 drive_id: Optional[pulumi.Input[str]] = None,
                 error_log_uri: Optional[pulumi.Input[str]] = None,
                 manifest_file: Optional[pulumi.Input[str]] = None,
                 manifest_hash: Optional[pulumi.Input[str]] = None,
                 manifest_uri: Optional[pulumi.Input[str]] = None,
                 percent_complete: Optional[pulumi.Input[int]] = None,
                 state: Optional[pulumi.Input[Union[str, 'DriveState']]] = None,
                 verbose_log_uri: Optional[pulumi.Input[str]] = None):
        """
        Provides information about the drive's status
        :param pulumi.Input[str] bit_locker_key: The BitLocker key used to encrypt the drive.
        :param pulumi.Input[float] bytes_succeeded: Bytes successfully transferred for the drive.
        :param pulumi.Input[str] copy_status: Detailed status about the data transfer process. This field is not returned in the response until the drive is in the Transferring state.
        :param pulumi.Input[str] drive_header_hash: The drive header hash value.
        :param pulumi.Input[str] drive_id: The drive's hardware serial number, without spaces.
        :param pulumi.Input[str] error_log_uri: A URI that points to the blob containing the error log for the data transfer operation.
        :param pulumi.Input[str] manifest_file: The relative path of the manifest file on the drive. 
        :param pulumi.Input[str] manifest_hash: The Base16-encoded MD5 hash of the manifest file on the drive.
        :param pulumi.Input[str] manifest_uri: A URI that points to the blob containing the drive manifest file. 
        :param pulumi.Input[int] percent_complete: Percentage completed for the drive. 
        :param pulumi.Input[Union[str, 'DriveState']] state: The drive's current state. 
        :param pulumi.Input[str] verbose_log_uri: A URI that points to the blob containing the verbose log for the data transfer operation. 
        """
        if bit_locker_key is not None:
            pulumi.set(__self__, "bit_locker_key", bit_locker_key)
        if bytes_succeeded is not None:
            pulumi.set(__self__, "bytes_succeeded", bytes_succeeded)
        if copy_status is not None:
            pulumi.set(__self__, "copy_status", copy_status)
        if drive_header_hash is not None:
            pulumi.set(__self__, "drive_header_hash", drive_header_hash)
        if drive_id is not None:
            pulumi.set(__self__, "drive_id", drive_id)
        if error_log_uri is not None:
            pulumi.set(__self__, "error_log_uri", error_log_uri)
        if manifest_file is not None:
            pulumi.set(__self__, "manifest_file", manifest_file)
        if manifest_hash is not None:
            pulumi.set(__self__, "manifest_hash", manifest_hash)
        if manifest_uri is not None:
            pulumi.set(__self__, "manifest_uri", manifest_uri)
        if percent_complete is not None:
            pulumi.set(__self__, "percent_complete", percent_complete)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if verbose_log_uri is not None:
            pulumi.set(__self__, "verbose_log_uri", verbose_log_uri)

    @property
    @pulumi.getter(name="bitLockerKey")
    def bit_locker_key(self) -> Optional[pulumi.Input[str]]:
        """
        The BitLocker key used to encrypt the drive.
        """
        return pulumi.get(self, "bit_locker_key")

    @bit_locker_key.setter
    def bit_locker_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bit_locker_key", value)

    @property
    @pulumi.getter(name="bytesSucceeded")
    def bytes_succeeded(self) -> Optional[pulumi.Input[float]]:
        """
        Bytes successfully transferred for the drive.
        """
        return pulumi.get(self, "bytes_succeeded")

    @bytes_succeeded.setter
    def bytes_succeeded(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "bytes_succeeded", value)

    @property
    @pulumi.getter(name="copyStatus")
    def copy_status(self) -> Optional[pulumi.Input[str]]:
        """
        Detailed status about the data transfer process. This field is not returned in the response until the drive is in the Transferring state.
        """
        return pulumi.get(self, "copy_status")

    @copy_status.setter
    def copy_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "copy_status", value)

    @property
    @pulumi.getter(name="driveHeaderHash")
    def drive_header_hash(self) -> Optional[pulumi.Input[str]]:
        """
        The drive header hash value.
        """
        return pulumi.get(self, "drive_header_hash")

    @drive_header_hash.setter
    def drive_header_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "drive_header_hash", value)

    @property
    @pulumi.getter(name="driveId")
    def drive_id(self) -> Optional[pulumi.Input[str]]:
        """
        The drive's hardware serial number, without spaces.
        """
        return pulumi.get(self, "drive_id")

    @drive_id.setter
    def drive_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "drive_id", value)

    @property
    @pulumi.getter(name="errorLogUri")
    def error_log_uri(self) -> Optional[pulumi.Input[str]]:
        """
        A URI that points to the blob containing the error log for the data transfer operation.
        """
        return pulumi.get(self, "error_log_uri")

    @error_log_uri.setter
    def error_log_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "error_log_uri", value)

    @property
    @pulumi.getter(name="manifestFile")
    def manifest_file(self) -> Optional[pulumi.Input[str]]:
        """
        The relative path of the manifest file on the drive. 
        """
        return pulumi.get(self, "manifest_file")

    @manifest_file.setter
    def manifest_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "manifest_file", value)

    @property
    @pulumi.getter(name="manifestHash")
    def manifest_hash(self) -> Optional[pulumi.Input[str]]:
        """
        The Base16-encoded MD5 hash of the manifest file on the drive.
        """
        return pulumi.get(self, "manifest_hash")

    @manifest_hash.setter
    def manifest_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "manifest_hash", value)

    @property
    @pulumi.getter(name="manifestUri")
    def manifest_uri(self) -> Optional[pulumi.Input[str]]:
        """
        A URI that points to the blob containing the drive manifest file. 
        """
        return pulumi.get(self, "manifest_uri")

    @manifest_uri.setter
    def manifest_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "manifest_uri", value)

    @property
    @pulumi.getter(name="percentComplete")
    def percent_complete(self) -> Optional[pulumi.Input[int]]:
        """
        Percentage completed for the drive. 
        """
        return pulumi.get(self, "percent_complete")

    @percent_complete.setter
    def percent_complete(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "percent_complete", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[Union[str, 'DriveState']]]:
        """
        The drive's current state. 
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[Union[str, 'DriveState']]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="verboseLogUri")
    def verbose_log_uri(self) -> Optional[pulumi.Input[str]]:
        """
        A URI that points to the blob containing the verbose log for the data transfer operation. 
        """
        return pulumi.get(self, "verbose_log_uri")

    @verbose_log_uri.setter
    def verbose_log_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verbose_log_uri", value)


@pulumi.input_type
class ExportArgs:
    def __init__(__self__, *,
                 blob_listblob_path: Optional[pulumi.Input[str]] = None,
                 blob_path: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 blob_path_prefix: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        A property containing information about the blobs to be exported for an export job. This property is required for export jobs, but must not be specified for import jobs.
        :param pulumi.Input[str] blob_listblob_path: The relative URI to the block blob that contains the list of blob paths or blob path prefixes as defined above, beginning with the container name. If the blob is in root container, the URI must begin with $root. 
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blob_path: A collection of blob-path strings.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blob_path_prefix: A collection of blob-prefix strings.
        """
        if blob_listblob_path is not None:
            pulumi.set(__self__, "blob_listblob_path", blob_listblob_path)
        if blob_path is not None:
            pulumi.set(__self__, "blob_path", blob_path)
        if blob_path_prefix is not None:
            pulumi.set(__self__, "blob_path_prefix", blob_path_prefix)

    @property
    @pulumi.getter(name="blobListblobPath")
    def blob_listblob_path(self) -> Optional[pulumi.Input[str]]:
        """
        The relative URI to the block blob that contains the list of blob paths or blob path prefixes as defined above, beginning with the container name. If the blob is in root container, the URI must begin with $root. 
        """
        return pulumi.get(self, "blob_listblob_path")

    @blob_listblob_path.setter
    def blob_listblob_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blob_listblob_path", value)

    @property
    @pulumi.getter(name="blobPath")
    def blob_path(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A collection of blob-path strings.
        """
        return pulumi.get(self, "blob_path")

    @blob_path.setter
    def blob_path(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blob_path", value)

    @property
    @pulumi.getter(name="blobPathPrefix")
    def blob_path_prefix(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A collection of blob-prefix strings.
        """
        return pulumi.get(self, "blob_path_prefix")

    @blob_path_prefix.setter
    def blob_path_prefix(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blob_path_prefix", value)


@pulumi.input_type
class JobDetailsArgs:
    def __init__(__self__, *,
                 backup_drive_manifest: Optional[pulumi.Input[bool]] = None,
                 cancel_requested: Optional[pulumi.Input[bool]] = None,
                 delivery_package: Optional[pulumi.Input['PackageInfomationArgs']] = None,
                 diagnostics_path: Optional[pulumi.Input[str]] = None,
                 drive_list: Optional[pulumi.Input[Sequence[pulumi.Input['DriveStatusArgs']]]] = None,
                 export: Optional[pulumi.Input['ExportArgs']] = None,
                 incomplete_blob_list_uri: Optional[pulumi.Input[str]] = None,
                 job_type: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 percent_complete: Optional[pulumi.Input[int]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 return_address: Optional[pulumi.Input['ReturnAddressArgs']] = None,
                 return_package: Optional[pulumi.Input['PackageInfomationArgs']] = None,
                 return_shipping: Optional[pulumi.Input['ReturnShippingArgs']] = None,
                 shipping_information: Optional[pulumi.Input['ShippingInformationArgs']] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None):
        """
        Specifies the job properties
        :param pulumi.Input[bool] backup_drive_manifest: Default value is false. Indicates whether the manifest files on the drives should be copied to block blobs.
        :param pulumi.Input[bool] cancel_requested: Indicates whether a request has been submitted to cancel the job.
        :param pulumi.Input['PackageInfomationArgs'] delivery_package: Contains information about the package being shipped by the customer to the Microsoft data center. 
        :param pulumi.Input[str] diagnostics_path: The virtual blob directory to which the copy logs and backups of drive manifest files (if enabled) will be stored.
        :param pulumi.Input[Sequence[pulumi.Input['DriveStatusArgs']]] drive_list: List of up to ten drives that comprise the job. The drive list is a required element for an import job; it is not specified for export jobs.
        :param pulumi.Input['ExportArgs'] export: A property containing information about the blobs to be exported for an export job. This property is included for export jobs only.
        :param pulumi.Input[str] incomplete_blob_list_uri: A blob path that points to a block blob containing a list of blob names that were not exported due to insufficient drive space. If all blobs were exported successfully, then this element is not included in the response.
        :param pulumi.Input[str] job_type: The type of job
        :param pulumi.Input[str] log_level: Default value is Error. Indicates whether error logging or verbose logging will be enabled.
        :param pulumi.Input[int] percent_complete: Overall percentage completed for the job.
        :param pulumi.Input[str] provisioning_state: Specifies the provisioning state of the job.
        :param pulumi.Input['ReturnAddressArgs'] return_address: Specifies the return address information for the job. 
        :param pulumi.Input['PackageInfomationArgs'] return_package: Contains information about the package being shipped from the Microsoft data center to the customer to return the drives. The format is the same as the deliveryPackage property above. This property is not included if the drives have not yet been returned. 
        :param pulumi.Input['ReturnShippingArgs'] return_shipping: Specifies the return carrier and customer's account with the carrier. 
        :param pulumi.Input['ShippingInformationArgs'] shipping_information: Contains information about the Microsoft datacenter to which the drives should be shipped. 
        :param pulumi.Input[str] state: Current state of the job.
        :param pulumi.Input[str] storage_account_id: The resource identifier of the storage account where data will be imported to or exported from.
        """
        if backup_drive_manifest is not None:
            pulumi.set(__self__, "backup_drive_manifest", backup_drive_manifest)
        if cancel_requested is not None:
            pulumi.set(__self__, "cancel_requested", cancel_requested)
        if delivery_package is not None:
            pulumi.set(__self__, "delivery_package", delivery_package)
        if diagnostics_path is not None:
            pulumi.set(__self__, "diagnostics_path", diagnostics_path)
        if drive_list is not None:
            pulumi.set(__self__, "drive_list", drive_list)
        if export is not None:
            pulumi.set(__self__, "export", export)
        if incomplete_blob_list_uri is not None:
            pulumi.set(__self__, "incomplete_blob_list_uri", incomplete_blob_list_uri)
        if job_type is not None:
            pulumi.set(__self__, "job_type", job_type)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)
        if percent_complete is not None:
            pulumi.set(__self__, "percent_complete", percent_complete)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)
        if return_address is not None:
            pulumi.set(__self__, "return_address", return_address)
        if return_package is not None:
            pulumi.set(__self__, "return_package", return_package)
        if return_shipping is not None:
            pulumi.set(__self__, "return_shipping", return_shipping)
        if shipping_information is not None:
            pulumi.set(__self__, "shipping_information", shipping_information)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)

    @property
    @pulumi.getter(name="backupDriveManifest")
    def backup_drive_manifest(self) -> Optional[pulumi.Input[bool]]:
        """
        Default value is false. Indicates whether the manifest files on the drives should be copied to block blobs.
        """
        return pulumi.get(self, "backup_drive_manifest")

    @backup_drive_manifest.setter
    def backup_drive_manifest(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "backup_drive_manifest", value)

    @property
    @pulumi.getter(name="cancelRequested")
    def cancel_requested(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether a request has been submitted to cancel the job.
        """
        return pulumi.get(self, "cancel_requested")

    @cancel_requested.setter
    def cancel_requested(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cancel_requested", value)

    @property
    @pulumi.getter(name="deliveryPackage")
    def delivery_package(self) -> Optional[pulumi.Input['PackageInfomationArgs']]:
        """
        Contains information about the package being shipped by the customer to the Microsoft data center. 
        """
        return pulumi.get(self, "delivery_package")

    @delivery_package.setter
    def delivery_package(self, value: Optional[pulumi.Input['PackageInfomationArgs']]):
        pulumi.set(self, "delivery_package", value)

    @property
    @pulumi.getter(name="diagnosticsPath")
    def diagnostics_path(self) -> Optional[pulumi.Input[str]]:
        """
        The virtual blob directory to which the copy logs and backups of drive manifest files (if enabled) will be stored.
        """
        return pulumi.get(self, "diagnostics_path")

    @diagnostics_path.setter
    def diagnostics_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diagnostics_path", value)

    @property
    @pulumi.getter(name="driveList")
    def drive_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DriveStatusArgs']]]]:
        """
        List of up to ten drives that comprise the job. The drive list is a required element for an import job; it is not specified for export jobs.
        """
        return pulumi.get(self, "drive_list")

    @drive_list.setter
    def drive_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DriveStatusArgs']]]]):
        pulumi.set(self, "drive_list", value)

    @property
    @pulumi.getter
    def export(self) -> Optional[pulumi.Input['ExportArgs']]:
        """
        A property containing information about the blobs to be exported for an export job. This property is included for export jobs only.
        """
        return pulumi.get(self, "export")

    @export.setter
    def export(self, value: Optional[pulumi.Input['ExportArgs']]):
        pulumi.set(self, "export", value)

    @property
    @pulumi.getter(name="incompleteBlobListUri")
    def incomplete_blob_list_uri(self) -> Optional[pulumi.Input[str]]:
        """
        A blob path that points to a block blob containing a list of blob names that were not exported due to insufficient drive space. If all blobs were exported successfully, then this element is not included in the response.
        """
        return pulumi.get(self, "incomplete_blob_list_uri")

    @incomplete_blob_list_uri.setter
    def incomplete_blob_list_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incomplete_blob_list_uri", value)

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of job
        """
        return pulumi.get(self, "job_type")

    @job_type.setter
    def job_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_type", value)

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[pulumi.Input[str]]:
        """
        Default value is Error. Indicates whether error logging or verbose logging will be enabled.
        """
        return pulumi.get(self, "log_level")

    @log_level.setter
    def log_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_level", value)

    @property
    @pulumi.getter(name="percentComplete")
    def percent_complete(self) -> Optional[pulumi.Input[int]]:
        """
        Overall percentage completed for the job.
        """
        return pulumi.get(self, "percent_complete")

    @percent_complete.setter
    def percent_complete(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "percent_complete", value)

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the provisioning state of the job.
        """
        return pulumi.get(self, "provisioning_state")

    @provisioning_state.setter
    def provisioning_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provisioning_state", value)

    @property
    @pulumi.getter(name="returnAddress")
    def return_address(self) -> Optional[pulumi.Input['ReturnAddressArgs']]:
        """
        Specifies the return address information for the job. 
        """
        return pulumi.get(self, "return_address")

    @return_address.setter
    def return_address(self, value: Optional[pulumi.Input['ReturnAddressArgs']]):
        pulumi.set(self, "return_address", value)

    @property
    @pulumi.getter(name="returnPackage")
    def return_package(self) -> Optional[pulumi.Input['PackageInfomationArgs']]:
        """
        Contains information about the package being shipped from the Microsoft data center to the customer to return the drives. The format is the same as the deliveryPackage property above. This property is not included if the drives have not yet been returned. 
        """
        return pulumi.get(self, "return_package")

    @return_package.setter
    def return_package(self, value: Optional[pulumi.Input['PackageInfomationArgs']]):
        pulumi.set(self, "return_package", value)

    @property
    @pulumi.getter(name="returnShipping")
    def return_shipping(self) -> Optional[pulumi.Input['ReturnShippingArgs']]:
        """
        Specifies the return carrier and customer's account with the carrier. 
        """
        return pulumi.get(self, "return_shipping")

    @return_shipping.setter
    def return_shipping(self, value: Optional[pulumi.Input['ReturnShippingArgs']]):
        pulumi.set(self, "return_shipping", value)

    @property
    @pulumi.getter(name="shippingInformation")
    def shipping_information(self) -> Optional[pulumi.Input['ShippingInformationArgs']]:
        """
        Contains information about the Microsoft datacenter to which the drives should be shipped. 
        """
        return pulumi.get(self, "shipping_information")

    @shipping_information.setter
    def shipping_information(self, value: Optional[pulumi.Input['ShippingInformationArgs']]):
        pulumi.set(self, "shipping_information", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Current state of the job.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource identifier of the storage account where data will be imported to or exported from.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)


@pulumi.input_type
class PackageInfomationArgs:
    def __init__(__self__, *,
                 carrier_name: pulumi.Input[str],
                 drive_count: pulumi.Input[int],
                 ship_date: pulumi.Input[str],
                 tracking_number: pulumi.Input[str]):
        """
        Contains information about the package being shipped by the customer to the Microsoft data center.
        :param pulumi.Input[str] carrier_name: The name of the carrier that is used to ship the import or export drives.
        :param pulumi.Input[int] drive_count: The number of drives included in the package.
        :param pulumi.Input[str] ship_date: The date when the package is shipped.
        :param pulumi.Input[str] tracking_number: The tracking number of the package.
        """
        pulumi.set(__self__, "carrier_name", carrier_name)
        pulumi.set(__self__, "drive_count", drive_count)
        pulumi.set(__self__, "ship_date", ship_date)
        pulumi.set(__self__, "tracking_number", tracking_number)

    @property
    @pulumi.getter(name="carrierName")
    def carrier_name(self) -> pulumi.Input[str]:
        """
        The name of the carrier that is used to ship the import or export drives.
        """
        return pulumi.get(self, "carrier_name")

    @carrier_name.setter
    def carrier_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "carrier_name", value)

    @property
    @pulumi.getter(name="driveCount")
    def drive_count(self) -> pulumi.Input[int]:
        """
        The number of drives included in the package.
        """
        return pulumi.get(self, "drive_count")

    @drive_count.setter
    def drive_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "drive_count", value)

    @property
    @pulumi.getter(name="shipDate")
    def ship_date(self) -> pulumi.Input[str]:
        """
        The date when the package is shipped.
        """
        return pulumi.get(self, "ship_date")

    @ship_date.setter
    def ship_date(self, value: pulumi.Input[str]):
        pulumi.set(self, "ship_date", value)

    @property
    @pulumi.getter(name="trackingNumber")
    def tracking_number(self) -> pulumi.Input[str]:
        """
        The tracking number of the package.
        """
        return pulumi.get(self, "tracking_number")

    @tracking_number.setter
    def tracking_number(self, value: pulumi.Input[str]):
        pulumi.set(self, "tracking_number", value)


@pulumi.input_type
class ReturnAddressArgs:
    def __init__(__self__, *,
                 city: pulumi.Input[str],
                 country_or_region: pulumi.Input[str],
                 email: pulumi.Input[str],
                 phone: pulumi.Input[str],
                 postal_code: pulumi.Input[str],
                 recipient_name: pulumi.Input[str],
                 street_address1: pulumi.Input[str],
                 state_or_province: Optional[pulumi.Input[str]] = None,
                 street_address2: Optional[pulumi.Input[str]] = None):
        """
        Specifies the return address information for the job.
        :param pulumi.Input[str] city: The city name to use when returning the drives.
        :param pulumi.Input[str] country_or_region: The country or region to use when returning the drives. 
        :param pulumi.Input[str] email: Email address of the recipient of the returned drives.
        :param pulumi.Input[str] phone: Phone number of the recipient of the returned drives.
        :param pulumi.Input[str] postal_code: The postal code to use when returning the drives.
        :param pulumi.Input[str] recipient_name: The name of the recipient who will receive the hard drives when they are returned. 
        :param pulumi.Input[str] street_address1: The first line of the street address to use when returning the drives. 
        :param pulumi.Input[str] state_or_province: The state or province to use when returning the drives.
        :param pulumi.Input[str] street_address2: The second line of the street address to use when returning the drives. 
        """
        pulumi.set(__self__, "city", city)
        pulumi.set(__self__, "country_or_region", country_or_region)
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "phone", phone)
        pulumi.set(__self__, "postal_code", postal_code)
        pulumi.set(__self__, "recipient_name", recipient_name)
        pulumi.set(__self__, "street_address1", street_address1)
        if state_or_province is not None:
            pulumi.set(__self__, "state_or_province", state_or_province)
        if street_address2 is not None:
            pulumi.set(__self__, "street_address2", street_address2)

    @property
    @pulumi.getter
    def city(self) -> pulumi.Input[str]:
        """
        The city name to use when returning the drives.
        """
        return pulumi.get(self, "city")

    @city.setter
    def city(self, value: pulumi.Input[str]):
        pulumi.set(self, "city", value)

    @property
    @pulumi.getter(name="countryOrRegion")
    def country_or_region(self) -> pulumi.Input[str]:
        """
        The country or region to use when returning the drives. 
        """
        return pulumi.get(self, "country_or_region")

    @country_or_region.setter
    def country_or_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "country_or_region", value)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        Email address of the recipient of the returned drives.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Input[str]:
        """
        Phone number of the recipient of the returned drives.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: pulumi.Input[str]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> pulumi.Input[str]:
        """
        The postal code to use when returning the drives.
        """
        return pulumi.get(self, "postal_code")

    @postal_code.setter
    def postal_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "postal_code", value)

    @property
    @pulumi.getter(name="recipientName")
    def recipient_name(self) -> pulumi.Input[str]:
        """
        The name of the recipient who will receive the hard drives when they are returned. 
        """
        return pulumi.get(self, "recipient_name")

    @recipient_name.setter
    def recipient_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "recipient_name", value)

    @property
    @pulumi.getter(name="streetAddress1")
    def street_address1(self) -> pulumi.Input[str]:
        """
        The first line of the street address to use when returning the drives. 
        """
        return pulumi.get(self, "street_address1")

    @street_address1.setter
    def street_address1(self, value: pulumi.Input[str]):
        pulumi.set(self, "street_address1", value)

    @property
    @pulumi.getter(name="stateOrProvince")
    def state_or_province(self) -> Optional[pulumi.Input[str]]:
        """
        The state or province to use when returning the drives.
        """
        return pulumi.get(self, "state_or_province")

    @state_or_province.setter
    def state_or_province(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_or_province", value)

    @property
    @pulumi.getter(name="streetAddress2")
    def street_address2(self) -> Optional[pulumi.Input[str]]:
        """
        The second line of the street address to use when returning the drives. 
        """
        return pulumi.get(self, "street_address2")

    @street_address2.setter
    def street_address2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "street_address2", value)


@pulumi.input_type
class ReturnShippingArgs:
    def __init__(__self__, *,
                 carrier_account_number: pulumi.Input[str],
                 carrier_name: pulumi.Input[str]):
        """
        Specifies the return carrier and customer's account with the carrier.
        :param pulumi.Input[str] carrier_account_number: The customer's account number with the carrier.
        :param pulumi.Input[str] carrier_name: The carrier's name.
        """
        pulumi.set(__self__, "carrier_account_number", carrier_account_number)
        pulumi.set(__self__, "carrier_name", carrier_name)

    @property
    @pulumi.getter(name="carrierAccountNumber")
    def carrier_account_number(self) -> pulumi.Input[str]:
        """
        The customer's account number with the carrier.
        """
        return pulumi.get(self, "carrier_account_number")

    @carrier_account_number.setter
    def carrier_account_number(self, value: pulumi.Input[str]):
        pulumi.set(self, "carrier_account_number", value)

    @property
    @pulumi.getter(name="carrierName")
    def carrier_name(self) -> pulumi.Input[str]:
        """
        The carrier's name.
        """
        return pulumi.get(self, "carrier_name")

    @carrier_name.setter
    def carrier_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "carrier_name", value)


@pulumi.input_type
class ShippingInformationArgs:
    def __init__(__self__, *,
                 city: pulumi.Input[str],
                 country_or_region: pulumi.Input[str],
                 postal_code: pulumi.Input[str],
                 recipient_name: pulumi.Input[str],
                 state_or_province: pulumi.Input[str],
                 street_address1: pulumi.Input[str],
                 phone: Optional[pulumi.Input[str]] = None,
                 street_address2: Optional[pulumi.Input[str]] = None):
        """
        Contains information about the Microsoft datacenter to which the drives should be shipped.
        :param pulumi.Input[str] city: The city name to use when returning the drives.
        :param pulumi.Input[str] country_or_region: The country or region to use when returning the drives. 
        :param pulumi.Input[str] postal_code: The postal code to use when returning the drives.
        :param pulumi.Input[str] recipient_name: The name of the recipient who will receive the hard drives when they are returned. 
        :param pulumi.Input[str] state_or_province: The state or province to use when returning the drives.
        :param pulumi.Input[str] street_address1: The first line of the street address to use when returning the drives. 
        :param pulumi.Input[str] phone: Phone number of the recipient of the returned drives.
        :param pulumi.Input[str] street_address2: The second line of the street address to use when returning the drives. 
        """
        pulumi.set(__self__, "city", city)
        pulumi.set(__self__, "country_or_region", country_or_region)
        pulumi.set(__self__, "postal_code", postal_code)
        pulumi.set(__self__, "recipient_name", recipient_name)
        pulumi.set(__self__, "state_or_province", state_or_province)
        pulumi.set(__self__, "street_address1", street_address1)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if street_address2 is not None:
            pulumi.set(__self__, "street_address2", street_address2)

    @property
    @pulumi.getter
    def city(self) -> pulumi.Input[str]:
        """
        The city name to use when returning the drives.
        """
        return pulumi.get(self, "city")

    @city.setter
    def city(self, value: pulumi.Input[str]):
        pulumi.set(self, "city", value)

    @property
    @pulumi.getter(name="countryOrRegion")
    def country_or_region(self) -> pulumi.Input[str]:
        """
        The country or region to use when returning the drives. 
        """
        return pulumi.get(self, "country_or_region")

    @country_or_region.setter
    def country_or_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "country_or_region", value)

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> pulumi.Input[str]:
        """
        The postal code to use when returning the drives.
        """
        return pulumi.get(self, "postal_code")

    @postal_code.setter
    def postal_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "postal_code", value)

    @property
    @pulumi.getter(name="recipientName")
    def recipient_name(self) -> pulumi.Input[str]:
        """
        The name of the recipient who will receive the hard drives when they are returned. 
        """
        return pulumi.get(self, "recipient_name")

    @recipient_name.setter
    def recipient_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "recipient_name", value)

    @property
    @pulumi.getter(name="stateOrProvince")
    def state_or_province(self) -> pulumi.Input[str]:
        """
        The state or province to use when returning the drives.
        """
        return pulumi.get(self, "state_or_province")

    @state_or_province.setter
    def state_or_province(self, value: pulumi.Input[str]):
        pulumi.set(self, "state_or_province", value)

    @property
    @pulumi.getter(name="streetAddress1")
    def street_address1(self) -> pulumi.Input[str]:
        """
        The first line of the street address to use when returning the drives. 
        """
        return pulumi.get(self, "street_address1")

    @street_address1.setter
    def street_address1(self, value: pulumi.Input[str]):
        pulumi.set(self, "street_address1", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[str]]:
        """
        Phone number of the recipient of the returned drives.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter(name="streetAddress2")
    def street_address2(self) -> Optional[pulumi.Input[str]]:
        """
        The second line of the street address to use when returning the drives. 
        """
        return pulumi.get(self, "street_address2")

    @street_address2.setter
    def street_address2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "street_address2", value)


