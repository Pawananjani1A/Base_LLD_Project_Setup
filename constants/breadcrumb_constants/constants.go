package breadcrumb_constants

const (
	StageSegmentation        = "Segmentation Layer"
	ActionSegmentationError  = "Error in segmentation"
	ActionSegmentationPassed = "Passed segmentation"

	StageCategoryCohort        = "Category Cohort Layer"
	ActionCategoryCohortError  = "Error in category cohort"
	ActionCategoryCohortPassed = "Passed category cohort"

	StageRanker                 = "Ranker Layer"
	ActionModelRankerCall       = "Model ranker called"
	ActionRankerVariantSelected = "Ranking variant selected"
	ActionRankerFlowDecision    = "Flow decision made"
	ActionRankerQrlExecution    = "QRL ranking executed"
	ActionRankerProductsMerged  = "Products merged"
)

const (
	StageProductRetrieval  = "Atlas Product Retrieval Layer"
	ActionProductRetrieved = "Retrieved from Atlas"

	StageRetrievalSetBucketing         = "Retrieval Set Bucketing Layer"
	ActionClassifiedAsOrganic          = "Classified into Organic bucket"
	ActionClassifiedAsSemantic         = "Classified into Semantic bucket"
	ActionMergedFromOrganicAndSemantic = "Merged from Organic and Semantic buckets"

	StageRelevance          = "Relevance Layer"
	ActionRelevanceReranked = "Reranked by relevance score"

	StageProductReranking = "Reranking Layer"
	ActionRerankAssigned  = "Sorted by QRL score"

	StageRuleBasedReranking   = "Rule Reranking Layer"
	ActionReorderedDueToRules = "Rank adjusted by rule-based reranking"

	StageCategoryUnderstanding      = "L3 Understanding Layer"
	ActionReorderedDueToL3Bucketing = "Reordered by L3 bucket priority"

	StageServiceabilityFiltering = "Unserviceable Bucket Layer"
	ActionDeboostedUnserviceable = "Deboosted due to store unavailability"

	StageVIPProductBoosting = "VIP Deal Boosting Layer"
	ActionBoostedVIPDeals   = "Boosted for VIP deal qualification"

	StagePreviousBoughtLayer    = "Previous Bought Layer"
	StagePreviousBoughtLayerV2  = "Previous Bought Layer V2"
	ActionBoostedPreviousBought = "Boosted based on previous purchases"
	ActionDeboostedDueToOS      = "Deboosted for previous purchase and out-of-stock"

	StagePTAttributeLayer         = "PT Attribute Understanding Layer"
	ActionReorderedByPTAttributes = "Rank adjusted by product type attributes"

	StagePostQRLMatchReordering = "Post QRL Match Reordering Layer"
	ActionReorderedByQRLMatch   = "Reordered after QRL match analysis"

	StagePostFilterSortLayer      = "Post Filter Sorting Layer"
	ActionMergedPostFilterBuckets = "Merged from filtered buckets"

	StageAdsDeboosting = "Ads Deboosting Layer"
	ActionDeboostedDueToAds = "Deboosted orgranic products due to ads"

	StageAdsFlyWheelDeboosting = "Ads FlyWheel Deboosting Layer"
	ActionDeboostedDueToAdsFlyWheel = "Deboosted orgranic products due to ads flywheel"

	StageFinalDelivery           = "Final Delivery Layer"
	ActionPostToPaginationAndPas = "Forwarded to pagination and PAS layers"

	StagePasLayer          = "PAS Layer"
	ActionForwardedToPas   = "Forwarded to PAS"
	ActionRetrievedFromPas = "Retrieved from PAS"
	ActionDroppedByPas     = "Dropped by PAS"

	StageRerankingBasedOnStock       = "Reranking Based on Stock Layer"
	ActionDeboostBelowThresholdStock = "Deboosted due to stock below threshold"

	StageOutOfStockRemovalLayer          = "Out of Stock Removal Layer"
	ActionRetainedAfterOutOfStockRemoval = "Retained after out-of-stock filtering"
)
