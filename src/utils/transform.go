package utils

// SerializeModelToDTO serializes a slice of models into a slice of DTOs using a provided serialization function.
//
// This function takes a slice of models of type T and a serialization function that converts a model of type T into a DTO of type U.
// It applies the serialization function to each model in the slice and returns a new slice of DTOs.
//
// The serialization function is a closure that takes a model of type T as an argument and returns a DTO of type U.
// Generic and can be used with any type of models and DTOs.
func SerializeModelToDTO[T any, U any](models []T, serializeFunc func(T) U) []U {
	var result []U
	for _, model := range models {
		result = append(result, serializeFunc(model))
	}
	return result
}

// TransformToDTO converts a slice of items(model) to a slice of DTOs Type provided.
// It takes a slice of T and a function to convert T to DTO.
// Returns a slice of DTO objects.
func TransformToDTO[T any, DTO any](items []T, newDTOFunc func(T) DTO) []DTO {
	return SerializeModelToDTO(items, newDTOFunc)
}

func DerefPtr[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}
	var zero T
	return zero
}
