package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	username   string
	expiration time.Time
}

// this map stores the users sessions. For larger scale applications, you can use a database or cache for this purpose
var sessions = map[string]Session{}

func (session Session) isExpired() bool {
	return session.expiration.Before(time.Now())
}

func ReadCookie(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {

	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return w, r

		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return w, r

	}
	sessionToken := c.Value

	// We then get the name of the user from our session map, where we set the session token
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return w, r

	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return w, r

	}
	// Finally, return the welcome message to the user
	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
	return w, r
}

func CreateCookie(w http.ResponseWriter, r *http.Request, username string) (http.ResponseWriter, *http.Request) {
	// Create a new random session token
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	sessions[sessionToken] = Session{
		username:   username,
		expiration: expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   sessionToken,
		Expires: expiresAt,
		// HttpOnly: true,
		// Secure:   true,
		// SameSite: http.SameSiteLaxMode,
	})
	return w, r
}

func UpdateCookie(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	// (BEGIN) The code from this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return w, r
		}
		w.WriteHeader(http.StatusBadRequest)
		return w, r
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return w, r
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return w, r
	}
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// If the previous session is valid, create a new session token for the current user
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	sessions[newSessionToken] = Session{
		username:   userSession.username,
		expiration: expiresAt,
	}

	// Delete the older session token
	delete(sessions, sessionToken)

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	return w, r
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return w, r

		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return w, r

	}
	sessionToken := c.Value

	// remove the users session from the session map
	delete(sessions, sessionToken)

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
	return w, r

}
