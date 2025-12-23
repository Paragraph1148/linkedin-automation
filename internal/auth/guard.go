package auth

import (
	"errors"
	"strings"
	"github.com/go-rod/rod"
)

func EnsureAuthenticated(page *rod.Page) error {
	url := page.MustInfo().URL

	if strings.Contains(url, "/login") {
		return errors.New("not authenticated: redirected to login")
	}

	if strings.Contains(url, "checkpoint") || strings.Contains(url, "challenge") {
		return errors.New("authentication checkpoint detected")
	}

	return nil
}