# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .disk import *
from .disk_access import *
from .disk_access_a_private_endpoint_connection import *
from .disk_encryption_set import *
from .gallery import *
from .gallery_application import *
from .gallery_application_version import *
from .gallery_image import *
from .gallery_image_version import *
from .get_disk import *
from .get_disk_access import *
from .get_disk_access_a_private_endpoint_connection import *
from .get_disk_encryption_set import *
from .get_gallery import *
from .get_gallery_application import *
from .get_gallery_application_version import *
from .get_gallery_image import *
from .get_gallery_image_version import *
from .get_snapshot import *
from .snapshot import *
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
            if typ == "azure-native:compute/v20200930:Disk":
                return Disk(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:DiskAccess":
                return DiskAccess(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:DiskAccessAPrivateEndpointConnection":
                return DiskAccessAPrivateEndpointConnection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:DiskEncryptionSet":
                return DiskEncryptionSet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:Gallery":
                return Gallery(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:GalleryApplication":
                return GalleryApplication(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:GalleryApplicationVersion":
                return GalleryApplicationVersion(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:GalleryImage":
                return GalleryImage(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:GalleryImageVersion":
                return GalleryImageVersion(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:compute/v20200930:Snapshot":
                return Snapshot(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "compute/v20200930", _module_instance)

_register_module()
