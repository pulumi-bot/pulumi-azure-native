# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .configuration_store import *
from .get_configuration_store import *
from .get_key_value import *
from .get_private_endpoint_connection import *
from .key_value import *
from .list_configuration_store_key_value import *
from .list_configuration_store_keys import *
from .private_endpoint_connection import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    latest,
    v20190201preview,
    v20191001,
    v20191101preview,
    v20200601,
    v20200701preview,
)

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-native:appconfiguration:ConfigurationStore":
                return ConfigurationStore(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:appconfiguration:KeyValue":
                return KeyValue(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:appconfiguration:PrivateEndpointConnection":
                return PrivateEndpointConnection(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "appconfiguration", _module_instance)

_register_module()
