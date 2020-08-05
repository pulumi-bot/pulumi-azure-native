# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PacketCapture(pulumi.CustomResource):
    etag: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    Name of the packet capture session.
    """
    properties: pulumi.Output[dict]
    """
    Describes the properties of a packet capture session.
      * `bytes_to_capture_per_packet` (`float`) - Number of bytes captured per packet, the remaining bytes are truncated.
      * `filters` (`list`)
        * `local_ip_address` (`str`) - Local IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5"? for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Default = null.
        * `local_port` (`str`) - Local port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Default = null.
        * `protocol` (`str`) - Protocol to be filtered on.
        * `remote_ip_address` (`str`) - Local IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Default = null.
        * `remote_port` (`str`) - Remote port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Default = null.

      * `provisioning_state` (`str`) - The provisioning state of the packet capture session.
      * `storage_location` (`dict`) - Describes the storage location for a packet capture session.
        * `file_path` (`str`) - A valid local path on the targeting VM. Must include the name of the capture file (*.cap). For linux virtual machine it must start with /var/captures. Required if no storage ID is provided, otherwise optional.
        * `storage_id` (`str`) - The ID of the storage account to save the packet capture session. Required if no local file path is provided.
        * `storage_path` (`str`) - The URI of the storage path to save the packet capture. Must be a well-formed URI describing the location to save the packet capture.

      * `target` (`str`) - The ID of the targeted resource, only VM is currently supported.
      * `time_limit_in_seconds` (`float`) - Maximum duration of the capture session in seconds.
      * `total_bytes_per_session` (`float`) - Maximum size of the capture output.
    """
    def __init__(__self__, resource_name, opts=None, bytes_to_capture_per_packet=None, filters=None, name=None, network_watcher_name=None, resource_group_name=None, storage_location=None, target=None, time_limit_in_seconds=None, total_bytes_per_session=None, __props__=None, __name__=None, __opts__=None):
        """
        Information about packet capture session.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bytes_to_capture_per_packet: Number of bytes captured per packet, the remaining bytes are truncated.
        :param pulumi.Input[str] name: The name of the packet capture session.
        :param pulumi.Input[str] network_watcher_name: The name of the network watcher.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] storage_location: Describes the storage location for a packet capture session.
        :param pulumi.Input[str] target: The ID of the targeted resource, only VM is currently supported.
        :param pulumi.Input[float] time_limit_in_seconds: Maximum duration of the capture session in seconds.
        :param pulumi.Input[float] total_bytes_per_session: Maximum size of the capture output.

        The **filters** object supports the following:

          * `local_ip_address` (`pulumi.Input[str]`) - Local IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5"? for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Default = null.
          * `local_port` (`pulumi.Input[str]`) - Local port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Default = null.
          * `protocol` (`pulumi.Input[str]`) - Protocol to be filtered on.
          * `remote_ip_address` (`pulumi.Input[str]`) - Local IP Address to be filtered on. Notation: "127.0.0.1" for single address entry. "127.0.0.1-127.0.0.255" for range. "127.0.0.1;127.0.0.5;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Default = null.
          * `remote_port` (`pulumi.Input[str]`) - Remote port to be filtered on. Notation: "80" for single port entry."80-85" for range. "80;443;" for multiple entries. Multiple ranges not currently supported. Mixing ranges with multiple entries not currently supported. Default = null.

        The **storage_location** object supports the following:

          * `file_path` (`pulumi.Input[str]`) - A valid local path on the targeting VM. Must include the name of the capture file (*.cap). For linux virtual machine it must start with /var/captures. Required if no storage ID is provided, otherwise optional.
          * `storage_id` (`pulumi.Input[str]`) - The ID of the storage account to save the packet capture session. Required if no local file path is provided.
          * `storage_path` (`pulumi.Input[str]`) - The URI of the storage path to save the packet capture. Must be a well-formed URI describing the location to save the packet capture.
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

            __props__['bytes_to_capture_per_packet'] = bytes_to_capture_per_packet
            __props__['filters'] = filters
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if network_watcher_name is None:
                raise TypeError("Missing required property 'network_watcher_name'")
            __props__['network_watcher_name'] = network_watcher_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if storage_location is None:
                raise TypeError("Missing required property 'storage_location'")
            __props__['storage_location'] = storage_location
            if target is None:
                raise TypeError("Missing required property 'target'")
            __props__['target'] = target
            __props__['time_limit_in_seconds'] = time_limit_in_seconds
            __props__['total_bytes_per_session'] = total_bytes_per_session
            __props__['etag'] = None
            __props__['properties'] = None
        super(PacketCapture, __self__).__init__(
            'azurerm:network/v20180401:PacketCapture',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PacketCapture resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PacketCapture(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
