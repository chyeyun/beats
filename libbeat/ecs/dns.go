// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ecs

// Fields describing DNS queries and answers.
// DNS events should either represent a single DNS query prior to getting
// answers (`dns.type:query`) or they should represent a full exchange and
// contain the query details as well as all of the answers that were provided
// for this query (`dns.type:answer`).
type Dns struct {
	// The type of DNS event captured, query or answer.
	// If your source of DNS events only gives you DNS queries, you should only
	// create dns events of type `dns.type:query`.
	// If your source of DNS events gives you answers as well, you should
	// create one event per query (optionally as soon as the query is seen).
	// And a second event containing all query details as well as an array of
	// answers.
	Type string `ecs:"type"`

	// The DNS packet identifier assigned by the program that generated the
	// query. The identifier is copied to the response.
	ID string `ecs:"id"`

	// The DNS operation code that specifies the kind of query in the message.
	// This value is set by the originator of a query and copied into the
	// response.
	OpCode string `ecs:"op_code"`

	// Array of 2 letter DNS header flags.
	// Expected values are: AA, TC, RD, RA, AD, CD, DO.
	HeaderFlags []string `ecs:"header_flags"`

	// The DNS response code.
	ResponseCode string `ecs:"response_code"`

	// The name being queried.
	// If the name field contains non-printable characters (below 32 or above
	// 126), those characters should be represented as escaped base 10 integers
	// (\DDD). Back slashes and quotes should be escaped. Tabs, carriage
	// returns, and line feeds should be converted to \t, \r, and \n
	// respectively.
	QuestionName string `ecs:"question.name"`

	// The type of record being queried.
	QuestionType string `ecs:"question.type"`

	// The class of records being queried.
	QuestionClass string `ecs:"question.class"`

	// The highest registered domain, stripped of the subdomain.
	// For example, the registered domain for "foo.example.com" is
	// "example.com".
	// This value can be determined precisely with a list like the public
	// suffix list (http://publicsuffix.org). Trying to approximate this by
	// simply taking the last two labels will not work well for TLDs such as
	// "co.uk".
	QuestionRegisteredDomain string `ecs:"question.registered_domain"`

	// The effective top level domain (eTLD), also known as the domain suffix,
	// is the last part of the domain name. For example, the top level domain
	// for example.com is "com".
	// This value can be determined precisely with a list like the public
	// suffix list (http://publicsuffix.org). Trying to approximate this by
	// simply taking the last label will not work well for effective TLDs such
	// as "co.uk".
	QuestionTopLevelDomain string `ecs:"question.top_level_domain"`

	// The subdomain is all of the labels under the registered_domain.
	// If the domain has multiple levels of subdomain, such as
	// "sub2.sub1.example.com", the subdomain field should contain "sub2.sub1",
	// with no trailing period.
	QuestionSubdomain string `ecs:"question.subdomain"`

	// An array containing an object for each answer section returned by the
	// server.
	// The main keys that should be present in these objects are defined by
	// ECS. Records that have more information may contain more keys than what
	// ECS defines.
	// Not all DNS data sources give all details about DNS answers. At minimum,
	// answer objects must contain the `data` key. If more information is
	// available, map as much of it to ECS as possible, and add any additional
	// fields to the answer objects as custom fields.
	Answers []map[string]interface{} `ecs:"answers"`

	// The domain name to which this resource record pertains.
	// If a chain of CNAME is being resolved, each answer's `name` should be
	// the one that corresponds with the answer's `data`. It should not simply
	// be the original `question.name` repeated.
	AnswersName string `ecs:"answers.name"`

	// The type of data contained in this resource record.
	AnswersType string `ecs:"answers.type"`

	// The class of DNS data contained in this resource record.
	AnswersClass string `ecs:"answers.class"`

	// The time interval in seconds that this resource record may be cached
	// before it should be discarded. Zero values mean that the data should not
	// be cached.
	AnswersTtl int64 `ecs:"answers.ttl"`

	// The data describing the resource.
	// The meaning of this data depends on the type and class of the resource
	// record.
	AnswersData string `ecs:"answers.data"`

	// Array containing all IPs seen in `answers.data`.
	// The `answers` array can be difficult to use, because of the variety of
	// data formats it can contain. Extracting all IP addresses seen in there
	// to `dns.resolved_ip` makes it possible to index them as IP addresses,
	// and makes them easier to visualize and query for.
	ResolvedIP []string `ecs:"resolved_ip"`
}
