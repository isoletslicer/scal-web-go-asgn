package controller

import (
	"net/http"
	"strconv"
	"strings"

	"finalprojekmygram/apps/model/domain"
	"finalprojekmygram/apps/model/request"
	"finalprojekmygram/apps/model/response"
	"finalprojekmygram/apps/service"

	utilities "finalprojekmygram/internal/utilities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type CommentControllerImpl struct {
	CommentService service.CommentService
	PhotoService   service.PhotoService
}

func NewCommentController(commentService service.CommentService, photoService service.PhotoService) CommentController {
	return &CommentControllerImpl{CommentService: commentService, PhotoService: photoService}
}

func (commentController *CommentControllerImpl) Create(ctx *gin.Context) {

	var req request.CommentCreateRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := err.(validator.ValidationErrors)
		fieldErrorResponse := make(map[string]interface{})

		for _, v := range validationError {
			fieldErrorResponse[strings.ToLower(v.Field())] = utilities.GetValidationErrorMsg(v)
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  fieldErrorResponse,
		})

		return
	}

	photoID := req.PhotoID

	if _, err := commentController.PhotoService.GetOne(photoID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Success: false,
			Message: "Not Found",
			Errors: gin.H{
				"message": "Record not found",
			},
		})

		return
	}

	comment := domain.Comment{
		Message: req.Message,
		PhotoID: req.PhotoID,
		UserID:  userID,
	}

	if err := commentController.CommentService.Create(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Comment Created!",
		Data: response.CommentCreateResponse{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PhotoID:   comment.PhotoID,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
		},
	})
}

func (commentController *CommentControllerImpl) GetAll(ctx *gin.Context) {

	comments, err := commentController.CommentService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})
	}

	commentsResponse := []response.CommentGetAllResponse{}
	for _, comment := range comments {
		commentsResponse = append(commentsResponse, response.CommentGetAllResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdatedAt: comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User: response.CommentUserGetAllResponse{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: response.CommentPhotoGetAllResponse{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserID:   comment.Photo.UserID,
			},
		})
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Get All comments!",
		Data:    commentsResponse,
	})
}

func (commentController *CommentControllerImpl) Update(ctx *gin.Context) {

	var (
		req   request.CommentUpdateRequest
		photo domain.Photo
		err   error
	)

	commentID, _ := strconv.ParseUint(ctx.Param("commentId"), 10, 32)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if err = ctx.ShouldBindJSON(&req); err != nil {
		validationError := err.(validator.ValidationErrors)
		fieldErrorResponse := make(map[string]interface{})

		for _, v := range validationError {
			fieldErrorResponse[strings.ToLower(v.Field())] = utilities.GetValidationErrorMsg(v)
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  fieldErrorResponse,
		})

		return
	}

	comment := domain.Comment{
		ID:      uint(commentID),
		UserID:  userID,
		Message: req.Message,
	}

	if photo, err = commentController.CommentService.Update(comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Comment Updated!",
		Data: response.CommentUpdateResponse{
			ID:        photo.ID,
			UserID:    photo.UserID,
			Title:     photo.Title,
			PhotoUrl:  photo.PhotoUrl,
			Caption:   photo.Caption,
			UpdatedAt: photo.UpdatedAt,
		},
	})
}

func (commentController *CommentControllerImpl) Delete(ctx *gin.Context) {

	commentID, _ := strconv.ParseUint(ctx.Param("commentId"), 10, 32)

	if err := commentController.CommentService.Delete(uint(commentID)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Comment Deleted!",
		Data: response.CommentDeleteResponse{
			Message: "Your comment has been successfully deleted",
		},
	})
}
