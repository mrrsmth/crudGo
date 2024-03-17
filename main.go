package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Roadmap struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Technology string `json:"technology"`
	Theme      string `json:"theme"`
	Bool       bool   `json:"bool"`
}

func main() {
	db, err := gorm.Open("mysql", "root:secret@tcp(localhost:3306)/road?parseTime=true")
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&Roadmap{})

	roadmap := []Roadmap{
		{Technology: "HTML", Theme: "Элементы в HTML", Bool: false},
		{Technology: "", Theme: "Формы, валидация форм", Bool: false},
		{Technology: "", Theme: "Семантическая верстка", Bool: false},
		{Technology: "CSS", Theme: "Селекторы", Bool: false},
		{Technology: "", Theme: "Свойства", Bool: false},
		{Technology: "", Theme: "Позиционирование элементов, Flexbox, Grid", Bool: false},
		{Technology: "", Theme: "Трансформации, переходы, анимации", Bool: false},
		{Technology: "", Theme: "Адаптивный дизайн и медиазапросы", Bool: false},
		{Technology: "", Theme: "CSS-препроцессоры(sass, scss, less)", Bool: false},
		{Technology: "", Theme: "БЭМ", Bool: false},
		{Technology: "JavaScript", Theme: "Типы данных, преобразования типов", Bool: false},
		{Technology: "", Theme: "Условное ветвление, логические операторы, циклы", Bool: false},
		{Technology: "", Theme: "Функции, функциональные выражения, стрелочные функции, поднятие", Bool: false},
		{Technology: "", Theme: "Замыкание, IIFE", Bool: false},
		{Technology: "", Theme: "Строки, шаблонные строки, регулярные выражения", Bool: false},
		{Technology: "", Theme: "Массивы, методы массивов, перебор массивов", Bool: false},
		{Technology: "", Theme: "Объекты, методы объектов, сравнение объектов, ссылки", Bool: false},
		{Technology: "", Theme: "Классы, наследование, статические свойства и методы, защищенные свойства и методы", Bool: false},
		{Technology: "", Theme: "Колбэки, промисы, обработка ошибок, микротаски, азупс/амаи, event loop", Bool: false},
		{Technology: "", Theme: "Взаимодействие DOM (создание, добавление, изменение и удаление элементов веб-станиць), браузерные события, распространение событий", Bool: false},
		{Technology: "", Theme: "Хранение данных - сооке, session storage, local storage", Bool: false},
		{Technology: "", Theme: "Дебаг в Chrome DevTools", Bool: false},
		{Technology: "", Theme: "XMLHttpRequest, FetchApi, WebSocket", Bool: false},
		{Technology: "Angular", Theme: "Компоненты, модули, загрузка приложения", Bool: false},
		{Technology: "", Theme: "Привязка данных", Bool: false},
		{Technology: "", Theme: "Привязка к событиям дочернего компонента", Bool: false},
		{Technology: "", Theme: "Жизненный цикл компонента", Bool: false},
		{Technology: "", Theme: "Атрибутивные и структурные директивы, создание директив", Bool: false},
		{Technology: "", Theme: "Сервисы и dependency injection", Bool: false},
		{Technology: "", Theme: "Работа с формами", Bool: false},
		{Technology: "", Theme: "НТТР и взаимодействие с сервером", Bool: false},
		{Technology: "", Theme: "Маршрутизация", Bool: false},
		{Technology: "", Theme: "Pipes", Bool: false},
	}

	for _, data := range roadmap {
		db.Create(&data)
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.GET("/roadmap", func(c *gin.Context) {
		var roadmap []Roadmap
		db.Find(&roadmap)

		c.JSON(http.StatusOK, roadmap)
	})

	router.POST("/roadmap", func(c *gin.Context) {
		var newRoadmap Roadmap

		if err := c.ShouldBindJSON(&newRoadmap); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&newRoadmap)
		c.JSON(http.StatusOK, newRoadmap)
	})

	router.PUT("/roadmap/:id", func(c *gin.Context) {
		id := c.Param("id")

		var updatedRoadmap Roadmap
		if err := c.ShouldBindJSON(&updatedRoadmap); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Model(&Roadmap{}).Where("id = ?", id).Updates(updatedRoadmap)
		c.JSON(http.StatusOK, updatedRoadmap)
	})

	router.DELETE("/roadmap/:id", func(c *gin.Context) {
		id := c.Param("id") // Получаем идентификатор из URL-параметра
		var roadmap Roadmap
		db.Delete(&roadmap, id) // Удаляем запись из базы данных

		c.Status(http.StatusNoContent) // Возвращаем статус "204 No Content" для успешного удаления
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}
