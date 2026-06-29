package apidoc

// ListListings godoc
// @Summary      List listings
// @Description  Returns mentorship listings filtered by status and optional OSS project.
// @Tags         listings
// @Produce      json
// @Param        status       query  string  false  "Listing status"  default(active)
// @Param        oss_project  query  string  false  "Filter by OSS project name"
// @Success      200          {array}   Listing
// @Router       /listings [get]
func ListListingsDoc() {}

// GetListing godoc
// @Summary      Get listing
// @Description  Returns a single listing by ID.
// @Tags         listings
// @Produce      json
// @Param        id   path      string  true  "Listing ID"
// @Success      200  {object}  Listing
// @Failure      404  {object}  ErrorResp
// @Router       /listings/{id} [get]
func GetListingDoc() {}

// CreateListing godoc
// @Summary      Create listing
// @Description  Creates a mentorship listing (approved mentor only).
// @Tags         listings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateListingReq  true  "Listing payload"
// @Success      201      {object}  Listing
// @Failure      403      {object}  ErrorResp
// @Router       /listings [post]
func CreateListingDoc() {}

// UpdateListing godoc
// @Summary      Update listing
// @Description  Updates a listing owned by the authenticated mentor.
// @Tags         listings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string            true  "Listing ID"
// @Param        request  body      UpdateListingReq  true  "Fields to update"
// @Success      200      {object}  Listing
// @Failure      403      {object}  ErrorResp
// @Failure      404      {object}  ErrorResp
// @Router       /listings/{id} [patch]
func UpdateListingDoc() {}

// ListPublicListings godoc
// @Summary      List public listings
// @Description  Public alias for active listings (cached).
// @Tags         listings
// @Produce      json
// @Param        status       query  string  false  "Listing status"  default(active)
// @Param        oss_project  query  string  false  "Filter by OSS project name"
// @Success      200          {array}   Listing
// @Router       /public/listings [get]
func ListPublicListingsDoc() {}
