# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuthenticationType',
    'ProjectSourcePlatform',
    'ProjectTargetPlatform',
    'ServerLevelPermissionsGroup',
]


class AuthenticationType(str, Enum):
    """
    Authentication type to use for connection
    """
    NONE = "None"
    WINDOWS_AUTHENTICATION = "WindowsAuthentication"
    SQL_AUTHENTICATION = "SqlAuthentication"
    ACTIVE_DIRECTORY_INTEGRATED = "ActiveDirectoryIntegrated"
    ACTIVE_DIRECTORY_PASSWORD = "ActiveDirectoryPassword"


class ProjectSourcePlatform(str, Enum):
    """
    Source platform for the project
    """
    SQL = "SQL"
    UNKNOWN = "Unknown"


class ProjectTargetPlatform(str, Enum):
    """
    Target platform for the project
    """
    SQLDB = "SQLDB"
    UNKNOWN = "Unknown"


class ServerLevelPermissionsGroup(str, Enum):
    """
    Permission group for validations
    """
    DEFAULT = "Default"
    MIGRATION_FROM_SQL_SERVER_TO_AZURE_DB = "MigrationFromSqlServerToAzureDB"
