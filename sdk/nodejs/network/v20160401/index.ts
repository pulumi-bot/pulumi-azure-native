// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getRecordSet";
export * from "./getZone";
export * from "./recordSet";
export * from "./zone";

// Export enums:
export * from "../../types/enums/network/v20160401";

// Import resources to register:
import { RecordSet } from "./recordSet";
import { Zone } from "./zone";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:network/v20160401:RecordSet":
                return new RecordSet(name, <any>undefined, { urn })
            case "azure-native:network/v20160401:Zone":
                return new Zone(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "network/v20160401", _module)
