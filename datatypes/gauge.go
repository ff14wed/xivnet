package datatypes

// Gauge defines the data array for an add status effect block
type Gauge struct {
	ClassJob  byte
	Duration2 uint16 // In millseconds. Indicates duration of gauge 3. (Enochian)
	Duration  uint16 // In milliseconds. Indicates duration of gauge 0.
	Gauge0    int8   // Main gauge. Negative number can mean umbral ice for instance.
	Gauge1    int8   // For example, umbral hearts lives in this gauge
	Gauge2    byte   // Enochian is set via bitfields in this gauge.
	U1        uint32
	U2        uint32
}
