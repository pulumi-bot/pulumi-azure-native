# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Task(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource. This cannot be changed after the resource is created.
    """
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties of a task.
      * `agent_configuration` (`dict`) - The machine configuration of the run agent.
        * `cpu` (`float`) - The CPU configuration in terms of number of cores required for the run.

      * `creation_date` (`str`) - The creation date of task.
      * `credentials` (`dict`) - The properties that describes a set of credentials that will be used when this run is invoked.
        * `custom_registries` (`dict`) - Describes the credential parameters for accessing other custom registries. The key
          for the dictionary item will be the registry login server (myregistry.azurecr.io) and
          the value of the item will be the registry credentials for accessing the registry.
        * `source_registry` (`dict`) - Describes the credential parameters for accessing the source registry.
          * `login_mode` (`str`) - The authentication mode which determines the source registry login scope. The credentials for the source registry
            will be generated using the given scope. These credentials will be used to login to
            the source registry during the run.

      * `platform` (`dict`) - The platform properties against which the run has to happen.
        * `architecture` (`str`) - The OS architecture.
        * `os` (`str`) - The operating system type required for the run.
        * `variant` (`str`) - Variant of the CPU.

      * `provisioning_state` (`str`) - The provisioning state of the task.
      * `status` (`str`) - The current status of task.
      * `step` (`dict`) - The properties of a task step.
        * `base_image_dependencies` (`list`) - List of base image dependencies for a step.
          * `digest` (`str`) - The sha256-based digest of the image manifest.
          * `registry` (`str`) - The registry login server.
          * `repository` (`str`) - The repository name.
          * `tag` (`str`) - The tag name.
          * `type` (`str`) - The type of the base image dependency.

        * `context_access_token` (`str`) - The token (git PAT or SAS token of storage account blob) associated with the context for a step.
        * `context_path` (`str`) - The URL(absolute or relative) of the source context for the task step.
        * `type` (`str`) - The type of the step.

      * `timeout` (`float`) - Run timeout in seconds.
      * `trigger` (`dict`) - The properties that describe all triggers for the task.
        * `base_image_trigger` (`dict`) - The trigger based on base image dependencies.
          * `base_image_trigger_type` (`str`) - The type of the auto trigger for base image dependency updates.
          * `name` (`str`) - The name of the trigger.
          * `status` (`str`) - The current status of trigger.

        * `source_triggers` (`list`) - The collection of triggers based on source code repository.
          * `name` (`str`) - The name of the trigger.
          * `source_repository` (`dict`) - The properties that describes the source(code) for the task.
            * `branch` (`str`) - The branch name of the source code.
            * `repository_url` (`str`) - The full URL to the source code repository
            * `source_control_auth_properties` (`dict`) - The authorization properties for accessing the source code repository and to set up
              webhooks for notifications.
              * `expires_in` (`float`) - Time in seconds that the token remains valid
              * `refresh_token` (`str`) - The refresh token used to refresh the access token.
              * `scope` (`str`) - The scope of the access token.
              * `token` (`str`) - The access token used to access the source control provider.
              * `token_type` (`str`) - The type of Auth token.

            * `source_control_type` (`str`) - The type of source control service.

          * `source_trigger_events` (`list`) - The source event corresponding to the trigger.
          * `status` (`str`) - The current status of trigger.
    """
    tags: pulumi.Output[dict]
    """
    The tags of the resource.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, registry_name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The task that has the ARM resource and task properties.
        The task will have all information to schedule a run against it.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the resource. This cannot be changed after the resource is created.
        :param pulumi.Input[str] name: The name of the container registry task.
        :param pulumi.Input[dict] properties: The properties of a task.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to which the container registry belongs.
        :param pulumi.Input[dict] tags: The tags of the resource.

        The **properties** object supports the following:

          * `agent_configuration` (`pulumi.Input[dict]`) - The machine configuration of the run agent.
            * `cpu` (`pulumi.Input[float]`) - The CPU configuration in terms of number of cores required for the run.

          * `credentials` (`pulumi.Input[dict]`) - The properties that describes a set of credentials that will be used when this run is invoked.
            * `custom_registries` (`pulumi.Input[dict]`) - Describes the credential parameters for accessing other custom registries. The key
              for the dictionary item will be the registry login server (myregistry.azurecr.io) and
              the value of the item will be the registry credentials for accessing the registry.
            * `source_registry` (`pulumi.Input[dict]`) - Describes the credential parameters for accessing the source registry.
              * `login_mode` (`pulumi.Input[str]`) - The authentication mode which determines the source registry login scope. The credentials for the source registry
                will be generated using the given scope. These credentials will be used to login to
                the source registry during the run.

          * `platform` (`pulumi.Input[dict]`) - The platform properties against which the run has to happen.
            * `architecture` (`pulumi.Input[str]`) - The OS architecture.
            * `os` (`pulumi.Input[str]`) - The operating system type required for the run.
            * `variant` (`pulumi.Input[str]`) - Variant of the CPU.

          * `status` (`pulumi.Input[str]`) - The current status of task.
          * `step` (`pulumi.Input[dict]`) - The properties of a task step.
            * `context_access_token` (`pulumi.Input[str]`) - The token (git PAT or SAS token of storage account blob) associated with the context for a step.
            * `context_path` (`pulumi.Input[str]`) - The URL(absolute or relative) of the source context for the task step.

          * `timeout` (`pulumi.Input[float]`) - Run timeout in seconds.
          * `trigger` (`pulumi.Input[dict]`) - The properties that describe all triggers for the task.
            * `base_image_trigger` (`pulumi.Input[dict]`) - The trigger based on base image dependencies.
              * `base_image_trigger_type` (`pulumi.Input[str]`) - The type of the auto trigger for base image dependency updates.
              * `name` (`pulumi.Input[str]`) - The name of the trigger.
              * `status` (`pulumi.Input[str]`) - The current status of trigger.

            * `source_triggers` (`pulumi.Input[list]`) - The collection of triggers based on source code repository.
              * `name` (`pulumi.Input[str]`) - The name of the trigger.
              * `source_repository` (`pulumi.Input[dict]`) - The properties that describes the source(code) for the task.
                * `branch` (`pulumi.Input[str]`) - The branch name of the source code.
                * `repository_url` (`pulumi.Input[str]`) - The full URL to the source code repository
                * `source_control_auth_properties` (`pulumi.Input[dict]`) - The authorization properties for accessing the source code repository and to set up
                  webhooks for notifications.
                  * `expires_in` (`pulumi.Input[float]`) - Time in seconds that the token remains valid
                  * `refresh_token` (`pulumi.Input[str]`) - The refresh token used to refresh the access token.
                  * `scope` (`pulumi.Input[str]`) - The scope of the access token.
                  * `token` (`pulumi.Input[str]`) - The access token used to access the source control provider.
                  * `token_type` (`pulumi.Input[str]`) - The type of Auth token.

                * `source_control_type` (`pulumi.Input[str]`) - The type of source control service.

              * `source_trigger_events` (`pulumi.Input[list]`) - The source event corresponding to the trigger.
              * `status` (`pulumi.Input[str]`) - The current status of trigger.
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if registry_name is None:
                raise TypeError("Missing required property 'registry_name'")
            __props__['registry_name'] = registry_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(Task, __self__).__init__(
            'azurerm:containerregistry/v20180901:Task',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Task resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Task(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop