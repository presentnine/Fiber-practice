package handlers

import (
	"fiber-practice/database"
	"fiber-practice/dtos"
	"fiber-practice/models"
	"github.com/gofiber/fiber/v2"
)

func UserGet(c *fiber.Ctx) error {
	user := new(models.User)

	if err := database.DB.First(&user, c.Params("id")).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

func UserCreate(c *fiber.Ctx) error {
	userCreateRequestDto := new(dtos.UserCreateRequestDto)

	if err := c.BodyParser(userCreateRequestDto); err != nil {
		return err
	}

	user := models.User{
		Nickname: userCreateRequestDto.Nickname,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success":              true,
		"userCreateRequestDto": user,
	})
}

func UserUpdate(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	userUpdateRequestDto := new(dtos.UserUpdateRequestDto)
	if err := c.BodyParser(userUpdateRequestDto); err != nil {
		return err
	}

	user := new(models.User)
	if err := database.DB.First(&user, userId).Error; err != nil {
		return err
	}
	user.Nickname = userUpdateRequestDto.Nickname

	if err := database.DB.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

func UserDelete(c *fiber.Ctx) error {
	if err := database.DB.First(&models.User{}, c.Params("id")).Error; err != nil {
		return err
	}
	if err := database.DB.Delete(&models.User{}, c.Params("id")).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
