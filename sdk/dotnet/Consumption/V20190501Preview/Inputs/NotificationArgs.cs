// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Consumption.V20190501Preview.Inputs
{

    /// <summary>
    /// The notification associated with a budget.
    /// </summary>
    public sealed class NotificationArgs : Pulumi.ResourceArgs
    {
        [Input("contactEmails", required: true)]
        private InputList<string>? _contactEmails;

        /// <summary>
        /// Email addresses to send the budget notification to when the threshold is exceeded.
        /// </summary>
        public InputList<string> ContactEmails
        {
            get => _contactEmails ?? (_contactEmails = new InputList<string>());
            set => _contactEmails = value;
        }

        [Input("contactGroups")]
        private InputList<string>? _contactGroups;

        /// <summary>
        /// Action groups to send the budget notification to when the threshold is exceeded.
        /// </summary>
        public InputList<string> ContactGroups
        {
            get => _contactGroups ?? (_contactGroups = new InputList<string>());
            set => _contactGroups = value;
        }

        [Input("contactRoles")]
        private InputList<string>? _contactRoles;

        /// <summary>
        /// Contact roles to send the budget notification to when the threshold is exceeded.
        /// </summary>
        public InputList<string> ContactRoles
        {
            get => _contactRoles ?? (_contactRoles = new InputList<string>());
            set => _contactRoles = value;
        }

        /// <summary>
        /// The notification is enabled or not.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The comparison operator.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        /// <summary>
        /// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
        /// </summary>
        [Input("threshold", required: true)]
        public Input<double> Threshold { get; set; } = null!;

        public NotificationArgs()
        {
        }
    }
}
