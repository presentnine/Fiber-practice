package handlers

import (
	"fiber-practice/database"
	"fiber-practice/dtos"
	"fiber-practice/models"
	"github.com/gofiber/fiber/v2"
)

func UserGet(c *fiber.Ctx) error {
	db := database.Get()
	user := new(models.User)

	if err := db.First(&user, c.Params("id")).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

func UserCreate(c *fiber.Ctx) error {
	db := database.Get()
	userCreateRequestDto := new(dtos.UserCreateRequestDto)

	if err := c.BodyParser(userCreateRequestDto); err != nil {
		return err
	}

	user := models.User{
		Nickname: userCreateRequestDto.Nickname,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success":              true,
		"userCreateRequestDto": user,
	})
}

func UserUpdate(c *fiber.Ctx) error {
	db := database.Get()

	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	userUpdateRequestDto := new(dtos.UserUpdateRequestDto)
	if err := c.BodyParser(userUpdateRequestDto); err != nil {
		return err
	}

	user := new(models.User)
	if err := db.First(&user, userId).Error; err != nil {
		return err
	}
	user.Nickname = userUpdateRequestDto.Nickname

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

func UserDelete(c *fiber.Ctx) error {
	db := database.Get()

	if err := db.First(&models.User{}, c.Params("id")).Error; err != nil {
		return err
	}

	if err := db.Delete(&models.User{}, c.Params("id")).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
