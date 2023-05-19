package controllers

import (
	"github.com/RianIhsan/ex-go-crud-icc/database"
	"github.com/RianIhsan/ex-go-crud-icc/models"
	"github.com/gofiber/fiber/v2"
)

func Reads(c *fiber.Ctx) error {
	var user []models.User

	database.DB.Find(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data user tersedia",
		"user":    user,
	})
}

func Read(c *fiber.Ctx) error {
	var user models.User

	userId := c.Params("id")
	if userId == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id tidak boleh kosong",
		})
	}

	if err := database.DB.Where("id = ? ", userId).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data user tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data user tersedia",
		"user":    user,
	})
}

func Create(c *fiber.Ctx) error {
	var userReq models.UserReq

	if err := c.BodyParser(&userReq); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user models.User
	user.Nama = userReq.Nama
	user.Kelas = userReq.Kelas
	user.Semester = userReq.Semester
	user.Prodi = userReq.Prodi
	user.Wa = userReq.Wa

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil dibuat",
		"user":    user,
	})
}

func Update(c *fiber.Ctx) error {
	userUpdate := new(models.UserReq)

	if err := c.BodyParser(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error body parse update",
		})
	}

	userId := c.Params("id")
	user := models.User{}

	if err := database.DB.First(&user, "id = ? ", userId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data user tidak ditemuakn",
		})
	}

	user.Nama = userUpdate.Nama
	user.Kelas = userUpdate.Kelas
	user.Semester = userUpdate.Semester
	user.Prodi = userUpdate.Prodi
	user.Wa = userUpdate.Wa

	if err := database.DB.Save(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data gagal di simpan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil diupdate",
		"user":    user,
	})
}

func Delete(c *fiber.Ctx) error {
	user := models.User{}

	userId := c.Params("id")

	if err := database.DB.First(&user, userId).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data gagal dihapus",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}
