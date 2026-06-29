package apidoc

// ApplyMentor godoc
// @Summary      Apply to become a mentor
// @Description  Submits a mentor application with GitHub contribution data.
// @Tags         mentors
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      MentorApplyReq  false  "Optional GitHub username override"
// @Success      201      {object}  MentorApplication
// @Failure      400      {object}  ErrorResp
// @Failure      401      {object}  ErrorResp
// @Failure      409      {object}  ErrorResp
// @Router       /mentors/apply [post]
func ApplyMentorDoc() {}

// ListMentorApplications godoc
// @Summary      List mentor applications
// @Description  Returns mentor applications (admin only).
// @Tags         mentors
// @Produce      json
// @Security     BearerAuth
// @Param        status  query     string  false  "Filter by status"  default(pending)
// @Success      200     {array}   MentorApplication
// @Failure      403     {object}  ErrorResp
// @Router       /mentors/admin/applications [get]
func ListMentorApplicationsDoc() {}

// ReviewMentorApplication godoc
// @Summary      Review mentor application
// @Description  Approves or rejects a mentor application (admin only).
// @Tags         mentors
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string                true  "Application ID"
// @Param        request  body      ReviewApplicationReq  true  "Review decision"
// @Success      200      {object}  StatusResp
// @Failure      400      {object}  ErrorResp
// @Failure      403      {object}  ErrorResp
// @Failure      404      {object}  ErrorResp
// @Router       /mentors/admin/applications/{id} [patch]
func ReviewMentorApplicationDoc() {}
