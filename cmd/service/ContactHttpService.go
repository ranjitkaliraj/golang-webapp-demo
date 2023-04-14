package service

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranjitkaliraj/go-test-crm-service/cmd/contacts"
)

type ContactHttpService struct {
	commandHandler *contacts.ContactCommandHandler
	query          *contacts.ContactQuery
}

func NewContactHttpService() *ContactHttpService {
	repository := contacts.NewContactRepository()
	return &ContactHttpService{commandHandler: contacts.NewContactCommandHandler(repository), query: contacts.NewContactQuery(repository)}
}

func (h *ContactHttpService) Create(c *gin.Context) {
	var contact contacts.Contact
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.commandHandler.HandleCreateCommand(context.Background(), &contacts.AddContactCommand{Contact: contact})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *ContactHttpService) Update(c *gin.Context) {
	id := c.Param("id")
	var update contacts.Contact
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.commandHandler.HandleUpdateCommand(context.Background(), &contacts.UpdateContactCommand{Id: id, Contact: update})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *ContactHttpService) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.commandHandler.HandleDeleteCommand(context.Background(), &contacts.DeleteContactCommand{ContactId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *ContactHttpService) Get(c *gin.Context) {
	id := c.Param("id")
	result, err := h.query.GetContact(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *ContactHttpService) List(c *gin.Context) {
	results, err := h.query.GetContacts(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
