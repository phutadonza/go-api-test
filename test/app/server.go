package main

import (
	"test/app/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// สร้าง instance ของ Echo framework
	e := echo.New()

	// สร้าง handler
	userHandler := handlers.NewUserHandler()

	// เส้น API สำหรับการดึงข้อมูลผู้ใช้ทั้งหมด
	e.GET("/users", userHandler.GetAllUsers)

	// เส้น API สำหรับการดึงข้อมูลผู้ใช้จาก ID
	e.GET("/users/:id", userHandler.GetUserByID)

	// เส้น API สำหรับการสร้างผู้ใช้ใหม่
	e.POST("/users", userHandler.CreateUser)

	// เส้น API สำหรับการลบผู้ใช้
	e.DELETE("/users/:id", userHandler.DeleteUser)

	// เริ่มต้นเซิร์ฟเวอร์
	e.Start(":8000")
}
