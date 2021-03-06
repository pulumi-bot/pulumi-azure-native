# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .cluster import *
from .data_export import *
from .data_source import *
from .get_cluster import *
from .get_data_export import *
from .get_data_source import *
from .get_linked_service import *
from .get_linked_storage_account import *
from .get_machine_group import *
from .get_query import *
from .get_query_pack import *
from .get_saved_search import *
from .get_storage_insight_config import *
from .get_workspace import *
from .linked_service import *
from .linked_storage_account import *
from .list_workspace_keys import *
from .machine_group import *
from .query import *
from .query_pack import *
from .saved_search import *
from .storage_insight_config import *
from .workspace import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    latest,
    v20150320,
    v20151101preview,
    v20190801preview,
    v20190901preview,
    v20200301preview,
    v20200801,
    v20201001,
)

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-native:operationalinsights:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:DataExport":
                return DataExport(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:DataSource":
                return DataSource(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:LinkedService":
                return LinkedService(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:LinkedStorageAccount":
                return LinkedStorageAccount(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:MachineGroup":
                return MachineGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:Query":
                return Query(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:QueryPack":
                return QueryPack(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:SavedSearch":
                return SavedSearch(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:StorageInsightConfig":
                return StorageInsightConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:operationalinsights:Workspace":
                return Workspace(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "operationalinsights", _module_instance)

_register_module()
