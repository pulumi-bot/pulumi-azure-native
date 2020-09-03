// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.V20170101
{
    /// <summary>
    /// The connector mapping resource format.
    /// 
    /// ## Example Usage
    /// ### ConnectorMappings_CreateOrUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var connectorMapping = new AzureRM.CustomerInsights.V20170101.ConnectorMapping("connectorMapping", new AzureRM.CustomerInsights.V20170101.ConnectorMappingArgs
    ///         {
    ///             ConnectorName = "testConnector8858",
    ///             Description = "Test mapping",
    ///             DisplayName = "testMapping12491",
    ///             EntityType = "Interaction",
    ///             EntityTypeName = "TestInteractionType2967",
    ///             HubName = "sdkTestHub",
    ///             MappingName = "testMapping12491",
    ///             MappingProperties = new AzureRM.CustomerInsights.V20170101.Inputs.ConnectorMappingPropertiesArgs
    ///             {
    ///                 Availability = new AzureRM.CustomerInsights.V20170101.Inputs.ConnectorMappingAvailabilityArgs
    ///                 {
    ///                     Frequency = "Hour",
    ///                     Interval = 5,
    ///                 },
    ///                 CompleteOperation = new AzureRM.CustomerInsights.V20170101.Inputs.ConnectorMappingCompleteOperationArgs
    ///                 {
    ///                     CompletionOperationType = "DeleteFile",
    ///                     DestinationFolder = "fakePath",
    ///                 },
    ///                 ErrorManagement = new AzureRM.CustomerInsights.V20170101.Inputs.ConnectorMappingErrorManagementArgs
    ///                 {
    ///                     ErrorLimit = 10,
    ///                     ErrorManagementType = "StopImport",
    ///                 },
    ///                 FileFilter = "unknown",
    ///                 FolderPath = "http://sample.dne/file",
    ///                 Format = new AzureRM.CustomerInsights.V20170101.Inputs.ConnectorMappingFormatArgs
    ///                 {
    ///                     ColumnDelimiter = "|",
    ///                     FormatType = "TextFormat",
    ///                 },
    ///                 HasHeader = false,
    ///                 Structure = 
    ///                 {
    ///                     new AzureRM.CustomerInsights.V20170101.Inputs.ConnectorMappingStructureArgs
    ///                     {
    ///                         ColumnName = "unknown1",
    ///                         IsEncrypted = false,
    ///                         PropertyName = "unknwon1",
    ///                     },
    ///                     new AzureRM.CustomerInsights.V20170101.Inputs.ConnectorMappingStructureArgs
    ///                     {
    ///                         ColumnName = "unknown2",
    ///                         IsEncrypted = true,
    ///                         PropertyName = "unknwon2",
    ///                     },
    ///                 },
    ///             },
    ///             ResourceGroupName = "TestHubRG",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ConnectorMapping : Pulumi.CustomResource
    {
        /// <summary>
        /// The connector mapping name
        /// </summary>
        [Output("connectorMappingName")]
        public Output<string> ConnectorMappingName { get; private set; } = null!;

        /// <summary>
        /// The connector name.
        /// </summary>
        [Output("connectorName")]
        public Output<string> ConnectorName { get; private set; } = null!;

        /// <summary>
        /// Type of connector.
        /// </summary>
        [Output("connectorType")]
        public Output<string?> ConnectorType { get; private set; } = null!;

        /// <summary>
        /// The created time.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// The DataFormat ID.
        /// </summary>
        [Output("dataFormatId")]
        public Output<string> DataFormatId { get; private set; } = null!;

        /// <summary>
        /// The description of the connector mapping.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Display name for the connector mapping.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Defines which entity type the file should map to.
        /// </summary>
        [Output("entityType")]
        public Output<string> EntityType { get; private set; } = null!;

        /// <summary>
        /// The mapping entity name.
        /// </summary>
        [Output("entityTypeName")]
        public Output<string> EntityTypeName { get; private set; } = null!;

        /// <summary>
        /// The last modified time.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// The properties of the mapping.
        /// </summary>
        [Output("mappingProperties")]
        public Output<Outputs.ConnectorMappingPropertiesResponseResult> MappingProperties { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The next run time based on customer's settings.
        /// </summary>
        [Output("nextRunTime")]
        public Output<string> NextRunTime { get; private set; } = null!;

        /// <summary>
        /// The RunId.
        /// </summary>
        [Output("runId")]
        public Output<string> RunId { get; private set; } = null!;

        /// <summary>
        /// State of connector mapping.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The hub name.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectorMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectorMapping(string name, ConnectorMappingArgs args, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/v20170101:ConnectorMapping", name, args ?? new ConnectorMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectorMapping(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/v20170101:ConnectorMapping", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:customerinsights/latest:ConnectorMapping"},
                    new Pulumi.Alias { Type = "azurerm:customerinsights/v20170426:ConnectorMapping"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConnectorMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectorMapping Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConnectorMapping(name, id, options);
        }
    }

    public sealed class ConnectorMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the connector.
        /// </summary>
        [Input("connectorName", required: true)]
        public Input<string> ConnectorName { get; set; } = null!;

        /// <summary>
        /// Type of connector.
        /// </summary>
        [Input("connectorType")]
        public Input<string>? ConnectorType { get; set; }

        /// <summary>
        /// The description of the connector mapping.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name for the connector mapping.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Defines which entity type the file should map to.
        /// </summary>
        [Input("entityType", required: true)]
        public Input<string> EntityType { get; set; } = null!;

        /// <summary>
        /// The mapping entity name.
        /// </summary>
        [Input("entityTypeName", required: true)]
        public Input<string> EntityTypeName { get; set; } = null!;

        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public Input<string> HubName { get; set; } = null!;

        /// <summary>
        /// The name of the connector mapping.
        /// </summary>
        [Input("mappingName", required: true)]
        public Input<string> MappingName { get; set; } = null!;

        /// <summary>
        /// The properties of the mapping.
        /// </summary>
        [Input("mappingProperties", required: true)]
        public Input<Inputs.ConnectorMappingPropertiesArgs> MappingProperties { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ConnectorMappingArgs()
        {
        }
    }
}
