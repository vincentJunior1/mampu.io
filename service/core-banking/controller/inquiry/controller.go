package inquiry

import "github.com/gin-gonic/gin"

func InitInquiryController(cfg ControllerConfig) (ControllerInterface, error) {
	return &controller{}, nil
}
func (c *controller) Post(ctx *gin.Context)    {}
func (c *controller) GetByID(ctx *gin.Context) {}
func (c *controller) GetAll(ctx *gin.Context)  {}
