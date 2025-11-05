package service_constants

const (
	RawSearchKey = "rawSearchKey"
	IndexName    = "indexName"

	GrpcMetadataUserId       = "user_id"
	GrpcMetadataRequestID    = "X_request_id"
	GrpcMetadataDebugRequest = "debug"
	AutoSuggestionIndexName  = "autosuggestions_3"
	SearchServiceName        = "search-service"
	CanaryValue              = "canary"

	GetAutoSuggestionsFromInhouseSearchClient = "GetAutoSuggestionsFromInhouseSearchClient"
	GetAutoSuggestionsFromInhouseSearch       = "GetAutoSuggestionsFromInhouseSearch"

	GetSearchProductsFromAlgoliaOrInHouseSearch              = "GetSearchProductsFromAlgoliaOrInHouseSearch"
	GetSearchProductsFromInHouseSearchForSuperStore          = "GetSearchProductsFromInHouseSearchForSuperStore"
	GetSearchResultsFromInhouseSearch                        = "GetSearchResultsFromInhouseSearch"
	GetSearchResultsFromInhouseSearchForSuperStore           = "GetSearchResultsFromInhouseSearchForSuperStore"
	GetSearchResultsFromInHouseSearchForPharmacySaltMappings = "GetSearchResultsFromInHouseSearchForPharmacySaltMappings"
	GetSPIdResultsFromInHouseSearchForPVId                   = "GetSPIdResultsFromInHouseSearchForPVId"
	GetSPIdResultsFromInHouseSearchForFilters                = "GetSPIdResultsFromInHouseSearchForFilters"
	GetScopedSearchResultsForSuperStore                      = "GetScopedSearchResultsForSuperStore"
	ReorderScopedSearchItems                                 = "ReorderScopedSearchItems"
	RuleContextsFieldKey                                     = "ruleContexts"
	StoreId                                                  = "storeId"
	PrimarySubcategoryName                                   = "primarySubcategoryName"
	Brand                                                    = "product.brand"
	BrandId                                                  = "product.brandId"
	SearchFiltersProductBrandIdKey                           = "brandIds"
	PrimaryCategoryName                                      = "primaryCategoryName"
	PrimaryCategoryId                                        = "primaryCategoryId"
	PrimarySubCategoryId                                     = "primarySubcategoryId"
	L3CategoriesDetailName                                   = "l3CategoriesDetail.name"
	L3CategoriesDetailId                                     = "l3CategoriesDetail.id"
	Properties                                               = "properties"
	ProductVariantId                                         = "productVariant.id"
	Enabled                                                  = "enabled"
	True                                                     = "true"
	CityID                                                   = "cityId"
	L4AttributeOpensearchPrefix                              = "l4Attributes"
	BuildWidgetSegmentName                                   = "BuildWidget"
	AddImageToWidgetSegmentName                              = "AddImageToWidget"

	ScopedCategory     = "ScopedCategory"
	ScopedRestCategory = "ScopedRestCategory"

	ConsumeSearchDataIntoInHouseSearch                       = "ConsumeSearchDataIntoInHouseSearch"
	SearchDataConsumerIndexingDlqTransaction                 = "SearchDataConsumerIndexingDlqTransaction"
	SearchDataIndexingProducerTransaction                    = "SearchDataIndexingProducerTransaction"
	BulkSearchDataIndexingDlqTransaction                     = "BulkSearchDataIndexingDlqTransaction"
	CmsConsumerDlqTransaction                                = "CmsConsumerDlqTransaction"
	QueryCategorizationConsumerDlqTransaction                = "QueryCategorizationConsumerDlqTransaction"
	FilterAttachmentConsumerDlqTransaction                   = "FilterAttachmentConsumerDlqTransaction"
	AlgoliaRequestDlqTransaction                             = "AlgoliaRequestDlqTransaction"
	SearchDataConsumerStoreProductPropertiesTransaction      = "SearchDataConsumerStoreProductPropertiesTransaction"
	BrowseSPRankingEventDlqTransaction                       = "BrowseSPRankingEventDlqTransaction"
	BrowseSPRankingEventProducerTransaction                  = "BrowseSPRankingEventProducerTransaction"
	BrowseTagsEventDlqTransaction                            = "BrowseTagsEventDlqTransaction"
	BrowseTagsEventProducerTransaction                       = "BrowseTagsEventProducerTransaction"
	CrossSellWidgetsEventDlqTransaction                      = "CrossSellWidgetsEventDlqTransaction"
	SearchQueryDescriptionUpdatesConsumerDlqTransaction      = "SearchQueryDescriptionUpdatesConsumerDlqTransaction"
	CmsPVL4AttributeConsumerTransaction                      = "CmsPVL4AttributeConsumerTransaction"
	CmsBundlePVL4AttributeConsumerTransaction                = "CmsBundlePVL4AttributeConsumerTransaction"
	CmsPVL4AttributeConsumerDlqTransaction                   = "CmsPVL4AttributeConsumerDlqTransaction"
	CmsL4AttributeConsumerInternalEventProcessingTransaction = "CmsL4AttributeConsumerInternalEventProcessing"
	UserSearchAnalyticsDlqTransaction                        = "UserSearchAnalyticsDlqTransaction"
	UserSearchAnalyticsV2DlqTransaction                      = "UserSearchAnalyticsV2DlqTransaction"
	StoreStressDlqTransaction                                = "StoreStressDlqTransaction"
	SearchPageMetricsProducerTransaction                     = "SearchPageMetricsProducerTransaction"
	PricingFeatureStoreDlqTransaction                        = "PricingFeatureStoreDlqTransaction"

	SearchFiltersL3CategoriesDetailNameKey = "l3CategoriesDetail.name"
	SearchFiltersL3CategoriesDetailIdKey   = "l3CategoriesDetail.id"
	SearchFiltersProductBrandKey           = "product.brand"
	SearchFiltersPropertiesKey             = "properties"
	SearchFiltersSaltNameAndCompositionKey = "pharmaSaltNameAndComposition"
	SearchFiltersAilmentsKey               = "pharmaAilments"
	SearchFiltersMedicineTypeKey           = "pharmaMedicineType"
	SearchFiltersDosageFormKey             = "pharmaDosageForm"
	SearchFiltersNoOfSalts                 = "pharmaNoOfSalts"
	SearchFiltersDrugType                  = "pharmaDrugType"
	SearchFiltersSourceFeed                = "sourceFeed"
	QclExperimentTestVariantAbAlias        = "Inhouse_query_processing_layer@test"

	ObjectID = "objectID"
	Id       = "id"

	//CDP Properties
	CategoryTier     = "categoryTier"
	Gender           = "gender"
	GenderMale       = "male"
	GenderFemale     = "female"
	GenderOthers     = "others"
	UserProfileScore = "userProfileScore"

	SearchOperationsDeleteDocuments = "SearchOperationsDeleteDocuments"
	SearchOperationsGetTaskStatus   = "SearchOperationsGetTaskStatus"
	SearchOperationsCancelTask      = "SearchOperationsCancelTask"
)

// Operator constants defined in search platform
const (
	Equal    = "equal"
	NotEqual = "not_equal"
	AndEqual = "and_equal"
)

const (
	DataDeletedSuccessfullyMessage   = "Data deleted successfully"
	NewL4MigratedSuccessfullyMessage = "New L4Migrated successfully"
)

const (
	YourPicksMaxProduct = 30
)

// PCA Ads banner constants
const (
	PREMIUM_CONTEXTUAL_ADS_V2_WIDGETID       = "PREMIUM_CONTEXTUAL_ADS_V2"
	PCA_SPONSORED_BRANDS_WIDGETNAME          = "PCA_SPONSORED_BRANDS"
	PREMIUM_CONTEXTUAL_WIDGET_WIDGETID       = "PREMIUM_CONTEXTUAL_WIDGET"
	PREMIUM_CONTEXTUAL_ADS_SEARCH_WIDGETNAME = "PREMIUM_CONTEXTUAL_ADS_SEARCH"
)
