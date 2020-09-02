// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Databricks.Latest.Outputs
{

    [OutputType]
    public sealed class WorkspaceCustomParametersResponseResult
    {
        /// <summary>
        /// The ID of a Azure Machine Learning workspace to link with Databricks workspace
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponseResult? AmlWorkspaceId;
        /// <summary>
        /// The name of the Private Subnet within the Virtual Network
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponseResult? CustomPrivateSubnetName;
        /// <summary>
        /// The name of a Public Subnet within the Virtual Network
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponseResult? CustomPublicSubnetName;
        /// <summary>
        /// The ID of a Virtual Network where this Databricks Cluster should be created
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponseResult? CustomVirtualNetworkId;
        /// <summary>
        /// Should the Public IP be Disabled?
        /// </summary>
        public readonly Outputs.WorkspaceCustomBooleanParameterResponseResult? EnableNoPublicIp;
        /// <summary>
        /// Contains the encryption details for Customer-Managed Key (CMK) enabled workspace.
        /// </summary>
        public readonly Outputs.WorkspaceEncryptionParameterResponseResult? Encryption;
        /// <summary>
        /// Prepare the workspace for encryption. Enables the Managed Identity for managed storage account.
        /// </summary>
        public readonly Outputs.WorkspaceCustomBooleanParameterResponseResult? PrepareEncryption;
        /// <summary>
        /// A boolean indicating whether or not the DBFS root file system will be enabled with secondary layer of encryption with platform managed keys for data at rest.
        /// </summary>
        public readonly Outputs.WorkspaceCustomBooleanParameterResponseResult? RequireInfrastructureEncryption;

        [OutputConstructor]
        private WorkspaceCustomParametersResponseResult(
            Outputs.WorkspaceCustomStringParameterResponseResult? amlWorkspaceId,

            Outputs.WorkspaceCustomStringParameterResponseResult? customPrivateSubnetName,

            Outputs.WorkspaceCustomStringParameterResponseResult? customPublicSubnetName,

            Outputs.WorkspaceCustomStringParameterResponseResult? customVirtualNetworkId,

            Outputs.WorkspaceCustomBooleanParameterResponseResult? enableNoPublicIp,

            Outputs.WorkspaceEncryptionParameterResponseResult? encryption,

            Outputs.WorkspaceCustomBooleanParameterResponseResult? prepareEncryption,

            Outputs.WorkspaceCustomBooleanParameterResponseResult? requireInfrastructureEncryption)
        {
            AmlWorkspaceId = amlWorkspaceId;
            CustomPrivateSubnetName = customPrivateSubnetName;
            CustomPublicSubnetName = customPublicSubnetName;
            CustomVirtualNetworkId = customVirtualNetworkId;
            EnableNoPublicIp = enableNoPublicIp;
            Encryption = encryption;
            PrepareEncryption = prepareEncryption;
            RequireInfrastructureEncryption = requireInfrastructureEncryption;
        }
    }
}