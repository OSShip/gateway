package apidoc

// GetProfile godoc
// @Summary      Get user profile
// @Description  Returns a user profile. Email is included when requesting your own profile.
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  Profile
// @Failure      404  {object}  ErrorResp
// @Router       /users/{id}/profile [get]
func GetProfileDoc() {}

// UpdateProfile godoc
// @Summary      Update own profile
// @Description  Updates display name, bio, or GitHub username.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      UpdateProfileReq  true  "Profile fields"
// @Success      200      {object}  Profile
// @Failure      401      {object}  ErrorResp
// @Router       /users/me [patch]
func UpdateProfileDoc() {}

// AddContribution godoc
// @Summary      Link a GitHub PR contribution
// @Description  Records a pull request URL and verifies merge status when possible.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      ContributionReq  true  "PR URL"
// @Success      201      {object}  Contribution
// @Failure      400      {object}  ErrorResp
// @Failure      401      {object}  ErrorResp
// @Router       /users/me/contributions [post]
func AddContributionDoc() {}

// GetEnrollments godoc
// @Summary      List user enrollments
// @Description  Returns enrollments for a user.
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {array}   Enrollment
// @Router       /users/{id}/enrollments [get]
func GetEnrollmentsDoc() {}

// CreateEnrollment godoc
// @Summary      Enroll in a listing
// @Description  Creates a pending enrollment for checkout.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateEnrollmentReq  true  "Listing to enroll in"
// @Success      201      {object}  Enrollment
// @Failure      401      {object}  ErrorResp
// @Failure      409      {object}  ErrorResp
// @Router       /users/enrollments [post]
func CreateEnrollmentDoc() {}

// ActivateEnrollment godoc
// @Summary      Activate enrollment
// @Description  Marks an enrollment active after successful payment.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id       path      string                 true  "Enrollment ID"
// @Param        request  body      ActivateEnrollmentReq  false "Optional checkout session ID"
// @Success      200      {object}  StatusResp
// @Failure      404      {object}  ErrorResp
// @Router       /users/enrollments/{id}/activate [post]
func ActivateEnrollmentDoc() {}
