// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200801.Outputs
{

    [OutputType]
    public sealed class RoutingStorageContainerPropertiesResponseResult
    {
        /// <summary>
        /// Method used to authenticate against the storage endpoint
        /// </summary>
        public readonly string? AuthenticationType;
        /// <summary>
        /// Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
        /// </summary>
        public readonly int? BatchFrequencyInSeconds;
        /// <summary>
        /// The connection string of the storage account.
        /// </summary>
        public readonly string? ConnectionString;
        /// <summary>
        /// The name of storage container in the storage account.
        /// </summary>
        public readonly string ContainerName;
        /// <summary>
        /// Encoding that is used to serialize messages to blobs. Supported values are 'avro', 'avrodeflate', and 'JSON'. Default value is 'avro'.
        /// </summary>
        public readonly string? Encoding;
        /// <summary>
        /// The url of the storage endpoint. It must include the protocol https://
        /// </summary>
        public readonly string? EndpointUri;
        /// <summary>
        /// File name format for the blob. Default format is {iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}. All parameters are mandatory but can be reordered.
        /// </summary>
        public readonly string? FileNameFormat;
        /// <summary>
        /// Id of the storage container endpoint
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
        /// </summary>
        public readonly int? MaxChunkSizeInBytes;
        /// <summary>
        /// The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the resource group of the storage account.
        /// </summary>
        public readonly string? ResourceGroup;
        /// <summary>
        /// The subscription identifier of the storage account.
        /// </summary>
        public readonly string? SubscriptionId;

        [OutputConstructor]
        private RoutingStorageContainerPropertiesResponseResult(
            string? authenticationType,

            int? batchFrequencyInSeconds,

            string? connectionString,

            string containerName,

            string? encoding,

            string? endpointUri,

            string? fileNameFormat,

            string? id,

            int? maxChunkSizeInBytes,

            string name,

            string? resourceGroup,

            string? subscriptionId)
        {
            AuthenticationType = authenticationType;
            BatchFrequencyInSeconds = batchFrequencyInSeconds;
            ConnectionString = connectionString;
            ContainerName = containerName;
            Encoding = encoding;
            EndpointUri = endpointUri;
            FileNameFormat = fileNameFormat;
            Id = id;
            MaxChunkSizeInBytes = maxChunkSizeInBytes;
            Name = name;
            ResourceGroup = resourceGroup;
            SubscriptionId = subscriptionId;
        }
    }
}
