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
		"data":    user,
	})
}

func Read(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
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
		"data":    user,
	})
}

func Create(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
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
  c.Set("Access-Control-Allow-Origin", "*")
  user := new(models.UserReq)

  if err := c.BodyParser(&user); err != nil {
    return err
  }

  id := c.Params("id")
  
  if id == ""{
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"Id Tidak boleh koseng",
    })
  }



  updateUser := models.User{}

  if err := database.DB.First(&updateUser, "id = ?", id).Error; err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"User Tidak ada",
    })
  }

  updateUser.Nama = user.Nama
  updateUser.Kelas = user.Kelas
  updateUser.Semester = user.Semester
  updateUser.Prodi = user.Prodi
  updateUser.Wa = user.Wa


  if err := database.DB.Model(&updateUser).Error; err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"ERROR ID",
    })
  }

  result := database.DB.Where("id = ?", id).Updates(&updateUser)
  if result.Error != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"Tidak Bisa mendapatkan user",
    })
  }

  if result.RowsAffected == 0 {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"Tidak ada data yang diubah",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message":"Data berhasil diubah",
    "data":updateUser,
  })
}

func Delete(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
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
