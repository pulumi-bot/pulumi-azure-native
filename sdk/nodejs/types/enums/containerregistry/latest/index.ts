// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const Action = {
    Allow: "Allow",
} as const;

/**
 * The action of virtual network rule.
 */
export type Action = (typeof Action)[keyof typeof Action];

export const Architecture = {
    Amd64: "amd64",
    X86: "x86",
    Arm: "arm",
} as const;

/**
 * The OS architecture.
 */
export type Architecture = (typeof Architecture)[keyof typeof Architecture];

export const BaseImageTriggerType = {
    All: "All",
    Runtime: "Runtime",
} as const;

/**
 * The type of the auto trigger for base image dependency updates.
 */
export type BaseImageTriggerType = (typeof BaseImageTriggerType)[keyof typeof BaseImageTriggerType];

export const DefaultAction = {
    Allow: "Allow",
    Deny: "Deny",
} as const;

/**
 * The default action of allow or deny when no other rules match.
 */
export type DefaultAction = (typeof DefaultAction)[keyof typeof DefaultAction];

export const OS = {
    Windows: "Windows",
    Linux: "Linux",
} as const;

/**
 * The operating system type required for the run.
 */
export type OS = (typeof OS)[keyof typeof OS];

export const PolicyStatus = {
    Enabled: "enabled",
    Disabled: "disabled",
} as const;

/**
 * The value that indicates whether the policy is enabled or not.
 */
export type PolicyStatus = (typeof PolicyStatus)[keyof typeof PolicyStatus];

export const ResourceIdentityType = {
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned, UserAssigned",
    None: "None",
} as const;

/**
 * The identity type.
 */
export type ResourceIdentityType = (typeof ResourceIdentityType)[keyof typeof ResourceIdentityType];

export const SecretObjectType = {
    Opaque: "Opaque",
    Vaultsecret: "Vaultsecret",
} as const;

/**
 * The type of the secret object which determines how the value of the secret object has to be
 * interpreted.
 */
export type SecretObjectType = (typeof SecretObjectType)[keyof typeof SecretObjectType];

export const SkuName = {
    Classic: "Classic",
    Basic: "Basic",
    Standard: "Standard",
    Premium: "Premium",
} as const;

/**
 * The SKU name of the container registry. Required for registry creation.
 */
export type SkuName = (typeof SkuName)[keyof typeof SkuName];

export const SourceControlType = {
    Github: "Github",
    VisualStudioTeamService: "VisualStudioTeamService",
} as const;

/**
 * The type of source control service.
 */
export type SourceControlType = (typeof SourceControlType)[keyof typeof SourceControlType];

export const SourceRegistryLoginMode = {
    None: "None",
    Default: "Default",
} as const;

/**
 * The authentication mode which determines the source registry login scope. The credentials for the source registry
 * will be generated using the given scope. These credentials will be used to login to
 * the source registry during the run.
 */
export type SourceRegistryLoginMode = (typeof SourceRegistryLoginMode)[keyof typeof SourceRegistryLoginMode];

export const SourceTriggerEvent = {
    Commit: "commit",
    Pullrequest: "pullrequest",
} as const;

export type SourceTriggerEvent = (typeof SourceTriggerEvent)[keyof typeof SourceTriggerEvent];

export const TaskStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * The current status of task.
 */
export type TaskStatus = (typeof TaskStatus)[keyof typeof TaskStatus];

export const TokenType = {
    PAT: "PAT",
    OAuth: "OAuth",
} as const;

/**
 * The type of Auth token.
 */
export type TokenType = (typeof TokenType)[keyof typeof TokenType];

export const TriggerStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * The current status of trigger.
 */
export type TriggerStatus = (typeof TriggerStatus)[keyof typeof TriggerStatus];

export const TrustPolicyType = {
    Notary: "Notary",
} as const;

/**
 * The type of trust policy.
 */
export type TrustPolicyType = (typeof TrustPolicyType)[keyof typeof TrustPolicyType];

export const Variant = {
    V6: "v6",
    V7: "v7",
    V8: "v8",
} as const;

/**
 * Variant of the CPU.
 */
export type Variant = (typeof Variant)[keyof typeof Variant];

export const WebhookAction = {
    Push: "push",
    Delete: "delete",
    Quarantine: "quarantine",
    Chart_push: "chart_push",
    Chart_delete: "chart_delete",
} as const;

export type WebhookAction = (typeof WebhookAction)[keyof typeof WebhookAction];

export const WebhookStatus = {
    Enabled: "enabled",
    Disabled: "disabled",
} as const;

/**
 * The status of the webhook at the time the operation was called.
 */
export type WebhookStatus = (typeof WebhookStatus)[keyof typeof WebhookStatus];
