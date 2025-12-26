package stealth

import (
	"fmt"

	"github.com/go-rod/rod"
)

func ApplyFingerprint(page *rod.Page, fp FingerprintProfile) error {
	script := fmt.Sprintf(`
		// webdriver
		Object.defineProperty(navigator, 'webdriver', { get: () => undefined });

		// languages
		Object.defineProperty(navigator, 'languages', {
			get: () => %q
		});

		// platform
		Object.defineProperty(navigator, 'platform', {
			get: () => "%s"
		});

		// hardwareConcurrency
		Object.defineProperty(navigator, 'hardwareConcurrency', {
			get: () => %d
		});

		// plugins
		Object.defineProperty(navigator, 'plugins', {
			get: () => new Array(%d)
		});

		// WebGL spoof
		const getParameter = WebGLRenderingContext.prototype.getParameter;
		WebGLRenderingContext.prototype.getParameter = function(parameter) {
			if (parameter === 37445) return "%s"; // UNMASKED_VENDOR_WEBGL
			if (parameter === 37446) return "%s"; // UNMASKED_RENDERER_WEBGL
			return getParameter.call(this, parameter);
		};
	`, fp.Languages, fp.Platform, fp.HardwareC, fp.Plugins, fp.WebGLVend, fp.WebGLRend)

	_, err := page.Evaluate(script)
	return err
}
