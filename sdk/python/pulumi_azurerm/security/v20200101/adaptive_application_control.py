# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AdaptiveApplicationControl(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Location where the resource is stored
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Represents a VM/server group and set of rules to be allowed running on a machine
      * `configuration_status` (`str`) - The configuration status of the VM/server group or machine or rule on the machine
      * `enforcement_mode` (`str`) - The application control policy enforcement/protection mode of the VM/server group
      * `issues` (`list`)
        * `issue` (`str`) - An alert that VMs/servers within a group can have
        * `number_of_vms` (`float`) - The number of machines in the VM/server group that have this alert

      * `path_recommendations` (`list`)
        * `action` (`str`) - The recommendation action of the VM/server or rule
        * `common` (`bool`) - Whether the path is commonly run on the machine
        * `configuration_status` (`str`) - The configuration status of the VM/server group or machine or rule on the machine
        * `file_type` (`str`) - The type of the file (for Linux files - Executable is used)
        * `path` (`str`) - The full path to whitelist
        * `publisher_info` (`dict`) - Represents the publisher information of a process/rule
          * `binary_name` (`str`) - The "OriginalName" field taken from the file's version resource
          * `product_name` (`str`) - The product name taken from the file's version resource
          * `publisher_name` (`str`) - The Subject field of the x.509 certificate used to sign the code, using the following fields -  O = Organization, L = Locality, S = State or Province, and C = Country
          * `version` (`str`) - The binary file version taken from the file's version resource

        * `type` (`str`) - The type of the rule to be allowed
        * `user_sids` (`list`)
        * `usernames` (`list`)
          * `recommendation_action` (`str`) - The recommendation action of the VM/server or rule
          * `username` (`str`) - Represents a user that is recommended to be allowed for a certain rule

      * `protection_mode` (`dict`) - The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.
        * `exe` (`str`) - The application control policy enforcement/protection mode of the VM/server group
        * `executable` (`str`) - The application control policy enforcement/protection mode of the VM/server group
        * `msi` (`str`) - The application control policy enforcement/protection mode of the VM/server group
        * `script` (`str`) - The application control policy enforcement/protection mode of the VM/server group

      * `recommendation_status` (`str`) - The recommendation status of the VM/server group or VM/server
      * `source_system` (`str`) - The source type of the VM/server group
      * `vm_recommendations` (`list`)
        * `configuration_status` (`str`) - The configuration status of the VM/server group or machine or rule on the machine
        * `enforcement_support` (`str`) - The VM/server supportability of Enforce feature
        * `recommendation_action` (`str`) - The recommendation action of the VM/server or rule
        * `resource_id` (`str`) - The full azure resource id of the machine
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, asc_location=None, enforcement_mode=None, name=None, path_recommendations=None, protection_mode=None, vm_recommendations=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a AdaptiveApplicationControl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asc_location: The location where ASC stores the data of the subscription. can be retrieved from Get locations
        :param pulumi.Input[str] enforcement_mode: The application control policy enforcement/protection mode of the VM/server group
        :param pulumi.Input[str] name: Name of an application control VM/server group
        :param pulumi.Input[dict] protection_mode: The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.

        The **path_recommendations** object supports the following:

          * `action` (`pulumi.Input[str]`) - The recommendation action of the VM/server or rule
          * `common` (`pulumi.Input[bool]`) - Whether the path is commonly run on the machine
          * `configuration_status` (`pulumi.Input[str]`) - The configuration status of the VM/server group or machine or rule on the machine
          * `file_type` (`pulumi.Input[str]`) - The type of the file (for Linux files - Executable is used)
          * `path` (`pulumi.Input[str]`) - The full path to whitelist
          * `publisher_info` (`pulumi.Input[dict]`) - Represents the publisher information of a process/rule
            * `binary_name` (`pulumi.Input[str]`) - The "OriginalName" field taken from the file's version resource
            * `product_name` (`pulumi.Input[str]`) - The product name taken from the file's version resource
            * `publisher_name` (`pulumi.Input[str]`) - The Subject field of the x.509 certificate used to sign the code, using the following fields -  O = Organization, L = Locality, S = State or Province, and C = Country
            * `version` (`pulumi.Input[str]`) - The binary file version taken from the file's version resource

          * `type` (`pulumi.Input[str]`) - The type of the rule to be allowed
          * `user_sids` (`pulumi.Input[list]`)
          * `usernames` (`pulumi.Input[list]`)
            * `recommendation_action` (`pulumi.Input[str]`) - The recommendation action of the VM/server or rule
            * `username` (`pulumi.Input[str]`) - Represents a user that is recommended to be allowed for a certain rule

        The **protection_mode** object supports the following:

          * `exe` (`pulumi.Input[str]`) - The application control policy enforcement/protection mode of the VM/server group
          * `executable` (`pulumi.Input[str]`) - The application control policy enforcement/protection mode of the VM/server group
          * `msi` (`pulumi.Input[str]`) - The application control policy enforcement/protection mode of the VM/server group
          * `script` (`pulumi.Input[str]`) - The application control policy enforcement/protection mode of the VM/server group

        The **vm_recommendations** object supports the following:

          * `configuration_status` (`pulumi.Input[str]`) - The configuration status of the VM/server group or machine or rule on the machine
          * `enforcement_support` (`pulumi.Input[str]`) - The VM/server supportability of Enforce feature
          * `recommendation_action` (`pulumi.Input[str]`) - The recommendation action of the VM/server or rule
          * `resource_id` (`pulumi.Input[str]`) - The full azure resource id of the machine
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

            if asc_location is None:
                raise TypeError("Missing required property 'asc_location'")
            __props__['asc_location'] = asc_location
            __props__['enforcement_mode'] = enforcement_mode
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['path_recommendations'] = path_recommendations
            __props__['protection_mode'] = protection_mode
            __props__['vm_recommendations'] = vm_recommendations
            __props__['location'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(AdaptiveApplicationControl, __self__).__init__(
            'azurerm:security/v20200101:AdaptiveApplicationControl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AdaptiveApplicationControl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AdaptiveApplicationControl(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop