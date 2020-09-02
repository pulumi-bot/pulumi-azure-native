// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearning.V20160501Preview.Inputs
{

    /// <summary>
    /// The swagger 2.0 schema describing a single service input or output. See Swagger specification: http://swagger.io/specification/
    /// </summary>
    public sealed class TableSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Swagger schema description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The format, if 'type' is not 'object'
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("properties")]
        private InputMap<Inputs.ColumnSpecificationArgs>? _properties;

        /// <summary>
        /// The set of columns within the data table.
        /// </summary>
        public InputMap<Inputs.ColumnSpecificationArgs> Properties
        {
            get => _properties ?? (_properties = new InputMap<Inputs.ColumnSpecificationArgs>());
            set => _properties = value;
        }

        /// <summary>
        /// Swagger schema title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The type of the entity described in swagger.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TableSpecificationArgs()
        {
        }
    }
}
