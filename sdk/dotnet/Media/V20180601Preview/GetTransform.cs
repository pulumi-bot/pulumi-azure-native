// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview
{
    public static class GetTransform
    {
        public static Task<GetTransformResult> InvokeAsync(GetTransformArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransformResult>("azurerm:media/v20180601preview:getTransform", args ?? new GetTransformArgs(), options.WithVersion());
    }


    public sealed class GetTransformArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

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

        public GetTransformArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTransformResult
    {
        /// <summary>
        /// The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// An optional verbose description of the Transform.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An array of one or more TransformOutputs that the Transform should generate.
        /// </summary>
        public readonly ImmutableArray<Outputs.TransformOutputResponseResult> Outputs;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTransformResult(
            string created,

            string? description,

            string lastModified,

            string name,

            ImmutableArray<Outputs.TransformOutputResponseResult> outputs,

            string type)
        {
            Created = created;
            Description = description;
            LastModified = lastModified;
            Name = name;
            Outputs = outputs;
            Type = type;
        }
    }
}
