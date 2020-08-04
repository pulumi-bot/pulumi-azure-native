# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Domain(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource
    """
    location: pulumi.Output[str]
    """
    Resource Location
    """
    name: pulumi.Output[str]
    """
    Resource Name
    """
    properties: pulumi.Output[dict]
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, id=None, kind=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a domain

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] kind: Kind of resource
        :param pulumi.Input[str] location: Resource Location
        :param pulumi.Input[str] name: Name of the domain
        :param pulumi.Input[str] resource_group_name: &gt;Name of the resource group
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[str] type: Resource type

        The **properties** object supports the following:

          * `auto_renew` (`pulumi.Input[bool]`) - If true then domain will renewed automatically
          * `consent` (`pulumi.Input[dict]`) - Legal agreement consent
            * `agreed_at` (`pulumi.Input[str]`) - Timestamp when the agreements were accepted
            * `agreed_by` (`pulumi.Input[str]`) - Client IP address
            * `agreement_keys` (`pulumi.Input[list]`) - List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements Api under TopLevelDomain resource

          * `contact_admin` (`pulumi.Input[dict]`) - Admin contact information
            * `address_mailing` (`pulumi.Input[dict]`) - Mailing address
              * `address1` (`pulumi.Input[str]`) - Address 1
              * `address2` (`pulumi.Input[str]`) - Address 2
              * `city` (`pulumi.Input[str]`) - City
              * `country` (`pulumi.Input[str]`) - Country
              * `postal_code` (`pulumi.Input[str]`) - Postal code
              * `state` (`pulumi.Input[str]`) - State

            * `email` (`pulumi.Input[str]`) - Email address
            * `fax` (`pulumi.Input[str]`) - Fax number
            * `job_title` (`pulumi.Input[str]`) - Job title
            * `name_first` (`pulumi.Input[str]`) - First name
            * `name_last` (`pulumi.Input[str]`) - Last name
            * `name_middle` (`pulumi.Input[str]`) - Middle name
            * `organization` (`pulumi.Input[str]`) - Organization
            * `phone` (`pulumi.Input[str]`) - Phone number

          * `contact_billing` (`pulumi.Input[dict]`) - Billing contact information
          * `contact_registrant` (`pulumi.Input[dict]`) - Registrant contact information
          * `contact_tech` (`pulumi.Input[dict]`) - Technical contact information
          * `created_time` (`pulumi.Input[str]`) - Domain creation timestamp
          * `domain_not_renewable_reasons` (`pulumi.Input[list]`) - Reasons why domain is not renewable
          * `expiration_time` (`pulumi.Input[str]`) - Domain expiration timestamp
          * `last_renewed_time` (`pulumi.Input[str]`) - Timestamp when the domain was renewed last time
          * `managed_host_names` (`pulumi.Input[list]`) - All hostnames derived from the domain and assigned to Azure resources
            * `azure_resource_name` (`pulumi.Input[str]`) - Name of the Azure resource the hostname is assigned to. If it is assigned to a traffic manager then it will be the traffic manager name otherwise it will be the website name
            * `azure_resource_type` (`pulumi.Input[str]`) - Type of the Azure resource the hostname is assigned to
            * `custom_host_name_dns_record_type` (`pulumi.Input[str]`) - Type of the Dns record
            * `host_name_type` (`pulumi.Input[str]`) - Type of the hostname
            * `name` (`pulumi.Input[str]`) - Name of the hostname
            * `site_names` (`pulumi.Input[list]`) - List of sites the hostname is assigned to. This list will have more than one site only if the hostname is pointing to a Traffic Manager

          * `name_servers` (`pulumi.Input[list]`) - Name servers
          * `privacy` (`pulumi.Input[bool]`) - If true then domain privacy is enabled for this domain
          * `provisioning_state` (`pulumi.Input[str]`) - Domain provisioning state
          * `ready_for_dns_record_management` (`pulumi.Input[bool]`) - If true then Azure can assign this domain to Web Apps. This value will be true if domain registration status is active and it is hosted on name servers Azure has programmatic access to
          * `registration_status` (`pulumi.Input[str]`) - Domain registration status
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

            __props__['id'] = id
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = type
        super(Domain, __self__).__init__(
            'azurerm:domainregistration/v20150801:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Domain(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
