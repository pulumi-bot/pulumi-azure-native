// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CustomerInsights.V20170101.Inputs
{

    /// <summary>
    /// Connector mapping property format.
    /// </summary>
    public sealed class ConnectorMappingFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The oData language.
        /// </summary>
        [Input("acceptLanguage")]
        public Input<string>? AcceptLanguage { get; set; }

        /// <summary>
        /// Character separating array elements.
        /// </summary>
        [Input("arraySeparator")]
        public Input<string>? ArraySeparator { get; set; }

        /// <summary>
        /// The character that signifies a break between columns.
        /// </summary>
        [Input("columnDelimiter")]
        public Input<string>? ColumnDelimiter { get; set; }

        /// <summary>
        /// The type mapping format.
        /// </summary>
        [Input("formatType", required: true)]
        public Input<Pulumi.AzureNextGen.CustomerInsights.V20170101.FormatTypes> FormatType { get; set; } = null!;

        /// <summary>
        /// Quote character, used to indicate enquoted fields.
        /// </summary>
        [Input("quoteCharacter")]
        public Input<string>? QuoteCharacter { get; set; }

        /// <summary>
        /// Escape character for quotes, can be the same as the quoteCharacter.
        /// </summary>
        [Input("quoteEscapeCharacter")]
        public Input<string>? QuoteEscapeCharacter { get; set; }

        public ConnectorMappingFormatArgs()
        {
        }
    }
}
