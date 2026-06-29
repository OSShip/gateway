package apidoc

// Checkout godoc
// @Summary      Create Stripe Checkout session
// @Description  Starts checkout for a listing enrollment. Uses Stripe Connect when the mentor is onboarded.
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CheckoutReq  true  "Checkout payload"
// @Success      200      {object}  CheckoutResp
// @Failure      403      {object}  ErrorResp
// @Router       /payments/checkout [post]
func CheckoutDoc() {}

// GetLedger godoc
// @Summary      Listing ledger
// @Description  Returns payment ledger entries for a listing (mentor or admin).
// @Tags         payments
// @Produce      json
// @Security     BearerAuth
// @Param        listingId  path      string  true  "Listing ID"
// @Success      200        {array}   LedgerEntry
// @Failure      403        {object}  ErrorResp
// @Router       /payments/ledger/{listingId} [get]
func GetLedgerDoc() {}

// PayoutSummary godoc
// @Summary      Public payout summary
// @Description  Returns anonymized platform payout aggregates.
// @Tags         payments
// @Produce      json
// @Success      200  {object}  PayoutSummary
// @Router       /public/payout-summary [get]
func PayoutSummaryDoc() {}

// ConnectOnboard godoc
// @Summary      Start Stripe Connect onboarding
// @Description  Creates or resumes Stripe Express onboarding for a mentor.
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      ConnectOnboardReq  true  "Return and refresh URLs"
// @Success      200      {object}  ConnectOnboardResp
// @Router       /payments/connect/onboard [post]
func ConnectOnboardDoc() {}

// ConnectStatus godoc
// @Summary      Stripe Connect status
// @Description  Returns whether the mentor has completed Stripe Connect onboarding.
// @Tags         payments
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  ConnectStatus
// @Router       /payments/connect/status [get]
func ConnectStatusDoc() {}
