package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      	int  			
	Name   		string    		`json:"name" binding:"required,min=5"`
	Email  		string 			`json:"email" binding:"required,min=5"`
}

  




  
  // Add name field