# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .cluster import *
from .database import *
from .event_hub_connection import *
from .get_cluster import *
from .get_database import *
from .get_event_hub_connection import *
from .list_database_principals import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-native:kusto/v20170907privatepreview:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:kusto/v20170907privatepreview:Database":
                return Database(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:kusto/v20170907privatepreview:EventHubConnection":
                return EventHubConnection(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "kusto/v20170907privatepreview", _module_instance)

_register_module()
