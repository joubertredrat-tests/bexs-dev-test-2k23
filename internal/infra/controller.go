package infra

import (
	"errors"
	"fmt"
	"joubertredrat/bexs-dev-test-2k23/internal/application"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	QUERY_STRING_OFFSET       = "offset"
	QUERY_STRING_LIMIT        = "limit"
	DEFAULT_PAGINATION_OFFSET = 1
	DEFAULT_PAGINATION_LIMIT  = 10
)

type ApiBaseController struct {
}

func NewApiBaseController() ApiBaseController {
	return ApiBaseController{}
}

func (c ApiBaseController) HandleStatus(ctx *gin.Context) {
	t := time.Now()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   DatetimeCanonical(&t),
	})
}

func (c ApiBaseController) HandleNotFound(ctx *gin.Context) {
	t := time.Now()
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "page not found",
		"time":  DatetimeCanonical(&t),
	})
}

type PartnerController struct {
}

func NewPartnerController() PartnerController {
	return PartnerController{}
}

func (c PartnerController) HandleCreate(usecase application.UsecaseCreatePartner) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request CreatePartnerRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			responseWithError(ctx, err)
			return
		}

		partner, err := usecase.Execute(ctx, application.UsecaseCreatePartnerInput{
			ID:          request.ID,
			TradingName: request.TradingName,
			Document:    request.Document,
			Currency:    request.Currency,
		})

		if err != nil {
			switch err.(type) {
			case domain.ErrInvalidCurrency:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			case application.ErrPartnerAlreadyExists:
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": err.Error(),
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
			}
			return
		}

		ctx.JSON(http.StatusCreated, PartnerResponse{
			ID:          partner.ID,
			TradingName: partner.TradingName,
			Document:    partner.Document,
			Currency:    partner.Currency.Value,
		})
	}
}

type PaymentController struct {
}

func NewPaymentController() PaymentController {
	return PaymentController{}
}

func (c PaymentController) HandleList(usecase application.UsecaseListPayments) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		offset := ctx.DefaultQuery(QUERY_STRING_OFFSET, fmt.Sprintf("%d", DEFAULT_PAGINATION_OFFSET))
		limit := ctx.DefaultQuery(QUERY_STRING_LIMIT, fmt.Sprintf("%d", DEFAULT_PAGINATION_LIMIT))
		o64, err := strconv.ParseUint(offset, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid offset query string",
			})
			return
		}
		l64, err := strconv.ParseUint(limit, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid limit query string",
			})
			return
		}

		pagination := domain.NewPagination(uint(o64), uint(l64))

		payments, err := usecase.Execute(ctx, pagination)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		response := []PaymentResponse{}
		for _, payment := range payments {
			response = append(response, PaymentResponse{
				ID:            payment.ID,
				PartnerID:     payment.PartnerID,
				Amount:        payment.Amount.Value,
				ForeignAmount: payment.ForeignAmount.Value,
				Consumer: ConsumerResponse{
					Name:       payment.Consumer.Name,
					NationalID: payment.Consumer.NationalID,
				},
			})
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func (c PaymentController) HandleCreate(usecase application.UsecaseCreatePayment) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request CreatePaymentRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			responseWithError(ctx, err)
			return
		}

		payment, err := usecase.Execute(ctx, application.UsecaseCreatePaymentInput{
			PartnerID:          request.PartnerID,
			Amount:             request.Amount,
			ConsumerName:       request.Consumer.Name,
			ConsumerNationalID: request.Consumer.NationalID,
		})

		if err != nil {
			fmt.Println(err)
			switch err.(type) {
			case application.ErrPartnerNotFound:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			case application.ErrPaymentDuplicated:
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": err.Error(),
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
			}
			return
		}

		ctx.JSON(http.StatusCreated, PaymentResponse{
			ID:            payment.ID,
			PartnerID:     payment.PartnerID,
			Amount:        payment.Amount.Value,
			ForeignAmount: payment.ForeignAmount.Value,
			Consumer: ConsumerResponse{
				Name:       payment.Consumer.Name,
				NationalID: payment.Consumer.NationalID,
			},
		})
	}
}

func (c PaymentController) HandleGet(usecase application.UsecaseGetPayment) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ID := ctx.Param("id")
		if ID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Payment ID required",
			})
			return
		}

		payment, err := usecase.Execute(ctx, ID)

		if err != nil {
			fmt.Println(err)
			switch err.(type) {
			case application.ErrPaymentNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
			}
			return
		}

		ctx.JSON(http.StatusOK, PaymentResponse{
			ID:            payment.ID,
			PartnerID:     payment.PartnerID,
			Amount:        payment.Amount.Value,
			ForeignAmount: payment.ForeignAmount.Value,
			Consumer: ConsumerResponse{
				Name:       payment.Consumer.Name,
				NationalID: payment.Consumer.NationalID,
			},
		})
	}
}

func RegisterCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func responseWithError(c *gin.Context, err error) {
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		c.JSON(http.StatusBadRequest, gin.H{"errors": getValidatorErrors(verr)})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func getValidatorErrors(verr validator.ValidationErrors) []RequestValidationError {
	var errs []RequestValidationError

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}

		errs = append(errs, RequestValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}
