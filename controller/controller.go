package controller

import (
	"context"
	"fmt"
	"log"
	pb "myapp/grpc"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	grpc "google.golang.org/grpc"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	client pb.GroupServiceClient
	db     *gorm.DB
)

// GetClient is ..
func GetClient() {
	conn, err := grpc.Dial("149.28.91.164:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client = pb.NewGroupServiceClient(conn)
}

// ConnectDB is ..
func ConnectDB() {
	dsn := "root:654321@tcp(149.28.91.164:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		db = database
	}
	db.AutoMigrate(&models.User{})
	fmt.Printf("connet succeed %v", db)
}

// GetUser is ..
func GetUser(c *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// user.ID = id

	if DbResult := db.Where("id = ?", id).First(&user); DbResult.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"CreateUser": "fail",
			"Error":      DbResult.Error.Error(),
		})
		return
	}
	// user.Name = c.Query("name")
	// user.GroupID = c.Query("group_id")

	log.Println(user)

	result, err := client.GetGroup(context.Background(), &pb.GroupRequest{GroupID: user.GroupID, GroupName: ""})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"ID":    user.ID,
		"Name":  user.Name,
		"Group": result.GroupName,
	})
}

// CreateUser is ..
func CreateUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(user)

	if DbResult := db.Create(&user); DbResult.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"CreateUser": "fail",
			"Error":      DbResult.Error.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"CreateUser": "succeed",
		})
	}

}

// UpdateUser is ..
func UpdateUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if DbResult := db.Save(&user); DbResult.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"UpdateUser": "fail",
			"Error":      DbResult.Error.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"UpdateUser": "succeed",
		})
	}

}

// DeleteUser is ..
func DeleteUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if DbResult := db.Where("id = ?", user.ID).Delete(&user); DbResult.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"DeleteUser": "fail",
			"Error":      DbResult.Error.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"DeleteUser": "succeed",
		})
	}

}

// GetGroup is ..
func GetGroup(c *gin.Context) {
	//var group models.UserGroup
	var pbGroup pb.GroupRequest

	pbGroup.GroupID = c.Query("group_id")
	pbGroup.GroupName = c.Query("group_name")

	result, err := client.GetGroup(context.Background(), &pbGroup)
	if err != nil {
		log.Fatal(err)
	}
	if result.Status {
		log.Printf("true")
	} else {
		log.Printf("false")
	}

	log.Printf(result.GroupID)
	log.Printf(result.GroupName)
	c.JSON(http.StatusOK, gin.H{
		"GetGroup": true,
	})
}

// CreateGroup is ..
func CreateGroup(c *gin.Context) {
	var group models.UserGroup
	var pbGroup pb.GroupRequest

	// pbGroup.GroupID = c.Query("group_id")
	// pbGroup.GroupName = c.Query("group_name")

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println("show json : " + group.GroupID)
	pbGroup.GroupID = group.GroupID
	pbGroup.GroupName = group.GroupName
	result, err := client.CreateGroup(context.Background(), &pbGroup)
	if err != nil {
		log.Fatal(err)
	}
	if result.Status {
		log.Printf("true")
	} else {
		log.Printf("false")
	}
	log.Printf(result.GroupID)
	log.Printf(result.GroupName)
	c.JSON(http.StatusOK, gin.H{
		"CreateGroup": true,
	})

}

// DeleteGroup is ..
func DeleteGroup(c *gin.Context) {
	var pbGroup pb.GroupRequest

	// pbGroup.GroupID = c.Query("group_id")
	// pbGroup.GroupName = c.Query("group_name")

	if err := c.ShouldBindJSON(&pbGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := client.DeleteGroup(context.Background(), &pbGroup)
	if err != nil {
		log.Fatal(err)
	}
	if result.Status {
		log.Printf("true")
	} else {
		log.Printf("false")
	}
	log.Printf(result.GroupID)
	log.Printf(result.GroupName)
	c.JSON(http.StatusOK, gin.H{
		"CreateGroup": true,
	})
}
