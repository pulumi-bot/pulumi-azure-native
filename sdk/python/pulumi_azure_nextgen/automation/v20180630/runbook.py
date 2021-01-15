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

__all__ = ['Runbook']


class Runbook(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 draft: Optional[pulumi.Input[pulumi.InputType['RunbookDraftArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_activity_trace: Optional[pulumi.Input[int]] = None,
                 log_progress: Optional[pulumi.Input[bool]] = None,
                 log_verbose: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 publish_content_link: Optional[pulumi.Input[pulumi.InputType['ContentLinkArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 runbook_name: Optional[pulumi.Input[str]] = None,
                 runbook_type: Optional[pulumi.Input[Union[str, 'RunbookTypeEnum']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Definition of the runbook type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account.
        :param pulumi.Input[str] description: Gets or sets the description of the runbook.
        :param pulumi.Input[pulumi.InputType['RunbookDraftArgs']] draft: Gets or sets the draft runbook properties.
        :param pulumi.Input[str] location: Gets or sets the location of the resource.
        :param pulumi.Input[int] log_activity_trace: Gets or sets the activity-level tracing options of the runbook.
        :param pulumi.Input[bool] log_progress: Gets or sets progress log option.
        :param pulumi.Input[bool] log_verbose: Gets or sets verbose log option.
        :param pulumi.Input[str] name: Gets or sets the name of the resource.
        :param pulumi.Input[pulumi.InputType['ContentLinkArgs']] publish_content_link: Gets or sets the published runbook content link.
        :param pulumi.Input[str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[str] runbook_name: The runbook name.
        :param pulumi.Input[Union[str, 'RunbookTypeEnum']] runbook_type: Gets or sets the type of the runbook.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Gets or sets the tags attached to the resource.
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

            if automation_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            __props__['description'] = description
            __props__['draft'] = draft
            __props__['location'] = location
            __props__['log_activity_trace'] = log_activity_trace
            __props__['log_progress'] = log_progress
            __props__['log_verbose'] = log_verbose
            __props__['name'] = name
            __props__['publish_content_link'] = publish_content_link
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if runbook_name is None and not opts.urn:
                raise TypeError("Missing required property 'runbook_name'")
            __props__['runbook_name'] = runbook_name
            if runbook_type is None and not opts.urn:
                raise TypeError("Missing required property 'runbook_type'")
            __props__['runbook_type'] = runbook_type
            __props__['tags'] = tags
            __props__['creation_time'] = None
            __props__['etag'] = None
            __props__['job_count'] = None
            __props__['last_modified_by'] = None
            __props__['last_modified_time'] = None
            __props__['output_types'] = None
            __props__['parameters'] = None
            __props__['provisioning_state'] = None
            __props__['state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:automation/latest:Runbook"), pulumi.Alias(type_="azure-nextgen:automation/v20151031:Runbook"), pulumi.Alias(type_="azure-nextgen:automation/v20190601:Runbook")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Runbook, __self__).__init__(
            'azure-nextgen:automation/v20180630:Runbook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Runbook':
        """
        Get an existing Runbook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Runbook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the creation time.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def draft(self) -> pulumi.Output[Optional['outputs.RunbookDraftResponse']]:
        """
        Gets or sets the draft runbook properties.
        """
        return pulumi.get(self, "draft")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the etag of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="jobCount")
    def job_count(self) -> pulumi.Output[Optional[int]]:
        """
        Gets or sets the job count of the runbook.
        """
        return pulumi.get(self, "job_count")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the last modified by.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the last modified time.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The Azure Region where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logActivityTrace")
    def log_activity_trace(self) -> pulumi.Output[Optional[int]]:
        """
        Gets or sets the option to log activity trace of the runbook.
        """
        return pulumi.get(self, "log_activity_trace")

    @property
    @pulumi.getter(name="logProgress")
    def log_progress(self) -> pulumi.Output[Optional[bool]]:
        """
        Gets or sets progress log option.
        """
        return pulumi.get(self, "log_progress")

    @property
    @pulumi.getter(name="logVerbose")
    def log_verbose(self) -> pulumi.Output[Optional[bool]]:
        """
        Gets or sets verbose log option.
        """
        return pulumi.get(self, "log_verbose")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputTypes")
    def output_types(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Gets or sets the runbook output types.
        """
        return pulumi.get(self, "output_types")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.RunbookParameterResponse']]]:
        """
        Gets or sets the runbook parameters.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the provisioning state of the runbook.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publishContentLink")
    def publish_content_link(self) -> pulumi.Output[Optional['outputs.ContentLinkResponse']]:
        """
        Gets or sets the published runbook content link.
        """
        return pulumi.get(self, "publish_content_link")

    @property
    @pulumi.getter(name="runbookType")
    def runbook_type(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the type of the runbook.
        """
        return pulumi.get(self, "runbook_type")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the state of the runbook.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

