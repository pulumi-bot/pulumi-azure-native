// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./defaultRollout";
export * from "./getDefaultRollout";
export * from "./getNotificationRegistration";
export * from "./getProviderRegistration";
export * from "./getResourceTypeRegistration";
export * from "./getSkus";
export * from "./getSkusNestedResourceTypeFirst";
export * from "./getSkusNestedResourceTypeSecond";
export * from "./getSkusNestedResourceTypeThird";
export * from "./notificationRegistration";
export * from "./operationByProviderRegistration";
export * from "./providerRegistration";
export * from "./resourceTypeRegistration";
export * from "./skus";
export * from "./skusNestedResourceTypeFirst";
export * from "./skusNestedResourceTypeSecond";
export * from "./skusNestedResourceTypeThird";

// Export enums:
export * from "../../types/enums/providerhub/latest";

// Import resources to register:
import { DefaultRollout } from "./defaultRollout";
import { NotificationRegistration } from "./notificationRegistration";
import { OperationByProviderRegistration } from "./operationByProviderRegistration";
import { ProviderRegistration } from "./providerRegistration";
import { ResourceTypeRegistration } from "./resourceTypeRegistration";
import { Skus } from "./skus";
import { SkusNestedResourceTypeFirst } from "./skusNestedResourceTypeFirst";
import { SkusNestedResourceTypeSecond } from "./skusNestedResourceTypeSecond";
import { SkusNestedResourceTypeThird } from "./skusNestedResourceTypeThird";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:providerhub/latest:DefaultRollout":
                return new DefaultRollout(name, <any>undefined, { urn })
            case "azure-native:providerhub/latest:NotificationRegistration":
                return new NotificationRegistration(name, <any>undefined, { urn })
            case "azure-native:providerhub/latest:OperationByProviderRegistration":
                return new OperationByProviderRegistration(name, <any>undefined, { urn })
            case "azure-native:providerhub/latest:ProviderRegistration":
                return new ProviderRegistration(name, <any>undefined, { urn })
            case "azure-native:providerhub/latest:ResourceTypeRegistration":
                return new ResourceTypeRegistration(name, <any>undefined, { urn })
            case "azure-native:providerhub/latest:Skus":
                return new Skus(name, <any>undefined, { urn })
            case "azure-native:providerhub/latest:SkusNestedResourceTypeFirst":
                return new SkusNestedResourceTypeFirst(name, <any>undefined, { urn })
            case "azure-native:providerhub/latest:SkusNestedResourceTypeSecond":
                return new SkusNestedResourceTypeSecond(name, <any>undefined, { urn })
            case "azure-native:providerhub/latest:SkusNestedResourceTypeThird":
                return new SkusNestedResourceTypeThird(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "providerhub/latest", _module)
