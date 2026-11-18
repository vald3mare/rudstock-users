package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-platform/httpx"
)

type Identity struct {
	UserID        uuid.UUID `json:"user_id"`
	SessionID     uuid.UUID `json:"session_id"`
	Role          string    `json:"role"`
	Email         *string   `json:"email"`
	EmailVerified *bool     `json:"email_verified"`
}

type Client struct {
	verifyURL string
	http      *http.Client
}

type identityKey struct{}

func IdentityFrom(ctx context.Context) (Identity, bool) {
	identity, ok := ctx.Value(identityKey{}).(Identity)
	return identity, ok
}

func NewClient(baseURL string) (*Client, error) {
	base, err := url.Parse(baseURL)
	if err != nil || (base.Scheme != "http" && base.Scheme != "https") || base.Hostname() == "" ||
		base.User != nil || base.RawQuery != "" || base.ForceQuery || base.Fragment != "" {
		return nil, errors.New("AUTH_URL must be an HTTP base URL")
	}
	return &Client{
		verifyURL: base.JoinPath("auth", "verify").String(),
		http: &http.Client{
			Timeout: 3 * time.Second,
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}, nil
}

func (c *Client) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fields := strings.Fields(r.Header.Get("Authorization"))
		if len(fields) != 2 || !strings.EqualFold(fields[0], "Bearer") {
			writeUnauthorized(w, r, "unauthorized", "Authorization required")
			return
		}

		req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, c.verifyURL, nil)
		if err != nil {
			writeUnavailable(w, r, err)
			return
		}
		req.Header.Set("Authorization", "Bearer "+fields[1])
		resp, err := c.http.Do(req)
		if err != nil {
			writeUnavailable(w, r, err)
			return
		}
		defer resp.Body.Close()
		body := json.NewDecoder(io.LimitReader(resp.Body, 16<<10))

		switch resp.StatusCode {
		case http.StatusOK:
			var identity Identity
			if err := body.Decode(&identity); err != nil || identity.UserID == uuid.Nil || identity.SessionID == uuid.Nil {
				writeUnavailable(w, r, errors.New("invalid identity from auth"))
				return
			}
			ctx := context.WithValue(r.Context(), identityKey{}, identity)
			next.ServeHTTP(w, r.WithContext(ctx))
		case http.StatusUnauthorized:
			var result httpx.ErrorResponse
			if body.Decode(&result) == nil && result.Slug == "token-expired" {
				writeUnauthorized(w, r, "token-expired", "Access token expired")
				return
			}
			writeUnauthorized(w, r, "invalid-token", "Invalid access token")
		case http.StatusForbidden:
			httpx.WriteJSON(w, r, http.StatusForbidden, httpx.ErrorResponse{
				Slug: "forbidden",
				Msg:  "Access forbidden",
			})
		default:
			writeUnavailable(w, r, fmt.Errorf("verify returned HTTP %d", resp.StatusCode))
		}
	})
}

func writeUnauthorized(w http.ResponseWriter, r *http.Request, slug, message string) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	httpx.WriteJSON(w, r, http.StatusUnauthorized, httpx.ErrorResponse{Slug: slug, Msg: message})
}

func writeUnavailable(w http.ResponseWriter, r *http.Request, err error) {
	slog.Warn("auth verification failed", "error", err)
	httpx.WriteJSON(w, r, http.StatusServiceUnavailable, httpx.ErrorResponse{
		Slug: "auth-unavailable",
		Msg:  "Authentication service is unavailable",
	})
}
