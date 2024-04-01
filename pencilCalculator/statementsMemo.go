package pencilCalculator

import (
	"context"
	"sync"
)

type result struct {
	value PencilResult
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func NewStatementsMemo(s map[string]*Statement) *StatementsMemo {
	return &StatementsMemo{statementMap: s, resultsCache: make(map[string]*entry)}
}

type StatementsMemo struct {
	statementMap map[string]*Statement
	mu           sync.Mutex // guards resultsCache
	resultsCache map[string]*entry
}

func (memo *StatementsMemo) Evaluate(key string, ctx context.Context) (value PencilResult, err error) {
	memo.mu.Lock()
	statement := memo.statementMap[key]
	e := memo.resultsCache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.resultsCache[key] = e
		memo.mu.Unlock()

		e.res.value = statement.Evaluate(ctx, memo)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
