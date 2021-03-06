// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./deviceSecurityGroup";
export * from "./getDeviceSecurityGroup";
export * from "./getIotSecuritySolution";
export * from "./iotSecuritySolution";

// Export enums:
export * from "../../types/enums/security/v20190801";

// Import resources to register:
import { DeviceSecurityGroup } from "./deviceSecurityGroup";
import { IotSecuritySolution } from "./iotSecuritySolution";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:security/v20190801:DeviceSecurityGroup":
                return new DeviceSecurityGroup(name, <any>undefined, { urn })
            case "azure-native:security/v20190801:IotSecuritySolution":
                return new IotSecuritySolution(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "security/v20190801", _module)
