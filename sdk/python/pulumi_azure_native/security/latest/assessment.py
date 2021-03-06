# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['Assessment']

warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:security:Assessment'.""", DeprecationWarning)


class Assessment(pulumi.CustomResource):
    warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:security:Assessment'.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 assessment_name: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['SecurityAssessmentMetadataPropertiesArgs']]] = None,
                 partners_data: Optional[pulumi.Input[pulumi.InputType['SecurityAssessmentPartnerDataArgs']]] = None,
                 resource_details: Optional[pulumi.Input[Union[pulumi.InputType['AzureResourceDetailsArgs'], pulumi.InputType['OnPremiseResourceDetailsArgs'], pulumi.InputType['OnPremiseSqlResourceDetailsArgs']]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['AssessmentStatusArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Security assessment on a resource
        Latest API Version: 2020-01-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_data: Additional data regarding the assessment
        :param pulumi.Input[str] assessment_name: The Assessment Key - Unique key for the assessment type
        :param pulumi.Input[pulumi.InputType['SecurityAssessmentMetadataPropertiesArgs']] metadata: Describes properties of an assessment metadata.
        :param pulumi.Input[pulumi.InputType['SecurityAssessmentPartnerDataArgs']] partners_data: Data regarding 3rd party partner integration
        :param pulumi.Input[Union[pulumi.InputType['AzureResourceDetailsArgs'], pulumi.InputType['OnPremiseResourceDetailsArgs'], pulumi.InputType['OnPremiseSqlResourceDetailsArgs']]] resource_details: Details of the resource that was assessed
        :param pulumi.Input[str] resource_id: The identifier of the resource.
        :param pulumi.Input[pulumi.InputType['AssessmentStatusArgs']] status: The result of the assessment
        """
        pulumi.log.warn("""Assessment is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:security:Assessment'.""")
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

            __props__['additional_data'] = additional_data
            __props__['assessment_name'] = assessment_name
            __props__['metadata'] = metadata
            __props__['partners_data'] = partners_data
            if resource_details is None and not opts.urn:
                raise TypeError("Missing required property 'resource_details'")
            __props__['resource_details'] = resource_details
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__['status'] = status
            __props__['display_name'] = None
            __props__['links'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:security/latest:Assessment"), pulumi.Alias(type_="azure-native:security:Assessment"), pulumi.Alias(type_="azure-nextgen:security:Assessment"), pulumi.Alias(type_="azure-native:security/v20190101preview:Assessment"), pulumi.Alias(type_="azure-nextgen:security/v20190101preview:Assessment"), pulumi.Alias(type_="azure-native:security/v20200101:Assessment"), pulumi.Alias(type_="azure-nextgen:security/v20200101:Assessment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Assessment, __self__).__init__(
            'azure-native:security/latest:Assessment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Assessment':
        """
        Get an existing Assessment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_data"] = None
        __props__["display_name"] = None
        __props__["links"] = None
        __props__["metadata"] = None
        __props__["name"] = None
        __props__["partners_data"] = None
        __props__["resource_details"] = None
        __props__["status"] = None
        __props__["type"] = None
        return Assessment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalData")
    def additional_data(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Additional data regarding the assessment
        """
        return pulumi.get(self, "additional_data")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        User friendly display name of the assessment
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output['outputs.AssessmentLinksResponse']:
        """
        Links relevant to the assessment
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional['outputs.SecurityAssessmentMetadataPropertiesResponse']]:
        """
        Describes properties of an assessment metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partnersData")
    def partners_data(self) -> pulumi.Output[Optional['outputs.SecurityAssessmentPartnerDataResponse']]:
        """
        Data regarding 3rd party partner integration
        """
        return pulumi.get(self, "partners_data")

    @property
    @pulumi.getter(name="resourceDetails")
    def resource_details(self) -> pulumi.Output[Any]:
        """
        Details of the resource that was assessed
        """
        return pulumi.get(self, "resource_details")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.AssessmentStatusResponse']:
        """
        The result of the assessment
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

