package utils

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func FormatTime(str v1.Time) string {
	return str.Format("2006-01-02 15:04:05")
}
