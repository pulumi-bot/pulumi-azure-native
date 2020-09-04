// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801Preview
{
    public static class ListConnectionKeys
    {
        public static Task<ListConnectionKeysResult> InvokeAsync(ListConnectionKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListConnectionKeysResult>("azurerm:web/v20150801preview:listConnectionKeys", args ?? new ListConnectionKeysArgs(), options.WithVersion());
    }


    public sealed class ListConnectionKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The connection name.
        /// </summary>
        [Input("connectionName", required: true)]
        public string ConnectionName { get; set; } = null!;

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Input("kind")]
        public string? Kind { get; set; }

        /// <summary>
        /// Resource Location
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// time span for how long the keys will be valid
        /// </summary>
        [Input("validityTimeSpan")]
        public string? ValidityTimeSpan { get; set; }

        public ListConnectionKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListConnectionKeysResult
    {
        /// <summary>
        /// Connection Key
        /// </summary>
        public readonly string? ConnectionKey;
        /// <summary>
        /// Tokens/Claim
        /// </summary>
        public readonly ImmutableDictionary<string, ImmutableDictionary<string, object>>? ParameterValues;

        [OutputConstructor]
        private ListConnectionKeysResult(
            string? connectionKey,

            ImmutableDictionary<string, ImmutableDictionary<string, object>>? parameterValues)
        {
            ConnectionKey = connectionKey;
            ParameterValues = parameterValues;
        }
    }
}
