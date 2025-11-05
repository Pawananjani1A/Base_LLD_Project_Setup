package repo_constants

const (
	TopSearch          = "top_search"
	TopSearchHistory   = "top_search_history"
	QclIgnore          = "qcl_ignore"
	QclIgnoreHistory   = "qcl_ignore_history"
	QclMongoPstValue   = "pst_value"
	QclMongoPstValueV2 = "pst_value_v2"
	Created            = "CREATED"
	Updated            = "UPDATED"
	Deleted            = "DELETED"

	Offset          = "offset"
	Limit           = "limit"
	QueryIntentType = "queryIntentType"
	CreatedOn       = "createdOn"
	Timeout         = "timeout"

	// below are mongo collection names
	CorrectedTerms                          = "corrected_terms"
	CollectionNameSubCategories             = "subcategories"
	CollectionNameQueryCategorization       = "query_categorization"
	CollectionNameFilterAttachments         = "filter_attachments_v2"
	CollectionSearchQueryDescriptions       = "search_query_descriptions"
	CollectionNameUserSearchRelevance       = "user_search_relevance"
	CollectionNameCrossSellWidgets          = "cross_sell_widgets"
	CollectionNameAttachments               = "attachments"
	CollectionNameQueryIntentConfig         = "query_intent_config"
	CollectionQueryToCalculatedSaltMappings = "query_to_calculated_salt_mappings"
	CollectionQueryRedirection              = "query_redirection"

	// collection fields for filter attachments
	FilterAttachmentQueryField     = "query"
	FilterAttachmentTypeField      = "filterType"
	FilterAttachmentTypeL3Category = "l3Category"

	CrossSellFilterImagesIdField = "_id"

	// collection fields for search_query_descriptions
	FilterSearchQueryDescriptionsIdField = "_id"
	PharmaSearchDescriptionsEnabled      = "enabled"
	BulkSaveUserSearchRelevance          = "BulkSaveUserSearchRelevance"

	// model fields
	OstKeySnakeCase    = "ost_key"
	OstKeyCamelCase    = "ostKey"
	PstValueSnakeCase  = "pst_value"
	PstValueCamelCase  = "pstValue"
	UpdatedBySnakeCase = "updated_by"
	UpdatedAtSnakeCase = "updated_at"
	UpdatedAtCamelCase = "updatedAt"
	CreatedOnSnakeCase = "created_on"
	CreatedOnCamelCase = "createdOn"
	IsActiveSnakeCase  = "is_active"
	IsActiveCamelCase  = "isActive"

	// filter attachments
	FilterQueryToCalculatedSaltMappingsIdField = "_id"

	// query categorization
	QueryCategorizationIdField      = "_id"
	QueryCategorizationEnabledField = "enabled"

	// mongo collection cross sell widgets
	CollectionCrossSellWidgetsFieldQuery = "query"
)

// qcl methods
const (
	IsTopSearchExist                                                 = "IsTopSearchExist"
	IsQclIgnoreExist                                                 = "IsQclIgnoreExist"
	SaveTopSearchTerms                                               = "SaveTopSearchTerms"
	GetAllTopSearchTerms                                             = "GetAllTopSearchTerms"
	GetAllQclIgnoreTerms                                             = "GetAllQclIgnoreTerms"
	SaveHistory                                                      = "SaveHistory"
	DeleteTopSearchTerms                                             = "DeleteTopSearchTerms"
	SaveQclIgnoreTerms                                               = "SaveQclIgnoreTerms"
	DeleteQclIgnoreTerms                                             = "DeleteQclIgnoreTerms"
	GetAllSubCategories                                              = "GetAllSubCategories"
	GetSubcategoryByIdWithProjections                                = "GetSubcategoryByIdWithProjections"
	GetAllContainsTobaccoSubCategoriesService                        = "GetAllContainsTobaccoSubCategoriesService"
	SaveFilterAttachmentDetails                                      = "SaveFilterAttachmentDetails"
	GetFilterImageByFilterId                                         = "GetFilterImageByFilterId"
	GetFilterImageByFilterIds                                        = "GetFilterImageByFilterIds"
	DeleteFilterAttachment                                           = "DeleteFilterAttachment"
	SaveSubCategories                                                = "SaveSubCategories"
	SaveQueryCategorizationDetails                                   = "SaveQueryCategorizationDetails"
	GetTobaccoSubcategoryById                                        = "GetTobaccoSubcategoryById"
	GetQueryCategorizationByQuery                                    = "GetQueryCategorizationByQuery"
	SaveTobaccoSubCategory                                           = "SaveTobaccoSubCategory"
	DeleteSubCategory                                                = "DeleteSubCategory"
	DeleteQueryCategorization                                        = "DeleteQueryCategorization"
	AddDataToBatchForProcessing                                      = "AddDataToBatchForProcessing"
	IndexSearchKafkaMessage                                          = "IndexSearchKafkaMessage"
	IndexSearchEvent                                                 = "IndexSearchEvent"
	HandleExploreTagsOnCityIdPVIdUpdateEvent                         = "HandleExploreTagsOnCityIdPVIdUpdateEvent"
	GetAllStoresBasicDetailsWithFilter                               = "GetAllStoresBasicDetailsWithFilter"
	AddKafkaMessageToBatchForProcessing                              = "AddKafkaMessageToBatchForProcessing"
	ProcessSearchDocuments                                           = "processSearchDocuments"
	PostExtractionCallbackKafka                                      = "PostExtractionCallbackKafka"
	DoesOriginalSearchTermExistInDB                                  = "DoesOriginalSearchTermExistInDB"
	FetchAllExistingSpellCorrectionSearchTerm                        = "FetchAllExistingSpellCorrectionSearchTerm"
	FetchProcessedSearchTermsForOriginalSearchTerm                   = "FetchProcessedSearchTermsForOriginalSearchTerm"
	ToggleSpellCorrectionSearchTerm                                  = "ToggleSpellCorrectionSearchTerm"
	GetSpellCorrectionSearchTerm                                     = "GetSpellCorrectionSearchTerm"
	CreateSpellCorrectionSearchTerm                                  = "CreateSpellCorrectionSearchTerm"
	UpdateSpellCorrectionSearchTerm                                  = "UpdateSpellCorrectionSearchTerm"
	SaveUserSearchCategoryRelevance                                  = "SaveUserSearchCategoryRelevance"
	BuildSubCategoryMap                                              = "buildSubCategoryMap"
	PasEventExtractSearchRelevantDataFromSearchEvent                 = "PasEventExtractSearchRelevantDataFromSearchEvent"
	RankingEventExtractSearchRelevantDataFromSearchEvent             = "RankingEventExtractSearchRelevantDataFromSearchEvent"
	PasCronExtractSearchRelevantDataFromSearchEvent                  = "PasCronExtractSearchRelevantDataFromSearchEvent"
	PropertiesEventExtractSearchRelevantDataFromSearchEvent          = "PropertiesEventExtractSearchRelevantDataFromSearchEvent"
	BrowseRankingParamsEventExtractSearchRelevantDataFromSearchEvent = "BrowseRankingParamsEventExtractSearchRelevantDataFromSearchEvent"
	TagsEventExtractSearchRelevantDataFromSearchEvent                = "TagsEventExtractSearchRelevantDataFromSearchEvent"
	ExploreTagsEventExtractSearchRelevantDataFromSearchEvent         = "ExploreTagsEventExtractSearchRelevantDataFromSearchEvent"
	GetOpsOverriddenQueries                                          = "GetOpsOverriddenQueries"
	SaveUserSearchAttributeRelevance                                 = "SaveUserSearchAttributeRelevance"
	SaveUserSearchSubThemeRelevance                                  = "SaveUserSearchSubThemeRelevance"
	SaveUserSearchFeatureRelevance                                   = "SaveUserSearchFeatureRelevance"

	GetCrossSellWidgetDetails                             = "GetCrossSellWidgetDetails"
	SaveCrossSellWidgetDetails                            = "SaveCrossSellWidgetDetails"
	GetCrossSellAttachmentsByAttachmentIds                = "GetCrossSellAttachmentsByAttachmentIds"
	SaveAndMapCrossSellAttachments                        = "SaveAndMapCrossSellAttachments"
	SaveCrossSellWidgetDetailsAndAttachmentsInTransaction = "SaveCrossSellWidgetDetailsAndAttachmentsInTransaction"

	GetQueryIntentConfig           = "GetQueryIntentConfig"
	GetQueryIntentConfigFieldQuery = "GetQueryIntentConfigFieldQuery"
	GetQueryIntentConfigFieldCity

	SaveSearchQueryDescriptionDetails            = "SaveSearchQueryDescriptionDetails"
	GetSearchQueryDescriptionDetails             = "GetSearchQueryDescriptionDetails"
	GetAllUserSearchRelevanceWithFilterAndOffset = "GetAllUserSearchRelevanceWithFilterAndOffset"
	SavePharmaSaltMappingsData                   = "SavePharmaSaltMappingsData"
	UpsertQueryTermRedirection                   = "UpsertQueryTermRedirection"
	GetQueryTermRedirections                     = "GetQueryTermRedirections"
)
