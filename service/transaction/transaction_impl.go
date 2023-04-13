package transaction

import (
	"context"
	"sync"
	"time"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/dimasyudhana/alterra-group-project-2/repository/book"
	"github.com/dimasyudhana/alterra-group-project-2/repository/transaction"
	"github.com/go-playground/validator"
)

type trx struct {
	repo      transaction.TrasanctionRepoInterface
	book      book.Repository
	dep       dependecy.Depend
	validator *validator.Validate
}

func NewTrxService(repo transaction.TrasanctionRepoInterface, book book.Repository, dep dependecy.Depend) TrxServiceInterface {
	return &trx{repo: repo, dep: dep, validator: validator.New(), book: book}
}

func (t *trx) Create(ctx context.Context, reqs []int, uid int) error {
	errchan := make(chan error, len(reqs))
	wg := sync.WaitGroup{}
	wg.Add(len(reqs))
	for _, val := range reqs {
		go func(id int) {
			defer wg.Done()
			time := time.Now().Add(3 * 24 * time.Hour).UTC().Local().Format("2006-01-02 15:04:05")
			req := entities.Transaction{BorrowerId: uid, EndDate: time}
			trxid, err := t.repo.Create(t.dep.Db.WithContext(ctx), req)
			if err := t.book.UpdateByBookID(t.dep.Db.WithContext(ctx), uint(id), entities.Book{Status: "rented"}); err != nil {
				errchan <- err
				return
			}
			if err != nil {
				t.dep.Log.Errorf("Service : %v", err)
				errchan <- err
				return
			}
			if err := t.repo.InsertTrxBook(t.dep.Db.WithContext(ctx), entities.TransactionBook{TransactionId: trxid, BookId: id}); err != nil {
				errchan <- err
				return
			}
			errchan <- nil
		}(val)
	}
	wg.Wait()
	close(errchan)
	for err := range errchan {
		if err != nil {
			return err
		}
	}
	return nil
}
