package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type Security struct {
	//Общие атрибуты
	SecId     string  // Код ценной бумаги
	ISIN      *string // ISIN код
	FullName  string  // Полное наименование
	ShortName string  // Краткое наименование
	RegNumber string  // Номер государственной регистрации

	TypeCode string // Тип бумаги
	TypeName string // Вид/категория ценной бумаги

	GroupCode string // Код типа инструмента
	GroupName string // Типа инструмента

	IssueDate *time.Time // Дата начала торгов

	// Облигации
	FaceUnit         string           // Валюта номинала
	InitialFaceValue *decimal.Decimal // Первоначальная номинальная стоимость
	FaceValue        *decimal.Decimal // Номинальная стоимость

	CouponValue     *decimal.Decimal // Сумма купона, в валюте номинала
	CouponFrequency *int             // Периодичность выплаты купона в год

	MaturityDate *time.Time // Дата погашения

}

func (s Security) String() string {

	isin := "<nil>"
	if s.ISIN != nil {
		isin = *s.ISIN
	}

	issue := "<nil>"
	if s.IssueDate != nil {
		issue = s.IssueDate.Format(time.DateOnly)
	}
	initialFaceValue := "<nil>"
	if s.InitialFaceValue != nil {
		initialFaceValue = s.InitialFaceValue.String()
	}
	faceValue := "<nil>"
	if s.FaceValue != nil {
		faceValue = s.FaceValue.String()
	}

	couponValue := "<nil>"
	if s.CouponValue != nil {
		couponValue = s.CouponValue.String()
	}
	mat := "<nil>"
	if s.MaturityDate != nil {
		mat = s.MaturityDate.Format(time.DateOnly)
	}

	couponFrequency := "<nil>"
	if s.CouponFrequency != nil {
		couponFrequency = strconv.Itoa(*s.CouponFrequency)
	}
	return fmt.Sprintf(
		"Security{SecId:%q ISIN:%q FullName:%q ShortName:%q RegNumber:%q TypeCode:%q TypeName:%q GroupCode:%q GroupName:%q IssueDate:%s "+
			"FaceUnit:%q InitialFaceValue:%s FaceValue:%s CouponValue:%s CouponFrequency:%s MaturityDate:%s}",
		s.SecId, isin, s.FullName, s.ShortName, s.RegNumber, s.TypeCode, s.TypeName, s.GroupCode, s.GroupName, issue,
		s.FaceUnit, initialFaceValue, faceValue, couponValue, couponFrequency, mat,
	)
}
