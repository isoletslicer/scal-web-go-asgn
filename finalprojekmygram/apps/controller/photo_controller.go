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

type PhotoController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PhotoControllerImpl struct {
	PhotoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) PhotoController {
	return &PhotoControllerImpl{PhotoService: photoService}
}

func (photoController *PhotoControllerImpl) Create(ctx *gin.Context) {

	var req request.PhotoCreateRequest

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

	photo := domain.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   userID,
	}

	if err := photoController.PhotoService.Create(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Photo Created!",
		Data: response.PhotoCreateResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
		},
	})
}

func (photoController *PhotoControllerImpl) GetAll(ctx *gin.Context) {

	photos, err := photoController.PhotoService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})
	}

	photosResponse := []response.PhotoGetAllResponse{}
	for _, photo := range photos {
		photosResponse = append(photosResponse, response.PhotoGetAllResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: response.PhotoUserGetAllReponse{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Get All Photos!",
		Data:    photosResponse,
	})
}

func (photoController *PhotoControllerImpl) Update(ctx *gin.Context) {
	var (
		req          request.PhotoUpdateRequest
		updatedPhoto domain.Photo
		err          error
	)

	photoID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

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

	photo := domain.Photo{
		ID:       uint(photoID),
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}

	if updatedPhoto, err = photoController.PhotoService.Update(photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Photo Updated!",
		Data: response.PhotoUpdateResponse{
			ID:        updatedPhoto.ID,
			UserID:    updatedPhoto.UserID,
			Title:     updatedPhoto.Title,
			PhotoUrl:  updatedPhoto.PhotoUrl,
			Caption:   updatedPhoto.Caption,
			UpdatedAt: updatedPhoto.UpdatedAt,
		},
	})
}

func (photoController *PhotoControllerImpl) Delete(ctx *gin.Context) {
	photoID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err := photoController.PhotoService.Delete(uint(photoID)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Photo Deleted!",
		Data: response.PhotoDeleteResponse{
			Message: "your photo has been successfully deleted",
		},
	})
}
