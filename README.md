## Demonstration Video

Video link:
https://github.com/paragraph1148/linkedin-automation/blob/main/demo/final.mp4

# LinkedIn Automation System (Go + Rod)

A stealth-focused browser automation system built in Go using Rod, designed to simulate human-like interaction patterns while safely handling detection scenarios.

This project focuses on **behavioral realism, anti-detection techniques, and defensive system design**, rather than simple scripting.

---

## Disclaimer

This project is a technical proof-of-concept created for educational purposes.  
Automating LinkedIn violates LinkedIn’s Terms of Service.  
This system is not intended for production or real account usage.

---

## Key Highlights

- Human-like interaction simulation (mouse movement, typing, scrolling)
- Multi-layer anti-detection system
- Automatic checkpoint detection with safe fallback
- Mock mode for deterministic and risk-free execution
- Modular Go architecture with clean separation of concerns
- Persistent state and rate-limiting across sessions

---

## What This System Solves

Most automation scripts are deterministic and easily detectable.

This system is designed to behave more like a **cautious human user**, introducing:

- Non-linear interaction patterns
- Timing variability
- Behavioral noise
- Defensive execution strategies

---

## System Architecture

The project follows a modular, package-based architecture with clear separation of concerns:

- `cmd/main.go`  
  Entry point responsible for orchestration, configuration loading, and execution flow.

- `internal/auth`  
  Handles authentication, session cookies, and checkpoint detection.

- `internal/browser`  
  Manages browser launch configuration and fingerprint setup.

- `internal/stealth`  
  Implements human-like behavior (timing, cursor movement, scrolling, delays).

- `internal/search`  
  Handles profile search, pagination, parsing, and deduplication.

- `internal/messaging`  
  Simulates a messaging workflow with persistent state tracking.

- `internal/config`  
  Loads configuration from YAML and environment variables.

Each module follows **single-responsibility principles** and is designed for clarity and extensibility.

---

## Stealth & Anti-Detection Techniques

### Core Techniques

- **Human-like mouse movement**  
  Bézier-style cursor paths, non-center targeting, natural pauses

- **Randomized timing patterns**  
  Think time before actions and reaction delays after actions

- **Browser fingerprint masking**  
  User-Agent override, viewport randomization, JS-level tweaks

### Additional Techniques

- Random scrolling behavior
- Realistic typing simulation
- Hover-before-click interactions
- Activity scheduling (business hours only)
- Hourly and daily rate limiting
- Persistent session state
- Cookie reuse
- Automatic checkpoint detection

---

## Activity Scheduling & Rate Limiting

To mimic real user behavior:

- Runs only during configured business hours
- Enforces hourly and daily action limits
- Limits persist across restarts
- Gracefully stops execution when limits are reached

---

## Checkpoint Detection & Safe Fallback

If LinkedIn triggers a security checkpoint:

- Live automation stops immediately
- No retries are attempted
- System switches to mock mode automatically

This ensures **safe execution and avoids account risk**.

---

## Mock Mode

Mock mode simulates LinkedIn using local HTML.

- Enabled via:
  ```bash
  USE_MOCK=1
  ```
- Automatically triggered on checkpoint detection

- Allows deterministic and safe testing

- Demonstrates parsing, deduplication, and workflow logic

---

## Configuration

Configuration is managed via config.yaml.

Example:

schedule:
start_hour: 9
end_hour: 18

limits:
daily_connections: 10

Sensitive data (credentials) is handled via environment variables.

---

## Build & Run

Prerequisites

- Go 1.20+

- Chromium (handled automatically by Rod)

Build

- go build ./...

Run (Safe Mock Mode)

- USE_MOCK=1 go run cmd/main.go

Run (Educational Only)

- go run cmd/main.go

---

## Engineering Approach

This system emphasizes:

- Defensive design over aggressive automation

- Behavior simulation over deterministic scripting

- Modular architecture for testability and extensibility

- Safe fallbacks instead of retry loops

---

## Future Improvements

- Adaptive behavior based on feedback signals

- ML-based anomaly detection avoidance

- Proxy rotation and distributed execution

- Smarter interaction strategies based on page context

---

## What This Demonstrates

- Advanced browser automation concepts

- Human-like interaction modeling

- Anti-detection awareness

- Clean and scalable Go architecture

- Real-world system design thinking
