package err

type BadRequest struct {
	Err string
}

func (err BadRequest) Error() string {
	return err.Err
}

func NewErr(err string) BadRequest {
	return BadRequest{err}
}
