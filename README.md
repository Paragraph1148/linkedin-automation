## Demonstration Video

Video link:
https://github.com/paragraph1148/linkedin-automation/blob/main/demo/final.mp4

# LinkedIn Automation Assignment (Educational Proof of Concept)

IMPORTANT DISCLAIMER  
This project is a technical proof-of-concept created strictly for educational and evaluation purposes.  
Automating LinkedIn violates LinkedIn’s Terms of Service.  
This tool must not be used in production or on real LinkedIn accounts.

---

## Project Overview

This repository demonstrates a Go-based browser automation proof-of-concept using the Rod library.  
The primary focus is human-like behavior simulation, anti-detection techniques, and clean Go architecture.

The project intentionally prioritizes:

- Stealth engineering
- Behavioral realism
- Defensive execution
- Safe fallback mechanisms

---

## Objectives

- Demonstrate advanced browser automation concepts in Go
- Implement multiple stealth and anti-detection techniques
- Simulate realistic human interaction patterns
- Maintain clean, modular, and testable architecture
- Provide safe execution paths that avoid account risk

---

## Architecture Overview

The project follows a modular, package-based architecture with clear separation of concerns:

- cmd/main.go  
  Entry point responsible for orchestration, configuration loading, and execution flow.

- internal/auth  
  Handles authentication, session cookies, authentication guards, and checkpoint detection.

- internal/browser  
  Responsible for browser launch configuration and fingerprint setup.

- internal/stealth  
  Implements human-like behavior simulation, timing control, mouse movement, scrolling, and rate limiting.

- internal/search  
  Handles profile search, pagination, parsing, deduplication, and mock HTML parsing.

- internal/messaging  
  Implements a mocked messaging workflow with persistent state tracking.

- internal/config  
  Loads and validates configuration from YAML and environment variables.

- internal/meta  
  Contains capability metadata used for demonstration and logging.

Each package follows single-responsibility principles with explicit and well-defined boundaries.

---

## Stealth and Anti-Detection Techniques

This project implements more than eight stealth techniques, including all mandatory requirements.

### Mandatory Techniques

- Human-like mouse movement  
  Bézier-style cursor movement, non-center targeting, natural pauses

- Randomized timing patterns  
  Think time before actions and reaction delays after actions

- Browser fingerprint masking  
  User-Agent override, viewport randomization, JavaScript-level fingerprint adjustments

### Additional Techniques

- Random scrolling behavior
- Realistic typing simulation
- Mouse hovering before interactions
- Activity scheduling (business hours only)
- Hourly and daily rate limiting
- Persistent state across runs
- Session cookie reuse
- Automatic checkpoint detection

---

## Activity Scheduling and Rate Limiting

To simulate realistic human usage patterns:

- Automation runs only during configured business hours
- Hourly and daily action limits are enforced
- Limits persist across application restarts
- When limits are reached, execution stops gracefully

---

## Checkpoint Detection and Safe Fallback

If LinkedIn triggers a security checkpoint or challenge:

- Live automation stops immediately
- No retries are attempted
- Execution automatically switches to mock mode
- The demo continues without account risk

This design protects accounts and demonstrates defensive automation engineering.

---

## Mock Mode

Mock mode simulates LinkedIn pages using local HTML files.

- Enabled explicitly via environment variable:
  USE_MOCK=1
- Automatically enabled when a checkpoint is detected
- Allows deterministic, safe execution
- Demonstrates parsing, deduplication, state tracking, and messaging logic

---

## Messaging System (Mocked)

The messaging system is intentionally mocked.

- Simulates accepted connections
- Sends follow-up messages using templates
- Tracks message state persistently
- Demonstrates workflow design without violating Terms of Service

---

## Configuration

Configuration is loaded from `config.yaml` with sensible defaults.

Example:
schedule:
start_hour: 9
end_hour: 18

limits:
daily_connections: 10

Environment variables are used for credentials and are never committed.

---

## Build and Run Instructions

### Prerequisites

- Go version 1.20 or newer
- Chromium (automatically handled by Rod)

### Build

go build ./...

### Run in Mock Mode (Safe)

USE_MOCK=1 go run cmd/main.go

### Run Live (Educational Only)

go run cmd/main.go

If a checkpoint is detected, the system automatically falls back to mock mode.

---

## Design Decisions

- Headless mode is used for compatibility on immutable Linux systems such as Fedora Silverblue
- Checkpoint fallback ensures safe, responsible execution
- Mock-first execution allows reproducible demonstrations
- Explicit scheduling and rate limits reduce bot-like behavior
- Clean Go patterns improve maintainability and clarity

---

## Demonstration Video

The demonstration video includes:

- Project setup
- Mock mode execution
- Stealth behavior overview
- Rate limiting behavior
- Checkpoint detection and fallback explanation

The video link is provided in the submission.

---

## Summary

This project is not a production automation tool.

It is a carefully designed educational proof-of-concept that demonstrates:

- Advanced browser automation concepts
- Human-like interaction simulation
- Anti-detection awareness
- Clean and modular Go architecture
- Responsible engineering practices

---

## License

MIT License
