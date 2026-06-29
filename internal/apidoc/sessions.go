package apidoc

// CreateSession godoc
// @Summary      Schedule a session
// @Description  Creates a mentorship session with a Jitsi room (mentor only).
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateSessionReq  true  "Session schedule"
// @Success      201      {object}  Session
// @Failure      401      {object}  ErrorResp
// @Failure      403      {object}  ErrorResp
// @Router       /sessions [post]
func CreateSessionDoc() {}

// ListSessionsByListing godoc
// @Summary      List sessions for a listing
// @Description  Returns scheduled sessions for a mentorship listing.
// @Tags         sessions
// @Produce      json
// @Param        listingId  path      string  true  "Listing ID"
// @Success      200        {array}   Session
// @Router       /sessions/listings/{listingId} [get]
func ListSessionsByListingDoc() {}

// UpdateSession godoc
// @Summary      Update session
// @Description  Updates scheduled time or status.
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string            true  "Session ID"
// @Param        request  body      UpdateSessionReq  true  "Fields to update"
// @Success      200      {object}  StatusResp
// @Router       /sessions/{id} [patch]
func UpdateSessionDoc() {}

// JoinSession godoc
// @Summary      Join session
// @Description  Returns the Jitsi URL for an authorized participant.
// @Tags         sessions
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Session ID"
// @Success      200  {object}  JoinSessionResp
// @Failure      401  {object}  ErrorResp
// @Failure      403  {object}  ErrorResp
// @Failure      404  {object}  ErrorResp
// @Router       /sessions/{id}/join [post]
func JoinSessionDoc() {}

// AddProgress godoc
// @Summary      Add progress note
// @Description  Records a progress entry for an enrollment.
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      ProgressReq  true  "Progress entry"
// @Success      201      {object}  map[string]string
// @Failure      401      {object}  ErrorResp
// @Failure      403      {object}  ErrorResp
// @Router       /sessions/progress [post]
func AddProgressDoc() {}

// ListProgress godoc
// @Summary      List progress entries
// @Description  Returns progress notes for an enrollment.
// @Tags         sessions
// @Produce      json
// @Security     BearerAuth
// @Param        enrollment_id  query     string  true  "Enrollment ID"
// @Success      200            {array}   ProgressEntry
// @Failure      400            {object}  ErrorResp
// @Failure      401            {object}  ErrorResp
// @Failure      403            {object}  ErrorResp
// @Router       /sessions/progress [get]
func ListProgressDoc() {}
