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
    'ContactDetailsResponse',
    'CopyLogDetailsResponse',
    'DestinationAccountDetailsResponse',
    'ErrorResponse',
    'JobDetailsResponse',
    'JobErrorDetailsResponse',
    'JobSecretsResponseResult',
    'JobStagesResponse',
    'NotificationPreferenceResponse',
    'PackageShippingDetailsResponse',
    'PreferencesResponse',
    'ShippingAddressResponse',
    'SkuResponse',
    'UnencryptedCredentialsResponseResult',
]

@pulumi.output_type
class ContactDetailsResponse(dict):
    """
    Contact Details.
    """
    def __init__(__self__, *,
                 contact_name: str,
                 email_list: List[str],
                 phone: str,
                 mobile: Optional[str] = None,
                 notification_preference: Optional[List['outputs.NotificationPreferenceResponse']] = None,
                 phone_extension: Optional[str] = None):
        """
        Contact Details.
        :param str contact_name: Contact name of the person.
        :param List[str] email_list: List of Email-ids to be notified about job progress.
        :param str phone: Phone number of the contact person.
        :param str mobile: Mobile number of the contact person.
        :param List['NotificationPreferenceResponseArgs'] notification_preference: Notification preference for a job stage.
        :param str phone_extension: Phone extension number of the contact person.
        """
        pulumi.set(__self__, "contact_name", contact_name)
        pulumi.set(__self__, "email_list", email_list)
        pulumi.set(__self__, "phone", phone)
        if mobile is not None:
            pulumi.set(__self__, "mobile", mobile)
        if notification_preference is not None:
            pulumi.set(__self__, "notification_preference", notification_preference)
        if phone_extension is not None:
            pulumi.set(__self__, "phone_extension", phone_extension)

    @property
    @pulumi.getter(name="contactName")
    def contact_name(self) -> str:
        """
        Contact name of the person.
        """
        return pulumi.get(self, "contact_name")

    @property
    @pulumi.getter(name="emailList")
    def email_list(self) -> List[str]:
        """
        List of Email-ids to be notified about job progress.
        """
        return pulumi.get(self, "email_list")

    @property
    @pulumi.getter
    def phone(self) -> str:
        """
        Phone number of the contact person.
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter
    def mobile(self) -> Optional[str]:
        """
        Mobile number of the contact person.
        """
        return pulumi.get(self, "mobile")

    @property
    @pulumi.getter(name="notificationPreference")
    def notification_preference(self) -> Optional[List['outputs.NotificationPreferenceResponse']]:
        """
        Notification preference for a job stage.
        """
        return pulumi.get(self, "notification_preference")

    @property
    @pulumi.getter(name="phoneExtension")
    def phone_extension(self) -> Optional[str]:
        """
        Phone extension number of the contact person.
        """
        return pulumi.get(self, "phone_extension")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CopyLogDetailsResponse(dict):
    """
    Details for log generated during copy.
    """
    def __init__(__self__, *,
                 copy_log_details_type: str):
        """
        Details for log generated during copy.
        :param str copy_log_details_type: Indicates the type of job details.
        """
        pulumi.set(__self__, "copy_log_details_type", copy_log_details_type)

    @property
    @pulumi.getter(name="copyLogDetailsType")
    def copy_log_details_type(self) -> str:
        """
        Indicates the type of job details.
        """
        return pulumi.get(self, "copy_log_details_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DestinationAccountDetailsResponse(dict):
    """
    Details of the destination of the data
    """
    def __init__(__self__, *,
                 data_destination_type: str,
                 account_id: Optional[str] = None):
        """
        Details of the destination of the data
        :param str data_destination_type: Data Destination Type.
        :param str account_id: Arm Id of the destination where the data has to be moved.
        """
        pulumi.set(__self__, "data_destination_type", data_destination_type)
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)

    @property
    @pulumi.getter(name="dataDestinationType")
    def data_destination_type(self) -> str:
        """
        Data Destination Type.
        """
        return pulumi.get(self, "data_destination_type")

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[str]:
        """
        Arm Id of the destination where the data has to be moved.
        """
        return pulumi.get(self, "account_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ErrorResponse(dict):
    """
    Top level error for the job.
    """
    def __init__(__self__, *,
                 code: str,
                 message: str):
        """
        Top level error for the job.
        :param str code: Error code that can be used to programmatically identify the error.
        :param str message: Describes the error in detail and provides debugging information.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Error code that can be used to programmatically identify the error.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        Describes the error in detail and provides debugging information.
        """
        return pulumi.get(self, "message")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobDetailsResponse(dict):
    """
    Job details.
    """
    def __init__(__self__, *,
                 chain_of_custody_sas_key: str,
                 contact_details: 'outputs.ContactDetailsResponse',
                 copy_log_details: List['outputs.CopyLogDetailsResponse'],
                 delivery_package: 'outputs.PackageShippingDetailsResponse',
                 destination_account_details: List['outputs.DestinationAccountDetailsResponse'],
                 error_details: List['outputs.JobErrorDetailsResponse'],
                 job_details_type: str,
                 job_stages: List['outputs.JobStagesResponse'],
                 return_package: 'outputs.PackageShippingDetailsResponse',
                 reverse_shipment_label_sas_key: str,
                 shipping_address: 'outputs.ShippingAddressResponse',
                 expected_data_size_in_tera_bytes: Optional[float] = None,
                 preferences: Optional['outputs.PreferencesResponse'] = None):
        """
        Job details.
        :param str chain_of_custody_sas_key: Shared access key to download the chain of custody logs
        :param 'ContactDetailsResponseArgs' contact_details: Contact details for notification and shipping.
        :param List['CopyLogDetailsResponseArgs'] copy_log_details: List of copy log details.
        :param 'PackageShippingDetailsResponseArgs' delivery_package: Delivery package shipping details.
        :param List['DestinationAccountDetailsResponseArgs'] destination_account_details: Destination account details.
        :param List['JobErrorDetailsResponseArgs'] error_details: Error details for failure. This is optional.
        :param str job_details_type: Indicates the type of job details.
        :param List['JobStagesResponseArgs'] job_stages: List of stages that run in the job.
        :param 'PackageShippingDetailsResponseArgs' return_package: Return package shipping details.
        :param str reverse_shipment_label_sas_key: Shared access key to download the return shipment label
        :param 'ShippingAddressResponseArgs' shipping_address: Shipping address of the customer.
        :param float expected_data_size_in_tera_bytes: The expected size of the data, which needs to be transferred in this job, in terabytes.
        :param 'PreferencesResponseArgs' preferences: Preferences for the order.
        """
        pulumi.set(__self__, "chain_of_custody_sas_key", chain_of_custody_sas_key)
        pulumi.set(__self__, "contact_details", contact_details)
        pulumi.set(__self__, "copy_log_details", copy_log_details)
        pulumi.set(__self__, "delivery_package", delivery_package)
        pulumi.set(__self__, "destination_account_details", destination_account_details)
        pulumi.set(__self__, "error_details", error_details)
        pulumi.set(__self__, "job_details_type", job_details_type)
        pulumi.set(__self__, "job_stages", job_stages)
        pulumi.set(__self__, "return_package", return_package)
        pulumi.set(__self__, "reverse_shipment_label_sas_key", reverse_shipment_label_sas_key)
        pulumi.set(__self__, "shipping_address", shipping_address)
        if expected_data_size_in_tera_bytes is not None:
            pulumi.set(__self__, "expected_data_size_in_tera_bytes", expected_data_size_in_tera_bytes)
        if preferences is not None:
            pulumi.set(__self__, "preferences", preferences)

    @property
    @pulumi.getter(name="chainOfCustodySasKey")
    def chain_of_custody_sas_key(self) -> str:
        """
        Shared access key to download the chain of custody logs
        """
        return pulumi.get(self, "chain_of_custody_sas_key")

    @property
    @pulumi.getter(name="contactDetails")
    def contact_details(self) -> 'outputs.ContactDetailsResponse':
        """
        Contact details for notification and shipping.
        """
        return pulumi.get(self, "contact_details")

    @property
    @pulumi.getter(name="copyLogDetails")
    def copy_log_details(self) -> List['outputs.CopyLogDetailsResponse']:
        """
        List of copy log details.
        """
        return pulumi.get(self, "copy_log_details")

    @property
    @pulumi.getter(name="deliveryPackage")
    def delivery_package(self) -> 'outputs.PackageShippingDetailsResponse':
        """
        Delivery package shipping details.
        """
        return pulumi.get(self, "delivery_package")

    @property
    @pulumi.getter(name="destinationAccountDetails")
    def destination_account_details(self) -> List['outputs.DestinationAccountDetailsResponse']:
        """
        Destination account details.
        """
        return pulumi.get(self, "destination_account_details")

    @property
    @pulumi.getter(name="errorDetails")
    def error_details(self) -> List['outputs.JobErrorDetailsResponse']:
        """
        Error details for failure. This is optional.
        """
        return pulumi.get(self, "error_details")

    @property
    @pulumi.getter(name="jobDetailsType")
    def job_details_type(self) -> str:
        """
        Indicates the type of job details.
        """
        return pulumi.get(self, "job_details_type")

    @property
    @pulumi.getter(name="jobStages")
    def job_stages(self) -> List['outputs.JobStagesResponse']:
        """
        List of stages that run in the job.
        """
        return pulumi.get(self, "job_stages")

    @property
    @pulumi.getter(name="returnPackage")
    def return_package(self) -> 'outputs.PackageShippingDetailsResponse':
        """
        Return package shipping details.
        """
        return pulumi.get(self, "return_package")

    @property
    @pulumi.getter(name="reverseShipmentLabelSasKey")
    def reverse_shipment_label_sas_key(self) -> str:
        """
        Shared access key to download the return shipment label
        """
        return pulumi.get(self, "reverse_shipment_label_sas_key")

    @property
    @pulumi.getter(name="shippingAddress")
    def shipping_address(self) -> 'outputs.ShippingAddressResponse':
        """
        Shipping address of the customer.
        """
        return pulumi.get(self, "shipping_address")

    @property
    @pulumi.getter(name="expectedDataSizeInTeraBytes")
    def expected_data_size_in_tera_bytes(self) -> Optional[float]:
        """
        The expected size of the data, which needs to be transferred in this job, in terabytes.
        """
        return pulumi.get(self, "expected_data_size_in_tera_bytes")

    @property
    @pulumi.getter
    def preferences(self) -> Optional['outputs.PreferencesResponse']:
        """
        Preferences for the order.
        """
        return pulumi.get(self, "preferences")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobErrorDetailsResponse(dict):
    """
    Job Error Details for providing the information and recommended action.
    """
    def __init__(__self__, *,
                 error_code: float,
                 error_message: str,
                 exception_message: str,
                 recommended_action: str):
        """
        Job Error Details for providing the information and recommended action.
        :param float error_code: Code for the error.
        :param str error_message: Message for the error.
        :param str exception_message: Contains the non localized exception message
        :param str recommended_action: Recommended action for the error.
        """
        pulumi.set(__self__, "error_code", error_code)
        pulumi.set(__self__, "error_message", error_message)
        pulumi.set(__self__, "exception_message", exception_message)
        pulumi.set(__self__, "recommended_action", recommended_action)

    @property
    @pulumi.getter(name="errorCode")
    def error_code(self) -> float:
        """
        Code for the error.
        """
        return pulumi.get(self, "error_code")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> str:
        """
        Message for the error.
        """
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter(name="exceptionMessage")
    def exception_message(self) -> str:
        """
        Contains the non localized exception message
        """
        return pulumi.get(self, "exception_message")

    @property
    @pulumi.getter(name="recommendedAction")
    def recommended_action(self) -> str:
        """
        Recommended action for the error.
        """
        return pulumi.get(self, "recommended_action")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobSecretsResponseResult(dict):
    """
    The base class for the secrets
    """
    def __init__(__self__, *,
                 job_secrets_type: str):
        """
        The base class for the secrets
        :param str job_secrets_type: Used to indicate what type of job secrets object.
        """
        pulumi.set(__self__, "job_secrets_type", job_secrets_type)

    @property
    @pulumi.getter(name="jobSecretsType")
    def job_secrets_type(self) -> str:
        """
        Used to indicate what type of job secrets object.
        """
        return pulumi.get(self, "job_secrets_type")


@pulumi.output_type
class JobStagesResponse(dict):
    """
    Job stages.
    """
    def __init__(__self__, *,
                 display_name: str,
                 error_details: List['outputs.JobErrorDetailsResponse'],
                 job_stage_details: Mapping[str, Any],
                 stage_name: str,
                 stage_status: str,
                 stage_time: str):
        """
        Job stages.
        :param str display_name: Display name of the job stage.
        :param List['JobErrorDetailsResponseArgs'] error_details: Error details for the stage.
        :param Mapping[str, Any] job_stage_details: Job Stage Details
        :param str stage_name: Name of the job stage.
        :param str stage_status: Status of the job stage.
        :param str stage_time: Time for the job stage in UTC ISO 8601 format.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "error_details", error_details)
        pulumi.set(__self__, "job_stage_details", job_stage_details)
        pulumi.set(__self__, "stage_name", stage_name)
        pulumi.set(__self__, "stage_status", stage_status)
        pulumi.set(__self__, "stage_time", stage_time)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name of the job stage.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="errorDetails")
    def error_details(self) -> List['outputs.JobErrorDetailsResponse']:
        """
        Error details for the stage.
        """
        return pulumi.get(self, "error_details")

    @property
    @pulumi.getter(name="jobStageDetails")
    def job_stage_details(self) -> Mapping[str, Any]:
        """
        Job Stage Details
        """
        return pulumi.get(self, "job_stage_details")

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> str:
        """
        Name of the job stage.
        """
        return pulumi.get(self, "stage_name")

    @property
    @pulumi.getter(name="stageStatus")
    def stage_status(self) -> str:
        """
        Status of the job stage.
        """
        return pulumi.get(self, "stage_status")

    @property
    @pulumi.getter(name="stageTime")
    def stage_time(self) -> str:
        """
        Time for the job stage in UTC ISO 8601 format.
        """
        return pulumi.get(self, "stage_time")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class NotificationPreferenceResponse(dict):
    """
    Notification preference for a job stage.
    """
    def __init__(__self__, *,
                 send_notification: bool,
                 stage_name: str):
        """
        Notification preference for a job stage.
        :param bool send_notification: Notification is required or not.
        :param str stage_name: Name of the stage.
        """
        pulumi.set(__self__, "send_notification", send_notification)
        pulumi.set(__self__, "stage_name", stage_name)

    @property
    @pulumi.getter(name="sendNotification")
    def send_notification(self) -> bool:
        """
        Notification is required or not.
        """
        return pulumi.get(self, "send_notification")

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> str:
        """
        Name of the stage.
        """
        return pulumi.get(self, "stage_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PackageShippingDetailsResponse(dict):
    """
    Shipping details.
    """
    def __init__(__self__, *,
                 carrier_name: str,
                 tracking_id: str,
                 tracking_url: str):
        """
        Shipping details.
        :param str carrier_name: Name of the carrier.
        :param str tracking_id: Tracking Id of shipment.
        :param str tracking_url: Url where shipment can be tracked.
        """
        pulumi.set(__self__, "carrier_name", carrier_name)
        pulumi.set(__self__, "tracking_id", tracking_id)
        pulumi.set(__self__, "tracking_url", tracking_url)

    @property
    @pulumi.getter(name="carrierName")
    def carrier_name(self) -> str:
        """
        Name of the carrier.
        """
        return pulumi.get(self, "carrier_name")

    @property
    @pulumi.getter(name="trackingId")
    def tracking_id(self) -> str:
        """
        Tracking Id of shipment.
        """
        return pulumi.get(self, "tracking_id")

    @property
    @pulumi.getter(name="trackingUrl")
    def tracking_url(self) -> str:
        """
        Url where shipment can be tracked.
        """
        return pulumi.get(self, "tracking_url")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PreferencesResponse(dict):
    """
    Preferences related to the order
    """
    def __init__(__self__, *,
                 preferred_data_center_region: Optional[List[str]] = None):
        """
        Preferences related to the order
        """
        if preferred_data_center_region is not None:
            pulumi.set(__self__, "preferred_data_center_region", preferred_data_center_region)

    @property
    @pulumi.getter(name="preferredDataCenterRegion")
    def preferred_data_center_region(self) -> Optional[List[str]]:
        return pulumi.get(self, "preferred_data_center_region")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ShippingAddressResponse(dict):
    """
    Shipping address where customer wishes to receive the device.
    """
    def __init__(__self__, *,
                 country: str,
                 postal_code: str,
                 street_address1: str,
                 address_type: Optional[str] = None,
                 city: Optional[str] = None,
                 company_name: Optional[str] = None,
                 state_or_province: Optional[str] = None,
                 street_address2: Optional[str] = None,
                 street_address3: Optional[str] = None,
                 zip_extended_code: Optional[str] = None):
        """
        Shipping address where customer wishes to receive the device.
        :param str country: Name of the Country.
        :param str postal_code: Postal code.
        :param str street_address1: Street Address line 1.
        :param str address_type: Type of address.
        :param str city: Name of the City.
        :param str company_name: Name of the company.
        :param str state_or_province: Name of the State or Province.
        :param str street_address2: Street Address line 2.
        :param str street_address3: Street Address line 3.
        :param str zip_extended_code: Extended Zip Code.
        """
        pulumi.set(__self__, "country", country)
        pulumi.set(__self__, "postal_code", postal_code)
        pulumi.set(__self__, "street_address1", street_address1)
        if address_type is not None:
            pulumi.set(__self__, "address_type", address_type)
        if city is not None:
            pulumi.set(__self__, "city", city)
        if company_name is not None:
            pulumi.set(__self__, "company_name", company_name)
        if state_or_province is not None:
            pulumi.set(__self__, "state_or_province", state_or_province)
        if street_address2 is not None:
            pulumi.set(__self__, "street_address2", street_address2)
        if street_address3 is not None:
            pulumi.set(__self__, "street_address3", street_address3)
        if zip_extended_code is not None:
            pulumi.set(__self__, "zip_extended_code", zip_extended_code)

    @property
    @pulumi.getter
    def country(self) -> str:
        """
        Name of the Country.
        """
        return pulumi.get(self, "country")

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> str:
        """
        Postal code.
        """
        return pulumi.get(self, "postal_code")

    @property
    @pulumi.getter(name="streetAddress1")
    def street_address1(self) -> str:
        """
        Street Address line 1.
        """
        return pulumi.get(self, "street_address1")

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> Optional[str]:
        """
        Type of address.
        """
        return pulumi.get(self, "address_type")

    @property
    @pulumi.getter
    def city(self) -> Optional[str]:
        """
        Name of the City.
        """
        return pulumi.get(self, "city")

    @property
    @pulumi.getter(name="companyName")
    def company_name(self) -> Optional[str]:
        """
        Name of the company.
        """
        return pulumi.get(self, "company_name")

    @property
    @pulumi.getter(name="stateOrProvince")
    def state_or_province(self) -> Optional[str]:
        """
        Name of the State or Province.
        """
        return pulumi.get(self, "state_or_province")

    @property
    @pulumi.getter(name="streetAddress2")
    def street_address2(self) -> Optional[str]:
        """
        Street Address line 2.
        """
        return pulumi.get(self, "street_address2")

    @property
    @pulumi.getter(name="streetAddress3")
    def street_address3(self) -> Optional[str]:
        """
        Street Address line 3.
        """
        return pulumi.get(self, "street_address3")

    @property
    @pulumi.getter(name="zipExtendedCode")
    def zip_extended_code(self) -> Optional[str]:
        """
        Extended Zip Code.
        """
        return pulumi.get(self, "zip_extended_code")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    The Sku.
    """
    def __init__(__self__, *,
                 name: str,
                 display_name: Optional[str] = None,
                 family: Optional[str] = None):
        """
        The Sku.
        :param str name: The sku name.
        :param str display_name: The display name of the sku.
        :param str family: The sku family.
        """
        pulumi.set(__self__, "name", name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if family is not None:
            pulumi.set(__self__, "family", family)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The sku name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The display name of the sku.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def family(self) -> Optional[str]:
        """
        The sku family.
        """
        return pulumi.get(self, "family")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UnencryptedCredentialsResponseResult(dict):
    """
    Unencrypted credentials for accessing device.
    """
    def __init__(__self__, *,
                 job_name: str,
                 job_secrets: 'outputs.JobSecretsResponseResult'):
        """
        Unencrypted credentials for accessing device.
        :param str job_name: Name of the job.
        :param 'JobSecretsResponseArgs' job_secrets: Secrets related to this job.
        """
        pulumi.set(__self__, "job_name", job_name)
        pulumi.set(__self__, "job_secrets", job_secrets)

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> str:
        """
        Name of the job.
        """
        return pulumi.get(self, "job_name")

    @property
    @pulumi.getter(name="jobSecrets")
    def job_secrets(self) -> 'outputs.JobSecretsResponseResult':
        """
        Secrets related to this job.
        """
        return pulumi.get(self, "job_secrets")


