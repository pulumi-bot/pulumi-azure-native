# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Task']


class Task(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_configuration: Optional[pulumi.Input[pulumi.InputType['AgentPropertiesArgs']]] = None,
                 agent_pool_name: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[pulumi.InputType['CredentialsArgs']]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['IdentityPropertiesArgs']]] = None,
                 is_system_task: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_template: Optional[pulumi.Input[str]] = None,
                 platform: Optional[pulumi.Input[pulumi.InputType['PlatformPropertiesArgs']]] = None,
                 registry_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 step: Optional[pulumi.Input[Union[pulumi.InputType['DockerBuildStepArgs'], pulumi.InputType['EncodedTaskStepArgs'], pulumi.InputType['FileTaskStepArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 task_name: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[float]] = None,
                 trigger: Optional[pulumi.Input[pulumi.InputType['TriggerPropertiesArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The task that has the ARM resource and task properties.
        The task will have all information to schedule a run against it.

        ## Example Usage
        ### Tasks_Create

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        task = azurerm.containerregistry.latest.Task("task",
            agent_configuration={
                "cpu": 2,
            },
            identity={
                "type": "SystemAssigned",
            },
            is_system_task=False,
            location="eastus",
            log_template="acr/tasks:{{.Run.OS}}",
            platform={
                "architecture": "amd64",
                "os": "Linux",
            },
            registry_name="myRegistry",
            resource_group_name="myResourceGroup",
            status="Enabled",
            step={
                "arguments": [
                    {
                        "isSecret": False,
                        "name": "mytestargument",
                        "value": "mytestvalue",
                    },
                    {
                        "isSecret": True,
                        "name": "mysecrettestargument",
                        "value": "mysecrettestvalue",
                    },
                ],
                "contextPath": "src",
                "dockerFilePath": "src/DockerFile",
                "imageNames": ["azurerest:testtag"],
                "isPushEnabled": True,
                "noCache": False,
                "type": "Docker",
            },
            tags={
                "testkey": "value",
            },
            task_name="mytTask",
            trigger={
                "baseImageTrigger": {
                    "baseImageTriggerType": "Runtime",
                    "name": "myBaseImageTrigger",
                    "updateTriggerEndpoint": "https://user:pass@mycicd.webhook.com?token=foo",
                    "updateTriggerPayloadType": "Token",
                },
                "sourceTriggers": [{
                    "name": "mySourceTrigger",
                    "sourceRepository": {
                        "branch": "master",
                        "repositoryUrl": "https://github.com/Azure/azure-rest-api-specs",
                        "sourceControlAuthProperties": {
                            "token": "xxxxx",
                            "tokenType": "PAT",
                        },
                        "sourceControlType": "Github",
                    },
                    "sourceTriggerEvents": ["commit"],
                }],
                "timerTriggers": [{
                    "name": "myTimerTrigger",
                    "schedule": "30 9 * * 1-5",
                }],
            })

        ```
        ### Tasks_Create_QuickTask

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        task = azurerm.containerregistry.latest.Task("task",
            is_system_task=True,
            location="eastus",
            log_template="acr/tasks:{{.Run.OS}}",
            registry_name="myRegistry",
            resource_group_name="myResourceGroup",
            status="Enabled",
            tags={
                "testkey": "value",
            },
            task_name="quicktask")

        ```
        ### Tasks_Create_WithSystemAndUserIdentities

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        task = azurerm.containerregistry.latest.Task("task",
            agent_configuration={
                "cpu": 2,
            },
            identity={
                "type": "SystemAssigned, UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": {},
                },
            },
            is_system_task=False,
            location="eastus",
            platform={
                "architecture": "amd64",
                "os": "Linux",
            },
            registry_name="myRegistry",
            resource_group_name="myResourceGroup",
            status="Enabled",
            step={
                "arguments": [
                    {
                        "isSecret": False,
                        "name": "mytestargument",
                        "value": "mytestvalue",
                    },
                    {
                        "isSecret": True,
                        "name": "mysecrettestargument",
                        "value": "mysecrettestvalue",
                    },
                ],
                "contextPath": "src",
                "dockerFilePath": "src/DockerFile",
                "imageNames": ["azurerest:testtag"],
                "isPushEnabled": True,
                "noCache": False,
                "type": "Docker",
            },
            tags={
                "testkey": "value",
            },
            task_name="mytTask",
            trigger={
                "baseImageTrigger": {
                    "baseImageTriggerType": "Runtime",
                    "name": "myBaseImageTrigger",
                    "updateTriggerEndpoint": "https://user:pass@mycicd.webhook.com?token=foo",
                    "updateTriggerPayloadType": "Default",
                },
                "sourceTriggers": [{
                    "name": "mySourceTrigger",
                    "sourceRepository": {
                        "branch": "master",
                        "repositoryUrl": "https://github.com/Azure/azure-rest-api-specs",
                        "sourceControlAuthProperties": {
                            "token": "xxxxx",
                            "tokenType": "PAT",
                        },
                        "sourceControlType": "Github",
                    },
                    "sourceTriggerEvents": ["commit"],
                }],
                "timerTriggers": [{
                    "name": "myTimerTrigger",
                    "schedule": "30 9 * * 1-5",
                }],
            })

        ```
        ### Tasks_Create_WithUserIdentities

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        task = azurerm.containerregistry.latest.Task("task",
            agent_configuration={
                "cpu": 2,
            },
            identity={
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
                    "/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": {},
                },
            },
            is_system_task=False,
            location="eastus",
            platform={
                "architecture": "amd64",
                "os": "Linux",
            },
            registry_name="myRegistry",
            resource_group_name="myResourceGroup",
            status="Enabled",
            step={
                "arguments": [
                    {
                        "isSecret": False,
                        "name": "mytestargument",
                        "value": "mytestvalue",
                    },
                    {
                        "isSecret": True,
                        "name": "mysecrettestargument",
                        "value": "mysecrettestvalue",
                    },
                ],
                "contextPath": "src",
                "dockerFilePath": "src/DockerFile",
                "imageNames": ["azurerest:testtag"],
                "isPushEnabled": True,
                "noCache": False,
                "type": "Docker",
            },
            tags={
                "testkey": "value",
            },
            task_name="mytTask",
            trigger={
                "baseImageTrigger": {
                    "baseImageTriggerType": "Runtime",
                    "name": "myBaseImageTrigger",
                    "updateTriggerEndpoint": "https://user:pass@mycicd.webhook.com?token=foo",
                    "updateTriggerPayloadType": "Default",
                },
                "sourceTriggers": [{
                    "name": "mySourceTrigger",
                    "sourceRepository": {
                        "branch": "master",
                        "repositoryUrl": "https://github.com/Azure/azure-rest-api-specs",
                        "sourceControlAuthProperties": {
                            "token": "xxxxx",
                            "tokenType": "PAT",
                        },
                        "sourceControlType": "Github",
                    },
                    "sourceTriggerEvents": ["commit"],
                }],
                "timerTriggers": [{
                    "name": "myTimerTrigger",
                    "schedule": "30 9 * * 1-5",
                }],
            })

        ```
        ### Tasks_Create_WithUserIdentities_WithSystemIdentity

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        task = azurerm.containerregistry.latest.Task("task",
            agent_configuration={
                "cpu": 2,
            },
            identity={
                "type": "SystemAssigned",
            },
            is_system_task=False,
            location="eastus",
            platform={
                "architecture": "amd64",
                "os": "Linux",
            },
            registry_name="myRegistry",
            resource_group_name="myResourceGroup",
            status="Enabled",
            step={
                "arguments": [
                    {
                        "isSecret": False,
                        "name": "mytestargument",
                        "value": "mytestvalue",
                    },
                    {
                        "isSecret": True,
                        "name": "mysecrettestargument",
                        "value": "mysecrettestvalue",
                    },
                ],
                "contextPath": "src",
                "dockerFilePath": "src/DockerFile",
                "imageNames": ["azurerest:testtag"],
                "isPushEnabled": True,
                "noCache": False,
                "type": "Docker",
            },
            tags={
                "testkey": "value",
            },
            task_name="mytTask",
            trigger={
                "baseImageTrigger": {
                    "baseImageTriggerType": "Runtime",
                    "name": "myBaseImageTrigger",
                },
                "sourceTriggers": [{
                    "name": "mySourceTrigger",
                    "sourceRepository": {
                        "branch": "master",
                        "repositoryUrl": "https://github.com/Azure/azure-rest-api-specs",
                        "sourceControlAuthProperties": {
                            "token": "xxxxx",
                            "tokenType": "PAT",
                        },
                        "sourceControlType": "Github",
                    },
                    "sourceTriggerEvents": ["commit"],
                }],
                "timerTriggers": [{
                    "name": "myTimerTrigger",
                    "schedule": "30 9 * * 1-5",
                }],
            })

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AgentPropertiesArgs']] agent_configuration: The machine configuration of the run agent.
        :param pulumi.Input[str] agent_pool_name: The dedicated agent pool for the task.
        :param pulumi.Input[pulumi.InputType['CredentialsArgs']] credentials: The properties that describes a set of credentials that will be used when this run is invoked.
        :param pulumi.Input[pulumi.InputType['IdentityPropertiesArgs']] identity: Identity for the resource.
        :param pulumi.Input[bool] is_system_task: The value of this property indicates whether the task resource is system task or not.
        :param pulumi.Input[str] location: The location of the resource. This cannot be changed after the resource is created.
        :param pulumi.Input[str] log_template: The template that describes the repository and tag information for run log artifact.
        :param pulumi.Input[pulumi.InputType['PlatformPropertiesArgs']] platform: The platform properties against which the run has to happen.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to which the container registry belongs.
        :param pulumi.Input[str] status: The current status of task.
        :param pulumi.Input[Union[pulumi.InputType['DockerBuildStepArgs'], pulumi.InputType['EncodedTaskStepArgs'], pulumi.InputType['FileTaskStepArgs']]] step: The properties of a task step.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        :param pulumi.Input[str] task_name: The name of the container registry task.
        :param pulumi.Input[float] timeout: Run timeout in seconds.
        :param pulumi.Input[pulumi.InputType['TriggerPropertiesArgs']] trigger: The properties that describe all triggers for the task.
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

            __props__['agent_configuration'] = agent_configuration
            __props__['agent_pool_name'] = agent_pool_name
            __props__['credentials'] = credentials
            __props__['identity'] = identity
            __props__['is_system_task'] = is_system_task
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['log_template'] = log_template
            __props__['platform'] = platform
            if registry_name is None:
                raise TypeError("Missing required property 'registry_name'")
            __props__['registry_name'] = registry_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['status'] = status
            __props__['step'] = step
            __props__['tags'] = tags
            if task_name is None:
                raise TypeError("Missing required property 'task_name'")
            __props__['task_name'] = task_name
            __props__['timeout'] = timeout
            __props__['trigger'] = trigger
            __props__['creation_date'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:containerregistry/v20180901:Task"), pulumi.Alias(type_="azurerm:containerregistry/v20190401:Task"), pulumi.Alias(type_="azurerm:containerregistry/v20190601preview:Task")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Task, __self__).__init__(
            'azurerm:containerregistry/latest:Task',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Task':
        """
        Get an existing Task resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Task(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentConfiguration")
    def agent_configuration(self) -> pulumi.Output[Optional['outputs.AgentPropertiesResponse']]:
        """
        The machine configuration of the run agent.
        """
        return pulumi.get(self, "agent_configuration")

    @property
    @pulumi.getter(name="agentPoolName")
    def agent_pool_name(self) -> pulumi.Output[Optional[str]]:
        """
        The dedicated agent pool for the task.
        """
        return pulumi.get(self, "agent_pool_name")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        The creation date of task.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional['outputs.CredentialsResponse']]:
        """
        The properties that describes a set of credentials that will be used when this run is invoked.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.IdentityPropertiesResponse']]:
        """
        Identity for the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="isSystemTask")
    def is_system_task(self) -> pulumi.Output[Optional[bool]]:
        """
        The value of this property indicates whether the task resource is system task or not.
        """
        return pulumi.get(self, "is_system_task")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the resource. This cannot be changed after the resource is created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logTemplate")
    def log_template(self) -> pulumi.Output[Optional[str]]:
        """
        The template that describes the repository and tag information for run log artifact.
        """
        return pulumi.get(self, "log_template")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def platform(self) -> pulumi.Output[Optional['outputs.PlatformPropertiesResponse']]:
        """
        The platform properties against which the run has to happen.
        """
        return pulumi.get(self, "platform")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the task.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        The current status of task.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def step(self) -> pulumi.Output[Optional[Any]]:
        """
        The properties of a task step.
        """
        return pulumi.get(self, "step")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[float]]:
        """
        Run timeout in seconds.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def trigger(self) -> pulumi.Output[Optional['outputs.TriggerPropertiesResponse']]:
        """
        The properties that describe all triggers for the task.
        """
        return pulumi.get(self, "trigger")

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

