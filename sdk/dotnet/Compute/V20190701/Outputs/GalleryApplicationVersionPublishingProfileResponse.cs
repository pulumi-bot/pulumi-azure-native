// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20190701.Outputs
{

    [OutputType]
    public sealed class GalleryApplicationVersionPublishingProfileResponse
    {
        /// <summary>
        /// Optional. Whether or not this application reports health.
        /// </summary>
        public readonly bool? EnableHealthCheck;
        /// <summary>
        /// The end of life date of the gallery Image Version. This property can be used for decommissioning purposes. This property is updatable.
        /// </summary>
        public readonly string? EndOfLifeDate;
        /// <summary>
        /// If set to true, Virtual Machines deployed from the latest version of the Image Definition won't use this Image Version.
        /// </summary>
        public readonly bool? ExcludeFromLatest;
        public readonly Outputs.UserArtifactManageResponse? ManageActions;
        /// <summary>
        /// The timestamp for when the gallery Image Version is published.
        /// </summary>
        public readonly string PublishedDate;
        /// <summary>
        /// The number of replicas of the Image Version to be created per region. This property would take effect for a region when regionalReplicaCount is not specified. This property is updatable.
        /// </summary>
        public readonly int? ReplicaCount;
        /// <summary>
        /// The source image from which the Image Version is going to be created.
        /// </summary>
        public readonly Outputs.UserArtifactSourceResponse Source;
        /// <summary>
        /// Specifies the storage account type to be used to store the image. This property is not updatable.
        /// </summary>
        public readonly string? StorageAccountType;
        /// <summary>
        /// The target regions where the Image Version is going to be replicated to. This property is updatable.
        /// </summary>
        public readonly ImmutableArray<Outputs.TargetRegionResponse> TargetRegions;

        [OutputConstructor]
        private GalleryApplicationVersionPublishingProfileResponse(
            bool? enableHealthCheck,

            string? endOfLifeDate,

            bool? excludeFromLatest,

            Outputs.UserArtifactManageResponse? manageActions,

            string publishedDate,

            int? replicaCount,

            Outputs.UserArtifactSourceResponse source,

            string? storageAccountType,

            ImmutableArray<Outputs.TargetRegionResponse> targetRegions)
        {
            EnableHealthCheck = enableHealthCheck;
            EndOfLifeDate = endOfLifeDate;
            ExcludeFromLatest = excludeFromLatest;
            ManageActions = manageActions;
            PublishedDate = publishedDate;
            ReplicaCount = replicaCount;
            Source = source;
            StorageAccountType = storageAccountType;
            TargetRegions = targetRegions;
        }
    }
}
