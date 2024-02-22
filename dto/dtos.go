package dto

import "time"

// Identifier struct with detailed fields
type Identifier struct {
	ID        string          `json:"id"`        // Unique identifier for the instance
	Extension interface{}     `json:"extension"` // from Element: extension
	Use       string          `json:"use"`       // usual | official | temp | secondary | old (If known)
	Type      CodeableConcept `json:"type"`      // Description of identifier
	System    string          `json:"system"`    // The namespace for the identifier value
	Value     string          `json:"value"`     // The value that is unique
	Period    Period          `json:"period"`    // Time period when id is/was valid for use
	Assigner  *Reference      `json:"assigner"`  // Organization that issued id (may be just text)
}

// CodeableConcept struct with detailed fields
type CodeableConcept struct {
	ID        string      `json:"id"`        // Unique identifier for the instance
	Extension interface{} `json:"extension"` // from Element: extension
	Coding    []Coding    `json:"coding"`    // Code defined by a terminology system
	Text      string      `json:"text"`      // Plain text representation of the concept
}

// Coding struct with detailed fields
type Coding struct {
	ID           string      `json:"id"`           // Unique identifier for the instance
	Extension    interface{} `json:"extension"`    // from Element: extension
	System       string      `json:"system"`       // Identity of the terminology system
	Version      string      `json:"version"`      // Version of the system - if relevant
	Code         string      `json:"code"`         // Symbol in syntax defined by the system
	Display      string      `json:"display"`      // Representation defined by the system
	UserSelected bool        `json:"userSelected"` // If this coding was chosen directly by the user
}

// Quantity struct with detailed fields
type Quantity struct {
	ID     string  `json:"id"`     // Unique identifier for the instance
	Value  float64 `json:"value"`  // Numerical value (with implicit precision)
	Unit   string  `json:"unit"`   // Unit representation
	System string  `json:"system"` // System that defines the unit
	Code   string  `json:"code"`   // Coded form of the unit
}

// Duration struct updated to match Quantity fields
type Duration struct {
	ID     string  `json:"id"`     // Unique identifier for the instance
	Value  float64 `json:"value"`  // Duration length (with implicit precision)
	Unit   string  `json:"unit"`   // Unit representation (s, min, h, d, etc.)
	System string  `json:"system"` // System that defines the unit (UCUM)
	Code   string  `json:"code"`   // Coded form of the unit
}

// Range struct with detailed fields and added ID
type Range struct {
	ID   string   `json:"id"`   // Unique identifier for the instance
	Low  Quantity `json:"low"`  // Low limit
	High Quantity `json:"high"` // High limit
}

// Reference struct with detailed fields
type Reference struct {
	ID         string      `json:"id"`         // Unique identifier for the instance
	Extension  interface{} `json:"extension"`  // from Element: extension
	Reference  string      `json:"reference"`  // Literal reference, Relative, internal or absolute URL
	Type       string      `json:"type"`       // Type the reference refers to (e.g., "Patient")
	Identifier Identifier  `json:"identifier"` // Logical reference, when literal reference is not known
	Display    string      `json:"display"`    // Text alternative for the resource
}

// Period struct represents a time span
type Period struct {
	ID    string    `json:"id"`    // Unique identifier for the instance
	Start time.Time `json:"start"` // Start of the period
	End   time.Time `json:"end"`   // End of the period
}

// Timing struct with detailed fields
type Timing struct {
	ID     string          `json:"id"`     // Unique identifier for the instance
	Event  []time.Time     `json:"event"`  // When the event occurs
	Repeat Repeat          `json:"repeat"` // When the event is to occur
	Code   CodeableConcept `json:"code"`   // BID | TID | QID | AM | PM | QD | QOD | +
}

type Repeat struct {
	BoundsDuration Duration `json:"boundsDuration"`
	BoundsRange    Range    `json:"boundsRange"`
	BoundsPeriod   Period   `json:"boundsPeriod"`
	Count          int      `json:"count"`        // Number of times to repeat
	CountMax       int      `json:"countMax"`     // Maximum number of times to repeat
	Duration       float64  `json:"duration"`     // How long when it happens
	DurationMax    float64  `json:"durationMax"`  // How long when it happens (Max)
	DurationUnit   string   `json:"durationUnit"` // Unit of time (UCUM)
	Frequency      int      `json:"frequency"`    // Event occurs frequency times per period
	FrequencyMax   int      `json:"frequencyMax"` // Event occurs up to frequencyMax times per period
	Period         float64  `json:"period"`       // Event occurs frequency times per period
	PeriodMax      float64  `json:"periodMax"`    // Upper limit of period
	PeriodUnit     string   `json:"periodUnit"`   // Unit of time (UCUM)
	DayOfWeek      []string `json:"dayOfWeek"`    // Days of week
	TimeOfDay      []string `json:"timeOfDay"`    // Time of day for action
	When           []string `json:"when"`         // Code for time period of occurrence
	Offset         int      `json:"offset"`       // Minutes from event (before or after)
}

// Annotation struct with detailed fields
type Annotation struct {
	ID              string    `json:"id"`              // Unique identifier for the instance
	AuthorReference Reference `json:"authorReference"` // Practitioner|Patient|RelatedPerson|Organization
	AuthorString    string    `json:"authorString"`
	Time            time.Time `json:"time"` // When the annotation was made
	Text            string    `json:"text"` // The annotation - text content (as markdown)
}

// Dosage struct with detailed fields
type Dosage struct {
	ID                       string            `json:"id"`                    // Unique identifier for the instance
	Sequence                 int               `json:"sequence"`              // The order of the dosage instructions
	Text                     string            `json:"text"`                  // Free text dosage instructions e.g. SIG
	AdditionalInstruction    []CodeableConcept `json:"additionalInstruction"` // Supplemental instruction - e.g. "with meals"
	PatientInstruction       string            `json:"patientInstruction"`    // Patient or consumer oriented instructions
	Timing                   Timing            `json:"timing"`                // When medication should be administered
	AsNeededBoolean          bool              `json:"asNeededBoolean"`
	AsNeededCodeableConcept  CodeableConcept   `json:"asNeededCodeableConcept"`
	Site                     CodeableConcept   `json:"site"`                     // Body site to administer to
	Route                    CodeableConcept   `json:"route"`                    // How drug should enter body
	Method                   CodeableConcept   `json:"method"`                   // Technique for administering medication
	DoseAndRate              []DoseAndRate     `json:"doseAndRate"`              // Amount of medication administered
	MaxDosePerPeriod         Ratio             `json:"maxDosePerPeriod"`         // Upper limit on medication per unit of time
	MaxDosePerAdministration Quantity          `json:"maxDosePerAdministration"` // Upper limit on medication per administration
	MaxDosePerLifetime       Quantity          `json:"maxDosePerLifetime"`       // Upper limit on medication per lifetime of the patient
}

type DoseAndRate struct {
	Type         CodeableConcept `json:"type"` // The kind of dose or rate specified
	DoseRange    Range           `json:"doseRange"`
	DoseQuantity Quantity        `json:"doseQuantity"`
	RateRatio    Ratio           `json:"rateRatio"`
	RateRange    Range           `json:"rateRange"`
	RateQuantity Quantity        `json:"rateQuantity"`
}

type Ratio struct {
	Numerator   Quantity `json:"numerator"`
	Denominator Quantity `json:"denominator"`
}

type MedicationAdministration struct {
	ID                        string            `json:"id"`           // Unique identifier
	ResourceType              string            `json:"resourceType"` // Type of resource
	Identifier                []Identifier      `json:"identifier"`   // External identifier
	Instantiates              []string          `json:"instantiates"` // Instantiates protocol or definition
	PartOf                    []Reference       `json:"partOf"`       // Part of referenced event
	Status                    string            `json:"status"`       // in-progress | not-done | on-hold | completed | entered-in-error | stopped | unknown
	StatusReason              []CodeableConcept `json:"statusReason"` // Reason administration not performed
	Category                  CodeableConcept   `json:"category"`     // Type of medication usage
	MedicationCodeableConcept CodeableConcept   `json:"medicationCodeableConcept"`
	MedicationReference       Reference         `json:"medicationReference"`
	Subject                   Reference         `json:"subject"`               // Who received medication
	Context                   Reference         `json:"context"`               // Encounter or Episode of Care administered as part of
	SupportingInformation     []Reference       `json:"supportingInformation"` // Additional information to support administration
	EffectiveDateTime         string            `json:"effectiveDateTime"`
	EffectivePeriod           Period            `json:"effectivePeriod"`
	Performer                 []Performer       `json:"performer"`
	ReasonCode                []CodeableConcept `json:"reasonCode"`      // Reason administration performed
	ReasonReference           []Reference       `json:"reasonReference"` // Condition or observation that supports why
	Request                   Reference         `json:"request"`         // Request administration performed against
	Device                    []Reference       `json:"device"`          // Device used to administer
	Note                      []Annotation      `json:"note"`            // Information about the administration
	Dosage                    Dosage            `json:"dosage"`          // Details of how medication was taken
	EventHistory              []Reference       `json:"eventHistory"`    // A list of events of interest in the lifecycle
}

type Performer struct {
	Function CodeableConcept `json:"function"` // Type of performance
	Actor    Reference       `json:"actor"`    // Who performed the medication administration
}

type List struct {
	ID           string          `json:"id"`           // Unique identifier
	ResourceType string          `json:"resourceType"` // Type of resource
	Identifier   []Identifier    `json:"identifier"`   // Business identifier
	Status       string          `json:"status"`       // current | retired | entered-in-error
	Mode         string          `json:"mode"`         // working | snapshot | changes
	Title        string          `json:"title"`        // Descriptive name for the list
	Code         CodeableConcept `json:"code"`         // What the purpose of this list is
	Subject      Reference       `json:"subject"`      // If all resources have the same subject
	Encounter    Reference       `json:"encounter"`    // Context in which list created
	Date         string          `json:"date"`         // When the list was prepared
	Source       Reference       `json:"source"`       // Who and/or what defined the list contents (aka Author)
	OrderedBy    CodeableConcept `json:"orderedBy"`    // What order the list has
	Note         []Annotation    `json:"note"`         // Comments about the list
	Entry        []Entry         `json:"entry"`        // Entries in the list
	EmptyReason  CodeableConcept `json:"emptyReason"`  // Why list is empty
}

type Entry struct {
	Flag    CodeableConcept `json:"flag"`    // Status/Workflow information about this item
	Deleted bool            `json:"deleted"` // If this item is actually marked as deleted
	Date    string          `json:"date"`    // When item added to list
	Item    Reference       `json:"item"`    // Actual entry
}

type MedicationStatement struct {
	ID                        string            `json:"id"`           // Unique identifier
	ResourceType              string            `json:"resourceType"` // Type of resource
	Identifier                []Identifier      `json:"identifier"`   // External identifier
	BasedOn                   []Reference       `json:"basedOn"`      // Fulfils plan, proposal or order
	PartOf                    []Reference       `json:"partOf"`       // Part of referenced event
	Status                    string            `json:"status"`       // active | completed | entered-in-error | intended | stopped | on-hold | unknown | not-taken
	StatusReason              []CodeableConcept `json:"statusReason"` // Reason for current status
	Category                  CodeableConcept   `json:"category"`     // Type of medication usage
	MedicationCodeableConcept CodeableConcept   `json:"medicationCodeableConcept"`
	MedicationReference       Reference         `json:"medicationReference"`
	Subject                   Reference         `json:"subject"` // Who is/was taking the medication
	Context                   Reference         `json:"context"` // Encounter / Episode associated with MedicationStatement
	EffectiveDateTime         string            `json:"effectiveDateTime"`
	EffectivePeriod           Period            `json:"effectivePeriod"`
	DateAsserted              string            `json:"dateAsserted"`      // When the statement was asserted?
	InformationSource         Reference         `json:"informationSource"` // Person or organization that provided the information
	DerivedFrom               []Reference       `json:"derivedFrom"`       // Additional supporting information
	ReasonCode                []CodeableConcept `json:"reasonCode"`        // Reason for why the medication is being/was taken
	ReasonReference           []Reference       `json:"reasonReference"`   // Condition or observation that supports why
	Note                      []Annotation      `json:"note"`              // Further information about the statement
	Dosage                    []Dosage          `json:"dosage"`            // Details of how medication is/was taken or should be taken
}
