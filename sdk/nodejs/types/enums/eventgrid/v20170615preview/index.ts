// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EndpointType = {
    WebHook: "WebHook",
} as const;

/**
 * Type of the endpoint for the event subscription destination
 */
export type EndpointType = (typeof EndpointType)[keyof typeof EndpointType];
