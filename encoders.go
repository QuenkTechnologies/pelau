package pelau

import (
	"encoding/json"
)

//JSON creates a JSON format Encoder
func JSON(res Response) Encoder {

	return json.NewEncoder(res)

}
