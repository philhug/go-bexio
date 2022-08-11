package bexio

import "fmt"

type InvoiceService struct {
	client *Client
}

func (c *InvoiceService) ListInvoices() ([]Invoice, error) {
	req, err := c.client.newRequest("GET", "kb_invoice", nil)
	if err != nil {
		return nil, err
	}
	var invoice []Invoice
	_, err = c.client.do(req, &invoice)
	return invoice, err
}

func (c *InvoiceService) GetInvoice(id int) (Invoice, error) {
	var invoice Invoice
	req, err := c.client.newRequest("GET", fmt.Sprintf("kb_invoice/%d", id), nil)
	if err != nil {
		return invoice, err
	}
	_, err = c.client.do(req, &invoice)
	return invoice, err
}

func (c *InvoiceService) CreateInvoice(invoice Invoice) (Invoice, error) {
	req, err := c.client.newRequest("POST", "kb_invoice", invoice)
	if err != nil {
		return invoice, err
	}
	_, err = c.client.do(req, &invoice)
	return invoice, err
}
