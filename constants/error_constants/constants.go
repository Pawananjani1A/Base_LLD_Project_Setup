package error_constants

// In House Auto Suggestion Flow Constants
const (
	SearchQueryNotPresentError      = "search key not present"
	UserIdMissingError              = "user id missing"
	IndexNameMissingError           = "index name missing"
	InvalidPageNumber               = "invalid page number: page number cannot be greater than max limit"
	EmptyResponseFromSearchPlatform = "empty response from search-platform"
	InvalidTenantHeader             = "Invalid Tenant"
)

// Category Cohort Based Ranking Experiment Constants
const (
	ExperimentConfigNotFound                          = "experiment Config not found in experiments Map"
	CategoryNotFoundError                             = "Could not find %s category in relevance info"
	CdpPropertiesFetchingError                        = "Error fetching cdp properties for the user"
	CohortNotFoundError                               = "Cohort Not found for highest confidence relevance in string format in cdp response"
	NoValidCohortExperimentsFoundForUser              = "NoValidCohortExperimentsFoundForUser"
	RankingVariantConfigNotFound                      = "Ranking variant not found in config"
	CohortNotFoundInExperimentMap                     = "Cohort type not found in experiment map"
	DataNotFoundInExperimentMap                       = "Data not found in experiment map"
	RelevanceCategoryMismatchError                    = "Relevance category not found in cdp category map"
	InvalidMarketplaceErr                             = "Invalid Marketplace type for given experiment : "
	TooManyCohortExperiments                          = "User is part of more than one experiment in marketplace : "
	RelevanceInfoNotPresent                           = "Relevance Info is not present for this query"
	UserProfileScoreNotFoundInFloatFormat             = "User profile score not found in float format for this user"
	InvalidUserProfileScore                           = "Invalid user profile score %v for this query"
	InvalidExperimentConfigForUps                     = "Invalid experiment config for triggering user profile score fallback"
	FallbackCategoryPropertiesMissingInCdpResponseErr = "Fallback Category properties not found in cdp response"
	NoValidExperimentForMarketPlaceType               = "No valid experiment found for the user in marketplace : "
	UserIdNotPresent                                  = "UserIdNotPresent"
)

// Store Product Properties Consumer Constants
const (
	ErrorPropertiesIsEmpty = "store product properties is empty"
)

// L3 Understanding errors
const (
	UserNotPartOfRequiredExperimentVariant = "User is not part of required experiment variant"
	UserNotPartOfExperiment                = "User is not part of experiment"
)

// Deboosting SecondaryStore Errors
const (
	PrimaryStoreIdMissing           = "Primary Store Id is missing"
	SearchHeaderIsEmpty             = "Search Header is empty"
	PrimaryStoreIdMissingInStoreIds = "Primary Store Id is missing in Store Ids"
	SuperStoreParamsIsEmpty         = "SuperstoreParams is Empty"
	ZeroUnserviceableStoresPresent  = "Zero Unserviceable Stores Present"
	AllStoresUnserviceable          = "Both Dark store and superstores are unserviceable"
	DynamicFlagDisabled             = "Dynamic Flag Disabled"
)

// rate limiting errors
const (
	RateLimitExceededError = "rate limit exceeded"
	RedisEvalError         = "redis eval error"
)

// VIP Deal Horizontal Card
const (
	ThemeBasedBuyingEnabled = "Theme base buying is enabled"
)

const (
	InvalidEventConfigKey              = "InvalidEventConfigKey"
	ErrorLoadingLocationZone           = "Error loading location zone"
	EventExpiredError                  = "Event expired"
	SpecialMasterDynamicFlagSetToFalse = "Special Event Dynamic Flag Set to False"
	ErrorBuildingWidget                = "Error buidling widget"
)

const ERROR = "ERROR"
