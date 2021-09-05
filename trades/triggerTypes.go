package trades

type triggerType uint8

const (
	PctMoveUp triggerType = iota
	PctMoveDown
	TargetCeiling
	TargetFloor
	IndustryCorrelation
	SP500Correlation

	// ongoing checks
	PctProfit
	PctLoss
	Delta
	DeltaUp
	DeltaDown
	Theta
	ThetaUp
	ThetaDown
	SentimentChange
	IndustryDivergence

	// Trend signals
	Ema200CrossUp
	Ema200CrossDown
	Ema50CrossUp
	Ema50CrossDown

	// Interest signals
	VolumePctUp    // outlier spike in volume as compared to historical volume
	VolumePctDown  // outlier drop in volume as compared to historical volume
	PercentageUp   // outlier price spike (percentage) as compared to historical beta
	PercentageDown // outlier price drop (percentage) as compared to historical beta

	// fundamentals
	PECeiling          // P/E ratio upper boundary
	PEFloor            // P/E ratio upper boundary
	TTMEarningsCeiling // TTMEarnings upper boundary
	TTMEarningsFloor   // TTMEarnings lower boundary
	MarketCapCeiling
	MarketCapFloor

	// Sentiment
	EarningsSurpriseUp
	EarningsSurpriseDown
	NewsSentimentUp
	NewsSentimentDown
	TwitterSentimentUp
	TwitterSentimentDown
	SearchSentimentUp
	SearchSentimentDown
)
