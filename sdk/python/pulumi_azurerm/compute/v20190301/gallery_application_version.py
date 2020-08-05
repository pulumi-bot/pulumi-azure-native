# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GalleryApplicationVersion(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Describes the properties of a gallery Image Version.
      * `provisioning_state` (`str`) - The provisioning state, which only appears in the response.
      * `publishing_profile` (`dict`) - The publishing profile of a gallery Image Version.
        * `content_type` (`str`) - Optional. May be used to help process this file. The type of file contained in the source, e.g. zip, json, etc.
        * `enable_health_check` (`bool`) - Optional. Whether or not this application reports health.
        * `end_of_life_date` (`str`) - The end of life date of the gallery Image Version. This property can be used for decommissioning purposes. This property is updatable.
        * `exclude_from_latest` (`bool`) - If set to true, Virtual Machines deployed from the latest version of the Image Definition won't use this Image Version.
        * `published_date` (`str`) - The timestamp for when the gallery Image Version is published.
        * `replica_count` (`float`) - The number of replicas of the Image Version to be created per region. This property would take effect for a region when regionalReplicaCount is not specified. This property is updatable.
        * `source` (`dict`) - The source image from which the Image Version is going to be created.
          * `file_name` (`str`) - Required. The fileName of the artifact.
          * `media_link` (`str`) - Required. The mediaLink of the artifact, must be a readable storage blob.

        * `storage_account_type` (`str`) - Specifies the storage account type to be used to store the image. This property is not updatable.
        * `target_regions` (`list`) - The target regions where the Image Version is going to be replicated to. This property is updatable.
          * `name` (`str`) - The name of the region.
          * `regional_replica_count` (`float`) - The number of replicas of the Image Version to be created per region. This property is updatable.
          * `storage_account_type` (`str`) - Specifies the storage account type to be used to store the image. This property is not updatable.

      * `replication_status` (`dict`) - This is the replication status of the gallery Image Version.
        * `aggregated_state` (`str`) - This is the aggregated replication status based on all the regional replication status flags.
        * `summary` (`list`) - This is a summary of replication status for each region.
          * `details` (`str`) - The details of the replication status.
          * `progress` (`float`) - It indicates progress of the replication job.
          * `region` (`str`) - The region to which the gallery Image Version is being replicated to.
          * `state` (`str`) - This is the regional replication state.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, gallery_application_name=None, gallery_name=None, location=None, name=None, publishing_profile=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Specifies information about the gallery Application Version that you want to create or update.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gallery_application_name: The name of the gallery Application Definition in which the Application Version is to be created.
        :param pulumi.Input[str] gallery_name: The name of the Shared Application Gallery in which the Application Definition resides.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the gallery Application Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
        :param pulumi.Input[dict] publishing_profile: The publishing profile of a gallery Image Version.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags

        The **publishing_profile** object supports the following:

          * `content_type` (`pulumi.Input[str]`) - Optional. May be used to help process this file. The type of file contained in the source, e.g. zip, json, etc.
          * `enable_health_check` (`pulumi.Input[bool]`) - Optional. Whether or not this application reports health.
          * `end_of_life_date` (`pulumi.Input[str]`) - The end of life date of the gallery Image Version. This property can be used for decommissioning purposes. This property is updatable.
          * `exclude_from_latest` (`pulumi.Input[bool]`) - If set to true, Virtual Machines deployed from the latest version of the Image Definition won't use this Image Version.
          * `replica_count` (`pulumi.Input[float]`) - The number of replicas of the Image Version to be created per region. This property would take effect for a region when regionalReplicaCount is not specified. This property is updatable.
          * `source` (`pulumi.Input[dict]`) - The source image from which the Image Version is going to be created.
            * `file_name` (`pulumi.Input[str]`) - Required. The fileName of the artifact.
            * `media_link` (`pulumi.Input[str]`) - Required. The mediaLink of the artifact, must be a readable storage blob.

          * `storage_account_type` (`pulumi.Input[str]`) - Specifies the storage account type to be used to store the image. This property is not updatable.
          * `target_regions` (`pulumi.Input[list]`) - The target regions where the Image Version is going to be replicated to. This property is updatable.
            * `name` (`pulumi.Input[str]`) - The name of the region.
            * `regional_replica_count` (`pulumi.Input[float]`) - The number of replicas of the Image Version to be created per region. This property is updatable.
            * `storage_account_type` (`pulumi.Input[str]`) - Specifies the storage account type to be used to store the image. This property is not updatable.
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

            if gallery_application_name is None:
                raise TypeError("Missing required property 'gallery_application_name'")
            __props__['gallery_application_name'] = gallery_application_name
            if gallery_name is None:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if publishing_profile is None:
                raise TypeError("Missing required property 'publishing_profile'")
            __props__['publishing_profile'] = publishing_profile
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(GalleryApplicationVersion, __self__).__init__(
            'azurerm:compute/v20190301:GalleryApplicationVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing GalleryApplicationVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return GalleryApplicationVersion(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
