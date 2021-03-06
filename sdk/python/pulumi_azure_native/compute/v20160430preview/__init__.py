# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .availability_set import *
from .disk import *
from .get_availability_set import *
from .get_disk import *
from .get_image import *
from .get_snapshot import *
from .get_virtual_machine import *
from .get_virtual_machine_extension import *
from .get_virtual_machine_scale_set import *
from .image import *
from .snapshot import *
from .virtual_machine import *
from .virtual_machine_extension import *
from .virtual_machine_scale_set import *
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
            if typ == "azure-native:compute/v20160430preview:AvailabilitySet":
                return AvailabilitySet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20160430preview:Disk":
                return Disk(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20160430preview:Image":
                return Image(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20160430preview:Snapshot":
                return Snapshot(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20160430preview:VirtualMachine":
                return VirtualMachine(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20160430preview:VirtualMachineExtension":
                return VirtualMachineExtension(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20160430preview:VirtualMachineScaleSet":
                return VirtualMachineScaleSet(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "compute/v20160430preview", _module_instance)

_register_module()
