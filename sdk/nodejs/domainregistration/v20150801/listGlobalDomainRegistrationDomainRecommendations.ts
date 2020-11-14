// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listGlobalDomainRegistrationDomainRecommendations(args?: ListGlobalDomainRegistrationDomainRecommendationsArgs, opts?: pulumi.InvokeOptions): Promise<ListGlobalDomainRegistrationDomainRecommendationsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:domainregistration/v20150801:listGlobalDomainRegistrationDomainRecommendations", {
        "keywords": args.keywords,
        "maxDomainRecommendations": args.maxDomainRecommendations,
    }, opts);
}

export interface ListGlobalDomainRegistrationDomainRecommendationsArgs {
    /**
     * Keywords to be used for generating domain recommendations
     */
    readonly keywords?: string;
    /**
     * Maximum number of recommendations
     */
    readonly maxDomainRecommendations?: number;
}

/**
 * Collection of domain name identifiers
 */
export interface ListGlobalDomainRegistrationDomainRecommendationsResult {
    /**
     * Link to next page of resources
     */
    readonly nextLink?: string;
    /**
     * Collection of resources
     */
    readonly value?: outputs.domainregistration.v20150801.NameIdentifierResponse[];
}
