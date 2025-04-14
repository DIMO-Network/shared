package logfields

const (
	FunctionName = "func"
	// VehicleTokenID to track the vehicle's token id. type: uint64
	VehicleTokenID = "vehicleTokenId"
	// AftermarketTokenID to track the aftermarket device token id. Type: uint64.
	AftermarketTokenID = "aftermarketTokenId"
	VIN                = "vin"
	WMI                = "wmi"
	// Payload can be used to track different payloads from stream or requests/responses. Type: string or raw json
	Payload      = "payload"
	CountryCode  = "countryCode"
	DefinitionID = "definitionId"
)

const (
	// CloudEventID unique id for the event. Type: string
	CloudEventID       = "cloudeventId"
	CloudEventSource   = "cloudeventSource"
	CloudEventType     = "cloudeventType"
	CloudEventSubject  = "cloudeventSubject"
	CloudEventProducer = "cloudeventProducer"
)

const (
	GRPCStatusCode = "grpcStatusCode"
	GRPCMethod     = "grpcMethod"
)
