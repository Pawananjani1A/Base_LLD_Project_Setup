package api_constants

// http stats logger constants
const (
	RequestUrl                = "RequestUrl"
	RequestMethod             = "RequestMethod"
	RemoteAddr                = "RemoteAddr"
	RequestId                 = "RequestId"
	DeltaTime                 = "DeltaTime"
	RequestTime               = "RequestTime"
	ResponseTime              = "ResponseTime"
	UserId                    = "UserId"
	SearchTerm                = "SearchTerm"
	QCLSearchTerm             = "QCLSearchTerm"
	StoreId                   = "StoreId"
	StoreProductId            = "StoreProductId"
	SearchKey                 = "SearchKey"
	Category                  = "Category"
	Subcategory               = "Subcategory"
	CityId                    = "CityId"
	CorrectedAtlasQueryTerm   = "queryString"
	RequestTimedout           = "requestTimedout"
	QclSearchTerm             = "QclSearchTerm"
	Source                    = "Source"
	Mode                      = "Mode"
	PageNumber                = "PageNumber"
	EmailId                   = "EmailId"
	Tenant                    = "Tenant"
	AppVersion                = "AppVersion"
	BundleVersion             = "BundleVersion"
	DeviceId                  = "DeviceId"
	Platform                  = "Platform"
	SearchMode                = "SearchMode"
	MarketPlaceType           = "MarketPlaceType"
	UserSessionId             = "UserSessionId"
	BatchIndex                = "BatchIndex"
	MemoryProfiling           = "MemoryProfiling"
	CvrAutosuggestionsCount   = "CvrAutosuggestionsCount"
	CtrAutosuggestionsCount   = "CtrAutosuggestionsCount"
	TotalAutosuggestionsCount = "TotalAutosuggestionsCount"
	NeuronModelName           = "NeuronModelName"
	CvrErrorMessage           = "CvrErrorMessage"
	CtrErrorMessage           = "CtrErrorMessage"
	ProductVariantId          = "ProductVariantId"
)

// search api params
const (
	Query            = "query"
	RPI              = "RPI"
	ATCPI            = "ATCPI"
	VDR              = "VDR"
	SSFR             = "SSFR"
	Rank5Variant     = "Rank5"
	Rank6Variant     = "Rank6"
	Rank7Variant     = "Rank7"
	Rank8Variant     = "Rank8"
	Rank9Variant     = "Rank9"
	Rank10Variant    = "Rank10"
	Rank11Variant    = "Rank11"
	Rank12Variant    = "Rank12"
	Rank13Variant    = "Rank13"
	Rank14Variant    = "Rank14"
	Rank15Variant    = "Rank15"
	Rank16Variant    = "Rank16"
	Rank17Variant    = "Rank17"
	Rank18Variant    = "Rank18"
	Rank19Variant    = "Rank19"
	Rank20Variant    = "Rank20"
	Rank21Variant    = "Rank21"
	Rank22Variant    = "Rank22"
	Rank23Variant    = "Rank23"
	Rank24Variant    = "Rank24"
	Rank25Variant    = "Rank25"
	Rank26Variant    = "Rank26"
	Rank27Variant    = "Rank27"
	Rank28Variant    = "Rank28"
	Rank29Variant    = "Rank29"
	Rank30Variant    = "Rank30"
	Rank31Variant    = "Rank31"
	Rank32Variant    = "Rank32"
	Rank33Variant    = "Rank33"
	Rank34Variant    = "Rank34"
	Rank35Variant    = "Rank35"
	Rank36Variant    = "Rank36"
	Rank37Variant    = "Rank37"
	Rank38Variant    = "Rank38"
	Rank39Variant    = "Rank39"
	Rank40Variant    = "Rank40"
	SortBy           = "sortBy"
	SortByValue      = "sortByValue"
	EntityType       = "entityType"
	EntityId         = "entityId"
	EntitySearchTerm = "searchTerm"
)

// consumer flow constants
const (
	PvL4V2Flow       = "pv_l4_consumer_v2_flow"
	PvL4V2BundleFlow = "pv_l4_consumer_v2_bundle_flow"
)

// Segmentation api path constants
const (
	SegmentationServiceApiPath           = "/api/v1"
	CheckTypeIdsSegmentIdsMappingApiPath = "/segments/user/evaluate"
)

// ads service APIs
const (
	SearchSuggestionAds = "/api/v1/search-suggestion" // API to fetch ads is search suggestions mode typed
	CategoryAds         = "/api/v1/category"          // API to fetch ads from category
)

// recipe service API
const (
	SearchSuggestionRecipe = "/api/v1/recipe/suggestion"
)

// In House Auto Suggestion Flow Constants
const (
	UserEnabledforInhouseAutoSuggestionsFlow = "UserEnabledforInhouseAutoSuggestionsFlow"
)
const QueryCityID = "cityId"

const (
	UpsertSucceeded = "UpsertSucceeded"
)

const (
	RateLimitAllowed = "RateLimitAllowed"
)

// Debugger Constants
const (
	DebugEnabledRequestFlag = "is_debug_request"
	BreadcrumbEnabledFlag   = "is_breadcrumb_tracing_enabled"
	DebugProxyUserId        = "debug_proxy_user_id"
	DebugProxyRequestId     = "debug_proxy_request_id"
)
