// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./availabilitySet";
export * from "./disk";
export * from "./getAvailabilitySet";
export * from "./getDisk";
export * from "./getImage";
export * from "./getSnapshot";
export * from "./getVirtualMachine";
export * from "./getVirtualMachineExtension";
export * from "./getVirtualMachineScaleSet";
export * from "./image";
export * from "./snapshot";
export * from "./virtualMachine";
export * from "./virtualMachineExtension";
export * from "./virtualMachineScaleSet";

// Export enums:
export * from "../../types/enums/compute/v20160430preview";

// Import resources to register:
import { AvailabilitySet } from "./availabilitySet";
import { Disk } from "./disk";
import { Image } from "./image";
import { Snapshot } from "./snapshot";
import { VirtualMachine } from "./virtualMachine";
import { VirtualMachineExtension } from "./virtualMachineExtension";
import { VirtualMachineScaleSet } from "./virtualMachineScaleSet";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:compute/v20160430preview:AvailabilitySet":
                return new AvailabilitySet(name, <any>undefined, { urn })
            case "azure-native:compute/v20160430preview:Disk":
                return new Disk(name, <any>undefined, { urn })
            case "azure-native:compute/v20160430preview:Image":
                return new Image(name, <any>undefined, { urn })
            case "azure-native:compute/v20160430preview:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "azure-native:compute/v20160430preview:VirtualMachine":
                return new VirtualMachine(name, <any>undefined, { urn })
            case "azure-native:compute/v20160430preview:VirtualMachineExtension":
                return new VirtualMachineExtension(name, <any>undefined, { urn })
            case "azure-native:compute/v20160430preview:VirtualMachineScaleSet":
                return new VirtualMachineScaleSet(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "compute/v20160430preview", _module)
