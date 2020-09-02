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
    return pulumi.runtime.invoke("azurerm:security/latest:getAssessmentMetadataInSubscription", {
        "assessmentMetadataName": args.assessmentMetadataName,
    }, opts);
}

export interface GetAssessmentMetadataInSubscriptionArgs {
    /**
     * The Assessment Key - Unique key for the assessment type
     */
    readonly assessmentMetadataName: string;
}

/**
 * Security assessment metadata
 */
export interface GetAssessmentMetadataInSubscriptionResult {
    /**
     * BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
     */
    readonly assessmentType: string;
    readonly category?: string[];
    /**
     * Human readable description of the assessment
     */
    readonly description?: string;
    /**
     * User friendly display name of the assessment
     */
    readonly displayName: string;
    /**
     * The implementation effort required to remediate this assessment
     */
    readonly implementationEffort?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Describes the partner that created the assessment
     */
    readonly partnerData?: outputs.security.latest.SecurityAssessmentMetadataPartnerDataResponse;
    /**
     * Azure resource ID of the policy definition that turns this assessment calculation on
     */
    readonly policyDefinitionId: string;
    /**
     * True if this assessment is in preview release status
     */
    readonly preview?: boolean;
    /**
     * Human readable description of what you should do to mitigate this security issue
     */
    readonly remediationDescription?: string;
    /**
     * The severity level of the assessment
     */
    readonly severity: string;
    readonly threats?: string[];
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * The user impact of the assessment
     */
    readonly userImpact?: string;
}