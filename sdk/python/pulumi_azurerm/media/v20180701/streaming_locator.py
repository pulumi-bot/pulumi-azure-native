# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class StreamingLocator(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the Streaming Locator.
      * `alternative_media_id` (`str`) - Alternative Media ID of this Streaming Locator
      * `asset_name` (`str`) - Asset Name
      * `content_keys` (`list`) - The ContentKeys used by this Streaming Locator.
        * `id` (`str`) - ID of Content Key
        * `label_reference_in_streaming_policy` (`str`) - Label of Content Key as specified in the Streaming Policy
        * `policy_name` (`str`) - ContentKeyPolicy used by Content Key
        * `tracks` (`list`) - Tracks which use this Content Key
          * `track_selections` (`list`) - TrackSelections is a track property condition list which can specify track(s)
            * `operation` (`str`) - Track property condition operation
            * `property` (`str`) - Track property type
            * `value` (`str`) - Track property value

        * `type` (`str`) - Encryption type of Content Key
        * `value` (`str`) - Value of Content Key

      * `created` (`str`) - The creation time of the Streaming Locator.
      * `default_content_key_policy_name` (`str`) - Name of the default ContentKeyPolicy used by this Streaming Locator.
      * `end_time` (`str`) - The end time of the Streaming Locator.
      * `filters` (`list`) - A list of asset or account filters which apply to this streaming locator
      * `start_time` (`str`) - The start time of the Streaming Locator.
      * `streaming_locator_id` (`str`) - The StreamingLocatorId of the Streaming Locator.
      * `streaming_policy_name` (`str`) - Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        A Streaming Locator resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[str] name: The Streaming Locator name.
        :param pulumi.Input[dict] properties: Properties of the Streaming Locator.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.

        The **properties** object supports the following:

          * `alternative_media_id` (`pulumi.Input[str]`) - Alternative Media ID of this Streaming Locator
          * `asset_name` (`pulumi.Input[str]`) - Asset Name
          * `content_keys` (`pulumi.Input[list]`) - The ContentKeys used by this Streaming Locator.
            * `id` (`pulumi.Input[str]`) - ID of Content Key
            * `label_reference_in_streaming_policy` (`pulumi.Input[str]`) - Label of Content Key as specified in the Streaming Policy
            * `value` (`pulumi.Input[str]`) - Value of Content Key

          * `default_content_key_policy_name` (`pulumi.Input[str]`) - Name of the default ContentKeyPolicy used by this Streaming Locator.
          * `end_time` (`pulumi.Input[str]`) - The end time of the Streaming Locator.
          * `filters` (`pulumi.Input[list]`) - A list of asset or account filters which apply to this streaming locator
          * `start_time` (`pulumi.Input[str]`) - The start time of the Streaming Locator.
          * `streaming_locator_id` (`pulumi.Input[str]`) - The StreamingLocatorId of the Streaming Locator.
          * `streaming_policy_name` (`pulumi.Input[str]`) - Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(StreamingLocator, __self__).__init__(
            'azurerm:media/v20180701:StreamingLocator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing StreamingLocator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return StreamingLocator(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
