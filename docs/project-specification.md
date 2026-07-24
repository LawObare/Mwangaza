# Project Specification

## Overview

Mwangaza is a farm advisory dashboard that uses satellite data to generate irrigation, rainfall, and temperature recommendations for small-scale farmers in Kenya.

## Target Users

- **Farmers** — Receive SMS alerts on feature phones (no smartphone required)
- **Admin** — Uses the Flutter dashboard to monitor farms and send broadcasts

## Core Features

1. Satellite data ingestion from SpaceIoTBox API
2. Recommendation engine (irrigation, rainfall, temperature, NDVI analysis)
3. SMS alerts via Africa's Talking
4. Farm management and mapping
5. Dashboard with real-time stats

## Tech Stack

- **Frontend:** Flutter + OpenStreetMap
- **Backend:** Go + Gin
- **Database:** SQLite
- **Satellite:** SpaceIoTBox API
- **SMS:** Africa's Talking

## Constraints

- 3-hour hackathon build
- Offline demo mode via mock data
- Must work without internet (USE_MOCK_DATA=true)