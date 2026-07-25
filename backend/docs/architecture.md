# Architecture

## Overview

Mwangaza is a farm advisory dashboard with a Go backend and Flutter frontend.

## Components

- **Flutter Dashboard** — Admin UI with stats, farm map, SMS history
- **Go Backend (Gin)** — REST API server with SQLite persistence
- **SpaceIoTBox API** — External satellite data provider
- **Africa's Talking** — SMS gateway for farmer notifications

## Data Flow

1. Flutter dashboard requests data via HTTP to Go backend
2. Backend queries SQLite or calls external services
3. Satellite data is fetched from SpaceIoTBox API
4. Recommendation engine analyzes data and generates advice
5. SMS alerts are sent via Africa's Talking

## Decision Engine

The recommendation service uses satellite data to generate farm-specific advice based on thresholds for soil moisture, temperature, rain probability, and NDVI.