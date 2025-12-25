# LinkedIn Automation – Technical Proof of Concept (Go + Rod)

> ** Educational Use Only – Do NOT use in production**

This repository contains a **technical proof-of-concept** demonstrating advanced browser automation, anti-detection strategies, and clean Go architecture using the **Rod** browser automation library.

The project is built **strictly for educational and evaluation purposes** to showcase automation engineering skills, human-like interaction modeling, and stealth mechanisms.
It **must not** be used against real LinkedIn accounts.

---

## Critical Disclaimer

### Educational Purpose Only

This project exists solely to demonstrate:

- Browser automation patterns
- Anti-bot detection avoidance concepts
- Clean, modular Go code architecture

### Terms of Service Violation

Automating LinkedIn violates their Terms of Service.
Using this code on real accounts **may result in permanent bans or legal consequences**.

### Do Not Use in Production

This project is **not production-ready** and intentionally avoids real-world deployment safeguards.

---

## Project Objective

Build a Go-based automation framework that simulates **authentic human behavior** while navigating a modern, bot-protected website.

The focus is on:

- Stealth & anti-detection techniques
- Realistic interaction simulation
- Maintainable, testable architecture

---

## Architecture Overview

```
cmd/
 └── main.go              # Application entry point

internal/
 ├── auth/                # Login, cookies, checkpoint detection
 ├── search/              # Search & profile parsing
 ├── messaging/           # Follow-up message handling (mocked)
 ├── stealth/             # Human-like behavior & anti-detection
 ├── state/               # Persistent state (JSON)
 ├── config/              # Config loading & validation
 └── logger/              # Structured logging
```

### Design Principles

- Clear separation of concerns
- Testability via mock mode
- Graceful failure handling
- Explicit demonstration of stealth techniques

---

## Stealth & Anti-Detection Techniques

This project implements **8+ anti-detection strategies**, combining mandatory and additional techniques.

### Mandatory Techniques

1. **Human-Like Mouse Movement**

   - Bézier-curve based cursor paths
   - Natural overshoot & micro-corrections
   - Avoids straight-line motion

2. **Randomized Timing Patterns**

   - Variable think times
   - Randomized delays between actions
   - Non-deterministic interaction intervals

3. **Browser Fingerprint Masking**

   - Custom / randomized User-Agent
   - Randomized viewport dimensions
   - `navigator.webdriver` removal via JS injection

### Additional Stealth Techniques

4. **Random Scrolling Behavior**

   - Variable scroll speeds
   - Occasional scroll-back actions

5. **Realistic Typing Simulation**

   - Per-character delays
   - Occasional typos with backspaces

6. **Mouse Hovering & Idle Movement**

   - Random hover events
   - Cursor wandering during idle time

7. **Rate Limiting & Throttling**

   - Daily action quotas
   - Cooldown periods between actions

8. **Activity Scheduling**

   - Operates only during configurable business hours
   - Simulated breaks between sessions

> ⚠️ These implementations are demonstrative and intentionally simplified to highlight concepts rather than evade real systems.

---

## Mock Mode (Safe Testing)

To avoid real website interaction, the project supports **mock mode**.

### Enable Mock Mode

```bash
USE_MOCK=1 go run cmd/main.go
```

Mock mode:

- Loads static HTML fixtures
- Tests parsing & logic safely
- Enables rapid iteration without live requests

---

## Configuration

### `.env.example`

```env
LINKEDIN_EMAIL=your_email
LINKEDIN_PASSWORD=your_password
USE_MOCK=1
```

### `config.yaml`

```yaml
browser:
  headless: false
  user_agent: random

timing:
  min_delay_ms: 800
  max_delay_ms: 3000

limits:
  daily_connections: 10
```

Environment variables override config values.

---

## State Persistence

Application state is persisted in JSON to allow:

- Resume after crashes
- Duplicate profile detection
- Daily limit enforcement

Tracked data includes:

- Sent connection requests
- Accepted connections
- Messages sent
- Last run timestamp

---

## Logging

Structured logging using Go’s `slog`:

- Levels: debug, info, warn, error
- Contextual metadata (profile URL, action type)
- Timestamped entries

Example:

```
INFO  sending connection request profile=https://...
WARN  daily limit reached
```

---

## Demonstration Video

**Demo Video:** _(link to be added)_

The video walkthrough demonstrates:

- Project setup
- Config & mock mode
- Stealth techniques
- Logging & state persistence

---

## Limitations

- No real LinkedIn automation is performed
- CAPTCHA & advanced checkpoint handling is mocked
- Selectors and flows are simplified intentionally

These trade-offs are **deliberate** to prioritize architectural clarity and safety.

---

## Evaluation Focus

This project prioritizes:

- Stealth technique implementation quality
- Code structure & clarity
- Demonstration of automation engineering concepts

It is **not** intended as a production-grade automation tool.

---

## Final Note

This repository demonstrates **how** advanced automation systems are built—not how to exploit platforms.

**Do not use this code against live accounts.**
