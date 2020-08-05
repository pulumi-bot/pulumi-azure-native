# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Job(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
    """
    name: pulumi.Output[str]
    """
    Name of the object.
    """
    properties: pulumi.Output[dict]
    """
    Properties of a job.
      * `cancellation_reason` (`str`) - Reason for cancellation.
      * `delivery_info` (`dict`) - Delivery Info of Job.
        * `scheduled_date_time` (`str`) - Scheduled date time.

      * `delivery_type` (`str`) - Delivery type of Job.
      * `details` (`dict`) - Details of a job run. This field will only be sent for expand details filter.
        * `chain_of_custody_sas_key` (`str`) - Shared access key to download the chain of custody logs
        * `contact_details` (`dict`) - Contact details for notification and shipping.
          * `contact_name` (`str`) - Contact name of the person.
          * `email_list` (`list`) - List of Email-ids to be notified about job progress.
          * `mobile` (`str`) - Mobile number of the contact person.
          * `notification_preference` (`list`) - Notification preference for a job stage.
            * `send_notification` (`bool`) - Notification is required or not.
            * `stage_name` (`str`) - Name of the stage.

          * `phone` (`str`) - Phone number of the contact person.
          * `phone_extension` (`str`) - Phone extension number of the contact person.

        * `copy_log_details` (`list`) - List of copy log details.
          * `copy_log_details_type` (`str`) - Indicates the type of job details.

        * `delivery_package` (`dict`) - Delivery package shipping details.
          * `carrier_name` (`str`) - Name of the carrier.
          * `tracking_id` (`str`) - Tracking Id of shipment.
          * `tracking_url` (`str`) - Url where shipment can be tracked.

        * `destination_account_details` (`list`) - Destination account details.
          * `account_id` (`str`) - Arm Id of the destination where the data has to be moved.
          * `data_destination_type` (`str`) - Data Destination Type.
          * `share_password` (`str`) - Share password to be shared by all shares in SA.

        * `error_details` (`list`) - Error details for failure. This is optional.
          * `error_code` (`float`) - Code for the error.
          * `error_message` (`str`) - Message for the error.
          * `exception_message` (`str`) - Contains the non localized exception message
          * `recommended_action` (`str`) - Recommended action for the error.

        * `expected_data_size_in_terabytes` (`float`) - The expected size of the data, which needs to be transferred in this job, in terabytes.
        * `job_details_type` (`str`) - Indicates the type of job details.
        * `job_stages` (`list`) - List of stages that run in the job.
          * `display_name` (`str`) - Display name of the job stage.
          * `error_details` (`list`) - Error details for the stage.
          * `job_stage_details` (`dict`) - Job Stage Details
          * `stage_name` (`str`) - Name of the job stage.
          * `stage_status` (`str`) - Status of the job stage.
          * `stage_time` (`str`) - Time for the job stage in UTC ISO 8601 format.

        * `preferences` (`dict`) - Preferences for the order.
          * `preferred_data_center_region` (`list`) - Preferred Data Center Region.
          * `transport_preferences` (`dict`) - Preferences related to the shipment logistics of the sku.
            * `preferred_shipment_type` (`str`) - Indicates Shipment Logistics type that the customer preferred.

        * `return_package` (`dict`) - Return package shipping details.
        * `reverse_shipment_label_sas_key` (`str`) - Shared access key to download the return shipment label
        * `shipping_address` (`dict`) - Shipping address of the customer.
          * `address_type` (`str`) - Type of address.
          * `city` (`str`) - Name of the City.
          * `company_name` (`str`) - Name of the company.
          * `country` (`str`) - Name of the Country.
          * `postal_code` (`str`) - Postal code.
          * `state_or_province` (`str`) - Name of the State or Province.
          * `street_address1` (`str`) - Street Address line 1.
          * `street_address2` (`str`) - Street Address line 2.
          * `street_address3` (`str`) - Street Address line 3.
          * `zip_extended_code` (`str`) - Extended Zip Code.

      * `error` (`dict`) - Top level error for the job.
        * `code` (`str`) - Error code that can be used to programmatically identify the error.
        * `message` (`str`) - Describes the error in detail and provides debugging information.

      * `is_cancellable` (`bool`) - Describes whether the job is cancellable or not.
      * `is_cancellable_without_fee` (`bool`) - Flag to indicate cancellation of scheduled job.
      * `is_deletable` (`bool`) - Describes whether the job is deletable or not.
      * `is_shipping_address_editable` (`bool`) - Describes whether the shipping address is editable or not.
      * `start_time` (`str`) - Time at which the job was started in UTC ISO 8601 format.
      * `status` (`str`) - Name of the stage which is in progress.
    """
    sku: pulumi.Output[dict]
    """
    The sku type.
      * `display_name` (`str`) - The display name of the sku.
      * `family` (`str`) - The sku family.
      * `name` (`str`) - The sku name.
    """
    tags: pulumi.Output[dict]
    """
    The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
    """
    type: pulumi.Output[str]
    """
    Type of the object.
    """
    def __init__(__self__, resource_name, opts=None, delivery_info=None, delivery_type=None, details=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Job Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] delivery_info: Delivery Info of Job.
        :param pulumi.Input[str] delivery_type: Delivery type of Job.
        :param pulumi.Input[dict] details: Details of a job run. This field will only be sent for expand details filter.
        :param pulumi.Input[str] location: The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
        :param pulumi.Input[str] name: The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        :param pulumi.Input[str] resource_group_name: The Resource Group Name
        :param pulumi.Input[dict] sku: The sku type.
        :param pulumi.Input[dict] tags: The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).

        The **delivery_info** object supports the following:

          * `scheduled_date_time` (`pulumi.Input[str]`) - Scheduled date time.

        The **details** object supports the following:

          * `contact_details` (`pulumi.Input[dict]`) - Contact details for notification and shipping.
            * `contact_name` (`pulumi.Input[str]`) - Contact name of the person.
            * `email_list` (`pulumi.Input[list]`) - List of Email-ids to be notified about job progress.
            * `mobile` (`pulumi.Input[str]`) - Mobile number of the contact person.
            * `notification_preference` (`pulumi.Input[list]`) - Notification preference for a job stage.
              * `send_notification` (`pulumi.Input[bool]`) - Notification is required or not.
              * `stage_name` (`pulumi.Input[str]`) - Name of the stage.

            * `phone` (`pulumi.Input[str]`) - Phone number of the contact person.
            * `phone_extension` (`pulumi.Input[str]`) - Phone extension number of the contact person.

          * `destination_account_details` (`pulumi.Input[list]`) - Destination account details.
            * `account_id` (`pulumi.Input[str]`) - Arm Id of the destination where the data has to be moved.
            * `data_destination_type` (`pulumi.Input[str]`) - Data Destination Type.
            * `share_password` (`pulumi.Input[str]`) - Share password to be shared by all shares in SA.

          * `expected_data_size_in_terabytes` (`pulumi.Input[float]`) - The expected size of the data, which needs to be transferred in this job, in terabytes.
          * `job_details_type` (`pulumi.Input[str]`) - Indicates the type of job details.
          * `preferences` (`pulumi.Input[dict]`) - Preferences for the order.
            * `preferred_data_center_region` (`pulumi.Input[list]`) - Preferred Data Center Region.
            * `transport_preferences` (`pulumi.Input[dict]`) - Preferences related to the shipment logistics of the sku.
              * `preferred_shipment_type` (`pulumi.Input[str]`) - Indicates Shipment Logistics type that the customer preferred.

          * `shipping_address` (`pulumi.Input[dict]`) - Shipping address of the customer.
            * `address_type` (`pulumi.Input[str]`) - Type of address.
            * `city` (`pulumi.Input[str]`) - Name of the City.
            * `company_name` (`pulumi.Input[str]`) - Name of the company.
            * `country` (`pulumi.Input[str]`) - Name of the Country.
            * `postal_code` (`pulumi.Input[str]`) - Postal code.
            * `state_or_province` (`pulumi.Input[str]`) - Name of the State or Province.
            * `street_address1` (`pulumi.Input[str]`) - Street Address line 1.
            * `street_address2` (`pulumi.Input[str]`) - Street Address line 2.
            * `street_address3` (`pulumi.Input[str]`) - Street Address line 3.
            * `zip_extended_code` (`pulumi.Input[str]`) - Extended Zip Code.

        The **sku** object supports the following:

          * `display_name` (`pulumi.Input[str]`) - The display name of the sku.
          * `family` (`pulumi.Input[str]`) - The sku family.
          * `name` (`pulumi.Input[str]`) - The sku name.
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

            __props__['delivery_info'] = delivery_info
            __props__['delivery_type'] = delivery_type
            __props__['details'] = details
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(Job, __self__).__init__(
            'azurerm:databox/v20190901:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Job(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
