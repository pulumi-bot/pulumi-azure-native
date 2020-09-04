// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20191201Preview.Outputs
{

    [OutputType]
    public sealed class EventRequestMessageResponseResult
    {
        /// <summary>
        /// The content of the event request message.
        /// </summary>
        public readonly Outputs.EventContentResponseResult? Content;
        /// <summary>
        /// The headers of the event request message.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Headers;
        /// <summary>
        /// The HTTP method used to send the event request message.
        /// </summary>
        public readonly string? Method;
        /// <summary>
        /// The URI used to send the event request message.
        /// </summary>
        public readonly string? RequestUri;
        /// <summary>
        /// The HTTP message version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private EventRequestMessageResponseResult(
            Outputs.EventContentResponseResult? content,

            ImmutableDictionary<string, string>? headers,

            string? method,

            string? requestUri,

            string? version)
        {
            Content = content;
            Headers = headers;
            Method = method;
            RequestUri = requestUri;
            Version = version;
        }
    }
}
