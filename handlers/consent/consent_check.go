package consent

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go/v2"
	"misso/config"
	"misso/consts"
	"misso/global"
	"misso/utils"
	"net/http"
	"time"
)

func ConsentCheck(ctx *gin.Context) {
	oauth2challenge := ctx.Query("consent_challenge") // OAuth2 login
	if oauth2challenge == "" {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"error": "Necessary challenge not provided",
		})
		return
	}

	// Get OAuth2 Request
	global.Logger.Debugf("Getting OAuth2 Consent Request...")
	consentReq, _, err := global.Hydra.Admin.OAuth2Api.GetOAuth2ConsentRequest(context.Background()).ConsentChallenge(oauth2challenge).Execute()
	if err != nil {
		global.Logger.Errorf("Failed to get required consent request with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to get required consent request",
		})
		return
	}

	// Check if is able to skip
	if consentReq.Skip != nil && *consentReq.Skip {
		global.Logger.Debugf("Consent is able to skip, skipping...")
		// Skip consent confirm
		acceptReq, _, err := global.Hydra.Admin.OAuth2Api.AcceptOAuth2ConsentRequest(context.Background()).ConsentChallenge(oauth2challenge).AcceptOAuth2ConsentRequest(client.AcceptOAuth2ConsentRequest{
			GrantScope:               consentReq.RequestedScope,
			GrantAccessTokenAudience: consentReq.RequestedAccessTokenAudience,
		}).Execute()
		if err != nil {
			global.Logger.Errorf("Failed to accept consent request with error: %v", err)
			ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"error": "Failed to accept consent request",
			})
			return
		}

		// Redirect to target uri
		ctx.Redirect(http.StatusTemporaryRedirect, acceptReq.RedirectTo)

		global.Logger.Debugf("User should now be redirecting to target URI.")
	} else {
		global.Logger.Debugf("Consent is NOT able to skip, preparing authorize UI...")
		// Generate CSRF token
		global.Logger.Debugf("Generating CSRF token...")
		csrf := utils.RandString(32)
		sessKey := fmt.Sprintf(consts.REDIS_KEY_CONSENT_CSRF, oauth2challenge)
		err := global.Redis.Set(context.Background(), sessKey, csrf, time.Duration(config.Config.Time.RequestValid)*time.Second).Err()
		if err != nil {
			global.Logger.Errorf("Failed to save csrf into redis with error: %v", err)
			ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"error": "Failed to save csrf",
			})
			return
		}

		// Retrieve context
		global.Logger.Debugf("Retrieving context...")

		userinfo, err := utils.GetUserinfo(*consentReq.Subject)
		if err != nil {
			global.Logger.Errorf("Failed to retrieve userinfo with error: %v", err)
			ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"error": "Failed to get userinfo",
			})
			return
		}

		// Show the consent UI
		global.Logger.Debugf("Rendering consent UI...")
		templateFields := gin.H{
			"user":      *userinfo,
			"challenge": oauth2challenge,
			"csrf":      csrf,
			"scopes":    consentReq.RequestedScope,
		}

		if consentReq.Client.LogoUri != nil && *consentReq.Client.LogoUri != "" {
			templateFields["logo"] = *consentReq.Client.LogoUri
		}
		if consentReq.Client.ClientName != nil && *consentReq.Client.ClientName != "" {
			templateFields["clientName"] = *consentReq.Client.ClientName
		} else {
			templateFields["clientName"] = *consentReq.Client.ClientId
		}
		if consentReq.Client.PolicyUri != nil && *consentReq.Client.PolicyUri != "" {
			templateFields["clientPolicy"] = *consentReq.Client.PolicyUri
		}
		if consentReq.Client.TosUri != nil && *consentReq.Client.TosUri != "" {
			templateFields["clientTos"] = *consentReq.Client.TosUri
		}
		ctx.HTML(http.StatusOK, "consent.tmpl", templateFields)

		global.Logger.Debugf("User should now see Consent UI.")
	}

}
