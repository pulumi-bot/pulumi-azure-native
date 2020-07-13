// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.Outputs
{

    [OutputType]
    public sealed class CustomImagePropertiesResponse
    {
        /// <summary>
        /// The author of the custom image.
        /// </summary>
        public readonly string? Author;
        /// <summary>
        /// The creation date of the custom image.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// Storage information about the plan related to this custom image
        /// </summary>
        public readonly Outputs.CustomImagePropertiesFromPlanResponse? CustomImagePlan;
        /// <summary>
        /// Storage information about the data disks present in the custom image
        /// </summary>
        public readonly ImmutableArray<Outputs.DataDiskStorageTypeInfoResponse> DataDiskStorageInfo;
        /// <summary>
        /// The description of the custom image.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
        /// </summary>
        public readonly bool? IsPlanAuthorized;
        /// <summary>
        /// The Managed Image Id backing the custom image.
        /// </summary>
        public readonly string? ManagedImageId;
        /// <summary>
        /// The Managed Snapshot Id backing the custom image.
        /// </summary>
        public readonly string? ManagedSnapshotId;
        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        public readonly string UniqueIdentifier;
        /// <summary>
        /// The VHD from which the image is to be created.
        /// </summary>
        public readonly Outputs.CustomImagePropertiesCustomResponse? Vhd;
        /// <summary>
        /// The virtual machine from which the image is to be created.
        /// </summary>
        public readonly Outputs.CustomImagePropertiesFromVmResponse? Vm;

        [OutputConstructor]
        private CustomImagePropertiesResponse(
            string? author,

            string creationDate,

            Outputs.CustomImagePropertiesFromPlanResponse? customImagePlan,

            ImmutableArray<Outputs.DataDiskStorageTypeInfoResponse> dataDiskStorageInfo,

            string? description,

            bool? isPlanAuthorized,

            string? managedImageId,

            string? managedSnapshotId,

            string provisioningState,

            string uniqueIdentifier,

            Outputs.CustomImagePropertiesCustomResponse? vhd,

            Outputs.CustomImagePropertiesFromVmResponse? vm)
        {
            Author = author;
            CreationDate = creationDate;
            CustomImagePlan = customImagePlan;
            DataDiskStorageInfo = dataDiskStorageInfo;
            Description = description;
            IsPlanAuthorized = isPlanAuthorized;
            ManagedImageId = managedImageId;
            ManagedSnapshotId = managedSnapshotId;
            ProvisioningState = provisioningState;
            UniqueIdentifier = uniqueIdentifier;
            Vhd = vhd;
            Vm = vm;
        }
    }
}