// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getPrivateEndpointConnection";
export * from "./getProtectedItem";
export * from "./getProtectionContainer";
export * from "./getProtectionPolicy";
export * from "./privateEndpointConnection";
export * from "./protectedItem";
export * from "./protectionContainer";
export * from "./protectionPolicy";

// Export enums:
export * from "../../types/enums/recoveryservices/v20210101";

// Import resources to register:
import { PrivateEndpointConnection } from "./privateEndpointConnection";
import { ProtectedItem } from "./protectedItem";
import { ProtectionContainer } from "./protectionContainer";
import { ProtectionPolicy } from "./protectionPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-nextgen:recoveryservices/v20210101:PrivateEndpointConnection":
                return new PrivateEndpointConnection(name, <any>undefined, { urn })
            case "azure-nextgen:recoveryservices/v20210101:ProtectedItem":
                return new ProtectedItem(name, <any>undefined, { urn })
            case "azure-nextgen:recoveryservices/v20210101:ProtectionContainer":
                return new ProtectionContainer(name, <any>undefined, { urn })
            case "azure-nextgen:recoveryservices/v20210101:ProtectionPolicy":
                return new ProtectionPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-nextgen", "recoveryservices/v20210101", _module)
