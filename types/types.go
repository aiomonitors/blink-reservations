package types

type AvailEvents struct {
	Success bool   `json:"success"`
	Slots   []Slot `json:"slots"`  
}

type Slot struct {
	ID              int64  `json:"id"`             
	EventInstanceID int64  `json:"eventInstanceId"`
	EventID         int64  `json:"eventId"`        
	MosoID          string `json:"mosoId"`         
	Date            int64  `json:"date"`           
	StartDateUTC    string `json:"startDateUtc"`   
	EndDateUTC      string `json:"endDateUtc"`     
	StartTime       string `json:"startTime"`      
	EndTime         string `json:"endTime"`        
	RemainingSpots  int64  `json:"remainingSpots"` 
	TTL             int64  `json:"ttl"`            
}

// Refresh Response stuff
type RefreshResponse struct {
	ExternalPages          []ExternalPage        `json:"externalPages"`         
	Features               []FeatureElement      `json:"features"`              
	HasCompletedIntake     bool                  `json:"hasCompletedIntake"`    
	HasGatedAccess         HasGatedAccess        `json:"hasGatedAccess"`        
	HasSeenBookmarkTooltip bool                  `json:"hasSeenBookmarkTooltip"`
	HasSeenOnBoarding      bool                  `json:"hasSeenOnBoarding"`     
	Meta                   Meta                  `json:"meta"`                  
	Offers                 []Offer               `json:"offers"`                
	Pillars                []string              `json:"pillars"`               
	RefreshToken           string                `json:"refreshToken"`          
	Token                  string                `json:"token"`                 
	UserAccess             string                `json:"userAccess"`            
	UUID                   string                `json:"uuid"`                  
	Member                 Member                `json:"member"`                
	MembershipTypeDisplay  MembershipTypeDisplay `json:"membershipTypeDisplay"` 
	TermsAgreed            bool                  `json:"termsAgreed"`           
}

type ExternalPage struct {
	Title    string `json:"title"`   
	LinkID   string `json:"linkId"`  
	Priority string `json:"priority"`
	Link     string `json:"link"`    
}

type FeatureElement struct {
	FeatureID          string      `json:"featureId"`                   
	FeatureName        string      `json:"featureName"`                 
	FeatureType        FeatureType `json:"featureType"`                 
	Order              *int64      `json:"order,omitempty"`             
	DisplayInMenu      *bool       `json:"displayInMenu,omitempty"`     
	EnabledMemberTypes []Label     `json:"enabledMemberTypes"`          
	FeatureDescription *string     `json:"featureDescription,omitempty"`
	EnabledLocations   []string    `json:"enabledLocations"`            
}

type HasGatedAccess struct {
	TrialEnabled bool `json:"trialEnabled"`
	TrialExpired bool `json:"trialExpired"`
}

type Member struct {
	RoleState              string                `json:"RoleState"`             
	Interests              interface{}           `json:"Interests"`             
	Name                   string                `json:"Name"`                  
	Addresses              []Address             `json:"Addresses"`             
	AccountStatus          string                `json:"AccountStatus"`         
	TaxExemptID            int64                 `json:"TaxExemptId"`           
	MembershipAgreements   []MembershipAgreement `json:"MembershipAgreements"`  
	MemberSince            string                `json:"MemberSince"`           
	ClientAccountIDS       []int64               `json:"ClientAccountIds"`      
	Cards                  []Card                `json:"Cards"`                 
	Status                 string                `json:"Status"`                
	EnforceRequiredFields  bool                  `json:"EnforceRequiredFields"` 
	PhoneNumbers           []PhoneNumber         `json:"PhoneNumbers"`          
	ProfilePicImageBase64  interface{}           `json:"ProfilePicImageBase64"` 
	ReferredByRoleID       interface{}           `json:"ReferredByRoleId"`      
	CreditLimit            int64                 `json:"CreditLimit"`           
	RoleID                 string                `json:"RoleId"`                
	PartyID                int64                 `json:"PartyId"`               
	Characteristics        []Characteristic      `json:"Characteristics"`       
	ProfilePicURL          string                `json:"ProfilePicUrl"`         
	TaxExemptionName       string                `json:"TaxExemptionName"`      
	EmailAddresses         []EmailAddress        `json:"EmailAddresses"`        
	Balance                int64                 `json:"Balance"`               
	DefaultClientAccountID int64                 `json:"DefaultClientAccountId"`
	Location               Location              `json:"Location"`              
}

type Address struct {
	Status      []Status      `json:"Status"`     
	TypeID      int64         `json:"TypeId"`     
	Type        string        `json:"Type"`       
	Address2    interface{}   `json:"Address2"`   
	StateCode   string        `json:"StateCode"`  
	Address3    interface{}   `json:"Address3"`   
	PostalCode  string        `json:"PostalCode"` 
	Address1    string        `json:"Address1"`   
	City        string        `json:"City"`       
	Purposes    []interface{} `json:"Purposes"`   
	CountryCode string        `json:"CountryCode"`
}

type Status struct {
	Status bool   `json:"Status"`
	Name   string `json:"Name"`  
}

type Card struct {
	Status string `json:"Status"`
	CardID string `json:"CardId"`
	ID     int64  `json:"Id"`    
}

type Characteristic struct {
	CharacteristicTypeID   int64  `json:"CharacteristicTypeId"`  
	Value                  string `json:"Value"`                 
	CharacteristicTypeName string `json:"CharacteristicTypeName"`
	CharacteristicValueID  *int64 `json:"CharacteristicValueId"` 
}

type EmailAddress struct {
	Status       []Status  `json:"Status"`      
	TypeID       int64     `json:"TypeId"`      
	Type         string    `json:"Type"`        
	Purposes     []Purpose `json:"Purposes"`    
	EmailAddress string    `json:"EmailAddress"`
}

type Purpose struct {
	Status               bool   `json:"Status"`              
	ContactPurposeTypeID int64  `json:"ContactPurposeTypeId"`
	Name                 string `json:"Name"`                
}

type Location struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

type MembershipAgreement struct {
	MemberAgreementID     int64       `json:"MemberAgreementId"`    
	EditableStartDate     string      `json:"EditableStartDate"`    
	MemberAgreementStatus string      `json:"MemberAgreementStatus"`
	AgreementID           int64       `json:"AgreementId"`          
	CancellationDate      interface{} `json:"CancellationDate"`     
	AgreementName         Label       `json:"AgreementName"`        
	ObligationDate        string      `json:"ObligationDate"`       
	Location              Location    `json:"Location"`             
}

type PhoneNumber struct {
	AdditionalInformation string    `json:"AdditionalInformation"`
	Status                []Status  `json:"Status"`               
	TypeID                int64     `json:"TypeId"`               
	Type                  string    `json:"Type"`                 
	PhoneNumber           string    `json:"PhoneNumber"`          
	TextMessageOk         bool      `json:"TextMessageOk"`        
	Purposes              []Purpose `json:"Purposes"`             
	CountryCode           string    `json:"CountryCode"`          
}

type MembershipTypeDisplay struct {
	Label          Label  `json:"label"`         
	PrimaryColor   string `json:"primaryColor"`  
	SecondaryColor string `json:"secondaryColor"`
}

type Meta struct {
	HasSeenOnboarding                       string `json:"hasSeenOnboarding"`                        
	HasSeenFeatureHighlightPersonalTraining string `json:"hasSeenFeatureHighlight-personal-training"`
	HasSeenFeatureHighlightBookPtQueens2    string `json:"hasSeenFeatureHighlight-book-pt-queens-2"` 
	DeclinedPostLoginIntake                 string `json:"declinedPostLoginIntake"`                  
	HasSeenFeatureHighlightBookPtQueens     string `json:"hasSeenFeatureHighlight-book-pt-queens"`   
	LastSeenFeatureHighlightTimestamp       string `json:"lastSeenFeatureHighlightTimestamp"`        
}

type Offer struct {
	AppendUserID bool   `json:"appendUserId"`
	CtaCopy      string `json:"ctaCopy"`     
	CtaLink      string `json:"ctaLink"`     
	Offer        string `json:"offer"`       
	PartnerID    string `json:"partnerId"`   
}

type Label string
const (
	Blue Label = "Blue"
	Gray Label = "Gray"
	Green Label = "Green"
)

type FeatureType string
const (
	Feature FeatureType = "Feature"
)

// Member Reservation Stuff
type MemberReservationResponse struct {
	Success       bool           `json:"success"`
	Registrations []Registration `json:"registrations"`
}

type Registration struct {
	EventRegistrationInstanceID int64       `json:"EventRegistrationInstanceId"`
	EventInstanceID             int64       `json:"EventInstanceId"`
	EventID                     int64       `json:"EventId"`
	EventDescription            string      `json:"EventDescription"`
	EventCode                   string      `json:"EventCode"`
	CategoryName                string      `json:"CategoryName"`
	CategoryID                  int64       `json:"CategoryId"`
	ShortDescription            interface{} `json:"ShortDescription"`
	LongDescription             interface{} `json:"LongDescription"`
	TermsAndConditions          interface{} `json:"TermsAndConditions"`
	ThumbnailImage              interface{} `json:"ThumbnailImage"`
	FullSizeImage               interface{} `json:"FullSizeImage"`
	LocationID                  int64       `json:"LocationId"`
	StartDateUTC                string      `json:"StartDateUtc"`
	EndDateUTC                  string      `json:"EndDateUtc"`
	StartDate                   string      `json:"StartDate"`
	EndDate                     string      `json:"EndDate"`
	StartTime                   string      `json:"StartTime"`
	EndTime                     string      `json:"EndTime"`
	CreatedDateUTC              string      `json:"CreatedDateUtc"`
	CreatedDate                 string      `json:"CreatedDate"`
	BusinessUnitCode            string      `json:"BusinessUnitCode"`
	InstanceType                string      `json:"InstanceType"`
	PaymentMethodID             int64       `json:"PaymentMethodId"`
	PaymentStatusID             int64       `json:"PaymentStatusId"`
}
