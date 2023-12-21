package middlewares

import (
	"net/http"
	"strconv"

	"finalprojekmygram/apps/model/domain"
	"finalprojekmygram/apps/model/response"
	"finalprojekmygram/apps/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CommentAuthz(commentService service.CommentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			comment domain.Comment
			err     error
		)

		commentID, _ := strconv.ParseUint(ctx.Param("commentId"), 10, 32)
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if comment, err = commentService.GetOne(uint(commentID)); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
				Success: false,
				Message: "Bad Request",
				Errors: gin.H{
					"message": "Comment not found",
				},
			})

			return
		}

		if comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
				Success: false,
				Message: "Forbidden",
				Errors: gin.H{
					"message": "You don't have permission",
				},
			})

			return
		}
	}
}

func SocialMediaAuthz(socialMediaService service.SocialMediaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			socialMedia domain.SocialMedia
			err         error
		)

		socialMediaId, _ := strconv.Atoi(ctx.Param("id"))
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if socialMedia, err = socialMediaService.GetOne(uint(socialMediaId)); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
				Success: false,
				Message: "Bad Request",
				Errors: gin.H{
					"message": "Social media not found",
				},
			})

			return
		}

		if socialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
				Success: false,
				Message: "Forbidden",
				Errors: gin.H{
					"message": "You don't have permission",
				},
			})

			return
		}
	}
}

func PhotoAuthz(photoService service.PhotoService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			photo domain.Photo
			err   error
		)

		photoID, _ := strconv.Atoi(ctx.Param("id"))
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if photo, err = photoService.GetOne(uint(photoID)); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
				Success: false,
				Message: "Bad Request",
				Errors: gin.H{
					"message": "Photo not found",
				},
			})

			return
		}

		if photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
				Success: false,
				Message: "Forbidden",
				Errors: gin.H{
					"message": "You don't have permission",
				},
			})

			return
		}
	}
}
