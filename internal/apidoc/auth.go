package apidoc

// Register godoc
// @Summary      Register a new user
// @Description  Creates an account and returns a JWT.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      RegisterReq  true  "Registration payload"
// @Success      201      {object}  TokenResp
// @Failure      400      {object}  ErrorResp
// @Failure      409      {object}  ErrorResp
// @Router       /auth/register [post]
func RegisterDoc() {}

// Login godoc
// @Summary      Login
// @Description  Authenticates with email and password.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      LoginReq  true  "Credentials"
// @Success      200      {object}  TokenResp
// @Failure      401      {object}  ErrorResp
// @Router       /auth/login [post]
func LoginDoc() {}

// Refresh godoc
// @Summary      Refresh JWT
// @Description  Issues a new token from a valid existing JWT.
// @Tags         auth
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  RefreshResp
// @Failure      401  {object}  ErrorResp
// @Router       /auth/refresh [post]
func RefreshDoc() {}

// Me godoc
// @Summary      Current user
// @Description  Returns the authenticated user profile.
// @Tags         auth
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  User
// @Failure      401  {object}  ErrorResp
// @Failure      404  {object}  ErrorResp
// @Router       /auth/me [get]
func MeDoc() {}

// GitHubOAuthStart godoc
// @Summary      Start GitHub OAuth
// @Description  Redirects to GitHub authorization.
// @Tags         auth
// @Produce      json
// @Success      302  {string}  string  "Redirect to GitHub"
// @Router       /auth/oauth/github [get]
func GitHubOAuthStartDoc() {}

// GitHubOAuthCallback godoc
// @Summary      GitHub OAuth callback
// @Description  Handles GitHub OAuth redirect and issues a JWT.
// @Tags         auth
// @Produce      json
// @Param        code   query  string  true  "Authorization code"
// @Param        state  query  string  true  "CSRF state"
// @Success      302    {string}  string  "Redirect with token"
// @Router       /auth/oauth/github/callback [get]
func GitHubOAuthCallbackDoc() {}
