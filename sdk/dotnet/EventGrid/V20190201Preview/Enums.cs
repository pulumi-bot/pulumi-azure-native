// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.EventGrid.V20190201Preview
{
    /// <summary>
    /// The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
    /// </summary>
    [EnumType]
    public readonly struct AdvancedFilterOperatorType : IEquatable<AdvancedFilterOperatorType>
    {
        private readonly string _value;

        private AdvancedFilterOperatorType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AdvancedFilterOperatorType NumberIn { get; } = new AdvancedFilterOperatorType("NumberIn");
        public static AdvancedFilterOperatorType NumberNotIn { get; } = new AdvancedFilterOperatorType("NumberNotIn");
        public static AdvancedFilterOperatorType NumberLessThan { get; } = new AdvancedFilterOperatorType("NumberLessThan");
        public static AdvancedFilterOperatorType NumberGreaterThan { get; } = new AdvancedFilterOperatorType("NumberGreaterThan");
        public static AdvancedFilterOperatorType NumberLessThanOrEquals { get; } = new AdvancedFilterOperatorType("NumberLessThanOrEquals");
        public static AdvancedFilterOperatorType NumberGreaterThanOrEquals { get; } = new AdvancedFilterOperatorType("NumberGreaterThanOrEquals");
        public static AdvancedFilterOperatorType BoolEquals { get; } = new AdvancedFilterOperatorType("BoolEquals");
        public static AdvancedFilterOperatorType StringIn { get; } = new AdvancedFilterOperatorType("StringIn");
        public static AdvancedFilterOperatorType StringNotIn { get; } = new AdvancedFilterOperatorType("StringNotIn");
        public static AdvancedFilterOperatorType StringBeginsWith { get; } = new AdvancedFilterOperatorType("StringBeginsWith");
        public static AdvancedFilterOperatorType StringEndsWith { get; } = new AdvancedFilterOperatorType("StringEndsWith");
        public static AdvancedFilterOperatorType StringContains { get; } = new AdvancedFilterOperatorType("StringContains");

        public static bool operator ==(AdvancedFilterOperatorType left, AdvancedFilterOperatorType right) => left.Equals(right);
        public static bool operator !=(AdvancedFilterOperatorType left, AdvancedFilterOperatorType right) => !left.Equals(right);

        public static explicit operator string(AdvancedFilterOperatorType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AdvancedFilterOperatorType other && Equals(other);
        public bool Equals(AdvancedFilterOperatorType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the endpoint for the dead letter destination
    /// </summary>
    [EnumType]
    public readonly struct DeadLetterEndPointType : IEquatable<DeadLetterEndPointType>
    {
        private readonly string _value;

        private DeadLetterEndPointType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeadLetterEndPointType StorageBlob { get; } = new DeadLetterEndPointType("StorageBlob");

        public static bool operator ==(DeadLetterEndPointType left, DeadLetterEndPointType right) => left.Equals(right);
        public static bool operator !=(DeadLetterEndPointType left, DeadLetterEndPointType right) => !left.Equals(right);

        public static explicit operator string(DeadLetterEndPointType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeadLetterEndPointType other && Equals(other);
        public bool Equals(DeadLetterEndPointType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the endpoint for the event subscription destination
    /// </summary>
    [EnumType]
    public readonly struct EndpointType : IEquatable<EndpointType>
    {
        private readonly string _value;

        private EndpointType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EndpointType WebHook { get; } = new EndpointType("WebHook");
        public static EndpointType EventHub { get; } = new EndpointType("EventHub");
        public static EndpointType StorageQueue { get; } = new EndpointType("StorageQueue");
        public static EndpointType HybridConnection { get; } = new EndpointType("HybridConnection");
        public static EndpointType ServiceBusQueue { get; } = new EndpointType("ServiceBusQueue");

        public static bool operator ==(EndpointType left, EndpointType right) => left.Equals(right);
        public static bool operator !=(EndpointType left, EndpointType right) => !left.Equals(right);

        public static explicit operator string(EndpointType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EndpointType other && Equals(other);
        public bool Equals(EndpointType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The event delivery schema for the event subscription.
    /// </summary>
    [EnumType]
    public readonly struct EventDeliverySchema : IEquatable<EventDeliverySchema>
    {
        private readonly string _value;

        private EventDeliverySchema(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EventDeliverySchema EventGridSchema { get; } = new EventDeliverySchema("EventGridSchema");
        public static EventDeliverySchema CloudEventV01Schema { get; } = new EventDeliverySchema("CloudEventV01Schema");
        public static EventDeliverySchema CustomInputSchema { get; } = new EventDeliverySchema("CustomInputSchema");

        public static bool operator ==(EventDeliverySchema left, EventDeliverySchema right) => left.Equals(right);
        public static bool operator !=(EventDeliverySchema left, EventDeliverySchema right) => !left.Equals(right);

        public static explicit operator string(EventDeliverySchema value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EventDeliverySchema other && Equals(other);
        public bool Equals(EventDeliverySchema other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This determines the format that Event Grid should expect for incoming events published to the topic.
    /// </summary>
    [EnumType]
    public readonly struct InputSchema : IEquatable<InputSchema>
    {
        private readonly string _value;

        private InputSchema(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InputSchema EventGridSchema { get; } = new InputSchema("EventGridSchema");
        public static InputSchema CustomEventSchema { get; } = new InputSchema("CustomEventSchema");
        public static InputSchema CloudEventV01Schema { get; } = new InputSchema("CloudEventV01Schema");

        public static bool operator ==(InputSchema left, InputSchema right) => left.Equals(right);
        public static bool operator !=(InputSchema left, InputSchema right) => !left.Equals(right);

        public static explicit operator string(InputSchema value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InputSchema other && Equals(other);
        public bool Equals(InputSchema other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the custom mapping
    /// </summary>
    [EnumType]
    public readonly struct InputSchemaMappingType : IEquatable<InputSchemaMappingType>
    {
        private readonly string _value;

        private InputSchemaMappingType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InputSchemaMappingType Json { get; } = new InputSchemaMappingType("Json");

        public static bool operator ==(InputSchemaMappingType left, InputSchemaMappingType right) => left.Equals(right);
        public static bool operator !=(InputSchemaMappingType left, InputSchemaMappingType right) => !left.Equals(right);

        public static explicit operator string(InputSchemaMappingType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InputSchemaMappingType other && Equals(other);
        public bool Equals(InputSchemaMappingType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
