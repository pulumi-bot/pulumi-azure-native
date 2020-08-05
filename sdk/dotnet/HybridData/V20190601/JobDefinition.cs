// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridData.V20190601
{
    /// <summary>
    /// Job Definition.
    /// </summary>
    public partial class JobDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the object.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// JobDefinition properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.JobDefinitionPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a JobDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JobDefinition(string name, JobDefinitionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:hybriddata/v20190601:JobDefinition", name, args ?? new JobDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private JobDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:hybriddata/v20190601:JobDefinition", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing JobDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JobDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new JobDefinition(name, id, options);
        }
    }

    public sealed class JobDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("customerSecrets")]
        private InputList<Inputs.CustomerSecretArgs>? _customerSecrets;

        /// <summary>
        /// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        /// </summary>
        public InputList<Inputs.CustomerSecretArgs> CustomerSecrets
        {
            get => _customerSecrets ?? (_customerSecrets = new InputList<Inputs.CustomerSecretArgs>());
            set => _customerSecrets = value;
        }

        /// <summary>
        /// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        /// </summary>
        [Input("dataManagerName", required: true)]
        public Input<string> DataManagerName { get; set; } = null!;

        [Input("dataServiceInput")]
        private InputMap<object>? _dataServiceInput;

        /// <summary>
        /// A generic json used differently by each data service type.
        /// </summary>
        public InputMap<object> DataServiceInput
        {
            get => _dataServiceInput ?? (_dataServiceInput = new InputMap<object>());
            set => _dataServiceInput = value;
        }

        /// <summary>
        /// The data service type of the job definition.
        /// </summary>
        [Input("dataServiceName", required: true)]
        public Input<string> DataServiceName { get; set; } = null!;

        /// <summary>
        /// Data Sink Id associated to the job definition.
        /// </summary>
        [Input("dataSinkId", required: true)]
        public Input<string> DataSinkId { get; set; } = null!;

        /// <summary>
        /// Data Source Id associated to the job definition.
        /// </summary>
        [Input("dataSourceId", required: true)]
        public Input<string> DataSourceId { get; set; } = null!;

        /// <summary>
        /// Last modified time of the job definition.
        /// </summary>
        [Input("lastModifiedTime")]
        public Input<string>? LastModifiedTime { get; set; }

        /// <summary>
        /// The job definition name to be created or updated.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// This is the preferred geo location for the job to run.
        /// </summary>
        [Input("runLocation")]
        public Input<string>? RunLocation { get; set; }

        [Input("schedules")]
        private InputList<Inputs.ScheduleArgs>? _schedules;

        /// <summary>
        /// Schedule for running the job definition
        /// </summary>
        public InputList<Inputs.ScheduleArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.ScheduleArgs>());
            set => _schedules = value;
        }

        /// <summary>
        /// State of the job definition.
        /// </summary>
        [Input("state", required: true)]
        public Input<string> State { get; set; } = null!;

        /// <summary>
        /// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
        /// </summary>
        [Input("userConfirmation")]
        public Input<string>? UserConfirmation { get; set; }

        public JobDefinitionArgs()
        {
        }
    }
}
