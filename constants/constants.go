package constants

const ServiceName = "user-search-service"

const (
	PRIMARY_SUBCATEGORY_AS_MILK = "MILK"
	V1_PRE_SEARCH_PATH          = "v1/pre-search"
	V2_PRE_SEARCH_PATH          = "v2/pre-search"
	V1_SEARCH_PATH              = "v1/search"
	V2_SEARCH_PATH              = "v2/search"
	V3_SEARCH_PATH              = "v3/search"
)

// Connection Pooling
const (
	PrometheusServiceName = "user_search_service"
	API                   = "api"
	WORKER                = "worker"
	CRON                  = "cron"
	RPC                   = "rpc"
	USSAdminBackend       = "uss-admin-backend"
)

// Kafka Event Names
const (
	EventNameLoggerKey              = "eventName"
	EventTypeLoggerKey              = "eventType"
	StoreSearchProductsEvent        = "store_search_term_products"
	CitySearchProductsEvent         = "city_search_term_products"
	StoreStressSearchInitiativeName = "STRESS_BASED_SEARCH_RESULTS"
	StoreStressInitiativeName       = "INITIATIVE"
	UserSearchCategoryRelevance     = "userSearchCategoryRelevance"
	UserSearchAttributeRelevance    = "userSearchAttributeRelevance"
	UserSearchSubThemeRelevance     = "userSearchSubThemeRelevance"
	PricingFeatureStoreEventName    = "PRICING_FEATURE_STORE_EVENT"
	UserSearchFeatureRelevance      = "userSearchFeatureRelevance"
)

// Selective Logger Flow Names
const (
	AtlasFlow          = "atlas_flow"
	RecommendationFlow = "recommendation_flow"
	AdsFlow            = "ads_flow"
	PasFlow            = "pas_flow"
	RankerFlow         = "ranker_flow"
	BoostingFlow       = "boosting_flow"
	AutosuggestionFlow = "autosuggestion_flow"
)

const (
	Semantic              = "semantic"
	SemanticRecallPercent = "semantic_recall_percent_value"
)

type EntityType string

const (
	STORE EntityType = "store"
	CITY  EntityType = "city"
)

// Slack Channel Constants
const (
	UserSearchQaAlertChannel   = "qa-user-search-backend-alerts"
	UserSearchProdAlertChannel = "user-search-backend-alerts"
)

// compatible components
const (
	BPC_GROUP_ENABLED              = "BPC_GROUP_DETAILS"
	RERANKING_QCL_RELATED_PRODUCTS = "RERANKING_QCL_RELATED_PRODUCTS"
	PAAN_BANNER_WIDGET_COMPONENT   = "PAAN_BANNER_WIDGETIZED"
	AUTOSUGGESTION_PAGE_ENABLED    = "AUTOSUGGESTION_PAGE_ENABLED"
	AUTOSUGGESTION_AD_PIP          = "AUTOSUGGESTION_AD_PIP"
	PREVIOUSLY_BOUGHT_ENABLED      = "PREVIOUSLY_BOUGHT_ENABLED"
	ROLLUPS_UOM                    = "ROLLUPS_UOM"
	NEW_ROLLUPS_ENABLED            = "NEW_ROLLUPS_ENABLED"
	AUTOSUGGESTION_RECIPE_PIP      = "AUTOSUGGESTION_RECIPE_PIP"
	VBD_COMPATIBLE_COMPONENT       = "VBD"
	SCOPED_SEARCH_V1               = "SCOPED_SEARCH_V1"
	PRE_SEARCH                     = "PRE_SEARCH"
	QUERY_DESCRIPTION_WIDGET       = "QUERY_DESCRIPTION_WIDGET"
	MEDS_WITH_SIMILAR_SALT_WIDGET  = "MEDS_WITH_SIMILAR_SALT_WIDGET"
	SEARCH_PRODUCT_GRID            = "SEARCH_PRODUCT_GRID"
	SEARCH_PRODUCT_GRID_V2         = "SEARCH_PRODUCT_GRID_V2"
	SearchRelook                   = "SEARCH_RELOOK"
	LiquidationWidget              = "LIQUIDATION_WIDGET"
	L3_UNDERSTANDING_LAYOUT        = "L3_UNDERSTANDING_LAYOUT"
	OPEN_BOX_DELIVERY              = "OPEN_BOX_DELIVERY"
	SRP_REDIRECT                   = "SRP_REDIRECT"
	SEARCH_MONETIZATION            = "SEARCH_MONETIZATION"
	VIP_DEAL_TILE_TAKEOVER         = "VIP_DEAL_TILE_TAKEOVER"
	VIP_DEAL_HORIZONTAL_CARD       = "VIP_DEAL_HORIZONTAL_CARD"
	PRODUCT_CARD_ADD_SLOT          = "PRODUCT_CARD_ADD_SLOT"
	DISCOUNTED_ADDONS_ENABLED      = "DISCOUNTED_ADDONS_ENABLED"
	RAKHI_SRP_WIDGET               = "RAKHI_SRP_WIDGET"
	PCA_SINGLE_PRODUCT_VARIANT     = "PCA_SINGLE_PRODUCT_VARIANT"
	AISLE_EXPERIENCE               = "AISLE_SEARCH"
	SEARCH_UI_CLEANUP              = "SEARCH_UI_CLEANUP"
)

// Qrl, Previous bought not done reasons
const (
	GetSearchResultsAPITransaction                  = "GetSearchResultsApiTransactionV3"
	ExperimentClientError                           = "ExperimentationClientError"
	ExperimentNotFound                              = "ExperimentNotFoundForUser"
	ExperimentNotFoundFor2x2And3x3                  = "2x2And3x3ExperimentNotFoundForUser"
	ExperimentVariantNotFound                       = "ExperimentVariantNotFoundForUser"
	UserNotFound                                    = "UserNotFound"
	UserNotEnabledForExperiment                     = "UserNotEnabledForExperimentation"
	QueryIntentEnabled                              = "QueryIntentEnabled"
	MarketplaceIsSuperSaver                         = "MarketplaceIsSuperSaver"
	TenantNotZepto                                  = "TenantNotZepto"
	ErrorFetchingStoreProductsFromStoreCache        = "ErrorFetchingStoreProductsFromCache"
	ErrorFetchingStoreProductsFromCityCache         = "ErrorFetchingStoreProductsFromCityCache"
	ErrorFetchingSemanticStoreProductsFromCityCache = "ErrorFetchingSemanticStoreProductsFromCityCache"
	NoStoreProductsFoundForStoreSearchTerm          = "NoProductsFoundForStoreSearchTermInCache"
	CompatibleComponentsNotFound                    = "CompatibleComponentsNotFound"
	CompatibleComponentsNotFoundFor2x2              = "CompatibleComponentsNotFoundFor2x2"
	CompatibleComponentsNotFoundFor3x3              = "CompatibleComponentsNotFoundFor3x3"
	PreviousBoughtItemsAlreadyBoosted               = "PreviousBoughtItemsAlreadyBoosted"
	PreviousProductsNotFoundInTopNAlgoliaProducts   = "PreviousProductsNotFoundInTop%dAlgoliaProducts"
	PreviousProductsNotFoundInBucketProducts        = "PreviousProductsNotFoundInBucketProducts"
	RecommendedProductsInTopNAlgoliaProducts        = "RecommendedProductsInTop%dAlgoliaProducts"
	ZeroPreviousBoughtProduct                       = "ZeroPreviousBoughtProduct"
	NoStoreProductsFoundForCitySearchTerm           = "NoProductsFoundForCitySearchTermInCache"
	NoSemanticStoreProductsFoundForCitySearchTerm   = "NoSemanticProductsFoundForCitySearchTermInCache"
	NoProductsFoundInPasByGivenPvIDs                = "NoProductsFoundInPasByGivenPvIDs"
	DisabledDueToRelevanceMismatch                  = "DisabledDueToRelevanceMismatch"
	SearchModeDisabledForScopedSearch               = "SearchModeDisabledForScopedSearch"
	ErrorFetchingFeedProductsFromInhouse            = "ErrorFetchingFeedProductsFromInhouse"
	ZeroInStockResultsInPrimaryScopedFeed           = "ZeroInStockResultsInPrimaryScopedFeed"
	NoYourPicksFilterSelected                       = "NoYourPicksFilterSelected"
	SuperSaverMode                                  = "SearchInSuperSaverMode"
	ErrorParsingCityId                              = "ErrorParsingCityId"
	ErrorParsingPVId                                = "ErrorParsingPVId"
	ErrorFetchingExperimentInfo                     = "ErrorFetchingExperimentInfo"
	RecommendationsDataNotAvailable                 = "RecommendationsDataNotAvailable"
	QueryIntentNotDone                              = "QueryIntentNotDone"
	GetSearchResultsFromFiltersAPITransaction       = "GetSearchResultsFromFiltersApiTransactionV3"
	RankerNotApplicableForV2Flow                    = "RankerNotApplicableForV2Flow"
	ModelResponseError                              = "ModelResponseError"
	GoroutineFailureForNeuronCall                   = "GoroutineFailureForNeuronCall"
	ProductCountCriteriaFailed                      = "ProductCountCriteriaFailed"
	ProductGridInfo2x2DueToPLPForFTB                = "ProductGridInfo2x2DueToPLPForFTB"
	ZeptoCafeTenant                                 = "ZeptoCafeTenant"

	GetSearchResultsAPIV2Transaction           = "GetSearchResultsApiTransactionV2"
	GetResultsAPITransaction                   = "GetResultsApiTransactionV1"
	GetVipDealHorizontalCardSKURecommendations = "GetVipDealHorizontalCardSKURecommendations"
	GetNonActivatedL2PVIdsForUser              = "GetNonActivatedL2PVIdsForUser"
	GetL3CrossSellRecommendationsInfo          = "GetL3CrossSellRecommendationsInfo"

	InvalidMarketPlaceFound    = "InvalidMarketPlaceFound"
	CvrModelResultsIncomplete  = "Cvr model returned results for %v autosuggestions, total autosuggestions are %v"
	CtrModelResultsIncomplete  = "Ctr model returned results for %v autosuggestions, total autosuggestions are %v"
	USSRateLimitingTransaction = "USSRateLimitingTransaction"

	GetRelevanceScoreForQueryV2          = "GetRelevanceScoreForQueryV2"
	HandleUserSearchRelevanceEventV2     = "HandleUserSearchRelevanceEventV2"
	HandleBulkUserSearchRelevanceEventV2 = "HandleBulkUserSearchRelevanceEventV2"
	GetOpsOverriddenQueries              = "GetOpsOverriddenQueries"

	GetPricingFeatureStoreMiscAPIResultsAPITransaction = "GetPricingFeatureStoreMiscAPIResultsAPITransaction"
	FashionRevampLayoutEnabled                         = "FashionRevampLayoutEnabled"
)

// meds with similar salt global meta constants
const (
	CountOfRecallItemsOrFinalEnrichedProductsLessThan4 = "CountOfRecallItemsOrFinalEnrichedProductsLessThan4"
)

// tobacco widget
const (
	TobaccoWidgetTitle                   = "Paan Corner"
	TobaccoWidgetId                      = "TOBACCO_BANNER"
	TobaccoWidgetName                    = "TOBACCO_BANNER_SEARCH"
	TobaccoWidgetType                    = "tobacco_banner"
	TobaccoWidgetPreTitle                = "To order"
	TobaccoWidgetPostTitle               = "products"
	TobaccoWidgetCtaText                 = "Get Zepto Lite"
	TobaccoWidgetCtaTextIOS              = "View items"
	TobaccoWidgetSubtitleRedirectLiteIOS = "Click on the button below to view items"
	TobaccoWidgetSubtitleRedirectLite    = "Visit the below link to order"
	TobaccoWidgetSubtitleCopy            = "Copy the below link and open it on browser"
	TobaccoWidgetSubtitleRedirect        = "Order online by clicking on the link below"
	TobaccoWidgetInfo                    = "Playstore guidelines prohibit the sale of Paan Corner products on apps. Use Zepto Web to order directly"
	TobaccoWidgetInfoLite                = "Playstore guidelines prohibit the sale of Paan Corner products on apps. Use Zepto Lite to order directly."
)

// tobacco widget v2
const (
	TobaccoWidgetIdV2   = "PREMIUM_PROMO"
	TobaccoWidgetNameV2 = "PAAN_CORNER_PREMIUM_PROMO"
)

const (
	QueryReRankingExperiments                     = "QUERY_RERANKING"
	RelatedProductsExperiments                    = "RELATED_PRODUCTS"
	AutoSuggestionExperimentKey                   = "AUTO_SUGGESTION"
	UserPreviousBoughtExperimentKey               = "PREVIOUSLY_BOUGHT"
	InhouseAutoSuggestionExperimentKey            = "INHOUSE_AUTOSUGGESTION_FLOW"
	YouMightLikeExperimentKey                     = "YOU_MIGHT_LIKE_EXPERIMENT"
	InhouseAutoSuggestionIndexExperimentKey       = "INHOUSE_AUTOSUGGESTION_INDEX"
	PlpOnSearchExperimentKey                      = "PLP_ON_SEARCH"
	RelevanceExperimentKey                        = "RELEVANCE_SCORE"
	BlockedAdsExperimentName                      = "TOS_BLOCKED_KEYWORDS_ADS_EXPERIMENT"
	PcaExperimentName                             = "PCAExperimentBackend"
	PcaAdsV3ExperimentName                        = "PCA_ADS_V3_EXPERIMENT"
	TOSAdsExperimentName                          = "TOSExperiment"
	NewLayoutAdsExperimentName                    = "NewLayoutAdsExperimentName"
	AdsCategoryDensityExperimentName              = "AdsCategoryDensityExperiment"
	AdsSlotRevampExperimentName                   = "AdsSlotRevampExperiment"
	OldLayoutAdsExperimentName                    = "OldLayoutAdsExperiment"
	IsTOSPriorityEnabled                          = "is_tos_priority_enabled"
	SearchPageAdsExperimentKey                    = "SEARCH_PAGE_ADS"
	SearchRevampExperimentKey                     = "SEARCH_REVAMP_EXPERIMENT"
	SearchRevampRecallExperimentKey               = "SEARCH_REVAMP_RECALL_EXPERIMENT"
	TemplateLatencyImprovementExperimentKey       = "template_latency_improvement"
	L3SearchRelevanceExperimentKey                = "L3_SEARCH_RELEVANCE_EXPERIMENT"
	PostAtcCrossSell                              = "POST_ATC_CROSS_SELL_EXPERIMENT_NAME"
	DsPoweredRelatedProductsFeedExperimentKey     = "DSPoweredRelatedProductsFeedExperiment"
	QueryIntentExperimentKey                      = "QUERY_INTENT"
	SearchKeywordCompleteMatchExperimentKey       = "sk_exact_experiment"
	L3SearchEnableExperimentKey                   = "l3_search_enabled_experiment"
	ScopedSearchExperimentKey                     = "scoped_search_experiment"
	PreAtcCrossSell                               = "PRE_ATC_CROSS_SELL_EXPERIMENT"
	PreSearchPageCrossSellExperimentName          = "pre_search_cross_sell_experiment"
	QueryIntentM2ExperimentKey                    = "QUERY_INTENT_M2"
	PharmaSearchEnabledExperimentName             = "pharma_search_enabled"
	OrganicOosExperimentName                      = "ORGANIC_OOS_EXPERIMENT"
	SemanticExperimentName                        = "SEMANTIC_EXPERIMENT"
	SemanticExperimentV2Name                      = "SEMANTIC_EXPERIMENT_V2"
	SearchPageLiquidationExperimentName           = "SEARCH_PAGE_LIQUIDATION_EXPERIMENT"
	SRPPostAtcRecoCrossSell                       = "SRP_POST_ATC_RECO_MIGRATE_EXP"
	QCLFuzzyDisableExperimentKey                  = "qcl_fuzzy_disable_experiment"
	ProductGridExperimentKey                      = "PRODUCT_GRID_EXPERIMENT"
	ProductGrid3x3ExperimentKey                   = "PRODUCT_GRID_3X3_EXPERIMENT"
	QCLV3ExperimentName                           = "qcl_v3_experiment"
	BrowsePostAtcRecommendationExperimentKey      = "BrowsePostAtcRecommendationExperiment"
	ShowTobaccoProductsForAndroidExperimentKey    = "show_tobacco_products_for_android_experiment"
	OpenBoxDeliveryExperimentKey                  = "pasOpenBoxDeliveryExperiment"
	UserRecommendationsV2ExperimentKey            = "user_recommendation_v2"
	ThemeBasedBuyingExperimentKey                 = "theme_based_buying_experiment"
	NonQueryIntentCrossSellWidgetExperimentKey    = "non_query_intent_cross_sell_widget_experiment"
	QueryIntentCrossSellWidgetExperimentKey       = "query_intent_cross_sell_widget_experiment"
	RankerV2ExperimentKey                         = "ranker_v2_experiment"
	PvidSearchCatalogueExperimentName             = "PVID_SEARCH_CATALOGUE_EXPERIMENT"
	SearchPageLBentoExperimentName                = "SEARCH_PAGE_BENTO_EXPERIMENT"
	PostATCBottomSheetExperimentKey               = "POST_ATC_BOTTOM_SHEET_EXPERIMENT"
	RelatedProductMigrationExperimentKey          = "related_products_city_level_migration"
	Ads3x3LayoutExperiment                        = "ADS_3x3_LAYOUT_EXPERIMENT"
	Ads2x2LayoutExperiment                        = "ADS_2x2_LAYOUT_EXPERIMENT"
	ModelRankerExperimentKey                      = "model_ranker_experiment"
	ThemeBasedGodProductExperimentKey             = "theme_based_god_product_experiment"
	ThemeBasedPremiumCarouselExperimentKey        = "theme_based_premium_carousel_experiment"
	ThemeBasedRecallRankerExperimentKey           = "theme_based_recall_ranker_experiment"
	GPTThemeExperimentKey                         = "gpt_theme_experiment"
	FTBPersonalisationExperimentKey               = "ftb_personalisation_experiment"
	ThemePaginationExperimentKey                  = "theme_pagination_experiment"
	SearchMonetizationExperimentKey               = "search_monetization_experiment"
	ModelRankerExperimentSSKey                    = "model_ranker_experiment_ss"
	SemanticExperimentV3Name                      = "SEMANTIC_EXPERIMENT_V3"
	RulePriorityOverL3ExperimentNameKey           = "RULE_PRIORITY_OVER_L3_EXPERIMENT"
	MonetRpiRankingExperimentNameKey              = "monet_rpi_ranking_experiment"
	AttributeUnderstandingExperimentNameKey       = "SEARCH_ATTRIBUTE_UNDERSTANDING_V1"
	VipDealsProductTileExperimentKey              = "VIP_DEALS_PRODUCT_TILE_EXPERIMENT"
	BundleHandlingExperimentKey                   = "BUNDLE_HANDLING_EXPERIMENT"
	OosDeboostingExperimentNameKey                = "OOS_DEBOOSTING_EXPERIMENT"
	VipDealsHorizontalCardExperimentKey           = "VIP_DEALS_HORIZONTAL_CARD_EXPERIMENT"
	AutosuggestionRankerExperimentNameKey         = "AUTOSUGGESTION_LIVE_RANKER"
	SemanticExperimentV4Name                      = "SEMANTIC_EXPERIMENT_V4"
	SemanticExperimentV5Name                      = "SEMANTIC_EXPERIMENT_V5"
	AdsSlotChangeExperimentName                   = "ADS_SLOT_CHANGE_ADTAG_PC"
	SpeicalEventExperimentKey                     = "SPECIAL_EVENT_EXPERIMENT"
	SearchVectorBucketizationV1                   = "SearchVectorBucketizationV1"
	ClmFallbackRankerExperimentName               = "CLM_FALLBACK_RANKER"
	ExploreExploitExperimentName                  = "EXPLORE_EXPLOIT_RANKER"
	ItemizationDeboostingBelowInstockThreshold    = "ItemizationDeboostingBelowInstockThreshold"
	TailOfTailRelevanceV1                         = "TailOfTailRelevanceV1"
	RTVectorSearch                                = "RTVectorSearch"
	EmbeddingModelRTVectorSearch                  = "EmbeddingModelRTVectorSearch"
	RankingModelExperimentKey                     = "RANKING_MODEL_EXPERIMENT"
	RelevanceL3DeciderExperimentKey               = "RELEVANCE_L3_DECIDER"
	AutosuggestionWeightedRankerExperimentNameKey = "AUTOSUGGESTION_WEIGHTED_RANKER"
	PricingFeatureStoreRankingExperimentNameKey   = "PRICING_FEATURE_STORE_RANKING_EXPERIMENT"
	StoreWiseBatchingModelRanker                  = "STORE_WISE_BATCHING_MODEL_RANKER"
	FashionWidgetExperimentName                   = "FASHION_WIDGET_EXPERIMENT"
	RelevanceVersionExperiment                    = "RELEVANCE_VERSION_EXPERIMENT"
	HeadQueryFuzzyRemovalExperimentName           = "HEAD_QUERY_FUZZY_REMOVAL_EXPERIMENT"
	Tier2TailOfTailRelevanceV1                    = "Tier2TailOfTailRelevanceV1"
	AdsDeboostingExperiment                       = "ADS_DEBOOSTING_EXPERIMENT"
	AdsFlyWheelExperiment                         = "ADS_FLY_WHEEL_EXPERIMENT"
	PreviouslyBoughtExperimentV2Name              = "PREVIOUSLY_BOUGHT_EXPERIMENT_V2"
	BuyAgainWithPredictiveCartExperimentName      = "BUY_AGAIN_WITH_PREDICTIVE_CART_EXPERIMENT"
	NewQclV3ExperimentName                        = "NEW_QCL_V3_EXPERIMENT"
	AisleExperimentName                           = "SEARCH_AISLE_EXPERIMENT"
	ExploreExpansionExperimentNameKey             = "EXPLORE_EXPANSION"
	SearchSuperSaver3x3LayoutExperimentName       = "SEARCH_SUPERSAVER_3X3_LAYOUT_EXPERIMENT"
)

// QCL Meta
const (
	QclOverrideMeta         = "OstOverrideBucket"
	QclDsTeamMeta           = "OstDsTeamBucket"
	OstMatchedWithPst       = "OstMatchedWithPst"
	OstMatchedWithPstV2     = "OstMatchedWithPstV2"
	OstMatchedWithPstV3     = "OstMatchedWithPstV3"
	AllPSTsInActive         = "AllPSTsInActive"
	SearchModeNotApplicable = "SearchModeNotApplicable"
	QCLV2PSTNotFound        = "OstPstMappingNotFound"
	QCLErrorFound           = "InternalDbError"
	QCLV3PSTNotFound        = "OstPstMappingV3NotFound"
)

// QCL Algo Thresholds
const (
	AllChars = "abcdefghijklmnopqrstuvwxyz1234567890"
)

// Pca Slots Variants
const (
	PcaControlVariant = "pca_c_0_1"
	Variant1_2_3      = "pca_v1_2_3"
	Variant1_1_3      = "pca_v1_1_3"
	Variant2_2_3      = "pca_v2_2_3"
	Variant2_1_3      = "pca_v2_1_3"
	Variant3_2_3      = "pca_v3_2_3"
	Variant3_1_3      = "pca_v3_1_3"
)

// span variables
const (
	UserSearchServiceTracer                                       = "UserSearchServiceTracer"
	QclAlgorithm                                                  = "QclAlgorithm"
	IsCompatibleComponentsFound                                   = "IsCompatibleComponentsFound"
	ExperimentVariantName                                         = "ExperimentVariantName"
	IsInternalUser                                                = "IsInternalUser"
	SearchTermLength                                              = "searchTermLength"
	TopSearchMapSize                                              = "TopSearchMapSize"
	CorrectTermsWithSameChar                                      = "CorrectTermsWithSameChar"
	RelatedProductsSearchTerm                                     = "RelatedProductsSearchTerm"
	ReRankingListProductCountInCache                              = "ReRankingListProductCountInCache"
	ReRankingListSemanticProductCountInCache                      = "ReRankingListSemanticProductCountInCache"
	AlgoliaSearchProductCount                                     = "AlgoliaSearchProductCount"
	AlgoliaSemanticSearchProductCount                             = "AlgoliaSemanticSearchProductCount"
	CommonProductCountInCacheAndAlgolia                           = "CommonProductCountInCacheAndAlgolia"
	CommonSemanticProductCountInCacheAndAlgolia                   = "CommonSemanticProductCountInCacheAndAlgolia"
	StoreProductRankChangedCount                                  = "StoreProductRankChangedCount"
	SemanticStoreProductRankChangedCount                          = "SemanticStoreProductRankChangedCount"
	RelatedProductsDefaultTerm                                    = "RelatedProductsDefaultTerm"
	RelatedProductsDefaultTermCount                               = "RelatedProductsDefaultTermCount"
	RelatedProductsSubcategoryTerm                                = "RelatedProductsSubcategoryTerm"
	RelatedProductsSubcategoryTermCount                           = "RelatedProductsSubcategoryTermCount"
	RelatedProductsReason                                         = "RelatedProductsReason"
	RelatedProductsEnabled                                        = "RelatedProductsEnabled"
	AutoSuggestionPageEnabled                                     = "AutoSuggestionPageEnabled"
	AutoSuggestionTestVariant                                     = "AutoSuggestionTestVariant"
	SearchMode                                                    = "SearchMode"
	CityName                                                      = "CityName"
	IsSearchKeyPrefixMatched                                      = "IsSearchKeyPrefixMatched"
	ProductBoostCount                                             = "productBoostCount"
	OosProductsCount                                              = "OosProductsCount"
	InStockProductsCount                                          = "InStockProductsCount"
	StoreId                                                       = "StoreId"
	StoreIds                                                      = "StoreIds"
	UserId                                                        = "UserId"
	FilterInStockAndOutOfStock                                    = "filterInStockAndOutOfStock"
	FilterProductsForBulkSearch                                   = "filterProductsForBulkSearch"
	GetSearchResultsV3SuperStore                                  = "GetSearchResultsV3SuperStore"
	GetSearchResultsV3ForFilters                                  = "GetSearchResultsV3ForFilters"
	MakeFiltersFromAtlasResponse                                  = "MakeFiltersFromAtlasResponse"
	GetStoreProductsUsingPvIds                                    = "GetStoreProductsUsingPvIds"
	FetchStoreProductsFromPasUsingPvIds                           = "FetchStoreProductsFromPasUsingPvIds"
	FetchStoreProductsFromPasUsingPvIdsBatchWise                  = "FetchStoreProductsFromPasUsingPvIdsBatchWise"
	HandleCmsSubCategoryData                                      = "HandleCmsSubCategoryData"
	GetQueryCategorizationDetails                                 = "GetQueryCategorizationDetails"
	GetQueryCategorizationDetailsV2                               = "GetQueryCategorizationDetailsV2"
	GetFilterAttachmentDetails                                    = "GetFilterAttachmentDetails"
	HandleQueryCategorizationData                                 = "HandleQueryCategorizationData"
	HandleFilterAttachmentData                                    = "HandleFilterAttachmentData"
	ProcessCMSCreateUpdateEvent                                   = "ProcessCMSCreateUpdateEvent"
	RerankingVariantName                                          = "RerankingVariantName"
	SemanticRerankingVariantName                                  = "SemanticRerankingVariantName"
	SemanticRerankingVariantNameField                             = "ranking_variant"
	SemanticRerankingVariantNotFound                              = "SemanticRerankingVariantNotFound"
	CityId                                                        = "CityId"
	ResponseCacheHitInterval                                      = "ResponseCacheHitInterval"
	GetDsPoweredRelatedProductsData                               = "GetDsPoweredRelatedProductsData"
	GetDsPoweredRelatedProductsDataForSuperStore                  = "GetDsPoweredRelatedProductsDataForSuperStore"
	SearchResultQuery                                             = "SearchResultQuery"
	AutosuggestionResultQuery                                     = "AutosuggestionResultQuery"
	SearchResultQuerySuperstore                                   = "SearchResultQuerySuperstore"
	FetchingDataForCacheSuperStoreFlow                            = "FetchingDataForCacheSuperStoreFlow"
	ProcessStoreProductsFromPas                                   = "ProcessStoreProductsFromPas"
	GetSearchProductsV3                                           = "GetSearchProductsV3"
	FetchSpFromPasUsingPvIdsForSuperStore                         = "FetchSpFromPasUsingPvIdsForSuperStore"
	FetchSpecialResponseFromPasUsingPvIDs                         = "FetchSpecialResponseFromPasUsingPvIDs"
	FetchThemeExperienceResponseFromPasUsingPvIDs                 = "FetchThemeExperienceResponseFromPasUsingPvIDs"
	SaveCrossSellWidgetDetailsTransaction                         = "save-cross-sell-widgets"
	GetCrossSellWidgetDetailsTransaction                          = "get-cross-sell-widgets"
	UpsellProductBoostCount                                       = "upsellProductBoostCount"
	HandleQueryDescriptionCreateUpdateEvent                       = "HandleQueryDescriptionCreateUpdateEvent"
	GetQueryDescriptionForSearchTerm                              = "GetQueryDescriptionForSearchTerm"
	HandlePVL4AttributeValueCreateUpdateEvent                     = "HandlePVL4AttributeValueCreateUpdateEvent"
	StoreSaltMappings                                             = "storeSaltMappings"
	RunCacheSaltMappingsForQueryTermsCron                         = "RunCacheSaltMappingsForQueryTermsCron"
	MakeMedsWithSimilarSaltCarouselAtlasCall                      = "makeMedsWithSimilarSaltCarouselAtlasCall"
	FetchStoreProductsForLiquidationWidget                        = "FetchStoreProductsForLiquidationWidget"
	FetchLiquidationStoreProductsFromPas                          = "FetchLiquidationStoreProductsFromPas"
	GetStoreProductsForLiquidation                                = "GetStoreProductsForLiquidation"
	PublishSearchResultMessage                                    = "PublishSearchResultMessage"
	PublishSearchPageMetricsMessage                               = "PublishSearchPageMetricsMessage"
	GetBulkExperimentDetails                                      = "GetBulkExperimentDetails"
	GetSearchAdsBannersPca                                        = "GetSearchAdsBannersPca"
	EnrichResponseWithIntent                                      = "EnrichResponseWithIntent"
	FooterSetForInvalidPage                                       = "FooterSetForInvalidPage"
	ProcessHitsPVL4                                               = "ProcessHitsPVL4"
	ThemeBasedEnrichResponseWithIntent                            = "ThemeBasedEnrichResponseWithIntent"
	DefaultRerankingVariantName                                   = "DefaultRerankingVariantName"
	StoreProductModelRankChangedCount                             = "StoreProductModelRankChangedCount"
	FetchStoreProductByStoreSubcategoryId                         = "FetchStoreProductByStoreSubcategoryId"
	FetchBentoSearchProductsFromPas                               = "FetchBentoSearchProductsFromPas"
	FetchStoreProductsForBentoWidget                              = "FetchStoreProductsForBentoWidget"
	FetchDataForThemeCacheSuperStoreFlow                          = "FetchingDataForThemeCacheSuperStoreFlow"
	ProcessHitsParentBundlePVs                                    = "ProcessHitsParentBundlePVs"
	HandlePVL4AttributeValueCreateUpdateEventForBundleHandling    = "HandlePVL4AttributeValueCreateUpdateEventForBundleHandling"
	QueryReRankingLayerExploreExploitWrapperForAtlasStoreProducts = "QueryReRankingLayerWrapperForAtlasStoreProducts"
)

const (
	Override = "OVERRIDE"
	DSTeam   = "DS_TEAM"
)

// search modes
const (
	Typed                   = "TYPED"
	RecentSearch            = "RECENT_SEARCH"
	Brand                   = "BRAND"
	Autosuggest             = "AUTOSUGGEST"
	ShowAllResults          = "SHOW_ALL_RESULTS"
	RelatedSearches         = "RELATED_SEARCHES"
	BottomSheetList         = "BOTTOM_SHEET_LIST"
	MainSearch              = "MainSearch"
	ThemeBasedSeeAllResults = "THEME_BASED_SEE_ALL"
	Recipe                  = "RECIPE"
	AisleTab                = "AISLE_TAB"
	QclIgnore               = "QCL_IGNORE"
)

// relevance info constants
const (
	EMPTY_STRING = ""
	OTHERS       = "others"
	VERSION_V1   = "v1"
	VERSION_V2   = "v2"

	//Operation Type strings for opsOverride details
	ADD     = "ADD"
	REMOVE  = "REMOVE"
	Version = "version"
)

// mediaType ads
const (
	Image                      = "IMAGE"
	Video                      = "VIDEO"
	GetPcaPanicRecoveryCommand = "GetPcaProducts"
)

// Related Products
const (
	RelatedProductsSearchKey                = "search_related_products"
	RelatedProductsSubcategoryKey           = "search_subcategory_details"
	RelatedProductsDefault                  = "default"
	RelatedProductsSubcategory              = "subcategory"
	RelatedProductEmptyList                 = "RelatedProductAlgoliaSearchResultNotFound"
	RelatedProductNotFound                  = "RelatedProductKeyNotFound"
	RelatedProductAllDuplicate              = "RelatedProductAllSubcategoryDuplicate"
	RelatedProductsTitle                    = "Other similar products"
	DsPoweredRelatedProductsTitle           = "Other related products"
	RelatedProductsMaxCsvRecord             = 50000
	ResetRpCache                            = "reset"
	RelatedProductsPasBatchSize             = 10
	RelatedProductsCacheKey                 = "search_related_products_subcategory_%s_%s"     // storeId_subcategoryName
	RelatedProductsSuperstoreCacheKey       = "superstore_related_products_subcategory_%s_%s" // storeId_subcategoryName
	RelatedProductsSuperstoreCacheKeyNew    = "superstore_rp_subcategory_%s_%s"               // storeId_subcategoryName
	RelatedProductsSearchKeySubCatIdMapping = "query_subcat_id_mapping"
	DsRelatedProductsCacheKey               = "ds_search_related_products_subcategory_%s_%s" // cityId_query
	QueryTermRedirectionMaxCsvRecord        = 500
	QueryTermRedirectionBatchSize           = 100

	RelatedProductsCityMigrationPvIdsCacheKey                = "related_products:%s_%s:pv_ids"                 // related_products:cityId_subcatId:pv_ids
	RelatedProductsCityMigrationStoreIdsCacheKey             = "related_products:%s_%s:store_ids"              // related_products:cityId_subcatId:store_ids
	RelatedProductsCityMigrationPvIdsCacheKeyUncompressed    = "related_products:%s_%s:pv_ids_uncompressed"    // related_products:cityId_subcatId:pv_ids
	RelatedProductsCityMigrationStoreIdsCacheKeyUncompressed = "related_products:%s_%s:store_ids_uncompressed" // related_products:cityId_subcatId:store_ids
)

// Experience plpInSearch constants
const (
	MaxLimitForSearchProducts             = 10
	YourSearchResults                     = "Search Results"
	ExploreRelatedCategories              = "RELATED CATEGORIES"
	PlpSearchType                         = "search"
	PlpWidgetType                         = "widget"
	PlpSubcategoryType                    = "subcategory"
	PlpSearchCompatibleComponents         = "PLP_ON_SEARCH"
	PageTypeNormal                        = "normal"
	PageTypeSubcategory                   = "search_with_subcategories"
	PageTypeThemeBasedSearch              = "theme_based_search"
	CoreLogic                             = "core"
	FallbackLogic                         = "fallback"
	AlgoliaNoProductsFound                = "AlgoliaNoProductsFound"
	SearchModeBrand                       = "BRAND"
	SearchModeBrandTypeIncompatible       = "SearchModeBrandTypeIncompatible"
	PlpThemeBasedSeeAllCall               = "PlpThemeBasedSeeAllCall"
	PlpCompatibleComponentsNotFound       = "PlpCompatibleComponentsNotFound"
	PlpExperimentNotEnabled               = "PlpExperimentNotEnabled"
	PlpPasCategoryDetailsNotFound         = "PlpPasCategoryDetailsNotFound"
	PlpInactiveCategoryFound              = "PlpInactiveCategoryFound"
	PlpInvalidSubcategoriesCount          = "PlpInvalidSubcategoriesCount"
	PlpBlacklistedCategory                = "PlpBlacklistedCategory"
	PlpFTBDisabledCategory                = "PlpFTBDisabledCategory"
	PlpInSearchNotMarketPlace             = "PlpInSearchNotMarketPlace"
	PlpInSearchNotZeptoNowOrSuperSaver    = "PlpInSearchNotZeptoNowOrSuperSaver"
	PlpCategoryIdDataNotFound             = "PlpCategoryIdDataNotFound"
	AnalyticsSearchProductsUsed           = "analytics_search_products_used"
	AnalyticsSearchProductsUsedNoUser     = "analytics_search_products_used_no_user"
	IsMainThemePageCall                   = "IsMainThemePageCall"
	IsMainThemePageOverridenToQueryIntent = "IsMainThemePageOverridenToQueryIntent"
	ThemeVariant                          = "ThemeVariant"
)

const (
	BoostingExperimentNotEnabled = "BoostingExperimentNotEnabled"
	RelevanceQueryNotInRedis     = "RelevanceQueryNotInRedis"
	RelevanceQueryNotInMongo     = "RelevanceQueryNotInMongo"
	FilteringBrandDisabled       = "FilteringBrandDisabled"
	DisableFlagPassed            = "DisableFlagPassed"
	BoostingEmptySearchProducts  = "BoostingEmptySearchProducts"
)

// Internal Users Experiment constants
const (
	InternalUserRelatedProductExperimentName              = "InternalUserRelatedProductExperiment"
	InternalUserRelatedProductExperimentId                = "a07e4971-d215-4d89-989a-3200c330ebbb"
	InternalUserRelatedProductTestVariant                 = "InternalUserRelatedProductTestVariant"
	InternalUserQrlExperimentName                         = "InternalUserQrlExperiment"
	InternalUserQrlExperimentId                           = "bae102b6-6d08-4ec7-98c8-1eb65163ac51"
	InternalUserQrlRpiTestVariant                         = "InternalUserQueryLevelRPI"
	InternalUserQrlAtcpiTestVariant                       = "InternalUserQueryLevelATCPI"
	InternalUserQrlVDRVariant                             = "InternalUserQueryLevelVDR"
	InternalUserQrlSSFRVariant                            = "InternalUserQueryLevelSSFR"
	InternalUserAutoSuggestionExperimentName              = "InternalUserAutoSuggestionExperiment"
	InternalUserAutoSuggestionExperimentId                = "1070bbfb-6c5c-418c-9093-e83e2afdca11"
	InternalUserAutoSuggestionTestVariant                 = "InternalUserAutoSuggestionTestVariant"
	InternalUserPreviousBoughtExperimentName              = "InternalUserPreviousBoughtExperiment"
	InternalUserPreviousBoughtExperimentId                = "1ae102b6-6d08-4ec7-98c8-1eb65163ac51"
	InternalUserPreviousBoughtTestVariant                 = "InternalUserPreviousBoughtTestVariant"
	InternalUserPreviousBoughtUpSellVariantExperimentName = "InternalUserPreviousBoughtUpSellVariantExperiment"
	InternalUserPreviousBoughtUpSellVariantExperimentId   = "1ae102a6-6d08-4ec7-98c8-1eb65163ac51"
	InternalUserPreviousBoughtUpSellTestVariant           = "InternalUserPreviousBoughtUpSellTestVariant"
	UserQrlControl                                        = "control"

	InternalUserPlpExperienceExperimentName = "InternalUserPlpExperienceExperiment"
	InternalUserPlpExperienceExperimentId   = "f2f3727a-00cf-4ee3-86d2-840853ad7dba"
	InternalUserPlpExperienceTestVariant    = "test1"
	PreviousBoughtForReRank                 = "PreviousBoughtForReRankLayer"
)

const (
	MediaTypeImage              = "IMAGE"
	SearchSuggestionIconName    = "icon"
	SearchSuggestionAdTagName   = "ad_tag"
	SearchSuggestionAdTagHeight = 14
	SearchSuggestionAdTagWidth  = 14
)

// Pre search Ads Experiment Constants
const (
	PreSearchAdsExperimentName = "PreSearchAdsExperiment"
	PreSearchAdsTestVariant1   = "Test1"
	PreSearchAdsTestVariant2   = "Test2"
	PreSearchAdsTestVariant3   = "Test3"
)

// Methods
const (
	ApplyQclLayer                                   = "ApplyQclLayer"
	GetRelatedProducts                              = "GetRelatedProducts"
	ApplyRelatedProducts                            = "ApplyRelatedProducts"
	ApplyQueryReRankingLayerWrapper                 = "ApplyQueryReRankingLayerWrapper"
	ApplyQueryReRankingLayerWrapperForSemantic      = "ApplyQueryReRankingLayerWrapperForSemantic"
	UpdateInternalUserTestVariant                   = "UpdateInternalUserTestVariant"
	GetSearchResponseAndOffsets                     = "GetSearchResponseAndOffsets"
	GetThemeSearchResponseAndOffsets                = "GetThemeSearchResponseAndOffsets"
	GenerateSearchResultResponseLayout              = "GenerateSearchResultResponseLayout"
	GenerateSearchResultResponseLayoutV2            = "GenerateSearchResultResponseLayoutV2"
	GetPCABannerForEnrichment                       = "GetPCABannerEnrichment"
	QueryReRankingLayerWrapperForAtlasStoreProducts = "QueryReRankingLayerWrapperForAtlasStoreProducts"
	IsApplicableForRankerV2Flow                     = "IsApplicableForRankerV2Flow"
	ModelQueryRankingLayerForAtlasSearchProducts    = "ModelQueryRankingLayerForAtlasSearchProducts"
	NeuronModelProcessBatch                         = "NeuronModelProcessBatch"
	NeuronModelCall                                 = "NeuronModelCall"
	NeuronModelProcessBatchV2                       = "NeuronModelProcessBatchV2"
	AutosuggestionCvrModelAllBatches                = "AutosuggestionCvrModelAllBatches"
	AutosuggestionCvrModelProcessBatch              = "AutosuggestionCvrModelProcessBatch"
	AutosuggestionCtrModelAllBatches                = "AutosuggestionCtrModelAllBatches"
	AutosuggestionCtrModelProcessBatch              = "AutosuggestionCtrModelProcessBatch"
	GetAutosuggestions                              = "GetAutosuggestions"
	HandlePricingFeatureStoreEvent                  = "HandlePricingFeatureStoreEvent"
	UpdatePricingFeatureStoreCacheData              = "UpdatePricingFeatureStoreCacheData"
	GetPricingFeatureStoreDataFromRedis             = "GetPricingFeatureStoreDataFromRedis"
	UpsertPricingFeatureStoreEventsScylla           = "UpsertPricingFeatureStoreEventsScylla"
	UpsertPricingFeatureStoreEventsRedis            = "UpsertPricingFeatureStoreEventsRedis"
	GetPricingFeatureStoreData                      = "GetPricingFeatureStoreData"
	GetPricingFeatureStoreDataScylla                = "GetPricingFeatureStoreDataScylla"
	ScyllaQueryBatchPVID                            = "ScyllaQueryBatchPVID"
)

// Redis Keys
const (
	Pune      = "pune"
	Mumbai    = "mumbai"
	Bengaluru = "bengaluru"
	Hyderabad = "hyderabad"
	Chennai   = "chennai"
	Kolkata   = "kolkata"
)

// Redis Keys for Ads
const (
	OnlinesalesPreSearchAdsDisable    = "onlinesales_ads_pre_search_disable"
	AutosuggestionSearchTermsRedisKey = "AUTOSUGGESTION_SEARCH_TERMS"
)

// constants for Zero UpSell Higher Variants
const (
	HigherVariantsNotFoundAndOutOfStock = "HigherVariantsNotFoundAndOutOfStock"
	OutOfStockHigherVariants            = "OutOfStockHigherVariants"
	HigherVariantsNotFound              = "HigherVariantsNotFound"
)

// environments
const (
	Development = "development"
	Production  = "production"
)

// pas constants
const (
	FetchStoreProductsFromPasRequestCount = "fetchStoreProductsFromPasRequestCount"
	FetchStoreProductsFromPasStatus       = "fetchStoreProductsFromPasStatus"
	FetchStoreProductsFromPasCountDiff    = "fetchStoreProductsFromPasCountDiff"
	TotalProductsToPas                    = "totalProductsToPas"
	TotalProductsFromPas                  = "totalProductsFromPas"
	TotalActiveProductsFromPas            = "totalActiveProductsFromPas"
	TotalInstockProductsFromPas           = "totalInstockProductsFromPas"
	TotalInActiveProductsFromPas          = "totalInActiveProductsFromPas"
	TotalOutOfStockProductsFromPas        = "totalOutOfStockProductsFromPas"
)

const FilteredStoreProductsForSubcat = "filteredStoreProductsForSubcat"

// widget ID
const (
	ProductGrid = "PRODUCT_GRID"
)

// user search relevance
const (
	DefaultRelevanceScore = 1
)

// UserAppVersionUpdate

const (
	ViewMetaTitleColor           = "#2B1E35"
	ViewMetaCtaTitleColor        = "#FF3269"
	ViewMetaTitle                = "App Update"
	ViewMetaAspectRatio          = 2.68
	ViewMetaMinimumItemsRequired = 1
	ViewMetaSpacerHeight         = 32
	ViewMetaPaddingLeft          = 16
	ViewMetaPaddingRight         = 16
	PageType                     = "normal"
	PremiumPromo                 = "PREMIUM_PROMO"
	FetchMode                    = "network"
)

const CityIdByStoreIdKey = "store_city_map_%s"

// Ad tag enhancement experiment constants
const (
	AdTagEnhancementExperimentName = "AD_TAG_ENHANCEMENT"
	AdTagEnhancementControlVariant = "Old_tag"
	AdTagEnhancementTestVariant1   = "New_tag_P1"
	AdTagEnhancementTestVariant2   = "New_tag_P3"
)

const AdsP3OldTagData = `[{"id":"764c51c8-4a17-492f-9e95-095d2316b037","tag_name":"Sponsored","position":"P3","attachment":{"name":"P3 - Ad.png","path":"inventory/product/55a80cdc-a78d-4408-9913-9d89423e171f-P3_-_Ad.png","lottie_path":null,"size_in_bytes":null,"width":28,"hide_padding":null,"id":"159424cf-72c4-43af-938a-29ea5f300540","media_type":"IMAGE","video_path":null,"height":28},"pdp_attachment":null,"is_active":true}]`
const AdsP3NewTagData = `[{"id":"764c51c8-4a17-492f-9e95-095d2316b037","tag_name":"Sponsored","position":"P3","attachment":{"name":"P1 - Ad.png","path":"inventory/product/55a80cdc-a78d-4408-9913-9d89423e171f-P1_V2_-_Ad.png","lottie_path":null,"size_in_bytes":null,"width":16,"hide_padding":null,"id":"159424cf-72c4-43af-938a-29ea5f300540","media_type":"IMAGE","video_path":null,"height":14},"pdp_attachment":null,"is_active":true}]`
const AdsP1NewTagData = `[{"id":"764c51c8-4a17-492f-9e95-095d2316b037","tag_name":"Sponsored","position":"P1","style":{"top":6,"right":6},"attachment":{"name":"P1 - Ad.png","path":"inventory/product/55a80cdc-a78d-4408-9913-9d89423e171f-P1_V2_-_Ad.png","lottie_path":null,"size_in_bytes":null,"width":16,"hide_padding":null,"id":"159424cf-72c4-43af-938a-29ea5f300540","media_type":"IMAGE","video_path":null,"height":14},"pdp_attachment":null,"is_active":true}]`

type ConsequenceType string

const (
	BOOST_THEN ConsequenceType = "BOOST_THEN"
	BURY_THEN  ConsequenceType = "BURY_THEN"
	BOOST_WITH ConsequenceType = "BOOST_WITH"
	BURY_WITH  ConsequenceType = "BURY_WITH"
)

// inhouse search txn constants
const (
	InHouseResponseStatus  = "SearchPlatformResponseStatus"
	DeviceId               = "DeviceId"
	IpAddressWithPort      = "IpAddressWithPort"
	Host                   = "Host"
	RequestUri             = "RequestUri"
	AppVersion             = "AppVersion"
	RequestId              = "RequestId"
	PrimarySubcategoryName = "PrimarySubcategoryName"
	IsRelatedSearch        = "IsRelatedSearch"
	RelatedProductsCount   = "RelatedProductsCount"
)

const AdsPageSize = 8
const StoreProductFetchSizeFromPas = 45
const BannerFetchSizeFromPas = 20
const AdsFetchSizeFromPas = 15
const IdealStoreProductsInPage = 22
const InstockLimitForLayout = 2
const AdsFetchSizeForZeroSearchProducts = 45

const CBP = "cbp"
const (
	L3CategoryDetailBoostAttribute = "l3CategoriesDetail.name"
	L3SearchRelevance              = "l3SearchRelevance"
	DefaultRecall                  = "defaultRecall"
)

// Events
type EventName string

const (
	SearchResultsPage EventName = "search_results_page_event"
)

// screen names
const (
	AutoSuggestion = "AUTOSUGGESTION"
	SearchPage     = "SEARCH_PAGE"
	PreSearch      = "PRE_SEARCH"
)

// TOS priority constants
const (
	HighPriority = "high"
)

const (
	DSRelatedProductsCacheTimeInHours = 4
)

const (
	HTTPLibraryType = "HTTP"
)

const (
	SearchPageRecipeExperiment            = "SEARCH_PAGE_RECIPE_EXPERIMENT_NAME"
	SearchPageRecipePosition              = "SEARCH_PAGE_RECIPE_POS"
	IsRecipeWidgetEnabled                 = "IsRecipeWidgetEnabled"
	MinRecipeCount                        = "MIN_RECIPE_COUNT"
	RecipeWidgetViewMeta                  = "RECIPE_WIDGET_VIEW_META"
	RecipeWidgetWhitelistedL3Category     = "RECIPE_WIDGET_WHITELISTED_L3_CATEGORY"
	SearchAdsV3Experiment                 = "SEARCH_ADS_V3_EXPERIMENT"
	SearchAdsV3Enabled                    = "SEARCH_ADS_V3_ENABLED"
	SearchPcaBannerAdsV3Experiment        = "SEARCH_PCA_BANNER_ADS_V3_EXPERIMENT"
	PCAVariantTypeExperiment              = "SEARCH_PCA_VARIANT_ADS_V3_EXPERIMENT"
	SearchPcaBannerAdsV3Enabled           = "SEARCH_PCA_BANNER_ADS_V3_ENABLED"
	PCAVariantTypeEnabled                 = "SEARCH_PCA_VARIANT_ADS_V3_ENABLED"
	PcaWidgetMinProductThreshold          = "PCA_WIDGET_MIN_PRODUCT_THRESHOLD"
	PcaBackfillExperimentName             = "PCA_BACKFILL_EXPERIMENT"
	PcaBackfillEnabled                    = "PCA_BACKFILL_ENABLED"
	UnifiedSearchAdsExperiment            = "UNIFIED_SEARCH_ADS_EXPERIMENT"
	UnifiedSearchAdsEnabled               = "UNIFIED_SEARCH_ADS_ENABLED"
	CrossSellPcaExperiment                = "CROSS_SELL_PCA_EXPERIMENT"
	IsCrossSellPcaEnabled                 = "IS_CROSS_SELL_PCA_ENABLED"
	PcaBackfillThreshold                  = "PCA_BACKFILL_THRESHOLD"
	PcaBackfillTosThreshold               = "PCA_BACKFILL_TOS_THRESHOLD"
	PcaBackfillLogoUrl                    = "PCA_BACKFILL_LOGO_URL"
	PcaBackfillTitle                      = "PCA_BACKFILL_TITLE"
	PcaBackfillPrimarySubtitle            = "PCA_BACKFILL_PRIMARY_SUBTITLE"
	PcaBackfillSecondarySubtitle          = "PCA_BACKFILL_SECONDARY_SUBTITLE"
	PcaBackFillAdTagEligible              = "PCA_BACKFILL_AD_TAG_ELIGIBLE"
	PcaV2BannerTitleStyle                 = "PCA_V2_BANNER_TITLE_STYLE"
	PcaV2BannerPrimarySubtitleStyle       = "PCA_V2_BANNER_PRIMARY_SUBTITLE_STYLE"
	PcaV2BannerSecondarySubtitleStyle     = "PCA_V2_BANNER_SECONDARY_SUBTITLE_STYLE"
	PcaV2BannerLogoStyle                  = "PCA_V2_BANNER_LOGO_STYLE"
	PcaV2BannerSubtitleStyle              = "PCA_V2_BANNER_SUBTITLE_STYLE"
	UserRecommendationV2Experiments       = "user_recommendations_v2_experiment_names"
	FilteredAdsExperiment                 = "FILTERED_ADS_EXPERIMENT"
	PCAL3RelevancyExperiment              = "PCA_L3_EXPERIMENT"
	AreFilteredAdsEnabled                 = "FILTERED_ADS_ENABLED"
	LiquidationWidgetViewMeta             = "LIQUIDATION_WIDGET_VIEW_META"
	VipDealHorizontalCardWidgetViewMeta   = "VIP_DEAL_HORIZONTAL_CARD_VIEW_META"
	PrebuiltCartWidgetCompatibleComponent = "PRE_BUILT_CART_ENABLED"
	YmalL3FallbackEnabled                 = "YMAL_L3_FALLBACK_ENABLED"

	EmbeddingModelRTVectorSearchFlag                  = "embedding_model_rt_vector_search_flag"
	EmbeddingModelRTVectorSearchExperimentName        = "embedding_model_rt_vector_search_experiment_name"
	EmbeddingModelRTVectorSearchExperimentVariantName = "embedding_model_rt_vector_search_experiment_variant_name"
)

// dynamic config constants
const (
	IsAisleExperienceEnabled                     = "is_aisle_experience_enabled"
	AisleExperienceExperimentName                = "aisle_experience_experiment_name"
	IsSuperSaver3x3LayoutScaledUp                = "is_search_super_saver_3x3_layout_scaled_up"
	SuperSaver3x3LayoutExperimentName            = "search_supersaver_3x3_layout_experiment_name"
	SuperSaver3x3LayoutExperimentTestVariantName = "search_supersaver_3x3_layout_experiment_test_variant"
	MinAisleTabCount                             = "min_aisle_tab_count"
	MaxAisleTabCount                             = "max_aisle_tab_count"
	AisleHeaderLayoutCacheTTL                    = "aisle_header_layout_cache_ttl"
	ApplicableSearchModesForAisleExperience      = "applicable_search_modes_for_aisle_experience"
	ApplicableMarketPlaceForAisleExperience      = "applicable_market_place_for_aisle_experience"
	MinProductsInAisleTab                        = "min_products_in_aisle_tab"
	IsOtherBucketOverRpAdsEnabled                = "is_other_bucket_switch_over_rp_ads_enabled"
	TertiaryStoreDeboostingExperimentName        = "tertiary_store_deboosting_experiment"
	TertiaryStoreDeboostingEnabled               = "is_tertiary_store_deboosting_enabled"
	EnableTertiaryStoreDeboostingLogs            = "enable_tertiary_store_deboosting_logs"
	EnableOOSCheckForStoreSelection              = "enable_oos_check_for_store_selection"
	IsSearchWidgetRemovalEnabled                 = "is_search_widget_removal_enabled"
	ExploreExpansionExperimentName               = "explore_expansion_experiment_name"
	IsExploreExpansionEnabled                    = "is_explore_expansion_enabled"
	ExploreExpansionQueryFeatures                = "explore_expansion_query_features"
	ExploreExpansionExperimentVariantNames       = "explore_expansion_experiment_variant_names"
	IsEmptyWidgetEnableForFrontendLmsFlow        = "is_empty_widget_enable_for_frontend_lms_flow"
	DefaultSearchMarketplace                     = "default_search_marketplace"
)

// Aisle Experience constants
const (
	QueryCategorizationAisleExperienceType = "aisleExperience"
	Tabs                                   = "TABS"
	AisleAllTabName                        = "All Results"
	AisleAllTabDisplayName                 = " All \nResults"
	AisleL3CategoryDetail                  = "l3CategoriesDetail.id"
	AisleL2CategoryDetail                  = "primarySubcategoryId"
	Aisle                                  = "aisle"

	// not shown reason
	AisleExperienceFeatureFlagOff          = "AisleExperienceFeatureFlagOff"
	AisleExperienceExperimentNotEnabled    = "AisleExperienceExperimentNotEnabled"
	AisleExperienceSearchModeNotSupported  = "AisleExperienceSearchModeNotSupported"
	AisleExperienceMarketplaceNotSupported = "AisleExperienceMarketplaceNotSupported"
	AisleTabMinimumThresholdNotMet         = "AisleTabMinimumThresholdNotMet"
)

// Marketplace Type
const (
	SuperSaver = "SUPER_SAVER"
	ZeptoNow   = "ZEPTO_NOW"
)

// MatchType : Query matching buckets
const (
	Span      = "SPAN"
	Partial   = "PARTIAL"
	Other     = "OTHER"
	SuperSpan = "SUPERSPAN"
)

// MatchType: EXACT, PARTIAL, OTHERS
const (
	BucketExact   = "EXACT"
	BucketPartial = "PARTIAL"
	BucketOther   = "OTHER"
)

// Cross Sell Widgets
const (
	CrossSellWidgetsExpirationTimeInHour     = 30 * 24
	InMemoryCacheDefaultExpirationTimeInHour = 4
	InMemoryCacheCleanUpIntervalInHour       = 1

	CrossSellWidgetsMaxRecommendationPerWidget = 40
)

const (
	SuggestedForYouWidgetType  = "SUGGESTED_FOR_YOU"
	DiscoverNewPicksWidgetType = "DISCOVER_NEW_PICKS"
	CoBoughtWidgetMongoKey     = "PEOPLE_ALSO_LOOK_FOR"

	SuggestedForYouTitle             = "Because you looked for "
	SuggestedForYouTitleVariant      = "body4"
	SuggestedForYouTitleColor        = "#757C8D"
	SuggestedForYouSearchTermVariant = "heading8"
	SuggestedForYouSearchTermColor   = "#757C8D"

	OnPageLoadActionType                               = "LOAD_WIDGET_DATA"
	OnPageLoadActionPageKey                            = "SEARCH"
	OnPageLoadActionPillsWidget                        = "PILLS_WIDGET"
	OnPageLoadActionSearchSuggestionsWidgetName        = "SEARCH_SUGGESTIONS"
	OnPageLoadActionItemActionType                     = "RESET_SEARCH_QUERY"
	OnPageLoadActionSuggestedSearchesWidget            = "SUGGESTED_SEARCHES_WIDGET"
	OnPageLoadActionSearchSuggestionsOverlayWidgetName = "SEARCH_SUGGESTIONS_OVERLAY"
	ShowSearchSuperSaverCoachMark                      = "SHOW_SEARCH_SUPER_SAVER_COACHMARK"
	SearchToggleCoachMark                              = "search_toggle_coachmark"
	SuperSaverToggleText                               = "Switch to Super Saver mode and grab the lowest prices in your city"
	ToggleExploreButtonText                            = "Switch & Explore"
)

// Recall Info Constants
const (
	OrganicRecallWithNoTypoAndExactMatch = "ORGANIC_RECALL_WITH_NO_TYPO_AND_EXACT_MATCH"
	OrganicRecallWithTypo                = "ORGANIC_RECALL_WITH_TYPO"
	OrganicRecallWithPartialMatch        = "ORGANIC_RECALL_WITH_PARTIAL_MATCH"
	NonOrganicRecallWithRelatedRecall    = "NON_ORGANIC_RECALL_WITH_RELATED_RECALL"
	ZeroRecall                           = "ZERO_RECALL"
	SingleOrganicWithRelatedRecall       = "SINGLE_ORGANIC_WITH_RELATED_RECALL"
	ScopedSearchFirstFeedRecall          = "SCOPED_CATEGORY_FIRST_FEED_RECALL"
)

// Pharma constants
const (
	// QueryFilters
	SaltNameAndCompositionsSearchFilter = "saltNameAndCompositions"
	MedicineTypeSearchFilter            = "medicineType"
	DosageFormSearchFilter              = "dosageForm"
	DrugTypeSearchFilter                = "drugType"

	// Query intent type
	Medicine = "medicine"
	Ailment  = "ailment"
	Salt     = "salt"
	Others   = "others"

	// Search query description error reasons
	PharmaQueryDescriptionNotFound                                   = "PharmaQueryDescriptionNotFound"
	QueryDescriptionWidgetCompatibleComponentNotFoundInRequestHeader = "QueryDescriptionWidgetCompatibleComponentNotFoundInRequestHeader"
)

// L4 Attribute Pharma Identifiers
const (
	PVCompositionWithStrength = "composition with strength"
	PVAilment                 = "ailment"
	PVDosageForm              = "dosage form"
	PVMedicineType            = "medicine type"
	PVDrugType                = "drug type"
	PVColour                  = "colour"
)

// Cron Job constants
const (
	// meds with similar salts cron
	OFFSET                = "offset"
	LIMIT                 = "limit"
	OFFSET_INITIAL_VALUE  = 0
	MEDICINE_QUERY_INTENT = "medicine"
	TIMEOUT               = "timeout"
)

// Search Query Descriptions SubHeader Constant
const (
	SubHeaderSaltsOrComposition = "Salt / Composition"
	SubHeaderMedicine           = "Medicine"
	SubHeaderAilment            = "Ailment"
)

const (
	IsCategoryAdsV3Enabled = "IsCategoryV3AdsEnabled"
	PcaV2Component         = "PCA_V2"
	PcaV2Phase2Component   = "PCA_V2:01"
	PcaWidgetPosition      = "PCA_WIDGET_POSITION"
)

const (
	CommaSeparator = ","
	EmptyString    = ""
)

const (
	BlockedSubCategoryIDs = "BLOCKED_SUBCATEGORY_IDS"
)

const CDPClientName = "USER-SEARCH-SERVICE"

// Liquidation Not Done Reasons
const (
	LiquidationExperimentNotFound            = "LiquidationExperimentNotFound"
	RelevanceCategoryOrSubcategoryNotPresent = "RelevanceCategoryOrSubcategoryNotPresent"
	LiquidationTimeSlotMismatch              = "LiquidationTimeSlotMismatch"
	OrganicProductsLessThan4                 = "OrganicProductsLessThan4"
	LiquidationMinProductsCriteriaFailed     = "LiquidationMinProductsCriteriaFailed"
	LiquidationL1NotAllowed                  = "LiquidationL1NotAllowed"
	LiquidationOnDefaultL1Category           = "default"
)

// Vip Deal Product Card Not Shown Reason
const (
	VipDealProductCardMinProductsCriteriaFailed = "VipDealProductCardMinProductsCriteriaFailed"
)

// Tenant Types Header
const (
	TenantZeptoCafe = "ZEPTOCAFE"
	TenantZepto     = "ZEPTO"
	TenantZeptoLoot = "ZEPTOLOOT"
)

const (
	ScopeSearchPrimaryCategory       = "primaryCategoryId"
	ScopedSearchZeptoCafeDisplayName = "Zepto Cafe"
)

const (
	FetchProcessedSearchTermsForOriginalSearchTerm = "fetch_processed_search_terms_for_original_search_term"
	GetUserSearchRelevanceForQueryFromCache        = "get_user_search_relevance_for_query_from_cache"
	GetUserSearchRelevanceForQueryFromMongo        = "get_user_search_relevance_for_query_from_mongo"
	GetFilterAttachmentFromMongo                   = "get_filter_attachment_from_mongo"
	GetCrossSellWidgetDetailsFromMongo             = "get_cross_sell_widget_details_from_mongo"
	GetCrossSellAttachmentsByAttachmentIdFromMongo = "get_cross_sell_attachments_by_attachment_id_from_mongo"
	GetQueryCategorizationByQueryFromMongo         = "get_query_categorization_by_query_from_mongo"
	GetPharmaQueryDescriptionDetailsFromMongo      = "get_pharma_query_description_details_from_mongo"
	GetQueryTermRedirectionFromMongo               = "get_query_term_redirection_from_mongo"
	GetOpsOverriddenRelevanceQueriesFromMongo      = "get_ops_overridden_relevance_queries_from_mongo"
)

// Analytics Worker Event Action
type Action string

const (
	Update Action = "UPDATE"
)

// Title Widgets
const (
	ExactMatchesWidgetTitle                      = `Showing results for "<b>%s<b>"`
	PartialMatchesWidgetTitle                    = "Similar Products"
	OtherSimilarWidgetsTitle                     = "More to Explore"
	OtherSimilarWidgetsTitle_ZeroPartialProducts = `Showing related products for "<b>%s<b>"`
)

// Widget Names
const (
	OrganicOosWidgetName        = `ORGANIC_OOS_SEARCH_WIDGET`
	RelatedProductWidgetName    = `RELATED_PRODUCT_WIDGET`
	PartialOrganicWidgetName    = `PARTIAL_ORGANIC_PRODUCT_WIDGET`
	LayoutStartWidgetName       = "SEARCH_TITLE_WIDGET"
	ProductGridWidgetNamePrefix = "SEARCHED_PRODUCTS_"
)

// OBD Banners constants
const (
	DoorstepCancellation                = "Doorstep Cancellation"
	ObdBannerDoorStepCancellation       = "obd_banner_door_step_cancellation"
	ObdBannerDoorStepCancellationPlpExp = "obd_banner_door_step_cancellation_plp_exp"
	TryAndBuy                           = "Try And Buy"
	ObdBannerTryAndBuy                  = "obd_banner_try_and_buy"
	ObdBannerTryAndBuyPlpExp            = "obd_banner_try_and_buy_plp_exp"
	OpenBoxDelivery                     = "Open Box Delivery"
	ObdBannerOpenBoxDelivery            = "obd_banner_open_box_delivery"
	ObdBannerOpenBoxDeliveryPlpExp      = "obd_banner_open_box_delivery_plp_exp"
)

const (
	Default = "default"
	Explore = "explore"
	Vector  = "vector"
)

const (
	Zero  = 0
	One   = 1
	Two   = 2
	Three = 3
	Ten   = 10
)

// User Recommendation Types
const (
	BuyAgainV2RecommendationType                   = "BUY_AGAIN_V2"
	GodProductsRecommendationType                  = "GOD_PRODUCTS"
	BuyAgainV2AndGodProductsRecommendationType     = "BUY_AGAIN_V2_AND_GOD_PRODUCTS"
	BuyAgainV2WithPredictiveCartRecommendationType = "BUY_AGAIN_V2_WITH_PREDICTIVE_CART"
)

const (
	L1 = "L1"
	L3 = "L3"
	L2 = "L2"
)

// Configurable SRP slot boosting experiment parameters
const (
	NumberOfSlotsToBoost = "no_of_slots_to_boost"
	RecallScreeningLimit = "recall_screening_limit"
	ShowForYouTag        = "show_for_you_tag"
	Filtering            = "filtering"
	PriorityOrder        = "priority_order"
)

const (
	ZeroGodProductsInRecall                       = "Zero God Products found in Recall"
	MaxPreviouslyBoughtBoostedProductsCountReason = "Previously Bought Boosted Products are 2"
)

const (
	ControlVariant = "control"
	TestVariant    = "test"
)

const (
	Rpi           = "rpi"
	Atcpi         = "atcpi"
	Vdr           = "vdr"
	Ssfr          = "ssfr"
	Rank5         = "rank5"
	Rank6         = "rank6"
	Rank7         = "rank7"
	Rank8         = "rank8"
	Rank9         = "rank9"
	Rank10        = "rank10"
	Rank11        = "rank11"
	Rank12        = "rank12"
	Rank13        = "rank13"
	Rank14        = "rank14"
	Rank15        = "rank15"
	Rank16        = "rank16"
	Rank17        = "rank17"
	Rank18        = "rank18"
	Rank19        = "rank19"
	Rank20        = "rank20"
	Rank21        = "rank21"
	Rank22        = "rank22"
	Rank23        = "rank23"
	Rank24        = "rank24"
	Rank25        = "rank25"
	Rank26        = "rank26"
	Rank27        = "rank27"
	Rank28        = "rank28"
	Rank29        = "rank29"
	Rank30        = "rank30"
	SemanticRank1 = "semantic_rank1"
	SemanticRank2 = "semantic_rank2"
	SemanticRank3 = "semantic_rank3"
	SemanticRank4 = "semantic_rank4"
	SemanticRank5 = "semantic_rank5"
)

// Cohort Reranking based experiment key
const (
	CategoryCohortBasedExperimentKey                       = "category_cohort_based_ranking_experiment"
	CategoryCohortE5ExperimentKey                          = "category_cohort_based_ranking_experiment_e5"
	CategoryCohortE5ExperimentKeySuperSaver                = "category_cohort_based_ranking_experiment_e5_ss"
	CategoryCohortE5ExperimentKeyZeptoNow                  = "category_cohort_based_ranking_experiment_e5_zn"
	CategoryCohortBasedGenderExperimentKey                 = "category_cohort_based_ranking_experiment_gender"
	CategoryCohortE5GenderExperimentKey                    = "category_cohort_based_ranking_experiment_e5_gender"
	CategoryCohortE5GenderExperimentKeySuperSaver          = "category_cohort_based_ranking_experiment_e5_ss_gender"
	CategoryCohortE5GenderExperimentKeyZeptoNow            = "category_cohort_based_ranking_experiment_e5_zn_gender"
	CategoryCohortE5NightExperimentKeySuperSaver           = "category_cohort_based_ranking_experiment_e5_ss_night"
	CategoryCohortE5NightExperimentKeyZeptoNow             = "category_cohort_based_ranking_experiment_e5_zn_night"
	CategoryCohortE5ExperimentKeySuperSaverSearchInsession = "category_cohort_based_ranking_experiment_e5_ss_search_insession"
	CategoryCohortE5ExperimentKeyZeptoNowSearchInsession   = "category_cohort_based_ranking_experiment_e5_zn_search_insession"

	CategoryCohortE5GenderExperimentKeyZeptoNowUpsFallback   = "category_cohort_based_ranking_experiment_e5_zn_gender_ups"
	CategoryCohortE5GenderExperimentKeySuperSaverUpsFallback = "category_cohort_based_ranking_experiment_e5_ss_gender_ups"

	SubcategoryCohortE5GenderExperimentKeyZeptoNowUpsFallback   = "category_cohort_based_ranking_experiment_e5_zn_gender_ups_l2"
	SubcategoryCohortE5GenderExperimentKeySuperSaverUpsFallback = "category_cohort_based_ranking_experiment_e5_ss_gender_ups_l2"

	//Remarks
	UserProfileScoreFallbackTriggered       = "User Profile Score Fallback Triggered for the user"
	MarketplaceFallbackTriggerred           = "Marketplace Fallback Triggered for the user"
	GenderFallbackTriggered                 = "Gender Fallback Triggered for the user"
	MissingRelevancePropertiesInCdpResponse = "Missing relevance properties in cdp response"
	MissingCohortTypeInCdpRelevanceProperty = "Missing cohort type in cdp relevance property"
	FallbackRelevanceMissing                = "Fallback relevance level %s missing in relevance info"
	FallbackRelevanceMappingMissing         = "Fallback relevance mapping missing in relevance info"
	FallbackRelevanceLevelTriggerred        = "Fallback relevance level %s triggerred for the user"
	FallbackClmRankerExperimentTriggerred   = "Fallback Clm Ranker Experiment Triggerred for the user"

	//Cohort types
	Entry       = "Entry"
	Premium     = "Premium"
	FTB         = "FTB"
	Mass        = "Mass"
	MassPremium = "Mass-Premium"

	//Relevance Levels
	Category    = "L1"
	Subcategory = "L2"
)

const (
	Invalid = "INVALID"
)

const (
	BentoWidgetUnlistedSubcategoryId = "subcategory_id"
	BentoWidgetTitle                 = "title"
	BentoWidgetDescription           = "description"
	BentoWidgetDeepLink              = "widget_deep_link"
)

// theme not shown reasons
const (
	ThemeNotShownMasterDynamicFlagSetToFalse                  = "Theme Master Dynamic Flag Set to False"
	ThemeNotShownMasterGPTFlagSetToFalse                      = "GPT Theme Master Dynamic Flag Set to False"
	ThemeNotShownCompatibleComponentMissingInRequest          = "THEME_BASED_EXPERIENCE Compatible Component Missing"
	ThemeNotShownSearchModeBrand                              = "Search Mode Brand"
	ThemeNotShownThemeBasedSeeAllPage                         = "Theme Based See All Page"
	ThemeNotShownQueryCategorizationDetailEmpty               = "Query Categorization Details Empty"
	ThemeNotShownNoDocumentsFound                             = "No Documents Found In Mongo"
	ThemeNotShownThemeBuyingExperimentationDetailEmpty        = "Theme Based Buying Experimentation Details Empty"
	ThemeNotShownThemeBuyingExperimentationTestVariantInvalid = "Theme Based Buying Experimentation Test Variant Invalid"
	ThemeNotShownExperienceEnabledFlagIsFalse                 = "themeBasedExperienceShown field Is false"
	ThemeNotShownGptExperienceEnabledFlagIsFalse              = "gptThemeExperienceShown field Is false"
	ThemeNotShownInvalidTenantType                            = "Theme Based Buying Invalid Tenant Type"
	ThemeNotShownScopedSearchContext                          = "Theme Based Buying Scoped Search Context is not nil"
	ThemeNotShownPlatformNotAndroidOrIos                      = "Theme Based Buying Platform Not android Or ios"
	ThemeNotShownSearchProductsAtlasEmpty                     = "Theme Based Buying Search Products Atlas Empty"
	ThemeNotShownOverrideToQueryIntent                        = "overrideToQueryIntentDueToLesserCarouselsFormation"
	ThemeNotShownPASEnrichmentFailed                          = "Theme Based Buying PAS Enrichment Failed"
	ThemeNotShownFTBUserInPLPVariant                          = "Theme Based Not Shown as user in FTB PLP Bucket"
	ThemeNotShownAsGPTThemeActive                             = "L3 Theme Not Shown as GPT Theme is active"

	GPTThemeNotShownFieldFlagFalse       = "GPT Theme Not Shown as categorization gptThemeExperienceShown is false"
	GPTThemeNotShownExperimentNotApplied = "GPT Theme Not Shown as user not in test bucket"
	GPTThemeNotShownDataAbsentFallback   = "GPT Theme Not Shown as fallback to SRP engaged because of missing data"
)

// model ranker constants
const (
	ApplicationVersion         = "app_version"
	DeviceBrand                = "device_brand"
	DeviceOs                   = "device_os"
	Gender                     = "gender"
	CategoryTier               = "category_tier"
	ProductVariantIds          = "product_variant_ids"
	CityID                     = "city_id"
	SearchTerm                 = "search_term"
	StoreID                    = "store_id"
	Month                      = "month"
	DayOfWeek                  = "day_of_week"
	Hour                       = "hour"
	IsWeekend                  = "is_weekend"
	Male                       = "male"
	Female                     = "female"
	OtherGender                = "others"
	MassCategoryTier           = "mass"
	PremiumCategoryTier        = "premium"
	EntryCategoryTier          = "entry"
	MassPremiumCategoryTier    = "mass-premium"
	CityNameNeuronModelRequest = "city_name"
	CityIntents                = "city_intents"
	DbMatchingScore            = "db_matching_score"
	PartialQuery               = "partial_query"
	IntentId                   = "intent_id"
)

var ValuesToBeSkippedForFilter = []string{"false"}

const (
	// compression algorithms
	Gzip   = "gzip"
	Lz4    = "lz4"
	Snappy = "snappy"
	Zstd   = "zstd"
)

const ProductLeafnameL4 = "product_leafname"
const OtherPredcitedL4 = "other"
const StopWordPredictedL4 = "stop_word"

// external client headers
const (
	XSourceIdentifier = "x-source-identifier"
)

// Time Zones
const (
	IndianTimeZone = "Asia/Kolkata"
)

// Experiment Prefixes
const (
	FTbExperimentPrefix                = "Search_reranking_FTB"
	StoreStressRankingExperimentPrefix = "demand_shaping"
)

// stress ranking not done reasons
const (
	UserNotEnabledForMarketplace                  = "user not enabled for marketplaceType: %v"
	UserBelongsToVariant                          = "user belongs to %v ranking variant"
	StoreStressLevelNotFoundForStoreInCache       = "store stress level is not found in redis for store"
	StoreStressVariantMappingNotFoundInExperiment = "store stress-ranking variant mapping is not found in experiment"
	StoreStressRankingNotEnabledForRankerV2       = "User belongs to ranker v2, stress ranking variant doesnt belong to applicable variants for ranker v2 flow. Thus fallback to ranker v1 flow"
	NoStoreStressMapping                          = "Stress Mapping for the %v store is NO_STRESS, thus stress mapping not applied"
)

var AttributesForAttributeUnderstandingMeta = []string{"price", "size", "weight", "colour family", "gender"}
var ProductLeafNameForAttributeUnderstandingMeta = "product_leafname"
var BrandForAttributeUnderstandingMeta = "brand"

const (
	VIP_DEALS = "vip_deals"
)

// Vip Deal Product Tile states
const (
	AtlasBoosted          = 1
	AtlasBoostedCardShown = 2
	CardShown             = 3
)

const (
	SearchInsessionRankingExperiments             = "search_insession_ranking_experiment"
	SearchInsessionRankingExperimentKeySuperSaver = "search_insession_ranking_experiment_ss"
	SearchInsessionRankingExperimentKeyZeptonow   = "search_insession_ranking_experiment_zn"
)

const (
	L3_BASED_THEMES  = "L3_BASED_THEMES"
	GPT_BASED_THEMES = "GPT_BASED_THEMES"
)

const (
	GPT_SUBTHEME_TITLE_ATTRIBUTE_NAME = "gpt_subtheme_title"
)

// Constants for External Methods
const (
	BulkIndexDocumentsMethod = "BulkIndexDocuments"
)

// StatusCodes for Error Tracking
const (
	SuccessResponse = "success"
)

// Platforms
const (
	Android = "android"
	IOS     = "ios"
	Web     = "web"
)

const (
	PriorityHigh = "high"
)

const (
	RequestIdHeaderKey          = "requestId"
	RequestIdHeaderKeySnakeCase = "request_id"
	XRequestIdHeaderKey         = "x-request-id"
	XDebugLogEnabledHeaderKey   = "x-debug-log-enabled"
	XIntentIdHeaderKey          = "x-intent-id"
)
const SamplingPercentageMax = 100

const (
	EventStartDateSuffix = "start_date"
	EventEndDateSuffix   = "end_date"
	EventWidgetPosition  = "widget_position"
)

const (
	BrowseLiquidationRankerOne = "liquidation_ranker_1"
	BrowseLiquidationRankerTwo = "liquidation_ranker_2"
)

const (
	EXPECTED_CSV_COLUMN_COUNT_RELEVANCE_BULK_UPLOAD = 6
)

const (
	SearchPageWidgetName = "SRP"
)

const (
	Subcategories = "subcategories"
	SearchFilters = "search_filters"
	L3Details     = "l3_details"
)

const (
	FetchModeLocal = "local"
)

// redis PIPE constants
const (
	PipeSet = "Pipe SET"
	Get     = "GET"
)

// Weighted Autosuggestion Ranker Constants
const (
	NormalizedRelevanceScore = "normalizedRelevanceScore"
	NormalizedCtrCvrScore    = "normalizedCtrCvrScore"
	Weight1                  = "weight1"
	Weight2                  = "weight2"
)
const (
	TAIL_OF_TAIL             = "TAIL_OF_TAIL"
	RT_VECTOR_SEARCH         = "RTVS"
	HEAD_QUERY_FUZZY_REMOVAL = "HEAD"
)

const (
	FashionRevampPasHeaderKey   = "fashion_experience_header"
	FashionRevampPasHeaderValue = "present"
)

const (
	AdsBannerKeyName = "BANNER"
)

const (
	AdsBannerCandidate1 = "candidate1"
	AdsBannerCandidate2 = "candidate2"
	AdsBannerCandidate3 = "candidate3"
	AdsBannerCandidate4 = "candidate4"
)

// Ads Banner Widget Names
const (
	AdsBannerTopWidgetName    = "ADS_BANNER_TOP_PREMIUM_PROMO"
	AdsBannerBottomWidgetName = "ADS_BANNER_BOTTOM_BANNER_GRID"
)

// ReRanking not Done reasons
const (
	RankingNotEnabled = "RankingNotEnabled"
)

// Explore Exploit constants
const (
	NotShownAsExperimentBucketNotTest         = "user_not_in_experiment_test_bucket"
	NotShownAsRelevanceQuerySegmentMismatched = "query segment not valid for explore exploit in segmentations"
	NotShownAsRankerV1FlowQRL                 = "qrl_v1_flow"
	NotShownAsFeatureStoreNotTest             = "user_not_in_feature_store_test_bucket"
	NotShownAsGlobalFlagIsFalse               = "explore_exploit_master_flag_is_false"

	// Experiment Field Values
	MaxMergeThresholdRatio = "maxMergeRatio"
	MinMergeThresholdRatio = "minMergeRatio"
	Offset                 = "offset"
	RankingVariant         = "rankVariant"

	PositionChange          = "pos_change"
	LexicalNormalizedScore  = "lexical_normalized_score"
	SemanticNormalizedScore = "semantic_normalized_score"

	// buckets for merge ratio
	ExploreExact   = "explore_exact"
	ExplorePartial = "explore_partial"
	ExploreOthers  = "explore_others"

	ExploitExact   = "exploit_exact"
	ExploitPartial = "exploit_partial"
	ExploitOthers  = "exploit_others"
)

const (
	DisableChecksInQueryRedirectionForSpecialCampaigns = "disable_checks_in_query_redirection_for_special_campaigns"
)

// Explore Expansion
const (
	NotShownAsRelevanceQueryFeaturesMismatched = "query features not valid for explore expansion in features"
	NotShownAsExperimentCheckFailed            = "experiment check failed"
)
