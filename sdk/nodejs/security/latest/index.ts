// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./adaptiveApplicationControl";
export * from "./assessment";
export * from "./assessmentMetadataInSubscription";
export * from "./deviceSecurityGroup";
export * from "./getAdaptiveApplicationControl";
export * from "./getAssessment";
export * from "./getAssessmentMetadataInSubscription";
export * from "./getDeviceSecurityGroup";
export * from "./getIotSecuritySolution";
export * from "./getJitNetworkAccessPolicy";
export * from "./getServerVulnerabilityAssessment";
export * from "./iotSecuritySolution";
export * from "./jitNetworkAccessPolicy";
export * from "./serverVulnerabilityAssessment";

// Export enums:
export * from "../../types/enums/security/latest";

// Import resources to register:
import { AdaptiveApplicationControl } from "./adaptiveApplicationControl";
import { Assessment } from "./assessment";
import { AssessmentMetadataInSubscription } from "./assessmentMetadataInSubscription";
import { DeviceSecurityGroup } from "./deviceSecurityGroup";
import { IotSecuritySolution } from "./iotSecuritySolution";
import { JitNetworkAccessPolicy } from "./jitNetworkAccessPolicy";
import { ServerVulnerabilityAssessment } from "./serverVulnerabilityAssessment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:security/latest:AdaptiveApplicationControl":
                return new AdaptiveApplicationControl(name, <any>undefined, { urn })
            case "azure-native:security/latest:Assessment":
                return new Assessment(name, <any>undefined, { urn })
            case "azure-native:security/latest:AssessmentMetadataInSubscription":
                return new AssessmentMetadataInSubscription(name, <any>undefined, { urn })
            case "azure-native:security/latest:DeviceSecurityGroup":
                return new DeviceSecurityGroup(name, <any>undefined, { urn })
            case "azure-native:security/latest:IotSecuritySolution":
                return new IotSecuritySolution(name, <any>undefined, { urn })
            case "azure-native:security/latest:JitNetworkAccessPolicy":
                return new JitNetworkAccessPolicy(name, <any>undefined, { urn })
            case "azure-native:security/latest:ServerVulnerabilityAssessment":
                return new ServerVulnerabilityAssessment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "security/latest", _module)
