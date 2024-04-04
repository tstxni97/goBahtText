package gobahttext

import (
	"testing"
)

func TestInputIsANumberThenCanConvertToString(t *testing.T) {
	result := BahtText(1)
	if result != "หนึ่งบาทถ้วน" {
		t.Errorf("Expected หนึ่งบาทถ้วน but got %s", result)
	}
}

func TestInputIntegerOneDigitCanConvert(t *testing.T) {
	result := BahtText(0.0)
	if result != "ศูนย์บาทถ้วน" {
		t.Errorf("Expected ศูนย์บาทถ้วน but got %s", result)
	}
	result = BahtText(1.00)
	if result != "หนึ่งบาทถ้วน" {
		t.Errorf("Expected หนึ่งบาทถ้วน but got %s", result)
	}
	result = BahtText(2.0)
	if result != "สองบาทถ้วน" {
		t.Errorf("Expected สองบาทถ้วน but got %s", result)
	}
	result = BahtText(5.00)
	if result != "ห้าบาทถ้วน" {
		t.Errorf("Expected ห้าบาทถ้วน but got %s", result)
	}
}

func TestNumberInMultipleCouldShowWordSip(t *testing.T) {
	result := BahtText(10.0)
	if result != "สิบบาทถ้วน" {
		t.Errorf("Expected สิบบาทถ้วน but got %s", result)
	}
	result = BahtText(20.0)
	if result != "ยี่สิบบาทถ้วน" {
		t.Errorf("Expected ยี่สิบบาทถ้วน but got %s", result)
	}
	result = BahtText(50.0)
	if result != "ห้าสิบบาทถ้วน" {
		t.Errorf("Expected ห้าสิบบาทถ้วน but got %s", result)
	}
	result = BahtText(11.0)
	if result != "สิบเอ็ดบาทถ้วน" {
		t.Errorf("Expected สิบเอ็ดบาทถ้วน but got %s", result)
	}
	result = BahtText(15.0)
	if result != "สิบห้าบาทถ้วน" {
		t.Errorf("Expected สิบห้าบาทถ้วน but got %s", result)
	}
	result = BahtText(32.0)
	if result != "สามสิบสองบาทถ้วน" {
		t.Errorf("Expected สิบห้าบาทถ้วน but got %s", result)
	}
}

func TestInputOneIntegerWithDecimalCanConvert(t *testing.T) {
	result := BahtText(1.10)
	if result != "หนึ่งบาทสิบสตางค์" {
		t.Errorf("Expected หนึ่งบาทสิบสตางค์ but got %s", result)
	}
	result = BahtText(1.01)
	if result != "หนึ่งบาทหนึ่งสตางค์" {
		t.Errorf("Expected หนึ่งบาทหนึ่งสตางค์ but got %s", result)
	}
	result = BahtText(1.02)
	if result != "หนึ่งบาทสองสตางค์" {
		t.Errorf("Expected หนึ่งบาทสองสตางค์ but got %s", result)
	}
	result = BahtText(1.03)
	if result != "หนึ่งบาทสามสตางค์" {
		t.Errorf("Expected หนึ่งบาทสามสตางค์ but got %s", result)
	}
	result = BahtText(2.50)
	if result != "สองบาทห้าสิบสตางค์" {
		t.Errorf("Expected สองบาทห้าสิบสตางค์ but got %s", result)
	}
	result = BahtText(3.75)
	if result != "สามบาทเจ็ดสิบห้าสตางค์" {
		t.Errorf("Expected สามบาทเจ็ดสิบห้าสตางค์ but got %s", result)
	}
	result = BahtText(4.81)
	if result != "สี่บาทแปดสิบเอ็ดสตางค์" {
		t.Errorf("Expected สี่บาทแปดสิบเอ็ดสตางค์ but got %s", result)
	}
	result = BahtText(2.21)
	if result != "สองบาทยี่สิบเอ็ดสตางค์" {
		t.Errorf("Expected สองบาทยี่สิบเอ็ดสตางค์ but got %s", result)
	}
}

func TestNumberHundredShouldShowRoi(t *testing.T) {
	result := BahtText(100.0)
	if result != "หนึ่งร้อยบาทถ้วน" {
		t.Errorf("Expected หนึ่งร้อยบาทถ้วน but got %s", result)
	}
	result = BahtText(101.0)
	if result != "หนึ่งร้อยหนึ่งบาทถ้วน" {
		t.Errorf("Expected หนึ่งร้อยหนึ่งบาทถ้วน but got %s", result)
	}
	result = BahtText(200.0)
	if result != "สองร้อยบาทถ้วน" {
		t.Errorf("Expected สองร้อยบาทถ้วน but got %s", result)
	}
	result = BahtText(201.0)
	if result != "สองร้อยหนึ่งบาทถ้วน" {
		t.Errorf("Expected สองร้อยหนึ่งบาทถ้วน but got %s", result)
	}
}

func TestNumberThousandShouldShowPan(t *testing.T) {
	result := BahtText(1000.0)
	if result != "หนึ่งพันบาทถ้วน" {
		t.Errorf("Expected หนึ่งพันบาทถ้วน but got %s", result)
	}
	result = BahtText(2000.10)
	if result != "สองพันบาทสิบสตางค์" {
		t.Errorf("Expected สองพันบาทสิบสตางค์ but got %s", result)
	}
	result = BahtText(3211.51)
	if result != "สามพันสองร้อยสิบเอ็ดบาทห้าสิบเอ็ดสตางค์" {
		t.Errorf("Expected สามพันสองร้อยสิบเอ็ดบาทห้าสิบเอ็ดสตางค์ but got %s", result)
	}
	result = BahtText(8000.31)
	if result != "แปดพันบาทสามสิบเอ็ดสตางค์" {
		t.Errorf("Expected แปดพันบาทสามสิบเอ็ดสตางค์ but got %s", result)
	}
}

func TestNumberTenThousandShouldShowMuern(t *testing.T) {
	result := BahtText(30000.0)
	if result != "สามหมื่นบาทถ้วน" {
		t.Errorf("Expected สามหมื่นบาทถ้วน but got %s", result)
	}
	result = BahtText(98765.10)
	if result != "เก้าหมื่นแปดพันเจ็ดร้อยหกสิบห้าบาทสิบสตางค์" {
		t.Errorf("Expected เก้าหมื่นแปดพันเจ็ดร้อยหกสิบห้าบาทสิบสตางค์ but got %s", result)
	}
	result = BahtText(30211.21)
	if result != "สามหมื่นสองร้อยสิบเอ็ดบาทยี่สิบเอ็ดสตางค์" {
		t.Errorf("Expected สามหมื่นสองร้อยสิบเอ็ดบาทยี่สิบเอ็ดสตางค์ but got %s", result)
	}
}

func TestNumberHundredThousandShouldShowSaan(t *testing.T) {
	result := BahtText(800000.0)
	if result != "แปดแสนบาทถ้วน" {
		t.Errorf("Expected แปดแสนบาทถ้วน but got %s", result)
	}
	result = BahtText(258065.81)
	if result != "สองแสนห้าหมื่นแปดพันหกสิบห้าบาทแปดสิบเอ็ดสตางค์" {
		t.Errorf("Expected สองแสนห้าหมื่นแปดพันหกสิบห้าบาทแปดสิบเอ็ดสตางค์ but got %s", result)
	}
}

func TestNumberMillionShouldShowLaan(t *testing.T) {
	result := BahtText(3500000.0)
	if result != "สามล้านห้าแสนบาทถ้วน" {
		t.Errorf("Expected สามล้านห้าแสนบาทถ้วน but got %s", result)
	}
}
func TestNumberMultipleMillionShouldShowMultipleLaan(t *testing.T) {
	result := BahtText(12000000.0)
	if result != "สิบสองล้านบาทถ้วน" {
		t.Errorf("Expected สิบสองล้านบาทถ้วน but got %s", result)
	}
	result = BahtText(21000000.0)
	if result != "ยี่สิบเอ็ดล้านบาทถ้วน" {
		t.Errorf("Expected ยี่สิบเอ็ดล้านบาทถ้วน but got %s", result)
	}
	result = BahtText(51000000000000.51)
	if result != "ห้าสิบเอ็ดล้านล้านบาทห้าสิบเอ็ดสตางค์" {
		t.Errorf("Expected ห้าสิบเอ็ดล้านล้านบาทห้าสิบเอ็ดสตางค์ but got %s", result)
	}
	result = BahtText(10000000680000.51)
	if result != "สิบล้านล้านหกแสนแปดหมื่นบาทห้าสิบเอ็ดสตางค์" {
		t.Errorf("Expected สิบล้านล้านหกแสนแปดหมื่นบาทห้าสิบเอ็ดสตางค์ but got %s", result)
	}
}

func TestNegativeMinusPrefixNumberShouldPrintLoob(t *testing.T) {
	result := BahtText(-1.10)
	if result != "ลบหนึ่งบาทสิบสตางค์" {
		t.Errorf("Expected ลบหนึ่งบาทสิบสตางค์ but got %s", result)
	}
	result = BahtText(-1000.0)
	if result != "ลบหนึ่งพันบาทถ้วน" {
		t.Errorf("Expected ลบหนึ่งพันบาทถ้วน but got %s", result)
	}
	result = BahtText(-258065.81)
	if result != "ลบสองแสนห้าหมื่นแปดพันหกสิบห้าบาทแปดสิบเอ็ดสตางค์" {
		t.Errorf("Expected ลบสองแสนห้าหมื่นแปดพันหกสิบห้าบาทแปดสิบเอ็ดสตางค์ but got %s", result)
	}
	result = BahtText(-10000000680000.51)
	if result != "ลบสิบล้านล้านหกแสนแปดหมื่นบาทห้าสิบเอ็ดสตางค์" {
		t.Errorf("Expected สิบล้านล้านหกแสนแปดหมื่นบาทห้าสิบเอ็ดสตางค์ but got %s", result)
	}
}
