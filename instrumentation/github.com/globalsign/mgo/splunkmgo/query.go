// Copyright Splunk Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package splunkmgo

import (
	"time"

	"github.com/globalsign/mgo"
)

// Query tracks info on the query and traces operations it performs.
type Query struct {
	*mgo.Query

	cfg *config
}

// All works like Iter.All.
func (q *Query) All(result interface{}) error {
	return q.Iter().All(result)
}

func (q *Query) Apply(change mgo.Change, result interface{}) (*mgo.ChangeInfo, error) {
	// FIXME: trace
	return q.Query.Apply(change, result)
}

func (q *Query) Batch(n int) *Query {
	return &Query{Query: q.Query.Batch(n), cfg: q.cfg.Copy()}
}

func (q *Query) Collation(collation *mgo.Collation) *Query {
	return &Query{Query: q.Query.Collation(collation), cfg: q.cfg.Copy()}
}

func (q *Query) Comment(comment string) *Query {
	return &Query{Query: q.Query.Comment(comment), cfg: q.cfg.Copy()}
}

// For iterates the job of documents coverted by the query evaluating f for
// each iteration and stopping if it returns a non-nil error.
//
// Deprecated: Use Iter instead.
func (q *Query) For(result interface{}, f func() error) error {
	return q.Iter().For(result, f)
}

func (q *Query) Hint(indexKey ...string) *Query {
	return &Query{Query: q.Query.Hint(indexKey...), cfg: q.cfg.Copy()}
}

func (q *Query) Iter() *Iter {
	return &Iter{Iter: q.Query.Iter(), cfg: q.cfg.Copy()}
}

func (q *Query) Limit(n int) *Query {
	return &Query{Query: q.Query.Limit(n), cfg: q.cfg.Copy()}
}

func (q *Query) LogReplay() *Query {
	return &Query{Query: q.Query.LogReplay(), cfg: q.cfg.Copy()}
}

// MapReduce traces and executes a map/reduce job for documents covered by the
// query. That kind of job is suitable for very flexible bulk aggregation of
// data performed at the server side via Javascript functions.
//
// Results from the job may be returned as a result of the query itself
// through the result parameter in case they'll certainly fit in memory and in
// a single document. If there's the possibility that the amount of data might
// be too large, results must be stored back in an alternative collection or
// even a separate database, by setting the Out field of the provided
// MapReduce job. In that case, provide nil as the result parameter.
func (q *Query) MapReduce(job *mgo.MapReduce, result interface{}) (*mgo.MapReduceInfo, error) {
	// FIXME: trace
	return q.Query.MapReduce(job, result)
}

func (q *Query) One(result interface{}) error {
	// FIXME: trace
	return q.Query.One(result)
}

func (q *Query) Prefetch(p float64) *Query {
	return &Query{Query: q.Query.Prefetch(p), cfg: q.cfg.Copy()}
}

func (q *Query) Select(selector interface{}) *Query {
	return &Query{Query: q.Query.Select(selector), cfg: q.cfg.Copy()}
}

func (q *Query) SetMaxScan(n int) *Query {
	return &Query{Query: q.Query.SetMaxScan(n), cfg: q.cfg.Copy()}
}

func (q *Query) SetMaxTime(d time.Duration) *Query {
	return &Query{Query: q.Query.SetMaxTime(d), cfg: q.cfg.Copy()}
}

func (q *Query) Skip(n int) *Query {
	return &Query{Query: q.Query.Skip(n), cfg: q.cfg.Copy()}
}

func (q *Query) Snapshot() *Query {
	return &Query{Query: q.Query.Snapshot(), cfg: q.cfg.Copy()}
}

func (q *Query) Sort(fields ...string) *Query {
	return &Query{Query: q.Query.Sort(fields...), cfg: q.cfg.Copy()}
}

func (q *Query) Tail(timeout time.Duration) *Iter {
	return &Iter{Iter: q.Query.Tail(timeout), cfg: q.cfg.Copy()}
}

// Iter stores informations about a Cursor and traces access to it.
type Iter struct {
	*mgo.Iter

	cfg *config
}
