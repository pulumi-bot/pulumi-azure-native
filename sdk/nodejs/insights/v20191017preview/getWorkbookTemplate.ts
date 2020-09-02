// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getWorkbookTemplate(args: GetWorkbookTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkbookTemplateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:insights/v20191017preview:getWorkbookTemplate", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetWorkbookTemplateArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Application Insights component resource.
     */
    readonly resourceName: string;
}

/**
 * An Application Insights workbook template definition.
 */
export interface GetWorkbookTemplateResult {
    /**
     * Information about the author of the workbook template.
     */
    readonly author?: string;
    /**
     * Workbook galleries supported by the template.
     */
    readonly galleries: outputs.insights.v20191017preview.WorkbookTemplateGalleryResponse[];
    /**
     * Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
     */
    readonly localized?: {[key: string]: outputs.insights.v20191017preview.WorkbookTemplateLocalizedGalleryResponse[]};
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Azure resource name.
     */
    readonly name: string;
    /**
     * Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
     */
    readonly priority?: number;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Valid JSON object containing workbook template payload.
     */
    readonly templateData: {[key: string]: any};
    /**
     * Azure resource type
     */
    readonly type: string;
}
