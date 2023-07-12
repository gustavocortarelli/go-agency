package costumer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavocortarelli/go-agency/internal/model"
	"github.com/gustavocortarelli/go-agency/internal/service/costumer"
	"log"
	"net/http"
	"strconv"
)

// Get Return costumer data by ID
//
// @Summary		Get costumer by ID
// @Description	Get costumer
// @Tags			costumer
// @Produce		json
// @Param			id	query		int	true	"Costumer ID"
// @Success		200	{object}	model.Costumer
// @Failure		400	{object}	model.Error
// @Failure		404	{object}	model.Error
// @Failure		500	{object}	model.Error
// @Router			/costumer/:id [get]
func Get(context *fiber.Ctx) error {
	log.Printf("Starting get costumer by ID request")
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		context.Status(http.StatusBadRequest).JSON(model.Error{
			Message: "Invalid ID on request",
			Code:    http.StatusBadRequest,
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

// GetAll Return All costumers
//
// @Summary		Return all costumers
// @Description	Return all costumers that was created on database
// @Tags			costumer
// @Produce		json
// @Success		200	{object}	model.CostumerData
// @Failure		400	{object}	model.Error
// @Failure		404	{object}	model.Error
// @Failure		500	{object}	model.Error
// @Router			/costumer [get]
func GetAll(context *fiber.Ctx) error {
	log.Printf("Starting get all costumers request")
	costumers, err := costumer.GetAll()
	log.Printf("Data was fetched")
	if err != nil {
		log.Printf("Error during get all costumers: %d", err)
		context.Status(http.StatusInternalServerError).JSON(model.Error{
			Message: "Could not get costumers",
			Code:    http.StatusInternalServerError,
		})
		return err
	}
	context.Status(http.StatusOK).JSON(costumers)
	log.Printf("Get all costumers request was done")
	return nil
}

// Create Insert costumer data on database
//
// @Summary		Create costumer
// @Description	Create a costumer record on database
// @Tags			costumer
// @Produce		json
// @Param			request body model.Costumer true "Costumer data"
// @Success		200	{object}	model.Costumer
// @Failure		400	{object}	model.Error
// @Failure		404	{object}	model.Error
// @Failure		500	{object}	model.Error
// @Router			/costumer [post]
func Create(context *fiber.Ctx) error {
	log.Printf("Starting create costumer request")
	cost := model.Costumer{}
	err := context.BodyParser(&cost)
	if err != nil {
		log.Printf("Error parsing body data: %d", err)
		context.Status(http.StatusBadRequest).JSON(model.Error{
			Message: "Error parsing body data: invalid request",
			Code:    http.StatusBadRequest,
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

// Delete Remove costumer record by ID
//
// @Summary		Remove costumer by ID
// @Description	Remove costumer data from database by ID
// @Tags			costumer
// @Produce		json
// @Param			id	query		int	true	"Costumer ID"
// @Success		200	{object}	model.Success
// @Failure		400	{object}	model.Error
// @Failure		404	{object}	model.Error
// @Failure		500	{object}	model.Error
// @Router			/costumer/:id [delete]
func Delete(context *fiber.Ctx) error {
	log.Printf("Starting delete costumer request")
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		context.Status(http.StatusBadRequest).JSON(model.Error{
			Message: "Invalid ID on delete request",
			Code:    http.StatusBadRequest,
		})
		return err
	}
	log.Printf("Deleting costumer...")
	err = costumer.Delete(int64(id))
	log.Printf("Costumer was deleted")
	context.Status(http.StatusOK).JSON(model.Success{Message: "Costumer was deleted"})
	log.Printf("Delete costumer request was done")
	return err
}

// Update Update costumer data
//
// @Summary		Update costumer data by ID
// @Description	Remove costumer data from database by ID
// @Tags			costumer
// @Produce		json
// @Param			id				query		int				true	"Costumer ID"
// @Param			request 		body 		model.Costumer 	true 	"Costumer data"
// @Success		200	{object}	model.Costumer
// @Failure		400	{object}	model.Error
// @Failure		404	{object}	model.Error
// @Failure		500	{object}	model.Error
// @Router			/costumer/:id [put]
func Update(context *fiber.Ctx) error {
	log.Printf("Starting update costumer request")
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		context.Status(http.StatusBadRequest).JSON(model.Error{
			Message: "Invalid ID on delete request",
			Code:    http.StatusBadRequest,
		})
		return err
	}

	cost := model.Costumer{}
	err = context.BodyParser(&cost)
	if err != nil {
		log.Printf("Error parsing body data: %d", err)
		context.Status(http.StatusBadRequest).JSON(model.Error{
			Message: "Error parsing body data: invalid request",
			Code:    http.StatusBadRequest,
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
