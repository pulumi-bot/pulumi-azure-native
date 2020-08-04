# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WebAppSiteExtension(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    SiteExtensionInfo resource specific properties
      * `authors` (`list`) - List of authors.
      * `comment` (`str`) - Site Extension comment.
      * `description` (`str`) - Detailed description.
      * `download_count` (`float`) - Count of downloads.
      * `extension_id` (`str`) - Site extension ID.
      * `extension_type` (`str`) - Site extension type.
      * `extension_url` (`str`) - Extension URL.
      * `feed_url` (`str`) - Feed URL.
      * `icon_url` (`str`) - Icon URL.
      * `installed_date_time` (`str`) - Installed timestamp.
      * `installer_command_line_params` (`str`) - Installer command line parameters.
      * `license_url` (`str`) - License URL.
      * `local_is_latest_version` (`bool`) - <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
      * `local_path` (`str`) - Local path.
      * `project_url` (`str`) - Project URL.
      * `provisioning_state` (`str`) - Provisioning state.
      * `published_date_time` (`str`) - Published timestamp.
      * `summary` (`str`) - Summary description.
      * `title` (`str`)
      * `version` (`str`) - Version information.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, name=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Site Extension Information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Site extension name.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['kind'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(WebAppSiteExtension, __self__).__init__(
            'azurerm:web/v20180201:WebAppSiteExtension',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WebAppSiteExtension resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppSiteExtension(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
