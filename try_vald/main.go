//
// Copyright (C) 2019-2021 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kpango/glg"
	"google.golang.org/grpc"

	"github.com/vdaas/vald-client-go/v1/payload"
	"github.com/vdaas/vald-client-go/v1/vald"
)

type Vector struct {
	ID     string
	Vector []float32
}

func main() {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := vald.NewValdClient(conn)

	vs := []Vector{
		// {
		// 	ID:     "1",
		// 	Vector: []float32{1, 2, 3, 4, 5},
		// },
		// {
		// 	ID:     "2",
		// 	Vector: []float32{5, 4, 3, 2, 1},
		// },
		// {
		// 	ID:     "3",
		// 	Vector: []float32{0, 0, 0, 0, 0},
		// },
		// {
		// 	ID:     "4",
		// 	Vector: []float32{1, 1, 1, 1, 1},
		// },
		// {
		// 	ID:     "5",
		// 	Vector: []float32{5, 5, 5, 5, 5},
		// },
	}

	for _, v := range vs {
		_, err := client.Insert(ctx, &payload.Insert_Request{
			Vector: &payload.Object_Vector{
				Id:     v.ID,
				Vector: v.Vector,
			},
			Config: &payload.Insert_Config{
				SkipStrictExistCheck: true,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(3 * time.Second)

	// Send searching vector and configuration object to the Vald Agent server via gRPC.
	res, err := client.Search(ctx, &payload.Search_Request{
		Vector: []float32{2, 2, 3, 4, 5},
		// Conditions for hitting the search.
		Config: &payload.Search_Config{
			Num:     2,    // the number of search results
			Radius:  -1,   // Radius is used to determine the space of search candidate radius for neighborhood vectors. -1 means infinite circle.
			Epsilon: 0.01, // Epsilon is used to determines how much to expand from search candidate radius.
		},
	})
	if err != nil {
		glg.Fatal(err)
	}

	b, _ := json.MarshalIndent(res.GetResults(), "", " ")
	fmt.Printf("%v\n", string(b))
}
