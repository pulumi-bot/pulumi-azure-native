# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListWebAppBackupStatusSecretsResult:
    """
    Backup description.
    """
    def __init__(__self__, kind=None, name=None, properties=None, type=None):
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        __self__.kind = kind
        """
        Kind of resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource Name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        BackupItem resource specific properties
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type.
        """


class AwaitableListWebAppBackupStatusSecretsResult(ListWebAppBackupStatusSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListWebAppBackupStatusSecretsResult(
            kind=self.kind,
            name=self.name,
            properties=self.properties,
            type=self.type)


def list_web_app_backup_status_secrets(kind=None, name=None, properties=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str kind: Kind of resource.
    :param str name: ID of backup.
    :param dict properties: BackupRequest resource specific properties
    :param str resource_group_name: Name of the resource group to which the resource belongs.

    The **properties** object supports the following:

      * `backup_name` (`str`) - Name of the backup.
      * `backup_schedule` (`dict`) - Schedule for the backup if it is executed periodically.
        * `frequency_interval` (`float`) - How often the backup should be executed (e.g. for weekly backup, this should be set to 7 and FrequencyUnit should be set to Day)
        * `frequency_unit` (`str`) - The unit of time for how often the backup should be executed (e.g. for weekly backup, this should be set to Day and FrequencyInterval should be set to 7)
        * `keep_at_least_one_backup` (`bool`) - True if the retention policy should always keep at least one backup in the storage account, regardless how old it is; false otherwise.
        * `retention_period_in_days` (`float`) - After how many days backups should be deleted.
        * `start_time` (`str`) - When the schedule should start working.

      * `databases` (`list`) - Databases included in the backup.
        * `connection_string` (`str`) - Contains a connection string to a database which is being backed up or restored. If the restore should happen to a new database, the database name inside is the new one.
        * `connection_string_name` (`str`) - Contains a connection string name that is linked to the SiteConfig.ConnectionStrings.
          This is used during restore with overwrite connection strings options.
        * `database_type` (`str`) - Database type (e.g. SqlAzure / MySql).
        * `name` (`str`)

      * `enabled` (`bool`) - True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
      * `storage_account_url` (`str`) - SAS URL to the container.
    """
    __args__ = dict()
    __args__['kind'] = kind
    __args__['name'] = name
    __args__['properties'] = properties
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:web/v20190801:listWebAppBackupStatusSecrets', __args__, opts=opts).value

    return AwaitableListWebAppBackupStatusSecretsResult(
        kind=__ret__.get('kind'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
