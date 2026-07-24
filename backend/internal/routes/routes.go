// Package routes registers all API endpoints on a Gin router.
//
// It takes the *sql.DB as a dependency, creates handler structs with it,
// and wires up middleware (CORS) and all route groups.
//
// Routes (under /api prefix):
//   GET    /health          — health_handler.HealthCheck
//   GET    /farms           — farm_handler.GetFarms
//   GET    /farms/:id       — farm_handler.GetFarm
//   GET    /satellite       — satellite_handler.GetSatelliteData
//   GET    /recommendation  — recommendation_handler.GetRecommendations
//   GET    /sms             — sms_handler.GetSMSHistory
//   POST   /sms/send        — sms_handler.SendSMS
//
// Called from main.go:
//   r := routes.Setup(db)
//   r.Run(":8080")
package routes