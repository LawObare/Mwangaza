# Presentation Guide

Key talking points for demoing the Mwangaza system:

## Problem

Small-scale farmers lack access to real-time environmental data and personalized advice.

## Solution

Mwangaza combines satellite data, a decision engine, and SMS alerts to deliver actionable farm recommendations.

## Demo Flow

1. **Dashboard** — Show overall stats (soil moisture, temperature, rain, NDVI)
2. **Farms** — Browse registered farms and their status
3. **Map** — Visualize farm locations with color-coded health markers
4. **Recommendations** — Show generated advice with priority levels
5. **SMS** — Demonstrate sending alerts to farmers' phones

## Architecture

Flutter frontend → Go API → SpaceIoTBox + SQLite + Africa's Talking