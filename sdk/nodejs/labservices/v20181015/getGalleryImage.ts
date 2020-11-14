// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getGalleryImage(args: GetGalleryImageArgs, opts?: pulumi.InvokeOptions): Promise<GetGalleryImageResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:labservices/v20181015:getGalleryImage", {
        "expand": args.expand,
        "galleryImageName": args.galleryImageName,
        "labAccountName": args.labAccountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGalleryImageArgs {
    /**
     * Specify the $expand query. Example: 'properties($select=author)'
     */
    readonly expand?: string;
    /**
     * The name of the gallery Image.
     */
    readonly galleryImageName: string;
    /**
     * The name of the lab Account.
     */
    readonly labAccountName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Represents an image from the Azure Marketplace
 */
export interface GetGalleryImageResult {
    /**
     * The author of the gallery image.
     */
    readonly author: string;
    /**
     * The creation date of the gallery image.
     */
    readonly createdDate: string;
    /**
     * The description of the gallery image.
     */
    readonly description: string;
    /**
     * The icon of the gallery image.
     */
    readonly icon: string;
    /**
     * The image reference of the gallery image.
     */
    readonly imageReference: outputs.labservices.v20181015.GalleryImageReferenceResponse;
    /**
     * Indicates whether this gallery image is enabled.
     */
    readonly isEnabled?: boolean;
    /**
     * Indicates whether this gallery has been overridden for this lab account
     */
    readonly isOverride?: boolean;
    /**
     * Indicates if the plan has been authorized for programmatic deployment.
     */
    readonly isPlanAuthorized?: boolean;
    /**
     * The details of the latest operation. ex: status, error
     */
    readonly latestOperationResult: outputs.labservices.v20181015.LatestOperationResultResponse;
    /**
     * The location of the resource.
     */
    readonly location?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The third party plan that applies to this image
     */
    readonly planId: string;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState?: string;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    readonly uniqueIdentifier?: string;
}
