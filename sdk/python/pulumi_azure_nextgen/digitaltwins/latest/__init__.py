# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .digital_twin import *
from .digital_twins_endpoint import *
from .get_digital_twin import *
from .get_digital_twins_endpoint import *
from .get_private_endpoint_connection import *
from .private_endpoint_connection import *
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
            if typ == "azure-nextgen:digitaltwins/latest:DigitalTwin":
                return DigitalTwin(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:digitaltwins/latest:DigitalTwinsEndpoint":
                return DigitalTwinsEndpoint(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:digitaltwins/latest:PrivateEndpointConnection":
                return PrivateEndpointConnection(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-nextgen", "digitaltwins/latest", _module_instance)

_register_module()
