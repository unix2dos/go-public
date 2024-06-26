package example

type IPostRepo interface{}

func NewPostRepo() (IPostRepo, func(), error) {
	return new(IPostRepo), nil, nil
}

type IPostUsecase interface{}
type postUsecase struct {
	repo IPostRepo
}

func NewPostUsecase(repo IPostRepo) (*postUsecase, func(), error) {
	return &postUsecase{repo: repo}, nil, nil
}

type PostService struct {
	usecase IPostUsecase
}

func NewPostService(u IPostUsecase) (*PostService, error) {
	return &PostService{usecase: u}, nil
}
