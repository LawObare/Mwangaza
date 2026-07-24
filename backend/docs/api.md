# API Documentation

This document describes all REST API endpoints exposed by the Mwangaza backend.

## Base URL

All endpoints are prefixed with `/api`.

## Endpoints

### GET /api/health
Returns server health status.

### GET /api/farms
Returns a list of all registered farms.

### GET /api/farms/:id
Returns a single farm by its ID.

### GET /api/satellite
Returns the latest satellite environmental data (soil moisture, temperature, rain probability, NDVI, wind speed).

### GET /api/recommendation
Returns all recommendations. Supports `?farm_id=` query parameter for filtering.

### GET /api/sms
Returns SMS send history.

### POST /api/sms/send
Sends an SMS to a farmer.

**Request body:** `{ "farm_id": int, "message": string, "phone_number": string }`

## Response Format

All responses follow: `{ "success": bool, "data": ..., "error": "..." }`