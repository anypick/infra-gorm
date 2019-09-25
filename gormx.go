package basegorm

/**
配置gorm的事务，对于gorm的事务处理有更好的方式，优雅的一批
 */

//const (
//	TX = "Tx"
//)
//
//type Runner struct {
//	Tx *gorm.DB
//}
//
//func TxRunner(fn func(runner *Runner) error) error {
//	db := GetDb()
//	tx := db.Begin()
//	runner := &Runner{Tx: tx}
//
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//		}
//	}()
//	if err := fn(runner); err != nil {
//		tx.Rollback()
//		return err
//	} else {
//		return tx.Commit().Error
//	}
//}
//
//func WithValueContext(parent context.Context, runner *Runner) context.Context {
//	return context.WithValue(parent, TX, runner)
//}
//
//
//func ExecuteContext(parent context.Context, fn func(runner *Runner) error) error {
//	if parent == nil {
//		return TxRunner(fn)
//	}
//	runner := parent.Value(TX).(*Runner)
//	return fn(runner)
//}