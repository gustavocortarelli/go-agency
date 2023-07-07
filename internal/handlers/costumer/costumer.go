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

	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid ID on request",
		})
		return err
	}

	costumerData, err := costumer.Get(int64(id))

	context.Status(http.StatusOK).JSON(costumerData)

	return nil
}

func GetAll(context *fiber.Ctx) error {
	costumers, err := costumer.GetAll()
	if err != nil {
		log.Printf("Error during get all costumers: %d", err)
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Could not get costumers",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(costumers)
	return nil
}

func Create(context *fiber.Ctx) error {
	cost := model.Costumer{}
	err := context.BodyParser(&cost)
	if err != nil {
		log.Printf("Error parsing body data: %d", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Error parsing body data: invalid request",
		})
		return err
	}

	cost, err = costumer.Create(cost)
	context.Status(http.StatusOK).JSON(cost)
	return err
}

func Delete(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid ID on delete request",
		})
		return err
	}
	err = costumer.Delete(int64(id))
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Costumer was deleted"})
	return err
}

func Update(context *fiber.Ctx) error {
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
	cost, err = costumer.Update(cost)
	context.Status(http.StatusOK).JSON(cost)
	return err
}
