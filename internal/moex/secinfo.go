package moexparser

import (
	"encoding/xml"
	"strconv"

	"github.com/boldlogic/moex-connector/internal/models"
	"github.com/boldlogic/packages/utils/dates"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type securityDescription struct {
	XMLName xml.Name       `xml:"document"`
	Data    []securityData `xml:"data"`
}

type securityData struct {
	ID   string                   `xml:"id,attr"`
	Rows []securityDescriptionRow `xml:"rows>row"`
}

type securityDescriptionRow struct {
	Name      string `xml:"name,attr"`
	Title     string `xml:"title,attr"`
	Value     string `xml:"value,attr"`
	Type      string `xml:"type,attr"`
	SortOrder string `xml:"sort_order,attr"`
	IsHidden  string `xml:"is_hidden,attr"`
	Precision string `xml:"precision,attr"`
}

func (p *Parser) setSecurityField(row securityDescriptionRow, sec *models.Security) {
	if sec == nil {
		p.logger.Error("security nil")
		return
	}

	v := row.Value
	switch row.Name {
	case "SECID":
		sec.SecId = v
	case "NAME":
		sec.FullName = v
	case "SHORTNAME":
		sec.ShortName = v
	case "ISIN":
		isin := v
		if v != "" {
			sec.ISIN = &isin
		}
	case "REGNUMBER":
		sec.RegNumber = v
	case "TYPENAME":
		sec.TypeName = v
	case "GROUP":
		sec.GroupCode = v
	case "TYPE":
		sec.TypeCode = v
	case "GROUPNAME":
		sec.GroupName = v

		//Облигации
	case "INITIALFACEVALUE":
		initialFaceValue, err := decimal.NewFromString(v)
		if err != nil {
			p.logger.Warn("INITIALFACEVALUE в security не распознан", zap.Error(err))
			return
		}
		sec.InitialFaceValue = &initialFaceValue

	case "FACEVALUE":
		faceValue, err := decimal.NewFromString(v)
		if err != nil {
			p.logger.Warn("FACEVALUE в security не распознан", zap.Error(err))
			return
		}
		sec.FaceValue = &faceValue

	case "FACEUNIT":
		sec.FaceUnit = v

	case "COUPONFREQUENCY":
		couponFrequency, err := strconv.Atoi(v)
		if err != nil {
			p.logger.Warn("COUPONFREQUENCY в security не распознан", zap.Error(err))

			return
		}
		sec.CouponFrequency = &couponFrequency

	case "COUPONVALUE":
		couponValue, err := decimal.NewFromString(v)
		if err != nil {
			p.logger.Warn("COUPONVALUE в security не распознан", zap.Error(err))

			return
		}

		sec.CouponValue = &couponValue

	case "MATDATE":
		parsed, err := dates.OptionalDatePtr(v, dates.ISODateFormat)
		if err != nil {
			p.logger.Warn("MATDATE в security не распознан", zap.Error(err))
		}

		sec.MaturityDate = parsed

	case "ISSUEDATE":
		parsed, err := dates.OptionalDatePtr(v, dates.ISODateFormat)
		if err != nil {
			p.logger.Warn("ISSUEDATE в security не распознан", zap.Error(err))
		}

		sec.IssueDate = parsed
	}

}

func (p *Parser) securityDataToSecurity(d securityData) models.Security {
	security := models.Security{}

	for _, row := range d.Rows {
		p.setSecurityField(row, &security)
	}
	return security
}

func (p *Parser) SecurityDescriptionXML(bdy []byte) (models.Security, error) {
	var raw securityDescription
	err := xml.Unmarshal(bdy, &raw)
	if err != nil {
		p.logger.Error("ошибка разбора XML security", zap.Error(err))
		return models.Security{}, err
	}

	var security models.Security
	for _, data := range raw.Data {
		if data.ID == "description" {
			security = p.securityDataToSecurity(data)
		}
	}
	return security, nil

}
