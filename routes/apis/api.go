package apis

import (
	locationGateway "mms-api/src/Location/gateway"
	resourceGateway "mms-api/src/Resource/gateway"
	"mms-api/src/Services/gateway"

	"github.com/gofiber/fiber/v2"
)

func ServicesApi(c *fiber.App) {
	gateway.ServicesRoute(c)
	resourceGateway.ResourceRoute(c)
	locationGateway.LocationRoute(c)
}
