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

type InternalServer struct {
	Err string
}

func (err InternalServer) Error() string {
	return err.Err
}

func NewErrInter(err string) InternalServer {
	return InternalServer{err}
}
