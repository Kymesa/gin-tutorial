package res

import (
	"gorm.io/gorm"
)

// Page representa una respuesta paginada
type Page[T any] struct {
	Content    []T   `json:"content"`    // registros
	Page       int   `json:"page"`       // página actual
	PageSize   int   `json:"pageSize"`   // tamaño de página
	Total      int64 `json:"total"`      // total de registros
	TotalPages int   `json:"totalPages"` // total de páginas
	HasNext    bool  `json:"hasNext"`    // siguiente página
	HasPrev    bool  `json:"hasPrev"`    // página anterior
}

// Paginate genera la paginación para cualquier entidad
func Paginate[T any](db *gorm.DB, page, pageSize int) (Page[T], error) {
	var result Page[T]
	var total int64
	var items []T

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// contar total de registros
	db.Model(new(T)).Count(&total)

	// obtener registros
	err := db.Limit(pageSize).Offset(offset).Order("created_at DESC").Find(&items).Error

	if err != nil {
		return result, err
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	result = Page[T]{
		Content:    items,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}

	return result, nil
}
