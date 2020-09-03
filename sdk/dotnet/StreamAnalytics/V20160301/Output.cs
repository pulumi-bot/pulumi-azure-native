// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StreamAnalytics.V20160301
{
    /// <summary>
    /// An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
    /// 
    /// ## Example Usage
    /// ### Create a DocumentDB output
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "Microsoft.Storage/DocumentDB",
    ///             },
    ///             JobName = "sj2331",
    ///             OutputName = "output3022",
    ///             ResourceGroupName = "sjrg7983",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a Power BI output
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "PowerBI",
    ///             },
    ///             JobName = "sj2331",
    ///             OutputName = "output3022",
    ///             ResourceGroupName = "sjrg7983",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a Service Bus Queue output with Avro serialization
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "Microsoft.ServiceBus/Queue",
    ///             },
    ///             JobName = "sj5095",
    ///             OutputName = "output3456",
    ///             ResourceGroupName = "sjrg3410",
    ///             Serialization = new AzureRM.StreamAnalytics.V20160301.Inputs.SerializationArgs
    ///             {
    ///                 Type = "Avro",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a Service Bus Topic output with CSV serialization
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "Microsoft.ServiceBus/Topic",
    ///             },
    ///             JobName = "sj7094",
    ///             OutputName = "output7886",
    ///             ResourceGroupName = "sjrg6450",
    ///             Serialization = new AzureRM.StreamAnalytics.V20160301.Inputs.SerializationArgs
    ///             {
    ///                 Type = "Csv",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a blob output with CSV serialization
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "Microsoft.Storage/Blob",
    ///             },
    ///             JobName = "sj900",
    ///             OutputName = "output1623",
    ///             ResourceGroupName = "sjrg5023",
    ///             Serialization = new AzureRM.StreamAnalytics.V20160301.Inputs.SerializationArgs
    ///             {
    ///                 Type = "Csv",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create an Azure Data Lake Store output with JSON serialization
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "Microsoft.DataLake/Accounts",
    ///             },
    ///             JobName = "sj3310",
    ///             OutputName = "output5195",
    ///             ResourceGroupName = "sjrg6912",
    ///             Serialization = new AzureRM.StreamAnalytics.V20160301.Inputs.SerializationArgs
    ///             {
    ///                 Type = "Json",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create an Azure SQL database output
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "Microsoft.Sql/Server/Database",
    ///             },
    ///             JobName = "sj6458",
    ///             OutputName = "output1755",
    ///             ResourceGroupName = "sjrg2157",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create an Azure Table output
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "Microsoft.Storage/Table",
    ///             },
    ///             JobName = "sj2790",
    ///             OutputName = "output958",
    ///             ResourceGroupName = "sjrg5176",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create an Event Hub output with JSON serialization
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new AzureRM.StreamAnalytics.V20160301.Output("output", new AzureRM.StreamAnalytics.V20160301.OutputArgs
    ///         {
    ///             Datasource = new AzureRM.StreamAnalytics.V20160301.Inputs.OutputDataSourceArgs
    ///             {
    ///                 Type = "Microsoft.ServiceBus/EventHub",
    ///             },
    ///             JobName = "sj3310",
    ///             OutputName = "output5195",
    ///             ResourceGroupName = "sjrg6912",
    ///             Serialization = new AzureRM.StreamAnalytics.V20160301.Inputs.SerializationArgs
    ///             {
    ///                 Type = "Json",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Output : Pulumi.CustomResource
    {
        /// <summary>
        /// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Output("datasource")]
        public Output<Outputs.OutputDataSourceResponseResult?> Datasource { get; private set; } = null!;

        /// <summary>
        /// Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
        /// </summary>
        [Output("diagnostics")]
        public Output<Outputs.DiagnosticsResponseResult> Diagnostics { get; private set; } = null!;

        /// <summary>
        /// The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Output("serialization")]
        public Output<Outputs.SerializationResponseResult?> Serialization { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Output resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Output(string name, OutputArgs args, CustomResourceOptions? options = null)
            : base("azurerm:streamanalytics/v20160301:Output", name, args ?? new OutputArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Output(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:streamanalytics/v20160301:Output", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:streamanalytics/latest:Output"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Output resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Output Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Output(name, id, options);
        }
    }

    public sealed class OutputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("datasource")]
        public Input<Inputs.OutputDataSourceArgs>? Datasource { get; set; }

        /// <summary>
        /// The name of the streaming job.
        /// </summary>
        [Input("jobName", required: true)]
        public Input<string> JobName { get; set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the output.
        /// </summary>
        [Input("outputName", required: true)]
        public Input<string> OutputName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("serialization")]
        public Input<Inputs.SerializationArgs>? Serialization { get; set; }

        public OutputArgs()
        {
        }
    }
}
