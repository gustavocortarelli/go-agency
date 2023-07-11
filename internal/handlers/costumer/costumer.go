package costumer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavocortarelli/go-agency/internal/model"
	"github.com/gustavocortarelli/go-agency/internal/service/costumer"
	"log"
	"net/http"
	"strconv"
)

func Get(context *fiber.Ctx) error {
	log.Printf("Starting get costumer by ID request")
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid ID on request",
		})
		return err
	}
	log.Printf("Getting costumer data: ID %d...", id)
	costumerData, err := costumer.Get(int64(id))
	log.Printf("Data was fetched")

	context.Status(http.StatusOK).JSON(costumerData)
	log.Printf("Get costumer by ID request was done")
	return nil
}

func GetAll(context *fiber.Ctx) error {
	log.Printf("Starting get all costumers request")
	costumers, err := costumer.GetAll()
	log.Printf("Data was fetched")
	if err != nil {
		log.Printf("Error during get all costumers: %d", err)
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Could not get costumers",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(costumers)
	log.Printf("Get all costumers request was done")
	return nil
}

func Create(context *fiber.Ctx) error {
	log.Printf("Starting create costumer request")
	cost := model.Costumer{}
	err := context.BodyParser(&cost)
	if err != nil {
		log.Printf("Error parsing body data: %d", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Error parsing body data: invalid request",
		})
		return err
	}
	log.Printf("Creating costumer...")
	cost, err = costumer.Create(cost)
	log.Printf("Costumer was created")
	context.Status(http.StatusOK).JSON(cost)
	log.Printf("Create costumer request was done")
	return err
}

func Delete(context *fiber.Ctx) error {
	log.Printf("Starting delete costumer request")
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid ID on delete request",
		})
		return err
	}
	log.Printf("Deleting costumer...")
	err = costumer.Delete(int64(id))
	log.Printf("Costumer was deleted")
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Costumer was deleted"})
	log.Printf("Delete costumer request was done")
	return err
}

func Update(context *fiber.Ctx) error {
	log.Printf("Starting update costumer request")
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid ID on delete request",
		})
		return err
	}

	cost := model.Costumer{}
	err = context.BodyParser(&cost)
	if err != nil {
		log.Printf("Error parsing body data: %d", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Error parsing body data: invalid request",
		})
		return err
	}
	cost.ID = int64(id)
	log.Printf("Updating costumer...")
	cost, err = costumer.Update(cost)
	log.Printf("Costumer was updated")
	context.Status(http.StatusOK).JSON(cost)
	log.Printf("Update costumer request was done")
	return err
}
