// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getAssessmentMetadataInSubscription(args: GetAssessmentMetadataInSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetAssessmentMetadataInSubscriptionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:security/v20200101:getAssessmentMetadataInSubscription", {
        "name": args.name,
    }, opts);
}

export interface GetAssessmentMetadataInSubscriptionArgs {
    /**
     * The Assessment Key - Unique key for the assessment type
     */
    readonly name: string;
}

/**
 * Security assessment metadata
 */
export interface GetAssessmentMetadataInSubscriptionResult {
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Describes properties of an assessment metadata.
     */
    readonly properties: outputs.security.v20200101.SecurityAssessmentMetadataPropertiesResponse;
    /**
     * Resource type
     */
    readonly type: string;
}