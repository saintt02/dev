package ids

import "strconv"

// StringToInt convierte un string a int y maneja cualquier error de conversión.
func StringToInt(id string) (int, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(idInt), nil
}