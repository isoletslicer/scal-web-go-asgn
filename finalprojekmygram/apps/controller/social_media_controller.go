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

type SocialMediaController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type SocialMediaControllerImpl struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) SocialMediaController {
	return &SocialMediaControllerImpl{SocialMediaService: socialMediaService}
}

func (socialMediaController *SocialMediaControllerImpl) Create(ctx *gin.Context) {
	var req request.SocialMediaCreateRequest

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

	socialMedia := domain.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         userID,
	}

	if err := socialMediaController.SocialMediaService.Create(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Social Media Created!",
		Data: response.SocialMediaCreateResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
		},
	})
}

func (socialMediaController *SocialMediaControllerImpl) GetAll(ctx *gin.Context) {

	var (
		socialMedias         []domain.SocialMedia
		err                  error
		socialMediasResponse = []response.SocialMediaGetAllResponse{}
	)

	if socialMedias, err = socialMediaController.SocialMediaService.GetAll(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	for _, socialMedia := range socialMedias {
		socialMediasResponse = append(socialMediasResponse, response.SocialMediaGetAllResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User: response.SocialMediaUserGetAllResponse{
				ID:       socialMedia.User.ID,
				Username: socialMedia.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Get All Social Media!",
		Data: gin.H{
			"social_medias": socialMediasResponse,
		},
	})
}

func (socialMediaController *SocialMediaControllerImpl) Update(ctx *gin.Context) {

	var (
		req                domain.SocialMedia
		updatedSocialMedia domain.SocialMedia
		err                error
	)

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
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

	socialMedia := domain.SocialMedia{
		ID:             uint(id),
		UserID:         userID,
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
	}

	if updatedSocialMedia, err = socialMediaController.SocialMediaService.Update(socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Social Media Updated!",
		Data: response.SocialMediaUpdateResponse{
			ID:             updatedSocialMedia.ID,
			Name:           updatedSocialMedia.Name,
			SocialMediaUrl: updatedSocialMedia.SocialMediaUrl,
			UserID:         updatedSocialMedia.UserID,
			UpdatedAt:      updatedSocialMedia.UpdatedAt,
		},
	})
}

func (socialMediaController *SocialMediaControllerImpl) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err := socialMediaController.SocialMediaService.Delete(uint(id)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Social Media Deleted!",
		Data: response.SocialMediaDeleteResponse{
			Message: "Your social media has been successfully deleted",
		},
	})
}
