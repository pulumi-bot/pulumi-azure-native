// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.MachineLearning.V20160501Preview.Inputs
{

    /// <summary>
    /// Swagger 2.0 schema for a column within the data table representing a web service input or output. See Swagger specification: http://swagger.io/specification/
    /// </summary>
    public sealed class ColumnSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("enum")]
        private InputList<object>? _enum;

        /// <summary>
        /// If the data type is categorical, this provides the list of accepted categories.
        /// </summary>
        public InputList<object> Enum
        {
            get => _enum ?? (_enum = new InputList<object>());
            set => _enum = value;
        }

        /// <summary>
        /// Additional format information for the data type.
        /// </summary>
        [Input("format")]
        public InputUnion<string, Pulumi.AzureNextGen.MachineLearning.V20160501Preview.ColumnFormat>? Format { get; set; }

        /// <summary>
        /// Data type of the column.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.MachineLearning.V20160501Preview.ColumnType> Type { get; set; } = null!;

        /// <summary>
        /// Flag indicating if the type supports null values or not.
        /// </summary>
        [Input("xMsIsnullable")]
        public Input<bool>? XMsIsnullable { get; set; }

        /// <summary>
        /// Flag indicating whether the categories are treated as an ordered set or not, if this is a categorical column.
        /// </summary>
        [Input("xMsIsordered")]
        public Input<bool>? XMsIsordered { get; set; }

        public ColumnSpecificationArgs()
        {
        }
    }
}
