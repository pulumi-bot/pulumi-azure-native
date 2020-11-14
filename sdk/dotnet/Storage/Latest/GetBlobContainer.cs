// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.Latest
{
    public static class GetBlobContainer
    {
        public static Task<GetBlobContainerResult> InvokeAsync(GetBlobContainerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBlobContainerResult>("azure-nextgen:storage/latest:getBlobContainer", args ?? new GetBlobContainerArgs(), options.WithVersion());
    }


    public sealed class GetBlobContainerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
        /// </summary>
        [Input("containerName", required: true)]
        public string ContainerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetBlobContainerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBlobContainerResult
    {
        /// <summary>
        /// Default the container to use specified encryption scope for all writes.
        /// </summary>
        public readonly string? DefaultEncryptionScope;
        /// <summary>
        /// Indicates whether the blob container was deleted.
        /// </summary>
        public readonly bool Deleted;
        /// <summary>
        /// Blob container deletion time.
        /// </summary>
        public readonly string DeletedTime;
        /// <summary>
        /// Block override of encryption scope from the container default.
        /// </summary>
        public readonly bool? DenyEncryptionScopeOverride;
        /// <summary>
        /// Resource Etag.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.
        /// </summary>
        public readonly bool HasImmutabilityPolicy;
        /// <summary>
        /// The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
        /// </summary>
        public readonly bool HasLegalHold;
        /// <summary>
        /// The ImmutabilityPolicy property of the container.
        /// </summary>
        public readonly Outputs.ImmutabilityPolicyPropertiesResponse ImmutabilityPolicy;
        /// <summary>
        /// Returns the date and time the container was last modified.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
        /// </summary>
        public readonly string LeaseDuration;
        /// <summary>
        /// Lease state of the container.
        /// </summary>
        public readonly string LeaseState;
        /// <summary>
        /// The lease status of the container.
        /// </summary>
        public readonly string LeaseStatus;
        /// <summary>
        /// The LegalHold property of the container.
        /// </summary>
        public readonly Outputs.LegalHoldPropertiesResponse LegalHold;
        /// <summary>
        /// A name-value pair to associate with the container as metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Metadata;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies whether data in the container may be accessed publicly and the level of access.
        /// </summary>
        public readonly string? PublicAccess;
        /// <summary>
        /// Remaining retention days for soft deleted blob container.
        /// </summary>
        public readonly int RemainingRetentionDays;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The version of the deleted blob container.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetBlobContainerResult(
            string? defaultEncryptionScope,

            bool deleted,

            string deletedTime,

            bool? denyEncryptionScopeOverride,

            string etag,

            bool hasImmutabilityPolicy,

            bool hasLegalHold,

            Outputs.ImmutabilityPolicyPropertiesResponse immutabilityPolicy,

            string lastModifiedTime,

            string leaseDuration,

            string leaseState,

            string leaseStatus,

            Outputs.LegalHoldPropertiesResponse legalHold,

            ImmutableDictionary<string, string>? metadata,

            string name,

            string? publicAccess,

            int remainingRetentionDays,

            string type,

            string version)
        {
            DefaultEncryptionScope = defaultEncryptionScope;
            Deleted = deleted;
            DeletedTime = deletedTime;
            DenyEncryptionScopeOverride = denyEncryptionScopeOverride;
            Etag = etag;
            HasImmutabilityPolicy = hasImmutabilityPolicy;
            HasLegalHold = hasLegalHold;
            ImmutabilityPolicy = immutabilityPolicy;
            LastModifiedTime = lastModifiedTime;
            LeaseDuration = leaseDuration;
            LeaseState = leaseState;
            LeaseStatus = leaseStatus;
            LegalHold = legalHold;
            Metadata = metadata;
            Name = name;
            PublicAccess = publicAccess;
            RemainingRetentionDays = remainingRetentionDays;
            Type = type;
            Version = version;
        }
    }
}
