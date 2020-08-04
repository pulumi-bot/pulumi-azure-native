# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AssetFilter(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The Media Filter properties.
      * `first_quality` (`dict`) - The first quality.
        * `bitrate` (`float`) - The first quality bitrate.

      * `presentation_time_range` (`dict`) - The presentation time range.
        * `end_timestamp` (`float`) - The absolute end time boundary.
        * `force_end_timestamp` (`bool`) - The indicator of forcing existing of end time stamp.
        * `live_backoff_duration` (`float`) - The relative to end right edge.
        * `presentation_window_duration` (`float`) - The relative to end sliding window.
        * `start_timestamp` (`float`) - The absolute start time boundary.
        * `timescale` (`float`) - The time scale of time stamps.

      * `tracks` (`list`) - The tracks selection conditions.
        * `track_selections` (`list`) - The track selections.
          * `operation` (`str`) - The track property condition operation.
          * `property` (`str`) - The track property type.
          * `value` (`str`) - The track property value.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, asset_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        An Asset Filter.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[str] asset_name: The Asset name.
        :param pulumi.Input[str] name: The Asset Filter name
        :param pulumi.Input[dict] properties: The Media Filter properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.

        The **properties** object supports the following:

          * `first_quality` (`pulumi.Input[dict]`) - The first quality.
            * `bitrate` (`pulumi.Input[float]`) - The first quality bitrate.

          * `presentation_time_range` (`pulumi.Input[dict]`) - The presentation time range.
            * `end_timestamp` (`pulumi.Input[float]`) - The absolute end time boundary.
            * `force_end_timestamp` (`pulumi.Input[bool]`) - The indicator of forcing existing of end time stamp.
            * `live_backoff_duration` (`pulumi.Input[float]`) - The relative to end right edge.
            * `presentation_window_duration` (`pulumi.Input[float]`) - The relative to end sliding window.
            * `start_timestamp` (`pulumi.Input[float]`) - The absolute start time boundary.
            * `timescale` (`pulumi.Input[float]`) - The time scale of time stamps.

          * `tracks` (`pulumi.Input[list]`) - The tracks selection conditions.
            * `track_selections` (`pulumi.Input[list]`) - The track selections.
              * `operation` (`pulumi.Input[str]`) - The track property condition operation.
              * `property` (`pulumi.Input[str]`) - The track property type.
              * `value` (`pulumi.Input[str]`) - The track property value.
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
            if asset_name is None:
                raise TypeError("Missing required property 'asset_name'")
            __props__['asset_name'] = asset_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(AssetFilter, __self__).__init__(
            'azurerm:media/v20200501:AssetFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AssetFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AssetFilter(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
