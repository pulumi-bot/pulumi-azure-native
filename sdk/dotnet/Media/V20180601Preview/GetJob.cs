// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview
{
    public static class GetJob
    {
        public static Task<GetJobResult> InvokeAsync(GetJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("azurerm:media/v20180601preview:getJob", args ?? new GetJobArgs(), options.WithVersion());
    }


    public sealed class GetJobArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The Job name.
        /// </summary>
        [Input("jobName", required: true)]
        public string JobName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Transform name.
        /// </summary>
        [Input("transformName", required: true)]
        public string TransformName { get; set; } = null!;

        public GetJobArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobResult
    {
        /// <summary>
        /// Customer provided correlation data that will be returned in Job completed events.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CorrelationData;
        /// <summary>
        /// The UTC date and time when the Job was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// Optional customer supplied description of the Job.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The inputs for the Job.
        /// </summary>
        public readonly Outputs.JobInputResponseResult Input;
        /// <summary>
        /// The UTC date and time when the Job was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The outputs for the Job.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobOutputResponseResult> Outputs;
        /// <summary>
        /// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
        /// </summary>
        public readonly string? Priority;
        /// <summary>
        /// The current state of the job.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetJobResult(
            ImmutableDictionary<string, string>? correlationData,

            string created,

            string? description,

            Outputs.JobInputResponseResult input,

            string lastModified,

            string name,

            ImmutableArray<Outputs.JobOutputResponseResult> outputs,

            string? priority,

            string state,

            string type)
        {
            CorrelationData = correlationData;
            Created = created;
            Description = description;
            Input = input;
            LastModified = lastModified;
            Name = name;
            Outputs = outputs;
            Priority = priority;
            State = state;
            Type = type;
        }
    }
}
