package controllers

import (
	"app/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = models.GetCollection(models.DB, "users")
var validate = validator.New()

func UserNameParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		x := fmt.Sprintf("my name is %s mohite", name)
		c.JSON(http.StatusOK, gin.H{
			"message": x,
		})
	}
}

func UserNameQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		x := fmt.Sprintf("my real name is %s", name)
		c.JSON(http.StatusOK, gin.H{
			"message": x,
		})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": validationErr.Error(),
			})
			return
		}

		newUser := models.User{
			Id:   primitive.NewObjectID(),
			Name: user.Name,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func GetAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
			}

			user = append(user, singleUser)
		}

		c.JSON(http.StatusOK, user)

	}

}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": validationErr.Error(),
			})
			return
		}
		objId, _ := primitive.ObjectIDFromHex(userId)

		update := bson.M{"name": user.Name}

		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if result.MatchedCount == 1 && result.ModifiedCount == 0 {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Name Already exists",
			})
			return

		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Updated Successfully",
		})

	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.Param("userId")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		if result.DeletedCount == 0 {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Not Found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted Successfully",
		})

	}
}
