package relhandler

func ExecuteRCValidate(manifest string) Response {
	if manifest == "" {
		return Response{
			Status: "not ok",
			Error: "empty manifest",
		}
	}
	//core logic here
	return Response{
		Status: "ok",
		Message: "manifest validated successfully",
	}
}