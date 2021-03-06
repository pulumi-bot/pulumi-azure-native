# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .get_key import *
from .get_managed_hsm import *
from .get_private_endpoint_connection import *
from .get_secret import *
from .get_vault import *
from .key import *
from .managed_hsm import *
from .private_endpoint_connection import *
from .secret import *
from .vault import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    latest,
    v20150601,
    v20161001,
    v20180214,
    v20180214preview,
    v20190901,
    v20200401preview,
)

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-native:keyvault:Key":
                return Key(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:keyvault:ManagedHsm":
                return ManagedHsm(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:keyvault:PrivateEndpointConnection":
                return PrivateEndpointConnection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:keyvault:Secret":
                return Secret(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:keyvault:Vault":
                return Vault(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "keyvault", _module_instance)

_register_module()
